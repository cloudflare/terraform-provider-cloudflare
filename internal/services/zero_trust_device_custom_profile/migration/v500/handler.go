// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// UpgradeFromV4 upgrades state from v4 SDKv2 provider (schema_version=0) to v5 (version=500).
//
// This handles the full transformation from legacy device profile resources to the custom profile:
// - Resource type unchanged (already moved via MoveState)
// - Schema version: 0 → 500
// - Field transformations: See TransformToCustomProfile for details
func UpgradeFromV4(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	// Read the source state using the v4 schema
	var source SourceDeviceProfileModel
	diags := req.State.Get(ctx, &source)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform the source state to the target custom profile structure
	target, transformDiags := TransformToCustomProfile(ctx, source)
	resp.Diagnostics.Append(transformDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set the transformed state
	diags = resp.State.Set(ctx, target)
	resp.Diagnostics.Append(diags...)
}

// UpgradeFromV5 is a no-op upgrade from v5 (version=1) to v5 (version=500).
//
// This handles the case where:
// - User has already migrated to v5 provider (version=1, migrations dormant)
// - Now enabling migrations for testing (TF_MIG_TEST=1, version=500)
// - State structure is already correct, just bump the version
func UpgradeFromV5(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	// No transformation needed - state is already in v5 format
	// Just copy the state as-is to bump the version from 1 → 500
	resp.State.Raw = req.State.Raw
}
