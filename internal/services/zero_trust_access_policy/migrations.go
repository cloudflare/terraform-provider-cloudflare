package zero_trust_access_policy

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_policy/migration/v500"
)

var _ resource.ResourceWithUpgradeState = (*ZeroTrustAccessPolicyResource)(nil)
var _ resource.ResourceWithMoveState = (*ZeroTrustAccessPolicyResource)(nil)

// MoveState handles moves from cloudflare_access_policy (v4) to cloudflare_zero_trust_access_policy (v5).
// This is triggered when users use the `moved` block (Terraform 1.8+):
//
//	moved {
//	    from = cloudflare_access_policy.example
//	    to   = cloudflare_zero_trust_access_policy.example
//	}
func (r *ZeroTrustAccessPolicyResource) MoveState(ctx context.Context) []resource.StateMover {
	v4Schema := v500.SourceAccessPolicySchema()
	return []resource.StateMover{
		{
			SourceSchema: &v4Schema,
			StateMover:   v500.MoveFromAccessPolicy,
		},
	}
}

// UpgradeState handles schema version upgrades.
// Both v4 cloudflare_access_policy and early v5 cloudflare_zero_trust_access_policy have schema_version=0.
//
// We use v5Schema at version 0, which means:
// - Early v5 state (v5.12-v5.15): works correctly, passes through as no-op
// - v4 state via `terraform state mv`: will FAIL (schema mismatch)
//
// For v4 → v5 migration, users MUST use `moved` blocks (Terraform 1.8+) which go through MoveState.
func (r *ZeroTrustAccessPolicyResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {

	v5Schema := ResourceSchema(ctx)

	return map[int64]resource.StateUpgrader{
		// Handle early v5 state with schema_version=0 (v5.12-v5.15) AND v4 state
		// PriorSchema is nil to allow raw JSON unmarshaling in handler
		// Handler detects v4 format (application_id/precedence/connection_rules=[])
		// and transforms using v4 source schema, or passes v5 state through unchanged
		0: {
			PriorSchema:   nil,
			StateUpgrader: v500.UpgradeFromSchemaV0,
		},
		// Handle upgrades from v5 with schema_version=1
		// This is a no-op since the schema is compatible.
		1: {
			PriorSchema:   &v5Schema,
			StateUpgrader: v500.UpgradeFromSchemaV1,
		},
	}
}
