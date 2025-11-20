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

type WorkersResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[WorkersResultDataSourceModel] `json:"result,computed"`
}

type WorkersDataSourceModel struct {
	AccountID types.String                                               `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[WorkersResultDataSourceModel] `tfsdk:"result"`
}

func (m *WorkersDataSourceModel) toListParams(_ context.Context) (params workers.BetaWorkerListParams, diags diag.Diagnostics) {
	params = workers.BetaWorkerListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type WorkersResultDataSourceModel struct {
	ID            types.String                                                     `tfsdk:"id" json:"id,computed"`
	CreatedOn     timetypes.RFC3339                                                `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Logpush       types.Bool                                                       `tfsdk:"logpush" json:"logpush,computed"`
	Name          types.String                                                     `tfsdk:"name" json:"name,computed"`
	Observability customfield.NestedObject[WorkersObservabilityDataSourceModel]    `tfsdk:"observability" json:"observability,computed"`
	References    customfield.NestedObject[WorkersReferencesDataSourceModel]       `tfsdk:"references" json:"references,computed"`
	Subdomain     customfield.NestedObject[WorkersSubdomainDataSourceModel]        `tfsdk:"subdomain" json:"subdomain,computed"`
	Tags          customfield.Set[types.String]                                    `tfsdk:"tags" json:"tags,computed"`
	TailConsumers customfield.NestedObjectSet[WorkersTailConsumersDataSourceModel] `tfsdk:"tail_consumers" json:"tail_consumers,computed"`
	UpdatedOn     timetypes.RFC3339                                                `tfsdk:"updated_on" json:"updated_on,computed" format:"date-time"`
}

type WorkersObservabilityDataSourceModel struct {
	Enabled          types.Bool                                                        `tfsdk:"enabled" json:"enabled,computed"`
	HeadSamplingRate types.Float64                                                     `tfsdk:"head_sampling_rate" json:"head_sampling_rate,computed"`
	Logs             customfield.NestedObject[WorkersObservabilityLogsDataSourceModel] `tfsdk:"logs" json:"logs,computed"`
}

type WorkersObservabilityLogsDataSourceModel struct {
	Enabled          types.Bool    `tfsdk:"enabled" json:"enabled,computed"`
	HeadSamplingRate types.Float64 `tfsdk:"head_sampling_rate" json:"head_sampling_rate,computed"`
	InvocationLogs   types.Bool    `tfsdk:"invocation_logs" json:"invocation_logs,computed"`
}

type WorkersReferencesDataSourceModel struct {
	DispatchNamespaceOutbounds customfield.NestedObjectList[WorkersReferencesDispatchNamespaceOutboundsDataSourceModel] `tfsdk:"dispatch_namespace_outbounds" json:"dispatch_namespace_outbounds,computed"`
	Domains                    customfield.NestedObjectList[WorkersReferencesDomainsDataSourceModel]                    `tfsdk:"domains" json:"domains,computed"`
	DurableObjects             customfield.NestedObjectList[WorkersReferencesDurableObjectsDataSourceModel]             `tfsdk:"durable_objects" json:"durable_objects,computed"`
	Queues                     customfield.NestedObjectList[WorkersReferencesQueuesDataSourceModel]                     `tfsdk:"queues" json:"queues,computed"`
	Workers                    customfield.NestedObjectList[WorkersReferencesWorkersDataSourceModel]                    `tfsdk:"workers" json:"workers,computed"`
}

type WorkersReferencesDispatchNamespaceOutboundsDataSourceModel struct {
	NamespaceID   types.String `tfsdk:"namespace_id" json:"namespace_id,computed"`
	NamespaceName types.String `tfsdk:"namespace_name" json:"namespace_name,computed"`
	WorkerID      types.String `tfsdk:"worker_id" json:"worker_id,computed"`
	WorkerName    types.String `tfsdk:"worker_name" json:"worker_name,computed"`
}

type WorkersReferencesDomainsDataSourceModel struct {
	ID            types.String `tfsdk:"id" json:"id,computed"`
	CertificateID types.String `tfsdk:"certificate_id" json:"certificate_id,computed"`
	Hostname      types.String `tfsdk:"hostname" json:"hostname,computed"`
	ZoneID        types.String `tfsdk:"zone_id" json:"zone_id,computed"`
	ZoneName      types.String `tfsdk:"zone_name" json:"zone_name,computed"`
}

type WorkersReferencesDurableObjectsDataSourceModel struct {
	NamespaceID   types.String `tfsdk:"namespace_id" json:"namespace_id,computed"`
	NamespaceName types.String `tfsdk:"namespace_name" json:"namespace_name,computed"`
	WorkerID      types.String `tfsdk:"worker_id" json:"worker_id,computed"`
	WorkerName    types.String `tfsdk:"worker_name" json:"worker_name,computed"`
}

type WorkersReferencesQueuesDataSourceModel struct {
	QueueConsumerID types.String `tfsdk:"queue_consumer_id" json:"queue_consumer_id,computed"`
	QueueID         types.String `tfsdk:"queue_id" json:"queue_id,computed"`
	QueueName       types.String `tfsdk:"queue_name" json:"queue_name,computed"`
}

type WorkersReferencesWorkersDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type WorkersSubdomainDataSourceModel struct {
	Enabled         types.Bool `tfsdk:"enabled" json:"enabled,computed"`
	PreviewsEnabled types.Bool `tfsdk:"previews_enabled" json:"previews_enabled,computed"`
}

type WorkersTailConsumersDataSourceModel struct {
	Name types.String `tfsdk:"name" json:"name,computed"`
}
