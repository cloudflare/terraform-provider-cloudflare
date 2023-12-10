package email_routing_rule

import "github.com/hashicorp/terraform-plugin-framework/types"

type EmailRoutingRuleModel struct {
	ZoneID   types.String                    `tfsdk:"zone_id"`
	ID       types.String                    `tfsdk:"id"`
	Tag      types.String                    `tfsdk:"tag"`
	Name     types.String                    `tfsdk:"name"`
	Priority types.Int64                     `tfsdk:"priority"`
	Enabled  types.Bool                      `tfsdk:"enabled"`
	Matcher  []*EmailRoutingRuleMatcherModel `tfsdk:"matcher"`
	Action   []*EmailRoutingRuleActionModel  `tfsdk:"action"`
}

type EmailRoutingRuleMatcherModel struct {
	Type  types.String `tfsdk:"type"`
	Field types.String `tfsdk:"field"`
	Value types.String `tfsdk:"value"`
}

type EmailRoutingRuleActionModel struct {
	Type  types.String `tfsdk:"type"`
	Value types.Set    `tfsdk:"value"`
}
