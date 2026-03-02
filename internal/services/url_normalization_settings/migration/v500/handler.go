package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV0 handles state upgrades from v4 SDKv2 provider (schema_version=0) to v5 (version=500).
//
// For url_normalization_settings, the schema is identical between v4 and v5
// (same fields, same types), so this is a no-op upgrade - just copy raw state through.
func UpgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading url_normalization_settings state from v4 SDKv2 provider (schema_version=0)")

	// No-op: schema is identical between v4 and v5, just copy raw state
	resp.State.Raw = req.State.Raw

	tflog.Info(ctx, "State upgrade from v4 to v5 completed successfully")
}

// UpgradeFromV1 handles state upgrades from v5 Plugin Framework provider (version=1) to v5 (version=500).
//
// This is a no-op upgrade since the schema is compatible - just bumps the version.
// This handler is only triggered.
func UpgradeFromV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading url_normalization_settings state from version=1 to version=500 (no-op)")

	// CRITICAL: For no-op upgrades, copy raw state directly
	resp.State.Raw = req.State.Raw

	tflog.Info(ctx, "State version bump from 1 to 500 completed")
}
