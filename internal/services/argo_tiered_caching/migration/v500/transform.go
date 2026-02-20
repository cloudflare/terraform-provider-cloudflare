package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// TransformArgoToTieredCaching converts source (legacy v4 cloudflare_argo) state to target (current v5 cloudflare_argo_tiered_caching) state.
// This function is shared by both UpgradeArgoToTieredCaching and MoveArgoToTieredCaching handlers.
//
// Transformation logic:
// - zone_id: Direct copy
// - tiered_caching → value: Rename
// - smart_routing: Removed (goes to separate resource)
// - id: Change from checksum to zone_id
// - editable: Add as true
// - modified_on: Add as null (will be populated on refresh)
func TransformArgoToTieredCaching(ctx context.Context, source SourceCloudflareArgoModel) (*TargetArgoTieredCachingModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Validate required fields
	if source.ZoneID.IsNull() || source.ZoneID.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"zone_id is required for argo migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	// Validate tiered_caching exists (this transformation only happens for scenario 4)
	if source.TieredCaching.IsNull() || source.TieredCaching.IsUnknown() || source.TieredCaching.ValueString() == "" {
		diags.AddError(
			"Missing tiered_caching attribute",
			"tiered_caching attribute is required for transforming to argo_tiered_caching. The source state is missing this field.",
		)
		return nil, diags
	}

	// Initialize target
	target := &TargetArgoTieredCachingModel{
		ZoneID: source.ZoneID,
		// ID changes from checksum to zone_id
		ID: source.ZoneID,
		// tiered_caching → value
		Value: source.TieredCaching,
		// Add computed fields
		Editable:   types.BoolValue(true),
		ModifiedOn: timetypes.NewRFC3339Null(),
	}

	// Note: source.SmartRouting is intentionally not copied
	// It goes to a separate cloudflare_argo_smart_routing resource

	return target, diags
}
