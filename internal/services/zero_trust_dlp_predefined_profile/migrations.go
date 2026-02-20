package zero_trust_dlp_predefined_profile

import (
	"context"
	"os"

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
// Both v4 SDKv2 and early v5 Plugin Framework used schema_version=0.
// This requires a conditional approach based on TF_MIG_TEST:
//
// Production (no TF_MIG_TEST): schema returns version 1
//   - Slot 0: no-op upgrader (safely bumps existing v5 users from 0→1)
//
// Testing (TF_MIG_TEST=1): schema returns version 500
//   - Slot 0: v4→v5 full transformation (v4 state has schema_version=0)
//   - Slot 1: v5 no-op (v5 users already bumped to version=1 in prod)
func (r *ZeroTrustDLPPredefinedProfileResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)

	if os.Getenv("TF_MIG_TEST") == "" {
		return map[int64]resource.StateUpgrader{
			0: {
				PriorSchema: &targetSchema,
				StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
					resp.State.Raw = req.State.Raw
				},
			},
		}
	}

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
