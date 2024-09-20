// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web3_hostname

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/web3"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Web3HostnamesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[Web3HostnamesResultDataSourceModel] `json:"result,computed"`
}

type Web3HostnamesDataSourceModel struct {
	ZoneID   types.String                                                     `tfsdk:"zone_id" path:"zone_id,required"`
	MaxItems types.Int64                                                      `tfsdk:"max_items"`
	Result   customfield.NestedObjectList[Web3HostnamesResultDataSourceModel] `tfsdk:"result"`
}

func (m *Web3HostnamesDataSourceModel) toListParams(_ context.Context) (params web3.HostnameListParams, diags diag.Diagnostics) {
	params = web3.HostnameListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type Web3HostnamesResultDataSourceModel struct {
	ID          types.String      `tfsdk:"id" json:"id,computed"`
	CreatedOn   timetypes.RFC3339 `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Description types.String      `tfsdk:"description" json:"description,computed"`
	Dnslink     types.String      `tfsdk:"dnslink" json:"dnslink,computed"`
	ModifiedOn  timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Name        types.String      `tfsdk:"name" json:"name,computed"`
	Status      types.String      `tfsdk:"status" json:"status,computed"`
	Target      types.String      `tfsdk:"target" json:"target,computed"`
}
