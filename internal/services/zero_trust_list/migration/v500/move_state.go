package v500

import (
	"context"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/migrations"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// MoveState handles state moves from cloudflare_teams_list (v4)
// to cloudflare_zero_trust_list (v5).
func MoveState(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
	if req.SourceTypeName != "cloudflare_teams_list" {
		return
	}

	if !isCloudflareProvider(req.SourceProviderAddress) {
		return
	}

	tflog.Info(ctx, "Starting state move from cloudflare_teams_list to cloudflare_zero_trust_list",
		map[string]interface{}{
			"source_type":     req.SourceTypeName,
			"source_provider": req.SourceProviderAddress,
		})

	var sourceState SourceTeamsListModel
	if migrations.DiagnoseMoveStateNilSourceState(req, resp) {
		return
	}
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
	tflog.Info(ctx, "State move from cloudflare_teams_list to cloudflare_zero_trust_list completed")
}

func isCloudflareProvider(addr string) bool {
	return strings.Contains(addr, "cloudflare/cloudflare") ||
		strings.Contains(addr, "registry.terraform.io/cloudflare")
}
