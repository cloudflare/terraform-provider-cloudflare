package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV0 handles state upgrades from v4 Framework provider (version=0) to v5 (version=500).
//
// This performs a transformation from v4 → v5 format, adding new fields with defaults.
// The v4 state has version=0 (Framework resource), and we transform it to v5 format.
func UpgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading r2_bucket state from v4 Framework provider (version=0) to v5 (version=500)")

	// Parse v4 state using v4 model
	var v4State SourceR2BucketModel
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

// UpgradeFromV500 handles state upgrades from v5 Plugin Framework provider (version=500) to v5 (version=500).
//
// This is a no-op upgrade since the schema is compatible - just a self-reference for consistency.
// This handler exists for completeness but should not be triggered in practice.
func UpgradeFromV500(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading r2_bucket state from version=500 to version=500 (no-op)")

	// CRITICAL: For no-op upgrades, copy raw state directly
	// This preserves all state data without any transformation
	resp.State.Raw = req.State.Raw

	tflog.Info(ctx, "State version bump from 500 to 500 completed (no changes)")
}
