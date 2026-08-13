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
// Schema version history:
// - v4 SDKv2 (cloudflare_teams_list): schema_version=0
// - v5 production (v5.0-v5.18): schema_version=2 (GetSchemaVersion(2, 500) returned 2)
// - v5 current: schema_version=500
//
// Upgrade paths:
// 0: v4 SDKv2 state -> merge items + items_with_description
// 1: v5 intermediate state -> no-op
// 2: v5 production state (v5.0-v5.18) -> no-op
func (r *ZeroTrustListResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)
	sourceSchema := v500.SourceTeamsListSchema()

	return map[int64]resource.StateUpgrader{
		// v4 SDKv2 provider (schema_version=0) -- merge items + items_with_description
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromV4,
		},
		// v5 intermediate state (version=1) -- no-op
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV5V1,
		},
		// v5 production state (schema_version=2, from GetSchemaVersion(2, 500) in v5.0-v5.18).
		// State is already in v5 format -- no transformation needed.
		2: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV5V1,
		},
	}
}
