// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone/migration/v500"
)

var _ resource.ResourceWithUpgradeState = (*ZoneResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
//
// Clear schema version separation:
//   - v4 SDKv2 provider: schema_version=0, flat structure
//   - v5 Plugin Framework provider: version=1 (production) or version=500 (test)
func (r *ZoneResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {

	// v5 schema for version=1 upgrader (override version to match production state)
	v5SchemaVersion1 := ResourceSchema(ctx)
	v5SchemaVersion1.Version = 1

	return map[int64]resource.StateUpgrader{
		// Handle state from v4 SDKv2 provider (schema_version=0)
		// PriorSchema: nil — v4 SDKv2 encoding is incompatible with Plugin Framework
		// schema types. The upgrader reads raw JSON directly via the source schema.
		0: {
			PriorSchema:   nil,
			StateUpgrader: v500.UpgradeFromV4,
		},

		// Handle state from v5 Plugin Framework provider (version=1)
		// Uses v5 PriorSchema, no-op version bump to 500
		1: {
			PriorSchema:   &v5SchemaVersion1,
			StateUpgrader: v500.UpgradeFromV5,
		},
	}
}
