package zero_trust_device_posture_rule

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_device_posture_rule/migration/v500"
)

var _ resource.ResourceWithUpgradeState = (*ZeroTrustDevicePostureRuleResource)(nil)
var _ resource.ResourceWithMoveState = (*ZeroTrustDevicePostureRuleResource)(nil)

// MoveState registers state movers for resource renames.
// This enables Terraform 1.8+ `moved` blocks to automatically trigger state migration
// from cloudflare_device_posture_rule to cloudflare_zero_trust_device_posture_rule.
func (r *ZeroTrustDevicePostureRuleResource) MoveState(ctx context.Context) []resource.StateMover {
	sourceSchema := v500.SourceDevicePostureRuleSchema()

	return []resource.StateMover{
		{
			SourceSchema: &sourceSchema,
			StateMover:   v500.MoveState,
		},
	}
}

// UpgradeState registers state upgraders for schema version changes.
//
// Both v4 and published v5 used schema_version=0 (neither set an explicit Version).
//
//   - Slot 0: no-op upgrader (safely bumps existing v5 users from 0→1)
//
// Testing: schema returns version 500
//   - Slot 0: v4→v5 full transformation (v4 state has schema_version=0)
//   - Slot 1: v5 no-op (v5 users already bumped to version=1 in prod)
//
// Note: v4 SDKv2 provider used resource type cloudflare_device_posture_rule,
// which is handled by MoveState, not UpgradeState.
func (r *ZeroTrustDevicePostureRuleResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)

	sourceSchema := v500.SourceDevicePostureRuleSchema()

	return map[int64]resource.StateUpgrader{
		// Handle state from v4 SDKv2 provider (schema_version=0)
		// Full transformation: list→single nested, blocks→attributes, etc.
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromV0,
		},

		// Handle state from v5 provider with version=1 (production dormant state)
		// No-op upgrade that just bumps the version to 500
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV1,
		},
	}
}
