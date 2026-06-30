package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromLegacyV0 handles state upgrades from the legacy cloudflare_d1_database resource (schema_version=0).
// This is triggered when the state schema version (0 from the v4 framework provider) is lower than
// the current v5 schema version (500).
//
// Key transformations:
//   - "id" is copied to "uuid" (v5 uses uuid for API calls)
//   - All new computed fields are initialized as null (refreshed from API on next plan/apply)
//   - All new optional fields are initialized as null
func UpgradeFromLegacyV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading d1_database state from legacy cloudflare_d1_database (schema_version=0)")

	// Parse the state using the v4 source schema
	var sourceState SourceCloudflareD1DatabaseModel
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

	tflog.Info(ctx, "State upgrade from legacy cloudflare_d1_database completed successfully")
}
