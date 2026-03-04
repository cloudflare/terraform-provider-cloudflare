package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceRegionalHostnameModel represents the legacy resource state from v4.x provider.
// Schema version: 0 (v4 provider default)
// Resource type: cloudflare_regional_hostname
//
// Fields are identical between v4 and v5 - no transformations needed.
type SourceRegionalHostnameModel struct {
	ID        types.String `tfsdk:"id"`
	Hostname  types.String `tfsdk:"hostname"`
	ZoneID    types.String `tfsdk:"zone_id"`
	Routing   types.String `tfsdk:"routing"`
	RegionKey types.String `tfsdk:"region_key"`
	CreatedOn types.String `tfsdk:"created_on"`
}

// TargetRegionalHostnameModel represents the current resource state from v5.x+ provider.
// Schema version: 500
// Resource type: cloudflare_regional_hostname
type TargetRegionalHostnameModel struct {
	ID        types.String `tfsdk:"id"`
	Hostname  types.String `tfsdk:"hostname"`
	ZoneID    types.String `tfsdk:"zone_id"`
	Routing   types.String `tfsdk:"routing"`
	RegionKey types.String `tfsdk:"region_key"`
	CreatedOn types.String `tfsdk:"created_on"`
}
