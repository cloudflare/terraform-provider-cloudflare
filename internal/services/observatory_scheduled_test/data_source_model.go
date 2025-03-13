// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package observatory_scheduled_test

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/speed"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type ObservatoryScheduledTestResultDataSourceEnvelope struct {
Result ObservatoryScheduledTestDataSourceModel `json:"result,computed"`
}

type ObservatoryScheduledTestDataSourceModel struct {
URL types.String `tfsdk:"url" path:"url,required"`
ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
Region types.String `tfsdk:"region" query:"region,computed_optional"`
Frequency types.String `tfsdk:"frequency" json:"frequency,computed"`
}

func (m *ObservatoryScheduledTestDataSourceModel) toReadParams(_ context.Context) (params speed.ScheduleGetParams, diags diag.Diagnostics) {
  params = speed.ScheduleGetParams{
    ZoneID: cloudflare.F(m.ZoneID.ValueString()),
  }

  if !m.Region.IsNull() {
    params.Region = cloudflare.F(speed.ScheduleGetParamsRegion(m.Region.ValueString()))
  }

  return
}
