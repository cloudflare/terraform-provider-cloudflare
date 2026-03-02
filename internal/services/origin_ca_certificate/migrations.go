// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_ca_certificate

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/origin_ca_certificate/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*OriginCACertificateResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
//
// This handles two upgrade paths:
// 1. v4 state (schema_version=0) → v5 (version=500): Full transformation
//   - hostnames: Set → List
//   - requested_validity: Int64 → Float64 (default 5475)
//   - min_days_for_renewal: Dropped
//
// 2. v5 state (version=1) → v5 (version=500): No-op upgrade
func (r *OriginCACertificateResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)
	sourceSchema := v500.SourceCloudflareOriginCACertificateSchema()

	return map[int64]resource.StateUpgrader{
		// Handle state from v4 SDKv2 provider (schema_version=0)
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromV4,
		},

		// Handle state from v5 Plugin Framework provider with version=1
		// This is a no-op upgrade that just bumps the version to 500
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV5,
		},
	}
}
