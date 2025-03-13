// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_transfers_acl

import (
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type DNSZoneTransfersACLResultEnvelope struct {
Result DNSZoneTransfersACLModel `json:"result"`
}

type DNSZoneTransfersACLModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
IPRange types.String `tfsdk:"ip_range" json:"ip_range,required"`
Name types.String `tfsdk:"name" json:"name,required"`
}

func (m DNSZoneTransfersACLModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m DNSZoneTransfersACLModel) MarshalJSONForUpdate(state DNSZoneTransfersACLModel) (data []byte, err error) {
  return apijson.MarshalForUpdate(m, state)
}
