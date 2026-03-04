package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x)
// ============================================================================

// SourceURLNormalizationSettingsModel represents the legacy resource state from v4.x provider.
// Schema version: 1 (v4 provider schema version)
// Resource type: cloudflare_url_normalization_settings
//
// Fields are identical between v4 and v5 - no transformations needed.
type SourceURLNormalizationSettingsModel struct {
	ID     types.String `tfsdk:"id"`
	ZoneID types.String `tfsdk:"zone_id"`
	Scope  types.String `tfsdk:"scope"`
	Type   types.String `tfsdk:"type"`
}

// ============================================================================
// Target Models (Current Provider - v5.x+)
// ============================================================================

// TargetURLNormalizationSettingsModel represents the current resource state from v5.x+ provider.
// Schema version: 500
// Resource type: cloudflare_url_normalization_settings
//
// Note: This matches the model in the parent package's model.go file.
type TargetURLNormalizationSettingsModel struct {
	ID     types.String `tfsdk:"id"`
	ZoneID types.String `tfsdk:"zone_id"`
	Scope  types.String `tfsdk:"scope"`
	Type   types.String `tfsdk:"type"`
}
