package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV1 handles state upgrades from v5 production version (schema_version=1) to v500.
// This is a no-op upgrade since the schema is compatible - just copy state through.
// This upgrader is only active when TF_MIG_TEST=1 sets schema version to 500.
func UpgradeFromV1(
	ctx context.Context,
	req resource.UpgradeStateRequest,
	resp *resource.UpgradeStateResponse,
) {
	tflog.Info(ctx, "Upgrading zero_trust_organization state from schema_version=1 to v500 (no-op)")
	// No-op upgrade: schema is compatible, just copy raw state through
	resp.State.Raw = req.State.Raw
}

// UpgradeFromLegacyV0 handles state upgrades from the legacy v4 resources to v5.
// This is triggered when users manually run state mv (Terraform < 1.8), which preserves
// the source schema_version=0 from the legacy provider.
//
// Handles BOTH v4 resource names (they share identical schemas):
//   - cloudflare_access_organization
//   - cloudflare_zero_trust_access_organization
//
// Note: schema_version=0 was the schema version in the v4 (SDKv2) provider.
// The state structure matches SourceCloudflareAccessOrganizationModel.
func UpgradeFromLegacyV0(
	ctx context.Context,
	req resource.UpgradeStateRequest,
	resp *resource.UpgradeStateResponse,
) {
	tflog.Info(ctx, "Upgrading zero_trust_organization state from legacy v4 (schema_version=0)")

	// Parse the v4 state (schema_version=0)
	var sourceState SourceCloudflareAccessOrganizationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &sourceState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform to v5 target
	targetState, diags := Transform(ctx, sourceState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set the upgraded state
	resp.Diagnostics.Append(resp.State.Set(ctx, targetState)...)

	tflog.Info(ctx, "State upgrade from legacy v4 completed successfully")
}
