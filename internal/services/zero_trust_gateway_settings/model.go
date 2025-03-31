// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_settings

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustGatewaySettingsResultEnvelope struct {
	Result ZeroTrustGatewaySettingsModel `json:"result"`
}

type ZeroTrustGatewaySettingsModel struct {
	ID        types.String         `tfsdk:"id" json:"-,computed"`
	AccountID types.String         `tfsdk:"account_id" path:"account_id,required"`
	Settings  jsontypes.Normalized `tfsdk:"settings" json:"settings,optional"`
	CreatedAt timetypes.RFC3339    `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UpdatedAt timetypes.RFC3339    `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m ZeroTrustGatewaySettingsModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustGatewaySettingsModel) MarshalJSONForUpdate(state ZeroTrustGatewaySettingsModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
