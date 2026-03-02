package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV0 handles state upgrades from early v5 provider (schema_version=0) to v5 (version=500).
//
// cloudflare_zone_subscription is a new v5 resource with no v4 counterpart (zone subscriptions
// were previously managed via the `plan` attribute on cloudflare_zone). All state at version 0
// is already in v5 format, so this is a no-op that just bumps the schema version.
func UpgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zone_subscription state from version=0 to version=500 (no-op)")

	// CRITICAL: For no-op upgrades, copy raw state directly.
	// This preserves all state data without any transformation.
	resp.State.Raw = req.State.Raw

	tflog.Info(ctx, "State version bump from 0 to 500 completed")
}

// UpgradeFromV5 handles state upgrades from v5 Plugin Framework provider (version=1) to v5 (version=500).
//
// This is a no-op upgrade since the schema is compatible - just bumps the version.
// This handler is only triggered.
func UpgradeFromV5(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zone_subscription state from version=1 to version=500 (no-op)")

	// CRITICAL: For no-op upgrades, copy raw state directly.
	// This preserves all state data without any transformation.
	resp.State.Raw = req.State.Raw

	tflog.Info(ctx, "State version bump from 1 to 500 completed")
}
