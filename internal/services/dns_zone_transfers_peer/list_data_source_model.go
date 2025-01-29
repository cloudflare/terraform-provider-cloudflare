// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_transfers_peer

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/dns"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DNSZoneTransfersPeersResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[DNSZoneTransfersPeersResultDataSourceModel] `json:"result,computed"`
}

type DNSZoneTransfersPeersDataSourceModel struct {
	AccountID types.String                                                             `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                              `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[DNSZoneTransfersPeersResultDataSourceModel] `tfsdk:"result"`
}

func (m *DNSZoneTransfersPeersDataSourceModel) toListParams(_ context.Context) (params dns.ZoneTransferPeerListParams, diags diag.Diagnostics) {
	params = dns.ZoneTransferPeerListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type DNSZoneTransfersPeersResultDataSourceModel struct {
	ID         types.String  `tfsdk:"id" json:"id,computed"`
	Name       types.String  `tfsdk:"name" json:"name,computed"`
	IP         types.String  `tfsdk:"ip" json:"ip,computed"`
	IxfrEnable types.Bool    `tfsdk:"ixfr_enable" json:"ixfr_enable,computed"`
	Port       types.Float64 `tfsdk:"port" json:"port,computed"`
	TSIGID     types.String  `tfsdk:"tsig_id" json:"tsig_id,computed"`
}
