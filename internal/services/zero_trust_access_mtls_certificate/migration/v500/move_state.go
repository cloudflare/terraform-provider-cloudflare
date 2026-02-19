package v500

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// MoveState handles state moves from cloudflare_access_mutual_tls_certificate to cloudflare_zero_trust_access_mtls_certificate.
// This is triggered when users use the `moved` block (Terraform 1.8+):
//
//	moved {
//	    from = cloudflare_access_mutual_tls_certificate.example
//	    to   = cloudflare_zero_trust_access_mtls_certificate.example
//	}
func MoveState(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
	if req.SourceTypeName != "cloudflare_access_mutual_tls_certificate" {
		return
	}

	if !isCloudflareProvider(req.SourceProviderAddress) {
		return
	}

	tflog.Info(ctx, "Starting state move from cloudflare_access_mutual_tls_certificate to cloudflare_zero_trust_access_mtls_certificate",
		map[string]interface{}{
			"source_type":           req.SourceTypeName,
			"source_schema_version": req.SourceSchemaVersion,
			"source_provider":       req.SourceProviderAddress,
		})

	var sourceState SourceAccessMutualTLSCertificateModel
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

	tflog.Info(ctx, "State move from cloudflare_access_mutual_tls_certificate completed successfully")
}

func isCloudflareProvider(addr string) bool {
	return strings.Contains(addr, "cloudflare/cloudflare") ||
		strings.Contains(addr, "registry.terraform.io/cloudflare")
}
