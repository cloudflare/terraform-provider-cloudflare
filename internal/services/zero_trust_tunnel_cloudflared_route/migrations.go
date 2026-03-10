package zero_trust_tunnel_cloudflared_route

import (
	"context"

	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_tunnel_cloudflared_route/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithMoveState = (*ZeroTrustTunnelCloudflaredRouteResource)(nil)
var _ resource.ResourceWithUpgradeState = (*ZeroTrustTunnelCloudflaredRouteResource)(nil)

// MoveState registers state movers for resource renames.
// Handles both v4 resource type names:
//   - cloudflare_tunnel_route (deprecated v4 name)
//   - cloudflare_zero_trust_tunnel_route (preferred v4 name)
//
// This is triggered when users use the `moved` block (Terraform 1.8+).
func (r *ZeroTrustTunnelCloudflaredRouteResource) MoveState(ctx context.Context) []resource.StateMover {
	sourceSchema := v500.SourceTunnelRouteSchema()
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
// 1. v4 state (schema_version=0) → v5 (version=500): Full transformation
// 2. v5 state (version=1) → v5 (version=500): No-op upgrade
func (r *ZeroTrustTunnelCloudflaredRouteResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)
	sourceSchema := v500.SourceTunnelRouteSchema()

	return map[int64]resource.StateUpgrader{
		// Handle state from v4 SDKv2 provider (schema_version=0)
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
