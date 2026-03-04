// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_organization

import (
	"context"

	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_organization/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithMoveState = (*ZeroTrustOrganizationResource)(nil)
var _ resource.ResourceWithUpgradeState = (*ZeroTrustOrganizationResource)(nil)

// MoveState handles moves from v4 resource names to cloudflare_zero_trust_organization (v500).
//
// Handles BOTH v4 resource names (they share identical schemas):
//   - cloudflare_access_organization (deprecated v4 name)
//   - cloudflare_zero_trust_access_organization (current v4 name)
//
// This is triggered when users use the `moved` block (Terraform 1.8+):
//
//	moved {
//	    from = cloudflare_access_organization.example
//	    to   = cloudflare_zero_trust_organization.example
//	}
func (r *ZeroTrustOrganizationResource) MoveState(ctx context.Context) []resource.StateMover {
	sourceSchema := v500.SourceCloudflareAccessOrganizationSchema()
	return []resource.StateMover{
		{
			SourceSchema: &sourceSchema,
			StateMover:   v500.MoveState,
		},
	}
}

// UpgradeState handles schema version upgrades for cloudflare_zero_trust_organization.
//
// This handles two upgrade paths:
// 1. Published v5 (schema_version=0) → Current v5 (version=500): No-op upgrade
//   - Published v5 releases had no Version field, defaulted to 0
//   - State already uses v5 schema format (objects)
//   - Just needs version bump
//
// 2. v5 with explicit version=1 (if any exist) → v5 (version=500): No-op upgrade
//
// Note: v4→v5 migration is NOT handled by UpgradeState.
// It's handled by MoveState, which transforms the state during resource rename
// (cloudflare_access_organization → cloudflare_zero_trust_organization).
func (r *ZeroTrustOrganizationResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)

	return map[int64]resource.StateUpgrader{
		// Handle published v5 (schema_version=0) → current v5 (version=500)
		//
		// Published v5 releases (before Version field was added) defaulted to version 0.
		// These states already use v5 schema format (login_design as object, etc.)
		// and just need a version bump to 500.
		//
		// Note: v4→v5 migration is NOT handled here. It's handled by MoveState which
		// transforms the state during resource rename.
		0: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV0,
		},

		// Handle state from v5 Plugin Framework provider with version=1 (if any)
		// This is a no-op upgrade that just bumps the version to 500.
		// Note: This may not be used in practice since published v5 used version 0.
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV1,
		},
	}
}
