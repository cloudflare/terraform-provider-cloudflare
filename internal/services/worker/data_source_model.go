// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package worker

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/workers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkerResultDataSourceEnvelope struct {
	Result WorkerDataSourceModel `json:"result,computed"`
}

type WorkerDataSourceModel struct {
	ID            types.String                                                    `tfsdk:"id" path:"worker_id,computed"`
	WorkerID      types.String                                                    `tfsdk:"worker_id" path:"worker_id,optional"`
	AccountID     types.String                                                    `tfsdk:"account_id" path:"account_id,required"`
	CreatedOn     timetypes.RFC3339                                               `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Logpush       types.Bool                                                      `tfsdk:"logpush" json:"logpush,computed"`
	Name          types.String                                                    `tfsdk:"name" json:"name,computed"`
	UpdatedOn     timetypes.RFC3339                                               `tfsdk:"updated_on" json:"updated_on,computed" format:"date-time"`
	Tags          customfield.Set[types.String]                                   `tfsdk:"tags" json:"tags,computed"`
	Observability customfield.NestedObject[WorkerObservabilityDataSourceModel]    `tfsdk:"observability" json:"observability,computed"`
	Subdomain     customfield.NestedObject[WorkerSubdomainDataSourceModel]        `tfsdk:"subdomain" json:"subdomain,computed"`
	TailConsumers customfield.NestedObjectSet[WorkerTailConsumersDataSourceModel] `tfsdk:"tail_consumers" json:"tail_consumers,computed"`
}

func (m *WorkerDataSourceModel) toReadParams(_ context.Context) (params workers.BetaWorkerGetParams, diags diag.Diagnostics) {
	params = workers.BetaWorkerGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type WorkerObservabilityDataSourceModel struct {
	Enabled          types.Bool                                                       `tfsdk:"enabled" json:"enabled,computed"`
	HeadSamplingRate types.Float64                                                    `tfsdk:"head_sampling_rate" json:"head_sampling_rate,computed"`
	Logs             customfield.NestedObject[WorkerObservabilityLogsDataSourceModel] `tfsdk:"logs" json:"logs,computed"`
}

type WorkerObservabilityLogsDataSourceModel struct {
	Enabled          types.Bool    `tfsdk:"enabled" json:"enabled,computed"`
	HeadSamplingRate types.Float64 `tfsdk:"head_sampling_rate" json:"head_sampling_rate,computed"`
	InvocationLogs   types.Bool    `tfsdk:"invocation_logs" json:"invocation_logs,computed"`
}

type WorkerSubdomainDataSourceModel struct {
	Enabled         types.Bool `tfsdk:"enabled" json:"enabled,computed"`
	PreviewsEnabled types.Bool `tfsdk:"previews_enabled" json:"previews_enabled,computed"`
}

type WorkerTailConsumersDataSourceModel struct {
	Name types.String `tfsdk:"name" json:"name,computed"`
}
