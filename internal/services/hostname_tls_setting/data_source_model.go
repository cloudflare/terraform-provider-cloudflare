// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hostname_tls_setting

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/hostnames"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type HostnameTLSSettingResultDataSourceEnvelope struct {
Result HostnameTLSSettingDataSourceModel `json:"result,computed"`
}

type HostnameTLSSettingDataSourceModel struct {
SettingID types.String `tfsdk:"setting_id" path:"setting_id,required"`
ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
Hostname types.String `tfsdk:"hostname" json:"hostname,computed"`
Status types.String `tfsdk:"status" json:"status,computed"`
UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
Value types.Float64 `tfsdk:"value" json:"value,computed"`
}

func (m *HostnameTLSSettingDataSourceModel) toReadParams(_ context.Context) (params hostnames.SettingTLSGetParams, diags diag.Diagnostics) {
  params = hostnames.SettingTLSGetParams{
    ZoneID: cloudflare.F(m.ZoneID.ValueString()),
  }

  return
}
