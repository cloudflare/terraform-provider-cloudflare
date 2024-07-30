// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package regional_hostname

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
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
	Hostname  types.String                              `tfsdk:"hostname" path:"hostname"`
	CreatedOn timetypes.RFC3339                         `tfsdk:"created_on" json:"created_on,computed"`
	RegionKey types.String                              `tfsdk:"region_key" json:"region_key,computed"`
	Filter    *RegionalHostnameFindOneByDataSourceModel `tfsdk:"filter"`
}

type RegionalHostnameFindOneByDataSourceModel struct {
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id"`
}
