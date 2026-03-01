package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV1 handles state upgrades from earlier v5 versions (schema_version=1) to current v500.
// This is a no-op upgrade since the schema is compatible - just copy state through.
func UpgradeFromV1(
	ctx context.Context,
	req resource.UpgradeStateRequest,
	resp *resource.UpgradeStateResponse,
) {
	tflog.Info(ctx, "Upgrading zero trust device managed networks state from schema_version=1")
	// No-op upgrade: schema is compatible, just copy raw state through
	resp.State.Raw = req.State.Raw
}

// UpgradeFromLegacyV0 handles state upgrades from the legacy cloudflare_device_managed_networks resource to cloudflare_zero_trust_device_managed_networks.
// This is triggered when users manually run `terraform state mv cloudflare_device_managed_networks.x cloudflare_zero_trust_device_managed_networks.x`
// (Terraform < 1.8), which preserves the source schema_version=0 from the legacy provider.
//
// Note: schema_version=0 (implicit) was the schema version of cloudflare_device_managed_networks in the legacy (SDKv2) provider
// before it was renamed. The state structure matches SourceCloudflareDeviceManagedNetworksModel.
func UpgradeFromLegacyV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading state from legacy cloudflare_device_managed_networks (schema_version=0)")

	// Parse the state (schema_version=0, source resource type)
	var sourceState SourceCloudflareDeviceManagedNetworksModel
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

	tflog.Info(ctx, "State upgrade from legacy cloudflare_device_managed_networks completed successfully")
}
