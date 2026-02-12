package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// UpgradeFromV0 handles state upgrades from v0 to v500.
// This is triggered for resources that were created with schema version 0
// (either from the old cloudflare_access_group or early cloudflare_zero_trust_access_group).
func UpgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	// Parse source state (v4 SDKv2 format)
	var sourceState SourceV4ZeroTrustAccessGroupModel
	resp.Diagnostics.Append(req.State.Get(ctx, &sourceState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform to target state (v5 Plugin Framework format)
	targetState, diags := Transform(ctx, sourceState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set upgraded state
	resp.Diagnostics.Append(resp.State.Set(ctx, targetState)...)
}
