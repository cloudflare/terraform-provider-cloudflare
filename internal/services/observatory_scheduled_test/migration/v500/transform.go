package v500

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts source (v4 SDKv2) state to target (v5 Plugin Framework) state.
//
// For observatory_scheduled_test, this is a simple pass-through transformation since:
// - All fields have the same names (no renames)
// - All fields have the same types (no type conversions)
// - No nested structures to transform
// - No deprecated fields to drop
//
// The only difference is that v5 adds new computed-only fields (schedule, test)
// which don't exist in v4 state. These will be populated by the API on first refresh.
func Transform(ctx context.Context, source SourceCloudflareObservatoryScheduledTestModel) (*TargetObservatoryScheduledTestModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Validate required fields
	if source.ZoneID.IsNull() || source.ZoneID.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"zone_id is required for observatory_scheduled_test migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}
	if source.URL.IsNull() || source.URL.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"url is required for observatory_scheduled_test migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	canonicalURL := canonicalizeObservatoryURL(source.URL.ValueString())

	// Create target with field copies and URL canonicalization
	target := &TargetObservatoryScheduledTestModel{
		ID:        types.StringValue(canonicalURL),
		ZoneID:    source.ZoneID,
		URL:       types.StringValue(canonicalURL),
		Region:    source.Region,
		Frequency: source.Frequency,
		// Schedule and Test are computed-only in v5, not present in v4
		// Leave them uninitialized - they will be populated by API on first refresh
	}

	return target, diags
}

func canonicalizeObservatoryURL(input string) string {
	value := strings.TrimSpace(input)
	value = strings.TrimPrefix(value, "https://")
	value = strings.TrimPrefix(value, "http://")
	value = strings.TrimSuffix(value, "/")
	if value == "" {
		return "/"
	}
	return value + "/"
}
