// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room_setting

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WaitingRoomSettingResultEnvelope struct {
	Result WaitingRoomSettingModel `json:"result,computed"`
}

type WaitingRoomSettingModel struct {
	ZoneID                    types.String `tfsdk:"zone_id" path:"zone_id"`
	SearchEngineCrawlerBypass types.Bool   `tfsdk:"search_engine_crawler_bypass" json:"search_engine_crawler_bypass"`
}
