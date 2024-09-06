// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room_rules

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WaitingRoomRulesResultEnvelope struct {
	Result WaitingRoomRulesModel `json:"result"`
}

type WaitingRoomRulesModel struct {
	WaitingRoomID types.String `tfsdk:"waiting_room_id" path:"waiting_room_id,required"`
	ZoneID        types.String `tfsdk:"zone_id" path:"zone_id,required"`
	RuleID        types.String `tfsdk:"rule_id" path:"rule_id,optional"`
	Action        types.String `tfsdk:"action" json:"action,required"`
	Expression    types.String `tfsdk:"expression" json:"expression,required"`
	Description   types.String `tfsdk:"description" json:"description,computed_optional"`
	Enabled       types.Bool   `tfsdk:"enabled" json:"enabled,computed_optional"`
}
