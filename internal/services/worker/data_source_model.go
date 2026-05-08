// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package worker

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/workers"
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
	DeployedOn    timetypes.RFC3339                                               `tfsdk:"deployed_on" json:"deployed_on,computed" format:"date-time"`
	Logpush       types.Bool                                                      `tfsdk:"logpush" json:"logpush,computed"`
	Name          types.String                                                    `tfsdk:"name" json:"name,computed"`
	UpdatedOn     timetypes.RFC3339                                               `tfsdk:"updated_on" json:"updated_on,computed" format:"date-time"`
	Tags          customfield.Set[types.String]                                   `tfsdk:"tags" json:"tags,computed"`
	Observability customfield.NestedObject[WorkerObservabilityDataSourceModel]    `tfsdk:"observability" json:"observability,computed"`
	References    customfield.NestedObject[WorkerReferencesDataSourceModel]       `tfsdk:"references" json:"references,computed"`
	Subdomain     customfield.NestedObject[WorkerSubdomainDataSourceModel]        `tfsdk:"subdomain" json:"subdomain,computed"`
	TailConsumers customfield.NestedObjectSet[WorkerTailConsumersDataSourceModel] `tfsdk:"tail_consumers" json:"tail_consumers,computed"`
	Filter        *WorkerFindOneByDataSourceModel                                 `tfsdk:"filter"`
}

func (m *WorkerDataSourceModel) toReadParams(_ context.Context) (params workers.BetaWorkerGetParams, diags diag.Diagnostics) {
	params = workers.BetaWorkerGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *WorkerDataSourceModel) toListParams(_ context.Context) (params workers.BetaWorkerListParams, diags diag.Diagnostics) {
	params = workers.BetaWorkerListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Filter.Order.IsNull() {
		params.Order = cloudflare.F(workers.BetaWorkerListParamsOrder(m.Filter.Order.ValueString()))
	}
	if !m.Filter.OrderBy.IsNull() {
		params.OrderBy = cloudflare.F(workers.BetaWorkerListParamsOrderBy(m.Filter.OrderBy.ValueString()))
	}

	return
}

type WorkerObservabilityDataSourceModel struct {
	Enabled          types.Bool                                                         `tfsdk:"enabled" json:"enabled,computed"`
	HeadSamplingRate types.Float64                                                      `tfsdk:"head_sampling_rate" json:"head_sampling_rate,computed"`
	Logs             customfield.NestedObject[WorkerObservabilityLogsDataSourceModel]   `tfsdk:"logs" json:"logs,computed"`
	Traces           customfield.NestedObject[WorkerObservabilityTracesDataSourceModel] `tfsdk:"traces" json:"traces,computed"`
}

type WorkerObservabilityLogsDataSourceModel struct {
	Destinations     customfield.List[types.String] `tfsdk:"destinations" json:"destinations,computed"`
	Enabled          types.Bool                     `tfsdk:"enabled" json:"enabled,computed"`
	HeadSamplingRate types.Float64                  `tfsdk:"head_sampling_rate" json:"head_sampling_rate,computed"`
	InvocationLogs   types.Bool                     `tfsdk:"invocation_logs" json:"invocation_logs,computed"`
	Persist          types.Bool                     `tfsdk:"persist" json:"persist,computed"`
}

type WorkerObservabilityTracesDataSourceModel struct {
	Destinations     customfield.List[types.String] `tfsdk:"destinations" json:"destinations,computed"`
	Enabled          types.Bool                     `tfsdk:"enabled" json:"enabled,computed"`
	HeadSamplingRate types.Float64                  `tfsdk:"head_sampling_rate" json:"head_sampling_rate,computed"`
	Persist          types.Bool                     `tfsdk:"persist" json:"persist,computed"`
}

type WorkerReferencesDataSourceModel struct {
	DispatchNamespaceOutbounds customfield.NestedObjectList[WorkerReferencesDispatchNamespaceOutboundsDataSourceModel] `tfsdk:"dispatch_namespace_outbounds" json:"dispatch_namespace_outbounds,computed"`
	Domains                    customfield.NestedObjectList[WorkerReferencesDomainsDataSourceModel]                    `tfsdk:"domains" json:"domains,computed"`
	DurableObjects             customfield.NestedObjectList[WorkerReferencesDurableObjectsDataSourceModel]             `tfsdk:"durable_objects" json:"durable_objects,computed"`
	Queues                     customfield.NestedObjectList[WorkerReferencesQueuesDataSourceModel]                     `tfsdk:"queues" json:"queues,computed"`
	Workers                    customfield.NestedObjectList[WorkerReferencesWorkersDataSourceModel]                    `tfsdk:"workers" json:"workers,computed"`
}

type WorkerReferencesDispatchNamespaceOutboundsDataSourceModel struct {
	NamespaceID   types.String `tfsdk:"namespace_id" json:"namespace_id,computed"`
	NamespaceName types.String `tfsdk:"namespace_name" json:"namespace_name,computed"`
	WorkerID      types.String `tfsdk:"worker_id" json:"worker_id,computed"`
	WorkerName    types.String `tfsdk:"worker_name" json:"worker_name,computed"`
}

type WorkerReferencesDomainsDataSourceModel struct {
	ID            types.String `tfsdk:"id" json:"id,computed"`
	CertificateID types.String `tfsdk:"certificate_id" json:"certificate_id,computed"`
	Hostname      types.String `tfsdk:"hostname" json:"hostname,computed"`
	ZoneID        types.String `tfsdk:"zone_id" json:"zone_id,computed"`
	ZoneName      types.String `tfsdk:"zone_name" json:"zone_name,computed"`
}

type WorkerReferencesDurableObjectsDataSourceModel struct {
	NamespaceID   types.String `tfsdk:"namespace_id" json:"namespace_id,computed"`
	NamespaceName types.String `tfsdk:"namespace_name" json:"namespace_name,computed"`
	WorkerID      types.String `tfsdk:"worker_id" json:"worker_id,computed"`
	WorkerName    types.String `tfsdk:"worker_name" json:"worker_name,computed"`
}

type WorkerReferencesQueuesDataSourceModel struct {
	QueueConsumerID types.String `tfsdk:"queue_consumer_id" json:"queue_consumer_id,computed"`
	QueueID         types.String `tfsdk:"queue_id" json:"queue_id,computed"`
	QueueName       types.String `tfsdk:"queue_name" json:"queue_name,computed"`
}

type WorkerReferencesWorkersDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type WorkerSubdomainDataSourceModel struct {
	Enabled         types.Bool `tfsdk:"enabled" json:"enabled,computed"`
	PreviewsEnabled types.Bool `tfsdk:"previews_enabled" json:"previews_enabled,computed"`
}

type WorkerTailConsumersDataSourceModel struct {
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type WorkerFindOneByDataSourceModel struct {
	Order   types.String `tfsdk:"order" query:"order,computed_optional"`
	OrderBy types.String `tfsdk:"order_by" query:"order_by,computed_optional"`
}
