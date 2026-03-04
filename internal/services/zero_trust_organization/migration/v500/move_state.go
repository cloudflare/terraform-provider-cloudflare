package v500

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// MoveState handles state moves from v4 resource names to v5 cloudflare_zero_trust_organization.
//
// Handles TWO v4 resource names (both were aliases with identical schemas):
//   - cloudflare_access_organization (deprecated v4 name)
//   - cloudflare_zero_trust_access_organization (current v4 name)
//
// This is triggered when users use the `moved` block (Terraform 1.8+):
//
//	moved {
//	    from = cloudflare_access_organization.example
//	    to   = cloudflare_zero_trust_organization.example
//	}
func MoveState(
	ctx context.Context,
	req resource.MoveStateRequest,
	resp *resource.MoveStateResponse,
) {
	// Verify source is one of the v4 resource names from cloudflare provider
	if req.SourceTypeName != "cloudflare_access_organization" &&
		req.SourceTypeName != "cloudflare_zero_trust_access_organization" {
		return
	}

	if !isCloudflareProvider(req.SourceProviderAddress) {
		return
	}

	tflog.Info(ctx, "Starting state move to cloudflare_zero_trust_organization",
		map[string]interface{}{
			"source_type":           req.SourceTypeName,
			"source_schema_version": req.SourceSchemaVersion,
			"source_provider":       req.SourceProviderAddress,
		})

	// Parse the source state using the v4 source model
	var sourceState SourceCloudflareAccessOrganizationModel
	resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform source state to target state
	targetState, diags := Transform(ctx, sourceState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set the target state
	resp.Diagnostics.Append(resp.TargetState.Set(ctx, targetState)...)

	tflog.Info(ctx, "State move to cloudflare_zero_trust_organization completed successfully",
		map[string]interface{}{
			"source_type": req.SourceTypeName,
		})
}

func isCloudflareProvider(addr string) bool {
	return strings.Contains(addr, "cloudflare/cloudflare") ||
		strings.Contains(addr, "registry.terraform.io/cloudflare")
}
