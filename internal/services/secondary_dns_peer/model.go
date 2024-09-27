// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secondary_dns_peer

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SecondaryDNSPeerResultEnvelope struct {
	Result SecondaryDNSPeerModel `json:"result"`
}

type SecondaryDNSPeerModel struct {
	ID         types.String  `tfsdk:"id" json:"id,computed"`
	AccountID  types.String  `tfsdk:"account_id" path:"account_id,required"`
	Name       types.String  `tfsdk:"name" json:"name,required"`
	IP         types.String  `tfsdk:"ip" json:"ip,computed_optional"`
	IxfrEnable types.Bool    `tfsdk:"ixfr_enable" json:"ixfr_enable,computed_optional"`
	Port       types.Float64 `tfsdk:"port" json:"port,computed_optional"`
	TSIGID     types.String  `tfsdk:"tsig_id" json:"tsig_id,computed_optional"`
}
