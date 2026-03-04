package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeArgoToSmartRouting handles state upgrades from v4 SDKv2 cloudflare_argo (schema_version=0) to v5 cloudflare_argo_smart_routing (version=500).
//
// This performs a full transformation from v4 cloudflare_argo → v5 cloudflare_argo_smart_routing.
// The v4 state has schema_version=0 (SDKv2 default), and we transform it to v5 format.
func UpgradeArgoToSmartRouting(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading argo_smart_routing state from v4 cloudflare_argo (schema_version=0)")

	// Parse v4 state using v4 model
	var v4State SourceCloudflareArgoModel
	resp.Diagnostics.Append(req.State.Get(ctx, &v4State)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform v4 cloudflare_argo → v5 cloudflare_argo_smart_routing
	v5State, diags := TransformArgoToSmartRouting(ctx, v4State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write transformed state
	resp.Diagnostics.Append(resp.State.Set(ctx, v5State)...)
	tflog.Info(ctx, "State upgrade from v4 cloudflare_argo to v5 cloudflare_argo_smart_routing completed successfully")
}

// UpgradeFromV5SmartRouting handles state upgrades from v5 Plugin Framework provider (version=1) to v5 (version=500).
//
// This is a no-op upgrade since the schema is compatible - just bumps the version.
// This handler is only triggered.
func UpgradeFromV5SmartRouting(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading argo_smart_routing state from version=1 to version=500 (no-op)")

	// CRITICAL: For no-op upgrades, copy raw state directly
	// This preserves all state data without any transformation
	resp.State.Raw = req.State.Raw

	tflog.Info(ctx, "State version bump from 1 to 500 completed")
}
