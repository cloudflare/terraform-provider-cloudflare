package zero_trust_dlp_predefined_profile

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_dlp_predefined_profile/migration/v500"
)

var _ resource.ResourceWithUpgradeState = (*ZeroTrustDLPPredefinedProfileResource)(nil)
var _ resource.ResourceWithMoveState = (*ZeroTrustDLPPredefinedProfileResource)(nil)

// MoveState registers state movers for resource renames.
// This enables Terraform 1.8+ `moved` blocks to automatically trigger state migration
// from cloudflare_dlp_profile or cloudflare_zero_trust_dlp_profile to cloudflare_zero_trust_dlp_predefined_profile.
func (r *ZeroTrustDLPPredefinedProfileResource) MoveState(ctx context.Context) []resource.StateMover {
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
// This handles two upgrade paths:
// 1. v4 state (schema_version=0) → v5 (version=500): Full transformation
// 2. v5 state (version=1) → v5 (version=500): No-op upgrade
func (r *ZeroTrustDLPPredefinedProfileResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)
	sourceSchema := v500.SourceCloudflareDLPProfileSchema()

	return map[int64]resource.StateUpgrader{
		// Handle state from v4 SDKv2 provider (schema_version=0)
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromV0,
		},

		// Handle state from v5 Plugin Framework provider with version=1
		// No-op upgrade that just bumps the version to 500
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV1,
		},
	}
}
