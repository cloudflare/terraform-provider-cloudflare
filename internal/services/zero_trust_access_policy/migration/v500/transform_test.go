package v500

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestTransformConditions_SkipsEmptyAuthMethodAndCommonName(t *testing.T) {
	ctx := context.Background()

	conds, diags := transformConditions(ctx, []SourceConditionGroupModel{
		{
			AuthMethod: types.StringValue(""),
			CommonName: types.StringValue("   "),
		},
	})

	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}

	if len(conds) != 0 {
		t.Fatalf("expected 0 conditions, got %d: %#v", len(conds), conds)
	}
}

func TestTransformConditions_IncludesNonEmptyAuthMethodAndCommonName(t *testing.T) {
	ctx := context.Background()

	conds, diags := transformConditions(ctx, []SourceConditionGroupModel{
		{
			AuthMethod: types.StringValue(" otp "),
			CommonName: types.StringValue(" device1.example.com "),
		},
	})

	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}

	if len(conds) != 2 {
		t.Fatalf("expected 2 conditions, got %d: %#v", len(conds), conds)
	}

	var gotAuthMethod, gotCommonName bool
	for _, c := range conds {
		if c.AuthMethod != nil {
			gotAuthMethod = true
			if c.AuthMethod.AuthMethod.ValueString() != "otp" {
				t.Fatalf("expected trimmed auth_method 'otp', got %q", c.AuthMethod.AuthMethod.ValueString())
			}
		}

		if c.CommonName != nil {
			gotCommonName = true
			if c.CommonName.CommonName.ValueString() != "device1.example.com" {
				t.Fatalf("expected trimmed common_name 'device1.example.com', got %q", c.CommonName.CommonName.ValueString())
			}
		}
	}

	if !gotAuthMethod {
		t.Fatal("expected auth_method condition, got none")
	}

	if !gotCommonName {
		t.Fatal("expected common_name condition, got none")
	}
}
