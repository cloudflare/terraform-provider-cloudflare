// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package worker

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkerResultEnvelope struct {
	Result WorkerModel `json:"result"`
}

type WorkerModel struct {
	ID            types.String                                          `tfsdk:"id" json:"id,computed"`
	AccountID     types.String                                          `tfsdk:"account_id" path:"account_id,required"`
	Name          types.String                                          `tfsdk:"name" json:"name,required"`
	Logpush       types.Bool                                            `tfsdk:"logpush" json:"logpush,computed_optional"`
	Tags          customfield.Set[types.String]                         `tfsdk:"tags" json:"tags,computed_optional"`
	Observability customfield.NestedObject[WorkerObservabilityModel]    `tfsdk:"observability" json:"observability,computed_optional"`
	Subdomain     customfield.NestedObject[WorkerSubdomainModel]        `tfsdk:"subdomain" json:"subdomain,computed_optional"`
	TailConsumers customfield.NestedObjectSet[WorkerTailConsumersModel] `tfsdk:"tail_consumers" json:"tail_consumers,computed_optional"`
	CreatedOn     timetypes.RFC3339                                     `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	UpdatedOn     timetypes.RFC3339                                     `tfsdk:"updated_on" json:"updated_on,computed" format:"date-time"`
	References    customfield.NestedObject[WorkerReferencesModel]       `tfsdk:"references" json:"references,computed"`
}

func (m WorkerModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m WorkerModel) MarshalJSONForUpdate(state WorkerModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type WorkerObservabilityModel struct {
	Enabled          types.Bool                                             `tfsdk:"enabled" json:"enabled,computed_optional"`
	HeadSamplingRate types.Float64                                          `tfsdk:"head_sampling_rate" json:"head_sampling_rate,computed_optional"`
	Logs             customfield.NestedObject[WorkerObservabilityLogsModel] `tfsdk:"logs" json:"logs,computed_optional"`
}

type WorkerObservabilityLogsModel struct {
	Enabled          types.Bool    `tfsdk:"enabled" json:"enabled,computed_optional"`
	HeadSamplingRate types.Float64 `tfsdk:"head_sampling_rate" json:"head_sampling_rate,computed_optional"`
	InvocationLogs   types.Bool    `tfsdk:"invocation_logs" json:"invocation_logs,computed_optional"`
}

type WorkerSubdomainModel struct {
	Enabled         types.Bool `tfsdk:"enabled" json:"enabled,computed_optional"`
	PreviewsEnabled types.Bool `tfsdk:"previews_enabled" json:"previews_enabled,computed_optional"`
}

type WorkerTailConsumersModel struct {
	Name types.String `tfsdk:"name" json:"name,required"`
}

type WorkerReferencesModel struct {
	DispatchNamespaceOutbounds customfield.NestedObjectList[WorkerReferencesDispatchNamespaceOutboundsModel] `tfsdk:"dispatch_namespace_outbounds" json:"dispatch_namespace_outbounds,computed"`
	Domains                    customfield.NestedObjectList[WorkerReferencesDomainsModel]                    `tfsdk:"domains" json:"domains,computed"`
	DurableObjects             customfield.NestedObjectList[WorkerReferencesDurableObjectsModel]             `tfsdk:"durable_objects" json:"durable_objects,computed"`
	Queues                     customfield.NestedObjectList[WorkerReferencesQueuesModel]                     `tfsdk:"queues" json:"queues,computed"`
	Workers                    customfield.NestedObjectList[WorkerReferencesWorkersModel]                    `tfsdk:"workers" json:"workers,computed"`
}

type WorkerReferencesDispatchNamespaceOutboundsModel struct {
	NamespaceID   types.String `tfsdk:"namespace_id" json:"namespace_id,computed"`
	NamespaceName types.String `tfsdk:"namespace_name" json:"namespace_name,computed"`
	WorkerID      types.String `tfsdk:"worker_id" json:"worker_id,computed"`
	WorkerName    types.String `tfsdk:"worker_name" json:"worker_name,computed"`
}

type WorkerReferencesDomainsModel struct {
	ID            types.String `tfsdk:"id" json:"id,computed"`
	CertificateID types.String `tfsdk:"certificate_id" json:"certificate_id,computed"`
	Hostname      types.String `tfsdk:"hostname" json:"hostname,computed"`
	ZoneID        types.String `tfsdk:"zone_id" json:"zone_id,computed"`
	ZoneName      types.String `tfsdk:"zone_name" json:"zone_name,computed"`
}

type WorkerReferencesDurableObjectsModel struct {
	NamespaceID   types.String `tfsdk:"namespace_id" json:"namespace_id,computed"`
	NamespaceName types.String `tfsdk:"namespace_name" json:"namespace_name,computed"`
	WorkerID      types.String `tfsdk:"worker_id" json:"worker_id,computed"`
	WorkerName    types.String `tfsdk:"worker_name" json:"worker_name,computed"`
}

type WorkerReferencesQueuesModel struct {
	QueueConsumerID types.String `tfsdk:"queue_consumer_id" json:"queue_consumer_id,computed"`
	QueueID         types.String `tfsdk:"queue_id" json:"queue_id,computed"`
	QueueName       types.String `tfsdk:"queue_name" json:"queue_name,computed"`
}

type WorkerReferencesWorkersModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}
