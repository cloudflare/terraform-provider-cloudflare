package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// MoveState handles moving state from legacy cloudflare_managed_headers to cloudflare_managed_transforms.
// This is triggered by Terraform 1.8+ when it encounters a `moved` block:
//
//	moved {
//	    from = cloudflare_managed_headers.example
//	    to   = cloudflare_managed_transforms.example
//	}
func MoveState(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
	tflog.Info(ctx, "Moving state from legacy cloudflare_managed_headers to cloudflare_managed_transforms")

	var sourceState SourceManagedHeadersModel
	resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	targetState, diags := Transform(ctx, sourceState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.TargetState.Set(ctx, targetState)...)
	tflog.Info(ctx, "State move from legacy cloudflare_managed_headers completed successfully")
}
