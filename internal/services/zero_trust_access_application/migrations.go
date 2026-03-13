package zero_trust_access_application

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_application/migration/v500"
)

var _ resource.ResourceWithMoveState = (*ZeroTrustAccessApplicationResource)(nil)
var _ resource.ResourceWithUpgradeState = (*ZeroTrustAccessApplicationResource)(nil)

// MoveState handles moves from cloudflare_access_application (v4) to cloudflare_zero_trust_access_application (v5).
// This is triggered when users use the `moved` block (Terraform 1.8+):
//
//	moved {
//	    from = cloudflare_access_application.example
//	    to   = cloudflare_zero_trust_access_application.example
//	}
func (r *ZeroTrustAccessApplicationResource) MoveState(ctx context.Context) []resource.StateMover {
	v4Schema := v500.SourceAccessApplicationSchema()
	return []resource.StateMover{
		{
			SourceSchema: &v4Schema,
			StateMover:   v500.MoveFromAccessApplication,
		},
	}
}

// UpgradeState handles schema version upgrades.
//
// State upgrade paths:
// 1. v4 state with schema_version=0 (from tf-migrate renaming type but not transforming attributes)
//   - Uses v4Schema to parse the state
//   - Performs full v4→v5 transformation (same as MoveFromAccessApplication)
//
// 2. Early v5 state with schema_version=1 (v5.12-v5.18)
//   - Uses v5Schema since state is already in v5 format
//   - No-op upgrade, just passes through
//
// Note: Early v5 (v5.12-v5.15) also had schema_version=0, but the resource type was
// cloudflare_zero_trust_access_application with v5-format attributes. However, tf-migrate
// produces state with the v5 type but v4-format attributes, so we use v4Schema for version 0.
func (r *ZeroTrustAccessApplicationResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {

	v4Schema := v500.SourceAccessApplicationSchema()
	v5Schema := ResourceSchema(ctx)
	return map[int64]resource.StateUpgrader{
		// Handle v4-format state with schema_version=0
		// This occurs when tf-migrate renames the resource type but doesn't transform attributes
		// Uses v4Schema to parse state, then performs full v4→v5 transformation
		0: {
			PriorSchema:   &v4Schema,
			StateUpgrader: v500.UpgradeFromV0,
		},
		// Handle upgrades from v5 with schema_version=1 (v5.12-v5.18)
		// This is a no-op since the schema is compatible.
		1: {
			PriorSchema:   &v5Schema,
			StateUpgrader: v500.UpgradeFromV1,
		},
	}
}
