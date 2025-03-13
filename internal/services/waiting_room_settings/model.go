// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room_settings

import (
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type WaitingRoomSettingsResultEnvelope struct {
Result WaitingRoomSettingsModel `json:"result"`
}

type WaitingRoomSettingsModel struct {
ID types.String `tfsdk:"id" json:"-,computed"`
ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
SearchEngineCrawlerBypass types.Bool `tfsdk:"search_engine_crawler_bypass" json:"search_engine_crawler_bypass,computed_optional"`
}

func (m WaitingRoomSettingsModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m WaitingRoomSettingsModel) MarshalJSONForUpdate(state WaitingRoomSettingsModel) (data []byte, err error) {
  return apijson.MarshalForUpdate(m, state)
}
