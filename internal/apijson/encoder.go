package apijson

import (
	"bytes"
	"context"
	stdjson "encoding/json"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
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
	if t == reflect.TypeOf((*big.Float)(nil)) {
		return func(plan reflect.Value, state reflect.Value) ([]byte, error) {
			return []byte(plan.Interface().(*big.Float).Text('g', 10)), nil
		}
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
		attrType := reflect.TypeOf((*attr.Value)(nil)).Elem()
		if t.Implements(attrType) {
			return e.newTerraformTypeEncoder(t)
		}
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

type terraformUnwrappingFunc func(val attr.Value) (any, diag.Diagnostics)

func (e *encoder) terraformUnwrappedEncoder(underlyingType reflect.Type, unwrap terraformUnwrappingFunc) encoderFunc {
	enc := e.typeEncoder(underlyingType)
	return handleNullAndUndefined(func(plan attr.Value, state attr.Value) ([]byte, error) {
		unwrappedPlan, diags := unwrap(plan)
		if diags.HasError() {
			return nil, errorFromDiagnostics(diags)
		}
		unwrappedState, diags := unwrap(state)
		if diags.HasError() {
			return nil, errorFromDiagnostics(diags)
		}
		return enc(reflect.ValueOf(unwrappedPlan), reflect.ValueOf(unwrappedState))
	})
}

func (e *encoder) terraformUnwrappedDynamicEncoder(unwrap terraformUnwrappingFunc) encoderFunc {
	return handleNullAndUndefined(func(plan attr.Value, state attr.Value) ([]byte, error) {
		unwrappedPlan, diags := unwrap(plan)
		if diags.HasError() {
			return nil, errorFromDiagnostics(diags)
		}
		unwrappedState, diags := unwrap(state)
		if diags.HasError() {
			return nil, errorFromDiagnostics(diags)
		}
		enc := e.typeEncoder(reflect.TypeOf(unwrappedPlan))
		return enc(reflect.ValueOf(unwrappedPlan), reflect.ValueOf(unwrappedState))
	})
}

func handleNullAndUndefined(innerFunc func(attr.Value, attr.Value) ([]byte, error)) encoderFunc {
	return func(plan reflect.Value, state reflect.Value) ([]byte, error) {
		var tfPlan attr.Value
		var tfState attr.Value
		if plan.IsValid() {
			tfPlan = plan.Interface().(attr.Value)
		}
		if state.IsValid() {
			tfState = state.Interface().(attr.Value)
		}
		planNull := !plan.IsValid() || tfPlan.IsNull()
		stateNull := !state.IsValid() || tfState.IsNull()
		planUnknown := plan.IsValid() && tfPlan.IsUnknown()
		stateUnknown := state.IsValid() && tfState.IsUnknown()

		if stateNull && planNull {
			return nil, nil
		} else if planNull {
			return explicitJsonNull, nil
		} else if planUnknown {
			return nil, nil
		} else {
			if stateNull || stateUnknown {
				tfState = tfPlan
			}
			return innerFunc(tfPlan, tfState)
		}
	}
}

func (e encoder) newTerraformTypeEncoder(t reflect.Type) encoderFunc {

	if t == reflect.TypeOf(basetypes.BoolValue{}) {
		return e.terraformUnwrappedEncoder(reflect.TypeOf(true), func(value attr.Value) (any, diag.Diagnostics) {
			return value.(basetypes.BoolValue).ValueBool(), diag.Diagnostics{}
		})
	} else if t == reflect.TypeOf(basetypes.Int64Value{}) {
		return e.terraformUnwrappedEncoder(reflect.TypeOf(int64(0)), func(value attr.Value) (any, diag.Diagnostics) {
			return value.(basetypes.Int64Value).ValueInt64(), diag.Diagnostics{}
		})
	} else if t == reflect.TypeOf(basetypes.Float64Value{}) {
		return e.terraformUnwrappedEncoder(reflect.TypeOf(float64(0)), func(value attr.Value) (any, diag.Diagnostics) {
			return value.(basetypes.Float64Value).ValueFloat64(), diag.Diagnostics{}
		})
	} else if t == reflect.TypeOf(basetypes.NumberValue{}) {
		return e.terraformUnwrappedEncoder(reflect.TypeOf(big.NewFloat(0)), func(value attr.Value) (any, diag.Diagnostics) {
			return value.(basetypes.NumberValue).ValueBigFloat(), diag.Diagnostics{}
		})
	} else if t == reflect.TypeOf(basetypes.StringValue{}) {
		return e.terraformUnwrappedEncoder(reflect.TypeOf(""), func(value attr.Value) (any, diag.Diagnostics) {
			return value.(basetypes.StringValue).ValueString(), diag.Diagnostics{}
		})
	} else if t == reflect.TypeOf(timetypes.RFC3339{}) {
		return e.terraformUnwrappedEncoder(reflect.TypeOf(time.Time{}), func(value attr.Value) (any, diag.Diagnostics) {
			return value.(timetypes.RFC3339).ValueRFC3339Time()
		})
	} else if t == reflect.TypeOf(basetypes.ListValue{}) {
		return e.terraformUnwrappedDynamicEncoder(func(value attr.Value) (any, diag.Diagnostics) {
			return value.(basetypes.ListValue).Elements(), diag.Diagnostics{}
		})
	} else if t == reflect.TypeOf(basetypes.TupleValue{}) {
		return e.terraformUnwrappedDynamicEncoder(func(value attr.Value) (any, diag.Diagnostics) {
			return value.(basetypes.TupleValue).Elements(), diag.Diagnostics{}
		})
	} else if t == reflect.TypeOf(basetypes.SetValue{}) {
		return e.terraformUnwrappedDynamicEncoder(func(value attr.Value) (any, diag.Diagnostics) {
			return value.(basetypes.SetValue).Elements(), diag.Diagnostics{}
		})
	} else if t == reflect.TypeOf(basetypes.MapValue{}) {
		return e.terraformUnwrappedDynamicEncoder(func(value attr.Value) (any, diag.Diagnostics) {
			return value.(basetypes.MapValue).Elements(), diag.Diagnostics{}
		})
	} else if t == reflect.TypeOf(basetypes.ObjectValue{}) {
		return e.terraformUnwrappedDynamicEncoder(func(value attr.Value) (any, diag.Diagnostics) {
			return value.(basetypes.ObjectValue).Attributes(), diag.Diagnostics{}
		})
	} else if t == reflect.TypeOf(basetypes.DynamicValue{}) {
		return func(plan reflect.Value, state reflect.Value) ([]byte, error) {
			tfPlan := plan.Interface().(basetypes.DynamicValue)
			tfState := state.Interface().(basetypes.DynamicValue)
			planNull := tfPlan.IsNull() || tfPlan.IsUnderlyingValueNull()
			stateMissing := tfState.IsNull() || tfState.IsUnderlyingValueNull() || tfState.IsUnderlyingValueNull() || tfState.IsUnderlyingValueUnknown()
			if stateMissing && planNull {
				return nil, nil
			} else if planNull {
				return explicitJsonNull, nil
			} else if tfPlan.IsUnknown() || tfPlan.IsUnderlyingValueUnknown() {
				return nil, nil
			} else {
				if stateMissing {
					tfState = tfPlan
				}
				unwrappedPlan := tfPlan.UnderlyingValue()
				unwrappedState := tfState.UnderlyingValue()
				enc := e.typeEncoder(reflect.TypeOf(unwrappedPlan))
				return enc(reflect.ValueOf(unwrappedPlan), reflect.ValueOf(unwrappedState))
			}
		}
	} else if t.Implements(reflect.TypeOf((*customfield.NestedObjectLike)(nil)).Elem()) {
		structType := reflect.PointerTo(t.Field(0).Type)
		return e.terraformUnwrappedEncoder(structType, func(value attr.Value) (any, diag.Diagnostics) {
			return value.(customfield.NestedObjectLike).ValueAny(context.TODO())
		})
	} else if t.Implements(reflect.TypeOf((*customfield.NestedObjectListLike)(nil)).Elem()) {
		return e.terraformUnwrappedDynamicEncoder(func(value attr.Value) (any, diag.Diagnostics) {
			return value.(customfield.NestedObjectListLike).AsStructSlice(context.TODO())
		})
	} else if t.Implements(reflect.TypeOf((*customfield.ListLike)(nil)).Elem()) {
		return e.terraformUnwrappedDynamicEncoder(func(value attr.Value) (any, diag.Diagnostics) {
			return value.(customfield.ListLike).ValueAttr(context.TODO())
		})
	} else if t.Implements(reflect.TypeOf((*customfield.NestedObjectMapLike)(nil)).Elem()) {
		return e.terraformUnwrappedDynamicEncoder(func(value attr.Value) (any, diag.Diagnostics) {
			return value.(customfield.NestedObjectMapLike).AsStructMap(context.TODO())
		})
	} else if t.Implements(reflect.TypeOf((*customfield.MapLike)(nil)).Elem()) {
		return e.terraformUnwrappedDynamicEncoder(func(value attr.Value) (any, diag.Diagnostics) {
			return value.(customfield.MapLike).ValueAttr(context.TODO())
		})
	} else if t == reflect.TypeOf(jsontypes.Normalized{}) {
		return handleNullAndUndefined(func(plan attr.Value, state attr.Value) ([]byte, error) {
			return []byte(plan.(jsontypes.Normalized).ValueString()), nil
		})
	}

	return func(plan reflect.Value, state reflect.Value) (json []byte, err error) {
		return nil, fmt.Errorf("unknown type received at terraform encoder: %s", t.String())
	}
}

func (e *encoder) newStructTypeEncoder(t reflect.Type) encoderFunc {
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
	return handleNullAndUndefined(func(value attr.Value, state attr.Value) (json []byte, err error) {
		val, errs := value.(timetypes.RFC3339).ValueRFC3339Time()
		if errs != nil {
			return nil, errorFromDiagnostics(errs)
		}
		return stdjson.Marshal(val.Format(format))
	})
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
func (e *encoder) encodeMapEntries(json []byte, plan reflect.Value, state reflect.Value) ([]byte, error) {
	type mapPair struct {
		key   []byte
		plan  reflect.Value
		state reflect.Value
	}

	pairs := []mapPair{}
	keyEncoder := e.typeEncoder(plan.Type().Key())

	iter := plan.MapRange()
	for iter.Next() {
		var encodedKeyString string
		if iter.Key().Type().Kind() == reflect.String {
			encodedKeyString = iter.Key().String()
		} else {
			var err error
			encodedKeyBytes, err := keyEncoder(iter.Key(), iter.Key())
			encodedKeyString = string(encodedKeyBytes)
			if err != nil {
				return nil, err
			}
		}
		encodedKey := []byte(sjsonReplacer.Replace(encodedKeyString))
		stateValue := state.MapIndex(iter.Key())
		pairs = append(pairs, mapPair{key: encodedKey, plan: iter.Value(), state: stateValue})
	}

	// Ensure deterministic output
	sort.Slice(pairs, func(i, j int) bool {
		return bytes.Compare(pairs[i].key, pairs[j].key) < 0
	})

	elementEncoder := e.typeEncoder(plan.Type().Elem())
	for _, pair := range pairs {
		encodedValue, err := elementEncoder(pair.plan, pair.state)
		if err != nil {
			return nil, err
		}
		if encodedValue == nil {
			// encode a nil for the property rather than omitting the key entirely
			encodedValue = explicitJsonNull
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

// If we want to set a literal key value into JSON using sjson, we need to make sure it doesn't have
// special characters that sjson interprets as a path.
var sjsonReplacer *strings.Replacer = strings.NewReplacer(".", "\\.", ":", "\\:", "*", "\\*")
