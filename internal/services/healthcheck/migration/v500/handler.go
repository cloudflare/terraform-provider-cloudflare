package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV4 handles state upgrades from v4 SDKv2 provider (schema_version=0) to v5 (version=500).
//
// This performs a full transformation from v4 flat structure to v5 nested structure.
// The v4 state has schema_version=0 (SDKv2 default), and we transform it to v5 format with
// http_config or tcp_config nested objects based on the healthcheck type.
func UpgradeFromV4(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading healthcheck state from v4 SDKv2 provider (schema_version=0)")

	// Parse v4 state using v4 model (flat structure)
	var v4State SourceHealthcheckModel
	resp.Diagnostics.Append(req.State.Get(ctx, &v4State)...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to parse v4 healthcheck state")
		return
	}

	tflog.Debug(ctx, "Parsed v4 state", map[string]interface{}{
		"zone_id": v4State.ZoneID.ValueString(),
		"name":    v4State.Name.ValueString(),
		"type":    v4State.Type.ValueString(),
	})

	// Transform v4 → v5
	v5State, diags := Transform(ctx, v4State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to transform v4 healthcheck state to v5 format")
		return
	}

	// Write transformed state
	resp.Diagnostics.Append(resp.State.Set(ctx, v5State)...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to write v5 healthcheck state")
		return
	}

	tflog.Info(ctx, "State upgrade from v4 to v5 completed successfully")
}

// UpgradeFromV5 handles state upgrades from v5 Plugin Framework provider (version=1) to v5 (version=500).
//
// This is a no-op upgrade since the schema is compatible - just bumps the version.
// This handler is only triggered.
func UpgradeFromV5(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading healthcheck state from version=1 to version=500 (no-op)")

	// CRITICAL: For no-op upgrades, copy raw state directly
	// This preserves all state data without any transformation
	resp.State.Raw = req.State.Raw

	tflog.Info(ctx, "State version bump from 1 to 500 completed")
}
