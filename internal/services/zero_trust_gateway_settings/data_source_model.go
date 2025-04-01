// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_settings

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustGatewaySettingsResultDataSourceEnvelope struct {
	Result ZeroTrustGatewaySettingsDataSourceModel `json:"result,computed"`
}

type ZeroTrustGatewaySettingsDataSourceModel struct {
	AccountID types.String         `tfsdk:"account_id" path:"account_id,required"`
	CreatedAt timetypes.RFC3339    `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UpdatedAt timetypes.RFC3339    `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Settings  jsontypes.Normalized `tfsdk:"settings" json:"settings,computed"`
}

func (m *ZeroTrustGatewaySettingsDataSourceModel) toReadParams(_ context.Context) (params zero_trust.GatewayConfigurationGetParams, diags diag.Diagnostics) {
	params = zero_trust.GatewayConfigurationGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
