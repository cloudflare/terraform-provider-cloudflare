// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer_pool

import (
	"context"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/load_balancer_pool/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*LoadBalancerPoolResource)(nil)

// UpgradeState handles schema version upgrades for cloudflare_load_balancer_pool.
//
// Schema version history:
// - v4 (SDKv2): schema_version=0 (implicit, no explicit version set)
// - early v5 (v5.0–v5.7, before GetSchemaVersion was added): schema_version=0 with v5 (object) shape
// - v5 production (v5.0–v5.18 after stepping stone): schema_version=1
// - v5 current: schema_version=500
//
// Upgrade paths:
//
//  1. schema_version=0 → 500: AMBIGUOUS between
//     - v4 SDKv2 format (load_shedding/origin_steering as JSON arrays [{...}])
//     - early v5 format (load_shedding/origin_steering as JSON objects {...})
//     PriorSchema must be nil because the Plugin Framework rejects the state
//     pre-handler if the prior-schema shape doesn't match the raw JSON (see #7098).
//     The handler reads req.RawState.JSON directly and routes to the correct path.
//
//  2. schema_version=1 → 500: No-op
//     State is already in v5 format; just bumps the version number.
func (r *LoadBalancerPoolResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)

	return map[int64]resource.StateUpgrader{
		// Handle schema_version=0 state. AMBIGUOUS between v4 SDKv2 (list shape)
		// and early v5 (object shape). PriorSchema=nil so the handler can inspect
		// req.RawState.JSON itself and pick the right path. See #7098.
		0: {
			StateUpgrader: v500.UpgradeFromV0Ambiguous,
		},
		// Handle upgrades from v5 production state (schema_version=1).
		// Users on v5.0–v5.18 had GetSchemaVersion(1, 500) which stored state
		// at version 1. State is already in v5 format — no transformation needed.
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV0,
		},
	}
}
