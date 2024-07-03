// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package regional_hostname

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RegionalHostnamesResultListDataSourceEnvelope struct {
	Result *[]*RegionalHostnamesItemsDataSourceModel `json:"result,computed"`
}

type RegionalHostnamesDataSourceModel struct {
	ZoneID   types.String                              `tfsdk:"zone_id" path:"zone_id"`
	MaxItems types.Int64                               `tfsdk:"max_items"`
	Items    *[]*RegionalHostnamesItemsDataSourceModel `tfsdk:"items"`
}

type RegionalHostnamesItemsDataSourceModel struct {
	CreatedOn types.String `tfsdk:"created_on" json:"created_on,computed"`
	Hostname  types.String `tfsdk:"hostname" json:"hostname,computed"`
	RegionKey types.String `tfsdk:"region_key" json:"region_key,computed"`
}
