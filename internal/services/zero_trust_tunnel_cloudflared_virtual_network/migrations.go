package zero_trust_tunnel_cloudflared_virtual_network

import (
	"context"
	"os"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_tunnel_cloudflared_virtual_network/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithMoveState = (*ZeroTrustTunnelCloudflaredVirtualNetworkResource)(nil)
var _ resource.ResourceWithUpgradeState = (*ZeroTrustTunnelCloudflaredVirtualNetworkResource)(nil)

// MoveState registers state movers for resource renames.
// Handles both v4 resource type names:
//   - cloudflare_tunnel_virtual_network             (deprecated v4 name)
//   - cloudflare_zero_trust_tunnel_virtual_network  (preferred v4 name)
//
// This is triggered when users use the `moved` block (Terraform 1.8+).
func (r *ZeroTrustTunnelCloudflaredVirtualNetworkResource) MoveState(ctx context.Context) []resource.StateMover {
	sourceSchema := v500.SourceVirtualNetworkSchema()
	return []resource.StateMover{
		{
			SourceSchema: &sourceSchema,
			StateMover:   v500.MoveState,
		},
	}
}

// UpgradeState registers state upgraders for schema version changes.
//
// This handles two upgrade paths:
// 1. v4 state (schema_version=0) → v5 (version=500): Ensure defaults for comment and is_default_network
// 2. v5 state (version=1) → v5 (version=500): No-op upgrade (when TF_MIG_TEST=1)
//
// In production (no TF_MIG_TEST), only a no-op upgrader is registered at slot 0
// to safely bump existing v5 users from version 0 to 1 without triggering the
// v4→v5 transformation (which would fail on v5-format state).
func (r *ZeroTrustTunnelCloudflaredVirtualNetworkResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	tfMigTest := os.Getenv("TF_MIG_TEST")

	targetSchema := ResourceSchema(ctx)

	if tfMigTest == "" {
		return map[int64]resource.StateUpgrader{
			0: {
				PriorSchema: &targetSchema,
				StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
					resp.State.Raw = req.State.Raw
				},
			},
		}
	}

	sourceSchema := v500.SourceVirtualNetworkSchema()

	return map[int64]resource.StateUpgrader{
		// Handle state from v4 provider (schema_version=0)
		// This is used when users manually run terraform state mv (Terraform < 1.8)
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromV4,
		},

		// Handle state from v5 Plugin Framework provider with version=1
		// This is a no-op upgrade that just bumps the version to 500
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV5,
		},
	}
}
