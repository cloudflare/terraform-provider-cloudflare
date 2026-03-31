package types

import (
	"context"
	"fmt"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// Alternative to basetypes.NewFloat64Value that takes a *big.Float with arbitrary precision
// directly instead of taking a float64 at 53-bit precision.
func NewFloat64ValueFromBigFloat(value *big.Float) (basetypes.Float64Value, error) {
	tfVal := tftypes.NewValue(tftypes.Number, value)
	out, err := basetypes.Float64Type{}.ValueFromTerraform(context.Background(), tfVal)
	if out != nil && err == nil {
		return out.(basetypes.Float64Value), nil
	}
	return basetypes.Float64Value{}, err
}

// Alternative to basetypes.NewFloat64Value that takes a float64 and converts it
// to a big.Float with precision and rounding to match that of HCL config parsing
// instead of 53-bit precision.
func NewFloat64Value(value float64) basetypes.Float64Value {
	// We ignore the error here since (Float64Type).ValueFromTerraform only errors
	// when input is not accurately representable in float64 and we are starting
	// with a float64.
	out, _ := NewFloat64ValueFromBigFloat(new(big.Float).SetPrec(512).SetMode(big.ToNearestEven).SetFloat64(value))
	return out
}

// Creates a basetypes.Float64Value from a string where the underlying *big.Float
// has precision matching that of HCL config parsing.
func NewFloat64ValueFromString(value string) (basetypes.Float64Value, error) {
	bf, err := parseBigFloatingPoint(value)
	if err != nil {
		return basetypes.Float64Value{}, fmt.Errorf("failed to parse as basetypes.Float64Value: %w", err)
	}
	return NewFloat64ValueFromBigFloat(bf)
}

// Creates a basetypes.Float64Value from a string where the underlying *big.Float
// has precision matching that of HCL config parsing. Panics on error.
func NewFloat64ValueFromStringUnsafe(value string) basetypes.Float64Value {
	out, err := NewFloat64ValueFromString(value)
	if err != nil {
		panic(err)
	}
	return out
}

// Creates a basetypes.NumberValue from a string where the underlying *big.Float
// has precision matching that of HCL config parsing.
func NewNumberValueFromString(value string) (basetypes.NumberValue, error) {
	bf, err := parseBigFloatingPoint(value)
	if err != nil {
		return basetypes.NumberValue{}, fmt.Errorf("failed to parse as basetypes.NumberValue: %w", err)
	}
	return basetypes.NewNumberValue(bf), nil
}

// Creates a basetypes.NumberValue from a string where the underlying *big.Float
// has precision matching that of HCL config parsing. Panics on error.
func NewNumberValueFromStringUnsafe(value string) basetypes.NumberValue {
	out, err := NewNumberValueFromString(value)
	if err != nil {
		panic(err)
	}
	return out
}

// parseBigFloatingPoint parses a string with floating point precision
// matching that of HCL config parsing.
func parseBigFloatingPoint(value string) (*big.Float, error) {
	bf, _, err := big.ParseFloat(value, 10, 512, big.ToNearestEven)
	return bf, err
}
