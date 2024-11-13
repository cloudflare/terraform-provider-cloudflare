// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room_settings

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/waiting_rooms"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WaitingRoomSettingsResultDataSourceEnvelope struct {
	Result WaitingRoomSettingsDataSourceModel `json:"result,computed"`
}

type WaitingRoomSettingsDataSourceModel struct {
	ZoneID                    types.String `tfsdk:"zone_id" path:"zone_id,required"`
	SearchEngineCrawlerBypass types.Bool   `tfsdk:"search_engine_crawler_bypass" json:"search_engine_crawler_bypass,computed"`
}

func (m *WaitingRoomSettingsDataSourceModel) toReadParams(_ context.Context) (params waiting_rooms.SettingGetParams, diags diag.Diagnostics) {
	params = waiting_rooms.SettingGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
