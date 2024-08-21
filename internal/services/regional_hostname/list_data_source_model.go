// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package regional_hostname

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/addressing"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RegionalHostnamesResultListDataSourceEnvelope struct {
	Result *[]*RegionalHostnamesResultDataSourceModel `json:"result,computed"`
}

type RegionalHostnamesDataSourceModel struct {
	ZoneID   types.String                               `tfsdk:"zone_id" path:"zone_id"`
	MaxItems types.Int64                                `tfsdk:"max_items"`
	Result   *[]*RegionalHostnamesResultDataSourceModel `tfsdk:"result"`
}

func (m *RegionalHostnamesDataSourceModel) toListParams() (params addressing.RegionalHostnameListParams, diags diag.Diagnostics) {
	params = addressing.RegionalHostnameListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type RegionalHostnamesResultDataSourceModel struct {
	CreatedOn timetypes.RFC3339 `tfsdk:"created_on" json:"created_on,computed"`
	Hostname  types.String      `tfsdk:"hostname" json:"hostname,computed"`
	RegionKey types.String      `tfsdk:"region_key" json:"region_key,computed"`
}
