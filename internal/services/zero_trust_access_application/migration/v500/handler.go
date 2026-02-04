package v500

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV0 handles state upgrades from earlier v5 versions (schema_version=0) to current version.
// This is a no-op upgrade since the schema is compatible - just copy state through.
func UpgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zero_trust_access_application state from schema_version=0 (no-op)")

	// No-op upgrade: schema is compatible, just copy raw state through
	// We use the raw state value directly to avoid issues with custom field type serialization
	resp.State.Raw = req.State.Raw

	tflog.Info(ctx, "State upgrade from schema_version=0 completed")
}

// UpgradeFromV4 handles state upgrades from v4 cloudflare_access_application (schema_version=2) to v5.
// This is triggered when users run `terraform state mv cloudflare_access_application.x cloudflare_zero_trust_access_application.x`
// (Terraform < 1.8), which preserves the source schema_version=2 from the v4 provider.
func UpgradeFromV4(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zero_trust_access_application state from v4 cloudflare_access_application (schema_version=2)")

	// Parse the v4 state using the v4 model
	var v4State SourceAccessApplicationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &v4State)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform v4 → v5
	v5State, diags := Transform(ctx, v4State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write transformed state
	resp.Diagnostics.Append(resp.State.Set(ctx, v5State)...)

	tflog.Info(ctx, "State upgrade from v4 cloudflare_access_application completed successfully")
}

// MoveFromAccessApplication handles moves from cloudflare_access_application (v4) to cloudflare_zero_trust_access_application (v5).
// This is triggered when users use the `moved` block (Terraform 1.8+):
//
//	moved {
//	    from = cloudflare_access_application.example
//	    to   = cloudflare_zero_trust_access_application.example
//	}
func MoveFromAccessApplication(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
	// Verify source is cloudflare_access_application from cloudflare provider
	if req.SourceTypeName != "cloudflare_access_application" {
		return
	}

	if !isCloudflareProvider(req.SourceProviderAddress) {
		return
	}

	tflog.Info(ctx, "Starting state move from cloudflare_access_application to cloudflare_zero_trust_access_application",
		map[string]interface{}{
			"source_type":           req.SourceTypeName,
			"source_schema_version": req.SourceSchemaVersion,
			"source_provider":       req.SourceProviderAddress,
		})

	// Parse the v4 state using the v4 schema
	var v4State SourceAccessApplicationModel
	resp.Diagnostics.Append(req.SourceState.Get(ctx, &v4State)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform v4 state to v5 state
	v5State, diags := Transform(ctx, v4State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set the v5 state
	resp.Diagnostics.Append(resp.TargetState.Set(ctx, v5State)...)

	tflog.Info(ctx, "State move from cloudflare_access_application to cloudflare_zero_trust_access_application completed successfully")
}

// isCloudflareProvider checks if the provider address is the Cloudflare provider.
func isCloudflareProvider(addr string) bool {
	return strings.Contains(addr, "cloudflare/cloudflare") ||
		strings.Contains(addr, "registry.terraform.io/cloudflare")
}
