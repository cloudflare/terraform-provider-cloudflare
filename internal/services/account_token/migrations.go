// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_token

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/account_token/migration/v500"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/account_token/migration/v501"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*AccountTokenResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
//
// This handles three upgrade paths:
//
// 1. v0 state (schema_version=0) → v501: Full transformation from early v5
//   - Converts policies[].resources from map to JSON string
//   - Removes policies[].id (computed field)
//   - Removes policies[].permission_groups[].meta and .name (computed fields)
//
// 2. v1 state (schema_version=1) → v501: Set→List migration with sorting
//   - v5.16-v5.18 stored state at schema_version=1 with Set-based policies/permission_groups
//   - Structurally identical to v500 state — reuse v500 schema for deserialization
//   - Must sort canonically for stable List ordering (same as v500→v501)
//   - Note: The framework runs only ONE upgrader (the slot matching the state version),
//     so we cannot rely on slot 500 running after slot 1. Slot 1 must do sorting itself.
//
// 3. v500 state (schema_version=500) → v501: Set→List migration
//   - Converts Set-typed policies and permission_groups to sorted Lists
func (r *AccountTokenResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	v0Schema := v500.SourceSchemaV0()
	v500Schema := v501.SourceSchemaV500()

	// v5.16-v5.18 stored state at schema_version=1 with the same Set-based schema
	// as v500. Reuse the v500 schema for deserialization.
	v1Schema := v501.SourceSchemaV500()
	v1Schema.Version = 1

	return map[int64]resource.StateUpgrader{
		// Handle state from early v5 releases (v5.10, v5.11) with map-based resources
		0: {
			PriorSchema:   &v0Schema,
			StateUpgrader: v500.UpgradeFromV0,
		},

		// Handle state from v5.16-v5.18 (schema_version=1, Set-based)
		// Same format as v500 — deserialize Sets, sort canonically, write as Lists
		1: {
			PriorSchema:   &v1Schema,
			StateUpgrader: v501.UpgradeFromV500,
		},

		// Handle state from v500 (Set-based) → v501 (List-based with FastSetType)
		500: {
			PriorSchema:   &v500Schema,
			StateUpgrader: v501.UpgradeFromV500,
		},
	}
}
