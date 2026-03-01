package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV4 handles state upgrades from the legacy cloudflare_tunnel_route resource (schema_version=0).
// This is triggered when users manually run `terraform state mv cloudflare_tunnel_route.x cloudflare_zero_trust_tunnel_cloudflared_route.x`
// (Terraform < 1.8), which preserves the source schema_version=0 from the legacy provider.
//
// Note: schema_version=0 was the only schema version of cloudflare_tunnel_route in the legacy (SDKv2) provider.
func UpgradeFromV4(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading tunnel route state from legacy cloudflare_tunnel_route (schema_version=0)")

	// Parse the state (schema_version=0, source resource type)
	var sourceState SourceTunnelRouteModel
	resp.Diagnostics.Append(req.State.Get(ctx, &sourceState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform to target
	targetState, diags := Transform(ctx, &sourceState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set the upgraded state
	resp.Diagnostics.Append(resp.State.Set(ctx, targetState)...)

	tflog.Info(ctx, "State upgrade from legacy cloudflare_tunnel_route completed successfully")
}

// UpgradeFromV5 handles state upgrades from earlier v5 versions (schema_version=1) to current v500.
// This is a no-op upgrade since the schema is compatible - just copy state through.
func UpgradeFromV5(
	ctx context.Context,
	req resource.UpgradeStateRequest,
	resp *resource.UpgradeStateResponse,
) {
	tflog.Info(ctx, "Upgrading tunnel route state from schema_version=1")
	// No-op upgrade: schema is compatible, just copy raw state through
	resp.State.Raw = req.State.Raw
}
