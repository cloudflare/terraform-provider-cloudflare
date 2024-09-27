// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secondary_dns_peer

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/secondary_dns"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SecondaryDNSPeersResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[SecondaryDNSPeersResultDataSourceModel] `json:"result,computed"`
}

type SecondaryDNSPeersDataSourceModel struct {
	AccountID types.String                                                         `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                          `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[SecondaryDNSPeersResultDataSourceModel] `tfsdk:"result"`
}

func (m *SecondaryDNSPeersDataSourceModel) toListParams(_ context.Context) (params secondary_dns.PeerListParams, diags diag.Diagnostics) {
	params = secondary_dns.PeerListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type SecondaryDNSPeersResultDataSourceModel struct {
	ID         types.String  `tfsdk:"id" json:"id,computed"`
	Name       types.String  `tfsdk:"name" json:"name,computed"`
	IP         types.String  `tfsdk:"ip" json:"ip,computed"`
	IxfrEnable types.Bool    `tfsdk:"ixfr_enable" json:"ixfr_enable,computed"`
	Port       types.Float64 `tfsdk:"port" json:"port,computed"`
	TSIGID     types.String  `tfsdk:"tsig_id" json:"tsig_id,computed"`
}
