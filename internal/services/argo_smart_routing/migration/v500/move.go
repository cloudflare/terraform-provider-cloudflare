package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// MoveArgoToSmartRouting handles moving state from legacy cloudflare_argo to cloudflare_argo_smart_routing.
// This is triggered by Terraform 1.8+ when it encounters a `moved` block:
//
//   moved {
//     from = cloudflare_argo.example
//     to   = cloudflare_argo_smart_routing.example
//   }
//
// This handler is called when:
// - Scenario 1: Neither smart_routing nor tiered_caching attributes exist
// - Scenario 2: Both smart_routing and tiered_caching exist (smart_routing is primary)
// - Scenario 3: Only smart_routing attribute exists
//
// For Scenario 4 (only tiered_caching), the moved block points to cloudflare_argo_tiered_caching instead.
func MoveArgoToSmartRouting(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
	tflog.Info(ctx, "Moving state from legacy cloudflare_argo to cloudflare_argo_smart_routing")

	// Parse the source state (legacy v4 cloudflare_argo format)
	var sourceState SourceCloudflareArgoModel
	resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform to target (current v5 cloudflare_argo_smart_routing format)
	targetState, diags := TransformArgoToSmartRouting(ctx, sourceState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set the moved state
	resp.Diagnostics.Append(resp.TargetState.Set(ctx, targetState)...)

	tflog.Info(ctx, "State move from legacy cloudflare_argo to cloudflare_argo_smart_routing completed successfully")
}
