// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_token

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/account_token/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*AccountTokenResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
//
// This handles two upgrade paths:
//
// 1. v0 state (schema_version=0) → v500: Full transformation
//   - Converts policies[].resources from map to JSON string
//   - Removes policies[].id (computed field)
//   - Removes policies[].permission_groups[].meta and .name (computed fields)
//
// 2. v1 state (schema_version=1) → v500: No-op upgrade (version bump only)
func (r *AccountTokenResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	v0Schema := v500.SourceSchemaV0()
	v1Schema := v500.SourceSchemaV1()

	return map[int64]resource.StateUpgrader{
		// Handle state from early v5 releases (v5.10, v5.11) with map-based resources
		0: {
			PriorSchema:   &v0Schema,
			StateUpgrader: v500.UpgradeFromV0,
		},

		// Handle state from dormant v5 (version=1) — no-op version bump
		1: {
			PriorSchema:   &v1Schema,
			StateUpgrader: v500.UpgradeFromV1,
		},
	}
}
