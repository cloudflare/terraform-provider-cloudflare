// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room_rules

import (
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type WaitingRoomRulesResultEnvelope struct {
Result *[]*WaitingRoomRulesRulesModel `json:"result"`
}

type WaitingRoomRulesModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
WaitingRoomID types.String `tfsdk:"waiting_room_id" path:"waiting_room_id,required"`
ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
Rules *[]*WaitingRoomRulesRulesModel `tfsdk:"rules" json:"rules,required"`
Action types.String `tfsdk:"action" json:"action,computed"`
Description types.String `tfsdk:"description" json:"description,computed"`
Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed"`
Expression types.String `tfsdk:"expression" json:"expression,computed"`
LastUpdated timetypes.RFC3339 `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
Version types.String `tfsdk:"version" json:"version,computed"`
}

func (m WaitingRoomRulesModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m.Rules)
}

func (m WaitingRoomRulesModel) MarshalJSONForUpdate(state WaitingRoomRulesModel) (data []byte, err error) {
  return apijson.MarshalForUpdate(m.Rules, state.Rules)
}

type WaitingRoomRulesRulesModel struct {
Action types.String `tfsdk:"action" json:"action,required"`
Expression types.String `tfsdk:"expression" json:"expression,required"`
Description types.String `tfsdk:"description" json:"description,computed_optional"`
Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed_optional"`
}
