package apijson

import (
	"bytes"
	"context"
	stdjson "encoding/json"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/tidwall/sjson"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
)

var explicitJsonNull = []byte("null")
var encoders sync.Map // map[encoderEntry]encoderFunc

// Marshals the given data to a JSON string.
// For null values, omits the property entirely.
func Marshal(value interface{}) ([]byte, error) {
	e := &encoder{dateFormat: time.RFC3339}
	return e.marshal(value, value)
}

// Marshals the given plan data to a JSON string.
// For null values, omits the property unless the corresponding state value was set.
func MarshalForUpdate(plan interface{}, state interface{}) ([]byte, error) {
	e := &encoder{dateFormat: time.RFC3339}
	return e.marshal(plan, state)
}

func MarshalRoot(value interface{}) ([]byte, error) {
	e := &encoder{root: true, dateFormat: time.RFC3339}
	return e.marshal(value, value)
}

type encoder struct {
	dateFormat string
	root       bool
}

type encoderFunc func(plan reflect.Value, state reflect.Value) ([]byte, error)

type encoderField struct {
	tag parsedStructTag
	fn  encoderFunc
	idx []int
}

type encoderEntry struct {
	reflect.Type
	dateFormat string
	root       bool
}

func errorFromDiagnostics(diags diag.Diagnostics) error {
	if diags == nil {
		return nil
	}
	messages := []string{}
	for _, err := range diags {
		messages = append(messages, err.Summary())
		messages = append(messages, err.Detail())
	}
	return errors.New(strings.Join(messages, " "))
}

func (e *encoder) marshal(plan interface{}, state interface{}) ([]byte, error) {
	planVal := reflect.ValueOf(plan)
	stateVal := reflect.ValueOf(state)
	if !planVal.IsValid() {
		return nil, nil
	}
	if !stateVal.IsValid() {
		return nil, nil
	}
	typ := planVal.Type()
	enc := e.typeEncoder(typ)
	return enc(planVal, stateVal)
}

func (e *encoder) typeEncoder(t reflect.Type) encoderFunc {
	entry := encoderEntry{
		Type:       t,
		dateFormat: e.dateFormat,
		root:       e.root,
	}

	if fi, ok := encoders.Load(entry); ok {
		return fi.(encoderFunc)
	}

	// To deal with recursive types, populate the map with an
	// indirect func before we build it. This type waits on the
	// real func (f) to be ready and then calls it. This indirect
	// func is only used for recursive types.
	var (
		wg sync.WaitGroup
		f  encoderFunc
	)
	wg.Add(1)
	fi, loaded := encoders.LoadOrStore(entry, encoderFunc(func(state reflect.Value, plan reflect.Value) ([]byte, error) {
		wg.Wait()
		return f(state, plan)
	}))
	if loaded {
		return fi.(encoderFunc)
	}

	// Compute the real encoder and replace the indirect func with it.
	f = e.newTypeEncoder(t)
	wg.Done()
	encoders.Store(entry, f)
	return f
}

func (e *encoder) newTypeEncoder(t reflect.Type) encoderFunc {
	if t.ConvertibleTo(reflect.TypeOf(time.Time{})) {
		return e.newTimeTypeEncoder()
	}
	if t != reflect.TypeOf(jsontypes.Normalized{}) && t.ConvertibleTo(reflect.TypeOf(timetypes.RFC3339{})) {
		return e.newCustomTimeTypeEncoder()
	}
	// if !e.root && t.Implements(reflect.TypeOf((*json.Marshaler)(nil)).Elem()) {
	// 	return marshalerEncoder
	// }
	e.root = false
	switch t.Kind() {
	case reflect.Pointer:
		inner := t.Elem()

		innerEncoder := e.typeEncoder(inner)
		return func(p reflect.Value, s reflect.Value) ([]byte, error) {
			// if state and plan are both nil or invalid, then don't marshal the field
			if !s.IsValid() || !p.IsValid() || (s.IsNil() && p.IsNil()) {
				return nil, nil
			}
			// if plan is nil but state isn't, then marshal the field as an explicit null
			if !s.IsNil() && p.IsNil() {
				return explicitJsonNull, nil
			}
			// if state is nil, then there is no value to unset. we still have to pass
			// some value in for state, so we pass in the plan value so it marshals as-is
			if s.IsNil() {
				s = p
			}
			return innerEncoder(p.Elem(), s.Elem())
		}
	case reflect.Struct:
		return e.newStructTypeEncoder(t)
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		return e.newArrayTypeEncoder(t)
	case reflect.Map:
		return e.newMapEncoder(t)
	case reflect.Interface:
		return e.newInterfaceEncoder()
	default:
		return e.newPrimitiveTypeEncoder(t)
	}
}

