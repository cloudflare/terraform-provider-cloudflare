// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_cron_trigger

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersCronTriggerResultDataSourceEnvelope struct {
	Result WorkersCronTriggerDataSourceModel `json:"result,computed"`
}

type WorkersCronTriggerDataSourceModel struct {
	AccountID  types.String                                   `tfsdk:"account_id" path:"account_id"`
	ScriptName types.String                                   `tfsdk:"script_name" path:"script_name"`
	Schedules  *[]*WorkersCronTriggerSchedulesDataSourceModel `tfsdk:"schedules" json:"schedules"`
}

type WorkersCronTriggerSchedulesDataSourceModel struct {
	CreatedOn  jsontypes.Normalized `tfsdk:"created_on" json:"created_on,computed"`
	Cron       jsontypes.Normalized `tfsdk:"cron" json:"cron,computed"`
	ModifiedOn jsontypes.Normalized `tfsdk:"modified_on" json:"modified_on,computed"`
}
