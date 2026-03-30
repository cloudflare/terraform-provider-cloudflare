package workers_route

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_route/migration/v500"
)

var _ resource.ResourceWithUpgradeState = (*WorkersRouteResource)(nil)
var _ resource.ResourceWithMoveState = (*WorkersRouteResource)(nil)

// MoveState handles moves from cloudflare_worker_route (v4 singular) to
// cloudflare_workers_route (v5 plural).
func (r *WorkersRouteResource) MoveState(ctx context.Context) []resource.StateMover {
	sourceSchema := v500.SourceWorkerRouteSchema()
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
// - v4 SDKv2: schema_version=0 (has "script_name" field)
// - v5 initial: schema_version=0 (has "script" field -- ambiguous with v4)
// - v5 production (v5.0-v5.18): schema_version=2 (GetSchemaVersion(2, 500) returned 2)
// - v5 current: schema_version=500
//
// Upgrade paths:
// 0: AMBIGUOUS -- v4 state (script_name) or early v5 state (script). Detected at runtime.
// 1: v5 state after initial release -- no-op
// 2: v5 production state (v5.0-v5.18) -- no-op
func (r *WorkersRouteResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	unionSchema := v500.UnionV0Schema()
	targetSchema := ResourceSchema(ctx)

	return map[int64]resource.StateUpgrader{
		// Version 0: ambiguous between v4 (script_name) and early v5 (script).
		// UpgradeFromV0 detects format at runtime via raw state inspection.
		0: {
			PriorSchema:   &unionSchema,
			StateUpgrader: v500.UpgradeFromV0,
		},
		// Version 1: v5 state after initial release -- no-op
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV1,
		},
		// Version 2: v5 production state (schema_version=2, from GetSchemaVersion(2, 500) in v5.0-v5.18).
		// State is already in v5 format -- no transformation needed.
		2: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV1,
		},
	}
}
