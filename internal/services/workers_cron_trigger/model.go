// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_cron_trigger

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersCronTriggerResultEnvelope struct {
	Result *[]*WorkersCronTriggerSchedulesModel `json:"result"`
}

type WorkersCronTriggerModel struct {
	ID         types.String                         `tfsdk:"id" json:"-,computed"`
	ScriptName types.String                         `tfsdk:"script_name" path:"script_name,required"`
	AccountID  types.String                         `tfsdk:"account_id" path:"account_id,required"`
	Schedules  *[]*WorkersCronTriggerSchedulesModel `tfsdk:"schedules" json:"schedules,required"`
}

func (m WorkersCronTriggerModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m.Schedules)
}

func (m WorkersCronTriggerModel) MarshalJSONForUpdate(state WorkersCronTriggerModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m.Schedules, state.Schedules)
}

type WorkersCronTriggerSchedulesModel struct {
	Cron types.String `tfsdk:"cron" json:"cron,required"`
}

type WorkersCronTriggerBodyModel struct {
	CreatedOn  types.String `tfsdk:"created_on" json:"created_on,computed"`
	Cron       types.String `tfsdk:"cron" json:"cron,optional"`
	ModifiedOn types.String `tfsdk:"modified_on" json:"modified_on,computed"`
}

type WorkersCronTriggerSchedulesModel struct {
	CreatedOn  types.String `tfsdk:"created_on" json:"created_on,computed"`
	Cron       types.String `tfsdk:"cron" json:"cron,computed"`
	ModifiedOn types.String `tfsdk:"modified_on" json:"modified_on,computed"`
}
