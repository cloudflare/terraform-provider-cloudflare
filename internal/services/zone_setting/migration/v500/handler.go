package v500

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV4ZoneSettingsOverride handles state upgrades from the v4
// cloudflare_zone_settings_override resource (schema_version=2) to v5
// cloudflare_zone_setting (version=500).
//
// This is triggered when the e2e migration pipeline renames the resource type
// in state from cloudflare_zone_settings_override → cloudflare_zone_setting
// while preserving the original schema_version=2 and v4 attributes.
//
// The v4 state contains ALL zone settings in a single "settings" block.
// Since UpgradeState cannot know which specific setting this resource
// should represent (the resource name is not available), it uses a heuristic:
// pick the first non-null, non-deprecated setting from settings[0].
//
// After upgrade, the resource will be in state under its old name (e.g.,
// "minimal") which won't match any v5 config resource (e.g., "minimal_brotli").
// Terraform will remove it from state and create the v5 resources fresh.
// This is expected and handled by MigrationV2TestStepAllowCreate.
func UpgradeFromV4ZoneSettingsOverride(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zone_setting state from v4 cloudflare_zone_settings_override (schema_version=2)")

	if req.RawState == nil {
		resp.Diagnostics.AddError(
			"Missing raw state",
			"RawState was nil for schema version 2 migration",
		)
		return
	}

	rawJSON := req.RawState.JSON
	if rawJSON == nil {
		resp.Diagnostics.AddError(
			"Missing raw state JSON",
			"RawState.JSON was nil for schema version 2 migration",
		)
		return
	}

	// Parse the raw v4 state JSON — same struct used by MoveState
	var v4State v4ZoneSettingsOverrideRaw
	if err := json.Unmarshal(rawJSON, &v4State); err != nil {
		resp.Diagnostics.AddError(
			"Failed to parse v4 state",
			fmt.Sprintf("Could not parse cloudflare_zone_settings_override state JSON: %s", err),
		)
		return
	}

	// Resolve zone_id
	zoneID := v4State.ZoneID
	if zoneID == "" {
		zoneID = v4State.ID
	}
	if zoneID == "" {
		resp.Diagnostics.AddError(
			"Missing zone_id",
			"The v4 state is missing zone_id. Cannot perform state upgrade.",
		)
		return
	}

	// Extract the first non-null setting from settings[0]
	if len(v4State.Settings) == 0 {
		resp.Diagnostics.AddError(
			"Missing settings block",
			"The v4 state has no settings block. Cannot perform state upgrade.",
		)
		return
	}

	settingID, settingValue, err := extractFirstSetting(v4State.Settings[0])
	if err != nil {
		resp.Diagnostics.AddError(
			"No migratable setting found",
			fmt.Sprintf("Could not find a non-null setting in v4 state: %s", err),
		)
		return
	}

	tflog.Info(ctx, "Extracted setting from v4 state for upgrade",
		map[string]interface{}{
			"zone_id":    zoneID,
			"setting_id": settingID,
		})

	targetValue := buildTargetValue(settingID, settingValue)

	target := &SourceZoneSettingModel{
		ID:            types.StringValue(settingID),
		SettingID:     types.StringValue(settingID),
		ZoneID:        types.StringValue(zoneID),
		Value:         targetValue,
		Enabled:       types.BoolNull(),
		Editable:      types.BoolValue(true),
		ModifiedOn:    timetypes.NewRFC3339Null(),
		TimeRemaining: types.Float64Null(),
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, target)...)

	tflog.Info(ctx, "State upgrade from v4 cloudflare_zone_settings_override completed",
		map[string]interface{}{
			"zone_id":    zoneID,
			"setting_id": settingID,
		})
}

// UpgradeFromV5 handles state upgrades from v5 Plugin Framework provider (version=1) to v5 (version=500).
//
// This is a no-op upgrade since the schema is compatible — it just bumps the version number.
// This handler is only triggered when TF_MIG_TEST=1 (GetSchemaVersion returns 500).
func UpgradeFromV5(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zone_setting state from version=1 to version=500 (no-op)")

	// CRITICAL: For no-op upgrades, copy raw state directly.
	// This preserves all state data without any transformation.
	resp.State.Raw = req.State.Raw

	tflog.Info(ctx, "State version bump from 1 to 500 completed")
}
