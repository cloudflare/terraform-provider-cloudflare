// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package turnstile_widget

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/turnstile_widget/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*TurnstileWidgetResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
//
// This handles two upgrade paths:
//
// 1. v4 state (schema_version=0) → v5 (version=500): Full transformation
//   - Transforms domains from Set to List with alphabetical sorting
//   - Copies sitekey from ID field (in v4, sitekey was the ID)
//   - Passes through all other fields unchanged
//   - Sets new v5 computed fields to null (will be refreshed from API)
//
// 2. v5 state (version=1) → v5 (version=500): No-op upgrade
func (r *TurnstileWidgetResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)
	sourceSchema := v500.SourceCloudfareTurnstileWidgetSchema()

	return map[int64]resource.StateUpgrader{
		// Handle state from v4 Plugin Framework provider (schema_version=0)
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromV4,
		},

		// Handle state from v5 Plugin Framework provider with version=1
		// No-op upgrade that just bumps the version to 500
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV5,
		},
	}
}
