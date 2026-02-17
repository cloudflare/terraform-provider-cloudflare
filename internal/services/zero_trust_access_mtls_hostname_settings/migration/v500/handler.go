package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV0 handles state upgrades from v4 SDKv2 provider (schema_version=0) to v5 (version=500).
func UpgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zero_trust_access_mtls_hostname_settings state from v4 SDKv2 provider (schema_version=0)")

	var v4State SourceMTLSHostnameSettingsModel
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
	tflog.Info(ctx, "State upgrade from v4 to v5 completed successfully")
}

// UpgradeFromV1 handles state upgrades from v5 Plugin Framework provider (version=1) to v5 (version=500).
// No-op - just bumps the version.
func UpgradeFromV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zero_trust_access_mtls_hostname_settings state from version=1 to version=500 (no-op)")
	resp.State.Raw = req.State.Raw
	tflog.Info(ctx, "State version bump from 1 to 500 completed")
}
