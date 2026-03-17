// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_settings

import (
	"context"

	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_gateway_settings/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func init() {
	// Provide target schema to migration package (avoids circular import)
	v500.V5TargetSchema = func(ctx context.Context) schema.Schema {
		return ResourceSchema(ctx)
	}
}

var _ resource.ResourceWithUpgradeState = (*ZeroTrustGatewaySettingsResource)(nil)
var _ resource.ResourceWithMoveState = (*ZeroTrustGatewaySettingsResource)(nil)

// MoveState registers state movers for the resource rename:
// cloudflare_teams_account → cloudflare_zero_trust_gateway_settings.
//
// This enables Terraform 1.8+ `moved` blocks to automatically trigger state migration
// when renaming resources from the old type to the new type.
func (r *ZeroTrustGatewaySettingsResource) MoveState(ctx context.Context) []resource.StateMover {
	sourceSchema := v500.SourceV4ZeroTrustGatewaySettingsSchema()

	return []resource.StateMover{
		{
			SourceSchema: &sourceSchema,
			StateMover:   v500.MoveFromCloudflareTeamsAccount,
		},
	}
}

// UpgradeState registers state upgraders for schema version changes.
//
// This handles three upgrade paths:
// 1. v4 state (schema_version=0) → v5 (version=500): Full transformation
//   - Flat boolean fields → nested under settings.*
//   - TypeList MaxItems:1 blocks → SingleNestedAttribute pointers
//
// 2. v5.16.0 state (schema_version=0) → v5 (version=500): No-op upgrade
//   - v5.16.0 was released with dormant state upgrader (no GetSchemaVersion)
//   - State already has settings object, no transformation needed
//
// 3. v5 state (version=1) → v5 (version=500): No-op upgrade
//
// IMPORTANT: Both v4 and v5.16.0 have schema_version=0, so the version 0
// upgrader must detect the format at runtime by inspecting the raw state.
func (r *ZeroTrustGatewaySettingsResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	// v5 schema at version=1 for no-op pass-through upgrader
	v5SchemaVersion1 := ResourceSchema(ctx)
	v5SchemaVersion1.Version = 1

	return map[int64]resource.StateUpgrader{
		// Handle BOTH v4 (schema_version=0) AND v5.16.0 (version=0) states
		// PriorSchema is nil because v4 and v5 have incompatible schemas:
		// - v4: flat structure with ListNestedAttribute blocks
		// - v5: nested under settings with SingleNestedAttribute
		// Handler uses RawState to detect format and process accordingly.
		0: {
			PriorSchema:   nil, // Use RawState - schemas are incompatible
			StateUpgrader: v500.UpgradeFromVersion0,
		},

		// Handle state from v5 Plugin Framework provider (version=1)
		// No-op version bump to 500
		1: {
			PriorSchema:   &v5SchemaVersion1,
			StateUpgrader: v500.UpgradeFromV5,
		},
	}
}
