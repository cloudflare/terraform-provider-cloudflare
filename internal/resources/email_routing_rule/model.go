// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_rule

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailRoutingRuleResultEnvelope struct {
	Result EmailRoutingRuleModel `json:"result,computed"`
}

type EmailRoutingRuleModel struct {
	ID             types.String                      `tfsdk:"id" json:"id,computed"`
	ZoneIdentifier types.String                      `tfsdk:"zone_identifier" path:"zone_identifier"`
	Actions        *[]*EmailRoutingRuleActionsModel  `tfsdk:"actions" json:"actions"`
	Matchers       *[]*EmailRoutingRuleMatchersModel `tfsdk:"matchers" json:"matchers"`
	Enabled        types.Bool                        `tfsdk:"enabled" json:"enabled"`
	Name           types.String                      `tfsdk:"name" json:"name"`
	Priority       types.Float64                     `tfsdk:"priority" json:"priority"`
}

type EmailRoutingRuleActionsModel struct {
	Type  types.String    `tfsdk:"type" json:"type"`
	Value *[]types.String `tfsdk:"value" json:"value"`
}

type EmailRoutingRuleMatchersModel struct {
	Field types.String `tfsdk:"field" json:"field"`
	Type  types.String `tfsdk:"type" json:"type"`
	Value types.String `tfsdk:"value" json:"value"`
}
