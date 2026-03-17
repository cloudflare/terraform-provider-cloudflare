// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_ssl

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/custom_ssl/migration/v500"
)

var _ resource.ResourceWithUpgradeState = (*CustomSSLResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
//
// Version history:
//   - 0: Existing v5 Plugin Framework state (before explicit versioning was applied).
//     Schema attributes are already flat (v5 format). No-op.
//   - 1: v4 SDKv2 provider state (SchemaVersion=1 was the final v4 schema version,
//     after the v4-internal v0→v1 upgrade that converted custom_ssl_options from TypeMap to TypeList).
//     Full transformation needed: unpack custom_ssl_options block, restructure geo_restrictions, etc.
//   - 2: Dormant production v5 state (GetSchemaVersion returns 2 in production).
//     No-op upgrade: schema is compatible, just bumps version to 500.
//   - 500: Active migration version.
//
// Why GetSchemaVersion(2, 500) instead of (1, 500)?
// The v4 resource already used SchemaVersion=1. Using preMigration=2 ensures that
// dormant v5 state (version=2) never collides with v4 state (version=1).
func (r *CustomSSLResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	sourceSchema := v500.SourceCloudflareCustomSSLSchema()
	targetSchema := ResourceSchema(ctx)

	// v5 dormant schema at version=2 (for the no-op upgrader key=2).
	v5SchemaVersion2 := ResourceSchema(ctx)
	v5SchemaVersion2.Version = 2

	return map[int64]resource.StateUpgrader{
		// Handle existing v5 state from before explicit versioning (version=0).
		// No-op: schema is already in v5 format.
		0: {
			PriorSchema: &targetSchema,
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				resp.State.Raw = req.State.Raw
			},
		},

		// Handle state from v4 SDKv2 provider (schema_version=1).
		// Performs full transformation: unpack custom_ssl_options block, convert types, etc.
		1: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromV4,
		},

		// Handle dormant v5 state (version=2, produced by GetSchemaVersion(2, 500) in production).
		// No-op: schema is compatible, just bumps version to 500.
		2: {
			PriorSchema:   &v5SchemaVersion2,
			StateUpgrader: v500.UpgradeFromV5,
		},
	}
}
