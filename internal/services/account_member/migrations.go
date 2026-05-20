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
// This handles three upgrade paths via two slots:
//
//   - Slot 0: AMBIGUOUS — v4 SDKv2 state OR v5.0-v5.15 state (both schema_version=0)
//     v4: has email_address, role_ids (no policies, no user)
//     v5.0-v5.15: has email, roles (List), policies (List with 'id' field), user
//     PriorSchema is nil; runtime detection via raw JSON disambiguates the two formats.
//
//   - Slot 1: v5.16-v5.18 stepping stone state (schema_version=1) → v500
//     Users who went through the v5.16+ stepping stone have v5 field names with
//     List/Set types and no policies 'id' field.
//     This performs List→Set conversion.
func (r *AccountMemberResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	v513Schema := v500.SourceV513Schema()

	return map[int64]resource.StateUpgrader{
		// Handle state at schema_version=0, which is AMBIGUOUS:
		// - v4 SDKv2 provider: email_address, role_ids
		// - v5.0-v5.15 production: email, roles, policies, user
		// PriorSchema=nil so we can inspect req.RawState.JSON to detect format.
		0: {
			StateUpgrader: v500.UpgradeFromV0Ambiguous,
		},
		// Handle state from v5.16-v5.18 stepping stone (schema_version=1)
		// State has: email, roles, policies, user (Set types, no policies 'id' field)
		1: {
			PriorSchema:   &v513Schema,
			StateUpgrader: v500.UpgradeFromV1,
		},
	}
}
