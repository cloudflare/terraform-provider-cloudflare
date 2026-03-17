package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV4 handles state upgrades from v4 SDKv2 provider (schema_version=0) to v5 (version=500).
//
// This performs a full transformation from v4 → v5 format.
// The v4 state has schema_version=0 (implicit, no Version field), and we transform it to v5 format.
//
// Key transformations:
// - enforce_twofactor: top-level bool → settings.enforce_twofactor (wrapped in nested object)
// - settings: NEW nested object (set from enforce_twofactor + null abuse_contact_email)
// - unit, managed_by, created_on: NEW computed fields (set to null, will be refreshed from API)
func UpgradeFromV4(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading account state from v4 provider (schema_version=0)")

	// Parse v4 state using v4 model
	var v4State SourceCloudflareAccountModel
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

// UpgradeFromV5 handles state upgrades from v5 Plugin Framework provider (version=1) to v5 (version=500).
//
// This is a no-op upgrade since the schema is compatible - just bumps the version.
// This handler is only triggered when TF_MIG_TEST=1 (GetSchemaVersion returns 500).
//
// The v5 state with version=1 is the "dormant" state before migration activation.
// When migrations are activated, this handler copies the state as-is to version=500.
func UpgradeFromV5(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading account state from version=1 to version=500 (no-op)")

	// CRITICAL: For no-op upgrades, copy raw state directly
	// This preserves all state data without any transformation
	resp.State.Raw = req.State.Raw

	tflog.Info(ctx, "State version bump from 1 to 500 completed")
}
