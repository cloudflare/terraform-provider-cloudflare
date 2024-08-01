// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room_rules

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WaitingRoomRulesListResultListDataSourceEnvelope struct {
	Result *[]*WaitingRoomRulesListResultDataSourceModel `json:"result,computed"`
}

type WaitingRoomRulesListDataSourceModel struct {
	WaitingRoomID types.String                                  `tfsdk:"waiting_room_id" path:"waiting_room_id"`
	ZoneID        types.String                                  `tfsdk:"zone_id" path:"zone_id"`
	MaxItems      types.Int64                                   `tfsdk:"max_items"`
	Result        *[]*WaitingRoomRulesListResultDataSourceModel `tfsdk:"result"`
}

type WaitingRoomRulesListResultDataSourceModel struct {
	ID          types.String      `tfsdk:"id" json:"id"`
	Action      types.String      `tfsdk:"action" json:"action"`
	Description types.String      `tfsdk:"description" json:"description,computed"`
	Enabled     types.Bool        `tfsdk:"enabled" json:"enabled,computed"`
	Expression  types.String      `tfsdk:"expression" json:"expression"`
	LastUpdated timetypes.RFC3339 `tfsdk:"last_updated" json:"last_updated,computed"`
	Version     types.String      `tfsdk:"version" json:"version"`
}
