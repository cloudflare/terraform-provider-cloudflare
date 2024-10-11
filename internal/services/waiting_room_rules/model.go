// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room_rules

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WaitingRoomRulesResultEnvelope struct {
	Result customfield.NestedObjectList[WaitingRoomRulesBodyModel] `json:"result"`
}

type WaitingRoomRulesModel struct {
	WaitingRoomID types.String                                            `tfsdk:"waiting_room_id" path:"waiting_room_id,required"`
	ZoneID        types.String                                            `tfsdk:"zone_id" path:"zone_id,required"`
	RuleID        types.String                                            `tfsdk:"rule_id" path:"rule_id,optional"`
	Action        types.String                                            `tfsdk:"action" json:"action,required"`
	Expression    types.String                                            `tfsdk:"expression" json:"expression,required"`
	Description   types.String                                            `tfsdk:"description" json:"description,computed_optional"`
	Enabled       types.Bool                                              `tfsdk:"enabled" json:"enabled,computed_optional"`
	Body          customfield.NestedObjectList[WaitingRoomRulesBodyModel] `tfsdk:"body" json:"body,computed_optional"`
}

func (m WaitingRoomRulesModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m.Body)
}

func (m WaitingRoomRulesModel) MarshalJSONForUpdate(state WaitingRoomRulesModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m.Body, state.Body)
}

type WaitingRoomRulesBodyModel struct {
	Action      types.String `tfsdk:"action" json:"action,required"`
	Expression  types.String `tfsdk:"expression" json:"expression,required"`
	Description types.String `tfsdk:"description" json:"description,computed_optional"`
	Enabled     types.Bool   `tfsdk:"enabled" json:"enabled,computed_optional"`
}
