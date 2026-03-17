package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV4 handles state upgrades from v4 SDKv2 provider (schema_version=0) to v5 (version=500).
//
// This performs a full transformation from v4 → v5 format.
// The v4 state has schema_version=0 (SDKv2 default), and we transform it to v5 format.
//
// For observatory_scheduled_test, this is a simple pass-through since all fields are direct copies.
func UpgradeFromV4(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading observatory_scheduled_test state from v4 SDKv2 provider (schema_version=0)")

	// Parse v4 state using v4 model
	var v4State SourceCloudflareObservatoryScheduledTestModel
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
// CRITICAL: For no-op upgrades, we copy raw state directly to preserve all data without transformation.
func UpgradeFromV5(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading observatory_scheduled_test state from version=1 to version=500 (no-op)")

	// No-op upgrade: schema is compatible, just copy raw state through
	// We use the raw state value directly to avoid issues with custom field type serialization
	resp.State.Raw = req.State.Raw

	tflog.Info(ctx, "State version bump from 1 to 500 completed")
}
