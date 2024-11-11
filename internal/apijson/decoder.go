package apijson

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"sync"
	"time"
	"unsafe"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/tidwall/gjson"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
)

// decoders is a synchronized map with roughly the following type:
// map[reflect.Type]decoderFunc
var decoders sync.Map

// Unmarshal is similar to [encoding/json.Unmarshal] and parses the JSON-encoded
// data and stores it in the given pointer.
func Unmarshal(raw []byte, to any) error {
	d := &decoderBuilder{dateFormat: time.RFC3339}
	return d.unmarshal(raw, to)
}

// UnmarshalComputed is similar to [Unmarshal], but leaves non-computed
// properties (e.g. required and optional) unchanged.
func UnmarshalComputed(raw []byte, to any) error {
	d := &decoderBuilder{dateFormat: time.RFC3339, unmarshalComputedOnly: true}
	return d.unmarshal(raw, to)
}

// UnmarshalRoot is like Unmarshal, but doesn't try to call MarshalJSON on the
// root element. Useful if a struct's UnmarshalJSON is overrode to use the
// behavior of this encoder versus the standard library.
func UnmarshalRoot(raw []byte, to any) error {
	d := &decoderBuilder{dateFormat: time.RFC3339, root: true}
	return d.unmarshal(raw, to)
}

type TerraformUpdateBehavior int

const (
	// always update the property from JSON
	Always TerraformUpdateBehavior = iota

	// if the value is Null or Undefined, then update the value, otherwise skip
	IfUnset

	// always leave this property unchanged, but possibly update nested values
	OnlyNested
)

// decoderBuilder contains the 'compile-time' state of the decoder.
type decoderBuilder struct {
	// Whether or not this is the first element and called by [UnmarshalRoot], see
	// the documentation there to see why this is necessary.
	root bool
	// The dateFormat (a format string for [time.Format]) which is chosen by the
	// last struct tag that was seen.
	dateFormat string

	// Only updates computed properties on structs
	unmarshalComputedOnly bool

	// This is used to control decoding behavior for computed and computed_optional
	// fields.
	updateBehavior TerraformUpdateBehavior
}

// decoderState contains the 'run-time' state of the decoder.
type decoderState struct {
	strict    bool
	exactness exactness
}

// Exactness refers to how close to the type the result was if deserialization
// was successful. This is useful in deserializing unions, where you want to try
// each entry, first with strict, then with looser validation, without actually
// having to do a lot of redundant work by marshalling twice (or maybe even more
// times).
type exactness int8

const (
	// Some values had to fudged a bit, for example by converting a string to an
	// int, or an enum with extra values.
	loose exactness = iota
	// There are some extra arguments, but other wise it matches the union.
	extras
	// Exactly right.
	exact
)

type decoderFunc func(node gjson.Result, value reflect.Value, state *decoderState) error

type decoderField struct {
	tag    parsedStructTag
	fn     decoderFunc
	idx    []int
	goname string
}

type decoderEntry struct {
	reflect.Type
	dateFormat     string
	root           bool
	tfSkipBehavior TerraformUpdateBehavior
}

func (d *decoderBuilder) unmarshal(raw []byte, to any) error {
	value := reflect.ValueOf(to).Elem()
	result := gjson.ParseBytes(raw)
	if !value.IsValid() {
		return fmt.Errorf("apijson: cannot marshal into invalid value")
	}
	return d.typeDecoder(value.Type())(result, value, &decoderState{strict: false, exactness: exact})
}

func (d *decoderBuilder) typeDecoder(t reflect.Type) decoderFunc {
	entry := decoderEntry{
		Type:           t,
		dateFormat:     d.dateFormat,
		root:           d.root,
		tfSkipBehavior: d.updateBehavior,
	}

	if fi, ok := decoders.Load(entry); ok {
		return fi.(decoderFunc)
	}

	// To deal with recursive types, populate the map with an
	// indirect func before we build it. This type waits on the
	// real func (f) to be ready and then calls it. This indirect
	// func is only used for recursive types.
	var (
		wg sync.WaitGroup
		f  decoderFunc
	)
	wg.Add(1)
	fi, loaded := decoders.LoadOrStore(entry, decoderFunc(func(node gjson.Result, v reflect.Value, state *decoderState) error {
		wg.Wait()
		return f(node, v, state)
	}))
	if loaded {
		return fi.(decoderFunc)
	}

	// Compute the real decoder and replace the indirect func with it.
	f = d.newTypeDecoder(t)
	wg.Done()
	decoders.Store(entry, f)
	return f
}

func unmarshalerDecoder(n gjson.Result, v reflect.Value, state *decoderState) error {
	return v.Interface().(json.Unmarshaler).UnmarshalJSON([]byte(n.Raw))
}

