package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x)
// ============================================================================

// SourceTieredCacheModel represents the legacy resource state from v4.x provider.
// Schema version: 0 (SDKv2 default)
// Resource type: cloudflare_tiered_cache
//
// In v4, the tiered_cache resource used cache_type with values: "smart", "generic", "off"
type SourceTieredCacheModel struct {
	ID         types.String      `tfsdk:"id"`
	ZoneID     types.String      `tfsdk:"zone_id"`
	CacheType  types.String      `tfsdk:"cache_type"`
	Editable   types.Bool        `tfsdk:"editable"`
	ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on"`
}

// ============================================================================
// Target Models (Current Provider - v5.x+)
// ============================================================================

// TargetTieredCacheModel represents the current resource state from v5.x+ provider.
// Schema version: 500
// Resource type: cloudflare_tiered_cache
//
// In v5, the tiered_cache resource uses value with values: "on", "off"
// Note: This should match the model in the parent package's model.go file.
type TargetTieredCacheModel struct {
	ID         types.String      `tfsdk:"id"`
	ZoneID     types.String      `tfsdk:"zone_id"`
	Value      types.String      `tfsdk:"value"`
	Editable   types.Bool        `tfsdk:"editable"`
	ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on"`
}
