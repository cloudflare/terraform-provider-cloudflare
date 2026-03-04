package url_normalization_settings

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/url_normalization_settings/migration/v500"
)

var _ resource.ResourceWithUpgradeState = (*URLNormalizationSettingsResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
//
// This handles two upgrade paths:
// 1. v4 state (schema_version=0) → v5 (version=500): No-op (schema identical)
// 2. v5 state (version=1) → v5 (version=500): No-op upgrade
//
// to safely bump existing v5 users from version 0 to 1 without triggering the
// v4→v5 upgrade path (which uses the v4 source schema).
func (r *URLNormalizationSettingsResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)

	sourceSchema := v500.SourceURLNormalizationSettingsSchema()

	return map[int64]resource.StateUpgrader{
		// Handle state from v4 SDKv2 provider (schema_version=0)
		// Schema is identical between v4 and v5, so this is a no-op
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
