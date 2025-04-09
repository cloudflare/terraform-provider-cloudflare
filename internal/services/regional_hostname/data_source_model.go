// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package regional_hostname

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/addressing"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RegionalHostnameResultDataSourceEnvelope struct {
	Result RegionalHostnameDataSourceModel `json:"result,computed"`
}

type RegionalHostnameDataSourceModel struct {
	ID        types.String      `tfsdk:"id" path:"hostname,computed"`
	Hostname  types.String      `tfsdk:"hostname" path:"hostname,computed_optional"`
	ZoneID    types.String      `tfsdk:"zone_id" path:"zone_id,required"`
	CreatedOn timetypes.RFC3339 `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	RegionKey types.String      `tfsdk:"region_key" json:"region_key,computed"`
}

func (m *RegionalHostnameDataSourceModel) toReadParams(_ context.Context) (params addressing.RegionalHostnameGetParams, diags diag.Diagnostics) {
	params = addressing.RegionalHostnameGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
