package v500

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// MoveState handles state moves from both v4 source resource types:
//   - cloudflare_tunnel             (deprecated v4 name)
//   - cloudflare_zero_trust_tunnel  (preferred v4 name)
func MoveState(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
	if !strings.Contains(req.SourceProviderAddress, "cloudflare/cloudflare") {
		return
	}

	if req.SourceTypeName != "cloudflare_tunnel" && req.SourceTypeName != "cloudflare_zero_trust_tunnel" {
		return
	}

	var source SourceTunnelCloudflaredModel
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
