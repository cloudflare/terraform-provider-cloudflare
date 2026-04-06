package v500

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV0Ambiguous handles schema_version=0 state that is AMBIGUOUS between
// v4 SDKv2 format and v5 production format.
//
// Both v4 and early v5 production stored state at schema_version=0, but with
// incompatible formats:
//   - v4: has "prefix_id" and "advertisement" fields (SDKv2 provider)
//   - v5: has "asn", "cidr", computed fields; no "prefix_id" or "advertisement"
//
// migrations.go registers this with PriorSchema=nil so the framework skips pre-decoding
// req.State entirely. Both paths operate exclusively on req.RawState.JSON.
//
// Detection: presence of "prefix_id" in raw JSON unambiguously identifies v4.
func UpgradeFromV0Ambiguous(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading cloudflare_byo_ip_prefix state from schema_version=0 (detecting v4 vs v5 format)")

	isV4, err := detectV4ByoIPPrefixState(req)
	if err != nil {
		tflog.Warn(ctx, "Could not detect byo_ip_prefix state format, defaulting to no-op", map[string]interface{}{
			"error": err.Error(),
		})
		// Fall through to no-op path
	}

	if isV4 {
		tflog.Info(ctx, "Detected v4 byo_ip_prefix format (prefix_id present), performing transformation via raw state")
		upgradeFromV4ViaRawState(ctx, req, resp)
		return
	}

	// v5 production state: re-decode raw JSON with the target schema.
	// req.State is nil here (PriorSchema=nil), so we must use req.RawState directly.
	tflog.Info(ctx, "Detected v5 byo_ip_prefix format, performing no-op version bump via raw state")
	targetType := resp.State.Schema.Type().TerraformType(ctx)
	rawValue, unmarshalErr := req.RawState.Unmarshal(targetType)
	if unmarshalErr != nil {
		resp.Diagnostics.AddError(
			"Failed to unmarshal v5 byo_ip_prefix state",
			"Could not parse raw state as v5 format: "+unmarshalErr.Error(),
		)
		return
	}
	resp.State.Raw = rawValue
	tflog.Info(ctx, "State version bump from 0 to 500 completed")
}

// upgradeFromV4ViaRawState performs the v4->v5 transformation by parsing req.RawState
// directly with the v4 schema. Used when PriorSchema in migrations.go is nil
// (so req.State is nil and cannot be used).
func upgradeFromV4ViaRawState(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	if req.RawState == nil {
		resp.Diagnostics.AddError("Missing raw state", "RawState was nil during v4 byo_ip_prefix upgrade")
		return
	}

	v4Schema := SourceCloudflareByoIPPrefixSchema()
	v4Type := v4Schema.Type().TerraformType(ctx)

	rawValue, err := req.RawState.Unmarshal(v4Type)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to unmarshal v4 byo_ip_prefix state",
			"Could not parse raw state as v4 format: "+err.Error(),
		)
		return
	}

	syntheticState := &tfsdk.State{Raw: rawValue, Schema: v4Schema}
	syntheticReq := resource.UpgradeStateRequest{
		RawState: req.RawState,
		State:    syntheticState,
	}

	UpgradeFromV0(ctx, syntheticReq, resp)
}

// UpgradeFromV0 handles state upgrades from v4 SDKv2 provider (schema_version=0) to v5 (version=500).
//
// This is triggered when req.State has been populated with the v4 source schema
// (either directly or via upgradeFromV4ViaRawState).
func UpgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading cloudflare_byo_ip_prefix state from schema_version=0 (v4) to schema_version=500 (v5)")

	// Parse the v4 state using the source schema
	var sourceState SourceCloudflareByoIPPrefixModel
	resp.Diagnostics.Append(req.State.Get(ctx, &sourceState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform to v5 state
	targetState, diags := Transform(ctx, sourceState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set the upgraded state
	resp.Diagnostics.Append(resp.State.Set(ctx, targetState)...)

	tflog.Info(ctx, "cloudflare_byo_ip_prefix state upgrade from schema_version=0 completed successfully")
}

// detectV4ByoIPPrefixState returns true if the raw state is in v4 format.
// v4 format has "prefix_id" (renamed to "id" in v5 and dropped as a separate field).
func detectV4ByoIPPrefixState(req resource.UpgradeStateRequest) (bool, error) {
	if req.RawState == nil || len(req.RawState.JSON) == 0 {
		return false, nil
	}

	var raw map[string]interface{}
	if err := json.Unmarshal(req.RawState.JSON, &raw); err != nil {
		return false, err
	}

	_, hasV4Field := raw["prefix_id"]
	return hasV4Field, nil
}
