package snippet

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/snippet/migration/v500"
)

var _ resource.ResourceWithUpgradeState = (*SnippetResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
//
// Schema version history:
// - v4 SDKv2: schema_version=1
// - v5 production (v5.0-v5.18): schema_version=2 (GetSchemaVersion(2, 500) returned 2)
// - v5 current: schema_version=500
//
// Upgrade paths:
// 0: v5 state before migration was added -> no-op
// 1: v4 SDKv2 state -> full transformation (name->snippet_name, files blocks->list)
// 2: v5 production state (v5.0-v5.18) -> no-op
func (r *SnippetResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)
	sourceSchema := v500.SourceCloudflareSnippetSchema()

	return map[int64]resource.StateUpgrader{
		// v5 Plugin Framework provider (version=0, before migration was added) -- no-op
		0: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV5,
		},
		// v4 SDKv2 provider (schema_version=1) -- full transformation:
		// name->snippet_name, main_module->metadata.main_module, files blocks->list
		1: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromV4,
		},
		// v5 production state (schema_version=2, from GetSchemaVersion(2, 500) in v5.0-v5.18).
		// State is already in v5 format -- no transformation needed.
		2: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV5,
		},
	}
}
