package v500

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV4 handles state upgrades from v4 SDKv2 provider (schema_version=1) to v5 (version=500).
//
// This performs a full transformation from v4 → v5 format including:
// - Field renames (default_pool_ids → default_pools, fallback_pool_id → fallback_pool)
// - Type conversions (Int64 → Float64 for ttl fields)
// - Structure transformations (arrays → single objects, sets → maps)
//
// The v4 state has schema_version=1 (from SDKv2 with migrations), and we transform it to v5 format.
func UpgradeFromV4(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading load_balancer state from v4 SDKv2 provider (schema_version=1)")

	// Parse v4 state using v4 source schema
	var v4State SourceCloudflareLoadBalancerModel
	resp.Diagnostics.Append(req.State.Get(ctx, &v4State)...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to parse v4 state during upgrade")
		return
	}

	// Transform v4 → v5
	v5State, diags := Transform(ctx, v4State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to transform v4 state to v5")
		return
	}

	// Write transformed state
	resp.Diagnostics.Append(resp.State.Set(ctx, v5State)...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to set upgraded v5 state")
		return
	}

	tflog.Info(ctx, "State upgrade from v4 to v5 completed successfully")
}

// UpgradeFromV5 handles state upgrades from v5 Plugin Framework provider (version=0) to v5 (version=500).
//
// This is a no-op upgrade since the schema is compatible - just bumps the version.
func UpgradeFromV5(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading load_balancer state from version=0 to version=500 (no-op)")

	// CRITICAL: For no-op upgrades, copy raw state directly
	// This preserves all state data without any transformation
	resp.State.Raw = req.State.Raw

	tflog.Info(ctx, "State version bump from 0 to 500 completed")
}

// UpgradeFromV1Ambiguous handles schema_version=1 state that is AMBIGUOUS between
// v4 SDKv2 format and v5 production format.
//
// Both v4 and v5 production (v5.0-v5.18 with GetSchemaVersion(1,500)) stored state
// at schema_version=1, but with incompatible formats:
//   - v4: adaptive_routing is an ARRAY (TypeList MaxItems:1), field is "default_pool_ids"
//   - v5: adaptive_routing is an OBJECT (SingleNestedAttribute), field is "default_pools"
//
// migrations.go registers this with PriorSchema=nil so the framework skips pre-decoding
// req.State entirely. Both paths operate exclusively on req.RawState.JSON.
//
// Detection: presence of "default_pool_ids" key in raw JSON unambiguously identifies v4.
func UpgradeFromV1Ambiguous(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading load_balancer state from schema_version=1 (detecting v4 vs v5 format)")

	isV4, err := detectV4LoadBalancerState(req)
	if err != nil {
		tflog.Warn(ctx, "Could not detect load_balancer state format, defaulting to no-op", map[string]interface{}{
			"error": err.Error(),
		})
		// Fall through to no-op path
	}

	if isV4 {
		tflog.Info(ctx, "Detected v4 load_balancer format (default_pool_ids present), performing transformation via raw state")
		upgradeFromV4ViaRawState(ctx, req, resp)
		return
	}

	// v5 production state: re-decode raw JSON with the target schema.
	// req.State is nil here (PriorSchema=nil), so we must use req.RawState directly.
	tflog.Info(ctx, "Detected v5 load_balancer format, performing no-op version bump via raw state")
	targetType := resp.State.Schema.Type().TerraformType(ctx)
	rawValue, unmarshalErr := req.RawState.Unmarshal(targetType)
	if unmarshalErr != nil {
		resp.Diagnostics.AddError(
			"Failed to unmarshal v5 load_balancer state",
			"Could not parse raw state as v5 format: "+unmarshalErr.Error(),
		)
		return
	}
	resp.State.Raw = rawValue
	tflog.Info(ctx, "State version bump from 1 to 500 completed")
}

// upgradeFromV4ViaRawState performs the v4->v5 transformation by parsing req.RawState
// directly with the v4 schema. Used when PriorSchema in migrations.go is nil
// (so req.State is nil and cannot be used to parse v4 arrays).
func upgradeFromV4ViaRawState(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	if req.RawState == nil {
		resp.Diagnostics.AddError("Missing raw state", "RawState was nil during v4 load_balancer upgrade")
		return
	}

	v4Schema := SourceCloudflareLoadBalancerSchema()
	v4Type := v4Schema.Type().TerraformType(ctx)

	rawValue, err := req.RawState.Unmarshal(v4Type)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to unmarshal v4 load_balancer state",
			"Could not parse raw state as v4 format: "+err.Error(),
		)
		return
	}

	syntheticState := &tfsdk.State{Raw: rawValue, Schema: v4Schema}
	syntheticReq := resource.UpgradeStateRequest{
		RawState: req.RawState,
		State:    syntheticState,
	}

	UpgradeFromV4(ctx, syntheticReq, resp)
}

// detectV4LoadBalancerState returns true if the raw state is in v4 format.
// v4 format has "default_pool_ids" (renamed to "default_pools" in v5).
func detectV4LoadBalancerState(req resource.UpgradeStateRequest) (bool, error) {
	if req.RawState == nil || len(req.RawState.JSON) == 0 {
		return false, nil
	}

	var raw map[string]interface{}
	if err := json.Unmarshal(req.RawState.JSON, &raw); err != nil {
		return false, err
	}

	_, hasV4Field := raw["default_pool_ids"]
	return hasV4Field, nil
}
