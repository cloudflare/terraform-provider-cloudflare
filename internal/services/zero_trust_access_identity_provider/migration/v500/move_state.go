package v500

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// MoveState handles moved blocks from cloudflare_access_identity_provider
// to cloudflare_zero_trust_access_identity_provider (Terraform 1.8+).
func MoveState(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
	if req.SourceTypeName != "cloudflare_access_identity_provider" {
		return
	}
	if !isCloudflareProvider(req.SourceProviderAddress) {
		return
	}

	tflog.Info(ctx, "Moving state from cloudflare_access_identity_provider")

	var sourceState SourceAccessIdentityProviderModel
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
}

func isCloudflareProvider(addr string) bool {
	return strings.Contains(addr, "cloudflare/cloudflare") ||
		strings.Contains(addr, "registry.terraform.io/cloudflare")
}
