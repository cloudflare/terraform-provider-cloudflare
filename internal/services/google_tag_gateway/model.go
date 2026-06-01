// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package google_tag_gateway

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type GoogleTagGatewayResultEnvelope struct {
	Result GoogleTagGatewayModel `json:"result"`
}

type GoogleTagGatewayModel struct {
	ID             types.String `tfsdk:"id" json:"-,computed"`
	ZoneID         types.String `tfsdk:"zone_id" path:"zone_id,required"`
	Enabled        types.Bool   `tfsdk:"enabled" json:"enabled,required"`
	Endpoint       types.String `tfsdk:"endpoint" json:"endpoint,required"`
	HideOriginalIP types.Bool   `tfsdk:"hide_original_ip" json:"hideOriginalIp,required"`
	MeasurementID  types.String `tfsdk:"measurement_id" json:"measurementId,required"`
	SetUpTag       types.Bool   `tfsdk:"set_up_tag" json:"setUpTag,optional"`
}

func (m GoogleTagGatewayModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m GoogleTagGatewayModel) MarshalJSONForUpdate(state GoogleTagGatewayModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
