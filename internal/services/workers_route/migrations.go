package workers_route

import (
	"context"
	"os"

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
// v5 uses GetSchemaVersion(2, 500). Upgrade paths:
//
// Version 0 (both V4 and V5 used version 0):
//   - V4 state: has "script_name" → rename to "script"
//   - V5 state: has "script" → strip V4-only fields, pass through
//   - Detection: V4 has "script_name" (non-null), V5 does not
//   - Union PriorSchema includes both "script_name" and "script" as optional strings
//
// Version 1:
//   - V5 state after initial release → no-op
func (r *WorkersRouteResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	if os.Getenv("TF_MIG_TEST") == "" {
		// Production mode: preserve existing pass-through upgraders
		targetSchema := ResourceSchema(ctx)
		return map[int64]resource.StateUpgrader{
			0: {
				PriorSchema: &targetSchema,
				StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
					resp.State.Raw = req.State.Raw
				},
			},
			1: {
				PriorSchema: &targetSchema,
				StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
					resp.State.Raw = req.State.Raw
				},
			},
		}
	}

	// Test mode (TF_MIG_TEST=1): full StateUpgrader migration
	unionSchema := v500.UnionV0Schema()
	targetSchema := ResourceSchema(ctx)
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema:   &unionSchema,
			StateUpgrader: v500.UpgradeFromV0,
		},
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV1,
		},
	}
}
