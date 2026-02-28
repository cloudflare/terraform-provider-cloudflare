// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package healthcheck

import (
	"context"
	"os"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/healthcheck/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*HealthcheckResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
//
// This handles two upgrade paths:
// 1. v4 state (schema_version=0) -> v5 (version=500): Full transformation
//    - Flat structure -> Nested http_config/tcp_config based on type
//    - Header Set -> Map transformation
// 2. v5 state (version=1) -> v5 (version=500): No-op upgrade (when TF_MIG_TEST=1)
//    - Just version bump, no transformation
//
// The separation of schema versions (v4=0, v5=1/500) eliminates the need for
// dual-format detection that was required in earlier implementations.
func (r *HealthcheckResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
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
	sourceSchema := v500.SourceHealthcheckSchema()

	return map[int64]resource.StateUpgrader{
		// Handle state from v4 SDKv2 provider (schema_version=0)
		// Full transformation: flat fields -> nested http_config/tcp_config
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromV4,
		},

		// Handle state from v5 Plugin Framework provider with version=1
		// This is a no-op upgrade that just bumps the version to 500
		// Only triggered when TF_MIG_TEST=1 (GetSchemaVersion returns 500)
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV5,
		},
	}
}
