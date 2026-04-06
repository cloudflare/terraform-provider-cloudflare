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
//   - Slot 1: v5.13 stepping stone state (schema_version=1) → v500
//     Users who went through v5.13+ stepping stone have v5 field names but different
//     schema structure (List vs Set, policies had 'id' field).
//     This performs List→Set conversion and removes the policy 'id' field.
//
// Note: Users on v5.0-v5.12 must upgrade to v5.13+ first (stepping stone)
// before upgrading to v5.19+. See migration guide for details.
func (r *AccountMemberResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	v4Schema := v500.SourceV4Schema()
	v513Schema := v500.SourceV513Schema()

	return map[int64]resource.StateUpgrader{
		// Handle state from v4 SDKv2 provider (schema_version=0)
		// State has: email_address, role_ids (no policies, no user)
		0: {
			PriorSchema:   &v4Schema,
			StateUpgrader: v500.UpgradeFromV4,
		},
		// Handle state from v5.13 stepping stone (schema_version=1)
		// State has: email, roles (List), policies (List with 'id' field), user
		// This transforms: List→Set, removes policy 'id' field
		1: {
			PriorSchema:   &v513Schema,
			StateUpgrader: v500.UpgradeFromV1,
		},
	}
}