func (e *encoder) newPrimitiveTypeEncoder(t reflect.Type) encoderFunc {
	switch t.Kind() {
	// Note that we could use `gjson` to encode these types but it would complicate our
	// code more and this current code shouldn't cause any issues
	case reflect.String:
		return func(p reflect.Value, s reflect.Value) ([]byte, error) {
			return stdjson.Marshal(p.String())
		}
	case reflect.Bool:
		return func(p reflect.Value, s reflect.Value) ([]byte, error) {
			if p.Bool() {
				return []byte("true"), nil
			}
			return []byte("false"), nil
		}
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		return func(p reflect.Value, s reflect.Value) ([]byte, error) {
			return []byte(strconv.FormatInt(p.Int(), 10)), nil
		}
	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return func(p reflect.Value, s reflect.Value) ([]byte, error) {
			return []byte(strconv.FormatUint(p.Uint(), 10)), nil
		}
	case reflect.Float32:
		return func(p reflect.Value, s reflect.Value) ([]byte, error) {
			return []byte(strconv.FormatFloat(p.Float(), 'f', -1, 32)), nil
		}
	case reflect.Float64:
		return func(p reflect.Value, s reflect.Value) ([]byte, error) {
			return []byte(strconv.FormatFloat(p.Float(), 'f', -1, 64)), nil
		}
	default:
		return func(p reflect.Value, s reflect.Value) ([]byte, error) {
			return nil, fmt.Errorf("unknown type received at primitive encoder: %s", t.String())
		}
	}
}

func (e *encoder) newArrayTypeEncoder(t reflect.Type) encoderFunc {
	itemEncoder := e.typeEncoder(t.Elem())

	return func(plan reflect.Value, state reflect.Value) ([]byte, error) {
		if state.IsNil() && plan.IsNil() {
			return nil, nil
		} else if plan.IsNil() {
			return explicitJsonNull, nil
		}

		json := []byte("[]")
		for i := 0; i < plan.Len(); i++ {
			planItem := plan.Index(i)

			var value, err = itemEncoder(planItem, planItem)
			if err != nil {
				return nil, err
			}
			if value == nil {
				// Assume that empty items should be inserted as `null` so that the output array
				// will be the same length as the input array
				value = explicitJsonNull
			}

			json, err = sjson.SetRawBytes(json, "-1", value)
			if err != nil {
				return nil, err
			}
		}

		return json, nil
	}
}

