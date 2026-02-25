package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts source (legacy v4) state to target (current v5) state.
//
// Transformation summary:
// - Pass-through: zone_id, value (with defensive default)
// - ID: Set to zone_id (matches v5 resource behavior)
// - New computed fields: editable, modified_on (set to null, will refresh from API)
func Transform(ctx context.Context, source *SourceCloudflareRegionalTieredCacheModel) (*TargetRegionalTieredCacheModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Validate required fields
	if source.ZoneID.IsNull() || source.ZoneID.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"zone_id is required for regional_tiered_cache migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	// Initialize target model with direct copies and transformations
	target := &TargetRegionalTieredCacheModel{
		// Core identifier - v5 resource sets ID = ZoneID (see resource.go lines 94, 143, 183)
		ID:     source.ZoneID,
		ZoneID: source.ZoneID,

		// Value field - direct copy with defensive default
		Value: source.Value,

		// New computed fields in v5 - set to null, will refresh from API on first apply
		Editable:   types.BoolNull(),
		ModifiedOn: timetypes.NewRFC3339Null(),
	}

	// Defensive: If value is missing (shouldn't happen in v4 since it's required),
	// default to "off" matching tf-migrate behavior and v5 default
	if target.Value.IsNull() || target.Value.IsUnknown() {
		target.Value = types.StringValue("off")
	}

	return target, diags
}
