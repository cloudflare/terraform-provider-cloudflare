package v500

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV0 handles state upgrades from v4 provider (schema_version=0) to v5 (version=500).
//
// Performs the full v4→v5 transformation:
//   - dns, origin_dns, edge_ips: array[0] → object (ListNestedBlock → SingleNestedAttribute)
//   - origin_port_range: block with start/end → origin_port DynamicAttribute string
//   - origin_port: integer → DynamicAttribute number wrapper
//   - timestamps: string → timetypes.RFC3339
//
// IMPORTANT: this function requires req.State to have been populated with the
// v4 PriorSchema (i.e. it must NOT be called directly from a slot where
// PriorSchema=nil). Use UpgradeFromV0Ambiguous for the production code path —
// it builds a synthetic req.State from req.RawState and then defers here.
func UpgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading spectrum_application state from v4 provider (schema_version=0)")

	// Parse v4 state using source model
	var v4State SourceSpectrumApplicationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &v4State)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform v4 → v5
	v5State, diags := Transform(ctx, v4State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write transformed state
	resp.Diagnostics.Append(resp.State.Set(ctx, v5State)...)
	tflog.Info(ctx, "State upgrade from v4 to v5 completed successfully")
}

// UpgradeFromV1 handles state upgrades from v5 provider (version=1) to v5 (version=500).
//
// This is a no-op upgrade. Version 1 is the "dormant" v5 state set in production;
// this bumps from 1 to 500.
func UpgradeFromV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading spectrum_application state from version=1 to version=500 (no-op)")

	// CRITICAL: For no-op upgrades, copy raw state directly
	resp.State.Raw = req.State.Raw

	tflog.Info(ctx, "State version bump from 1 to 500 completed")
}

// UpgradeFromV0Ambiguous handles schema_version=0 state that is AMBIGUOUS between
// the v4 SDKv2 provider format and early v5 (v5.0-v5.7) provider format.
//
// Both stored state at schema_version=0, but with incompatible JSON shapes for
// nested blocks such as `dns`, `origin_dns`, `edge_ips`, and `origin_port_range`:
//   - v4 (SDKv2):  arrays of one element, e.g. `"dns": [{"name":"...","type":"..."}]`
//   - early v5:    single objects, e.g.       `"dns": {"name":"...","type":"..."}`
//
// migrations.go registers this with PriorSchema=nil so the Plugin Framework
// skips the pre-handler unmarshal that would otherwise reject one shape or the
// other. We probe the raw JSON ourselves:
//
//  1. Try to unmarshal req.RawState with the TARGET (current v5) schema.
//     If this succeeds, the state is already v5-shaped → no-op upgrade.
//  2. Otherwise, parse with the v4 source schema and run the full v4→v500
//     Transform.
func UpgradeFromV0Ambiguous(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading spectrum_application state from schema_version=0 (detecting v4 vs early-v5 format)")

	if req.RawState == nil || len(req.RawState.JSON) == 0 {
		resp.Diagnostics.AddError(
			"Missing raw state",
			"RawState was nil or empty during spectrum_application schema_version=0 upgrade",
		)
		return
	}

	// Probe with the target schema first; this is the cheap, accurate test.
	targetType := resp.State.Schema.Type().TerraformType(ctx)
	if val, err := req.RawState.Unmarshal(targetType); err == nil {
		tflog.Info(ctx, "Detected early-v5 spectrum_application format - performing no-op upgrade via RawState")
		resp.State.Raw = val
		return
	} else {
		tflog.Debug(ctx, "Target schema unmarshal failed, falling back to v4 path", map[string]interface{}{
			"unmarshal_err": err.Error(),
		})
	}

	// Fallback: parse as v4 (SDKv2) state and transform.
	v4Schema := SourceSpectrumApplicationSchema()
	v4Type := v4Schema.Type().TerraformType(ctx)

	rawValue, err := req.RawState.Unmarshal(v4Type)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to unmarshal v4 spectrum_application state",
			fmt.Sprintf("Could not parse raw state as v4 (SDKv2) format: %s", err),
		)
		return
	}

	syntheticState := tfsdk.State{Raw: rawValue, Schema: v4Schema}
	syntheticReq := resource.UpgradeStateRequest{
		RawState: req.RawState,
		State:    &syntheticState,
	}

	UpgradeFromV0(ctx, syntheticReq, resp)
}

// isV4SpectrumApplicationState returns true when the raw state JSON has the v4
// SDKv2 array shape for `dns`, `origin_dns`, `edge_ips`, or `origin_port_range`.
// Used by tests; production uses the target-schema probe in UpgradeFromV0Ambiguous.
func isV4SpectrumApplicationState(req resource.UpgradeStateRequest) (bool, error) {
	if req.RawState == nil || len(req.RawState.JSON) == 0 {
		return false, nil
	}
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(req.RawState.JSON, &raw); err != nil {
		return false, err
	}
	for _, field := range []string{"dns", "origin_dns", "edge_ips", "origin_port_range"} {
		v, ok := raw[field]
		if !ok || len(v) == 0 || v[0] == 'n' {
			continue
		}
		if v[0] == '[' {
			return true, nil
		}
		if v[0] == '{' {
			return false, nil
		}
	}
	return false, nil
}