func (d *decoderBuilder) newTypeDecoder(t reflect.Type) decoderFunc {
	b := d.updateBehavior

	if t.ConvertibleTo(reflect.TypeOf(time.Time{})) {
		return d.newTimeTypeDecoder(t)
	}
	if t != reflect.TypeOf(jsontypes.Normalized{}) && t.ConvertibleTo(reflect.TypeOf(timetypes.RFC3339{})) {
		return d.newCustomTimeTypeDecoder(t)
	}
	if !d.root && t.Implements(reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()) {
		return unmarshalerDecoder
	}
	if t == reflect.TypeOf((*big.Float)(nil)).Elem() {
		return d.newBigFloatDecoder(t)
	}
	d.root = false

	if _, ok := unionRegistry[t]; ok {
		return d.newUnionDecoder(t)
	}

	switch t.Kind() {
	case reflect.Pointer:
		inner := t.Elem()
		innerDecoder := d.typeDecoder(inner)

		return func(n gjson.Result, v reflect.Value, state *decoderState) error {
			if !v.IsValid() {
				return fmt.Errorf("apijson: unexpected invalid reflection value %+#v", v)
			}

			if (v.IsNil() && b == OnlyNested) || (!v.IsNil() && b == IfUnset) || (v.IsNil() && n.Type == gjson.Null) {
				return nil
			}

			newValue := reflect.New(inner).Elem()
			if !v.IsNil() {
				newValue.Set(v.Elem())
			}
			err := innerDecoder(n, newValue, state)
			if err != nil {
				return err
			}

			v.Set(newValue.Addr())
			return nil
		}
	case reflect.Struct:
		return d.newStructTypeDecoder(t)
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		return d.newArrayTypeDecoder(t)
	case reflect.Map:
		return d.newMapDecoder(t)
	case reflect.Interface:
		return func(node gjson.Result, value reflect.Value, state *decoderState) error {
			if !value.IsValid() {
				return fmt.Errorf("apijson: unexpected invalid value %+#v", value)
			}
			if node.Value() != nil && value.CanSet() {
				value.Set(reflect.ValueOf(node.Value()))
			}
			return nil
		}
	default:
		return d.newPrimitiveTypeDecoder(t)
	}
}

// newUnionDecoder returns a decoderFunc that deserializes into a union using an
// algorithm roughly similar to Pydantic's [smart algorithm].
//
// Conceptually this is equivalent to choosing the best schema based on how 'exact'
// the deserialization is for each of the schemas.
//
// If there is a tie in the level of exactness, then the tie is broken
// left-to-right.
//
// [smart algorithm]: https://docs.pydantic.dev/latest/concepts/unions/#smart-mode
func (d *decoderBuilder) newUnionDecoder(t reflect.Type) decoderFunc {
	unionEntry, ok := unionRegistry[t]
	if !ok {
		panic("apijson: couldn't find union of type " + t.String() + " in union registry")
	}
	decoders := []decoderFunc{}
	for _, variant := range unionEntry.variants {
		decoder := d.typeDecoder(variant.Type)
		decoders = append(decoders, decoder)
	}
	return func(n gjson.Result, v reflect.Value, state *decoderState) error {
		// Set bestExactness to worse than loose
		bestExactness := loose - 1

		for idx, variant := range unionEntry.variants {
			decoder := decoders[idx]
			if variant.TypeFilter != n.Type {
				continue
			}
			if len(unionEntry.discriminatorKey) != 0 && n.Get(unionEntry.discriminatorKey).Value() != variant.DiscriminatorValue {
				continue
			}
			sub := decoderState{strict: state.strict, exactness: exact}
			inner := reflect.New(variant.Type).Elem()
			err := decoder(n, inner, &sub)
			if err != nil {
				continue
			}
			if sub.exactness == exact {
				v.Set(inner)
				return nil
			}
			if sub.exactness > bestExactness {
				v.Set(inner)
				bestExactness = sub.exactness
			}
		}

		if bestExactness < loose {
			return errors.New("apijson: was not able to coerce type as union")
		}

		if guardStrict(state, bestExactness != exact) {
			return errors.New("apijson: was not able to coerce type as union strictly")
		}

		return nil
	}
}

func (d *decoderBuilder) newBigFloatDecoder(_ reflect.Type) decoderFunc {
	return func(node gjson.Result, value reflect.Value, state *decoderState) error {
		f, _, err := big.ParseFloat(node.Raw, 10, 0, big.ToNearestEven)
		if err != nil {
			return fmt.Errorf("apijson: failed to parse big.Float: %v", err)
		}
		value.Set(reflect.ValueOf(f))
		return nil
	}
}

func (d *decoderBuilder) newMapDecoder(t reflect.Type) decoderFunc {
	updateBehavior := d.updateBehavior

	keyType := t.Key()
	itemType := t.Elem()
	itemDecoder := d.typeDecoder(itemType)

	return func(node gjson.Result, value reflect.Value, state *decoderState) (err error) {
		mapValue := reflect.MakeMapWithSize(t, len(node.Map()))

		extraKeys := map[reflect.Value]bool{}
		var nonEmpty bool

		if updateBehavior == OnlyNested {
			// populate existing values regardless of whether they are coming from the API
			for _, key := range value.MapKeys() {
				nonEmpty = true
				extraKeys[key] = true
				item := value.MapIndex(key)
				mapValue.SetMapIndex(key, item)
			}
		}

		node.ForEach(func(key, jsonValue gjson.Result) bool {
			// It's fine for us to just use `ValueOf` here because the key types will
			// always be primitive types so we don't need to decode it using the standard pattern
			keyValue := reflect.ValueOf(key.Value())
			if !keyValue.IsValid() {
				if err == nil {
					err = fmt.Errorf("apijson: received invalid key type %v", keyValue.String())
				}
				return false
			}
			if keyValue.Type() != keyType {
				if err == nil {
					err = fmt.Errorf("apijson: expected key type %v but got %v", keyType, keyValue.Type())
				}
				return false
			}

			if updateBehavior == OnlyNested && !extraKeys[keyValue] {
				// skip keys that aren't already in the map
				return true
			}

			existingValue := value.MapIndex(keyValue)
			itemValue := reflect.New(itemType).Elem()
			if existingValue.IsValid() {
				itemValue.Set(existingValue)
			}
			itemerr := itemDecoder(jsonValue, itemValue, state)
			if itemerr != nil {
				if err == nil {
					err = itemerr
				}
				return false
			}

			mapValue.SetMapIndex(keyValue, itemValue)
			extraKeys[keyValue] = false
			nonEmpty = true
			return true
		})

		if err != nil {
			return err
		}

		// set additional keys not present in JSON to a zero value (or null)
		for key, exists := range extraKeys {
			if !exists {
				continue
			}
			existingValue := value.MapIndex(key)
			itemValue := reflect.New(itemType).Elem()
			itemValue.Set(existingValue)
			itemerr := itemDecoder(gjson.Result{}, itemValue, state)
			if itemerr != nil {
				return itemerr
			}

			mapValue.SetMapIndex(key, itemValue)
		}
		if nonEmpty {
			value.Set(mapValue)
		}

		return nil
	}
}

