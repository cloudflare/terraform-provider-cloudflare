package v500

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV4 handles state upgrades from the v4 Plugin Framework provider (schema_version=1) to v5 (version=500).
//
// The v4 provider used Plugin Framework with ListNestedBlock (stored as arrays in state JSON).
// This function reads the v4 state using the prior schema, transforms to v5 format, and writes v5 state.
//
// IMPORTANT: This must only be called when req.State was populated with the v4 PriorSchema.
// When called from UpgradeFromV1Ambiguous (where PriorSchema is nil), use
// upgradeFromV4ViaRawState instead.
func UpgradeFromV4(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading cloudflare_ruleset state from v4 Plugin Framework provider (schema_version=1)")

	// Parse v4 state using v4 schema
	var v4State SourceV4RulesetModel
	resp.Diagnostics.Append(req.State.Get(ctx, &v4State)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform v4 -> v5
	v5State, diags := Transform(ctx, v4State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write transformed state
	resp.Diagnostics.Append(resp.State.Set(ctx, v5State)...)
	tflog.Info(ctx, "State upgrade from v4 to v5 completed successfully")
}

// UpgradeFromV5 handles state upgrades from v5 Plugin Framework provider (version=0) to v5 (version=500).
//
// This is a no-op upgrade since the schema is compatible - just bumps the version.
func UpgradeFromV5(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading cloudflare_ruleset state to version=500 (no-op)")
	resp.State.Raw = req.State.Raw
	tflog.Info(ctx, "State version bump to 500 completed")
}

// UpgradeFromV1Ambiguous handles schema_version=1 state that is AMBIGUOUS between
// v4 Plugin Framework format and v5 production format.
//
// Both v4 and v5 production (v5.0-v5.18 with GetSchemaVersion(1,500)) stored state
// at schema_version=1, but with incompatible formats:
//   - v4: rules[].action_parameters is an ARRAY (ListNestedBlock)
//   - v5: rules[].action_parameters is an OBJECT (SingleNestedAttribute)
//
// migrations.go registers this with PriorSchema=nil so the framework skips pre-decoding
// req.State entirely. Both paths operate exclusively on req.RawState.JSON.
//
// Detection: inspect rules[0].action_parameters in raw JSON.
// If array -> v4 format -> transform via raw state.
// If object (or absent) -> v5 format -> no-op re-decode with target schema.
func UpgradeFromV1Ambiguous(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading cloudflare_ruleset state from schema_version=1 (detecting v4 vs v5 format)")

	isV4, err := detectV4RulesetState(req)
	if err != nil {
		// Cannot determine format — default to no-op to avoid data loss.
		tflog.Warn(ctx, "Could not detect ruleset state format, defaulting to no-op", map[string]interface{}{
			"error": err.Error(),
		})
		// Fall through to no-op path
	}

	if isV4 {
		tflog.Info(ctx, "Detected v4 ruleset format (action_parameters is array), performing transformation via raw state")
		upgradeFromV4ViaRawState(ctx, req, resp)
		return
	}

	// v5 production state: re-decode raw JSON with the target schema.
	// req.State is nil here (PriorSchema=nil), so we must use req.RawState directly.
	tflog.Info(ctx, "Detected v5 ruleset format (action_parameters is object), performing no-op version bump via raw state")
	targetType := resp.State.Schema.Type().TerraformType(ctx)
	rawValue, unmarshalErr := req.RawState.Unmarshal(targetType)
	if unmarshalErr != nil {
		resp.Diagnostics.AddError(
			"Failed to unmarshal v5 ruleset state",
			"Could not parse raw state as v5 format: "+unmarshalErr.Error(),
		)
		return
	}
	resp.State.Raw = rawValue
	tflog.Info(ctx, "State version bump from 1 to 500 completed")
}

// upgradeFromV4ViaRawState performs the v4->v5 transformation by parsing req.RawState
// directly with the v4 schema. This is necessary when the upgrader was registered with
// PriorSchema=nil (so req.State is nil and cannot be used to parse v4 arrays).
func upgradeFromV4ViaRawState(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	if req.RawState == nil {
		resp.Diagnostics.AddError("Missing raw state", "RawState was nil during v4 ruleset upgrade")
		return
	}

	// Parse raw state using the v4 schema
	v4Schema := SourceV4RulesetSchema()
	v4Type := v4Schema.Type().TerraformType(ctx)

	rawValue, err := req.RawState.Unmarshal(v4Type)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to unmarshal v4 ruleset state",
			fmt.Sprintf("Could not parse raw state as v4 format: %s", err),
		)
		return
	}

	// Build a synthetic req with the v4-typed state so UpgradeFromV4 can call req.State.Get
	v4State := &tfsdk.State{Raw: rawValue, Schema: v4Schema}
	syntheticReq := resource.UpgradeStateRequest{
		RawState: req.RawState,
		State:    v4State,
	}

	UpgradeFromV4(ctx, syntheticReq, resp)
}

// detectV4RulesetState returns true if the raw state is in v4 format.
// v4 format: rules[0].action_parameters is a JSON array [].
// v5 format: rules[0].action_parameters is a JSON object {} or absent.
func detectV4RulesetState(req resource.UpgradeStateRequest) (bool, error) {
	if req.RawState == nil || len(req.RawState.JSON) == 0 {
		return false, nil
	}

	var raw map[string]interface{}
	if err := json.Unmarshal(req.RawState.JSON, &raw); err != nil {
		return false, err
	}

	rules, ok := raw["rules"].([]interface{})
	if !ok || len(rules) == 0 {
		// No rules — cannot determine format, assume v5 (safe default).
		return false, nil
	}

	firstRule, ok := rules[0].(map[string]interface{})
	if !ok {
		return false, nil
	}

	ap, exists := firstRule["action_parameters"]
	if !exists || ap == nil {
		// No action_parameters — could be either format; assume v5.
		return false, nil
	}

	switch ap.(type) {
	case []interface{}:
		return true, nil // v4: array
	case map[string]interface{}:
		return false, nil // v5: object
	default:
		return false, nil
	}
}
