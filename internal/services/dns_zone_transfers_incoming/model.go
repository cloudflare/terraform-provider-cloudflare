// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_transfers_incoming

import (
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DNSZoneTransfersIncomingResultEnvelope struct {
	Result DNSZoneTransfersIncomingModel `json:"result"`
}

type DNSZoneTransfersIncomingModel struct {
	ID                 types.String    `tfsdk:"id" json:"id,computed"`
	ZoneID             types.String    `tfsdk:"zone_id" path:"zone_id,required"`
	AutoRefreshSeconds types.Float64   `tfsdk:"auto_refresh_seconds" json:"auto_refresh_seconds,required"`
	Name               types.String    `tfsdk:"name" json:"name,required"`
	Peers              *[]types.String `tfsdk:"peers" json:"peers,required"`
	CheckedTime        types.String    `tfsdk:"checked_time" json:"checked_time,computed"`
	CreatedTime        types.String    `tfsdk:"created_time" json:"created_time,computed"`
	ModifiedTime       types.String    `tfsdk:"modified_time" json:"modified_time,computed"`
	SOASerial          types.Float64   `tfsdk:"soa_serial" json:"soa_serial,computed"`
}

func (m DNSZoneTransfersIncomingModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m DNSZoneTransfersIncomingModel) MarshalJSONForUpdate(state DNSZoneTransfersIncomingModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
