// Package v500 implements state migration from legacy provider (v4) to current provider (v5)
// for the cloudflare_load_balancer_monitor resource.
package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV4 handles state upgrades from v4 SDKv2 provider (schema_version=0) to v5 (version=500).
//
// This performs a full transformation from v4 → v5 format, including:
// - Header field transformation: TypeSet (nested) → MapAttribute
// - Default value additions for fields with v5 defaults
// - Direct copy for compatible fields
//
// The v4 state has schema_version=0 (SDKv2 default), and we transform it to v5 format.
func UpgradeFromV4(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading load_balancer_monitor state from v4 SDKv2 provider (schema_version=0)")

	// Parse v4 state using v4 model with source schema
	var v4State SourceLoadBalancerMonitorModel
	resp.Diagnostics.Append(req.State.Get(ctx, &v4State)...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to parse v4 state")
		return
	}

	tflog.Debug(ctx, "Successfully parsed v4 state", map[string]interface{}{
		"account_id": v4State.AccountID.ValueString(),
		"type":       v4State.Type.ValueString(),
	})

	// Transform v4 → v5
	v5State, diags := Transform(ctx, &v4State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to transform v4 state to v5")
		return
	}

	// Write transformed state
	resp.Diagnostics.Append(resp.State.Set(ctx, v5State)...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to set v5 state")
		return
	}

	tflog.Info(ctx, "State upgrade from v4 to v5 completed successfully")
}

// UpgradeFromV5 handles state upgrades from v5 Plugin Framework provider (version=1) to v5 (version=500).
//
// This is a no-op upgrade since the schema is compatible - just bumps the version.
// This handler is only triggered.
//
// CRITICAL: For no-op upgrades, we copy raw state directly to preserve all data without transformation.
func UpgradeFromV5(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading load_balancer_monitor state from version=1 to version=500 (no-op)")

	// CRITICAL: For no-op upgrades, copy raw state directly
	// This preserves all state data without any transformation
	resp.State.Raw = req.State.Raw

	tflog.Info(ctx, "State version bump from 1 to 500 completed")
}
