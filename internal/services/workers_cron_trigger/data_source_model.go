// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_cron_trigger

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/workers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersCronTriggerResultDataSourceEnvelope struct {
	Result WorkersCronTriggerDataSourceModel `json:"result,computed"`
}

type WorkersCronTriggerDataSourceModel struct {
	AccountID  types.String                           `tfsdk:"account_id" path:"account_id,required"`
	ScriptName types.String                           `tfsdk:"script_name" path:"script_name,required"`
	Schedules  customfield.List[jsontypes.Normalized] `tfsdk:"schedules" json:"schedules,computed"`
}

func (m *WorkersCronTriggerDataSourceModel) toReadParams(_ context.Context) (params workers.ScriptScheduleGetParams, diags diag.Diagnostics) {
	params = workers.ScriptScheduleGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
