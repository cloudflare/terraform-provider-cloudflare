// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_transfers_peer

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/dns"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type DNSZoneTransfersPeerResultDataSourceEnvelope struct {
Result DNSZoneTransfersPeerDataSourceModel `json:"result,computed"`
}

type DNSZoneTransfersPeerDataSourceModel struct {
ID types.String `tfsdk:"id" json:"-,computed"`
PeerID types.String `tfsdk:"peer_id" path:"peer_id,optional"`
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
IP types.String `tfsdk:"ip" json:"ip,computed"`
IxfrEnable types.Bool `tfsdk:"ixfr_enable" json:"ixfr_enable,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
Port types.Float64 `tfsdk:"port" json:"port,computed"`
TSIGID types.String `tfsdk:"tsig_id" json:"tsig_id,computed"`
}

func (m *DNSZoneTransfersPeerDataSourceModel) toReadParams(_ context.Context) (params dns.ZoneTransferPeerGetParams, diags diag.Diagnostics) {
  params = dns.ZoneTransferPeerGetParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  return
}

func (m *DNSZoneTransfersPeerDataSourceModel) toListParams(_ context.Context) (params dns.ZoneTransferPeerListParams, diags diag.Diagnostics) {
  params = dns.ZoneTransferPeerListParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  return
}
