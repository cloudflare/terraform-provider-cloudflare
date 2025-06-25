package customfield

import (
	"context"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var (
	_ basetypes.ListTypable  = (*NestedObjectListType[basetypes.StringValue])(nil)
	_ basetypes.ListValuable = (*NestedObjectList[basetypes.StringValue])(nil)
)

// NestedObjectListType represents a basetypes.ListType that declares the type of the nested data statically.
type NestedObjectListType[T any] struct {
	basetypes.ListType
}

func NewNestedObjectListType[T any](ctx context.Context) NestedObjectListType[T] {
	elemType, err := attrType[T](ctx)
	if err.HasError() {
		panic(fmt.Errorf("unexpected error creating NestedObjectListType: %v", err))
	}
	t := NestedObjectListType[T]{ListType: basetypes.ListType{ElemType: elemType}}
	return t
}

func (t NestedObjectListType[T]) Equal(o attr.Type) bool {
	other, ok := o.(NestedObjectListType[T])
	if !ok {
		return false
	}

	return t.ListType.Equal(other.ListType)
}

func (t NestedObjectListType[T]) String() string {
	var zero T
	return fmt.Sprintf("ListType[%T]", zero)
}

func (t NestedObjectListType[T]) ValueFromList(ctx context.Context, in basetypes.ListValue) (basetypes.ListValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	if in.IsNull() {
		return NullObjectList[T](ctx), diags
	}
	if in.IsUnknown() {
		return UnknownObjectList[T](ctx), diags
	}

	ty, d := attrType[T](ctx)
	diags.Append(d...)
	if diags.HasError() {
		return UnknownObjectList[T](ctx), diags
	}

	v, d := basetypes.NewListValue(ty, in.Elements())
	diags.Append(d...)
	if diags.HasError() {
		return UnknownObjectList[T](ctx), diags
	}

	value := NestedObjectList[T]{
		ListValue: v,
	}

	return value, diags
}

func (t NestedObjectListType[T]) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.ListType.ValueFromTerraform(ctx, in)

	if err != nil {
		return nil, err
	}

	listValue, ok := attrValue.(basetypes.ListValue)
	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	listValuable, diags := t.ValueFromList(ctx, listValue)
	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting ListValue to ListValuable: %v", diags)
	}

	return listValuable, nil
}

func (t NestedObjectListType[T]) ValueType(ctx context.Context) attr.Value {
	return NestedObjectList[T]{}
}

func (t NestedObjectListType[T]) NullValue(ctx context.Context) (attr.Value, diag.Diagnostics) {
	var diags diag.Diagnostics
	return NullObjectList[T](ctx), diags
}

type NestedObjectListLike interface {
	AsStructSlice(ctx context.Context) (any, diag.Diagnostics)
	NullValue(ctx context.Context) NestedObjectListLike
	UnknownValue(ctx context.Context) NestedObjectListLike
	KnownValue(ctx context.Context, T any) NestedObjectListLike
	IsNullOrUnknown() bool
}

var _ NestedObjectListLike = (*NestedObjectList[basetypes.StringValue])(nil)

// NestedObjectList represents a basetypes.ListValue that is defined by a struct.
type NestedObjectList[T any] struct {
	//lint:ignore U1000 the placeholder is for easy reflection-based-access
	placeholder T
	basetypes.ListValue
}

func (v NestedObjectList[T]) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	tv := v.ListValue
	if tv.ElementType(ctx) == nil {
		tv = NullObjectList[T](ctx).ListValue
	}
	return tv.ToTerraformValue(ctx)
}

func (v NestedObjectList[T]) NullValue(ctx context.Context) NestedObjectListLike {
	return NullObjectList[T](ctx)
}

func (v NestedObjectList[T]) UnknownValue(ctx context.Context) NestedObjectListLike {
	return UnknownObjectList[T](ctx)
}

func (v NestedObjectList[T]) KnownValue(ctx context.Context, anyValues any) NestedObjectListLike {
	r, diags := NewObjectList(ctx, anyValues.([]T))
	if diags.HasError() {
		panic(fmt.Errorf("unexpected error creating NestedObjectList: %v", diags))
	}
	return r
}

func (v NestedObjectList[T]) IsNullOrUnknown() bool {
	return v.IsNull() || v.IsUnknown()
}

func (v NestedObjectList[T]) Equal(o attr.Value) bool {
	other, ok := o.(NestedObjectList[T])
	if !ok {
		return false
	}

	return v.ListValue.Equal(other.ListValue)
}

func (v NestedObjectList[T]) Type(ctx context.Context) attr.Type {
	return NewNestedObjectListType[T](ctx)
}

