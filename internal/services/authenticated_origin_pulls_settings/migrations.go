// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package authenticated_origin_pulls_settings

import (
	"context"
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

	sourceSchema := v500.SourceCloudflareAuthenticatedOriginPullsSchema()

	return map[int64]resource.StateUpgrader{
		// Upgrade from v4 (schema version 0) to v5 (schema version 500)
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromLegacyV0,
		},
		// Upgrade from v5 (schema version 1) to v5 (schema version 500)
		1: {
			StateUpgrader: v500.UpgradeFromV1,
		},
	}
}
