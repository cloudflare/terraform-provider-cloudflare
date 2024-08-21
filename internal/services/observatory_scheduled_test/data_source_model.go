// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package observatory_scheduled_test

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/speed"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ObservatoryScheduledTestResultDataSourceEnvelope struct {
	Result ObservatoryScheduledTestDataSourceModel `json:"result,computed"`
}

type ObservatoryScheduledTestDataSourceModel struct {
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id"`
	URL       types.String `tfsdk:"url" path:"url,computed"`
	Region    types.String `tfsdk:"region" query:"region,computed"`
	Frequency types.String `tfsdk:"frequency" json:"frequency"`
}

func (m *ObservatoryScheduledTestDataSourceModel) toReadParams() (params speed.ScheduleGetParams, diags diag.Diagnostics) {
	params = speed.ScheduleGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
