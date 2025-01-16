// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_transfers_peer

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/dns"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DNSZoneTransfersPeerResultDataSourceEnvelope struct {
	Result DNSZoneTransfersPeerDataSourceModel `json:"result,computed"`
}

type DNSZoneTransfersPeerResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[DNSZoneTransfersPeerDataSourceModel] `json:"result,computed"`
}

type DNSZoneTransfersPeerDataSourceModel struct {
	AccountID  types.String                                  `tfsdk:"account_id" path:"account_id,optional"`
	PeerID     types.String                                  `tfsdk:"peer_id" path:"peer_id,optional"`
	ID         types.String                                  `tfsdk:"id" json:"id,computed"`
	IP         types.String                                  `tfsdk:"ip" json:"ip,computed"`
	IxfrEnable types.Bool                                    `tfsdk:"ixfr_enable" json:"ixfr_enable,computed"`
	Name       types.String                                  `tfsdk:"name" json:"name,computed"`
	Port       types.Float64                                 `tfsdk:"port" json:"port,computed"`
	TSIGID     types.String                                  `tfsdk:"tsig_id" json:"tsig_id,computed"`
	Filter     *DNSZoneTransfersPeerFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *DNSZoneTransfersPeerDataSourceModel) toReadParams(_ context.Context) (params dns.ZoneTransferPeerGetParams, diags diag.Diagnostics) {
	params = dns.ZoneTransferPeerGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *DNSZoneTransfersPeerDataSourceModel) toListParams(_ context.Context) (params dns.ZoneTransferPeerListParams, diags diag.Diagnostics) {
	params = dns.ZoneTransferPeerListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type DNSZoneTransfersPeerFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
}
