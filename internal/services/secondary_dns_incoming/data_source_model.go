// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secondary_dns_incoming

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/secondary_dns"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SecondaryDNSIncomingResultDataSourceEnvelope struct {
	Result SecondaryDNSIncomingDataSourceModel `json:"result,computed"`
}

type SecondaryDNSIncomingDataSourceModel struct {
	ZoneID             types.String    `tfsdk:"zone_id" path:"zone_id,required"`
	AutoRefreshSeconds types.Float64   `tfsdk:"auto_refresh_seconds" json:"auto_refresh_seconds,optional"`
	CheckedTime        types.String    `tfsdk:"checked_time" json:"checked_time,optional"`
	CreatedTime        types.String    `tfsdk:"created_time" json:"created_time,optional"`
	ID                 types.String    `tfsdk:"id" json:"id,optional"`
	ModifiedTime       types.String    `tfsdk:"modified_time" json:"modified_time,optional"`
	Name               types.String    `tfsdk:"name" json:"name,optional"`
	SOASerial          types.Float64   `tfsdk:"soa_serial" json:"soa_serial,optional"`
	Peers              *[]types.String `tfsdk:"peers" json:"peers,optional"`
}

func (m *SecondaryDNSIncomingDataSourceModel) toReadParams(_ context.Context) (params secondary_dns.IncomingGetParams, diags diag.Diagnostics) {
	params = secondary_dns.IncomingGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
