package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV4 handles state upgrades from v4 SDKv2 provider (schema_version=0) to v5 (version=500).
//
// Transforms cloudflare_teams_account v4 state to cloudflare_zero_trust_gateway_settings v5 format:
//   - Flat boolean fields → nested under settings.*
//   - TypeList MaxItems:1 blocks → SingleNestedAttribute pointers under settings.*
//   - notification_settings: message → msg (field rename)
//   - logging, proxy, ssh_session_log, payload_log: dropped (handled by tf-migrate)
func UpgradeFromV4(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zero_trust_gateway_settings state from v4 SDKv2 provider (schema_version=0)")

	// Parse v4 state using v4 source model
	var v4State SourceV4ZeroTrustGatewaySettingsModel
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
// This is a no-op upgrade that just bumps the schema version. It is only triggered when
// TF_MIG_TEST=1 causes GetSchemaVersion to return 500 instead of 1.
func UpgradeFromV5(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zero_trust_gateway_settings state from version=1 to version=500 (no-op)")

	// CRITICAL: For no-op upgrades, copy raw state directly to preserve all state data
	resp.State.Raw = req.State.Raw

	tflog.Info(ctx, "State version bump from 1 to 500 completed")
}
