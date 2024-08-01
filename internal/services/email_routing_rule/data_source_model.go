// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_rule

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailRoutingRuleResultDataSourceEnvelope struct {
	Result EmailRoutingRuleDataSourceModel `json:"result,computed"`
}

type EmailRoutingRuleResultListDataSourceEnvelope struct {
	Result *[]*EmailRoutingRuleDataSourceModel `json:"result,computed"`
}

type EmailRoutingRuleDataSourceModel struct {
	RuleIdentifier types.String                                `tfsdk:"rule_identifier" path:"rule_identifier"`
	ZoneIdentifier types.String                                `tfsdk:"zone_identifier" path:"zone_identifier"`
	Enabled        types.Bool                                  `tfsdk:"enabled" json:"enabled,computed"`
	ID             types.String                                `tfsdk:"id" json:"id,computed"`
	Priority       types.Float64                               `tfsdk:"priority" json:"priority,computed"`
	Tag            types.String                                `tfsdk:"tag" json:"tag,computed"`
	Name           types.String                                `tfsdk:"name" json:"name"`
	Actions        *[]*EmailRoutingRuleActionsDataSourceModel  `tfsdk:"actions" json:"actions"`
	Matchers       *[]*EmailRoutingRuleMatchersDataSourceModel `tfsdk:"matchers" json:"matchers"`
	Filter         *EmailRoutingRuleFindOneByDataSourceModel   `tfsdk:"filter"`
}

type EmailRoutingRuleActionsDataSourceModel struct {
	Type  types.String    `tfsdk:"type" json:"type,computed"`
	Value *[]types.String `tfsdk:"value" json:"value,computed"`
}

type EmailRoutingRuleMatchersDataSourceModel struct {
	Field types.String `tfsdk:"field" json:"field,computed"`
	Type  types.String `tfsdk:"type" json:"type,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}

type EmailRoutingRuleFindOneByDataSourceModel struct {
	ZoneIdentifier types.String  `tfsdk:"zone_identifier" path:"zone_identifier"`
	Enabled        types.Bool    `tfsdk:"enabled" query:"enabled"`
	Page           types.Float64 `tfsdk:"page" query:"page"`
	PerPage        types.Float64 `tfsdk:"per_page" query:"per_page"`
}