func (v NestedObjectList[T]) AsStructSlice(ctx context.Context) (any, diag.Diagnostics) {
	return v.AsStructSliceT(ctx)
}

func (v NestedObjectList[T]) AsStructSliceT(ctx context.Context) ([]T, diag.Diagnostics) {
	ts := []T{}
	if len(v.ListValue.Elements()) == 0 {
		return ts, nil
	}
	diags := v.ListValue.ElementsAs(ctx, &ts, true)
	if diags.HasError() {
		return nil, diags
	}
	return ts, nil
}

func (v NestedObjectList[T]) Value(ctx context.Context) ([]attr.Value, diag.Diagnostics) {
	return v.ListValue.Elements(), nil
}

func NullObjectList[T any](ctx context.Context) NestedObjectList[T] {
	var diags diag.Diagnostics

	elemType, d := attrType[T](ctx)
	diags.Append(d...)
	if diags.HasError() {
		panic(fmt.Errorf("unexpected error creating NestedObjectList: %v", diags))
	}
	return NestedObjectList[T]{ListValue: basetypes.NewListNull(elemType)}
}

func UnknownObjectList[T any](ctx context.Context) NestedObjectList[T] {
	var diags diag.Diagnostics

	elemType, d := attrType[T](ctx)
	diags.Append(d...)
	if diags.HasError() {
		panic(fmt.Errorf("unexpected error creating NestedObjectList: %v", diags))
	}
	return NestedObjectList[T]{ListValue: basetypes.NewListUnknown(elemType)}
}

func NewObjectList[T any](ctx context.Context, values []T) (NestedObjectList[T], diag.Diagnostics) {
	var diags diag.Diagnostics

	elemType, d := attrType[T](ctx)
	diags.Append(d...)
	if diags.HasError() {
		return UnknownObjectList[T](ctx), diags
	}
	listValue, d := basetypes.NewListValueFrom(ctx, elemType, values)
	diags.Append(d...)
	if diags.HasError() {
		return UnknownObjectList[T](ctx), diags
	}

	return NestedObjectList[T]{ListValue: listValue}, nil
}

func NewObjectListFromAttributes[T any](ctx context.Context, values []attr.Value) (NestedObjectList[T], diag.Diagnostics) {
	var diags diag.Diagnostics

	elemType, d := attrType[T](ctx)
	diags.Append(d...)
	if diags.HasError() {
		return UnknownObjectList[T](ctx), diags
	}
	listValue, d := basetypes.NewListValue(elemType, values)
	diags.Append(d...)
	if diags.HasError() {
		return UnknownObjectList[T](ctx), diags
	}

	return NestedObjectList[T]{ListValue: listValue}, nil
}

func NewObjectListMust[T any](ctx context.Context, values []T) NestedObjectList[T] {
	o, diags := NewObjectList(ctx, values)
	if diags.HasError() {
		panic(fmt.Errorf("unexpected error creating NestedObjectList: %v", diags))
	}
	return o
}

func NewObjectListFromValue[T any](ctx context.Context, value reflect.Value) (NestedObjectList[T], diag.Diagnostics) {
	var diags diag.Diagnostics

	elemType, d := attrTypeGeneric(ctx, reflect.Zero(value.Type().Elem()))
	diags.Append(d...)
	if diags.HasError() {
		return UnknownObjectList[T](ctx), diags
	}
	listValue, d := basetypes.NewListValueFrom(ctx, elemType, value.Interface())
	diags.Append(d...)
	if diags.HasError() {
		return UnknownObjectList[T](ctx), diags
	}

	return NestedObjectList[T]{ListValue: listValue}, nil
}

func NewObjectListFromValueMust[T any](ctx context.Context, value reflect.Value) NestedObjectList[T] {
	o, diags := NewObjectListFromValue[T](ctx, value)
	if diags.HasError() {
		panic(fmt.Errorf("unexpected error creating NestedObjectList: %v", diags))
	}
	return o
}

func attrType[T any](ctx context.Context) (types.ObjectType, diag.Diagnostics) {
	var diags diag.Diagnostics
	attrTypes, d := StructToAttributes[T](ctx)
	diags.Append(d...)
	if diags.HasError() {
		return types.ObjectType{}, diags
	}
	return types.ObjectType{AttrTypes: attrTypes}, diags
}

func attrTypeGeneric(ctx context.Context, value reflect.Value) (types.ObjectType, diag.Diagnostics) {
	var diags diag.Diagnostics
	attrTypes, d := StructFromAttributesGeneric(ctx, value)
	diags.Append(d...)
	if diags.HasError() {
		return types.ObjectType{}, diags
	}
	return types.ObjectType{AttrTypes: attrTypes}, diags
}
