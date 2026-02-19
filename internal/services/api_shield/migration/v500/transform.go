// Package v500 handles state migration from cloudflare_api_shield v4 (schema_version=0)
// to cloudflare_api_shield v5 (version=500).
package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// Transform converts source (legacy v4) state to target (current v5) state.
//
// Transformations:
//   - id: Direct copy (computed, set to zone_id)
//   - zone_id: Direct copy (required)
//   - auth_id_characteristics: Handle Optional → Required
//     - If null/missing in v4: Set to empty array in v5
//     - If present in v4: Direct copy (state structure identical)
//
// This function is called by UpgradeFromV4 handler.
func Transform(ctx context.Context, source SourceAPIShieldModel) (*TargetAPIShieldModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Step 1: Validate required fields
	if source.ZoneID.IsNull() || source.ZoneID.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"zone_id is required for api_shield migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	// Step 2: Initialize target with direct copies
	target := &TargetAPIShieldModel{
		ID:     source.ID,
		ZoneID: source.ZoneID,
	}

	// Step 3: Handle auth_id_characteristics (Optional in v4 → Required in v5)
	//
	// v4: Optional field, can be null
	// v5: Required field, must have value
	//
	// State structure is identical in both versions (array of objects),
	// only HCL syntax differs (block vs attribute).
	if source.AuthIDCharacteristics != nil && len(*source.AuthIDCharacteristics) > 0 {
		// v4 has auth_id_characteristics: Convert to target type
		// Field values are identical, just need to convert model types
		targetCharacteristics := make([]*TargetAuthIDCharacteristicsModel, 0, len(*source.AuthIDCharacteristics))
		for _, sourceChar := range *source.AuthIDCharacteristics {
			targetChar := &TargetAuthIDCharacteristicsModel{
				Name: sourceChar.Name,
				Type: sourceChar.Type,
			}
			targetCharacteristics = append(targetCharacteristics, targetChar)
		}
		target.AuthIDCharacteristics = &targetCharacteristics
	} else {
		// v4 missing auth_id_characteristics: Set empty array for v5 Required field
		// This satisfies v5 schema requirement while preserving semantic meaning
		emptyArray := make([]*TargetAuthIDCharacteristicsModel, 0)
		target.AuthIDCharacteristics = &emptyArray
	}

	return target, diags
}
