// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_cron_trigger

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersCronTriggerResultEnvelope struct {
	Result WorkersCronTriggerModel `json:"result,computed"`
}

type WorkersCronTriggerModel struct {
	ID         types.String                         `tfsdk:"id" json:"-,computed"`
	ScriptName types.String                         `tfsdk:"script_name" path:"script_name"`
	AccountID  types.String                         `tfsdk:"account_id" path:"account_id"`
	Schedules  *[]*WorkersCronTriggerSchedulesModel `tfsdk:"schedules" json:"schedules,computed"`
}

type WorkersCronTriggerSchedulesModel struct {
	CreatedOn  jsontypes.Normalized `tfsdk:"created_on" json:"created_on,computed"`
	Cron       jsontypes.Normalized `tfsdk:"cron" json:"cron,computed"`
	ModifiedOn jsontypes.Normalized `tfsdk:"modified_on" json:"modified_on,computed"`
}
