// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package worker_cron_trigger

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkerCronTriggerResultEnvelope struct {
	Result WorkerCronTriggerModel `json:"result,computed"`
}

type WorkerCronTriggerModel struct {
	AccountID  types.String                        `tfsdk:"account_id" path:"account_id"`
	ScriptName types.String                        `tfsdk:"script_name" path:"script_name"`
	Schedules  *[]*WorkerCronTriggerSchedulesModel `tfsdk:"schedules" json:"schedules,computed"`
}

type WorkerCronTriggerSchedulesModel struct {
	CreatedOn  types.String `tfsdk:"created_on" json:"created_on,computed"`
	Cron       types.String `tfsdk:"cron" json:"cron,computed"`
	ModifiedOn types.String `tfsdk:"modified_on" json:"modified_on,computed"`
}
