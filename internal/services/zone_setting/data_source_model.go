// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_setting

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/zones"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneSettingResultDataSourceEnvelope struct {
Result ZoneSettingDataSourceModel `json:"result,computed"`
}

type ZoneSettingDataSourceModel struct {
SettingID types.String `tfsdk:"setting_id" path:"setting_id,required"`
ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
Editable types.Bool `tfsdk:"editable" json:"editable,computed"`
Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed"`
ID types.String `tfsdk:"id" json:"id,computed"`
ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
TimeRemaining types.Float64 `tfsdk:"time_remaining" json:"time_remaining,computed"`
Value types.String `tfsdk:"value" json:"value,computed"`
}

func (m *ZoneSettingDataSourceModel) toReadParams(_ context.Context) (params zones.SettingGetParams, diags diag.Diagnostics) {
  params = zones.SettingGetParams{
    ZoneID: cloudflare.F(m.ZoneID.ValueString()),
  }

  return
}
