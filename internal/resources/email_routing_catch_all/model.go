// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_catch_all

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailRoutingCatchAllResultEnvelope struct {
	Result EmailRoutingCatchAllModel `json:"result,computed"`
}

type EmailRoutingCatchAllModel struct {
	ID             types.String                          `tfsdk:"id" json:"id,computed"`
	ZoneIdentifier types.String                          `tfsdk:"zone_identifier" path:"zone_identifier"`
	Actions        *[]*EmailRoutingCatchAllActionsModel  `tfsdk:"actions" json:"actions"`
	Matchers       *[]*EmailRoutingCatchAllMatchersModel `tfsdk:"matchers" json:"matchers"`
	Enabled        types.Bool                            `tfsdk:"enabled" json:"enabled"`
	Name           types.String                          `tfsdk:"name" json:"name"`
}

type EmailRoutingCatchAllActionsModel struct {
	Type  types.String    `tfsdk:"type" json:"type"`
	Value *[]types.String `tfsdk:"value" json:"value"`
}

type EmailRoutingCatchAllMatchersModel struct {
	Type types.String `tfsdk:"type" json:"type"`
}
