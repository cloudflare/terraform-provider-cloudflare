package v500

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV0 handles state upgrades from schema_version=0 to version=500.
// Must handle BOTH formats since both v4 and early v5 wrote version=0:
// - v4 SDKv2 states: actions as array (list) - needs full transformation
// - Early v5 states: actions as object - just version bump
//
// PriorSchema is nil to bypass Terraform's schema parsing, allowing us to
// work directly with raw JSON and detect the format ourselves.
func UpgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading page_rule state from schema_version=0 to version=500")

	// STEP 1: Get raw JSON state (PriorSchema is nil, so use RawState)
	if req.RawState == nil || req.RawState.JSON == nil {
		resp.Diagnostics.AddError(
			"Missing raw state",
			"RawState or RawState.JSON is nil",
		)
		return
	}

	rawJSON := req.RawState.JSON

	tflog.Debug(ctx, "Raw state JSON", map[string]interface{}{
		"json_length": len(rawJSON),
	})

	// Unmarshal into generic map to check actions type
	var stateMap map[string]interface{}
	if err := json.Unmarshal(rawJSON, &stateMap); err != nil {
		resp.Diagnostics.AddError(
			"Failed to unmarshal state JSON",
			err.Error(),
		)
		return
	}

	// Check if actions field exists and what type it is
	actions, hasActions := stateMap["actions"]
	if !hasActions || actions == nil {
		tflog.Info(ctx, "No actions field found, performing version bump only")
		// For states without actions, just return - no migration needed
		// This shouldn't happen in practice but handle gracefully
		return
	}

	// Check if actions is an array (v4) or object (early v5)
	_, isV4Array := actions.([]interface{})

	tflog.Info(ctx, "Actions format detected", map[string]interface{}{
		"is_array": isV4Array,
	})

	// STEP 2: Handle based on detected format
	if isV4Array {
		tflog.Info(ctx, "Detected v4 SDKv2 format (actions as array), performing full transformation")

		// Parse JSON into v4 model manually (json.Unmarshal doesn't work with framework types)
		v4State, diags := parseJSONToSourceModel(ctx, rawJSON)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		// Transform v4 → v5
		targetState, diags := Transform(ctx, v4State)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		// Set the transformed state
		resp.Diagnostics.Append(resp.State.Set(ctx, targetState)...)
		tflog.Info(ctx, "State upgrade from v4 SDKv2 completed successfully")
	} else {
		// It's an object - early v5 format
		tflog.Info(ctx, "Detected early v5 format (actions as object), performing version bump")

		// Parse early v5 JSON into target model manually
		targetState, diags := parseJSONToTargetModel(ctx, rawJSON)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		// Write the state - it's already in the correct format, just need version bump
		resp.Diagnostics.Append(resp.State.Set(ctx, targetState)...)
		tflog.Info(ctx, "State version bump from 0 to 500 completed (early v5)")
	}
}

// UpgradeFromV4 handles state upgrade from schema version 4 to 500.
// This is a no-op version bump for released v5 versions (v5.0.0-v5.16.0).
// State structure is identical, so we just copy the raw state.
//
// CRITICAL: Use raw state copy, not Get/Set cycle, to avoid custom field type serialization issues.
func UpgradeFromV4(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading page_rule state from version 4 to 500 (no-op)")
	// Raw state copy avoids custom field type serialization issues
	resp.State.Raw = req.State.Raw
	tflog.Debug(ctx, "State upgrade from version 4 completed (no changes)")
}

// UpgradeFromLegacyV3 handles state upgrades from the legacy cloudflare_page_rule resource (v4.x SDKv2 provider).
// This is triggered when Terraform encounters state with schema_version=3 (SDKv2 standard version).
//
// Note: schema_version=3 was used in the v4.x SDKv2 provider.
// The state structure matches SourceCloudflarePageRuleModel.
func UpgradeFromLegacyV3(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading page_rule state from legacy v4.x SDKv2 provider (schema_version=3)")

	// Parse the state (schema_version=3, SDKv2 format)
	var sourceState SourceCloudflarePageRuleModel
	resp.Diagnostics.Append(req.State.Get(ctx, &sourceState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform to target (v5 Plugin Framework format)
	targetState, diags := Transform(ctx, sourceState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set the upgraded state
	resp.Diagnostics.Append(resp.State.Set(ctx, targetState)...)

	tflog.Info(ctx, "State upgrade from legacy v4.x SDKv2 provider completed successfully")
}
