package v500

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// v4SettingNameMap maps v4 attribute names to v5 setting IDs.
// The v4 provider used "zero_rtt" but the Cloudflare API expects "0rtt".
var v4SettingNameMap = map[string]string{
	"zero_rtt": "0rtt",
}

// v4DeprecatedSettings lists settings that should be skipped during migration.
var v4DeprecatedSettings = map[string]bool{
	"universal_ssl": true,
}

// mapV4SettingName translates a v4 attribute name to a v5 setting ID.
func mapV4SettingName(v4Name string) string {
	if v5Name, ok := v4SettingNameMap[v4Name]; ok {
		return v5Name
	}
	return v4Name
}

// v4ZoneSettingsOverrideRaw is a minimal struct for parsing the v4
// cloudflare_zone_settings_override state JSON.
//
// The v4 state has SchemaVersion=2 and contains:
//   - id: zone ID
//   - zone_id: zone ID
//   - settings: list (max 1) of setting attributes
//   - initial_settings: list (max 1) of initial setting values (read-only)
type v4ZoneSettingsOverrideRaw struct {
	ID       string                   `json:"id"`
	ZoneID   string                   `json:"zone_id"`
	Settings []map[string]interface{} `json:"settings"`
}

// MoveZoneSettingsOverrideToZoneSetting handles moving state from the legacy
// cloudflare_zone_settings_override (v4) to cloudflare_zone_setting (v5).
//
// This is triggered by Terraform 1.8+ when it encounters a `moved` block:
//
//	moved {
//	  from = cloudflare_zone_settings_override.example
//	  to   = cloudflare_zone_setting.example_http3
//	}
//
// The v4 cloudflare_zone_settings_override is a one-to-many resource: one v4
// resource holds many settings in a settings {} block. The v5 provider has one
// cloudflare_zone_setting per setting.
//
// Since the MoveStateRequest does not include the target resource name, this
// handler uses a heuristic: it picks the first non-null, non-empty setting from
// settings[0]. This works correctly for single-setting migrations (the common
// case). For multi-setting migrations, tf-migrate deletes the v4 state and
// creates v5 resources fresh via terraform apply.
//
// Note: tf-migrate's TransformState for zone_setting returns "" (deletes the
// v4 state entry). This handler is only invoked when a user manually writes a
// `moved` block pointing from cloudflare_zone_settings_override to
// cloudflare_zone_setting.
func MoveZoneSettingsOverrideToZoneSetting(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
	// Only handle moves from cloudflare_zone_settings_override
	if req.SourceTypeName != "cloudflare_zone_settings_override" {
		return
	}

	tflog.Info(ctx, "Moving state from cloudflare_zone_settings_override to cloudflare_zone_setting",
		map[string]interface{}{
			"source_type":           req.SourceTypeName,
			"source_schema_version": req.SourceSchemaVersion,
		})

	// Use raw JSON parsing to avoid enumerating all 50+ v4 settings in a typed schema.
	// The SourceRawState.JSON contains the full v4 state as JSON bytes.
	if req.SourceRawState == nil {
		resp.Diagnostics.AddError(
			"Missing source state",
			"The source state for cloudflare_zone_settings_override is nil. Cannot perform state move.",
		)
		return
	}

	rawJSON := req.SourceRawState.JSON
	if rawJSON == nil {
		resp.Diagnostics.AddError(
			"Missing source state JSON",
			"The source state JSON for cloudflare_zone_settings_override is nil. Cannot perform state move.",
		)
		return
	}

	// Parse the raw v4 state JSON
	var v4State v4ZoneSettingsOverrideRaw
	if err := json.Unmarshal(rawJSON, &v4State); err != nil {
		resp.Diagnostics.AddError(
			"Failed to parse source state",
			fmt.Sprintf("Failed to parse cloudflare_zone_settings_override state JSON: %s", err),
		)
		return
	}

	// Validate zone_id
	zoneID := v4State.ZoneID
	if zoneID == "" {
		// Fall back to id field (v4 used zone_id as the resource ID)
		zoneID = v4State.ID
	}
	if zoneID == "" {
		resp.Diagnostics.AddError(
			"Missing zone_id",
			"The source state for cloudflare_zone_settings_override is missing zone_id. Cannot perform state move.",
		)
		return
	}

	// Extract settings[0] — the v4 state has at most one settings block
	if len(v4State.Settings) == 0 {
		resp.Diagnostics.AddError(
			"Missing settings block",
			"The source state for cloudflare_zone_settings_override has no settings block. Cannot perform state move.",
		)
		return
	}

	settings := v4State.Settings[0]

	// Find the first non-null, non-empty, non-deprecated setting.
	// This heuristic works for single-setting migrations (the common case).
	// For multi-setting migrations, tf-migrate handles state deletion and fresh creation.
	settingID, settingValue, err := extractFirstSetting(settings)
	if err != nil {
		resp.Diagnostics.AddError(
			"No migratable setting found",
			fmt.Sprintf("Could not find a non-null setting in cloudflare_zone_settings_override state: %s", err),
		)
		return
	}

	tflog.Info(ctx, "Extracted setting from v4 state",
		map[string]interface{}{
			"zone_id":    zoneID,
			"setting_id": settingID,
		})

	// Build the target v5 cloudflare_zone_setting state.
	// The value is stored as a dynamic attribute — wrap the raw value appropriately.
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

	resp.Diagnostics.Append(resp.TargetState.Set(ctx, target)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "State move from cloudflare_zone_settings_override to cloudflare_zone_setting completed",
		map[string]interface{}{
			"zone_id":    zoneID,
			"setting_id": settingID,
		})
}

