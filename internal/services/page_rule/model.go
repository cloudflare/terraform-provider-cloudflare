// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_rule

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PageRuleResultEnvelope struct {
	Result PageRuleModel `json:"result"`
}

type PageRuleModel struct {
	ZoneID     types.String             `tfsdk:"zone_id" path:"zone_id,required"`
	PageruleID types.String             `tfsdk:"pagerule_id" path:"pagerule_id,optional"`
	Actions    *[]*PageRuleActionsModel `tfsdk:"actions" json:"actions,required"`
	Targets    *[]*PageRuleTargetsModel `tfsdk:"targets" json:"targets,required"`
	Priority   types.Int64              `tfsdk:"priority" json:"priority,computed_optional"`
	Status     types.String             `tfsdk:"status" json:"status,computed_optional"`
	ID         types.String             `tfsdk:"id" json:"id,computed"`
}

type PageRuleActionsModel struct {
	ModifiedOn timetypes.RFC3339          `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Name       types.String               `tfsdk:"name" json:"name,optional"`
	Value      *PageRuleActionsValueModel `tfsdk:"value" json:"value,optional"`
}

type PageRuleActionsValueModel struct {
	Type types.String `tfsdk:"type" json:"type,optional"`
	URL  types.String `tfsdk:"url" json:"url,optional"`
}

type PageRuleTargetsModel struct {
	Constraint *PageRuleTargetsConstraintModel `tfsdk:"constraint" json:"constraint,required"`
	Target     types.String                    `tfsdk:"target" json:"target,required"`
}

type PageRuleTargetsConstraintModel struct {
	Operator types.String `tfsdk:"operator" json:"operator,computed_optional"`
	Value    types.String `tfsdk:"value" json:"value,required"`
}
