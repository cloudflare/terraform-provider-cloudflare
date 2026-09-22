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
	ID                 types.String                                            `tfsdk:"id" json:"id,computed"`
	AccountID          types.String                                            `tfsdk:"account_id" path:"account_id,required"`
	Name               types.String                                            `tfsdk:"name" json:"name,required"`
	Logpush            types.Bool                                              `tfsdk:"logpush" json:"logpush,computed_optional"`
	Tags               customfield.Set[types.String]                           `tfsdk:"tags" json:"tags,computed_optional"`
	Observability      customfield.NestedObject[WorkerObservabilityModel]      `tfsdk:"observability" json:"observability,computed_optional"`
	PreviewsBaseConfig customfield.NestedObject[WorkerPreviewsBaseConfigModel] `tfsdk:"previews_base_config" json:"previews_base_config,computed_optional"`
	Subdomain          customfield.NestedObject[WorkerSubdomainModel]          `tfsdk:"subdomain" json:"subdomain,computed_optional"`
	TailConsumers      customfield.NestedObjectSet[WorkerTailConsumersModel]   `tfsdk:"tail_consumers" json:"tail_consumers,computed_optional"`
	CreatedOn          timetypes.RFC3339                                       `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	DeployedOn         timetypes.RFC3339                                       `tfsdk:"deployed_on" json:"deployed_on,computed" format:"date-time"`
	UpdatedOn          timetypes.RFC3339                                       `tfsdk:"updated_on" json:"updated_on,computed" format:"date-time"`
	References         customfield.NestedObject[WorkerReferencesModel]         `tfsdk:"references" json:"references,computed"`
}

func (m WorkerModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m WorkerModel) MarshalJSONForUpdate(state WorkerModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type WorkerObservabilityModel struct {
	Enabled           types.Bool                                               `tfsdk:"enabled" json:"enabled,computed_optional"`
	HeadSamplingRate  types.Float64                                            `tfsdk:"head_sampling_rate" json:"head_sampling_rate,computed_optional"`
	Issues            customfield.NestedObject[WorkerObservabilityIssuesModel] `tfsdk:"issues" json:"issues,computed_optional"`
	Logs              customfield.NestedObject[WorkerObservabilityLogsModel]   `tfsdk:"logs" json:"logs,computed_optional"`
	RedactQueryString types.Bool                                               `tfsdk:"redact_query_string" json:"redact_query_string,computed_optional"`
	Traces            customfield.NestedObject[WorkerObservabilityTracesModel] `tfsdk:"traces" json:"traces,computed_optional"`
}

type WorkerObservabilityIssuesModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed_optional"`
}

type WorkerObservabilityLogsModel struct {
	Destinations     customfield.List[types.String] `tfsdk:"destinations" json:"destinations,computed_optional"`
	Enabled          types.Bool                     `tfsdk:"enabled" json:"enabled,computed_optional"`
	HeadSamplingRate types.Float64                  `tfsdk:"head_sampling_rate" json:"head_sampling_rate,computed_optional"`
	InvocationLogs   types.Bool                     `tfsdk:"invocation_logs" json:"invocation_logs,computed_optional"`
	Persist          types.Bool                     `tfsdk:"persist" json:"persist,computed_optional"`
}

type WorkerObservabilityTracesModel struct {
	Destinations      customfield.List[types.String] `tfsdk:"destinations" json:"destinations,computed_optional"`
	Enabled           types.Bool                     `tfsdk:"enabled" json:"enabled,computed_optional"`
	HeadSamplingRate  types.Float64                  `tfsdk:"head_sampling_rate" json:"head_sampling_rate,computed_optional"`
	Persist           types.Bool                     `tfsdk:"persist" json:"persist,computed_optional"`
	PropagationPolicy types.String                   `tfsdk:"propagation_policy" json:"propagation_policy,computed_optional"`
}

type WorkerPreviewsBaseConfigModel struct {
	CacheOptions  customfield.NestedObject[WorkerPreviewsBaseConfigCacheOptionsModel]     `tfsdk:"cache_options" json:"cache_options,computed_optional"`
	Env           *map[string]WorkerPreviewsBaseConfigEnvModel                            `tfsdk:"env" json:"env,optional"`
	Limits        customfield.NestedObject[WorkerPreviewsBaseConfigLimitsModel]           `tfsdk:"limits" json:"limits,computed_optional"`
	Logpush       types.Bool                                                              `tfsdk:"logpush" json:"logpush,optional"`
	Observability customfield.NestedObject[WorkerPreviewsBaseConfigObservabilityModel]    `tfsdk:"observability" json:"observability,computed_optional"`
	Placement     *WorkerPreviewsBaseConfigPlacementModel                                 `tfsdk:"placement" json:"placement,optional"`
	TailConsumers customfield.NestedObjectSet[WorkerPreviewsBaseConfigTailConsumersModel] `tfsdk:"tail_consumers" json:"tail_consumers,optional"`
}

