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
	_ basetypes.SetTypable  = (*NestedObjectSetType[basetypes.StringValue])(nil)
	_ basetypes.SetValuable = (*NestedObjectSet[basetypes.StringValue])(nil)
)

// NestedObjectSetType represents a basetypes.SetType that declares the type of the nested data statically.
type NestedObjectSetType[T any] struct {
	basetypes.SetType
}

func NewNestedObjectSetType[T any](ctx context.Context) NestedObjectSetType[T] {
	elemType, _ := attrType[T](ctx)
	t := NestedObjectSetType[T]{SetType: basetypes.SetType{ElemType: elemType}}
	return t
}

func (t NestedObjectSetType[T]) Equal(o attr.Type) bool {
	other, ok := o.(NestedObjectSetType[T])
	if !ok {
		return false
	}

	return t.SetType.Equal(other.SetType)
}

func (t NestedObjectSetType[T]) String() string {
	var zero T
	return fmt.Sprintf("SetType[%T]", zero)
}

func (t NestedObjectSetType[T]) ValueFromSet(ctx context.Context, in basetypes.SetValue) (basetypes.SetValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	if in.IsNull() {
		return NullObjectSet[T](ctx), diags
	}
	if in.IsUnknown() {
		return UnknownObjectSet[T](ctx), diags
	}

	ty, d := attrType[T](ctx)
	diags.Append(d...)
	if diags.HasError() {
		return UnknownObjectSet[T](ctx), diags
	}

	v, d := basetypes.NewSetValue(ty, in.Elements())
	diags.Append(d...)
	if diags.HasError() {
		return UnknownObjectSet[T](ctx), diags
	}

	value := NestedObjectSet[T]{
		SetValue: v,
	}

	return value, diags
}

func (t NestedObjectSetType[T]) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
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

func (t NestedObjectSetType[T]) ValueType(ctx context.Context) attr.Value {
	return NestedObjectSet[T]{}
}

func (t NestedObjectSetType[T]) NullValue(ctx context.Context) (attr.Value, diag.Diagnostics) {
	var diags diag.Diagnostics
	return NullObjectSet[T](ctx), diags
}

var _ NestedObjectListLike = (*NestedObjectSet[basetypes.StringValue])(nil)

// NestedObjectSet represents a basetypes.SetValue that is defined by a struct.
type NestedObjectSet[T any] struct {
	//lint:ignore U1000 the placeholder is for easy reflection-based-access
	placeholder T
	basetypes.SetValue
}

func (v NestedObjectSet[T]) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	tv := v.SetValue
	if tv.ElementType(ctx) == nil {
		tv = NullObjectSet[T](ctx).SetValue
	}
	return tv.ToTerraformValue(ctx)
}

func (v NestedObjectSet[T]) NullValue(ctx context.Context) NestedObjectListLike {
	return NullObjectSet[T](ctx)
}

func (v NestedObjectSet[T]) UnknownValue(ctx context.Context) NestedObjectListLike {
	return UnknownObjectSet[T](ctx)
}

func (v NestedObjectSet[T]) KnownValue(ctx context.Context, anyValues any) NestedObjectListLike {
	r, diags := NewObjectSet(ctx, anyValues.([]T))
	if diags.HasError() {
		panic(fmt.Errorf("unexpected error creating NestedObjectSet: %v", diags))
	}
	return r
}

func (v NestedObjectSet[T]) IsNullOrUnknown() bool {
	return v.IsNull() || v.IsUnknown()
}

func (v NestedObjectSet[T]) Equal(o attr.Value) bool {
	other, ok := o.(NestedObjectSet[T])
	if !ok {
		return false
	}

	return v.SetValue.Equal(other.SetValue)
}

func (v NestedObjectSet[T]) Type(ctx context.Context) attr.Type {
	return NewNestedObjectSetType[T](ctx)
}

func (v NestedObjectSet[T]) AsStructSlice(ctx context.Context) (any, diag.Diagnostics) {
	return v.AsStructSliceT(ctx)
}

func (v NestedObjectSet[T]) AsStructSliceT(ctx context.Context) ([]T, diag.Diagnostics) {
	ts := []T{}
	if len(v.SetValue.Elements()) == 0 {
		return ts, nil
	}
	diags := v.SetValue.ElementsAs(ctx, &ts, true)
	if diags.HasError() {
		return nil, diags
	}
	return ts, nil
}

func (v NestedObjectSet[T]) Value(ctx context.Context) ([]attr.Value, diag.Diagnostics) {
	return v.SetValue.Elements(), nil
}

func NullObjectSet[T any](ctx context.Context) NestedObjectSet[T] {
	var diags diag.Diagnostics

	elemType, d := attrType[T](ctx)
	diags.Append(d...)
	if diags.HasError() {
		panic(fmt.Errorf("unexpected error creating NestedObjectSet: %v", diags))
	}
	return NestedObjectSet[T]{SetValue: basetypes.NewSetNull(elemType)}
}

func UnknownObjectSet[T any](ctx context.Context) NestedObjectSet[T] {
	var diags diag.Diagnostics

	elemType, d := attrType[T](ctx)
	diags.Append(d...)
	if diags.HasError() {
		panic(fmt.Errorf("unexpected error creating NestedObjectSet: %v", diags))
	}
	return NestedObjectSet[T]{SetValue: basetypes.NewSetUnknown(elemType)}
}

func NewObjectSet[T any](ctx context.Context, values []T) (NestedObjectSet[T], diag.Diagnostics) {
	var diags diag.Diagnostics

	elemType, d := attrType[T](ctx)
	diags.Append(d...)
	if diags.HasError() {
		return UnknownObjectSet[T](ctx), diags
	}
	setValue, d := basetypes.NewSetValueFrom(ctx, elemType, values)
	diags.Append(d...)
	if diags.HasError() {
		return UnknownObjectSet[T](ctx), diags
	}

	return NestedObjectSet[T]{SetValue: setValue}, nil
}

func NewObjectSetFromAttributes[T any](ctx context.Context, values []attr.Value) (NestedObjectSet[T], diag.Diagnostics) {
	var diags diag.Diagnostics

	elemType, d := attrType[T](ctx)
	diags.Append(d...)
	if diags.HasError() {
		return UnknownObjectSet[T](ctx), diags
	}
	setValue, d := basetypes.NewSetValue(elemType, values)
	diags.Append(d...)
	if diags.HasError() {
		return UnknownObjectSet[T](ctx), diags
	}

	return NestedObjectSet[T]{SetValue: setValue}, nil
}

func NewObjectSetMust[T any](ctx context.Context, values []T) NestedObjectSet[T] {
	o, diags := NewObjectSet(ctx, values)
	if diags.HasError() {
		panic(fmt.Errorf("unexpected error creating NestedObjectSet: %v", diags))
	}
	return o
}
