// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web3_hostname

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/web3"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Web3HostnameResultDataSourceEnvelope struct {
	Result Web3HostnameDataSourceModel `json:"result,computed"`
}

type Web3HostnameDataSourceModel struct {
	ID          types.String      `tfsdk:"id" path:"identifier,computed"`
	Identifier  types.String      `tfsdk:"identifier" path:"identifier,optional"`
	ZoneID      types.String      `tfsdk:"zone_id" path:"zone_id,required"`
	CreatedOn   timetypes.RFC3339 `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Description types.String      `tfsdk:"description" json:"description,computed"`
	Dnslink     types.String      `tfsdk:"dnslink" json:"dnslink,computed"`
	ModifiedOn  timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Name        types.String      `tfsdk:"name" json:"name,computed"`
	Status      types.String      `tfsdk:"status" json:"status,computed"`
	Target      types.String      `tfsdk:"target" json:"target,computed"`
}

func (m *Web3HostnameDataSourceModel) toReadParams(_ context.Context) (params web3.HostnameGetParams, diags diag.Diagnostics) {
	params = web3.HostnameGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
