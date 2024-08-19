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

	m, d := StructToAttributes[T](ctx)
	diags.Append(d...)
	if diags.HasError() {
		return NestedObjectType[T]{}, diags
	}

	t := NestedObjectType[T]{}
	t.ObjectType = basetypes.ObjectType{AttrTypes: m}
	return t, diags
}

func NewNestedObjectType[T any](ctx context.Context) NestedObjectType[T] {
	t, _ := newNestedObjectType[T](ctx)
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
	return NestedObject[T]{}
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

func (v NestedObject[T]) NullValue(ctx context.Context) NestedObjectLike {
	return NullObject[T](ctx)
}

func (v NestedObject[T]) UnknownValue(ctx context.Context) NestedObjectLike {
	return UnknownObject[T](ctx)
}

func (v NestedObject[T]) KnownValue(ctx context.Context, t any) NestedObjectLike {
	r, _ := NewObject[T](ctx, t.(*T))
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
	t, _ := StructToAttributes[T](ctx)
	return NestedObject[T]{ObjectValue: basetypes.NewObjectNull(t)}
}

func UnknownObject[T any](ctx context.Context) NestedObject[T] {
	t, _ := StructToAttributes[T](ctx)
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
	o, _ := NewObject(ctx, t)
	return o
}

func StructToAttributes[T any](ctx context.Context) (map[string]attr.Type, diag.Diagnostics) {
	var diags diag.Diagnostics
	var t T
	val := reflect.ValueOf(t)
	typ := val.Type()

	if typ.Kind() == reflect.Ptr && typ.Elem().Kind() == reflect.Struct {
		val = reflect.New(typ.Elem()).Elem()
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		diags.Append(diag.NewErrorDiagnostic("Invalid type", fmt.Sprintf("%T has unsupported type: %s", t, typ)))
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
			diags.Append(diag.NewErrorDiagnostic("Invalid type", fmt.Sprintf(`%T is missing a "tfsdk" struct tag on %s`, t, field.Name)))
			return nil, diags
		}

		if v, ok := val.Field(i).Interface().(attr.Value); ok {
			attributeTypes[tag] = v.Type(ctx)
		}
	}

	return attributeTypes, nil
}
