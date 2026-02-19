package v500

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts source (v4 SDKv2) state to target (v5 Plugin Framework) state.
//
// Key transformation:
// - configuration: Unwraps array[0] → object (TypeList MaxItems:1 → SingleNestedAttribute)
//
// Pass-through fields (no transformation):
// - account_id, zone_id, mode, notes
//
// New computed fields (initialize to Null):
// - id, created_on, modified_on, allowed_modes, scope
func Transform(ctx context.Context, source SourceV4AccessRuleModel) (*TargetAccessRuleModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Initialize target model
	target := &TargetAccessRuleModel{}

	// Step 1: Pass-through fields (direct copy, no transformation)
	// Handle account_id/zone_id mutual exclusivity
	if !source.AccountID.IsNull() && source.AccountID.ValueString() != "" {
		// Account-level rule: set account_id, ensure zone_id is null
		target.AccountID = source.AccountID
		target.ZoneID = types.StringNull()
	} else if !source.ZoneID.IsNull() && source.ZoneID.ValueString() != "" {
		// Zone-level rule: set zone_id, ensure account_id is null
		target.ZoneID = source.ZoneID
		target.AccountID = types.StringNull()
	} else {
		// Fallback: copy as-is (shouldn't happen in valid v4 state)
		target.AccountID = source.AccountID
		target.ZoneID = source.ZoneID
	}

	target.Mode = source.Mode
	target.Notes = source.Notes

	// Step 2: CRITICAL TRANSFORMATION - Unwrap configuration array[0] → object
	// v4: configuration stored as array (TypeList MaxItems:1)
	// v5: configuration stored as object (SingleNestedAttribute)
	if len(source.Configuration) > 0 {
		// Extract first element from array
		sourceConfig := source.Configuration[0]

		// Create target configuration object
		target.Configuration = &TargetConfigurationModel{
			Target: sourceConfig.Target,
			Value:  sourceConfig.Value,
		}
	} else {
		// Edge case: Empty array (shouldn't happen, field is Required)
		// But handle gracefully
		target.Configuration = nil
	}

	// Step 3: Copy ID from v4 state (implicit in SDKv2 but present in state)
	target.ID = source.ID

	// Step 4: Initialize new v5 computed fields to Null
	// These fields don't exist in v4 state - API will populate them after apply
	target.CreatedOn = timetypes.NewRFC3339Null()
	target.ModifiedOn = timetypes.NewRFC3339Null()
	target.AllowedModes = customfield.NullList[types.String](ctx)
	target.Scope = customfield.NullObject[TargetScopeModel](ctx)

	return target, diags
}
