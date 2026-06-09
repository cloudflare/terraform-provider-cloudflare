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
	if req.SourceState == nil {
		resp.Diagnostics.AddError(
			"Unable to Read Source State",
			"The source state for "+req.SourceTypeName+" could not be decoded. "+
				"This typically occurs when the state file uses the legacy flatmap format "+
				"from Terraform versions prior to 0.12. Run 'terraform apply -refresh-only' "+
				"with the v4 provider to upgrade the state format, then retry the v5 migration.",
		)
		return
	}
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
