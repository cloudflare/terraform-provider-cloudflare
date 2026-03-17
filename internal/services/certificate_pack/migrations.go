// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package certificate_pack

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/certificate_pack/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*CertificatePackResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
//
// This handles two upgrade paths:
// 1. v4 state (schema_version=0) → v5 (version=500): Full transformation
//   - Drops wait_for_active_status field
//   - Transforms validation_records items (removes cname_target and cname_name)
//   - Converts types.Set to customfield.Set for hosts
//   - Handles null validation_records/validation_errors
//
// 2. v5 state (version=1) → v5 (version=500): No-op upgrade
//   - Just bumps the version number, no data transformation
func (r *CertificatePackResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)
	sourceSchema := v500.SourceCloudflareCertificatePackSchema()

	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromV4,
		},
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV5,
		},
	}
}
