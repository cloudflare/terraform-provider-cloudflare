// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v500

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV0 handles state upgrades from schema_version=1 to schema_version=500.
// This is a no-op upgrade since the schema is compatible - just copy state through.
//
// Why this exists: Terraform requires explicit upgraders to be defined for version tracking,
// even when the schema is identical. This ensures the schema_version is updated in the statefile.
func UpgradeFromV0(
	ctx context.Context,
	req resource.UpgradeStateRequest,
	resp *resource.UpgradeStateResponse,
) {
	tflog.Info(ctx, "Upgrading load_balancer_pool state from schema_version=1 (no-op for v5 same-version states)")
	// No-op upgrade: schema is compatible, just copy raw state through
	// We use the raw state value directly to avoid issues with custom field type serialization
	resp.State.Raw = req.State.Raw
}

// UpgradeFromV0Ambiguous handles schema_version=0 state that is AMBIGUOUS between
// the v4 SDKv2 provider format and early v5 (v5.0-v5.7) provider format.
//
// Both stored state at schema_version=0, but with incompatible JSON shapes:
//   - v4 (SDKv2): load_shedding, origin_steering, origins[].header as JSON arrays
//     (e.g. "load_shedding": [{"default_percent": 50}], single-element list)
//   - early v5: load_shedding, origin_steering as JSON objects (e.g. "load_shedding": {...})
//
// migrations.go registers this with PriorSchema=nil so the Plugin Framework
// skips the pre-handler unmarshal that would otherwise fail for one shape or
// the other. The handler then inspects req.RawState.JSON directly.
//
// Detection: try to unmarshal raw state with the TARGET schema. If it succeeds,
// the state is already v5-shaped → no-op. Otherwise, parse with the v4 source
// schema and run the full v4→v500 transform.
func UpgradeFromV0Ambiguous(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading load_balancer_pool state from schema_version=0 (detecting v4 vs early-v5 format)")

	if req.RawState == nil || len(req.RawState.JSON) == 0 {
		resp.Diagnostics.AddError(
			"Missing raw state",
			"RawState was nil or empty during load_balancer_pool schema_version=0 upgrade",
		)
		return
	}

	// First, try the target (current v5) schema. v5 state at schema_version=0
	// (early v5.0-v5.7) already has object-shaped nested attrs and will parse cleanly.
	targetType := resp.State.Schema.Type().TerraformType(ctx)
	if val, err := req.RawState.Unmarshal(targetType); err == nil {
		tflog.Info(ctx, "Detected early-v5 load_balancer_pool format - performing no-op upgrade via RawState")
		resp.State.Raw = val
		return
	} else {
		tflog.Debug(ctx, "Target schema unmarshal failed, falling back to v4 path", map[string]interface{}{
			"unmarshal_err": err.Error(),
		})
	}

	// Otherwise: parse with the v4 source schema and transform.
	upgradeFromV4ViaRawState(ctx, req, resp)
}

// upgradeFromV4ViaRawState performs the v4→v500 transformation by unmarshalling
// req.RawState with the v4 source schema and running Transform. Used when the
// upgrader was registered with PriorSchema=nil (req.State is empty).
func upgradeFromV4ViaRawState(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	v4Schema := SourceCloudflareLoadBalancerPoolSchema()
	v4Type := v4Schema.Type().TerraformType(ctx)

	rawValue, err := req.RawState.Unmarshal(v4Type)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to unmarshal v4 load_balancer_pool state",
			fmt.Sprintf("Could not parse raw state as v4 (SDKv2) format: %s", err),
		)
		return
	}

	syntheticState := tfsdk.State{Raw: rawValue, Schema: v4Schema}

	var sourceState SourceCloudflareLoadBalancerPoolModel
	resp.Diagnostics.Append(syntheticState.Get(ctx, &sourceState)...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to parse v4 state from raw value", map[string]interface{}{
			"diagnostics": resp.Diagnostics,
		})
		return
	}

	tflog.Debug(ctx, "Parsed v4 source state successfully", map[string]interface{}{
		"id":         sourceState.ID.ValueString(),
		"account_id": sourceState.AccountID.ValueString(),
		"name":       sourceState.Name.ValueString(),
	})

	targetState, diags := Transform(ctx, sourceState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to transform state", map[string]interface{}{
			"diagnostics": resp.Diagnostics,
		})
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, targetState)...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to set upgraded state", map[string]interface{}{
			"diagnostics": resp.Diagnostics,
		})
		return
	}

	tflog.Info(ctx, "State upgrade from v4 load_balancer_pool completed successfully")
}

// UpgradeFromLegacyV0 is kept for backwards compatibility with any external
// callers; new registrations should use UpgradeFromV0Ambiguous via migrations.go.
//
// Deprecated: use UpgradeFromV0Ambiguous, which works with PriorSchema=nil and
// handles both v4 SDKv2 and early-v5 schema_version=0 state correctly.
func UpgradeFromLegacyV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	UpgradeFromV0Ambiguous(ctx, req, resp)
}

// isV4LoadBalancerPoolState returns true when the raw state JSON has the v4
// (SDKv2) array shape for `load_shedding` or `origin_steering` rather than the
// v5 object shape. Useful for tests and diagnostics; the production path uses
// the target-schema unmarshal probe in UpgradeFromV0Ambiguous.
func isV4LoadBalancerPoolState(req resource.UpgradeStateRequest) (bool, error) {
	if req.RawState == nil || len(req.RawState.JSON) == 0 {
		return false, nil
	}
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(req.RawState.JSON, &raw); err != nil {
		return false, err
	}
	for _, field := range []string{"load_shedding", "origin_steering"} {
		v, ok := raw[field]
		if !ok || len(v) == 0 || v[0] == 'n' { // not present or null
			continue
		}
		if v[0] == '[' {
			return true, nil
		}
		if v[0] == '{' {
			return false, nil
		}
	}
	// Neither field present — be conservative: treat as v5 (no-op).
	return false, nil
}
