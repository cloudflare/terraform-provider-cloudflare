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
	_ basetypes.SetTypable  = (*SetType[basetypes.StringValue])(nil)
	_ basetypes.SetValuable = (*Set[basetypes.StringValue])(nil)
)

// SetType represents a basetypes.SetType that declares the type of the element statically.
type SetType[T attr.Value] struct {
	basetypes.SetType
}

func NewSetType[T attr.Value](ctx context.Context) SetType[T] {
	return SetType[T]{SetType: basetypes.SetType{ElemType: elemType[T](ctx)}}
}

func (t SetType[T]) Equal(o attr.Type) bool {
	other, ok := o.(SetType[T])
	if !ok {
		return false
	}

	return t.SetType.Equal(other.SetType)
}

func (t SetType[T]) String() string {
	var zero T
	return fmt.Sprintf("SetType[%T]", zero)
}

func (t SetType[T]) ValueFromSet(ctx context.Context, in basetypes.SetValue) (basetypes.SetValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	if in.IsNull() {
		return NullSet[T](ctx), diags
	}
	if in.IsUnknown() {
		return UnknownSet[T](ctx), diags
	}

	v, d := basetypes.NewSetValue(elemType[T](ctx), in.Elements())
	diags.Append(d...)
	if diags.HasError() {
		return UnknownSet[T](ctx), diags
	}

	value := Set[T]{
		SetValue: v,
	}

	return value, diags
}

func (t SetType[T]) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.SetType.ValueFromTerraform(ctx, in)

	if err != nil {
		return nil, err
	}

	setValue, ok := attrValue.(basetypes.SetValue)
	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	setValuable, diags := t.ValueFromSet(ctx, setValue)
	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting SetValue to SetValuable: %v", diags)
	}

	return setValuable, nil
}

func (t SetType[T]) ValueType(ctx context.Context) attr.Value {
	return UnknownSet[T](ctx)
}

func (t SetType[T]) NullValue(ctx context.Context) (attr.Value, diag.Diagnostics) {
	var diags diag.Diagnostics
	return NullSet[T](ctx), diags
}

var _ ListLike = (*Set[basetypes.StringValue])(nil)

// Set represents a basetypes.SetValue that is defined by a struct.
type Set[T attr.Value] struct {
	//lint:ignore U1000 the placeholder is for easy reflection-based-access
	placeholder T
	basetypes.SetValue
}

func (v Set[T]) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	tv := v.SetValue
	if tv.ElementType(ctx) == nil {
		tv = NullSet[T](ctx).SetValue
	}
	return tv.ToTerraformValue(ctx)
}

func (v Set[T]) NullValue(ctx context.Context) ListLike {
	return NullSet[T](ctx)
}

func (v Set[T]) UnknownValue(ctx context.Context) ListLike {
	return UnknownSet[T](ctx)
}

func (v Set[T]) KnownValue(ctx context.Context, values any) ListLike {
	r, _ := NewSet[T](ctx, values)
	return r
}

func (v Set[T]) IsNullOrUnknown() bool {
	return v.IsNull() || v.IsUnknown()
}

func (v Set[T]) Equal(o attr.Value) bool {
	other, ok := o.(Set[T])
	if !ok {
		return false
	}

	return v.SetValue.Equal(other.SetValue)
}

func (v Set[T]) Type(ctx context.Context) attr.Type {
	return NewSetType[T](ctx)
}

func (v Set[T]) ValueAttr(ctx context.Context) (any, diag.Diagnostics) {
	return v.Value(ctx)
}

func (v Set[T]) Value(ctx context.Context) ([]T, diag.Diagnostics) {
	ts := []T{}
	for _, elem := range v.SetValue.Elements() {
		ts = append(ts, elem.(T))
	}
	return ts, nil
}

func NullSet[T attr.Value](ctx context.Context) Set[T] {
	return Set[T]{SetValue: basetypes.NewSetNull(elemType[T](ctx))}
}

func UnknownSet[T attr.Value](ctx context.Context) Set[T] {
	return Set[T]{SetValue: basetypes.NewSetUnknown(elemType[T](ctx))}
}

func NewSet[T attr.Value](ctx context.Context, values any) (Set[T], diag.Diagnostics) {
	var diags diag.Diagnostics

	attrs, ok := values.([]attr.Value)
	if !ok {
		ts, ok := values.([]T)
		if !ok {
			diags.AddError("unexpected type of values", fmt.Sprintf("expected %T or []attr.Value, got %T", []T{}, values))
			return UnknownSet[T](ctx), diags
		}
		attrs = make([]attr.Value, len(ts))
		for i, v := range ts {
			attrs[i] = v
		}
	}

	setValue, d := basetypes.NewSetValue(elemType[T](ctx), attrs)
	diags.Append(d...)
	if diags.HasError() {
		return UnknownSet[T](ctx), diags
	}

	return Set[T]{SetValue: setValue}, nil
}

func NewSetMust[T attr.Value](ctx context.Context, values []attr.Value) Set[T] {
	o, diags := NewSet[T](ctx, values)
	if diags.HasError() {
		panic(fmt.Errorf("unexpected error creating a customfield.Set: %v", diags))
	}
	return o
}
