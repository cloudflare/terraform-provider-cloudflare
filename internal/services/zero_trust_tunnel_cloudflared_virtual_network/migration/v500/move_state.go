package v500

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// MoveState handles state moves from both v4 source resource types:
//   - cloudflare_tunnel_virtual_network             (deprecated v4 name)
//   - cloudflare_zero_trust_tunnel_virtual_network  (preferred v4 name)
func MoveState(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
	if !strings.Contains(req.SourceProviderAddress, "cloudflare/cloudflare") {
		return
	}

	if req.SourceTypeName != "cloudflare_tunnel_virtual_network" && req.SourceTypeName != "cloudflare_zero_trust_tunnel_virtual_network" {
		return
	}

	var source SourceVirtualNetworkModel
	resp.Diagnostics.Append(req.SourceState.Get(ctx, &source)...)
	if resp.Diagnostics.HasError() {
		return
	}

	target, diags := Transform(ctx, &source)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.TargetState.Set(ctx, target)...)
}
