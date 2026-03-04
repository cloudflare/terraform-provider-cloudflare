package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV0 handles state upgrades from v4 provider (schema_version=0) to v5 (version=500).
//
// performs the full v4→v5 transformation. In production (no migration mode), a simple
// no-op is registered at slot 0 instead (see migrations.go).
//
// Transforms:
// - dns, origin_dns, edge_ips: array[0] → object (ListNestedBlock → SingleNestedAttribute)
// - origin_port_range: block with start/end → origin_port DynamicAttribute string
// - origin_port: integer → DynamicAttribute number wrapper
// - timestamps: string → timetypes.RFC3339
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
// This is a no-op upgrade. Version 1 is the "dormant" v5 state set in production
// this bumps from 1 to 500.
func UpgradeFromV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading spectrum_application state from version=1 to version=500 (no-op)")

	// CRITICAL: For no-op upgrades, copy raw state directly
	resp.State.Raw = req.State.Raw

	tflog.Info(ctx, "State version bump from 1 to 500 completed")
}
