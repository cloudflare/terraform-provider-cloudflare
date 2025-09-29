package types

import (
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func floatValue(value attr.Value) (bool, *big.Float) {
	if value == nil {
		return false, nil
	}

	switch v := value.(type) {
	case basetypes.Float32Value:
		return true, big.NewFloat(float64(v.ValueFloat32()))
	case basetypes.Float64Value:
		return true, big.NewFloat(v.ValueFloat64())
	case basetypes.NumberValue:
		return true, v.ValueBigFloat()
	default:
		return false, nil
	}
}

func intValue(value attr.Value) (bool, *big.Int) {
	if value == nil {
		return false, nil
	}

	switch v := value.(type) {
	case basetypes.Int32Value:
		return true, big.NewInt(int64(v.ValueInt32()))
	case basetypes.Int64Value:
		return true, big.NewInt((v.ValueInt64()))
	case basetypes.NumberValue:
		if i, a := v.ValueBigFloat().Int(nil); a == big.Exact {
			return true, i
		}
		return false, nil
	default:
		return false, nil
	}
}

func IntValue(value attr.Value) (bool, *big.Int) {
	if ok, i := intValue(value); ok {
		return ok, i
	}
	if ok, f := floatValue(value); ok {
		if i, a := f.Int(nil); a == big.Exact {
			return true, i
		}
	}

	return false, nil
}

func FloatValue(value attr.Value) (bool, *big.Float) {
	if ok, f := floatValue(value); ok {
		return ok, f
	}
	if ok, i := intValue(value); ok {
		return ok, big.NewFloat(0).SetInt(i)
	}

	return false, nil
}

func ChildItems(value attr.Value) (bool, []attr.Value) {
	if value == nil {
		return false, nil
	}

	switch v := value.(type) {
	case basetypes.ListValue:
		return true, v.Elements()
	case basetypes.TupleValue:
		return true, v.Elements()
	case basetypes.SetValue:
		return true, v.Elements()
	default:
		return false, nil
	}
}

func ChildAttributes(value attr.Value) (bool, map[string]attr.Value) {
	if value == nil {
		return false, nil
	}

	switch v := value.(type) {
	case basetypes.MapValue:
		return true, v.Elements()
	case basetypes.ObjectValue:
		return true, v.Attributes()
	default:
		return false, nil
	}
}
