package customfield

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var (
	_ basetypes.ListTypable  = (*ListType[basetypes.StringValue])(nil)
	_ basetypes.ListValuable = (*List[basetypes.StringValue])(nil)
)

// ListType represents a basetypes.ListType that declares the type of the element statically.
type ListType[T attr.Value] struct {
	basetypes.ListType
}

func NewListType[T attr.Value](ctx context.Context) ListType[T] {
	return ListType[T]{ListType: basetypes.ListType{ElemType: elemType[T](ctx)}}
}

func (t ListType[T]) Equal(o attr.Type) bool {
	other, ok := o.(ListType[T])
	if !ok {
		return false
	}

	return t.ListType.Equal(other.ListType)
}

func (t ListType[T]) String() string {
	var zero T
	return fmt.Sprintf("ListType[%T]", zero)
}

func (t ListType[T]) ValueFromList(ctx context.Context, in basetypes.ListValue) (basetypes.ListValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	if in.IsNull() {
		return NullList[T](ctx), diags
	}
	if in.IsUnknown() {
		return UnknownList[T](ctx), diags
	}

	v, d := basetypes.NewListValue(elemType[T](ctx), in.Elements())
	diags.Append(d...)
	if diags.HasError() {
		return UnknownList[T](ctx), diags
	}

	value := List[T]{
		ListValue: v,
	}

	return value, diags
}

func (t ListType[T]) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
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

func (t ListType[T]) ValueType(ctx context.Context) attr.Value {
	return UnknownList[T](ctx)
}

func (t ListType[T]) NullValue(ctx context.Context) (attr.Value, diag.Diagnostics) {
	var diags diag.Diagnostics
	return NullList[T](ctx), diags
}

type ListLike interface {
	ValueAttr(ctx context.Context) (any, diag.Diagnostics)
	NullValue(ctx context.Context) ListLike
	UnknownValue(ctx context.Context) ListLike
	KnownValue(ctx context.Context, T any) ListLike
	IsNullOrUnknown() bool
}

var _ ListLike = (*List[basetypes.StringValue])(nil)

// List represents a basetypes.ListValue that is defined by a struct.
type List[T attr.Value] struct {
	//lint:ignore U1000 the placeholder is for easy reflection-based-access
	placeholder T
	basetypes.ListValue
}

func (v List[T]) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	tv := v.ListValue
	if tv.ElementType(ctx) == nil {
		tv = NullList[T](ctx).ListValue
	}
	return tv.ToTerraformValue(ctx)
}

func (v List[T]) NullValue(ctx context.Context) ListLike {
	return NullList[T](ctx)
}

func (v List[T]) UnknownValue(ctx context.Context) ListLike {
	return UnknownList[T](ctx)
}

func (v List[T]) KnownValue(ctx context.Context, values any) ListLike {
	r, _ := NewList[T](ctx, values)
	return r
}

func (v List[T]) IsNullOrUnknown() bool {
	return v.IsNull() || v.IsUnknown()
}

func (v List[T]) Equal(o attr.Value) bool {
	other, ok := o.(List[T])
	if !ok {
		return false
	}

	return v.ListValue.Equal(other.ListValue)
}

func (v List[T]) Type(ctx context.Context) attr.Type {
	return NewListType[T](ctx)
}

func (v List[T]) ValueAttr(ctx context.Context) (any, diag.Diagnostics) {
	return v.Value(ctx)
}

func (v List[T]) Value(ctx context.Context) ([]T, diag.Diagnostics) {
	ts := []T{}
	for _, elem := range v.ListValue.Elements() {
		ts = append(ts, elem.(T))
	}
	return ts, nil
}

func NullList[T attr.Value](ctx context.Context) List[T] {
	return List[T]{ListValue: basetypes.NewListNull(elemType[T](ctx))}
}

func UnknownList[T attr.Value](ctx context.Context) List[T] {
	return List[T]{ListValue: basetypes.NewListUnknown(elemType[T](ctx))}
}

func NewList[T attr.Value](ctx context.Context, values any) (List[T], diag.Diagnostics) {
	var diags diag.Diagnostics

	attrs, ok := values.([]attr.Value)
	if !ok {
		ts, ok := values.([]T)
		if !ok {
			diags.AddError("unexpected type of values", fmt.Sprintf("expected %T or []attr.Value, got %T", []T{}, values))
			return UnknownList[T](ctx), diags
		}
		attrs = make([]attr.Value, len(ts))
		for i, v := range ts {
			attrs[i] = v
		}
	}

	listValue, d := basetypes.NewListValue(elemType[T](ctx), attrs)
	diags.Append(d...)
	if diags.HasError() {
		return UnknownList[T](ctx), diags
	}

	return List[T]{ListValue: listValue}, nil
}

func NewListMust[T attr.Value](ctx context.Context, values []attr.Value) List[T] {
	o, diags := NewList[T](ctx, values)
	if diags.HasError() {
		panic(fmt.Errorf("unexpected error creating a customfield.List: %v", diags))
	}
	return o
}

func elemType[T attr.Value](ctx context.Context) attr.Type {
	var ty T
	return ty.Type(ctx)
}
