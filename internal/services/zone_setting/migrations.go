// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_setting

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone_setting/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*ZoneSettingResource)(nil)
var _ resource.ResourceWithMoveState = (*ZoneSettingResource)(nil)

// MoveState registers state movers for resource renames.
// This enables Terraform 1.8+ `moved` blocks to automatically trigger state migration
// from the legacy cloudflare_zone_settings_override (v4) to cloudflare_zone_setting (v5).
//
// The v4 cloudflare_zone_settings_override is a one-to-many resource: one v4 resource
// holds many settings in a settings {} block. The v5 provider has one
// cloudflare_zone_setting per setting.
//
// Since MoveStateRequest does not include the target resource name, the handler uses
// a heuristic: it picks the first non-null, non-empty setting from settings[0].
// This works correctly for single-setting migrations (the common case).
// For multi-setting migrations, tf-migrate deletes the v4 state and creates v5
// resources fresh via terraform apply.
func (r *ZoneSettingResource) MoveState(ctx context.Context) []resource.StateMover {
	return []resource.StateMover{
		{
			// No SourceSchema — use SourceRawState.JSON for raw JSON parsing.
			// This avoids enumerating all 50+ v4 settings in a typed schema.
			StateMover: v500.MoveZoneSettingsOverrideToZoneSetting,
		},
	}
}

// UpgradeState registers state upgraders for schema version changes.
//
// This handles the following upgrade paths for cloudflare_zone_setting:
//
//  1. v4 state (schema_version=2): The e2e migration pipeline renames
//     cloudflare_zone_settings_override → cloudflare_zone_setting in state while
//     preserving schema_version=2 and all v4 attributes. This upgrader reads the
//     v4 attributes and produces a valid v5 state (picking the first non-null setting).
//     The resulting state entry is orphaned and Terraform removes it, creating v5
//     resources fresh via terraform apply.
//
//  2. v5 state (version=1) → v5 (version=500): No-op upgrade (version bump only).
//     This is triggered when TF_MIG_TEST=1 causes GetSchemaVersion to return 500.
//
//  3. v5 state (version=0): No-op pass-through for early v5 releases.
func (r *ZoneSettingResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)
	sourceSchema := v500.SourceZoneSettingSchema()

	return map[int64]resource.StateUpgrader{
		// Handle state from v4 cloudflare_zone_settings_override (schema_version=2).
		// The e2e migration pipeline renames the resource type in state but preserves
		// schema_version=2 and the full v4 attributes. This upgrader reads the v4
		// attributes and produces a valid v5 cloudflare_zone_setting state by picking
		// the first non-null setting. The resulting state entry will be orphaned (its
		// name won't match any v5 config resource) and Terraform will remove it,
		// creating the v5 resources fresh via terraform apply.
		2: {
			// No PriorSchema — use RawState.JSON for raw JSON parsing to avoid
			// enumerating all 50+ v4 settings in a typed schema.
			StateUpgrader: v500.UpgradeFromV4ZoneSettingsOverride,
		},

		// Handle state from v5 Plugin Framework provider with version=1.
		// This is a no-op upgrade that just bumps the version to 500.
		// Only triggered when TF_MIG_TEST=1.
		1: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromV5,
		},

		// Keep the existing no-op upgraders for any state written with version=0
		// by early v5 releases that used the raw pass-through pattern.
		0: {
			PriorSchema: &targetSchema,
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				resp.State.Raw = req.State.Raw
			},
		},
	}
}
