// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package authenticated_origin_pulls_hostname_certificate

import (
	"context"

	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/authenticated_origin_pulls_hostname_certificate/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Interface assertions
var _ resource.ResourceWithUpgradeState = (*AuthenticatedOriginPullsHostnameCertificateResource)(nil)
var _ resource.ResourceWithMoveState = (*AuthenticatedOriginPullsHostnameCertificateResource)(nil)

// MoveState registers state movers for resource renames.
// This enables Terraform 1.8+ `moved` blocks to automatically trigger state migration.
func (r *AuthenticatedOriginPullsHostnameCertificateResource) MoveState(ctx context.Context) []resource.StateMover {
	sourceSchema := v500.SourceCloudflareAuthenticatedOriginPullsCertificateSchema()

	return []resource.StateMover{
		{
			SourceSchema: &sourceSchema,
			StateMover:   v500.MoveState,
		},
	}
}

// UpgradeState registers state upgraders for schema version changes.
func (r *AuthenticatedOriginPullsHostnameCertificateResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	v4Schema := v500.SourceCloudflareAuthenticatedOriginPullsCertificateSchema()
	v5Schema := v500.ResourceSchema(ctx)

	return map[int64]resource.StateUpgrader{
		// Upgrade from v4 (schema version 0) to v5 (schema version 500)
		// Handles migration from v4 provider per-hostname certificates
		0: {
			PriorSchema:   &v4Schema,
			StateUpgrader: v500.UpgradeFromV0,
		},
		// Upgrade from v5 (schema version 1) to v5 (schema version 500)
		// Handles resources already in v5 format that need schema version bump
		1: {
			PriorSchema:   &v5Schema,
			StateUpgrader: v500.UpgradeFromV1,
		},
	}
}
