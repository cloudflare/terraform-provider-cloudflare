package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x)
// ============================================================================

// SourceCloudflareArgoModel represents the legacy cloudflare_argo state from v4.x provider.
// Schema version: 0 (SDK v2 default)
// Resource type: cloudflare_argo
//
// This resource could have smart_routing, tiered_caching, or both attributes.
// The decision logic determines whether this becomes argo_smart_routing or argo_tiered_caching.
type SourceCloudflareArgoModel struct {
	ID             types.String `tfsdk:"id"`              // Checksum-based ID: stringChecksum(fmt.Sprintf("%s/argo", zoneID))
	ZoneID         types.String `tfsdk:"zone_id"`         // Zone identifier
	SmartRouting   types.String `tfsdk:"smart_routing"`   // Optional, "on" or "off"
	TieredCaching  types.String `tfsdk:"tiered_caching"`  // Optional, "on" or "off"
}

// ============================================================================
// Target Models (Current Provider - v5.x+)
// ============================================================================

// TargetArgoTieredCachingModel represents the current cloudflare_argo_tiered_caching state from v5.x+ provider.
// Schema version: 500
// Resource type: cloudflare_argo_tiered_caching
//
// This model matches the ArgoTieredCachingModel in the parent package's model.go file.
type TargetArgoTieredCachingModel struct {
	ID         types.String      `tfsdk:"id"`          // Zone ID (changed from checksum)
	ZoneID     types.String      `tfsdk:"zone_id"`     // Zone identifier
	Value      types.String      `tfsdk:"value"`       // Renamed from tiered_caching, required
	Editable   types.Bool        `tfsdk:"editable"`    // New computed field
	ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on"` // New computed field
}
