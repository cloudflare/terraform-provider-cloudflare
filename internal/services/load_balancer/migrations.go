// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer

import (
	"context"
	"os"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/load_balancer/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*LoadBalancerResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
//
// This handles two upgrade paths:
// 1. v4 state (schema_version=1) -> v5 (version=500): Full transformation
//    - Field renames, type conversions, structure transformations
// 2. v5 state (version=1) -> v5 (version=500): No-op upgrade (when TF_MIG_TEST=1)
//    - Just bumps the version, no transformation needed
//
// The separation of schema versions (v4=1, v5=1/500) with GetSchemaVersion allows
// controlled rollout: migrations are dormant in production (version=1) until enabled
// with TF_MIG_TEST=1 for testing (version=500).
func (r *LoadBalancerResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)

	if os.Getenv("TF_MIG_TEST") == "" {
		// Production mode: preserve existing upgraders only
		return map[int64]resource.StateUpgrader{
			0: {
				PriorSchema: &targetSchema,
				StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
					resp.State.Raw = req.State.Raw
				},
			},
		}
	}

	// Test mode (TF_MIG_TEST=1): full StateUpgrader migration
	sourceSchema := v500.SourceCloudflareLoadBalancerSchema()

	return map[int64]resource.StateUpgrader{
		// Handle state from v4 SDKv2 provider (schema_version=1)
		1: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromV4,
		},

		// Handle state from v5 Plugin Framework provider with version=0
		0: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV5,
		},
	}
}
