// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer

import (
	"context"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/load_balancer/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*LoadBalancerResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
//
// Schema version history:
// - v4 SDKv2: schema_version=1
// - v5 production (v5.0-v5.18): schema_version=1 (GetSchemaVersion(1, 500) returned 1)
// - v5 current: schema_version=500
//
// IMPORTANT: Both v4 and v5 production state are at schema_version=1 but have
// incompatible formats:
// - v4: adaptive_routing as ARRAY, field named "default_pool_ids"
// - v5: adaptive_routing as OBJECT, field named "default_pools"
//
// Upgrade paths:
// 0: v5 state (version=0) -> no-op
// 1: AMBIGUOUS - v4 SDKv2 format or v5 production format, detected at runtime
func (r *LoadBalancerResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)

	return map[int64]resource.StateUpgrader{
		// Handle v5 state at version=0 -- no-op
		0: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV5,
		},

		// Handle schema_version=1 -- ambiguous between v4 SDKv2 and v5 production.
		//
		// PriorSchema is nil so the framework skips pre-decoding req.State entirely.
		// Both paths are handled via req.RawState.JSON in UpgradeFromV1Ambiguous:
		// - v4 format (default_pool_ids present, adaptive_routing is array): full transform
		// - v5 format (default_pools present, adaptive_routing is object): no-op re-decode
		1: {
			PriorSchema:   nil,
			StateUpgrader: v500.UpgradeFromV1Ambiguous,
		},
	}
}
