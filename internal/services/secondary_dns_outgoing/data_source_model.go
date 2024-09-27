// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secondary_dns_outgoing

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/secondary_dns"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SecondaryDNSOutgoingResultDataSourceEnvelope struct {
	Result SecondaryDNSOutgoingDataSourceModel `json:"result,computed"`
}

type SecondaryDNSOutgoingDataSourceModel struct {
	ZoneID              types.String    `tfsdk:"zone_id" path:"zone_id,required"`
	CheckedTime         types.String    `tfsdk:"checked_time" json:"checked_time,optional"`
	CreatedTime         types.String    `tfsdk:"created_time" json:"created_time,optional"`
	ID                  types.String    `tfsdk:"id" json:"id,optional"`
	LastTransferredTime types.String    `tfsdk:"last_transferred_time" json:"last_transferred_time,optional"`
	Name                types.String    `tfsdk:"name" json:"name,optional"`
	SOASerial           types.Float64   `tfsdk:"soa_serial" json:"soa_serial,optional"`
	Peers               *[]types.String `tfsdk:"peers" json:"peers,optional"`
}

func (m *SecondaryDNSOutgoingDataSourceModel) toReadParams(_ context.Context) (params secondary_dns.OutgoingGetParams, diags diag.Diagnostics) {
	params = secondary_dns.OutgoingGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
