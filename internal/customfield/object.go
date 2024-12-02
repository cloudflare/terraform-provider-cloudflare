package customfield

import (
	"context"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var (
	_ basetypes.ObjectTypable  = (*NestedObjectType[struct{}])(nil)
	_ basetypes.ObjectValuable = (*NestedObject[struct{}])(nil)
)

// NestedObjectType represents a basetypes.ObjectType that is defined by a struct.
type NestedObjectType[T any] struct {
	basetypes.ObjectType
}

func newNestedObjectType[T any](ctx context.Context) (NestedObjectType[T], diag.Diagnostics) {
	var diags diag.Diagnostics
	t := NestedObjectType[T]{}

	m, d := StructToAttributes[T](ctx)
	diags.Append(d...)
	if diags.HasError() {
		return t, diags
	}

	t.ObjectType = basetypes.ObjectType{AttrTypes: m}
	return t, diags
}

func NewNestedObjectType[T any](ctx context.Context) NestedObjectType[T] {
	t, err := newNestedObjectType[T](ctx)
	if err.HasError() {
		panic(fmt.Errorf("unexpected error creating NestedObjectType: %v", err))
	}
	return t
}

func (t NestedObjectType[T]) Equal(o attr.Type) bool {
	other, ok := o.(NestedObjectType[T])
	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t NestedObjectType[T]) String() string {
	var zero T
	return fmt.Sprintf("ObjectType[%T]", zero)
}

func (t NestedObjectType[T]) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	if in.IsNull() {
		return NullObject[T](ctx), diags
	}
	if in.IsUnknown() {
		return UnknownObject[T](ctx), diags
	}

	m, d := StructToAttributes[T](ctx)
	diags.Append(d...)
	if diags.HasError() {
		return UnknownObject[T](ctx), diags
	}

	v, d := basetypes.NewObjectValue(m, in.Attributes())
	diags.Append(d...)
	if diags.HasError() {
		return UnknownObject[T](ctx), diags
	}

	value := NestedObject[T]{
		ObjectValue: v,
	}

	return value, diags
}

func (t NestedObjectType[T]) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.ObjectType.ValueFromTerraform(ctx, in)

	if err != nil {
		return nil, err
	}

	objectValue, ok := attrValue.(basetypes.ObjectValue)
	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	objectValuable, diags := t.ValueFromObject(ctx, objectValue)
	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting ObjectValue to ObjectValuable: %v", diags)
	}

	return objectValuable, nil
}

func (t NestedObjectType[T]) ValueType(ctx context.Context) attr.Value {
	return UnknownObject[T](ctx)
}

func (t NestedObjectType[T]) NullValue(ctx context.Context) (attr.Value, diag.Diagnostics) {
	var diags diag.Diagnostics
	return NullObject[T](ctx), diags
}

type NestedObjectLike interface {
	basetypes.ObjectValuable
	ValueAny(ctx context.Context) (any, diag.Diagnostics)
	NullValue(ctx context.Context) NestedObjectLike
	UnknownValue(ctx context.Context) NestedObjectLike
	KnownValue(ctx context.Context, t any) NestedObjectLike
}

var _ NestedObjectLike = (*NestedObject[struct{}])(nil)

// NestedObject represents a basetypes.ObjectValue that is defined by a struct.
type NestedObject[T any] struct {
	//lint:ignore U1000 the placeholder is for easy reflection-based-access
	placeholder T
	basetypes.ObjectValue
}

func (v NestedObject[T]) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	// if the struct is initialized as the zero value, initialize the object type
	// before returning.
	if v.ObjectValue.AttributeTypes(ctx) == nil {
		v.ObjectValue = NullObject[T](ctx).ObjectValue
	}
	return v.ObjectValue.ToObjectValue(ctx)
}

func (v NestedObject[T]) NullValue(ctx context.Context) NestedObjectLike {
	return NullObject[T](ctx)
}

func (v NestedObject[T]) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	tv := v.ObjectValue
	if v.ObjectValue.Equal(basetypes.ObjectValue{}) {
		tv = NullObject[T](ctx).ObjectValue
	}
	return tv.ToTerraformValue(ctx)
}

func (v NestedObject[T]) UnknownValue(ctx context.Context) NestedObjectLike {
	return UnknownObject[T](ctx)
}

func (v NestedObject[T]) KnownValue(ctx context.Context, t any) NestedObjectLike {
	r, _ := NewObject(ctx, t.(*T))
	return r
}

func (v NestedObject[T]) Equal(o attr.Value) bool {
	other, ok := o.(NestedObject[T])
	if !ok {
		return false
	}

	return v.ObjectValue.Equal(other.ObjectValue)
}

func (v NestedObject[T]) Type(ctx context.Context) attr.Type {
	return NewNestedObjectType[T](ctx)
}

func (v NestedObject[T]) ValueAny(ctx context.Context) (any, diag.Diagnostics) {
	return v.Value(ctx)
}

