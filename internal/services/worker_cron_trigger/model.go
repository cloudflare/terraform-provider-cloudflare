// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package worker_cron_trigger

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkerCronTriggerResultEnvelope struct {
	Result WorkerCronTriggerModel `json:"result,computed"`
}

type WorkerCronTriggerModel struct {
	ID         types.String                        `tfsdk:"id" json:"-,computed"`
	AccountID  types.String                        `tfsdk:"account_id" path:"account_id"`
	ScriptName types.String                        `tfsdk:"script_name" path:"script_name"`
	Schedules  *[]*WorkerCronTriggerSchedulesModel `tfsdk:"schedules" json:"schedules,computed"`
}

type WorkerCronTriggerSchedulesModel struct {
	CreatedOn  jsontypes.Normalized `tfsdk:"created_on" json:"created_on,computed"`
	Cron       jsontypes.Normalized `tfsdk:"cron" json:"cron,computed"`
	ModifiedOn jsontypes.Normalized `tfsdk:"modified_on" json:"modified_on,computed"`
}
