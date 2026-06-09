package v500

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// MoveState handles state moves from cloudflare_dlp_profile or cloudflare_zero_trust_dlp_profile
// to cloudflare_zero_trust_dlp_custom_profile.
// This is triggered when users use the `moved` block (Terraform 1.8+):
//
//	moved {
//	    from = cloudflare_dlp_profile.example
//	    to   = cloudflare_zero_trust_dlp_custom_profile.example
//	}
func MoveState(
	ctx context.Context,
	req resource.MoveStateRequest,
	resp *resource.MoveStateResponse,
) {
	if req.SourceTypeName != "cloudflare_dlp_profile" &&
		req.SourceTypeName != "cloudflare_zero_trust_dlp_profile" {
		return
	}

	if !isCloudflareProvider(req.SourceProviderAddress) {
		return
	}

	tflog.Info(ctx, "Starting state move to cloudflare_zero_trust_dlp_custom_profile",
		map[string]interface{}{
			"source_type":           req.SourceTypeName,
			"source_schema_version": req.SourceSchemaVersion,
			"source_provider":       req.SourceProviderAddress,
		})

	var sourceState SourceCloudflareDLPProfileModel
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

	tflog.Info(ctx, "State move to cloudflare_zero_trust_dlp_custom_profile completed successfully")
}

func isCloudflareProvider(addr string) bool {
	return strings.Contains(addr, "cloudflare/cloudflare") ||
		strings.Contains(addr, "registry.terraform.io/cloudflare")
}
