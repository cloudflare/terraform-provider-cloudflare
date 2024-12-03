// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_transfer_acl

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DNSZoneTransferACLResultEnvelope struct {
	Result DNSZoneTransferACLModel `json:"result"`
}

type DNSZoneTransferACLModel struct {
	ID        types.String `tfsdk:"id" json:"id,computed"`
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
	IPRange   types.String `tfsdk:"ip_range" json:"ip_range,required"`
	Name      types.String `tfsdk:"name" json:"name,required"`
}

func (m DNSZoneTransferACLModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m DNSZoneTransferACLModel) MarshalJSONForUpdate(state DNSZoneTransferACLModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
