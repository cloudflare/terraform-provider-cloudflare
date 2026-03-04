package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// MoveState handles moving state from legacy resource to current resource.
// This is triggered by Terraform 1.8+ when it encounters a `moved` block:
//
//   moved {
//     from = cloudflare_device_posture_integration.example
//     to   = cloudflare_zero_trust_device_posture_integration.example
//   }
//
// For Terraform < 1.8, users must manually run:
//   terraform state mv cloudflare_device_posture_integration.example cloudflare_zero_trust_device_posture_integration.example
//
// which will preserve the source schema_version and trigger UpgradeFromV4 instead.
func MoveState(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
	tflog.Info(ctx, "Moving state from legacy cloudflare_device_posture_integration to cloudflare_zero_trust_device_posture_integration")

	// Parse the source state (legacy v4 format)
	var sourceState SourceCloudflareDevicePostureIntegrationModel
	resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform to target (current v5 format)
	targetState, diags := Transform(ctx, sourceState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set the moved state
	resp.Diagnostics.Append(resp.TargetState.Set(ctx, targetState)...)

	tflog.Info(ctx, "State move from legacy cloudflare_device_posture_integration completed successfully")
}