func (v NestedObject[T]) Value(ctx context.Context) (*T, diag.Diagnostics) {
	var diags diag.Diagnostics

	ptr := new(T)

	diags.Append(v.ObjectValue.As(ctx, ptr, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil, diags
	}

	return ptr, diags
}

func NullObject[T any](ctx context.Context) NestedObject[T] {
	t, err := StructToAttributes[T](ctx)
	if err.HasError() {
		panic(fmt.Errorf("unexpected error creating NullObject: %v", err))
	}
	return NestedObject[T]{ObjectValue: basetypes.NewObjectNull(t)}
}

func UnknownObject[T any](ctx context.Context) NestedObject[T] {
	t, err := StructToAttributes[T](ctx)
	if err.HasError() {
		panic(fmt.Errorf("unexpected error creating UnknownObject: %v", err))
	}
	return NestedObject[T]{ObjectValue: basetypes.NewObjectUnknown(t)}
}

func NewObject[T any](ctx context.Context, t *T) (NestedObject[T], diag.Diagnostics) {
	var diags diag.Diagnostics

	m, d := StructToAttributes[T](ctx)
	diags.Append(d...)
	if diags.HasError() {
		return UnknownObject[T](ctx), diags
	}

	v, d := basetypes.NewObjectValueFrom(ctx, m, t)
	diags.Append(d...)
	if diags.HasError() {
		return UnknownObject[T](ctx), diags
	}

	return NestedObject[T]{ObjectValue: v}, diags
}

func NewObjectMust[T any](ctx context.Context, t *T) NestedObject[T] {
	o, diags := NewObject(ctx, t)
	if diags.HasError() {
		panic(fmt.Errorf("unexpected error creating NestedObject: %v", diags))
	}
	return o
}

func StructToAttributes[T any](ctx context.Context) (map[string]attr.Type, diag.Diagnostics) {
	var t T
	val := reflect.ValueOf(t)
	return StructFromAttributesGeneric(ctx, val)
}

func StructFromAttributesGeneric(ctx context.Context, val reflect.Value) (map[string]attr.Type, diag.Diagnostics) {
	var diags diag.Diagnostics
	typ := val.Type()

	if typ.Kind() == reflect.Ptr && typ.Elem().Kind() == reflect.Struct {
		val = reflect.New(typ.Elem()).Elem()
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		diags.Append(diag.NewErrorDiagnostic("Invalid type", fmt.Sprintf("%T has unsupported type: %s", val.Interface(), typ)))
		return nil, diags
	}

	attributeTypes := make(map[string]attr.Type)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if field.PkgPath != "" {
			continue
		}
		tag := field.Tag.Get(`tfsdk`)
		if tag == "-" {
			continue
		}
		if tag == "" {
			diags.Append(diag.NewErrorDiagnostic("Invalid type", fmt.Sprintf(`%T is missing a "tfsdk" struct tag on %s`, val.Interface(), field.Name)))
			return nil, diags
		}

		v := val.Field(i)

		attr, ok := v.Interface().(attr.Value)
		if ok {
			attributeTypes[tag] = attr.Type(ctx)
		} else {
			ty, diags := structFromValue(ctx, v)
			diags.Append(diags...)
			if diags.HasError() {
				return nil, diags
			}
			attributeTypes[tag] = ty
		}
	}
	return attributeTypes, nil
}

func structFromValue(ctx context.Context, v reflect.Value) (attr.Type, diag.Diagnostics) {
	var elemType attr.Type
	var diags diag.Diagnostics
	attr, ok := v.Interface().(attr.Value)
	if ok {
		return attr.Type(ctx), nil
	}

	t := v.Type()
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() == reflect.Struct {
		m, d := StructFromAttributesGeneric(ctx, v)
		diags.Append(d...)
		if diags.HasError() {
			return nil, diags
		}
		return basetypes.ObjectType{AttrTypes: m}, nil
	} else if t.Kind() == reflect.Slice {
		sliceType := t.Elem()
		if sliceType.Kind() == reflect.Ptr {
			sliceType = sliceType.Elem()
		}
		elemType, diags = structFromValue(ctx, reflect.New(sliceType).Elem())
		if diags.HasError() {
			return nil, diags
		}
		return basetypes.ListType{ElemType: elemType}, nil
	} else if t.Kind() == reflect.Map {
		keyType := t.Key()
		if keyType.Kind() != reflect.String {
			diags.Append(diag.NewErrorDiagnostic("Invalid type", fmt.Sprintf(`%T has unsupported key type: %s`, t, keyType.Kind())))
			return nil, diags
		}
		valueType := t.Elem()
		if valueType.Kind() == reflect.Ptr {
			valueType = valueType.Elem()
		}
		elemType, diags = structFromValue(ctx, reflect.New(valueType).Elem())
		if diags.HasError() {
			return nil, diags
		}
		return basetypes.MapType{ElemType: elemType}, nil
	} else {
		switch t.Kind() {
		case reflect.String:
			return basetypes.StringType{}, nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			fallthrough
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return basetypes.Int64Type{}, nil
		case reflect.Float32, reflect.Float64:
			return basetypes.Float64Type{}, nil
		case reflect.Bool:
			return basetypes.BoolType{}, nil
		default:
			diags.Append(diag.NewErrorDiagnostic("Invalid type", fmt.Sprintf(`%T has unsupported type: %s`, t, t.Kind())))
			return nil, diags
		}
	}
}
