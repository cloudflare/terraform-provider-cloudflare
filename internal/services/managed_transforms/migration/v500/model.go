package v500

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x: cloudflare_managed_headers)
// ============================================================================

// SourceManagedHeadersModel represents the legacy resource state from v4.x provider.
// Schema version: 0 (v4 provider default)
// Resource type: cloudflare_managed_headers (renamed to cloudflare_managed_transforms in v5)
//
// Key differences from v5:
// - managed_request_headers and managed_response_headers were optional (could be null)
// - In v5, both fields are required (empty sets allowed)
type SourceManagedHeadersModel struct {
	ID                     types.String                `tfsdk:"id"`
	ZoneID                 types.String                `tfsdk:"zone_id"`
	ManagedRequestHeaders  *[]*SourceHeaderEntryModel  `tfsdk:"managed_request_headers"`
	ManagedResponseHeaders *[]*SourceHeaderEntryModel  `tfsdk:"managed_response_headers"`
}

type SourceHeaderEntryModel struct {
	ID      types.String `tfsdk:"id"`
	Enabled types.Bool   `tfsdk:"enabled"`
}

// ============================================================================
// Target Models (Current Provider - v5.x+: cloudflare_managed_transforms)
// ============================================================================

// TargetManagedTransformsModel represents the current resource state from v5.x+ provider.
// Schema version: 500
// Resource type: cloudflare_managed_transforms
type TargetManagedTransformsModel struct {
	ID                     types.String                                                                  `tfsdk:"id"`
	ZoneID                 types.String                                                                  `tfsdk:"zone_id"`
	ManagedRequestHeaders  customfield.NestedObjectSet[TargetHeaderEntryModel]  `tfsdk:"managed_request_headers"`
	ManagedResponseHeaders customfield.NestedObjectSet[TargetHeaderEntryModel] `tfsdk:"managed_response_headers"`
}

type TargetHeaderEntryModel struct {
	ID      types.String `tfsdk:"id" json:"id,required"`
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled,required"`
}
