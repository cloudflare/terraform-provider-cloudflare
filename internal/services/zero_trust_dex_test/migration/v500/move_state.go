package v500

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// MoveState handles state moves from cloudflare_device_dex_test (v0) to cloudflare_zero_trust_dex_test (v500).
// This is triggered when users use the `moved` block (Terraform 1.8+):
//
//	moved {
//	    from = cloudflare_device_dex_test.example
//	    to   = cloudflare_zero_trust_dex_test.example
//	}
//
// For Terraform < 1.8, users must manually run:
//
//	terraform state mv cloudflare_device_dex_test.example cloudflare_zero_trust_dex_test.example
//
// which will preserve the source schema_version and trigger UpgradeFromV4 instead.
func MoveState(
	ctx context.Context,
	req resource.MoveStateRequest,
	resp *resource.MoveStateResponse,
) {
	// Verify source is cloudflare_device_dex_test from cloudflare provider
	if req.SourceTypeName != "cloudflare_device_dex_test" {
		return
	}

	if !isCloudflareProvider(req.SourceProviderAddress) {
		return
	}

	tflog.Info(ctx, "Starting state move from cloudflare_device_dex_test to cloudflare_zero_trust_dex_test",
		map[string]interface{}{
			"source_type":           req.SourceTypeName,
			"source_schema_version": req.SourceSchemaVersion,
			"source_provider":       req.SourceProviderAddress,
		})

	// Parse the source state
	var sourceState SourceCloudflareDeviceDexTestModel
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

	tflog.Info(ctx, "State move from cloudflare_device_dex_test to cloudflare_zero_trust_dex_test completed successfully")
}

func isCloudflareProvider(addr string) bool {
	return strings.Contains(addr, "cloudflare/cloudflare") ||
		strings.Contains(addr, "registry.terraform.io/cloudflare")
}
