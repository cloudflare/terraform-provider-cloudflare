// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpush_ownership_challenge

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/logpush_ownership_challenge/migration/v500"
)

var _ resource.ResourceWithUpgradeState = (*LogpushOwnershipChallengeResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
//
// This handles two upgrade paths:
// 1. v4 state (schema_version=0) → v5 (version=500): Full transformation
//   - Removes ownership_challenge_filename (v4 computed field)
//   - Sets filename, message, valid to Null (new v5 computed fields)
//
// 2. v5 state (version=1) → v5 (version=500): No-op upgrade
//
// The separation of schema versions (v4=0, v5=1/500) eliminates the need for
// dual-format detection that was required in earlier implementations.
func (r *LogpushOwnershipChallengeResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)
	sourceSchema := v500.SourceCloudflareLogpushOwnershipChallengeSchema()

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
