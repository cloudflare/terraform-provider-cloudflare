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
	_ basetypes.MapTypable  = (*MapType[basetypes.StringValue])(nil)
	_ basetypes.MapValuable = (*Map[basetypes.StringValue])(nil)
)

type MapLike interface {
	basetypes.MapValuable
	ValueAttr(ctx context.Context) (any, diag.Diagnostics)
	NullValue(ctx context.Context) MapLike
	UnknownValue(ctx context.Context) MapLike
	KnownValue(ctx context.Context, T any) MapLike
}

// MapType represents a basetypes.MapType that declares the type of the element statically.
type MapType[T attr.Value] struct {
	basetypes.MapType
}

func newMapType[T attr.Value](ctx context.Context) (MapType[T], diag.Diagnostics) {
	return MapType[T]{MapType: basetypes.MapType{ElemType: elemType[T](ctx)}}, nil
}

func NewMapType[T attr.Value](ctx context.Context) MapType[T] {
	t, _ := newMapType[T](ctx)
	return t
}

func (t MapType[T]) Equal(o attr.Type) bool {
	other, ok := o.(MapType[T])
	if !ok {
		return false
	}

	return t.MapType.Equal(other.MapType)
}

func (t MapType[T]) String() string {
	var zero T
	return fmt.Sprintf("MapType[%T]", zero)
}

func (t MapType[T]) ValueFromMap(ctx context.Context, in basetypes.MapValue) (basetypes.MapValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	if in.IsNull() {
		return NullMap[T](ctx), diags
	}
	if in.IsUnknown() {
		return UnknownMap[T](ctx), diags
	}

	v, d := basetypes.NewMapValue(elemType[T](ctx), in.Elements())
	diags.Append(d...)
	if diags.HasError() {
		return UnknownMap[T](ctx), diags
	}

	value := Map[T]{
		MapValue: v,
	}

	return value, diags
}

func (t MapType[T]) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.MapType.ValueFromTerraform(ctx, in)

	if err != nil {
		return nil, err
	}

	mapValue, ok := attrValue.(basetypes.MapValue)
	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	mapValuable, diags := t.ValueFromMap(ctx, mapValue)
	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting MapValue to MapValuable: %v", diags)
	}

	return mapValuable, nil
}

func (t MapType[T]) ValueType(ctx context.Context) attr.Value {
	return UnknownMap[T](ctx)
}

func (t MapType[T]) NullValue(ctx context.Context) (attr.Value, diag.Diagnostics) {
	var diags diag.Diagnostics
	return NullMap[T](ctx), diags
}

var _ MapLike = (*Map[basetypes.StringValue])(nil)

// Map represents a basetypes.MapValue that is defined by a struct.
type Map[T attr.Value] struct {
	//lint:ignore U1000 the placeholder is for easy reflection-based-access
	placeholder T
	basetypes.MapValue
}

func (v Map[T]) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	tv := v.MapValue
	if tv.ElementType(ctx) == nil {
		tv = NullMap[T](ctx).MapValue
	}
	return tv.ToTerraformValue(ctx)
}

func (v Map[T]) NullValue(ctx context.Context) MapLike {
	return NullMap[T](ctx)
}

func (v Map[T]) UnknownValue(ctx context.Context) MapLike {
	return UnknownMap[T](ctx)
}

func (v Map[T]) KnownValue(ctx context.Context, values any) MapLike {
	r, _ := NewMap[T](ctx, values)
	return r
}

func (v Map[T]) Equal(o attr.Value) bool {
	other, ok := o.(Map[T])
	if !ok {
		return false
	}

	return v.MapValue.Equal(other.MapValue)
}

func (v Map[T]) Type(ctx context.Context) attr.Type {
	return NewMapType[T](ctx)
}

func (v Map[T]) ValueAttr(ctx context.Context) (any, diag.Diagnostics) {
	return v.Value(ctx)
}

func (v Map[T]) Value(ctx context.Context) (map[string]T, diag.Diagnostics) {
	ts := map[string]T{}
	for key, elem := range v.MapValue.Elements() {
		ts[key] = elem.(T)
	}
	return ts, nil
}

func NullMap[T attr.Value](ctx context.Context) Map[T] {
	return Map[T]{MapValue: basetypes.NewMapNull(elemType[T](ctx))}
}

func UnknownMap[T attr.Value](ctx context.Context) Map[T] {
	return Map[T]{MapValue: basetypes.NewMapUnknown(elemType[T](ctx))}
}

func NewMap[T attr.Value](ctx context.Context, values any) (Map[T], diag.Diagnostics) {
	var diags diag.Diagnostics

	attrs, ok := values.(map[string]attr.Value)
	if !ok {
		ts, ok := values.(map[string]T)
		if !ok {
			diags.AddError("unexpected type of values", fmt.Sprintf("expected %T or map[string]attr.Value, got %T", map[string]T{}, values))
			return UnknownMap[T](ctx), diags
		}
		attrs = make(map[string]attr.Value, len(ts))
		for i, v := range ts {
			attrs[i] = v
		}
	}

	setValue, d := basetypes.NewMapValue(elemType[T](ctx), attrs)
	diags.Append(d...)
	if diags.HasError() {
		return UnknownMap[T](ctx), diags
	}

	return Map[T]{MapValue: setValue}, nil
}

func NewMapMust[T attr.Value](ctx context.Context, values map[string]T) Map[T] {
	o, diags := NewMap[T](ctx, values)
	if diags.HasError() {
		panic(fmt.Errorf("unexpected error creating a customfield.Map: %v", diags))
	}
	return o
}
