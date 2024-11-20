// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_rule

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PageRuleResultEnvelope struct {
	Result PageRuleModel `json:"result"`
}

type PageRuleModel struct {
	ID         types.String             `tfsdk:"id" json:"id,computed"`
	ZoneID     types.String             `tfsdk:"zone_id" path:"zone_id,required"`
	Actions    *[]*PageRuleActionsModel `tfsdk:"actions" json:"actions,required"`
	Targets    *[]*PageRuleTargetsModel `tfsdk:"targets" json:"targets,required"`
	Priority   types.Int64              `tfsdk:"priority" json:"priority,computed_optional"`
	Status     types.String             `tfsdk:"status" json:"status,computed_optional"`
	CreatedOn  timetypes.RFC3339        `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ModifiedOn timetypes.RFC3339        `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
}

func (m PageRuleModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m PageRuleModel) MarshalJSONForUpdate(state PageRuleModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type PageRuleActionsModel struct {
	ID    types.String `tfsdk:"id" json:"id,optional"`
	Value types.Int64  `tfsdk:"value" json:"value,optional"`
}

type PageRuleTargetsModel struct {
	Constraint *PageRuleTargetsConstraintModel `tfsdk:"constraint" json:"constraint,required"`
	Target     types.String                    `tfsdk:"target" json:"target,required"`
}

type PageRuleTargetsConstraintModel struct {
	Operator types.String `tfsdk:"operator" json:"operator,computed_optional"`
	Value    types.String `tfsdk:"value" json:"value,required"`
}
