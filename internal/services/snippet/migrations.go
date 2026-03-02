package snippet

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/snippet/migration/v500"
)

var _ resource.ResourceWithUpgradeState = (*SnippetResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
//
// v4 snippet had schema_version=1. v5 uses 500.
//
// This handles three upgrade paths:
// 1. v5 state (version=0) → current: No-op (existing v5 users before migration was added)
// 2. v4 state (schema_version=1) → current: Full transformation
//   - name → snippet_name
//   - main_module → metadata.main_module
//   - files blocks → files list attribute
//
func (r *SnippetResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)

	sourceSchema := v500.SourceCloudflareSnippetSchema()

	return map[int64]resource.StateUpgrader{
		// Handle state from v5 Plugin Framework provider (version=0, before migration was added)
		// This is a no-op upgrade
		0: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV5,
		},

		// Handle state from v4 SDKv2 provider (schema_version=1)
		// Full transformation: name→snippet_name, main_module→metadata.main_module, files blocks→list
		1: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromV4,
		},
	}
}
