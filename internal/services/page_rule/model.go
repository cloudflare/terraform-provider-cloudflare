// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_rule

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PageRuleResultEnvelope struct {
	Result PageRuleModel `json:"result,computed"`
}

type PageRuleResultDataSourceEnvelope struct {
	Result PageRuleDataSourceModel `json:"result,computed"`
}

type PageRulesResultDataSourceEnvelope struct {
	Result PageRulesDataSourceModel `json:"result,computed"`
}

type PageRuleModel struct {
	ZoneID     types.String             `tfsdk:"zone_id" path:"zone_id"`
	PageruleID types.String             `tfsdk:"pagerule_id" path:"pagerule_id"`
	Actions    *[]*PageRuleActionsModel `tfsdk:"actions" json:"actions"`
	Targets    *[]*PageRuleTargetsModel `tfsdk:"targets" json:"targets"`
	Priority   types.Int64              `tfsdk:"priority" json:"priority"`
	Status     types.String             `tfsdk:"status" json:"status"`
	ID         types.String             `tfsdk:"id" json:"id,computed"`
}

type PageRuleActionsModel struct {
	ModifiedOn types.String               `tfsdk:"modified_on" json:"modified_on,computed"`
	Name       types.String               `tfsdk:"name" json:"name"`
	Value      *PageRuleActionsValueModel `tfsdk:"value" json:"value"`
}

type PageRuleActionsValueModel struct {
	Type types.String `tfsdk:"type" json:"type"`
	URL  types.String `tfsdk:"url" json:"url"`
}

type PageRuleTargetsModel struct {
	Constraint *PageRuleTargetsConstraintModel `tfsdk:"constraint" json:"constraint"`
	Target     types.String                    `tfsdk:"target" json:"target"`
}

type PageRuleTargetsConstraintModel struct {
	Operator types.String `tfsdk:"operator" json:"operator"`
	Value    types.String `tfsdk:"value" json:"value"`
}

type PageRuleDataSourceModel struct {
}

type PageRulesDataSourceModel struct {
}
