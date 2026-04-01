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
// - v5 production (v5.0–v5.18): schema_version=1 (GetSchemaVersion(1, 500) returned 1)
// - v5 current: schema_version=500
//
// Upgrade paths:
// 1. v4 SDKv2 (schema_version=0) → v5 (500): Full transformation
//    - load_shedding: Array[0] → NestedObject
//    - origin_steering: Array[0] → NestedObject
//    - origins.header: Complex nested structure transformation
//    - check_regions: Set → List
//    - origins: Set → List
//
// 2. v5 production (schema_version=1) → v5 (500): No-op
//    - State is already in v5 format; just bumps the version number.
func (r *LoadBalancerPoolResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)
	sourceSchema := v500.SourceCloudflareLoadBalancerPoolSchema()

	return map[int64]resource.StateUpgrader{
		// Handle upgrades from legacy v4 SDKv2 provider (schema_version=0).
		// UpgradeFromLegacyV0 detects v4 vs early-v5 format at runtime and
		// either transforms (v4) or passes through (early v5 at version 0).
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromLegacyV0,
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
