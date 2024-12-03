// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_transfer_peer

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/dns"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DNSZoneTransferPeersResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[DNSZoneTransferPeersResultDataSourceModel] `json:"result,computed"`
}

type DNSZoneTransferPeersDataSourceModel struct {
	AccountID types.String                                                            `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                             `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[DNSZoneTransferPeersResultDataSourceModel] `tfsdk:"result"`
}

func (m *DNSZoneTransferPeersDataSourceModel) toListParams(_ context.Context) (params dns.ZoneTransferPeerListParams, diags diag.Diagnostics) {
	params = dns.ZoneTransferPeerListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type DNSZoneTransferPeersResultDataSourceModel struct {
	ID         types.String  `tfsdk:"id" json:"id,computed"`
	Name       types.String  `tfsdk:"name" json:"name,computed"`
	IP         types.String  `tfsdk:"ip" json:"ip,computed"`
	IxfrEnable types.Bool    `tfsdk:"ixfr_enable" json:"ixfr_enable,computed"`
	Port       types.Float64 `tfsdk:"port" json:"port,computed"`
	TSIGID     types.String  `tfsdk:"tsig_id" json:"tsig_id,computed"`
}
