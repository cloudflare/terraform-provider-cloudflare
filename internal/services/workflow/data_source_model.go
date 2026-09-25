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

type WorkflowResultDataSourceEnvelope struct {
	Result WorkflowDataSourceModel `json:"result,computed"`
}

type WorkflowDataSourceModel struct {
	ID           types.String                                                   `tfsdk:"id" path:"workflow_name,computed"`
	WorkflowName types.String                                                   `tfsdk:"workflow_name" path:"workflow_name,optional"`
	AccountID    types.String                                                   `tfsdk:"account_id" path:"account_id,optional"`
	ClassName    types.String                                                   `tfsdk:"class_name" json:"class_name,computed"`
	CreatedOn    timetypes.RFC3339                                              `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ModifiedOn   timetypes.RFC3339                                              `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Name         types.String                                                   `tfsdk:"name" json:"name,computed"`
	ScriptName   types.String                                                   `tfsdk:"script_name" json:"script_name,computed"`
	TriggeredOn  timetypes.RFC3339                                              `tfsdk:"triggered_on" json:"triggered_on,computed" format:"date-time"`
	Instances    customfield.Map[types.Float64]                                 `tfsdk:"instances" json:"instances,computed"`
	Schedules    customfield.NestedObjectList[WorkflowSchedulesDataSourceModel] `tfsdk:"schedules" json:"schedules,computed"`
	Filter       *WorkflowFindOneByDataSourceModel                              `tfsdk:"filter"`
}

func (m *WorkflowDataSourceModel) toReadParams(_ context.Context) (params workflows.WorkflowGetParams, diags diag.Diagnostics) {
	params = workflows.WorkflowGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *WorkflowDataSourceModel) toListParams(_ context.Context) (params workflows.WorkflowListParams, diags diag.Diagnostics) {
	params = workflows.WorkflowListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Filter.Search.IsNull() {
		params.Search = cloudflare.F(m.Filter.Search.ValueString())
	}

	return
}

type WorkflowSchedulesDataSourceModel struct {
	Cron         types.String `tfsdk:"cron" json:"cron,computed"`
	NextInstance types.String `tfsdk:"next_instance" json:"next_instance,computed"`
}

type WorkflowFindOneByDataSourceModel struct {
	Search types.String `tfsdk:"search" query:"search,optional"`
}
