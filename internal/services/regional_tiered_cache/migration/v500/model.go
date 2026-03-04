package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x)
// ============================================================================

// SourceCloudflareRegionalTieredCacheModel represents the source cloudflare_regional_tiered_cache state structure.
// This corresponds to schema_version=0 from the legacy (SDKv2) cloudflare provider.
// Used by UpgradeFromLegacyV0 to parse legacy state.
//
// Resource type: cloudflare_regional_tiered_cache (NOT renamed in v5)
// Schema version: 0 (default SDKv2 schema version)
//
// Note: This resource has a very simple structure with only two configurable fields.
type SourceCloudflareRegionalTieredCacheModel struct {
	// Core identifier
	// Note: In v4, ID is set to zone_id by the provider (see v4 resource.go line 42)
	ID types.String `tfsdk:"id"`

	// Required fields in v4
	ZoneID types.String `tfsdk:"zone_id"` // Zone identifier
	Value  types.String `tfsdk:"value"`   // "on" or "off" - Required in v4

	// v4 DOES NOT have these fields (they are new in v5):
	// - editable (Bool, computed)
	// - modified_on (RFC3339, computed)
}

// ============================================================================
// Target Models (Current Provider - v5.x+)
// ============================================================================

// TargetRegionalTieredCacheModel represents the target cloudflare_regional_tiered_cache state structure (v500).
// This corresponds to schema_version=500 in the current (Framework) provider.
//
// Resource type: cloudflare_regional_tiered_cache (same as v4, NOT renamed)
// Schema version: 500
//
// Note: This should match the model in ../../model.go (parent package).
// We duplicate it here to keep the migration package self-contained.
type TargetRegionalTieredCacheModel struct {
	// Core identifiers
	ID     types.String `tfsdk:"id"`      // Set to zone_id (matches v5 resource behavior)
	ZoneID types.String `tfsdk:"zone_id"` // Zone identifier

	// Configurable field (Changed from Required in v4 to Optional+Computed in v5)
	Value types.String `tfsdk:"value"` // "on" or "off" - Optional with default "off" in v5

	// New computed fields in v5 (did not exist in v4)
	Editable   types.Bool        `tfsdk:"editable"`    // Whether the setting is editable (API-assigned)
	ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on"` // Last modification timestamp (API-assigned)
}
