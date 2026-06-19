package customfield_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
)

func TestDynamicSemanticEqualsWithNullAndUnknownNumbers(t *testing.T) {
	ctx := context.Background()

	known := customfield.RawNormalizedDynamicValueFrom(basetypes.NewNumberValue(big.NewFloat(5)))
	otherKnown := customfield.RawNormalizedDynamicValueFrom(basetypes.NewNumberValue(big.NewFloat(5)))
	null := customfield.RawNormalizedDynamicValueFrom(basetypes.NewNumberNull())
	unknown := customfield.RawNormalizedDynamicValueFrom(basetypes.NewNumberUnknown())

	cases := map[string]struct {
		lhs      customfield.NormalizedDynamicValue
		rhs      customfield.NormalizedDynamicValue
		expected bool
	}{
		"null vs known":      {null, known, false},
		"known vs null":      {known, null, false},
		"unknown vs known":   {unknown, known, false},
		"known vs unknown":   {known, unknown, false},
		"null vs unknown":    {null, unknown, false},
		"null vs null":       {null, null, true},
		"unknown vs unknown": {unknown, unknown, true},
		"known vs known":     {known, otherKnown, true},
	}

	for name, testCase := range cases {
		t.Run(name, func(t *testing.T) {
			eq, diags := testCase.lhs.DynamicSemanticEquals(ctx, testCase.rhs)

			if diags.HasError() {
				t.Fatalf("unexpected diagnostics: %v", diags)
			}

			if eq != testCase.expected {
				t.Fatalf("expected %v, got %v", testCase.expected, eq)
			}
		})
	}
}
