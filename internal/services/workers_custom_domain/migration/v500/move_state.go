package v500

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// MoveStateV4toV500 handles moved blocks from cloudflare_worker_domain
// to cloudflare_workers_custom_domain.
func MoveStateV4toV500(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
	if req.SourceTypeName != "cloudflare_worker_domain" {
		return
	}

	if !isCloudflareProvider(req.SourceProviderAddress) {
		return
	}

	tflog.Info(ctx, "Starting state move from cloudflare_worker_domain to cloudflare_workers_custom_domain",
		map[string]interface{}{
			"source_type":           req.SourceTypeName,
			"source_schema_version": req.SourceSchemaVersion,
			"source_provider":       req.SourceProviderAddress,
		})

	var source SourceV4WorkersCustomDomainModel
	resp.Diagnostics.Append(req.SourceState.Get(ctx, &source)...)
	if resp.Diagnostics.HasError() {
		return
	}

	target, diags := TransformV4toV500(ctx, source)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.TargetState.Set(ctx, target)...)

	tflog.Info(ctx, "State move from cloudflare_worker_domain to cloudflare_workers_custom_domain completed")
}

func isCloudflareProvider(addr string) bool {
	return strings.Contains(addr, "cloudflare/cloudflare") ||
		strings.Contains(addr, "registry.terraform.io/cloudflare")
}
