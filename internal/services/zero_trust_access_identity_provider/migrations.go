package zero_trust_access_identity_provider

import (
	"context"

	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_identity_provider/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func init() {
	// Provide target schema to migration package (avoids circular import)
	v500.V5TargetSchema = func(ctx context.Context) schema.Schema {
		return ResourceSchema(ctx)
	}
}

var _ resource.ResourceWithMoveState = (*ZeroTrustAccessIdentityProviderResource)(nil)
var _ resource.ResourceWithUpgradeState = (*ZeroTrustAccessIdentityProviderResource)(nil)

// MoveState handles moves from cloudflare_access_identity_provider to cloudflare_zero_trust_access_identity_provider.
// This is triggered when users use the `moved` block (Terraform 1.8+):
//
//	moved {
//	    from = cloudflare_access_identity_provider.example
//	    to   = cloudflare_zero_trust_access_identity_provider.example
//	}
func (r *ZeroTrustAccessIdentityProviderResource) MoveState(ctx context.Context) []resource.StateMover {
	sourceSchema := v500.SourceAccessIdentityProviderSchema()
	return []resource.StateMover{
		{
			SourceSchema: &sourceSchema,
			StateMover:   v500.MoveState,
		},
	}
}

// UpgradeState registers state upgraders for schema version changes.
//
// This handles three upgrade paths:
// 1. v4 state (schema_version=0) → v5 (version=500): Full transformation
//   - config stored as ARRAY (ListNestedBlock) → OBJECT (SingleNestedAttribute)
//
// 2. v5.16.0 state (schema_version=0) → v5 (version=500): No-op upgrade
//   - v5.16.0 was released with dormant state upgrader (no GetSchemaVersion)
//   - State already has config as object, no transformation needed
//
// 3. v5 state (version=1) → v5 (version=500): No-op upgrade
//
// IMPORTANT: Both v4 and v5.16.0 have schema_version=0, so the version 0
// upgrader must detect the format at runtime by inspecting the raw state.
func (r *ZeroTrustAccessIdentityProviderResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	// v5 schema at version=1 for no-op pass-through upgrader
	v5SchemaVersion1 := ResourceSchema(ctx)
	v5SchemaVersion1.Version = 1

	return map[int64]resource.StateUpgrader{
		// Handle BOTH v4 (schema_version=0) AND v5.16.0 (version=0) states
		// PriorSchema is nil because v4 and v5 have incompatible schemas:
		// - v4: config as ListNestedBlock (array)
		// - v5: config as SingleNestedAttribute (object)
		// Handler uses RawState to detect format and process accordingly.
		0: {
			PriorSchema:   nil, // Use RawState - schemas are incompatible
			StateUpgrader: v500.UpgradeFromVersion0,
		},
		// Handle state from v5 Plugin Framework provider with version=1
		1: {
			PriorSchema:   &v5SchemaVersion1,
			StateUpgrader: v500.UpgradeFromV5,
		},
	}
}
