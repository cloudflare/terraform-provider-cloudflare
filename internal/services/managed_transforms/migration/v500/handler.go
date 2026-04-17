package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV0 handles state upgrades from schema_version=0 to v5 (version=500).
//
// There are two sources of schema_version=0 state:
//
//  1. State from v4 cloudflare_managed_headers moved via `terraform state mv`.
//     This state has the v4 schema (optional set nested blocks).
//
//  2. State from early v5 cloudflare_managed_transforms (versions 5.0.0–5.x.y before
//     the version was bumped to 500). These already have the correct v5 structure.
//
// We distinguish the two by attempting to unmarshal as the v5 target first. If that
// succeeds, the state is already correct — just bump the version. If it fails, fall back
// to the v4 source schema and run the full transformation.
func UpgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading managed_transforms state from schema_version=0")

	// Try to unmarshal as the v5 target (early v5 state that already has the right structure).
	var v5State TargetManagedTransformsModel
	v5Diags := req.State.Get(ctx, &v5State)
	if !v5Diags.HasError() && !v5State.ZoneID.IsNull() && !v5State.ZoneID.IsUnknown() {
		tflog.Info(ctx, "Upgrading managed_transforms state: detected early-v5 structure, bumping version (no-op)")
		resp.Diagnostics.Append(resp.State.Set(ctx, &v5State)...)
		return
	}

	// Fall back: unmarshal as v4 managed_headers source and transform.
	tflog.Info(ctx, "Upgrading managed_transforms state: detected v4 managed_headers structure, transforming")
	var v4State SourceManagedHeadersModel
	resp.Diagnostics.Append(req.State.Get(ctx, &v4State)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newV5State, diags := Transform(ctx, v4State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newV5State)...)
	tflog.Info(ctx, "State upgrade from v4 managed_headers completed successfully")
}

// UpgradeFromV1 handles state upgrades from v5 provider (version=1) to v5 (version=500).
//
// This is a no-op upgrade. Version 1 is the current v5 schema version.
func UpgradeFromV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading managed_transforms state from version=1 to version=500 (no-op)")

	resp.State.Raw = req.State.Raw

	tflog.Info(ctx, "State version bump from 1 to 500 completed")
}
