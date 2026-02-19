package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// Transform converts source (legacy v4 SDKv2) state to target (current v5 Plugin Framework) state.
// This function handles the field rename from "enabled" to "flag".
func Transform(ctx context.Context, source SourceCloudflareLogpullRetentionModel) (*TargetLogpullRetentionModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Step 1: Validate required fields
	if source.ZoneID.IsNull() || source.ZoneID.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"zone_id is required for logpull_retention migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	// Step 2: Initialize target with direct copies and field rename
	target := &TargetLogpullRetentionModel{
		ID:     source.ID,
		ZoneID: source.ZoneID,
		Flag:   source.Enabled, // Rename: enabled → flag
	}

	return target, diags
}
