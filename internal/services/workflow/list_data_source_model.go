// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflow

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/workflows"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkflowsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[WorkflowsResultDataSourceModel] `json:"result,computed"`
}

type WorkflowsDataSourceModel struct {
	AccountID types.String                                                 `tfsdk:"account_id" path:"account_id,optional"`
	Search    types.String                                                 `tfsdk:"search" query:"search,optional"`
	MaxItems  types.Int64                                                  `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[WorkflowsResultDataSourceModel] `tfsdk:"result"`
}

func (m *WorkflowsDataSourceModel) toListParams(_ context.Context) (params workflows.WorkflowListParams, diags diag.Diagnostics) {
	params = workflows.WorkflowListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Search.IsNull() {
		params.Search = cloudflare.F(m.Search.ValueString())
	}

	return
}

type WorkflowsResultDataSourceModel struct {
	ID          types.String                                                    `tfsdk:"id" json:"id,computed"`
	ClassName   types.String                                                    `tfsdk:"class_name" json:"class_name,computed"`
	CreatedOn   timetypes.RFC3339                                               `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Instances   customfield.Map[types.Float64]                                  `tfsdk:"instances" json:"instances,computed"`
	ModifiedOn  timetypes.RFC3339                                               `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Name        types.String                                                    `tfsdk:"name" json:"name,computed"`
	ScriptName  types.String                                                    `tfsdk:"script_name" json:"script_name,computed"`
	TriggeredOn timetypes.RFC3339                                               `tfsdk:"triggered_on" json:"triggered_on,computed" format:"date-time"`
	Schedules   customfield.NestedObjectList[WorkflowsSchedulesDataSourceModel] `tfsdk:"schedules" json:"schedules,computed"`
}

type WorkflowsSchedulesDataSourceModel struct {
	Cron         types.String `tfsdk:"cron" json:"cron,computed"`
	NextInstance types.String `tfsdk:"next_instance" json:"next_instance,computed"`
}
