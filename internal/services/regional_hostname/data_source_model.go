// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package regional_hostname

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/addressing"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RegionalHostnameResultDataSourceEnvelope struct {
	Result RegionalHostnameDataSourceModel `json:"result,computed"`
}

type RegionalHostnameResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[RegionalHostnameDataSourceModel] `json:"result,computed"`
}

type RegionalHostnameDataSourceModel struct {
	ZoneID    types.String                              `tfsdk:"zone_id" path:"zone_id"`
	Hostname  types.String                              `tfsdk:"hostname" path:"hostname,computed_optional"`
	CreatedOn timetypes.RFC3339                         `tfsdk:"created_on" json:"created_on,computed"`
	RegionKey types.String                              `tfsdk:"region_key" json:"region_key,computed"`
	Filter    *RegionalHostnameFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *RegionalHostnameDataSourceModel) toReadParams() (params addressing.RegionalHostnameGetParams, diags diag.Diagnostics) {
	params = addressing.RegionalHostnameGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

func (m *RegionalHostnameDataSourceModel) toListParams() (params addressing.RegionalHostnameListParams, diags diag.Diagnostics) {
	params = addressing.RegionalHostnameListParams{
		ZoneID: cloudflare.F(m.Filter.ZoneID.ValueString()),
	}

	return
}

type RegionalHostnameFindOneByDataSourceModel struct {
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id"`
}
