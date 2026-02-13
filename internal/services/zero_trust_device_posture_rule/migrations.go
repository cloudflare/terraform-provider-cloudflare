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
// Clear schema version separation:
// - v4 SDKv2 provider: schema_version=0, input/match as blocks
// - v5 Plugin Framework provider: version=1 (production) or version=500 (test)
func (r *ZeroTrustDevicePostureRuleResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	// v4 schema for version=0 upgrader
	v4Schema := v500.SourceDevicePostureRuleSchema()

	// v5 schema for version=1 upgrader (override version to match production state)
	v5SchemaVersion1 := ResourceSchema(ctx)
	v5SchemaVersion1.Version = 1

	return map[int64]resource.StateUpgrader{
		// Handle state from v4 SDKv2 provider (schema_version=0)
		// Uses v4 PriorSchema to parse, then transforms to v5
		0: {
			PriorSchema:   &v4Schema,
			StateUpgrader: v500.UpgradeFromV0,
		},

		// Handle state from v5 Plugin Framework provider (version=1)
		// Uses v5 PriorSchema, no-op version bump to 500
		1: {
			PriorSchema:   &v5SchemaVersion1,
			StateUpgrader: v500.UpgradeFromV1,
		},
	}
}
