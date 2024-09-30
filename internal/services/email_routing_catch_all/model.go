// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_catch_all

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailRoutingCatchAllResultEnvelope struct {
	Result EmailRoutingCatchAllModel `json:"result"`
}

type EmailRoutingCatchAllModel struct {
	ID       types.String                          `tfsdk:"id" json:"-,computed"`
	ZoneID   types.String                          `tfsdk:"zone_id" path:"zone_id,required"`
	Actions  *[]*EmailRoutingCatchAllActionsModel  `tfsdk:"actions" json:"actions,required"`
	Matchers *[]*EmailRoutingCatchAllMatchersModel `tfsdk:"matchers" json:"matchers,required"`
	Name     types.String                          `tfsdk:"name" json:"name,optional"`
	Enabled  types.Bool                            `tfsdk:"enabled" json:"enabled,computed_optional"`
	Tag      types.String                          `tfsdk:"tag" json:"tag,computed"`
}

type EmailRoutingCatchAllActionsModel struct {
	Type  types.String    `tfsdk:"type" json:"type,required"`
	Value *[]types.String `tfsdk:"value" json:"value,optional"`
}

type EmailRoutingCatchAllMatchersModel struct {
	Type types.String `tfsdk:"type" json:"type,required"`
}
