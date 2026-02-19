package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// TransformArgoToSmartRouting converts source (legacy v4 cloudflare_argo) state to target (current v5 cloudflare_argo_smart_routing) state.
// This function is shared by both UpgradeArgoToSmartRouting and MoveArgoToSmartRouting handlers.
//
// Transformation logic:
// - zone_id: Direct copy
// - smart_routing → value: Rename, default to "off" if missing
// - tiered_caching: Removed (goes to separate resource)
// - id: Change from checksum to zone_id
// - editable: Add as true
// - modified_on: Add as null (will be populated on refresh)
func TransformArgoToSmartRouting(ctx context.Context, source SourceCloudflareArgoModel) (*TargetArgoSmartRoutingModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Validate required fields
	if source.ZoneID.IsNull() || source.ZoneID.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"zone_id is required for argo migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	// Initialize target
	target := &TargetArgoSmartRoutingModel{
		ZoneID: source.ZoneID,
		// ID changes from checksum to zone_id
		ID: source.ZoneID,
		// Add computed fields
		Editable:   types.BoolValue(true),
		ModifiedOn: timetypes.NewRFC3339Null(),
	}

	// Handle smart_routing → value transformation
	// Default to "off" if smart_routing is missing/null/empty
	if !source.SmartRouting.IsNull() && !source.SmartRouting.IsUnknown() && source.SmartRouting.ValueString() != "" {
		target.Value = source.SmartRouting
	} else {
		target.Value = types.StringValue("off")
	}

	// Note: source.TieredCaching is intentionally not copied
	// It goes to a separate cloudflare_argo_tiered_caching resource

	return target, diags
}
