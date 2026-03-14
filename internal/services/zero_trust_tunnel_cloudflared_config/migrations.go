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
// This handles two upgrade paths:
// 1. v4 state (schema_version=0) → v5 (version=500): Full transformation
// 2. v5 state (version=1) → v5 (version=500): No-op upgrade
func (r *ZeroTrustTunnelCloudflaredConfigResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	// v5 schema for version=1 upgrader (override version to match production state)
	v5SchemaVersion1 := ResourceSchema(ctx)
	v5SchemaVersion1.Version = 1

	return map[int64]resource.StateUpgrader{
		// Handle BOTH version=0 formats:
		// - v4 SDKv2 state (config as list) -> full transform
		// - early v5 published state (config as object) -> no-op version bump (+ normalization)
		//
		// PriorSchema is nil to avoid brittle dual-shape schema decoding at framework boundary.
		// Dispatcher handles v5 decode locally and delegates v4 transform to migration package.
		0: {
			PriorSchema:   nil,
			StateUpgrader: upgradeFromV0,
		},

		// Handle state from v5 Plugin Framework provider (version=1)
		// No-op: schema is compatible, just bumps version to 500
		1: {
			PriorSchema:   &v5SchemaVersion1,
			StateUpgrader: v500.UpgradeFromV5,
		},
	}
}

func upgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	if req.RawState == nil {
		resp.Diagnostics.AddError("Missing raw state", "RawState was nil for schema version 0 migration")
		return
	}

	// Try decoding as early v5 (version=0) first.
	// If this succeeds, state is already v5-shaped and can be passed through.
	v5SchemaVersion0 := ResourceSchema(ctx)
	v5SchemaVersion0.Version = 0
	v5Type := v5SchemaVersion0.Type().TerraformType(ctx)
	v5RawValue, err := req.RawState.Unmarshal(v5Type)
	if err == nil {
		resp.State.Raw = v5RawValue
		return
	}

	// Not decodable as v5 version=0; treat as v4 SDKv2 and transform.
	v500.UpgradeFromV0(ctx, req, resp)
}
