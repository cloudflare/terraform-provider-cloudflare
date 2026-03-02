package snippet_rules

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/snippet_rules/migration/v500"
)

var _ resource.ResourceWithUpgradeState = (*SnippetRulesResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
//
// v4 snippet_rules had schema_version=1. v5 uses GetSchemaVersion(2, 500).
//
// Upgrade paths:
// 1. v5 state (version=0) → current: No-op (existing v5 users before migration was added)
// 2. v4 state (schema_version=1) → current: Transform (add computed fields id, last_updated)
func (r *SnippetRulesResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)

	sourceSchema := v500.SourceSnippetRulesSchema()

	return map[int64]resource.StateUpgrader{
		// v5 Plugin Framework provider (version=0, before migration was added) — no-op
		0: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV5,
		},
		// v4 SDKv2 provider (schema_version=1) — add computed fields
		1: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromV4,
		},
	}
}
