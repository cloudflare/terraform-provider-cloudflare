// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_setting

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/zones"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneSettingResultDataSourceEnvelope struct {
	Result ZoneSettingDataSourceModel `json:"result,computed"`
}

type ZoneSettingDataSourceModel struct {
	SettingID types.String `tfsdk:"setting_id" path:"setting_id,required"`
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id,required"`
}

func (m *ZoneSettingDataSourceModel) toReadParams(_ context.Context) (params zones.SettingGetParams, diags diag.Diagnostics) {
	params = zones.SettingGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
