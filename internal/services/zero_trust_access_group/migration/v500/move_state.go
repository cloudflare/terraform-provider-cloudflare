package v500

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// MoveState handles state moves from cloudflare_access_group (v0) to cloudflare_zero_trust_access_group (v500).
// This is triggered when users use the `moved` block (Terraform 1.8+):
//
//	moved {
//	    from = cloudflare_access_group.example
//	    to   = cloudflare_zero_trust_access_group.example
//	}
func MoveState(
	ctx context.Context,
	req resource.MoveStateRequest,
	resp *resource.MoveStateResponse,
) {
	// Verify source is cloudflare_access_group from cloudflare provider
	if req.SourceTypeName != "cloudflare_access_group" {
		return
	}

	if !isCloudflareProvider(req.SourceProviderAddress) {
		return
	}

	tflog.Info(ctx, "Starting state move from cloudflare_access_group to cloudflare_zero_trust_access_group",
		map[string]interface{}{
			"source_type":           req.SourceTypeName,
			"source_schema_version": req.SourceSchemaVersion,
			"source_provider":       req.SourceProviderAddress,
		})

	// Parse the source state (v4 SDKv2 format)
	var sourceState SourceV4ZeroTrustAccessGroupModel
	resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform source state to target state (v5 Plugin Framework format)
	targetState, diags := Transform(ctx, sourceState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set the target state
	resp.Diagnostics.Append(resp.TargetState.Set(ctx, targetState)...)

	tflog.Info(ctx, "State move from cloudflare_access_group to cloudflare_zero_trust_access_group completed successfully")
}

func isCloudflareProvider(addr string) bool {
	return strings.Contains(addr, "cloudflare/cloudflare") ||
		strings.Contains(addr, "registry.terraform.io/cloudflare")
}
