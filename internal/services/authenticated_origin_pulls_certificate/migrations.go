// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package authenticated_origin_pulls_certificate

import (
	"context"

	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/authenticated_origin_pulls_certificate/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*AuthenticatedOriginPullsCertificateResource)(nil)

func (r *AuthenticatedOriginPullsCertificateResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	v4Schema := v500.ResourceSchema(ctx)
	v5Schema := ResourceSchema(ctx)

	return map[int64]resource.StateUpgrader{
		// v4 → v5 migration (schema version 0 → 500)
		// Handles migration from v4 provider per-zone certificates
		0: {
			PriorSchema:   &v4Schema,
			StateUpgrader: v500.UpgradeFromV0,
		},
		// v5 → v5 no-op migration (schema version 1 → 500)
		// Handles resources already in v5 format that need schema version bump
		1: {
			PriorSchema:   &v5Schema,
			StateUpgrader: v500.UpgradeFromV1,
		},
	}
}
