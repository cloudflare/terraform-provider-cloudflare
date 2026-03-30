package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromLegacyV0 handles state upgrades from the legacy cloudflare_queue resource (schema_version=0).
// This is triggered when the state schema version (0 from SDKv2 provider) is lower than
// the current v5 schema version (500).
//
// Key transformations:
//   - "name" → "queue_name" (field rename)
//   - "id" → "queue_id" (id is also copied to the new queue_id field)
//   - All new computed fields are initialized as null (refreshed from API on next plan/apply)
func UpgradeFromLegacyV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading queue state from legacy cloudflare_queue (schema_version=0)")

	// Parse the state using the v4 source schema
	var sourceState SourceCloudflareQueueModel
	resp.Diagnostics.Append(req.State.Get(ctx, &sourceState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform to target
	targetState, diags := Transform(ctx, sourceState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set the upgraded state
	resp.Diagnostics.Append(resp.State.Set(ctx, targetState)...)

	tflog.Info(ctx, "State upgrade from legacy cloudflare_queue completed successfully")
}

// UpgradeFromV1 handles state upgrades from v5 production state (schema_version=1) to v5 (500).
//
// Users on v5.0–v5.18 had GetSchemaVersion(1, 500) which stored state at version 1.
// State is already in v5 format — no transformation needed, just bump the version.
func UpgradeFromV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading queue state from schema_version=1 to 500 (no-op)")
	resp.State.Raw = req.State.Raw
}
