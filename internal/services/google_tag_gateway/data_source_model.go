// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package google_tag_gateway

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/google_tag_gateway"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type GoogleTagGatewayResultDataSourceEnvelope struct {
	Result GoogleTagGatewayDataSourceModel `json:"result,computed"`
}

type GoogleTagGatewayDataSourceModel struct {
	ID             types.String `tfsdk:"id" path:"zone_id,computed"`
	ZoneID         types.String `tfsdk:"zone_id" path:"zone_id,required"`
	Enabled        types.Bool   `tfsdk:"enabled" json:"enabled,computed"`
	Endpoint       types.String `tfsdk:"endpoint" json:"endpoint,computed"`
	HideOriginalIP types.Bool   `tfsdk:"hide_original_ip" json:"hideOriginalIp,computed"`
	MeasurementID  types.String `tfsdk:"measurement_id" json:"measurementId,computed"`
	SetUpTag       types.Bool   `tfsdk:"set_up_tag" json:"setUpTag,computed"`
}

func (m *GoogleTagGatewayDataSourceModel) toReadParams(_ context.Context) (params google_tag_gateway.ConfigGetParams, diags diag.Diagnostics) {
	params = google_tag_gateway.ConfigGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
