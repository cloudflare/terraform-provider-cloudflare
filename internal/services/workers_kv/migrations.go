// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_kv

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_kv/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*WorkersKVResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
//
// This handles two upgrade paths:
// 1. v4 state (schema_version=0) → v5 (version=1 or 500): Full transformation (key → key_name)
// 2. Early v5 state (schema_version=0) → v5 (version=1 or 500): No-op version bump
// 3. v5 state (version=1) → v5 (version=500): No-op upgrade
//
// The v0 upgrader uses the target schema as PriorSchema and detects whether the state
// is v4 or v5 format by checking if key_name is populated after parsing.
func (r *WorkersKVResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)

	return map[int64]resource.StateUpgrader{
		0: {
			// Use target schema as PriorSchema so v5 state (with key_name) can be parsed.
			// v4 state (with key) will fail to parse, which we detect and handle via RawState.
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV0,
		},
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV1,
		},
	}
}
