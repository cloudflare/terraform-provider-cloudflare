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

type WorkflowsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[WorkflowsResultDataSourceModel] `json:"result,computed"`
}

type WorkflowsDataSourceModel struct {
	AccountID types.String                                                 `tfsdk:"account_id" path:"account_id,required"`
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
	ID          types.String                                                `tfsdk:"id" json:"id,computed"`
	ClassName   types.String                                                `tfsdk:"class_name" json:"class_name,computed"`
	CreatedOn   timetypes.RFC3339                                           `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Instances   customfield.NestedObject[WorkflowsInstancesDataSourceModel] `tfsdk:"instances" json:"instances,computed"`
	ModifiedOn  timetypes.RFC3339                                           `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Name        types.String                                                `tfsdk:"name" json:"name,computed"`
	ScriptName  types.String                                                `tfsdk:"script_name" json:"script_name,computed"`
	TriggeredOn timetypes.RFC3339                                           `tfsdk:"triggered_on" json:"triggered_on,computed" format:"date-time"`
}

type WorkflowsInstancesDataSourceModel struct {
	Complete        types.Float64 `tfsdk:"complete" json:"complete,computed"`
	Errored         types.Float64 `tfsdk:"errored" json:"errored,computed"`
	Paused          types.Float64 `tfsdk:"paused" json:"paused,computed"`
	Queued          types.Float64 `tfsdk:"queued" json:"queued,computed"`
	Running         types.Float64 `tfsdk:"running" json:"running,computed"`
	Terminated      types.Float64 `tfsdk:"terminated" json:"terminated,computed"`
	Waiting         types.Float64 `tfsdk:"waiting" json:"waiting,computed"`
	WaitingForPause types.Float64 `tfsdk:"waiting_for_pause" json:"waitingForPause,computed"`
}
