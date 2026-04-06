package zero_trust_dlp_custom_profile

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_dlp_custom_profile/migration/v500"
)

func init() {
	// Provide target schema to migration package (avoids circular import)
	v500.V5TargetSchema = func(ctx context.Context) schema.Schema {
		return ResourceSchema(ctx)
	}
}

var _ resource.ResourceWithUpgradeState = (*ZeroTrustDLPCustomProfileResource)(nil)
var _ resource.ResourceWithMoveState = (*ZeroTrustDLPCustomProfileResource)(nil)

// MoveState registers state movers for resource renames.
// This enables Terraform 1.8+ `moved` blocks to automatically trigger state migration
// from cloudflare_dlp_profile or cloudflare_zero_trust_dlp_profile to cloudflare_zero_trust_dlp_custom_profile.
func (r *ZeroTrustDLPCustomProfileResource) MoveState(ctx context.Context) []resource.StateMover {
	sourceSchema := v500.SourceCloudflareDLPProfileSchema()

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
//   - Transforms context_awareness from array to single nested
//
// 2. v5.16.0 state (schema_version=0) → v5 (version=500): No-op upgrade
//   - v5.16.0 was released with dormant state upgrader (no GetSchemaVersion)
//   - State already has context_awareness as object, no transformation needed
//
// 3. v5 state (version=1) → v5 (version=500): No-op upgrade
//
// IMPORTANT: Both v4 and v5.16.0 have schema_version=0, so the version 0
// upgrader must detect the format at runtime by inspecting the raw state.
func (r *ZeroTrustDLPCustomProfileResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)

	return map[int64]resource.StateUpgrader{
		// Handle BOTH v4 (schema_version=0) AND v5.16.0 (version=0) states
		// PriorSchema is nil because v4 and v5 have incompatible schemas:
		// - v4: context_awareness is array (ListNestedAttribute)
		// - v5: context_awareness is object (SingleNestedAttribute)
		// Handler uses RawState to detect format and process accordingly.
		0: {
			PriorSchema:   nil, // Use RawState - schemas are incompatible
			StateUpgrader: v500.UpgradeFromVersion0,
		},

		// Handle state from v5 Plugin Framework provider with version=1
		// No-op upgrade that just bumps the version to 500
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV1,
		},
	}
}
