// File generated for StateUpgrader migration from v4 to v5

package v500

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// MoveFromWorkersForPlatformsNamespace handles state moves from
// cloudflare_workers_for_platforms_namespace (deprecated v4) to
// cloudflare_workers_for_platforms_dispatch_namespace (v5).
// This is triggered when users use the `moved` block (Terraform 1.8+):
//
//	moved {
//	    from = cloudflare_workers_for_platforms_namespace.example
//	    to   = cloudflare_workers_for_platforms_dispatch_namespace.example
//	}
func MoveFromWorkersForPlatformsNamespace(
	ctx context.Context,
	req resource.MoveStateRequest,
	resp *resource.MoveStateResponse,
) {
	// Verify source is cloudflare_workers_for_platforms_namespace from cloudflare provider
	if req.SourceTypeName != "cloudflare_workers_for_platforms_namespace" {
		return
	}

	if !isCloudflareProvider(req.SourceProviderAddress) {
		return
	}

	tflog.Info(ctx, "Starting state move from cloudflare_workers_for_platforms_namespace to cloudflare_workers_for_platforms_dispatch_namespace",
		map[string]interface{}{
			"source_type":           req.SourceTypeName,
			"source_schema_version": req.SourceSchemaVersion,
			"source_provider":       req.SourceProviderAddress,
		})

	// Parse the source state
	var sourceState SourceWorkersForPlatformsNamespaceModel
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

	tflog.Info(ctx, "State move from cloudflare_workers_for_platforms_namespace to cloudflare_workers_for_platforms_dispatch_namespace completed successfully")
}

func isCloudflareProvider(addr string) bool {
	return strings.Contains(addr, "cloudflare/cloudflare") ||
		strings.Contains(addr, "registry.terraform.io/cloudflare")
}
