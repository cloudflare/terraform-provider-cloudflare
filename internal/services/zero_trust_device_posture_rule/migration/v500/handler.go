package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV0 handles state upgrades from v4 SDKv2 provider (schema_version=0) to v5 (version=500).
//
// This performs a full transformation from v4 → v5 format.
// The v4 state has schema_version=0 (SDKv2 default), and we transform it to v5 format.
func UpgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zero_trust_device_posture_rule state from v4 SDKv2 provider (schema_version=0)")

	// Parse v4 state using v4 model
	var v4State SourceDevicePostureRuleModel
	resp.Diagnostics.Append(req.State.Get(ctx, &v4State)...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to parse v4 state for zero_trust_device_posture_rule",
			map[string]interface{}{
				"diagnostics": resp.Diagnostics.Errors(),
			})
		return
	}

	tflog.Debug(ctx, "Parsed v4 state successfully",
		map[string]interface{}{
			"name": v4State.Name.ValueString(),
			"type": v4State.Type.ValueString(),
		})

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

// UpgradeFromV1 handles state upgrades from v5 Plugin Framework provider (version=1) to v5 (version=500).
//
// This is a no-op upgrade since the schema is compatible - just bumps the version.
// This handler is only triggered (GetSchemaVersion returns 500).
func UpgradeFromV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zero_trust_device_posture_rule state from version=1 to version=500 (no-op)")

	// CRITICAL: For no-op upgrades, copy raw state directly
	// This preserves all state data without any transformation
	resp.State.Raw = req.State.Raw

	tflog.Info(ctx, "State version bump from 1 to 500 completed")
}