func (d *decoderBuilder) newArrayTypeDecoder(t reflect.Type) decoderFunc {
	updateBehavior := d.updateBehavior
	itemDecoder := d.typeDecoder(t.Elem())

	return func(node gjson.Result, value reflect.Value, state *decoderState) (err error) {

		if node.Type == gjson.Null {
			if updateBehavior == Always {
				value.Set(reflect.Zero(t))
				return nil
			}
		}

		if node.Type != gjson.Null && !node.IsArray() {
			return fmt.Errorf("apijson: could not deserialize to an array")
		}

		arrayNode := node.Array()

		existingLen := value.Len()
		var numItems int

		// if we are only updating nested values, we won't change the length of the array
		if updateBehavior == OnlyNested || updateBehavior == IfUnset {
			numItems = existingLen
		} else {
			numItems = len(arrayNode)
		}

		// populate the array with the existing values
		arrayValue := reflect.MakeSlice(reflect.SliceOf(t.Elem()), numItems, numItems)
		for i := 0; i < existingLen && i < numItems; i++ {
			arrayValue.Index(i).Set(value.Index(i))
		}

		for i, itemNode := range arrayNode {
			if i >= numItems {
				break
			}
			err = itemDecoder(itemNode, arrayValue.Index(i), state)
			if err != nil {
				return err
			}
		}

		// set any additional values not in the JSON to a zero value (or null)
		for i := len(arrayNode); i < numItems; i++ {
			err = itemDecoder(gjson.Result{}, arrayValue.Index(i), state)
			if err != nil {
				return
			}
		}

		value.Set(arrayValue)
		return nil
	}
}

func (d *decoderBuilder) decodeTerraformPrimitive(nullValue func() any, decodeNonNull decoderFunc) decoderFunc {
	updateBehavior := d.updateBehavior
	return func(node gjson.Result, value reflect.Value, state *decoderState) error {
		var isNullOrUnknown bool
		attr, ok := value.Interface().(attr.Value)
		if !ok || (attr.IsNull() || attr.IsUnknown()) {
			isNullOrUnknown = true
		}

		if updateBehavior == IfUnset && !isNullOrUnknown {
			return nil
		}

		if node.Type == gjson.Null && (updateBehavior == Always || isNullOrUnknown) {
			value.Set(reflect.ValueOf(nullValue()))
			return nil
		}

		if updateBehavior == OnlyNested {
			return nil
		}

		return decodeNonNull(node, value, state)
	}
}

