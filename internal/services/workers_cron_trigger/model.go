// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_cron_trigger

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersCronTriggerResultEnvelope struct {
	Result WorkersCronTriggerModel `json:"result"`
}

type WorkersCronTriggerModel struct {
	ID         types.String                                                   `tfsdk:"id" json:"-,computed"`
	ScriptName types.String                                                   `tfsdk:"script_name" path:"script_name,required"`
	AccountID  types.String                                                   `tfsdk:"account_id" path:"account_id,required"`
	Cron       types.String                                                   `tfsdk:"cron" json:"cron,optional"`
	Schedules  customfield.NestedObjectList[WorkersCronTriggerSchedulesModel] `tfsdk:"schedules" json:"schedules,computed"`
}

type WorkersCronTriggerSchedulesModel struct {
	CreatedOn  types.String `tfsdk:"created_on" json:"created_on,computed"`
	Cron       types.String `tfsdk:"cron" json:"cron,computed"`
	ModifiedOn types.String `tfsdk:"modified_on" json:"modified_on,computed"`
}
