package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV0 handles state upgrades from schema version 0 to version 500.
//
// v0 state comes from early v5 releases (v5.10, v5.11) where:
// - policies[].resources was a map[string]string
// - policies[].id existed (computed)
// - policies[].permission_groups had meta + name (computed)
//
// This performs a full transformation:
// - Converts resources from map to JSON-encoded string
// - Removes policy.id, permission_groups.meta, permission_groups.name
func UpgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading account_token state from v0 to v500")

	// Parse v0 state
	var v0State SourceAccountTokenModelV0
	resp.Diagnostics.Append(req.State.Get(ctx, &v0State)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform v0 → v500
	v500State, diags := Transform(ctx, v0State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write transformed state
	resp.Diagnostics.Append(resp.State.Set(ctx, v500State)...)
	tflog.Info(ctx, "State upgrade from v0 to v500 completed successfully")
}

// UpgradeFromV1 handles state upgrades from schema version 1 to version 500.
//
// This is a no-op upgrade — v1 schema is compatible with v500.
// v1 is the "dormant" state version before migration activation.
// When migrations are activated (schema version → 500), this handler
// copies the state as-is.
func UpgradeFromV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading account_token state from v1 to v500 (no-op)")

	// CRITICAL: For no-op upgrades, copy raw state directly.
	// This preserves all state data without any transformation.
	resp.State.Raw = req.State.Raw

	tflog.Info(ctx, "State version bump from 1 to 500 completed")
}
