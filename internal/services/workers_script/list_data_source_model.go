// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_script

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/workers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersScriptsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[WorkersScriptsResultDataSourceModel] `json:"result,computed"`
}

type WorkersScriptsDataSourceModel struct {
	AccountID types.String                                                      `tfsdk:"account_id" path:"account_id,required"`
	Tags      types.String                                                      `tfsdk:"tags" query:"tags,optional"`
	MaxItems  types.Int64                                                       `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[WorkersScriptsResultDataSourceModel] `tfsdk:"result"`
}

func (m *WorkersScriptsDataSourceModel) toListParams(_ context.Context) (params workers.ScriptListParams, diags diag.Diagnostics) {
	params = workers.ScriptListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Tags.IsNull() {
		params.Tags = cloudflare.F(m.Tags.ValueString())
	}

	return
}

type WorkersScriptsResultDataSourceModel struct {
	ID                 types.String                                                             `tfsdk:"id" json:"id,computed"`
	CompatibilityDate  types.String                                                             `tfsdk:"compatibility_date" json:"compatibility_date,computed"`
	CompatibilityFlags customfield.Set[types.String]                                            `tfsdk:"compatibility_flags" json:"compatibility_flags,computed"`
	CreatedOn          timetypes.RFC3339                                                        `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Etag               types.String                                                             `tfsdk:"etag" json:"etag,computed"`
	Handlers           customfield.List[types.String]                                           `tfsdk:"handlers" json:"handlers,computed"`
	HasAssets          types.Bool                                                               `tfsdk:"has_assets" json:"has_assets,computed"`
	HasModules         types.Bool                                                               `tfsdk:"has_modules" json:"has_modules,computed"`
	LastDeployedFrom   types.String                                                             `tfsdk:"last_deployed_from" json:"last_deployed_from,computed"`
	Logpush            types.Bool                                                               `tfsdk:"logpush" json:"logpush,computed"`
	MigrationTag       types.String                                                             `tfsdk:"migration_tag" json:"migration_tag,computed"`
	ModifiedOn         timetypes.RFC3339                                                        `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	NamedHandlers      customfield.NestedObjectList[WorkersScriptsNamedHandlersDataSourceModel] `tfsdk:"named_handlers" json:"named_handlers,computed"`
	Observability      customfield.NestedObject[WorkersScriptsObservabilityDataSourceModel]     `tfsdk:"observability" json:"observability,computed"`
	Placement          customfield.NestedObject[WorkersScriptsPlacementDataSourceModel]         `tfsdk:"placement" json:"placement,computed"`
	PlacementMode      types.String                                                             `tfsdk:"placement_mode" json:"placement_mode,computed"`
	PlacementStatus    types.String                                                             `tfsdk:"placement_status" json:"placement_status,computed"`
	Routes             customfield.NestedObjectList[WorkersScriptsRoutesDataSourceModel]        `tfsdk:"routes" json:"routes,computed"`
	Tag                types.String                                                             `tfsdk:"tag" json:"tag,computed"`
	Tags               customfield.Set[types.String]                                            `tfsdk:"tags" json:"tags,computed"`
	TailConsumers      customfield.NestedObjectSet[WorkersScriptsTailConsumersDataSourceModel]  `tfsdk:"tail_consumers" json:"tail_consumers,computed"`
	UsageModel         types.String                                                             `tfsdk:"usage_model" json:"usage_model,computed"`
}

type WorkersScriptsNamedHandlersDataSourceModel struct {
	Handlers customfield.List[types.String] `tfsdk:"handlers" json:"handlers,computed"`
	Name     types.String                   `tfsdk:"name" json:"name,computed"`
}

type WorkersScriptsObservabilityDataSourceModel struct {
	Enabled          types.Bool                                                               `tfsdk:"enabled" json:"enabled,computed"`
	HeadSamplingRate types.Float64                                                            `tfsdk:"head_sampling_rate" json:"head_sampling_rate,computed"`
	Logs             customfield.NestedObject[WorkersScriptsObservabilityLogsDataSourceModel] `tfsdk:"logs" json:"logs,computed"`
}

type WorkersScriptsObservabilityLogsDataSourceModel struct {
	Enabled          types.Bool                     `tfsdk:"enabled" json:"enabled,computed"`
	InvocationLogs   types.Bool                     `tfsdk:"invocation_logs" json:"invocation_logs,computed"`
	Destinations     customfield.List[types.String] `tfsdk:"destinations" json:"destinations,computed"`
	HeadSamplingRate types.Float64                  `tfsdk:"head_sampling_rate" json:"head_sampling_rate,computed"`
	Persist          types.Bool                     `tfsdk:"persist" json:"persist,computed"`
}

type WorkersScriptsPlacementDataSourceModel struct {
	Mode           types.String      `tfsdk:"mode" json:"mode,computed"`
	LastAnalyzedAt timetypes.RFC3339 `tfsdk:"last_analyzed_at" json:"last_analyzed_at,computed" format:"date-time"`
	Status         types.String      `tfsdk:"status" json:"status,computed"`
	Region         types.String      `tfsdk:"region" json:"region,computed"`
	Hostname       types.String      `tfsdk:"hostname" json:"hostname,computed"`
	Host           types.String      `tfsdk:"host" json:"host,computed"`
}

type WorkersScriptsRoutesDataSourceModel struct {
	ID      types.String `tfsdk:"id" json:"id,computed"`
	Pattern types.String `tfsdk:"pattern" json:"pattern,computed"`
	Script  types.String `tfsdk:"script" json:"script,computed"`
}

type WorkersScriptsTailConsumersDataSourceModel struct {
	Service     types.String `tfsdk:"service" json:"service,computed"`
	Environment types.String `tfsdk:"environment" json:"environment,computed"`
	Namespace   types.String `tfsdk:"namespace" json:"namespace,computed"`
}
