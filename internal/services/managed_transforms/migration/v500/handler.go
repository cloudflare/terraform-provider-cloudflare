package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV0 handles state upgrades from v4 provider (schema_version=0) to v5 (version=500).
//
// The v4 resource cloudflare_managed_headers used schema_version=0.
// This is triggered when users manually run:
//
//	terraform state mv cloudflare_managed_headers.x cloudflare_managed_transforms.x
//
// which preserves the source schema_version=0.
func UpgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading managed_transforms state from v4 managed_headers (schema_version=0)")

	var v4State SourceManagedHeadersModel
	resp.Diagnostics.Append(req.State.Get(ctx, &v4State)...)
	if resp.Diagnostics.HasError() {
		return
	}

	v5State, diags := Transform(ctx, v4State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, v5State)...)
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
