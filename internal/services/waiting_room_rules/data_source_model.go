// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room_rules

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WaitingRoomRulesResultListDataSourceEnvelope struct {
	Result *[]*WaitingRoomRulesDataSourceModel `json:"result,computed"`
}

type WaitingRoomRulesDataSourceModel struct {
	ID          types.String                              `tfsdk:"id" json:"id"`
	Action      types.String                              `tfsdk:"action" json:"action"`
	Description types.String                              `tfsdk:"description" json:"description"`
	Enabled     types.Bool                                `tfsdk:"enabled" json:"enabled"`
	Expression  types.String                              `tfsdk:"expression" json:"expression"`
	LastUpdated timetypes.RFC3339                         `tfsdk:"last_updated" json:"last_updated"`
	Version     types.String                              `tfsdk:"version" json:"version"`
	Filter      *WaitingRoomRulesFindOneByDataSourceModel `tfsdk:"filter"`
}

type WaitingRoomRulesFindOneByDataSourceModel struct {
	ZoneID        types.String `tfsdk:"zone_id" path:"zone_id"`
	WaitingRoomID types.String `tfsdk:"waiting_room_id" path:"waiting_room_id"`
}
