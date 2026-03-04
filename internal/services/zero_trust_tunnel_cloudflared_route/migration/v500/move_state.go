package v500

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// MoveState handles state moves from legacy tunnel route resources to cloudflare_zero_trust_tunnel_cloudflared_route.
// This is triggered when users use the `moved` block (Terraform 1.8+):
//
//	moved {
//	    from = cloudflare_tunnel_route.example
//	    to   = cloudflare_zero_trust_tunnel_cloudflared_route.example
//	}
//
// For Terraform < 1.8, users must manually run:
//
//	terraform state mv cloudflare_tunnel_route.example cloudflare_zero_trust_tunnel_cloudflared_route.example
//
// which will preserve the source schema_version and trigger UpgradeFromV4 instead.
//
// Handles both v4 source types:
//   - cloudflare_tunnel_route (deprecated v4 name)
//   - cloudflare_zero_trust_tunnel_route (preferred v4 name)
func MoveState(
	ctx context.Context,
	req resource.MoveStateRequest,
	resp *resource.MoveStateResponse,
) {
	// Only handle cloudflare_tunnel_route or cloudflare_zero_trust_tunnel_route sources
	if req.SourceTypeName != "cloudflare_tunnel_route" && req.SourceTypeName != "cloudflare_zero_trust_tunnel_route" {
		return
	}

	if !isCloudflareProvider(req.SourceProviderAddress) {
		return
	}

	tflog.Info(ctx, "Starting state move to cloudflare_zero_trust_tunnel_cloudflared_route",
		map[string]interface{}{
			"source_type":           req.SourceTypeName,
			"source_schema_version": req.SourceSchemaVersion,
			"source_provider":       req.SourceProviderAddress,
		})

	// Parse the source state
	var sourceState SourceTunnelRouteModel
	resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform source state to target state
	targetState, diags := Transform(ctx, &sourceState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set the target state
	resp.Diagnostics.Append(resp.TargetState.Set(ctx, targetState)...)

	tflog.Info(ctx, "State move to cloudflare_zero_trust_tunnel_cloudflared_route completed successfully")
}

func isCloudflareProvider(addr string) bool {
	return strings.Contains(addr, "cloudflare/cloudflare") ||
		strings.Contains(addr, "registry.terraform.io/cloudflare")
}
