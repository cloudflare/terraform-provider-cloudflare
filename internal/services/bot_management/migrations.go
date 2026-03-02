package bot_management

import (
	"context"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/bot_management/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*BotManagementResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
//
// This handles two upgrade paths:
// 1. v4 state (schema_version=0) -> v5 (version=500): Full transformation
//   - Copies all v4 fields directly (same names and types)
//   - Sets new v5-only fields to null (bm_cookie_enabled, cf_robots_variant, etc.)
//   - Sets stale_zone_configuration to null
//
// 2. v5 state (version=1) -> v5 (version=500): No-op upgrade
//   - Just bumps the version number, no data transformation
func (r *BotManagementResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)

	sourceSchema := v500.SourceCloudflareBotManagementSchema()

	return map[int64]resource.StateUpgrader{
		// Handle state from v4 SDKv2 provider (schema_version=0)
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromV4,
		},
		// Handle state from v5 Plugin Framework provider with version=1
		// This is a no-op upgrade that just bumps the version to 500
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV5,
		},
	}
}
