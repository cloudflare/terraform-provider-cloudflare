// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_transfers_outgoing

import (
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type DNSZoneTransfersOutgoingResultEnvelope struct {
Result DNSZoneTransfersOutgoingModel `json:"result"`
}

type DNSZoneTransfersOutgoingModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
Name types.String `tfsdk:"name" json:"name,required"`
Peers *[]types.String `tfsdk:"peers" json:"peers,required"`
CheckedTime types.String `tfsdk:"checked_time" json:"checked_time,computed"`
CreatedTime types.String `tfsdk:"created_time" json:"created_time,computed"`
LastTransferredTime types.String `tfsdk:"last_transferred_time" json:"last_transferred_time,computed"`
SOASerial types.Float64 `tfsdk:"soa_serial" json:"soa_serial,computed"`
}

func (m DNSZoneTransfersOutgoingModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m DNSZoneTransfersOutgoingModel) MarshalJSONForUpdate(state DNSZoneTransfersOutgoingModel) (data []byte, err error) {
  return apijson.MarshalForUpdate(m, state)
}
