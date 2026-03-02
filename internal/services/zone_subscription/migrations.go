// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_subscription

import (
	"context"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone_subscription/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*ZoneSubscriptionResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
//
// cloudflare_zone_subscription is a new v5 resource with no v4 counterpart.
// All state upgraders are no-ops — the schema structure has not changed.
//
// This handles two upgrade paths:
// 1. Early v5 state (version=0) → v5 (version=500): No-op (state already in v5 format)
// 2. v5 state (version=1) → v5 (version=500): No-op upgrade
//
// to safely bump existing v5 users from version 0 to 1.
func (r *ZoneSubscriptionResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)

	// For version 1 upgrader, use the current v5 schema but override version to 1.
	// This is necessary because schema version returns 500,
	// but we need PriorSchema to match the state version being upgraded (version 1).
	v5SchemaVersion1 := ResourceSchema(ctx)
	v5SchemaVersion1.Version = 1

	return map[int64]resource.StateUpgrader{
		// Handle state at version 0 (early v5 adopters before schema versioning was added).
		// No transformation needed — state is already in v5 format.
		0: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV0,
		},

		// Handle state at version 1 (no-op upgrade to version 500).
		1: {
			PriorSchema:   &v5SchemaVersion1,
			StateUpgrader: v500.UpgradeFromV5,
		},
	}
}
