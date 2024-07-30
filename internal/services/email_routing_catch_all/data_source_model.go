// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_catch_all

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailRoutingCatchAllResultDataSourceEnvelope struct {
	Result EmailRoutingCatchAllDataSourceModel `json:"result,computed"`
}

type EmailRoutingCatchAllDataSourceModel struct {
	ZoneIdentifier types.String                                    `tfsdk:"zone_identifier" path:"zone_identifier"`
	ID             types.String                                    `tfsdk:"id" json:"id"`
	Actions        *[]*EmailRoutingCatchAllActionsDataSourceModel  `tfsdk:"actions" json:"actions"`
	Enabled        types.Bool                                      `tfsdk:"enabled" json:"enabled"`
	Matchers       *[]*EmailRoutingCatchAllMatchersDataSourceModel `tfsdk:"matchers" json:"matchers"`
	Name           types.String                                    `tfsdk:"name" json:"name"`
	Tag            types.String                                    `tfsdk:"tag" json:"tag"`
}

type EmailRoutingCatchAllActionsDataSourceModel struct {
	Type  types.String    `tfsdk:"type" json:"type,computed"`
	Value *[]types.String `tfsdk:"value" json:"value"`
}

type EmailRoutingCatchAllMatchersDataSourceModel struct {
	Type types.String `tfsdk:"type" json:"type,computed"`
}
