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
// This handles two upgrade paths:
// 1. v4 state (schema_version=1) → v5 (version=500): Full transformation
//    - Field renames, type conversions, structure transformations
// 2. v5 state (version=1) → v5 (version=500): No-op upgrade (when TF_MIG_TEST=1)
//    - Just bumps the version, no transformation needed
//
// The separation of schema versions (v4=1, v5=1/500) with GetSchemaVersion allows
// controlled rollout: migrations are dormant in production (version=1) until enabled
// with TF_MIG_TEST=1 for testing (version=500).
func (r *LoadBalancerResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	sourceSchema := v500.SourceCloudflareLoadBalancerSchema()
	targetSchema := ResourceSchema(ctx)

	return map[int64]resource.StateUpgrader{
		// Handle state from v4 SDKv2 provider (schema_version=1)
		// This was the version set in v4 for resources with state migrations
		1: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromV4,
		},

		// Handle state from v5 Plugin Framework provider with version=1
		// This is a no-op upgrade that just bumps the version to 500
		// Only triggered when TF_MIG_TEST=1 (GetSchemaVersion returns 500)
		// Note: We use version 1 (not 0) to distinguish v4 from v5 state
		0: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV5,
		},
	}
}
