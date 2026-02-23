// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dex_test

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_dex_test/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*ZeroTrustDEXTestResource)(nil)
var _ resource.ResourceWithMoveState = (*ZeroTrustDEXTestResource)(nil)

// MoveState registers state movers for resource renames.
// This enables Terraform 1.8+ `moved` blocks to automatically trigger state migration.
func (r *ZeroTrustDEXTestResource) MoveState(ctx context.Context) []resource.StateMover {
	sourceSchema := v500.SourceCloudflareDeviceDexTestSchema()

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
// 2. v5 state (version=1) → v5 (version=500): No-op upgrade (when TF_MIG_TEST=1)
//
// The separation of schema versions (v4=0, v5=1/500) eliminates the need for
// dual-format detection that was required in earlier implementations.
func (r *ZeroTrustDEXTestResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	sourceSchema := v500.SourceCloudflareDeviceDexTestSchema()
	targetSchema := ResourceSchema(ctx)

	return map[int64]resource.StateUpgrader{
		// Handle state from v4 SDKv2 provider (schema_version=0)
		// This performs full v4 → v5 transformation
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
