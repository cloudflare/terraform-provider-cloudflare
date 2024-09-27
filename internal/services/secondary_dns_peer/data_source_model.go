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

type SecondaryDNSPeerResultDataSourceEnvelope struct {
	Result SecondaryDNSPeerDataSourceModel `json:"result,computed"`
}

type SecondaryDNSPeerResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[SecondaryDNSPeerDataSourceModel] `json:"result,computed"`
}

type SecondaryDNSPeerDataSourceModel struct {
	AccountID  types.String                              `tfsdk:"account_id" path:"account_id,optional"`
	PeerID     types.String                              `tfsdk:"peer_id" path:"peer_id,optional"`
	ID         types.String                              `tfsdk:"id" json:"id,computed"`
	IP         types.String                              `tfsdk:"ip" json:"ip,computed"`
	IxfrEnable types.Bool                                `tfsdk:"ixfr_enable" json:"ixfr_enable,computed"`
	Name       types.String                              `tfsdk:"name" json:"name,computed"`
	Port       types.Float64                             `tfsdk:"port" json:"port,computed"`
	TSIGID     types.String                              `tfsdk:"tsig_id" json:"tsig_id,computed"`
	Filter     *SecondaryDNSPeerFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *SecondaryDNSPeerDataSourceModel) toReadParams(_ context.Context) (params secondary_dns.PeerGetParams, diags diag.Diagnostics) {
	params = secondary_dns.PeerGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *SecondaryDNSPeerDataSourceModel) toListParams(_ context.Context) (params secondary_dns.PeerListParams, diags diag.Diagnostics) {
	params = secondary_dns.PeerListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type SecondaryDNSPeerFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
}