func (d *decoderBuilder) newTerraformTypeDecoder(t reflect.Type) decoderFunc {
	ctx := context.TODO()

	b := d.updateBehavior

	if (t == reflect.TypeOf(basetypes.StringValue{})) {
		return d.decodeTerraformPrimitive(func() any { return types.StringNull() }, func(node gjson.Result, value reflect.Value, state *decoderState) error {
			if node.Type == gjson.String {
				value.Set(reflect.ValueOf(types.StringValue(node.String())))
				return nil
			}
			return fmt.Errorf("apijson: cannot deserialize types.StringValue")
		})
	}

	if (t == reflect.TypeOf(basetypes.Int64Value{})) {
		return d.decodeTerraformPrimitive(func() any { return types.Int64Null() }, func(node gjson.Result, value reflect.Value, state *decoderState) error {
			// use ParseFloat just to validate that it's a valid number
			_, err := strconv.ParseFloat(node.Str, 64)
			if node.Type == gjson.JSON || (node.Type == gjson.String && err != nil) {
				return fmt.Errorf("apijson: failed to parse types.Int64Value")
			}
			value.Set(reflect.ValueOf(types.Int64Value(node.Int())))
			return nil
		})
	}

	if (t == reflect.TypeOf(basetypes.Float64Value{})) {
		return d.decodeTerraformPrimitive(func() any { return types.Float64Null() }, func(node gjson.Result, value reflect.Value, state *decoderState) error {
			// use ParseFloat just to validate that it's a valid number
			_, err := strconv.ParseFloat(node.Str, 64)
			if node.Type == gjson.JSON || (node.Type == gjson.String && err != nil) {
				return fmt.Errorf("apijson: failed to parse types.Float64Value")
			}
			value.Set(reflect.ValueOf(types.Float64Value(node.Float())))
			return nil
		})
	}

	if (t == reflect.TypeOf(basetypes.NumberValue{})) {
		return d.decodeTerraformPrimitive(func() any { return types.NumberNull() }, func(node gjson.Result, value reflect.Value, state *decoderState) error {
			value.Set(reflect.ValueOf(types.NumberValue(big.NewFloat(node.Float()))))
			_, err := strconv.ParseFloat(node.Str, 64)
			if node.Type == gjson.JSON || (node.Type == gjson.String && err != nil) {
				return fmt.Errorf("apijson: failed to parse types.Float64Value")
			}
			return nil
		})
	}

	if (t == reflect.TypeOf(basetypes.BoolValue{})) {
		return d.decodeTerraformPrimitive(func() any { return types.BoolNull() }, func(node gjson.Result, value reflect.Value, state *decoderState) error {
			if node.Type == gjson.True || node.Type == gjson.False {
				value.Set(reflect.ValueOf(types.BoolValue(node.Bool())))
				return nil
			}
			return fmt.Errorf("cannot deserialize bool")
		})
	}

	if (t == reflect.TypeOf(basetypes.ListValue{})) {
		return func(node gjson.Result, value reflect.Value, state *decoderState) error {
			if !shouldUpdateNested(value, b) {
				return nil
			}
			eleType := value.Interface().(basetypes.ListValue).ElementType(ctx)
			if b == Always && node.Type == gjson.Null {
				value.Set(reflect.ValueOf(types.ListNull(eleType)))
				return nil
			}
			if node.Type == gjson.JSON {
				attr, err := d.inferTerraformAttrFromValue(node)
				if err != nil {
					return err
				}
				value.Set(reflect.ValueOf(attr))
				return nil
			}
			return fmt.Errorf("apijson: cannot deserialize unexpected type %s to types.ListValue", node.Type)
		}
	}

	if (t == reflect.TypeOf(basetypes.TupleValue{})) {
		return func(node gjson.Result, value reflect.Value, state *decoderState) error {
			if !shouldUpdateNested(value, b) {
				return nil
			}
			tuple := value.Interface().(basetypes.TupleValue)
			elementTypes := tuple.ElementTypes(ctx)
			if node.Type == gjson.Null {
				value.Set(reflect.ValueOf(types.TupleNull(elementTypes)))
				return nil
			} else if node.Type == gjson.JSON && node.IsArray() {
				nodes := node.Array()
				elements := make([]attr.Value, len(elementTypes))
				for i, elementType := range elementTypes {
					elements[i] = elementType.ValueType(ctx)
					element := &elements[i]
					if i >= len(nodes) {
						continue
					}
					decoder := d.newTerraformTypeDecoder(reflect.TypeOf(*element))
					err := decoder(nodes[i], reflect.ValueOf(element).Elem(), state)
					if err != nil {
						return err
					}
				}
				value.Set(reflect.ValueOf(types.TupleValueMust(elementTypes, elements)))
				return nil
			} else {
				return fmt.Errorf("apijson: cannot deserialize unexpected type %s to types.TupleValue", node.Type)
			}
		}
	}

	if (t == reflect.TypeOf(basetypes.SetValue{})) {
		return func(node gjson.Result, value reflect.Value, state *decoderState) error {
			if !shouldUpdateNested(value, b) {
				return nil
			}
			eleType := value.Interface().(basetypes.SetValue).ElementType(ctx)
			switch node.Type {
			case gjson.Null:
				value.Set(reflect.ValueOf(types.ListNull(eleType)))
				return nil
			case gjson.JSON:
				elementType, attributes, err := d.parseArrayOfValues(node)
				if err != nil {
					return err
				}
				setValue, diags := basetypes.NewSetValue(elementType, attributes)
				if diags.HasError() {
					return errorFromDiagnostics(diags)
				}
				value.Set(reflect.ValueOf(setValue))
				return nil
			default:
				return fmt.Errorf("apijson: cannot deserialize unexpected type %s to types.ListValue", node.Type)
			}
		}
	}

	if (t == reflect.TypeOf(basetypes.MapValue{})) {
		return func(node gjson.Result, value reflect.Value, state *decoderState) error {
			if !shouldUpdateNested(value, b) {
				return nil
			}
			eleType := value.Interface().(basetypes.MapValue).ElementType(ctx)
			switch node.Type {
			case gjson.Null:
				if b == Always {
					value.Set(reflect.ValueOf(types.ListNull(eleType)))
				}
				return nil
			case gjson.JSON:
				attributes := map[string]attr.Value{}
				loopErr := error(nil)
				node.ForEach(func(key, value gjson.Result) bool {
					attr, err := d.inferTerraformAttrFromValue(value)
					if err != nil {
						loopErr = err
						return false
					}
					attributes[key.String()] = attr
					return true
				})
				if loopErr != nil {
					return loopErr
				}
				mapValue, diags := basetypes.NewMapValue(eleType, attributes)
				if diags.HasError() {
					return errorFromDiagnostics(diags)
				}
				value.Set(reflect.ValueOf(mapValue))
				return nil
			default:
				return fmt.Errorf("apijson: cannot deserialize unexpected type %s to types.MapValue", node.Type)
			}
		}
	}

	if (t == reflect.TypeOf(basetypes.ObjectValue{})) {
		return func(node gjson.Result, value reflect.Value, state *decoderState) error {
			if !shouldUpdateNested(value, b) {
				return nil
			}
			objValue := value.Interface().(basetypes.ObjectValue)
			attrTypes := objValue.AttributeTypes(ctx)
			switch node.Type {
			case gjson.Null:
				if b == Always {
					value.Set(reflect.ValueOf(types.ObjectNull(attrTypes)))
				}
				return nil
			case gjson.JSON:
				if len(attrTypes) > 0 {
					attributes := objValue.Attributes()
					newAttributes := map[string]attr.Value{}
					for key, attrType := range attrTypes {
						value := attributes[key]
						jsonValue := node.Get(key)
						newValue := attrType.ValueType(ctx)
						if value == nil {
							value = newValue
						}
						dec := d.typeDecoder(reflect.TypeOf(value))
						err := dec(jsonValue, reflect.ValueOf(&newValue).Elem(), state)
						if err != nil {
							return err
						}
						newAttributes[key] = newValue
					}
					newObject, diags := basetypes.NewObjectValue(attrTypes, newAttributes)
					if diags.HasError() {
						return errorFromDiagnostics(diags)
					}
					value.Set(reflect.ValueOf(newObject))
					return nil
				} else {
					attr, err := d.inferTerraformAttrFromValue(node)
					if err != nil {
						return err
					}
					value.Set(reflect.ValueOf(attr))
					return nil
				}
			default:
				return fmt.Errorf("apijson: cannot deserialize unexpected type to types.ObjectValue")
			}
		}
	}

	if t.Implements(reflect.TypeOf((*customfield.NestedObjectLike)(nil)).Elem()) {
		structType := t.Field(0).Type
		decoderFunc := d.newStructTypeDecoder(structType)
		return func(node gjson.Result, value reflect.Value, state *decoderState) error {
			if !shouldUpdateNested(value, b) {
				return nil
			}
			objectValue := value.Interface().(customfield.NestedObjectLike)
			if node.Type == gjson.Null {
				if b == Always || objectValue.IsNull() || objectValue.IsUnknown() {
					nullValue := objectValue.NullValue(ctx)
					value.Set(reflect.ValueOf(nullValue))
					return nil
				}
			}

			structValue := reflect.New(structType)
			if !objectValue.IsNull() && !objectValue.IsUnknown() {
				objPtr, _ := objectValue.ValueAny(ctx)
				structValue = reflect.ValueOf(objPtr)
			}
			err := decoderFunc(node, structValue.Elem(), state)
			if err != nil {
				return err
			}
			updated := objectValue.KnownValue(ctx, structValue.Interface())
			value.Set(reflect.ValueOf(updated))
			return nil
		}
	}

	if t.Implements(reflect.TypeOf((*customfield.NestedObjectListLike)(nil)).Elem()) {
		structType := t.Field(0).Type
		structSliceType := reflect.SliceOf(structType)
		dec := d.typeDecoder(structSliceType)
		return func(node gjson.Result, value reflect.Value, state *decoderState) error {
			if !shouldUpdateNested(value, b) {
				return nil
			}
			existingObjectListValue := value.Interface().(customfield.NestedObjectListLike)
			if node.Type == gjson.Null {
				if b == Always {
					nullValue := existingObjectListValue.NullValue(ctx)
					value.Set(reflect.ValueOf(nullValue))
				}
			}

			newObjectListValue := reflect.New(structSliceType).Elem()
			existingAny, _ := existingObjectListValue.AsStructSlice(ctx)
			newObjectListValue.Set(reflect.ValueOf(existingAny))
			err := dec(node, newObjectListValue, state)
			if err != nil {
				return err
			}
			structInterface := newObjectListValue.Interface()
			updated := existingObjectListValue.KnownValue(ctx, structInterface)
			value.Set(reflect.ValueOf(updated))
			return nil
		}
	}

	if t.Implements(reflect.TypeOf((*customfield.ListLike)(nil)).Elem()) {
		structType := t.Field(0).Type
		sliceOfStruct := reflect.SliceOf(structType)
		dec := d.typeDecoder(sliceOfStruct)
		return func(node gjson.Result, value reflect.Value, state *decoderState) error {
			if !shouldUpdateNested(value, b) {
				return nil
			}
			objectListValue := value.Interface().(customfield.ListLike)
			if node.Type == gjson.Null {
				if b == Always || objectListValue.IsNullOrUnknown() {
					nullValue := objectListValue.NullValue(ctx)
					value.Set(reflect.ValueOf(nullValue))
					return nil
				}
			}
			lv, _ := objectListValue.ValueAttr(ctx)
			val := reflect.New(sliceOfStruct).Elem()
			val.Set(reflect.ValueOf(lv))
			err := dec(node, val, state)
			if err != nil {
				return err
			}
			newObjectList := objectListValue.KnownValue(ctx, val.Interface())
			value.Set(reflect.ValueOf(newObjectList))
			return nil
		}
	}

	if t.Implements(reflect.TypeOf((*customfield.NestedObjectMapLike)(nil)).Elem()) {
		structType := t.Field(0).Type
		structMapType := reflect.MapOf(reflect.TypeOf(""), structType)
		dec := d.typeDecoder(structMapType)
		return func(node gjson.Result, value reflect.Value, state *decoderState) error {
			if !shouldUpdateNested(value, b) {
				return nil
			}
			existingObjectMapValue := value.Interface().(customfield.NestedObjectMapLike)
			if node.Type == gjson.Null {
				if b == Always || existingObjectMapValue.IsNull() || existingObjectMapValue.IsUnknown() {
					nullValue := existingObjectMapValue.NullValue(ctx)
					value.Set(reflect.ValueOf(nullValue))
					return nil
				}
			}

			newObjectMapValue := reflect.New(structMapType).Elem()
			existingAny, _ := existingObjectMapValue.AsStructMap(ctx)
			newObjectMapValue.Set(reflect.ValueOf(existingAny))
			err := dec(node, newObjectMapValue, state)
			if err != nil {
				return err
			}
			structInterface := newObjectMapValue.Interface()
			updated := existingObjectMapValue.KnownValue(ctx, structInterface)
			value.Set(reflect.ValueOf(updated))
			return nil
		}
	}

	if t.Implements(reflect.TypeOf((*customfield.MapLike)(nil)).Elem()) {
		structType := t.Field(0).Type
		mapOfStruct := reflect.MapOf(reflect.TypeOf(""), structType)
		dec := d.typeDecoder(mapOfStruct)
		return func(node gjson.Result, value reflect.Value, state *decoderState) error {
			if !shouldUpdateNested(value, b) {
				return nil
			}
			objectMapValue := value.Interface().(customfield.MapLike)
			if node.Type == gjson.Null {
				if b == Always || objectMapValue.IsNull() || objectMapValue.IsUnknown() {
					nullValue := objectMapValue.NullValue(ctx)
					value.Set(reflect.ValueOf(nullValue))
					return nil
				}
			}
			mv, _ := objectMapValue.ValueAttr(ctx)
			val := reflect.New(mapOfStruct).Elem()
			val.Set(reflect.ValueOf(mv))
			err := dec(node, val, state)
			if err != nil {
				return err
			}
			newObjectMap := objectMapValue.KnownValue(ctx, val.Interface())
			value.Set(reflect.ValueOf(newObjectMap))
			return nil
		}
	}

	if (t == reflect.TypeOf(basetypes.DynamicValue{})) {
		return func(node gjson.Result, value reflect.Value, state *decoderState) error {
			if !shouldUpdatePrimitive(value, b) {
				return nil
			}
			dynamic := value.Interface().(basetypes.DynamicValue)
			underlying := dynamic.UnderlyingValue()
			if !shouldUpdatePrimitive(reflect.ValueOf(underlying), b) {
				return nil
			}
			if node.Type == gjson.Null && underlying == nil {
				// special case of null means we don't have an underlying type
				value.Set(reflect.ValueOf(types.DynamicNull()))
				return nil
			}
			if underlying != nil {
				underlyingValue := reflect.New(reflect.TypeOf(underlying)).Elem()
				underlyingValue.Set(reflect.ValueOf(underlying)) // set any existing type information
				// if we have an underlying value, we can use that type to decode
				dec := d.newTerraformTypeDecoder(reflect.TypeOf(underlying))
				err := dec(node, underlyingValue, state)
				if err != nil {
					return err
				}
				value.Set(reflect.ValueOf(types.DynamicValue(underlyingValue.Interface().(attr.Value))))
			} else {
				// just decode from the json itself
				attr, err := d.inferTerraformAttrFromValue(node)
				if err != nil {
					return err
				}

				value.Set(reflect.ValueOf(types.DynamicValue(attr)))
			}
			return nil
		}
	}

	if (t == reflect.TypeOf(jsontypes.Normalized{})) {
		return d.decodeTerraformPrimitive(func() any { return jsontypes.NewNormalizedNull() }, func(node gjson.Result, value reflect.Value, state *decoderState) error {
			raw := ""
			switch node.Type {
			case gjson.Number:
				fallthrough
			case gjson.String:
				raw = node.Raw
			default:
				raw = node.String()
			}
			value.Set(reflect.ValueOf(jsontypes.NewNormalizedValue(raw)))
			return nil
		})
	}

	return func(node gjson.Result, value reflect.Value, state *decoderState) error {
		return fmt.Errorf("apijson: cannot deserialize terraform type %v", t)
	}
}

