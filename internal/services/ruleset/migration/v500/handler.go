package v500

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV4 handles state upgrades from the v4 Plugin Framework provider to v5 (version=500).
//
// The v4 provider used Plugin Framework with ListNestedBlock (stored as arrays in state JSON).
// This function reads the v4 state using the prior schema, transforms to v5 format, and writes v5 state.
//
// IMPORTANT: This must only be called when req.State was populated with the v4 PriorSchema.
// When called from UpgradeFromAmbiguous (where PriorSchema is nil), use
// upgradeFromV4ViaRawState instead.
func UpgradeFromV4(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading cloudflare_ruleset state from v4 Plugin Framework provider")

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

// UpgradeFromAmbiguous handles schema versions 0 and 1, both of which can contain
// v4 list-block state or v5 object-shaped state.
//
// migrations.go registers this with PriorSchema=nil so the framework skips pre-decoding
// req.State entirely. Both paths operate exclusively on req.RawState.JSON.
//
// Detection examines singleton blocks in every rule, not just the first rule's
// action_parameters. v5 state is re-decoded unchanged with the target schema.
func UpgradeFromAmbiguous(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading cloudflare_ruleset state (detecting v4 vs v5 format)")

	isV4, err := detectV4RulesetState(req)
	if err != nil {
		resp.Diagnostics.AddError("Failed to detect ruleset state format", err.Error())
		return
	}

	if isV4 {
		tflog.Info(ctx, "Detected v4 ruleset format, performing transformation via raw state")
		upgradeFromV4ViaRawState(ctx, req, resp)
		return
	}

	// v5 production state: re-decode raw JSON with the target schema.
	// req.State is nil here (PriorSchema=nil), so we must use req.RawState directly.
	tflog.Info(ctx, "Detected v5 ruleset format, performing no-op version bump via raw state")
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
	tflog.Info(ctx, "State version bump to 500 completed")
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

// detectV4RulesetState only probes fields that changed from singleton list blocks
// to objects. Lists shared by both versions (rules, algorithms, etc.) are not markers.
func detectV4RulesetState(req resource.UpgradeStateRequest) (bool, error) {
	if req.RawState == nil || len(req.RawState.JSON) == 0 {
		return false, fmt.Errorf("raw ruleset state is missing or empty")
	}

	var raw map[string]interface{}
	if err := json.Unmarshal(req.RawState.JSON, &raw); err != nil {
		return false, err
	}
	if raw == nil {
		return false, fmt.Errorf("raw ruleset state must be a JSON object")
	}

	if raw["rules"] == nil {
		return false, nil
	}
	rules, ok := raw["rules"].([]interface{})
	if !ok {
		return false, fmt.Errorf("rules must be a JSON array")
	}

	var hasV4, hasV5 bool
	for i, value := range rules {
		if value == nil {
			continue
		}
		rule, ok := value.(map[string]interface{})
		if !ok {
			return false, fmt.Errorf("rules[%d] must be a JSON object", i)
		}
		// These computed metadata fields existed only in legacy v4 state.
		if _, exists := rule["version"]; exists {
			hasV4 = true
		}
		if _, exists := rule["last_updated"]; exists {
			hasV4 = true
		}
		for _, field := range []string{"action_parameters", "logging", "ratelimit", "exposed_credential_check"} {
			switch rule[field].(type) {
			case nil:
			case []interface{}:
				hasV4 = true
			case map[string]interface{}:
				hasV5 = true
			default:
				return false, fmt.Errorf("rules[%d].%s must be a JSON array, object, or null", i, field)
			}
		}
	}
	if hasV4 && hasV5 {
		return false, fmt.Errorf("ruleset state mixes v4 list blocks and v5 objects")
	}
	return hasV4, nil
}
