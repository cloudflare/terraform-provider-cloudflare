package v500

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestTransformDescriptionEmptyStringToNull(t *testing.T) {
	source := SourceDevicePostureRuleModel{
		ID:          types.StringValue("id"),
		AccountID:   types.StringValue("acc"),
		Name:        types.StringValue("rule"),
		Type:        types.StringValue("firewall"),
		Description: types.StringValue(""),
	}

	target, diags := Transform(context.Background(), source)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	if !target.Description.IsNull() {
		t.Fatalf("expected null description, got %s", target.Description.String())
	}
}

func TestTransformDescriptionNonEmptyPreserved(t *testing.T) {
	source := SourceDevicePostureRuleModel{
		ID:          types.StringValue("id"),
		AccountID:   types.StringValue("acc"),
		Name:        types.StringValue("rule"),
		Type:        types.StringValue("firewall"),
		Description: types.StringValue("my description"),
	}

	target, diags := Transform(context.Background(), source)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	if target.Description.IsNull() || target.Description.ValueString() != "my description" {
		t.Fatalf("expected preserved description, got %s", target.Description.String())
	}
}
