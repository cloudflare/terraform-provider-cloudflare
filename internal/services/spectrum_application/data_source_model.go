// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum_application

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/spectrum"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SpectrumApplicationResultDataSourceEnvelope struct {
	Result SpectrumApplicationDataSourceModel `json:"result,computed"`
}

type SpectrumApplicationDataSourceModel struct {
	AppID  types.String `tfsdk:"app_id" path:"app_id,required"`
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
}

func (m *SpectrumApplicationDataSourceModel) toReadParams(_ context.Context) (params spectrum.AppGetParams, diags diag.Diagnostics) {
	params = spectrum.AppGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
