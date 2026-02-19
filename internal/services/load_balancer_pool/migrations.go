// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer_pool

import (
	"context"
	"os"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/load_balancer_pool/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*LoadBalancerPoolResource)(nil)

// UpgradeState handles schema version upgrades for cloudflare_load_balancer_pool.
//
// Schema version history:
// - v4 (SDKv2): schema_version=0 (implicit, no explicit version set)
// - v5: schema_version=0→500 (controlled rollout via GetSchemaVersion)
//
// This handles migration from:
// 1. Legacy v4 SDKv2 provider (schema_version=0) → v5 Plugin Framework (schema_version=500)
//
// Key transformations:
// - load_shedding: Array[0] → NestedObject
// - origin_steering: Array[0] → NestedObject
// - origins.header: Complex nested structure transformation
// - check_regions: Set → List
// - origins: Set → List
//
// In production (no TF_MIG_TEST), only a no-op upgrader is registered at slot 0
// to safely bump existing v5 users from version 0 to 1 without triggering the
// v4→v5 transformation (which would fail on v5-format state).
func (r *LoadBalancerPoolResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)

	if os.Getenv("TF_MIG_TEST") == "" {
		return map[int64]resource.StateUpgrader{
			0: {
				PriorSchema: &targetSchema,
				StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
					resp.State.Raw = req.State.Raw
				},
			},
		}
	}

	sourceSchema := v500.SourceCloudflareLoadBalancerPoolSchema()

	return map[int64]resource.StateUpgrader{
		// Handle upgrades from legacy v4 SDKv2 provider (schema_version=0)
		// This is the actual state transformation that handles all the SDK v2 → Plugin Framework changes
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromLegacyV0,
		},
	}
}
