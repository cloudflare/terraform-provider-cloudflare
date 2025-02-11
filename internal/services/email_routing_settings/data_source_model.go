// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_settings

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/email_routing"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailRoutingSettingsResultDataSourceEnvelope struct {
	Result EmailRoutingSettingsDataSourceModel `json:"result,computed"`
}

type EmailRoutingSettingsDataSourceModel struct {
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
}

func (m *EmailRoutingSettingsDataSourceModel) toReadParams(_ context.Context) (params email_routing.EmailRoutingGetParams, diags diag.Diagnostics) {
	params = email_routing.EmailRoutingGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
