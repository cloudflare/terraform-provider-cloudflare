// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_transfer_outgoing

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/dns"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DNSZoneTransferOutgoingResultDataSourceEnvelope struct {
	Result DNSZoneTransferOutgoingDataSourceModel `json:"result,computed"`
}

type DNSZoneTransferOutgoingDataSourceModel struct {
	ZoneID              types.String                   `tfsdk:"zone_id" path:"zone_id,required"`
	CheckedTime         types.String                   `tfsdk:"checked_time" json:"checked_time,computed"`
	CreatedTime         types.String                   `tfsdk:"created_time" json:"created_time,computed"`
	ID                  types.String                   `tfsdk:"id" json:"id,computed"`
	LastTransferredTime types.String                   `tfsdk:"last_transferred_time" json:"last_transferred_time,computed"`
	Name                types.String                   `tfsdk:"name" json:"name,computed"`
	SOASerial           types.Float64                  `tfsdk:"soa_serial" json:"soa_serial,computed"`
	Peers               customfield.List[types.String] `tfsdk:"peers" json:"peers,computed"`
}

func (m *DNSZoneTransferOutgoingDataSourceModel) toReadParams(_ context.Context) (params dns.ZoneTransferOutgoingGetParams, diags diag.Diagnostics) {
	params = dns.ZoneTransferOutgoingGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
