// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package regional_hostname

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RegionalHostnameResultDataSourceEnvelope struct {
	Result RegionalHostnameDataSourceModel `json:"result,computed"`
}

type RegionalHostnameResultListDataSourceEnvelope struct {
	Result *[]*RegionalHostnameDataSourceModel `json:"result,computed"`
}

type RegionalHostnameDataSourceModel struct {
	ZoneID    types.String                              `tfsdk:"zone_id" path:"zone_id"`
	Hostname  types.String                              `tfsdk:"hostname" json:"hostname"`
	CreatedOn types.String                              `tfsdk:"created_on" json:"created_on,computed"`
	RegionKey types.String                              `tfsdk:"region_key" json:"region_key,computed"`
	FindOneBy *RegionalHostnameFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

type RegionalHostnameFindOneByDataSourceModel struct {
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id"`
}
