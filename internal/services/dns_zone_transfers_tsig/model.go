// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_transfers_tsig

import (
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DNSZoneTransfersTSIGResultEnvelope struct {
	Result DNSZoneTransfersTSIGModel `json:"result"`
}

type DNSZoneTransfersTSIGModel struct {
	ID        types.String `tfsdk:"id" json:"id,computed"`
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
	Algo      types.String `tfsdk:"algo" json:"algo,required"`
	Name      types.String `tfsdk:"name" json:"name,required"`
	Secret    types.String `tfsdk:"secret" json:"secret,required"`
}

func (m DNSZoneTransfersTSIGModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m DNSZoneTransfersTSIGModel) MarshalJSONForUpdate(state DNSZoneTransfersTSIGModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
