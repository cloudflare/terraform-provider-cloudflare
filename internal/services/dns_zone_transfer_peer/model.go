// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_transfer_peer

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DNSZoneTransferPeerResultEnvelope struct {
	Result DNSZoneTransferPeerModel `json:"result"`
}

type DNSZoneTransferPeerModel struct {
	ID         types.String  `tfsdk:"id" json:"id,computed"`
	AccountID  types.String  `tfsdk:"account_id" path:"account_id,required"`
	Name       types.String  `tfsdk:"name" json:"name,required"`
	IP         types.String  `tfsdk:"ip" json:"ip,optional"`
	IxfrEnable types.Bool    `tfsdk:"ixfr_enable" json:"ixfr_enable,optional"`
	Port       types.Float64 `tfsdk:"port" json:"port,optional"`
	TSIGID     types.String  `tfsdk:"tsig_id" json:"tsig_id,optional"`
}

func (m DNSZoneTransferPeerModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m DNSZoneTransferPeerModel) MarshalJSONForUpdate(state DNSZoneTransferPeerModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
