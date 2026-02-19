package tiered_cache

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/tiered_cache/migration/v500"
)

var _ resource.ResourceWithUpgradeState = (*TieredCacheResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
//
// This handles two upgrade paths:
// 1. v4 state (schema_version=0) → v5 (version=500): Full transformation (cache_type → value)
// 2. v5 state (version=1) → v5 (version=500): No-op upgrade (when TF_MIG_TEST=1)
//
// In production (no TF_MIG_TEST), only a no-op upgrader is registered at slot 0
// to safely bump existing v5 users from version 0 to 1 without triggering the
// v4→v5 transformation (which would fail on v5-format state).
func (r *TieredCacheResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
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

	sourceSchema := v500.SourceTieredCacheSchema()

	return map[int64]resource.StateUpgrader{
		// Handle state from v4 SDKv2 provider (schema_version=0)
		// Transforms cache_type to value: "smart"→"on", "off"→"off"
		// "generic" returns error (must use tf-migrate tool to split)
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromV0,
		},

		// Handle state from v5 Plugin Framework provider with version=1
		// This is a no-op upgrade that just bumps the version to 500
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV1,
		},
	}
}
