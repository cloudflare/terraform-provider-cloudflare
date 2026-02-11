package types

import (
	"math/big"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// TestIntValue_NullNumber reproduces the crash in CUSTESC-56992
// When a NumberValue is null/unknown, ValueBigFloat() returns nil, and calling
// .Int(nil) on it causes a nil pointer dereference (segfault).
func TestIntValue_NullNumber(t *testing.T) {
	// Create a null NumberValue - this is what happens when a dynamic attribute
	// is null during semantic equality checks
	nullNumber := types.NumberNull()

	// This should NOT crash - but it does in the buggy version
	ok, val := IntValue(nullNumber)

	// With the fix, this should return false, nil
	if ok {
		t.Errorf("expected ok=false for null NumberValue, got ok=true with val=%v", val)
	}
	if val != nil {
		t.Errorf("expected val=nil for null NumberValue, got %v", val)
	}
}

// TestIntValue_UnknownNumber tests that unknown NumberValue also doesn't crash
func TestIntValue_UnknownNumber(t *testing.T) {
	unknownNumber := types.NumberUnknown()

	ok, val := IntValue(unknownNumber)

	if ok {
		t.Errorf("expected ok=false for unknown NumberValue, got ok=true with val=%v", val)
	}
	if val != nil {
		t.Errorf("expected val=nil for unknown NumberValue, got %v", val)
	}
}

// TestFloatValue_NullNumber tests the same bug in floatValue()
func TestFloatValue_NullNumber(t *testing.T) {
	nullNumber := types.NumberNull()

	ok, val := FloatValue(nullNumber)

	if ok {
		t.Errorf("expected ok=false for null NumberValue, got ok=true with val=%v", val)
	}
	if val != nil {
		t.Errorf("expected val=nil for null NumberValue, got %v", val)
	}
}

// TestFloatValue_UnknownNumber tests that unknown NumberValue also doesn't crash
func TestFloatValue_UnknownNumber(t *testing.T) {
	unknownNumber := types.NumberUnknown()

	ok, val := FloatValue(unknownNumber)

	if ok {
		t.Errorf("expected ok=false for unknown NumberValue, got ok=true with val=%v", val)
	}
	if val != nil {
		t.Errorf("expected val=nil for unknown NumberValue, got %v", val)
	}
}

// TestIntValue_ValidNumber ensures normal NumberValue still works
func TestIntValue_ValidNumber(t *testing.T) {
	validNumber := types.NumberValue(big.NewFloat(42))

	ok, val := IntValue(validNumber)

	if !ok {
		t.Errorf("expected ok=true for valid NumberValue, got ok=false")
	}
	if val == nil || val.Int64() != 42 {
		t.Errorf("expected val=42 for valid NumberValue, got %v", val)
	}
}

// TestFloatValue_ValidNumber ensures normal NumberValue still works
func TestFloatValue_ValidNumber(t *testing.T) {
	validNumber := types.NumberValue(big.NewFloat(3.14))

	ok, val := FloatValue(validNumber)

	if !ok {
		t.Errorf("expected ok=true for valid NumberValue, got ok=false")
	}
	if val == nil {
		t.Errorf("expected non-nil val for valid NumberValue")
	}
	// Check approximately equal (floating point)
	expected := 3.14
	actual, _ := val.Float64()
	if actual < expected-0.001 || actual > expected+0.001 {
		t.Errorf("expected val≈3.14 for valid NumberValue, got %v", actual)
	}
}
