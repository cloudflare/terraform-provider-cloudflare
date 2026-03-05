// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_member

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/account_member/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*AccountMemberResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
//
// This handles two upgrade paths:
//
//   - Slot 0: v4 SDKv2 state (schema_version=0) → v500
//     Users upgrading from v4 provider have state with email_address, role_ids fields.
//     This performs a full transformation to v5 format.
//
//   - Slot 1: v5 stepping stone state (schema_version=1) → v500
//     Users who went through v5.17/v5.18 stepping stone already have v5 field names.
//     This is a no-op version bump.
//
// Note: Users on v5.0-v5.16 must upgrade to v5.17/v5.18 first (stepping stone)
// before upgrading to v5.19+. See migration guide for details.
func (r *AccountMemberResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	v4Schema := v500.SourceV4Schema()
	targetSchema := ResourceSchema(ctx)

	return map[int64]resource.StateUpgrader{
		// Handle state from v4 SDKv2 provider (schema_version=0)
		// State has: email_address, role_ids (no policies, no user)
		0: {
			PriorSchema:   &v4Schema,
			StateUpgrader: v500.UpgradeFromV4,
		},
		// Handle state from v5 stepping stone (schema_version=1)
		// State already has: email, roles, policies, user
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV1,
		},
	}
}