func (e *encoder) newStructTypeEncoder(t reflect.Type) encoderFunc {

	if (t == reflect.TypeOf(basetypes.StringValue{})) {
		return func(plan reflect.Value, state reflect.Value) (json []byte, err error) {
			var tfPlan = plan.Interface().(basetypes.StringValue)
			var tfState = state.Interface().(basetypes.StringValue)
			if tfState.IsNull() && tfPlan.IsNull() {
				return nil, nil
			} else if tfPlan.IsNull() {
				return explicitJsonNull, nil
			} else if tfPlan.IsUnknown() {
				return nil, nil
			} else {
				return stdjson.Marshal(tfPlan.ValueString())
			}
		}
	}

	if (t == reflect.TypeOf(basetypes.Int64Value{})) {
		return func(plan reflect.Value, state reflect.Value) (json []byte, err error) {
			var tfPlan = plan.Interface().(basetypes.Int64Value)
			var tfState = state.Interface().(basetypes.Int64Value)
			if tfState.IsNull() && tfPlan.IsNull() {
				return nil, nil
			} else if tfPlan.IsNull() {
				return explicitJsonNull, nil
			} else if tfPlan.IsUnknown() {
				return nil, nil
			} else {
				return []byte(fmt.Sprint(tfPlan.ValueInt64())), nil
			}
		}
	}

	if (t == reflect.TypeOf(basetypes.NumberValue{})) {
		return func(plan reflect.Value, state reflect.Value) (json []byte, err error) {
			var tfPlan = plan.Interface().(basetypes.NumberValue)
			var tfState = state.Interface().(basetypes.NumberValue)
			if tfState.IsNull() && tfPlan.IsNull() {
				return nil, nil
			} else if tfPlan.IsNull() {
				return explicitJsonNull, nil
			} else if tfPlan.IsUnknown() {
				return nil, nil
			} else {
				return []byte(fmt.Sprint(tfPlan.ValueBigFloat().Float64())), nil
			}
		}
	}

	if (t == reflect.TypeOf(basetypes.Float64Value{})) {
		return func(plan reflect.Value, state reflect.Value) (json []byte, err error) {
			var tfPlan = plan.Interface().(basetypes.Float64Value)
			var tfState = state.Interface().(basetypes.Float64Value)
			if tfState.IsNull() && tfPlan.IsNull() {
				return nil, nil
			} else if tfPlan.IsNull() {
				return explicitJsonNull, nil
			} else if tfPlan.IsUnknown() {
				return nil, nil
			} else {
				return []byte(fmt.Sprint(tfPlan.ValueFloat64())), nil
			}
		}
	}

	if (t == reflect.TypeOf(basetypes.BoolValue{})) {
		return func(plan reflect.Value, state reflect.Value) (json []byte, err error) {
			var tfPlan = plan.Interface().(basetypes.BoolValue)
			var tfState = state.Interface().(basetypes.BoolValue)
			if tfState.IsNull() && tfPlan.IsNull() {
				return nil, nil
			} else if tfPlan.IsNull() {
				return explicitJsonNull, nil
			} else if tfPlan.IsUnknown() {
				return nil, nil
			} else {
				return []byte(fmt.Sprint(tfPlan.ValueBool())), nil
			}
		}
	}

	if t.Implements(reflect.TypeOf((*customfield.NestedObjectLike)(nil)).Elem()) {
		structType := t.Field(0).Type
		return func(plan reflect.Value, state reflect.Value) (json []byte, err error) {
			var tfPlan = plan.Interface().(basetypes.ObjectValuable)
			var tfState = state.Interface().(basetypes.ObjectValuable)
			if tfState.IsNull() && tfPlan.IsNull() {
				return nil, nil
			} else if tfPlan.IsNull() {
				return explicitJsonNull, nil
			} else if tfPlan.IsUnknown() {
				return nil, nil
			} else {
				if tfState.IsNull() || tfState.IsUnknown() {
					state = plan
				}
				planStruct, _ := plan.Interface().(customfield.NestedObjectLike).ValueAny(context.TODO())
				stateStruct, _ := state.Interface().(customfield.NestedObjectLike).ValueAny(context.TODO())
				return e.newStructTypeEncoder(structType)(reflect.ValueOf(planStruct).Elem(), reflect.ValueOf(stateStruct).Elem())
			}
		}
	}

	if (t == reflect.TypeOf(basetypes.DynamicValue{})) {
		return func(plan reflect.Value, state reflect.Value) (json []byte, err error) {
			var tfPlan = plan.Interface().(basetypes.DynamicValue)
			var tfState = state.Interface().(basetypes.DynamicValue)

			planValue := tfPlan.UnderlyingValue()
			if tfPlan.IsUnderlyingValueNull() || tfPlan.IsUnderlyingValueUnknown() {
				planValue = nil
			}

			stateValue := tfState.UnderlyingValue()
			if tfState.IsUnderlyingValueNull() || tfState.IsUnderlyingValueUnknown() {
				stateValue = nil
			}

			// if the plan is set to a value, use it
			if planValue != nil {
				valueType := planValue.Type(context.TODO()).ValueType(context.TODO())
				if stateValue == nil {
					// state must be non-nil, so set it to the plan if it is
					// because the plan is set, the state will be ignored anyway
					stateValue = planValue
				}
				return e.newStructTypeEncoder(reflect.TypeOf(valueType))(reflect.ValueOf(planValue), reflect.ValueOf(stateValue))
			}

			// if state is set to a value, and the plan is null, explicitly unset it
			if stateValue != nil && (tfPlan.IsNull() || tfPlan.IsUnderlyingValueNull()) {
				return []byte("null"), nil
			}

			// otherwise, omit the field
			return nil, nil
		}
	}

	if (t == reflect.TypeOf(jsontypes.Normalized{})) {
		return func(plan reflect.Value, state reflect.Value) (json []byte, err error) {
			var tfPlan = plan.Interface().(jsontypes.Normalized)
			var tfState = state.Interface().(jsontypes.Normalized)
			if tfState.IsNull() && tfPlan.IsNull() {
				return nil, nil
			} else if tfPlan.IsNull() {
				return explicitJsonNull, nil
			} else if tfPlan.IsUnknown() {
				return nil, nil
			} else {
				return []byte(tfPlan.ValueString()), nil
			}
		}
	}

	encoderFields := []encoderField{}
	extraEncoder := (*encoderField)(nil)

	// This helper allows us to recursively collect field encoders into a flat
	// array. The parameter `index` keeps track of the access patterns necessary
	// to get to some field.
	var collectEncoderFields func(r reflect.Type, index []int)
	collectEncoderFields = func(r reflect.Type, index []int) {
		for i := 0; i < r.NumField(); i++ {
			idx := append(index, i)
			field := t.FieldByIndex(idx)
			if !field.IsExported() {
				continue
			}
			// If this is an embedded struct, traverse one level deeper to extract
			// the field and get their encoders as well.
			if field.Anonymous {
				collectEncoderFields(field.Type, idx)
				continue
			}
			// If json tag is not present, then we skip, which is intentionally
			// different behavior from the stdlib.
			ptag, ok := parseJSONStructTag(field)
			if !ok {
				continue
			}
			// We only want to support unexported field if they're tagged with
			// `extras` because that field shouldn't be part of the public API. We
			// also want to only keep the top level extras
			if ptag.extras && len(index) == 0 {
				extraEncoder = &encoderField{ptag, e.typeEncoder(field.Type.Elem()), idx}
				continue
			}
			if ptag.name == "-" {
				continue
			}
			// Computed fields come from the server
			if ptag.computed {
				continue
			}

			dateFormat, ok := parseFormatStructTag(field)
			oldFormat := e.dateFormat
			if ok {
				switch dateFormat {
				case "date-time":
					e.dateFormat = time.RFC3339
				case "date":
					e.dateFormat = "2006-01-02"
				}
			}
			encoderFields = append(encoderFields, encoderField{ptag, e.typeEncoder(field.Type), idx})
			e.dateFormat = oldFormat
		}
	}
	collectEncoderFields(t, []int{})

	// Ensure deterministic output by sorting by lexicographic order
	sort.Slice(encoderFields, func(i, j int) bool {
		return encoderFields[i].tag.name < encoderFields[j].tag.name
	})

	return func(plan reflect.Value, state reflect.Value) (json []byte, err error) {
		json = []byte("{}")

		for _, ef := range encoderFields {
			planField := plan.FieldByIndex(ef.idx)
			stateField, err := state.FieldByIndexErr(ef.idx)
			if err != nil {
				stateField = planField
			}
			encoded, err := ef.fn(planField, stateField)
			if err != nil {
				return nil, err
			}
			if encoded == nil {
				continue
			}
			json, err = sjson.SetRawBytes(json, ef.tag.name, encoded)
			if err != nil {
				return nil, err
			}
		}

		if extraEncoder != nil {
			json, err = e.encodeMapEntries(json, plan.FieldByIndex(extraEncoder.idx), state.FieldByIndex(extraEncoder.idx))
			if err != nil {
				return nil, err
			}
		}
		return
	}
}

