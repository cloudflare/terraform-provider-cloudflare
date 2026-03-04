// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_rule

import (
	"context"
	"os"

	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/access_rule/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*AccessRuleResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
//
// This handles two upgrade paths:
// 1. v4 state (schema_version=1) -> v5 (version=500): Full transformation
//   - Unwraps configuration array[0] -> configuration object
//   - Initializes new computed fields
//
// 2. v5 state (version=1) -> v5 (version=500): No-op upgrade (when TF_MIG_TEST=1)
//   - Just bumps version number, no transformation
//
// The separation of schema versions (v4=1, v5=1/500) with GetSchemaVersion
// allows controlled rollout and eliminates dual-format detection issues.
func (r *AccessRuleResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
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
	sourceSchema := v500.SourceV4AccessRuleSchema(ctx)

	return map[int64]resource.StateUpgrader{
		// Handle fresh v5 resources (version 0 -> 500)
		// When a v5 resource is created with TF_MIG_TEST not set, it might start at version 0
		// This is a no-op upgrade - just bumps version, no transformation needed
		0: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV5,
		},

		// Handle state from v4 SDKv2 provider (schema_version=1)
		// Performs full transformation: configuration array -> object
		// v4 states have schema_version=1 after their internal v0->v1 migration
		1: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromV4,
		},
	}
}
