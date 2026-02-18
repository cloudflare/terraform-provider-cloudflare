package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// MoveArgoToTieredCaching handles moving state from legacy cloudflare_argo to cloudflare_argo_tiered_caching.
// This is triggered by Terraform 1.8+ when it encounters a `moved` block:
//
//   moved {
//     from = cloudflare_argo.example
//     to   = cloudflare_argo_tiered_caching.example
//   }
//
// This handler is called when:
// - Scenario 4: Only tiered_caching attribute exists (no smart_routing)
//
// For all other scenarios, the moved block points to cloudflare_argo_smart_routing instead.
func MoveArgoToTieredCaching(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
	tflog.Info(ctx, "Moving state from legacy cloudflare_argo to cloudflare_argo_tiered_caching")

	// Parse the source state (legacy v4 cloudflare_argo format)
	var sourceState SourceCloudflareArgoModel
	resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform to target (current v5 cloudflare_argo_tiered_caching format)
	targetState, diags := TransformArgoToTieredCaching(ctx, sourceState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set the moved state
	resp.Diagnostics.Append(resp.TargetState.Set(ctx, targetState)...)

	tflog.Info(ctx, "State move from legacy cloudflare_argo to cloudflare_argo_tiered_caching completed successfully")
}