func (e *encoder) newTimeTypeEncoder() encoderFunc {
	format := e.dateFormat
	return func(value reflect.Value, state reflect.Value) (json []byte, err error) {
		return []byte(`"` + value.Convert(reflect.TypeOf(time.Time{})).Interface().(time.Time).Format(format) + `"`), nil
	}
}

func (e *encoder) newCustomTimeTypeEncoder() encoderFunc {
	format := e.dateFormat
	return func(value reflect.Value, state reflect.Value) (json []byte, err error) {
		val, errs := value.Convert(reflect.TypeOf(timetypes.RFC3339{})).Interface().(timetypes.RFC3339).ValueRFC3339Time()
		if errs != nil {
			return nil, errorFromDiagnostics(errs)
		}
		return stdjson.Marshal(val.Format(format))
	}
}

func (e encoder) newInterfaceEncoder() encoderFunc {
	return func(plan reflect.Value, state reflect.Value) ([]byte, error) {
		plan = plan.Elem()
		state = state.Elem()
		if !plan.IsValid() {
			return nil, nil
		}
		if !state.IsValid() {
			return nil, nil
		}
		return e.typeEncoder(plan.Type())(plan, state)
	}
}

// Given a []byte of json (may either be an empty object or an object that already contains entries)
// encode all of the entries in the map to the json byte array.
func (e *encoder) encodeMapEntries(json []byte, p reflect.Value, s reflect.Value) ([]byte, error) {
	type mapPair struct {
		key   []byte
		plan  reflect.Value
		state reflect.Value
	}

	pairs := []mapPair{}
	keyEncoder := e.typeEncoder(p.Type().Key())

	iter := p.MapRange()
	sIter := s.MapRange()
	for iter.Next() {
		sIter.Next()
		var encodedKey []byte
		if iter.Key().Type().Kind() == reflect.String {
			encodedKey = []byte(iter.Key().String())
		} else {
			var err error
			encodedKey, err = keyEncoder(iter.Key(), sIter.Key())
			if err != nil {
				return nil, err
			}
		}
		pairs = append(pairs, mapPair{key: encodedKey, plan: iter.Value(), state: sIter.Value()})
	}

	// Ensure deterministic output
	sort.Slice(pairs, func(i, j int) bool {
		return bytes.Compare(pairs[i].key, pairs[j].key) < 0
	})

	elementEncoder := e.typeEncoder(p.Type().Elem())
	for _, pair := range pairs {
		encodedValue, err := elementEncoder(pair.plan, pair.state)
		if err != nil {
			return nil, err
		}
		json, err = sjson.SetRawBytes(json, string(pair.key), encodedValue)
		if err != nil {
			return nil, err
		}
	}

	return json, nil
}

func (e *encoder) newMapEncoder(_ reflect.Type) encoderFunc {
	return func(plan reflect.Value, state reflect.Value) ([]byte, error) {
		if state.IsNil() && plan.IsNil() {
			return nil, nil
		} else if plan.IsNil() {
			return explicitJsonNull, nil
		}

		json := []byte("{}")
		var err error
		json, err = e.encodeMapEntries(json, plan, state)
		if err != nil {
			return nil, err
		}
		return json, nil
	}
}
