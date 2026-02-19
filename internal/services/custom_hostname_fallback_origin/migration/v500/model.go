package v500

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x)
// ============================================================================

// SourceCustomHostnameFallbackOriginModel represents the legacy resource state from v4.x provider.
// Schema version: 0 (v4 had no explicit version)
// Resource type: cloudflare_custom_hostname_fallback_origin
//
// Note: v4 stored timestamps as plain strings, while v5 uses timetypes.RFC3339.
// Note: v4 used types.List for errors, while v5 uses customfield.List.
// This model uses v4 types, then Transform() converts to v5 types.
type SourceCustomHostnameFallbackOriginModel struct {
	ID        types.String `tfsdk:"id"`
	ZoneID    types.String `tfsdk:"zone_id"`
	Origin    types.String `tfsdk:"origin"`
	CreatedAt types.String `tfsdk:"created_at"`
	Status    types.String `tfsdk:"status"`
	UpdatedAt types.String `tfsdk:"updated_at"`
	Errors    types.List   `tfsdk:"errors"`
}

// ============================================================================
// Target Models (Current Provider - v5.x+)
// ============================================================================

// TargetCustomHostnameFallbackOriginModel represents the current resource state from v5.x+ provider.
// Schema version: 500
// Resource type: cloudflare_custom_hostname_fallback_origin
//
// Note: This model is IDENTICAL to the v4 model - no field changes occurred between v4 and v5.
// This is duplicated here to keep the migration package self-contained and follow the standard pattern.
type TargetCustomHostnameFallbackOriginModel struct {
	ID        types.String                   `tfsdk:"id"`
	ZoneID    types.String                   `tfsdk:"zone_id"`
	Origin    types.String                   `tfsdk:"origin"`
	CreatedAt timetypes.RFC3339              `tfsdk:"created_at"`
	Status    types.String                   `tfsdk:"status"`
	UpdatedAt timetypes.RFC3339              `tfsdk:"updated_at"`
	Errors    customfield.List[types.String] `tfsdk:"errors"`
}
