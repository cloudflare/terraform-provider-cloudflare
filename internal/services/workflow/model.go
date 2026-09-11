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
	ID                types.String                   `tfsdk:"id" json:"-,computed"`
	Name              types.String                   `tfsdk:"name" json:"name,computed"`
	AccountID         types.String                   `tfsdk:"account_id" path:"account_id,required"`
	WorkflowName      types.String                   `tfsdk:"workflow_name" path:"workflow_name,required"`
	ClassName         types.String                   `tfsdk:"class_name" json:"class_name,required"`
	ScriptName        types.String                   `tfsdk:"script_name" json:"script_name,required"`
	Concurrency       *WorkflowConcurrencyModel      `tfsdk:"concurrency" json:"concurrency,optional,no_refresh"`
	DefaultRetention  *WorkflowDefaultRetentionModel `tfsdk:"default_retention" json:"default_retention,optional,no_refresh"`
	Limits            *WorkflowLimitsModel           `tfsdk:"limits" json:"limits,optional,no_refresh"`
	Schedules         *[]*WorkflowSchedulesModel     `tfsdk:"schedules" json:"schedules,optional"`
	CreatedOn         timetypes.RFC3339              `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	IsDeleted         types.Float64                  `tfsdk:"is_deleted" json:"is_deleted,computed,no_refresh"`
	ModifiedOn        timetypes.RFC3339              `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	TerminatorRunning types.Float64                  `tfsdk:"terminator_running" json:"terminator_running,computed,no_refresh"`
	TriggeredOn       timetypes.RFC3339              `tfsdk:"triggered_on" json:"triggered_on,computed" format:"date-time"`
	VersionID         types.String                   `tfsdk:"version_id" json:"version_id,computed,no_refresh"`
	Instances         customfield.Map[types.Float64] `tfsdk:"instances" json:"instances,computed"`
}

func (m WorkflowModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m WorkflowModel) MarshalJSONForUpdate(state WorkflowModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type WorkflowConcurrencyModel struct {
	Limit types.Int64 `tfsdk:"limit" json:"limit,optional"`
}

type WorkflowDefaultRetentionModel struct {
	ErrorRetention   customfield.NormalizedDynamicValue `tfsdk:"error_retention" json:"error_retention,optional"`
	SuccessRetention customfield.NormalizedDynamicValue `tfsdk:"success_retention" json:"success_retention,optional"`
}

type WorkflowLimitsModel struct {
	Steps types.Int64 `tfsdk:"steps" json:"steps,optional"`
}

type WorkflowSchedulesModel struct {
	Cron types.String `tfsdk:"cron" json:"cron,required"`
}
