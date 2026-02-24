// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_tunnel_cloudflared_config/migration/v500"
)

var _ resource.ResourceWithMoveState = (*ZeroTrustTunnelCloudflaredConfigResource)(nil)
var _ resource.ResourceWithUpgradeState = (*ZeroTrustTunnelCloudflaredConfigResource)(nil)

// MoveState handles moves from cloudflare_tunnel_config (deprecated v4 name) to
// cloudflare_zero_trust_tunnel_cloudflared_config (v5 resource type).
//
// This is triggered when tf-migrate generates the `moved` block (Terraform 1.8+):
//
//	moved {
//	    from = cloudflare_tunnel_config.example
//	    to   = cloudflare_zero_trust_tunnel_cloudflared_config.example
//	}
func (r *ZeroTrustTunnelCloudflaredConfigResource) MoveState(ctx context.Context) []resource.StateMover {
	sourceSchema := v500.SourceV4TunnelConfigSchema()
	return []resource.StateMover{
		{
			SourceSchema: &sourceSchema,
			StateMover:   v500.MoveStateFromTunnelConfig,
		},
	}
}

// UpgradeState handles schema version upgrades for cloudflare_zero_trust_tunnel_cloudflared_config.
//
// Version history:
//   - 0: v4 SDKv2 state (full transformation needed)
//   - 1: Dormant production v5 state (GetSchemaVersion returns 1 normally)
//   - 500: Active migration version (GetSchemaVersion returns 500 when TF_MIG_TEST=1)
func (r *ZeroTrustTunnelCloudflaredConfigResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	// v4 schema for version=0 upgrader
	v4Schema := v500.SourceV4TunnelConfigSchema()

	// v5 schema for version=1 upgrader (override version to match production state)
	v5SchemaVersion1 := ResourceSchema(ctx)
	v5SchemaVersion1.Version = 1

	return map[int64]resource.StateUpgrader{
		// Handle state from v4 SDKv2 provider (schema_version=0)
		// Performs full transformation: config/origin_request/access array→object,
		// ingress_rule→ingress rename, dropped fields, duration string→int64
		0: {
			PriorSchema:   &v4Schema,
			StateUpgrader: v500.UpgradeFromV4,
		},

		// Handle state from v5 Plugin Framework provider (version=1)
		// No-op: schema is compatible, just bumps version to 500
		1: {
			PriorSchema:   &v5SchemaVersion1,
			StateUpgrader: v500.UpgradeFromV5,
		},
	}
}
