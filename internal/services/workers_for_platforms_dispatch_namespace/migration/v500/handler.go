// File generated for StateUpgrader migration from v4 to v5

package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// UpgradeFromV0 handles the upgrade from v4 (version 0) to v5 (version 500).
// This handles both:
//   - cloudflare_workers_for_platforms_dispatch_namespace v4 state (upgrading provider directly)
//   - cloudflare_workers_for_platforms_namespace v4 state (when using terraform state mv instead of moved blocks)
func UpgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	// Step 1: Parse v4 state using source schema
	var v4State SourceWorkersForPlatformsNamespaceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &v4State)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Step 2: Transform v4 state to v5 state
	v5State, diags := Transform(ctx, v4State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Step 3: Set v5 state
	resp.Diagnostics.Append(resp.State.Set(ctx, v5State)...)
}

// UpgradeFromV1 handles the upgrade from v5 (version 1) to v5 (version 500).
// This is a no-op upgrade that just bumps the version number.
// Used for resources that were already on v5 provider before migration was enabled.
func UpgradeFromV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	// No transformation needed - just copy raw state through
	// This handles the case where state is already in v5 format but needs version bump
	resp.State.Raw = req.State.Raw
}
