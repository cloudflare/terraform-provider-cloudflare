// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflow

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/workflows"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkflowResultDataSourceEnvelope struct {
	Result WorkflowDataSourceModel `json:"result,computed"`
}

type WorkflowDataSourceModel struct {
	ID           types.String                                               `tfsdk:"id" path:"workflow_name,computed"`
	WorkflowName types.String                                               `tfsdk:"workflow_name" path:"workflow_name,optional"`
	AccountID    types.String                                               `tfsdk:"account_id" path:"account_id,required"`
	ClassName    types.String                                               `tfsdk:"class_name" json:"class_name,computed"`
	CreatedOn    timetypes.RFC3339                                          `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ModifiedOn   timetypes.RFC3339                                          `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Name         types.String                                               `tfsdk:"name" json:"name,computed"`
	ScriptName   types.String                                               `tfsdk:"script_name" json:"script_name,computed"`
	TriggeredOn  timetypes.RFC3339                                          `tfsdk:"triggered_on" json:"triggered_on,computed" format:"date-time"`
	Instances    customfield.NestedObject[WorkflowInstancesDataSourceModel] `tfsdk:"instances" json:"instances,computed"`
	Filter       *WorkflowFindOneByDataSourceModel                          `tfsdk:"filter"`
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

type WorkflowInstancesDataSourceModel struct {
	Complete        types.Float64 `tfsdk:"complete" json:"complete,computed"`
	Errored         types.Float64 `tfsdk:"errored" json:"errored,computed"`
	Paused          types.Float64 `tfsdk:"paused" json:"paused,computed"`
	Queued          types.Float64 `tfsdk:"queued" json:"queued,computed"`
	Running         types.Float64 `tfsdk:"running" json:"running,computed"`
	Terminated      types.Float64 `tfsdk:"terminated" json:"terminated,computed"`
	Waiting         types.Float64 `tfsdk:"waiting" json:"waiting,computed"`
	WaitingForPause types.Float64 `tfsdk:"waiting_for_pause" json:"waitingForPause,computed"`
}

type WorkflowFindOneByDataSourceModel struct {
	Search types.String `tfsdk:"search" query:"search,optional"`
}
