// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package authenticated_origin_pulls_settings

import (
	"context"
	"os"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/authenticated_origin_pulls_settings/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Interface assertions
var _ resource.ResourceWithUpgradeState = (*AuthenticatedOriginPullsSettingsResource)(nil)
var _ resource.ResourceWithMoveState = (*AuthenticatedOriginPullsSettingsResource)(nil)

// MoveState registers state movers for resource renames.
// This enables Terraform 1.8+ `moved` blocks to automatically trigger state migration.
func (r *AuthenticatedOriginPullsSettingsResource) MoveState(ctx context.Context) []resource.StateMover {
	sourceSchema := v500.SourceCloudflareAuthenticatedOriginPullsSchema()

	return []resource.StateMover{
		{
			SourceSchema: &sourceSchema,
			StateMover:   v500.MoveState,
		},
	}
}

// UpgradeState registers state upgraders for schema version changes.
func (r *AuthenticatedOriginPullsSettingsResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)

	if os.Getenv("TF_MIG_TEST") == "" {
		// Production mode: preserve existing upgraders only
		return map[int64]resource.StateUpgrader{
			0: {
				PriorSchema: &targetSchema,
				StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
					resp.State.Raw = req.State.Raw
				},
			},
		}
	}

	// Test mode (TF_MIG_TEST=1): full StateUpgrader migration
	sourceSchema := v500.SourceCloudflareAuthenticatedOriginPullsSchema()

	return map[int64]resource.StateUpgrader{
		// Upgrade from v4 (schema version 0) to v5 (schema version 500)
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromLegacyV0,
		},
		// Upgrade from v5 (schema version 1) to v5 (schema version 500)
		// This is a no-op upgrade to support TF_MIG_TEST rollout mechanism
		1: {
			StateUpgrader: v500.UpgradeFromV1,
		},
	}
}
