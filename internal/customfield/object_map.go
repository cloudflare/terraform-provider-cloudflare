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
	_ basetypes.MapTypable  = (*NestedObjectMapType[basetypes.StringValue])(nil)
	_ basetypes.MapValuable = (*NestedObjectMap[basetypes.StringValue])(nil)
)

// NestedObjectMapType represents a basetypes.MapType that declares the type of the nested data statically.
type NestedObjectMapType[T any] struct {
	basetypes.MapType
}

func NewNestedObjectMapType[T any](ctx context.Context) NestedObjectMapType[T] {
	elemType, _ := attrType[T](ctx)
	t := NestedObjectMapType[T]{MapType: basetypes.MapType{ElemType: elemType}}
	return t
}

func (t NestedObjectMapType[T]) Equal(o attr.Type) bool {
	other, ok := o.(NestedObjectMapType[T])
	if !ok {
		return false
	}

	return t.MapType.Equal(other.MapType)
}

func (t NestedObjectMapType[T]) String() string {
	var zero T
	return fmt.Sprintf("MapType[%T]", zero)
}

func (t NestedObjectMapType[T]) ValueFromMap(ctx context.Context, in basetypes.MapValue) (basetypes.MapValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	if in.IsNull() {
		return NullObjectMap[T](ctx), diags
	}
	if in.IsUnknown() {
		return UnknownObjectMap[T](ctx), diags
	}

	ty, d := attrType[T](ctx)
	diags.Append(d...)
	if diags.HasError() {
		return UnknownObjectMap[T](ctx), diags
	}

	v, d := basetypes.NewMapValue(ty, in.Elements())
	diags.Append(d...)
	if diags.HasError() {
		return UnknownObjectMap[T](ctx), diags
	}

	value := NestedObjectMap[T]{
		MapValue: v,
	}

	return value, diags
}

func (t NestedObjectMapType[T]) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.MapType.ValueFromTerraform(ctx, in)

	if err != nil {
		return nil, err
	}

	setValue, ok := attrValue.(basetypes.MapValue)
	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	setValuable, diags := t.ValueFromMap(ctx, setValue)
	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting MapValue to MapValuable: %v", diags)
	}

	return setValuable, nil
}

func (t NestedObjectMapType[T]) ValueType(ctx context.Context) attr.Value {
	return UnknownObjectMap[T](ctx)
}

func (t NestedObjectMapType[T]) NullValue(ctx context.Context) (attr.Value, diag.Diagnostics) {
	var diags diag.Diagnostics
	return NullObjectMap[T](ctx), diags
}

type NestedObjectMapLike interface {
	basetypes.MapValuable
	AsStructMap(ctx context.Context) (any, diag.Diagnostics)
	NullValue(ctx context.Context) NestedObjectMapLike
	UnknownValue(ctx context.Context) NestedObjectMapLike
	KnownValue(ctx context.Context, T any) NestedObjectMapLike
}

var _ NestedObjectMapLike = (*NestedObjectMap[basetypes.StringValue])(nil)

// NestedObjectMap represents a basetypes.MapValue that is defined by a struct.
type NestedObjectMap[T any] struct {
	//lint:ignore U1000 the placeholder is for easy reflection-based-access
	placeholder T
	basetypes.MapValue
}

func (v NestedObjectMap[T]) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	tv := v.MapValue
	if tv.ElementType(ctx) == nil {
		tv = NullObjectMap[T](ctx).MapValue
	}
	return tv.ToTerraformValue(ctx)
}

func (v NestedObjectMap[T]) NullValue(ctx context.Context) NestedObjectMapLike {
	return NullObjectMap[T](ctx)
}

func (v NestedObjectMap[T]) UnknownValue(ctx context.Context) NestedObjectMapLike {
	return UnknownObjectMap[T](ctx)
}

func (v NestedObjectMap[T]) KnownValue(ctx context.Context, anyValues any) NestedObjectMapLike {
	r, diags := NewObjectMap(ctx, anyValues.(map[string]T))
	if diags.HasError() {
		panic(fmt.Errorf("unexpected error creating NestedObjectMap: %v", diags))
	}
	return r
}

func (v NestedObjectMap[T]) Equal(o attr.Value) bool {
	other, ok := o.(NestedObjectMap[T])
	if !ok {
		return false
	}

	return v.MapValue.Equal(other.MapValue)
}

func (v NestedObjectMap[T]) Type(ctx context.Context) attr.Type {
	return NewNestedObjectMapType[T](ctx)
}

func (v NestedObjectMap[T]) AsStructMap(ctx context.Context) (any, diag.Diagnostics) {
	return v.AsStructMapT(ctx)
}

func (v NestedObjectMap[T]) AsStructMapT(ctx context.Context) (map[string]T, diag.Diagnostics) {
	ts := map[string]T{}
	if len(v.MapValue.Elements()) == 0 {
		return ts, nil
	}
	diags := v.MapValue.ElementsAs(ctx, &ts, true)
	if diags.HasError() {
		return nil, diags
	}
	return ts, nil
}

func (v NestedObjectMap[T]) Value(ctx context.Context) (map[string]attr.Value, diag.Diagnostics) {
	return v.MapValue.Elements(), nil
}

func NullObjectMap[T any](ctx context.Context) NestedObjectMap[T] {
	var diags diag.Diagnostics

	elemType, d := attrType[T](ctx)
	diags.Append(d...)
	if diags.HasError() {
		panic(fmt.Errorf("unexpected error creating NestedObjectMap: %v", diags))
	}
	return NestedObjectMap[T]{MapValue: basetypes.NewMapNull(elemType)}
}

func UnknownObjectMap[T any](ctx context.Context) NestedObjectMap[T] {
	var diags diag.Diagnostics

	elemType, d := attrType[T](ctx)
	diags.Append(d...)
	if diags.HasError() {
		panic(fmt.Errorf("unexpected error creating NestedObjectMap: %v", diags))
	}
	return NestedObjectMap[T]{MapValue: basetypes.NewMapUnknown(elemType)}
}

func NewObjectMap[T any](ctx context.Context, values map[string]T) (NestedObjectMap[T], diag.Diagnostics) {
	var diags diag.Diagnostics

	elemType, d := attrType[T](ctx)
	diags.Append(d...)
	if diags.HasError() {
		return UnknownObjectMap[T](ctx), diags
	}
	setValue, d := basetypes.NewMapValueFrom(ctx, elemType, values)
	diags.Append(d...)
	if diags.HasError() {
		return UnknownObjectMap[T](ctx), diags
	}

	return NestedObjectMap[T]{MapValue: setValue}, nil
}

func NewObjectMapFromAttributes[T any](ctx context.Context, values map[string]attr.Value) (NestedObjectMap[T], diag.Diagnostics) {
	var diags diag.Diagnostics

	elemType, d := attrType[T](ctx)
	diags.Append(d...)
	if diags.HasError() {
		return UnknownObjectMap[T](ctx), diags
	}
	setValue, d := basetypes.NewMapValue(elemType, values)
	diags.Append(d...)
	if diags.HasError() {
		return UnknownObjectMap[T](ctx), diags
	}

	return NestedObjectMap[T]{MapValue: setValue}, nil
}

func NewObjectMapMust[T any](ctx context.Context, values map[string]T) NestedObjectMap[T] {
	o, diags := NewObjectMap(ctx, values)
	if diags.HasError() {
		panic(fmt.Errorf("unexpected error creating NestedObjectMap: %v", diags))
	}
	return o
}
