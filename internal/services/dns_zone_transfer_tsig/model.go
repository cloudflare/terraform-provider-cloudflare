// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_transfer_tsig

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DNSZoneTransferTSIGResultEnvelope struct {
	Result DNSZoneTransferTSIGModel `json:"result"`
}

type DNSZoneTransferTSIGModel struct {
	ID        types.String `tfsdk:"id" json:"id,computed"`
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
	Algo      types.String `tfsdk:"algo" json:"algo,required"`
	Name      types.String `tfsdk:"name" json:"name,required"`
	Secret    types.String `tfsdk:"secret" json:"secret,required"`
}

func (m DNSZoneTransferTSIGModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m DNSZoneTransferTSIGModel) MarshalJSONForUpdate(state DNSZoneTransferTSIGModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
