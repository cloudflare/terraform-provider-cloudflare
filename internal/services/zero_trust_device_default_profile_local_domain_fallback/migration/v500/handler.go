package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// UpgradeFromV4 handles state migration from v4 (schema_version=0) to v5 (version=500).
//
// This is called when Terraform detects state with schema_version=0 (SDKv2 v4 provider)
// and needs to upgrade to version=500 (Plugin Framework v5 provider).
//
// The handler:
// 1. Parses v4 state using the source schema (SourceCloudflareFallbackDomainSchema)
// 2. Transforms the data to v5 structure (TransformToDefaultProfile)
// 3. Writes the transformed state back to Terraform
//
// Note: This handles resources WITHOUT policy_id (default profile path).
func UpgradeFromV4(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	// Parse v4 state using source schema
	var v4State SourceCloudflareFallbackDomainModel
	resp.Diagnostics.Append(req.State.Get(ctx, &v4State)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform v4 → v5 structure
	v5State, diags := TransformToDefaultProfile(ctx, v4State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write transformed state
	resp.Diagnostics.Append(resp.State.Set(ctx, v5State)...)
}

// UpgradeFromV5 handles state migration from v5 (version=1) to v5 (version=500).
//
// This is a no-op upgrade that just bumps the schema version when TF_MIG_TEST=1.
// It's needed because:
// - Production v5 resources use Version: 1
// - Migration testing requires Version: 500 to trigger migrations
// - We need a path to bump 1 → 500 without changing state structure
//
// The handler simply copies the raw state through without modification.
func UpgradeFromV5(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	// No-op: State structure is identical between v5 version=1 and version=500
	// Just copy the raw state through (version gets bumped automatically)
	resp.State.Raw = req.State.Raw
}