type WorkerPreviewsBaseConfigCacheOptionsModel struct {
	Enabled           types.Bool `tfsdk:"enabled" json:"enabled,computed_optional"`
	CrossVersionCache types.Bool `tfsdk:"cross_version_cache" json:"cross_version_cache,computed_optional"`
}

type WorkerPreviewsBaseConfigEnvModel struct {
	Type types.String `tfsdk:"type" json:"type,required"`
}

type WorkerPreviewsBaseConfigLimitsModel struct {
	CPUMs       types.Int64 `tfsdk:"cpu_ms" json:"cpu_ms,optional"`
	Subrequests types.Int64 `tfsdk:"subrequests" json:"subrequests,optional"`
}

type WorkerPreviewsBaseConfigObservabilityModel struct {
	Enabled           types.Bool                                                                 `tfsdk:"enabled" json:"enabled,computed_optional"`
	HeadSamplingRate  types.Float64                                                              `tfsdk:"head_sampling_rate" json:"head_sampling_rate,computed_optional"`
	Issues            customfield.NestedObject[WorkerPreviewsBaseConfigObservabilityIssuesModel] `tfsdk:"issues" json:"issues,computed_optional"`
	Logs              customfield.NestedObject[WorkerPreviewsBaseConfigObservabilityLogsModel]   `tfsdk:"logs" json:"logs,computed_optional"`
	RedactQueryString types.Bool                                                                 `tfsdk:"redact_query_string" json:"redact_query_string,computed_optional"`
	Traces            customfield.NestedObject[WorkerPreviewsBaseConfigObservabilityTracesModel] `tfsdk:"traces" json:"traces,computed_optional"`
}

type WorkerPreviewsBaseConfigObservabilityIssuesModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed_optional"`
}

type WorkerPreviewsBaseConfigObservabilityLogsModel struct {
	Destinations     customfield.List[types.String] `tfsdk:"destinations" json:"destinations,computed_optional"`
	Enabled          types.Bool                     `tfsdk:"enabled" json:"enabled,computed_optional"`
	HeadSamplingRate types.Float64                  `tfsdk:"head_sampling_rate" json:"head_sampling_rate,computed_optional"`
	InvocationLogs   types.Bool                     `tfsdk:"invocation_logs" json:"invocation_logs,computed_optional"`
	Persist          types.Bool                     `tfsdk:"persist" json:"persist,computed_optional"`
}

type WorkerPreviewsBaseConfigObservabilityTracesModel struct {
	Destinations      customfield.List[types.String] `tfsdk:"destinations" json:"destinations,computed_optional"`
	Enabled           types.Bool                     `tfsdk:"enabled" json:"enabled,computed_optional"`
	HeadSamplingRate  types.Float64                  `tfsdk:"head_sampling_rate" json:"head_sampling_rate,computed_optional"`
	Persist           types.Bool                     `tfsdk:"persist" json:"persist,computed_optional"`
	PropagationPolicy types.String                   `tfsdk:"propagation_policy" json:"propagation_policy,optional"`
}

type WorkerPreviewsBaseConfigPlacementModel struct {
	Mode     types.String                                     `tfsdk:"mode" json:"mode,optional"`
	Region   types.String                                     `tfsdk:"region" json:"region,optional"`
	Hostname types.String                                     `tfsdk:"hostname" json:"hostname,optional"`
	Host     types.String                                     `tfsdk:"host" json:"host,optional"`
	Target   *[]*WorkerPreviewsBaseConfigPlacementTargetModel `tfsdk:"target" json:"target,optional"`
}

type WorkerPreviewsBaseConfigPlacementTargetModel struct {
	Region   types.String `tfsdk:"region" json:"region,optional"`
	Hostname types.String `tfsdk:"hostname" json:"hostname,optional"`
	Host     types.String `tfsdk:"host" json:"host,optional"`
}

type WorkerPreviewsBaseConfigTailConsumersModel struct {
	Name types.String `tfsdk:"name" json:"name,required"`
}

type WorkerSubdomainModel struct {
	Enabled          types.Bool   `tfsdk:"enabled" json:"enabled,computed_optional"`
	PreviewURLSuffix types.String `tfsdk:"preview_url_suffix" json:"preview_url_suffix,computed"`
	PreviewsEnabled  types.Bool   `tfsdk:"previews_enabled" json:"previews_enabled,computed_optional"`
	URL              types.String `tfsdk:"url" json:"url,computed"`
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
