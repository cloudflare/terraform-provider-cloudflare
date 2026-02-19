// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package certificate_pack

import (
	"context"
	"os"

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
// 2. v5 state (version=1) → v5 (version=500): No-op upgrade (when TF_MIG_TEST=1)
//   - Just bumps the version number, no data transformation
//
// In production (no TF_MIG_TEST), only a no-op upgrader is registered at slot 0
// to safely bump existing v5 users from version 0 to 1 without triggering the
// v4→v5 transformation (which would fail on v5-format state).
func (r *CertificatePackResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
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
