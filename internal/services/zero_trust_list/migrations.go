package zero_trust_list

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_list/migration/v500"
)

var _ resource.ResourceWithUpgradeState = (*ZeroTrustListResource)(nil)
var _ resource.ResourceWithMoveState = (*ZeroTrustListResource)(nil)

// MoveState handles moves from cloudflare_teams_list (v4) to
// cloudflare_zero_trust_list (v5).
func (r *ZeroTrustListResource) MoveState(ctx context.Context) []resource.StateMover {
	sourceSchema := v500.SourceTeamsListSchema()
	return []resource.StateMover{
		{
			SourceSchema: &sourceSchema,
			StateMover:   v500.MoveState,
		},
	}
}

// UpgradeState registers state upgraders for schema version changes.
//
// v4 cloudflare_teams_list had schema_version=0. v5 uses GetSchemaVersion(2, 500).
//
// Upgrade paths:
// 0: v4 state (schema_version=0) → merge items + items_with_description
// 1: v5 state (version=1) → no-op
func (r *ZeroTrustListResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)

	sourceSchema := v500.SourceTeamsListSchema()
	return map[int64]resource.StateUpgrader{
		// v4 SDKv2 provider (schema_version=0) — merge items
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromV4,
		},
		// v5 state (version=1) — no-op
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV5V1,
		},
	}
}