func shouldUpdateNested(value reflect.Value, behavior TerraformUpdateBehavior) bool {
	switch behavior {
	case IfUnset:
		// even if the value is set, nested properties may be
		// unset, so we still update recursively.
		return true
	case OnlyNested:
		attr, ok := value.Interface().(attr.Value)
		if ok && attr.IsNull() {
			return false
		}
		// if the value is not null, we may have nested
		// properties to update.
		return true
	case Always:
		return true
	default:
		return true
	}
}

func shouldUpdatePrimitive(value reflect.Value, behavior TerraformUpdateBehavior) bool {
	switch behavior {
	case IfUnset:
		attr, ok := value.Interface().(attr.Value)
		if ok && (attr.IsNull() || attr.IsUnknown()) {
			return true
		}
		return false
	case OnlyNested:
		return false
	case Always:
		return true
	default:
		return true
	}
}

func (d *decoderBuilder) newStructTypeDecoder(t reflect.Type) decoderFunc {
	// map of json field name to struct field decoders
	decoderFields := map[string]decoderField{}
	extraDecoder := (*decoderField)(nil)
	inlineDecoder := (*decoderField)(nil)

	if t.Implements(reflect.TypeOf((*attr.Value)(nil)).Elem()) {
		return d.newTerraformTypeDecoder(t)
	}

	// This helper allows us to recursively collect field encoders into a flat
	// array. The parameter `index` keeps track of the access patterns necessary
	// to get to some field.
	var collectFieldDecoders func(r reflect.Type, index []int)
	collectFieldDecoders = func(r reflect.Type, index []int) {
		for i := 0; i < r.NumField(); i++ {
			idx := append(index, i)
			field := t.FieldByIndex(idx)
			if !field.IsExported() {
				continue
			}
			// If this is an embedded struct, traverse one level deeper to extract
			// the fields and get their encoders as well.
			if field.Anonymous {
				collectFieldDecoders(field.Type, idx)
				continue
			}
			// If json tag is not present, then we skip, which is intentionally
			// different behavior from the stdlib.
			ptag, ok := parseJSONStructTag(field)
			if !ok {
				continue
			}

			// sets the appropriate unmarshal behavior if we are only un-marshaling
			// computed properties.
			if d.unmarshalComputedOnly {
				// always skip non-computed fields, but update nested fields
				// if the value is not null
				if ptag.required || ptag.optional {
					d.updateBehavior = OnlyNested
				} else if ptag.computed_optional {
					// skip computed_optional fields only if they are non-null
					d.updateBehavior = IfUnset
				} else {
					// if the value is computed, we always update it.
					// note this is also set for untagged properties, so the default
					// is to update.
					d.updateBehavior = Always
				}
			}

			// We only want to support unexported fields if they're tagged with
			// `extras` because that field shouldn't be part of the public API. We
			// also want to only keep the top level extras
			if ptag.extras && len(index) == 0 {
				extraDecoder = &decoderField{ptag, d.typeDecoder(field.Type.Elem()), idx, field.Name}
				continue
			}
			if ptag.inline && len(index) == 0 {
				inlineDecoder = &decoderField{ptag, d.typeDecoder(field.Type), idx, field.Name}
				continue
			}
			if ptag.metadata {
				continue
			}

			oldFormat := d.dateFormat
			dateFormat, ok := parseFormatStructTag(field)
			if ok {
				switch dateFormat {
				case "date-time":
					d.dateFormat = time.RFC3339
				case "date":
					d.dateFormat = "2006-01-02"
				}
			}
			decoderFields[ptag.name] = decoderField{ptag, d.typeDecoder(field.Type), idx, field.Name}
			d.dateFormat = oldFormat
			d.updateBehavior = Always // reset the flag
		}
	}
	collectFieldDecoders(t, []int{})

	return func(node gjson.Result, value reflect.Value, state *decoderState) (err error) {
		if field := value.FieldByName("JSON"); field.IsValid() {
			if raw := field.FieldByName("raw"); raw.IsValid() {
				setUnexportedField(raw, node.Raw)
			}
		}

		if inlineDecoder != nil {
			dest := value.FieldByIndex(inlineDecoder.idx)

			if dest.IsValid() && node.Type != gjson.Null {
				err = inlineDecoder.fn(node, dest, state)
			}

			return err
		}

		typedExtraType := reflect.Type(nil)
		typedExtraFields := reflect.Value{}
		if extraDecoder != nil {
			typedExtraType = value.FieldByIndex(extraDecoder.idx).Type()
			typedExtraFields = reflect.MakeMap(typedExtraType)
		}

		nodeMap := node.Map()

		for fieldName, itemNode := range nodeMap {
			df, explicit := decoderFields[fieldName]
			var (
				dest reflect.Value
				fn   decoderFunc
			)
			if explicit {
				fn = df.fn
				dest = value.FieldByIndex(df.idx)
			}
			if !explicit && extraDecoder != nil {
				dest = reflect.New(typedExtraType.Elem()).Elem()
				fn = extraDecoder.fn
			}

			if dest.IsValid() {
				_ = fn(itemNode, dest, state)
			}

			if !explicit && extraDecoder != nil {
				typedExtraFields.SetMapIndex(reflect.ValueOf(fieldName), dest)
			}
		}

		// Handle struct fields that are not present in the JSON
		// this is in case they should be initialized to a "null" value
		// that is different from the zero value
		for fieldName, df := range decoderFields {
			_, existsInJson := nodeMap[fieldName]
			if existsInJson {
				continue
			}
			fn := df.fn
			dest := value.FieldByIndex(df.idx)

			// note that we don't include pointers to structs, because
			// that could be recursive and would cause an infinite loop.
			// if dest.IsValid() && dest.Kind() == reflect.Struct {
			if dest.IsValid() {
				_ = fn(gjson.Result{}, dest, state)
			}
		}

		if extraDecoder != nil && typedExtraFields.Len() > 0 {
			value.FieldByIndex(extraDecoder.idx).Set(typedExtraFields)
		}

		return nil
	}
}

