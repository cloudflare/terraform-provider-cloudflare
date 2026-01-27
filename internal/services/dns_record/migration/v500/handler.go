package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV0 handles state upgrades from earlier v500 versions (schema_version=0) to current v500.
// This is a no-op upgrade since the schema is compatible - just copy state through.
//
// Why this exists: Terraform requires explicit upgraders to be defined for version tracking,
// even when the schema is identical. This ensures the schema_version is updated in the statefile.
func UpgradeFromV0(
	ctx context.Context,
	req resource.UpgradeStateRequest,
	resp *resource.UpgradeStateResponse,
) {
	// No-op upgrade: schema is compatible, just copy raw state through
	// We use the raw state value directly to avoid issues with custom field type serialization
	resp.State.Raw = req.State.Raw
}

// UpgradeFromV4 handles state upgrades from v5 releases with schema_version=4 to current v500.
// This is a no-op upgrade since we only changed the version number from 4 to 500.
// All v5.x.x releases up to 5.16.0 used schema_version=4 for cloudflare_dns_record.
func UpgradeFromV4(
	ctx context.Context,
	req resource.UpgradeStateRequest,
	resp *resource.UpgradeStateResponse,
) {
	tflog.Info(ctx, "Upgrading DNS record state from schema_version=4 to schema_version=500")
	// No-op upgrade: schema is identical, just copy raw state through
	// We use the raw state value directly to avoid issues with custom field type serialization
	resp.State.Raw = req.State.Raw
}

// UpgradeFromLegacyV3 handles state upgrades from the legacy cloudflare_record resource to cloudflare_dns_record.
// This is triggered when users manually run `terraform state mv cloudflare_record.x cloudflare_dns_record.x`
// (Terraform < 1.8), which preserves the source schema_version=3 from the legacy provider.
//
// Note: schema_version=3 was the final schema version of cloudflare_record in the legacy (SDKv2) provider
// before it was deprecated. The state structure matches SourceCloudflareRecordModel.
func UpgradeFromLegacyV3(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading DNS record state from legacy cloudflare_record (schema_version=3)")

	// Parse the state (schema_version=3, source resource type)
	var sourceState SourceCloudflareRecordModel
	resp.Diagnostics.Append(req.State.Get(ctx, &sourceState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform to target
	targetState, diags := Transform(ctx, sourceState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set the upgraded state
	resp.Diagnostics.Append(resp.State.Set(ctx, targetState)...)

	tflog.Info(ctx, "State upgrade from legacy cloudflare_record completed successfully")
}
