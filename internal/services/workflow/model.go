// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflow

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkflowResultEnvelope struct {
	Result WorkflowModel `json:"result"`
}

type WorkflowModel struct {
	ID                types.String                                     `tfsdk:"id" json:"-,computed"`
	Name              types.String                                     `tfsdk:"name" json:"name,computed"`
	AccountID         types.String                                     `tfsdk:"account_id" path:"account_id,required"`
	WorkflowName      types.String                                     `tfsdk:"workflow_name" path:"workflow_name,required"`
	ClassName         types.String                                     `tfsdk:"class_name" json:"class_name,required"`
	ScriptName        types.String                                     `tfsdk:"script_name" json:"script_name,required"`
	CreatedOn         timetypes.RFC3339                                `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	IsDeleted         types.Float64                                    `tfsdk:"is_deleted" json:"is_deleted,computed,no_refresh"`
	ModifiedOn        timetypes.RFC3339                                `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	TerminatorRunning types.Float64                                    `tfsdk:"terminator_running" json:"terminator_running,computed,no_refresh"`
	TriggeredOn       timetypes.RFC3339                                `tfsdk:"triggered_on" json:"triggered_on,computed" format:"date-time"`
	VersionID         types.String                                     `tfsdk:"version_id" json:"version_id,computed,no_refresh"`
	Instances         customfield.NestedObject[WorkflowInstancesModel] `tfsdk:"instances" json:"instances,computed"`
}

func (m WorkflowModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m WorkflowModel) MarshalJSONForUpdate(state WorkflowModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type WorkflowInstancesModel struct {
	Complete        types.Float64 `tfsdk:"complete" json:"complete,computed"`
	Errored         types.Float64 `tfsdk:"errored" json:"errored,computed"`
	Paused          types.Float64 `tfsdk:"paused" json:"paused,computed"`
	Queued          types.Float64 `tfsdk:"queued" json:"queued,computed"`
	Running         types.Float64 `tfsdk:"running" json:"running,computed"`
	Terminated      types.Float64 `tfsdk:"terminated" json:"terminated,computed"`
	Waiting         types.Float64 `tfsdk:"waiting" json:"waiting,computed"`
	WaitingForPause types.Float64 `tfsdk:"waiting_for_pause" json:"waitingForPause,computed"`
}
