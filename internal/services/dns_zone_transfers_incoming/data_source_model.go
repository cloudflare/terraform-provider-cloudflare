// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_transfers_incoming

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/dns"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DNSZoneTransfersIncomingResultDataSourceEnvelope struct {
	Result DNSZoneTransfersIncomingDataSourceModel `json:"result,computed"`
}

type DNSZoneTransfersIncomingDataSourceModel struct {
	ZoneID             types.String                   `tfsdk:"zone_id" path:"zone_id,required"`
	AutoRefreshSeconds types.Float64                  `tfsdk:"auto_refresh_seconds" json:"auto_refresh_seconds,computed"`
	CheckedTime        types.String                   `tfsdk:"checked_time" json:"checked_time,computed"`
	CreatedTime        types.String                   `tfsdk:"created_time" json:"created_time,computed"`
	ID                 types.String                   `tfsdk:"id" json:"id,computed"`
	ModifiedTime       types.String                   `tfsdk:"modified_time" json:"modified_time,computed"`
	Name               types.String                   `tfsdk:"name" json:"name,computed"`
	SOASerial          types.Float64                  `tfsdk:"soa_serial" json:"soa_serial,computed"`
	Peers              customfield.List[types.String] `tfsdk:"peers" json:"peers,computed"`
}

func (m *DNSZoneTransfersIncomingDataSourceModel) toReadParams(_ context.Context) (params dns.ZoneTransferIncomingGetParams, diags diag.Diagnostics) {
	params = dns.ZoneTransferIncomingGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
