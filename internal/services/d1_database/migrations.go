package d1_database

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/d1_database/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*D1DatabaseResource)(nil)

// UpgradeState handles schema version upgrades for cloudflare_d1_database.
//
// Schema version history:
// - v4 (framework): schema_version=0
// - v5 current: schema_version=500
//
// Upgrade paths:
// 1. v4 framework (schema_version=0) -> v5 (500): Full transformation
//    - "id" is copied to "uuid" (v5 uses uuid for API calls)
//    - New computed/optional fields initialized as null
func (r *D1DatabaseResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	sourceSchema := v500.SourceCloudflareD1DatabaseSchema()

	return map[int64]resource.StateUpgrader{
		// Handle state from the legacy framework cloudflare_d1_database (schema_version=0).
		// Transforms: id -> uuid (copy).
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromLegacyV0,
		},
	}
}
