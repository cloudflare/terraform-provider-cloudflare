// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package regional_tiered_cache

import (
	"context"
	"os"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/regional_tiered_cache/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*RegionalTieredCacheResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
//
// This handles two upgrade paths:
// 1. v4 state (schema_version=0) → v5 (version=500): Full transformation
// 2. v5 state (version=1) → v5 (version=500): No-op upgrade (when TF_MIG_TEST=1)
//
// The separation of schema versions (v4=0, v5=1/500) eliminates the need for
// dual-format detection that was required in earlier implementations.
func (r *RegionalTieredCacheResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	sourceSchema := v500.SourceCloudflareRegionalTieredCacheSchema()
	targetSchema := ResourceSchema(ctx)

	// Safeguard: In production (TF_MIG_TEST not set), only handle v4→v5 migration
	// This prevents unnecessary state upgrades when schema version is 1 (dormant)
	if os.Getenv("TF_MIG_TEST") == "" {
		return map[int64]resource.StateUpgrader{
			0: {
				PriorSchema:   &sourceSchema,
				StateUpgrader: v500.UpgradeFromV4,
			},
		}
	}

	// Testing mode (TF_MIG_TEST=1): Handle both upgrade paths
	return map[int64]resource.StateUpgrader{
		// Handle state from v4 SDKv2 provider (schema_version=0)
		// Performs full transformation: v4 → v5 with new computed fields
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromV4,
		},

		// Handle state from v5 Plugin Framework provider with version=1
		// This is a no-op upgrade that just bumps the version to 500
		// Only triggered when TF_MIG_TEST=1 (GetSchemaVersion returns 500)
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV5,
		},
	}
}