// extractFirstSetting finds the first non-null, non-empty, non-deprecated setting
// from the v4 settings map. Returns the v5 setting ID and its raw value.
//
// The v4 settings map contains all zone settings, most of which will be null or
// empty strings (defaults). We look for the first one that has a meaningful value.
//
// Simple settings are strings (e.g., "on"/"off"). Nested settings (security_header,
// minify, etc.) are stored as []interface{} in the v4 state JSON.
func extractFirstSetting(settings map[string]interface{}) (settingID string, value interface{}, err error) {
	type candidate struct {
		name  string
		value interface{}
	}
	var candidates []candidate

	for k, v := range settings {
		// Skip deprecated settings
		if v4DeprecatedSettings[k] {
			continue
		}

		// Skip null values
		if v == nil {
			continue
		}

		// Skip empty strings
		if s, ok := v.(string); ok {
			if s == "" {
				continue
			}
			candidates = append(candidates, candidate{k, v})
			continue
		}

		// Skip empty slices (v4 nested blocks are stored as []interface{})
		if arr, ok := v.([]interface{}); ok {
			if len(arr) == 0 {
				continue
			}
			candidates = append(candidates, candidate{k, v})
			continue
		}

		// Skip empty maps
		if m, ok := v.(map[string]interface{}); ok {
			if len(m) == 0 {
				continue
			}
			candidates = append(candidates, candidate{k, v})
			continue
		}
	}

	if len(candidates) == 0 {
		return "", nil, fmt.Errorf("no non-null settings found in v4 state")
	}

	// Pick the first candidate (map iteration order is non-deterministic in Go,
	// but for single-setting migrations there will be exactly one candidate)
	best := candidates[0]
	v5Name := mapV4SettingName(best.name)
	return v5Name, best.value, nil
}

// buildTargetValue converts a raw v4 setting value to a v5 NormalizedDynamicValue.
//
// Simple string settings (e.g., "on"/"off") become string dynamic values.
// Nested settings (security_header, minify, etc.) stored as []interface{} in v4
// are converted to object dynamic values.
func buildTargetValue(settingID string, rawValue interface{}) customfield.NormalizedDynamicValue {
	switch v := rawValue.(type) {
	case string:
		// Simple string setting (e.g., http3 = "on")
		return customfield.RawNormalizedDynamicValueFrom(types.StringValue(v))

	case []interface{}:
		// Nested block stored as array in v4 state (e.g., security_header, minify)
		// The array has exactly one element (the block contents as a map)
		if len(v) == 0 {
			return customfield.RawNormalizedDynamicValueFrom(types.StringNull())
		}
		blockMap, ok := v[0].(map[string]interface{})
		if !ok {
			return customfield.RawNormalizedDynamicValueFrom(types.StringNull())
		}

		// Special handling for security_header: wrap in strict_transport_security
		if settingID == "security_header" {
			return buildSecurityHeaderValue(blockMap)
		}

		// For other nested blocks, convert to object
		return buildObjectValue(blockMap)

	case map[string]interface{}:
		// Direct map (shouldn't happen in v4 state, but handle gracefully)
		return buildObjectValue(v)

	default:
		// Unknown type — return null
		return customfield.RawNormalizedDynamicValueFrom(types.StringNull())
	}
}

// buildSecurityHeaderValue wraps the v4 security_header block contents in
// strict_transport_security, matching the v5 API structure.
//
// v4 state: settings[0].security_header = [{ enabled = true, max_age = 86400, ... }]
// v5 value: { strict_transport_security = { enabled = true, max_age = 86400, ... } }
func buildSecurityHeaderValue(blockMap map[string]interface{}) customfield.NormalizedDynamicValue {
	innerAttrs, innerTypes := buildAttrMap(blockMap)
	innerObj, diags := types.ObjectValue(innerTypes, innerAttrs)
	if diags.HasError() {
		return customfield.RawNormalizedDynamicValueFrom(types.StringNull())
	}

	outerObj, outerDiags := types.ObjectValue(
		map[string]attr.Type{
			"strict_transport_security": innerObj.Type(context.Background()),
		},
		map[string]attr.Value{
			"strict_transport_security": innerObj,
		},
	)
	if outerDiags.HasError() {
		return customfield.RawNormalizedDynamicValueFrom(types.StringNull())
	}

	return customfield.RawNormalizedDynamicValueFrom(outerObj)
}

// buildObjectValue converts a map of raw values to a NormalizedDynamicValue object.
func buildObjectValue(m map[string]interface{}) customfield.NormalizedDynamicValue {
	attrs, attrTypes := buildAttrMap(m)
	obj, diags := types.ObjectValue(attrTypes, attrs)
	if diags.HasError() {
		return customfield.RawNormalizedDynamicValueFrom(types.StringNull())
	}
	return customfield.RawNormalizedDynamicValueFrom(obj)
}

// buildAttrMap converts a raw map to framework attr.Value and attr.Type maps.
func buildAttrMap(m map[string]interface{}) (map[string]attr.Value, map[string]attr.Type) {
	attrValues := make(map[string]attr.Value, len(m))
	attrTypes := make(map[string]attr.Type, len(m))

	for k, v := range m {
		var val attr.Value
		switch typedVal := v.(type) {
		case string:
			val = types.StringValue(typedVal)
		case bool:
			val = types.BoolValue(typedVal)
		case float64:
			val = types.Float64Value(typedVal)
		case int64:
			val = types.Int64Value(typedVal)
		case nil:
			val = types.StringNull()
		default:
			// For complex types, convert to string representation
			val = types.StringValue(fmt.Sprintf("%v", typedVal))
		}
		attrValues[k] = val
		attrTypes[k] = val.Type(context.Background())
	}

	return attrValues, attrTypes
}
