// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room_rules

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WaitingRoomRulesResultEnvelope struct {
	Result WaitingRoomRulesModel `json:"result"`
}

type WaitingRoomRulesModel struct {
	WaitingRoomID types.String `tfsdk:"waiting_room_id" path:"waiting_room_id"`
	ZoneID        types.String `tfsdk:"zone_id" path:"zone_id"`
	RuleID        types.String `tfsdk:"rule_id" path:"rule_id"`
	Action        types.String `tfsdk:"action" json:"action"`
	Expression    types.String `tfsdk:"expression" json:"expression"`
	Description   types.String `tfsdk:"description" json:"description,computed_optional"`
	Enabled       types.Bool   `tfsdk:"enabled" json:"enabled,computed_optional"`
}
