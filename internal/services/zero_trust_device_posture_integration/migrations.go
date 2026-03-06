// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_posture_integration

import (
	"context"
	"os"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_device_posture_integration/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Interface assertions
var _ resource.ResourceWithUpgradeState = (*ZeroTrustDevicePostureIntegrationResource)(nil)
var _ resource.ResourceWithMoveState = (*ZeroTrustDevicePostureIntegrationResource)(nil)

// MoveState registers state movers for resource renames.
// This enables Terraform 1.8+ `moved` blocks to automatically trigger state migration.
func (r *ZeroTrustDevicePostureIntegrationResource) MoveState(ctx context.Context) []resource.StateMover {
	sourceSchema := v500.SourceCloudflareDevicePostureIntegrationSchema()

	return []resource.StateMover{
		{
			SourceSchema: &sourceSchema,
			StateMover:   v500.MoveState,
		},
	}
}

// UpgradeState registers state upgraders for schema version changes.
//
// This handles two upgrade paths:
// 1. v4 state (schema_version=0) → v5 (version=500): Full transformation
//    - Removes deprecated 'identifier' field
//    - Provides default for 'interval' if missing
//    - Transforms 'config' from array to pointer object
//    - Provides empty config if missing
//
// 2. v5 state (version=1) → v5 (version=500): No-op upgrade (when TF_MIG_TEST=1)
//    - Just bumps version number, no state changes
//
// The separation of schema versions (v4=0, v5=1/500) eliminates the need for
// dual-format detection that was required in earlier implementations.
func (r *ZeroTrustDevicePostureIntegrationResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
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

	sourceSchema := v500.SourceCloudflareDevicePostureIntegrationSchema()

	return map[int64]resource.StateUpgrader{
		// Handle state from v4 SDKv2 provider (schema_version=0)
		// Full transformation: v4 → v5
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
