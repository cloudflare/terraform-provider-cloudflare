// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_rule

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/pagerules"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PageRuleResultDataSourceEnvelope struct {
	Result PageRuleDataSourceModel `json:"result,computed"`
}

type PageRuleDataSourceModel struct {
	PageruleID types.String                                                 `tfsdk:"pagerule_id" path:"pagerule_id,required"`
	ZoneID     types.String                                                 `tfsdk:"zone_id" path:"zone_id,required"`
	CreatedOn  timetypes.RFC3339                                            `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ID         types.String                                                 `tfsdk:"id" json:"id,computed"`
	ModifiedOn timetypes.RFC3339                                            `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Priority   types.Int64                                                  `tfsdk:"priority" json:"priority,computed"`
	Status     types.String                                                 `tfsdk:"status" json:"status,computed"`
	Actions    customfield.NestedObjectList[PageRuleActionsDataSourceModel] `tfsdk:"actions" json:"actions,computed"`
	Targets    customfield.NestedObjectList[PageRuleTargetsDataSourceModel] `tfsdk:"targets" json:"targets,computed"`
}

func (m *PageRuleDataSourceModel) toReadParams(_ context.Context) (params pagerules.PageruleGetParams, diags diag.Diagnostics) {
	params = pagerules.PageruleGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type PageRuleActionsDataSourceModel struct {
	ID    types.String `tfsdk:"id" json:"id,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}

type PageRuleTargetsDataSourceModel struct {
	Constraint customfield.NestedObject[PageRuleTargetsConstraintDataSourceModel] `tfsdk:"constraint" json:"constraint,computed"`
	Target     types.String                                                       `tfsdk:"target" json:"target,computed"`
}

type PageRuleTargetsConstraintDataSourceModel struct {
	Operator types.String `tfsdk:"operator" json:"operator,computed"`
	Value    types.String `tfsdk:"value" json:"value,computed"`
}