func (d *decoderBuilder) newPrimitiveTypeDecoder(t reflect.Type) decoderFunc {
	switch t.Kind() {
	case reflect.String:
		return func(n gjson.Result, v reflect.Value, state *decoderState) error {
			v.SetString(n.String())
			if guardStrict(state, n.Type != gjson.String) {
				return fmt.Errorf("apijson: failed to parse string strictly")
			}
			// Everything that is not an object can be loosely stringified.
			if n.Type == gjson.JSON {
				return fmt.Errorf("apijson: failed to parse string")
			}
			if guardUnknown(state, v) {
				return fmt.Errorf("apijson: failed string enum validation")
			}
			return nil
		}
	case reflect.Bool:
		return func(n gjson.Result, v reflect.Value, state *decoderState) error {
			v.SetBool(n.Bool())
			if guardStrict(state, n.Type != gjson.True && n.Type != gjson.False) {
				return fmt.Errorf("apijson: failed to parse bool strictly")
			}
			// Numbers and strings that are either 'true' or 'false' can be loosely
			// deserialized as bool.
			if n.Type == gjson.String && (n.Raw != "true" && n.Raw != "false") || n.Type == gjson.JSON {
				return fmt.Errorf("apijson: failed to parse bool")
			}
			if guardUnknown(state, v) {
				return fmt.Errorf("apijson: failed bool enum validation")
			}
			return nil
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return func(n gjson.Result, v reflect.Value, state *decoderState) error {
			v.SetInt(n.Int())
			if guardStrict(state, n.Type != gjson.Number || n.Num != float64(int(n.Num))) {
				return fmt.Errorf("apijson: failed to parse int strictly")
			}
			// Numbers, booleans, and strings that maybe look like numbers can be
			// loosely deserialized as numbers.
			if n.Type == gjson.JSON || (n.Type == gjson.String && !canParseAsNumber(n.Str)) {
				return fmt.Errorf("apijson: failed to parse int")
			}
			if guardUnknown(state, v) {
				return fmt.Errorf("apijson: failed int enum validation")
			}
			return nil
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return func(n gjson.Result, v reflect.Value, state *decoderState) error {
			v.SetUint(n.Uint())
			if guardStrict(state, n.Type != gjson.Number || n.Num != float64(int(n.Num)) || n.Num < 0) {
				return fmt.Errorf("apijson: failed to parse uint strictly")
			}
			// Numbers, booleans, and strings that maybe look like numbers can be
			// loosely deserialized as uint.
			if n.Type == gjson.JSON || (n.Type == gjson.String && !canParseAsNumber(n.Str)) {
				return fmt.Errorf("apijson: failed to parse uint")
			}
			if guardUnknown(state, v) {
				return fmt.Errorf("apijson: failed uint enum validation")
			}
			return nil
		}
	case reflect.Float32, reflect.Float64:
		return func(n gjson.Result, v reflect.Value, state *decoderState) error {
			v.SetFloat(n.Float())
			if guardStrict(state, n.Type != gjson.Number) {
				return fmt.Errorf("apijson: failed to parse float strictly")
			}
			// Numbers, booleans, and strings that maybe look like numbers can be
			// loosely deserialized as floats.
			if n.Type == gjson.JSON || (n.Type == gjson.String && !canParseAsNumber(n.Str)) {
				return fmt.Errorf("apijson: failed to parse float")
			}
			if guardUnknown(state, v) {
				return fmt.Errorf("apijson: failed float enum validation")
			}
			return nil
		}
	default:
		return func(node gjson.Result, v reflect.Value, state *decoderState) error {
			return fmt.Errorf("unknown type received at primitive decoder: %s", t.String())
		}
	}
}

func decodeTime(format string, value string, state *decoderState) (*time.Time, error) {
	parsed, err := time.Parse(format, value)
	if err == nil {
		return &parsed, nil
	}

	if guardStrict(state, true) {
		return nil, err
	}

	layouts := []string{
		"2006-01-02",
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05Z0700",
		"2006-01-02T15:04:05",
		"2006-01-02 15:04:05Z07:00",
		"2006-01-02 15:04:05Z0700",
		"2006-01-02 15:04:05",
	}

	for _, layout := range layouts {
		parsed, err := time.Parse(layout, value)
		if err == nil {
			return &parsed, nil
		}
	}

	return nil, fmt.Errorf("unable to leniently parse date-time string: %s", value)
}

func (d *decoderBuilder) newTimeTypeDecoder(t reflect.Type) decoderFunc {
	format := d.dateFormat
	return func(n gjson.Result, v reflect.Value, state *decoderState) error {
		parsed, err := decodeTime(format, n.Str, state)
		if err == nil {
			v.Set(reflect.ValueOf(*parsed).Convert(t))
		}
		return err
	}
}

func (d *decoderBuilder) newCustomTimeTypeDecoder(t reflect.Type) decoderFunc {
	b := d.updateBehavior
	format := d.dateFormat
	return func(n gjson.Result, v reflect.Value, state *decoderState) error {
		if !shouldUpdatePrimitive(v, b) {
			return nil
		}
		if n.Type == gjson.Null {
			v.Set(reflect.ValueOf(timetypes.NewRFC3339Null()))
			return nil
		}
		parsed, err := decodeTime(format, n.Str, state)
		if err == nil {
			val := timetypes.NewRFC3339TimePointerValue(parsed)
			v.Set(reflect.ValueOf(val).Convert(t))
		}
		return err
	}
}

func setUnexportedField(field reflect.Value, value interface{}) {
	reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Set(reflect.ValueOf(value))
}

func guardStrict(state *decoderState, cond bool) bool {
	if !cond {
		return false
	}

	if state.strict {
		return true
	}

	state.exactness = loose
	return false
}

func canParseAsNumber(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}

func guardUnknown(state *decoderState, v reflect.Value) bool {
	if have, ok := v.Interface().(interface{ IsKnown() bool }); guardStrict(state, ok && !have.IsKnown()) {
		return true
	}
	return false
}

func (d *decoderBuilder) inferTerraformAttrFromValue(node gjson.Result) (attr.Value, error) {
	ctx := context.TODO()
	switch node.Type {
	case gjson.Null:
		return types.DynamicNull(), nil
	case gjson.True:
		return types.BoolValue(true), nil
	case gjson.False:
		return types.BoolValue(false), nil
	case gjson.Number:
		_, err := strconv.ParseInt(node.String(), 10, 64)
		if err == nil {
			return types.Int64Value(node.Int()), nil
		}
		return types.Float64Value(node.Float()), nil
	case gjson.String:
		return types.StringValue(node.String()), nil
	case gjson.JSON:
		if node.IsArray() {
			elementType, attributes, err := d.parseArrayOfValues(node)
			if err != nil {
				return nil, err
			}
			newVal, diags := basetypes.NewListValue(elementType, attributes)
			if diags.HasError() {
				return nil, errorFromDiagnostics(diags)
			}
			return newVal, nil
		} else if node.IsObject() {
			attributes := map[string]attr.Value{}
			attributeTypes := map[string]attr.Type{}
			loopErr := error(nil)
			node.ForEach(func(key, value gjson.Result) bool {
				attr, err := d.inferTerraformAttrFromValue(value)
				if err != nil {
					loopErr = err
					return false
				}
				attributes[key.String()] = attr
				attributeTypes[key.String()] = attr.Type(ctx)
				return true
			})
			if loopErr != nil {
				return nil, loopErr
			}
			newVal, diags := basetypes.NewObjectValue(attributeTypes, attributes)
			if diags.HasError() {
				return nil, errorFromDiagnostics(diags)
			}
			return newVal, nil
		}

	}
	return nil, fmt.Errorf("apijson: cannot infer terraform attribute from value")
}

func (d *decoderBuilder) parseArrayOfValues(node gjson.Result) (attr.Type, []attr.Value, error) {
	ctx := context.TODO()
	loopErr := error(nil)
	attributes := []attr.Value{}
	var elementType attr.Type
	node.ForEach(func(_, value gjson.Result) bool {
		val, err := d.inferTerraformAttrFromValue(value)
		if err != nil {
			loopErr = err
			return false
		}
		elementType = val.Type(ctx)
		attributes = append(attributes, val)
		return true
	})
	if loopErr != nil {
		return nil, nil, loopErr
	}
	return elementType, attributes, nil
}
