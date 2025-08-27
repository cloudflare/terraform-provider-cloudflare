// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tiered_cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.ResourceWithUpgradeState = (*TieredCacheResource)(nil)

// tieredCacheResourceSchemaV0 defines the v0 schema (v4 provider format)
// The v4 schema used cache_type instead of value with values "smart", "generic", "off"
var tieredCacheResourceSchemaV0 = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:      true,
			PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
		},
		"zone_id": schema.StringAttribute{
			Required:      true,
			PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
		},
		// v4 used cache_type with values "smart", "generic", "off"
		"cache_type": schema.StringAttribute{
			Optional: true,
			Validators: []validator.String{
				// We need to accept any string here to handle existing state
				// The actual validation is done in the transformation
			},
		},
		// v5 uses value with "on", "off" - include for compatibility
		"value": schema.StringAttribute{
			Optional: true,
		},
		"editable": schema.BoolAttribute{
			Computed: true,
		},
		"modified_on": schema.StringAttribute{
			Computed:   true,
			CustomType: timetypes.RFC3339Type{},
		},
	},
}

func (r *TieredCacheResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema:   &tieredCacheResourceSchemaV0,
			StateUpgrader: upgradeTieredCacheStateV0toV1,
		},
	}
}

// upgradeTieredCacheStateV0toV1 migrates from v4 provider state format to v5
// Primary change: cache_type → value with value transformation
func upgradeTieredCacheStateV0toV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Debug(ctx, "Starting tiered cache state upgrade from v0 to v1")

	// Handle both typed state and raw JSON for maximum compatibility
	var rawState map[string]interface{}
	if req.RawState != nil && len(req.RawState.JSON) > 0 {
		if err := json.Unmarshal(req.RawState.JSON, &rawState); err != nil {
			resp.Diagnostics.AddError(
				"Failed to parse raw state during migration",
				fmt.Sprintf("Could not unmarshal raw state: %s", err),
			)
			return
		}
		tflog.Debug(ctx, "Parsed raw state successfully", map[string]interface{}{
			"keys": getMapKeys(rawState),
		})
	}

	// Try to get the prior state using the typed schema first
	var priorStateData struct {
		ID         types.String      `tfsdk:"id"`
		ZoneID     types.String      `tfsdk:"zone_id"`
		CacheType  types.String      `tfsdk:"cache_type"`
		Value      types.String      `tfsdk:"value"`
		Editable   types.Bool        `tfsdk:"editable"`
		ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on"`
	}

	resp.Diagnostics.Append(req.State.Get(ctx, &priorStateData)...)
	if resp.Diagnostics.HasError() {
		tflog.Warn(ctx, "Failed to decode typed state, using raw state fallback")
		// Clear the error and continue with raw state
		resp.Diagnostics = resp.Diagnostics.Errors()[:0]
	}

	// Initialize the new state model
	newState := TieredCacheModel{}

	// Migrate basic attributes that don't change
	if !priorStateData.ID.IsNull() {
		newState.ID = priorStateData.ID
		tflog.Debug(ctx, "Migrated ID from typed state", map[string]interface{}{
			"id": priorStateData.ID.ValueString(),
		})
	} else if id, exists := rawState["id"]; exists {
		if idStr, ok := id.(string); ok {
			newState.ID = types.StringValue(idStr)
			tflog.Debug(ctx, "Migrated ID from raw state", map[string]interface{}{
				"id": idStr,
			})
		}
	}

	if !priorStateData.ZoneID.IsNull() {
		newState.ZoneID = priorStateData.ZoneID
	} else if zoneID, exists := rawState["zone_id"]; exists {
		if zoneIDStr, ok := zoneID.(string); ok {
			newState.ZoneID = types.StringValue(zoneIDStr)
		}
	}

	if !priorStateData.Editable.IsNull() {
		newState.Editable = priorStateData.Editable
	} else if editable, exists := rawState["editable"]; exists {
		if editableBool, ok := editable.(bool); ok {
			newState.Editable = types.BoolValue(editableBool)
		}
	}

	if !priorStateData.ModifiedOn.IsNull() {
		newState.ModifiedOn = priorStateData.ModifiedOn
	} else if modifiedOn, exists := rawState["modified_on"]; exists {
		if modifiedOnStr, ok := modifiedOn.(string); ok {
			// Attempt to parse the RFC3339 time
			modifiedOnValue, diag := timetypes.NewRFC3339Value(modifiedOnStr)
			if diag.HasError() {
				tflog.Warn(ctx, "Failed to parse modified_on time, leaving as null", map[string]interface{}{
					"modified_on": modifiedOnStr,
					"error":       diag.Errors()[0].Summary(),
				})
			} else {
				newState.ModifiedOn = modifiedOnValue
			}
		}
	}

	// Handle the main transformation: cache_type → value
	// Transformation rules:
	// - cache_type = "smart" → value = "on"
	// - cache_type = "generic" → value = "on"
	// - cache_type = "off" → value = "off"
	// - value = "on" → value = "on" (already migrated)
	// - value = "off" → value = "off" (already migrated)

	var finalValue string
	var valueSource string

	// Check if we already have the new 'value' field (partially migrated state)
	if !priorStateData.Value.IsNull() && priorStateData.Value.ValueString() != "" {
		finalValue = priorStateData.Value.ValueString()
		valueSource = "typed value field"
		tflog.Debug(ctx, "Found existing value field", map[string]interface{}{
			"value": finalValue,
		})
	} else if value, exists := rawState["value"]; exists {
		if valueStr, ok := value.(string); ok && valueStr != "" {
			finalValue = valueStr
			valueSource = "raw value field"
			tflog.Debug(ctx, "Found existing value in raw state", map[string]interface{}{
				"value": finalValue,
			})
		}
	}

	// If no value found, check for cache_type (v4 format)
	if finalValue == "" {
		if !priorStateData.CacheType.IsNull() {
			cacheType := priorStateData.CacheType.ValueString()
			finalValue = transformCacheTypeToValue(cacheType)
			valueSource = "typed cache_type field"
			tflog.Debug(ctx, "Transformed cache_type from typed state", map[string]interface{}{
				"cache_type": cacheType,
				"value":      finalValue,
			})
		} else if cacheType, exists := rawState["cache_type"]; exists {
			if cacheTypeStr, ok := cacheType.(string); ok {
				finalValue = transformCacheTypeToValue(cacheTypeStr)
				valueSource = "raw cache_type field"
				tflog.Debug(ctx, "Transformed cache_type from raw state", map[string]interface{}{
					"cache_type": cacheTypeStr,
					"value":      finalValue,
				})
			}
		}
	}

	// Set the final value
	if finalValue != "" {
		newState.Value = types.StringValue(finalValue)
		tflog.Info(ctx, "Successfully migrated tiered cache state", map[string]interface{}{
			"zone_id":      newState.ZoneID.ValueString(),
			"final_value":  finalValue,
			"value_source": valueSource,
		})
	} else {
		// Fallback to "off" if nothing found
		newState.Value = types.StringValue("off")
		tflog.Warn(ctx, "No cache_type or value found, defaulting to 'off'", map[string]interface{}{
			"zone_id": newState.ZoneID.ValueString(),
		})
	}

	// Set the upgraded state
	resp.Diagnostics.Append(resp.State.Set(ctx, &newState)...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to set upgraded state", map[string]interface{}{
			"errors": resp.Diagnostics.Errors(),
		})
		return
	}

	tflog.Debug(ctx, "Tiered cache state upgrade completed successfully")
}

// transformCacheTypeToValue transforms v4 cache_type values to v5 value values
func transformCacheTypeToValue(cacheType string) string {
	switch cacheType {
	case "smart", "generic":
		return "on"
	case "off":
		return "off"
	default:
		// Unknown value, default to "off"
		return "off"
	}
}

// getMapKeys returns the keys of a map for debugging
func getMapKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}