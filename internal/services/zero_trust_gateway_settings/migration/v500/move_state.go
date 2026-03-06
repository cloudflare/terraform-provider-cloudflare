package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// MoveFromCloudflareTeamsAccount handles moving state from the legacy
// cloudflare_teams_account resource type to cloudflare_zero_trust_gateway_settings.
//
// This is triggered by Terraform 1.8+ when it processes a `moved` block:
//
//	moved {
//	  from = cloudflare_teams_account.example
//	  to   = cloudflare_zero_trust_gateway_settings.example
//	}
//
// The source state is the v4 cloudflare_teams_account format (schema_version=0).
// The transformation logic is shared with UpgradeFromV4.
func MoveFromCloudflareTeamsAccount(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
	tflog.Info(ctx, "Moving state from legacy cloudflare_teams_account to cloudflare_zero_trust_gateway_settings")

	// Parse the source state using the v4 source schema
	var sourceState SourceV4ZeroTrustGatewaySettingsModel
	resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform v4 → v5 (same logic as UpgradeFromV4)
	targetState, diags := Transform(ctx, sourceState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set the moved state
	resp.Diagnostics.Append(resp.TargetState.Set(ctx, targetState)...)

	tflog.Info(ctx, "State move from cloudflare_teams_account completed successfully")
}
