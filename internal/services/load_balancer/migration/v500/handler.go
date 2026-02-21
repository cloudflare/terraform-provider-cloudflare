package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
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

// UpgradeFromV5 handles state upgrades from v5 Plugin Framework provider (version=1) to v5 (version=500).
//
// This is a no-op upgrade since the schema is compatible - just bumps the version.
// This handler is only triggered when TF_MIG_TEST=1 (GetSchemaVersion returns 500).
//
// The version bump from 1 → 500 allows Terraform to know the state has been validated
// and is ready for the new schema version, even though no actual transformation is needed.
func UpgradeFromV5(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading load_balancer state from version=1 to version=500 (no-op)")

	// CRITICAL: For no-op upgrades, copy raw state directly
	// This preserves all state data without any transformation
	resp.State.Raw = req.State.Raw

	tflog.Info(ctx, "State version bump from 1 to 500 completed")
}
