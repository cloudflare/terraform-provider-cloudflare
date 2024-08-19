// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room_rules

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WaitingRoomRulesResultDataSourceEnvelope struct {
	Result WaitingRoomRulesDataSourceModel `json:"result,computed"`
}

type WaitingRoomRulesDataSourceModel struct {
	WaitingRoomID types.String `tfsdk:"waiting_room_id" path:"waiting_room_id"`
	ZoneID        types.String `tfsdk:"zone_id" path:"zone_id"`
}
