// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_project

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/pages"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PagesProjectsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[PagesProjectsResultDataSourceModel] `json:"result,computed"`
}

type PagesProjectsDataSourceModel struct {
	AccountID types.String                                                     `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                      `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[PagesProjectsResultDataSourceModel] `tfsdk:"result"`
}

func (m *PagesProjectsDataSourceModel) toListParams(_ context.Context) (params pages.ProjectListParams, diags diag.Diagnostics) {
	params = pages.ProjectListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type PagesProjectsResultDataSourceModel struct {
	ID                   types.String                                                              `tfsdk:"id" json:"id,computed"`
	CanonicalDeployment  customfield.NestedObject[PagesProjectsCanonicalDeploymentDataSourceModel] `tfsdk:"canonical_deployment" json:"canonical_deployment,computed"`
	CreatedOn            timetypes.RFC3339                                                         `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	DeploymentConfigs    customfield.NestedObject[PagesProjectsDeploymentConfigsDataSourceModel]   `tfsdk:"deployment_configs" json:"deployment_configs,computed"`
	Framework            types.String                                                              `tfsdk:"framework" json:"framework,computed"`
	FrameworkVersion     types.String                                                              `tfsdk:"framework_version" json:"framework_version,computed"`
	LatestDeployment     customfield.NestedObject[PagesProjectsLatestDeploymentDataSourceModel]    `tfsdk:"latest_deployment" json:"latest_deployment,computed"`
	Name                 types.String                                                              `tfsdk:"name" json:"name,computed"`
	PreviewScriptName    types.String                                                              `tfsdk:"preview_script_name" json:"preview_script_name,computed"`
	ProductionBranch     types.String                                                              `tfsdk:"production_branch" json:"production_branch,computed"`
	ProductionScriptName types.String                                                              `tfsdk:"production_script_name" json:"production_script_name,computed"`
	UsesFunctions        types.Bool                                                                `tfsdk:"uses_functions" json:"uses_functions,computed"`
	BuildConfig          customfield.NestedObject[PagesProjectsBuildConfigDataSourceModel]         `tfsdk:"build_config" json:"build_config,computed"`
	Domains              customfield.List[types.String]                                            `tfsdk:"domains" json:"domains,computed"`
	Source               customfield.NestedObject[PagesProjectsSourceDataSourceModel]              `tfsdk:"source" json:"source,computed"`
	Subdomain            types.String                                                              `tfsdk:"subdomain" json:"subdomain,computed"`
}

type PagesProjectsCanonicalDeploymentDataSourceModel struct {
	ID                types.String                                                                               `tfsdk:"id" json:"id,computed"`
	Aliases           customfield.List[types.String]                                                             `tfsdk:"aliases" json:"aliases,computed"`
	BuildConfig       customfield.NestedObject[PagesProjectsCanonicalDeploymentBuildConfigDataSourceModel]       `tfsdk:"build_config" json:"build_config,computed"`
	CreatedOn         timetypes.RFC3339                                                                          `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	DeploymentTrigger customfield.NestedObject[PagesProjectsCanonicalDeploymentDeploymentTriggerDataSourceModel] `tfsdk:"deployment_trigger" json:"deployment_trigger,computed"`
	EnvVars           customfield.NestedObjectMap[PagesProjectsCanonicalDeploymentEnvVarsDataSourceModel]        `tfsdk:"env_vars" json:"env_vars,computed"`
	Environment       types.String                                                                               `tfsdk:"environment" json:"environment,computed"`
	IsSkipped         types.Bool                                                                                 `tfsdk:"is_skipped" json:"is_skipped,computed"`
	LatestStage       customfield.NestedObject[PagesProjectsCanonicalDeploymentLatestStageDataSourceModel]       `tfsdk:"latest_stage" json:"latest_stage,computed"`
	ModifiedOn        timetypes.RFC3339                                                                          `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	ProjectID         types.String                                                                               `tfsdk:"project_id" json:"project_id,computed"`
	ProjectName       types.String                                                                               `tfsdk:"project_name" json:"project_name,computed"`
	ShortID           types.String                                                                               `tfsdk:"short_id" json:"short_id,computed"`
	Source            customfield.NestedObject[PagesProjectsCanonicalDeploymentSourceDataSourceModel]            `tfsdk:"source" json:"source,computed"`
	Stages            customfield.NestedObjectList[PagesProjectsCanonicalDeploymentStagesDataSourceModel]        `tfsdk:"stages" json:"stages,computed"`
	URL               types.String                                                                               `tfsdk:"url" json:"url,computed"`
	UsesFunctions     types.Bool                                                                                 `tfsdk:"uses_functions" json:"uses_functions,computed"`
}

type PagesProjectsCanonicalDeploymentBuildConfigDataSourceModel struct {
	WebAnalyticsTag   types.String `tfsdk:"web_analytics_tag" json:"web_analytics_tag,computed"`
	WebAnalyticsToken types.String `tfsdk:"web_analytics_token" json:"web_analytics_token,computed"`
	BuildCaching      types.Bool   `tfsdk:"build_caching" json:"build_caching,computed"`
	BuildCommand      types.String `tfsdk:"build_command" json:"build_command,computed"`
	DestinationDir    types.String `tfsdk:"destination_dir" json:"destination_dir,computed"`
	RootDir           types.String `tfsdk:"root_dir" json:"root_dir,computed"`
}

type PagesProjectsCanonicalDeploymentDeploymentTriggerDataSourceModel struct {
	Metadata customfield.NestedObject[PagesProjectsCanonicalDeploymentDeploymentTriggerMetadataDataSourceModel] `tfsdk:"metadata" json:"metadata,computed"`
	Type     types.String                                                                                       `tfsdk:"type" json:"type,computed"`
}

type PagesProjectsCanonicalDeploymentDeploymentTriggerMetadataDataSourceModel struct {
	Branch        types.String `tfsdk:"branch" json:"branch,computed"`
	CommitDirty   types.Bool   `tfsdk:"commit_dirty" json:"commit_dirty,computed"`
	CommitHash    types.String `tfsdk:"commit_hash" json:"commit_hash,computed"`
	CommitMessage types.String `tfsdk:"commit_message" json:"commit_message,computed"`
}

type PagesProjectsCanonicalDeploymentEnvVarsDataSourceModel struct {
	Type  types.String `tfsdk:"type" json:"type,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}

type PagesProjectsCanonicalDeploymentLatestStageDataSourceModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed" format:"date-time"`
	Name      types.String      `tfsdk:"name" json:"name,computed"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed" format:"date-time"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
}

type PagesProjectsCanonicalDeploymentSourceDataSourceModel struct {
	Config customfield.NestedObject[PagesProjectsCanonicalDeploymentSourceConfigDataSourceModel] `tfsdk:"config" json:"config,computed"`
	Type   types.String                                                                          `tfsdk:"type" json:"type,computed"`
}

type PagesProjectsCanonicalDeploymentSourceConfigDataSourceModel struct {
	DeploymentsEnabled           types.Bool                     `tfsdk:"deployments_enabled" json:"deployments_enabled,computed"`
	Owner                        types.String                   `tfsdk:"owner" json:"owner,computed"`
	OwnerID                      types.String                   `tfsdk:"owner_id" json:"owner_id,computed"`
	PathExcludes                 customfield.List[types.String] `tfsdk:"path_excludes" json:"path_excludes,computed"`
	PathIncludes                 customfield.List[types.String] `tfsdk:"path_includes" json:"path_includes,computed"`
	PrCommentsEnabled            types.Bool                     `tfsdk:"pr_comments_enabled" json:"pr_comments_enabled,computed"`
	PreviewBranchExcludes        customfield.List[types.String] `tfsdk:"preview_branch_excludes" json:"preview_branch_excludes,computed"`
	PreviewBranchIncludes        customfield.List[types.String] `tfsdk:"preview_branch_includes" json:"preview_branch_includes,computed"`
	PreviewDeploymentSetting     types.String                   `tfsdk:"preview_deployment_setting" json:"preview_deployment_setting,computed"`
	ProductionBranch             types.String                   `tfsdk:"production_branch" json:"production_branch,computed"`
	ProductionDeploymentsEnabled types.Bool                     `tfsdk:"production_deployments_enabled" json:"production_deployments_enabled,computed"`
	RepoID                       types.String                   `tfsdk:"repo_id" json:"repo_id,computed"`
	RepoName                     types.String                   `tfsdk:"repo_name" json:"repo_name,computed"`
}

type PagesProjectsCanonicalDeploymentStagesDataSourceModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed" format:"date-time"`
	Name      types.String      `tfsdk:"name" json:"name,computed"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed" format:"date-time"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
}

type PagesProjectsDeploymentConfigsDataSourceModel struct {
	Preview    customfield.NestedObject[PagesProjectsDeploymentConfigsPreviewDataSourceModel]    `tfsdk:"preview" json:"preview,computed"`
	Production customfield.NestedObject[PagesProjectsDeploymentConfigsProductionDataSourceModel] `tfsdk:"production" json:"production,computed"`
}

type PagesProjectsDeploymentConfigsPreviewDataSourceModel struct {
	AlwaysUseLatestCompatibilityDate types.Bool                                                                                               `tfsdk:"always_use_latest_compatibility_date" json:"always_use_latest_compatibility_date,computed"`
	BuildImageMajorVersion           types.Int64                                                                                              `tfsdk:"build_image_major_version" json:"build_image_major_version,computed"`
	CompatibilityDate                types.String                                                                                             `tfsdk:"compatibility_date" json:"compatibility_date,computed"`
	CompatibilityFlags               customfield.List[types.String]                                                                           `tfsdk:"compatibility_flags" json:"compatibility_flags,computed"`
	EnvVars                          customfield.NestedObjectMap[PagesProjectsDeploymentConfigsPreviewEnvVarsDataSourceModel]                 `tfsdk:"env_vars" json:"env_vars,computed"`
	FailOpen                         types.Bool                                                                                               `tfsdk:"fail_open" json:"fail_open,computed"`
	UsageModel                       types.String                                                                                             `tfsdk:"usage_model" json:"usage_model,computed"`
	AIBindings                       customfield.NestedObjectMap[PagesProjectsDeploymentConfigsPreviewAIBindingsDataSourceModel]              `tfsdk:"ai_bindings" json:"ai_bindings,computed"`
	AnalyticsEngineDatasets          customfield.NestedObjectMap[PagesProjectsDeploymentConfigsPreviewAnalyticsEngineDatasetsDataSourceModel] `tfsdk:"analytics_engine_datasets" json:"analytics_engine_datasets,computed"`
	Browsers                         customfield.NestedObjectMap[PagesProjectsDeploymentConfigsPreviewBrowsersDataSourceModel]                `tfsdk:"browsers" json:"browsers,computed"`
	D1Databases                      customfield.NestedObjectMap[PagesProjectsDeploymentConfigsPreviewD1DatabasesDataSourceModel]             `tfsdk:"d1_databases" json:"d1_databases,computed"`
	DurableObjectNamespaces          customfield.NestedObjectMap[PagesProjectsDeploymentConfigsPreviewDurableObjectNamespacesDataSourceModel] `tfsdk:"durable_object_namespaces" json:"durable_object_namespaces,computed"`
	HyperdriveBindings               customfield.NestedObjectMap[PagesProjectsDeploymentConfigsPreviewHyperdriveBindingsDataSourceModel]      `tfsdk:"hyperdrive_bindings" json:"hyperdrive_bindings,computed"`
	KVNamespaces                     customfield.NestedObjectMap[PagesProjectsDeploymentConfigsPreviewKVNamespacesDataSourceModel]            `tfsdk:"kv_namespaces" json:"kv_namespaces,computed"`
	Limits                           customfield.NestedObject[PagesProjectsDeploymentConfigsPreviewLimitsDataSourceModel]                     `tfsdk:"limits" json:"limits,computed"`
	MTLSCertificates                 customfield.NestedObjectMap[PagesProjectsDeploymentConfigsPreviewMTLSCertificatesDataSourceModel]        `tfsdk:"mtls_certificates" json:"mtls_certificates,computed"`
	Placement                        customfield.NestedObject[PagesProjectsDeploymentConfigsPreviewPlacementDataSourceModel]                  `tfsdk:"placement" json:"placement,computed"`
	QueueProducers                   customfield.NestedObjectMap[PagesProjectsDeploymentConfigsPreviewQueueProducersDataSourceModel]          `tfsdk:"queue_producers" json:"queue_producers,computed"`
	R2Buckets                        customfield.NestedObjectMap[PagesProjectsDeploymentConfigsPreviewR2BucketsDataSourceModel]               `tfsdk:"r2_buckets" json:"r2_buckets,computed"`
	Services                         customfield.NestedObjectMap[PagesProjectsDeploymentConfigsPreviewServicesDataSourceModel]                `tfsdk:"services" json:"services,computed"`
	VectorizeBindings                customfield.NestedObjectMap[PagesProjectsDeploymentConfigsPreviewVectorizeBindingsDataSourceModel]       `tfsdk:"vectorize_bindings" json:"vectorize_bindings,computed"`
	WranglerConfigHash               types.String                                                                                             `tfsdk:"wrangler_config_hash" json:"wrangler_config_hash,computed"`
}

type PagesProjectsDeploymentConfigsPreviewEnvVarsDataSourceModel struct {
	Type  types.String `tfsdk:"type" json:"type,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}

type PagesProjectsDeploymentConfigsPreviewAIBindingsDataSourceModel struct {
	ProjectID types.String `tfsdk:"project_id" json:"project_id,computed"`
}

type PagesProjectsDeploymentConfigsPreviewAnalyticsEngineDatasetsDataSourceModel struct {
	Dataset types.String `tfsdk:"dataset" json:"dataset,computed"`
}

type PagesProjectsDeploymentConfigsPreviewBrowsersDataSourceModel struct {
}

type PagesProjectsDeploymentConfigsPreviewD1DatabasesDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type PagesProjectsDeploymentConfigsPreviewDurableObjectNamespacesDataSourceModel struct {
	NamespaceID types.String `tfsdk:"namespace_id" json:"namespace_id,computed"`
}

type PagesProjectsDeploymentConfigsPreviewHyperdriveBindingsDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type PagesProjectsDeploymentConfigsPreviewKVNamespacesDataSourceModel struct {
	NamespaceID types.String `tfsdk:"namespace_id" json:"namespace_id,computed"`
}

type PagesProjectsDeploymentConfigsPreviewLimitsDataSourceModel struct {
	CPUMs types.Int64 `tfsdk:"cpu_ms" json:"cpu_ms,computed"`
}

type PagesProjectsDeploymentConfigsPreviewMTLSCertificatesDataSourceModel struct {
	CertificateID types.String `tfsdk:"certificate_id" json:"certificate_id,computed"`
}

type PagesProjectsDeploymentConfigsPreviewPlacementDataSourceModel struct {
	Mode types.String `tfsdk:"mode" json:"mode,computed"`
}

type PagesProjectsDeploymentConfigsPreviewQueueProducersDataSourceModel struct {
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type PagesProjectsDeploymentConfigsPreviewR2BucketsDataSourceModel struct {
	Name         types.String `tfsdk:"name" json:"name,computed"`
	Jurisdiction types.String `tfsdk:"jurisdiction" json:"jurisdiction,computed"`
}

type PagesProjectsDeploymentConfigsPreviewServicesDataSourceModel struct {
	Environment types.String `tfsdk:"environment" json:"environment,computed"`
	Service     types.String `tfsdk:"service" json:"service,computed"`
	Entrypoint  types.String `tfsdk:"entrypoint" json:"entrypoint,computed"`
}

type PagesProjectsDeploymentConfigsPreviewVectorizeBindingsDataSourceModel struct {
	IndexName types.String `tfsdk:"index_name" json:"index_name,computed"`
}

type PagesProjectsDeploymentConfigsProductionDataSourceModel struct {
	AlwaysUseLatestCompatibilityDate types.Bool                                                                                                  `tfsdk:"always_use_latest_compatibility_date" json:"always_use_latest_compatibility_date,computed"`
	BuildImageMajorVersion           types.Int64                                                                                                 `tfsdk:"build_image_major_version" json:"build_image_major_version,computed"`
	CompatibilityDate                types.String                                                                                                `tfsdk:"compatibility_date" json:"compatibility_date,computed"`
	CompatibilityFlags               customfield.List[types.String]                                                                              `tfsdk:"compatibility_flags" json:"compatibility_flags,computed"`
	EnvVars                          customfield.NestedObjectMap[PagesProjectsDeploymentConfigsProductionEnvVarsDataSourceModel]                 `tfsdk:"env_vars" json:"env_vars,computed"`
	FailOpen                         types.Bool                                                                                                  `tfsdk:"fail_open" json:"fail_open,computed"`
	UsageModel                       types.String                                                                                                `tfsdk:"usage_model" json:"usage_model,computed"`
	AIBindings                       customfield.NestedObjectMap[PagesProjectsDeploymentConfigsProductionAIBindingsDataSourceModel]              `tfsdk:"ai_bindings" json:"ai_bindings,computed"`
	AnalyticsEngineDatasets          customfield.NestedObjectMap[PagesProjectsDeploymentConfigsProductionAnalyticsEngineDatasetsDataSourceModel] `tfsdk:"analytics_engine_datasets" json:"analytics_engine_datasets,computed"`
	Browsers                         customfield.NestedObjectMap[PagesProjectsDeploymentConfigsProductionBrowsersDataSourceModel]                `tfsdk:"browsers" json:"browsers,computed"`
	D1Databases                      customfield.NestedObjectMap[PagesProjectsDeploymentConfigsProductionD1DatabasesDataSourceModel]             `tfsdk:"d1_databases" json:"d1_databases,computed"`
	DurableObjectNamespaces          customfield.NestedObjectMap[PagesProjectsDeploymentConfigsProductionDurableObjectNamespacesDataSourceModel] `tfsdk:"durable_object_namespaces" json:"durable_object_namespaces,computed"`
	HyperdriveBindings               customfield.NestedObjectMap[PagesProjectsDeploymentConfigsProductionHyperdriveBindingsDataSourceModel]      `tfsdk:"hyperdrive_bindings" json:"hyperdrive_bindings,computed"`
	KVNamespaces                     customfield.NestedObjectMap[PagesProjectsDeploymentConfigsProductionKVNamespacesDataSourceModel]            `tfsdk:"kv_namespaces" json:"kv_namespaces,computed"`
	Limits                           customfield.NestedObject[PagesProjectsDeploymentConfigsProductionLimitsDataSourceModel]                     `tfsdk:"limits" json:"limits,computed"`
	MTLSCertificates                 customfield.NestedObjectMap[PagesProjectsDeploymentConfigsProductionMTLSCertificatesDataSourceModel]        `tfsdk:"mtls_certificates" json:"mtls_certificates,computed"`
	Placement                        customfield.NestedObject[PagesProjectsDeploymentConfigsProductionPlacementDataSourceModel]                  `tfsdk:"placement" json:"placement,computed"`
	QueueProducers                   customfield.NestedObjectMap[PagesProjectsDeploymentConfigsProductionQueueProducersDataSourceModel]          `tfsdk:"queue_producers" json:"queue_producers,computed"`
	R2Buckets                        customfield.NestedObjectMap[PagesProjectsDeploymentConfigsProductionR2BucketsDataSourceModel]               `tfsdk:"r2_buckets" json:"r2_buckets,computed"`
	Services                         customfield.NestedObjectMap[PagesProjectsDeploymentConfigsProductionServicesDataSourceModel]                `tfsdk:"services" json:"services,computed"`
	VectorizeBindings                customfield.NestedObjectMap[PagesProjectsDeploymentConfigsProductionVectorizeBindingsDataSourceModel]       `tfsdk:"vectorize_bindings" json:"vectorize_bindings,computed"`
	WranglerConfigHash               types.String                                                                                                `tfsdk:"wrangler_config_hash" json:"wrangler_config_hash,computed"`
}

type PagesProjectsDeploymentConfigsProductionEnvVarsDataSourceModel struct {
	Type  types.String `tfsdk:"type" json:"type,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}

type PagesProjectsDeploymentConfigsProductionAIBindingsDataSourceModel struct {
	ProjectID types.String `tfsdk:"project_id" json:"project_id,computed"`
}

type PagesProjectsDeploymentConfigsProductionAnalyticsEngineDatasetsDataSourceModel struct {
	Dataset types.String `tfsdk:"dataset" json:"dataset,computed"`
}

type PagesProjectsDeploymentConfigsProductionBrowsersDataSourceModel struct {
}

type PagesProjectsDeploymentConfigsProductionD1DatabasesDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type PagesProjectsDeploymentConfigsProductionDurableObjectNamespacesDataSourceModel struct {
	NamespaceID types.String `tfsdk:"namespace_id" json:"namespace_id,computed"`
}

type PagesProjectsDeploymentConfigsProductionHyperdriveBindingsDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type PagesProjectsDeploymentConfigsProductionKVNamespacesDataSourceModel struct {
	NamespaceID types.String `tfsdk:"namespace_id" json:"namespace_id,computed"`
}

type PagesProjectsDeploymentConfigsProductionLimitsDataSourceModel struct {
	CPUMs types.Int64 `tfsdk:"cpu_ms" json:"cpu_ms,computed"`
}

type PagesProjectsDeploymentConfigsProductionMTLSCertificatesDataSourceModel struct {
	CertificateID types.String `tfsdk:"certificate_id" json:"certificate_id,computed"`
}

type PagesProjectsDeploymentConfigsProductionPlacementDataSourceModel struct {
	Mode types.String `tfsdk:"mode" json:"mode,computed"`
}

type PagesProjectsDeploymentConfigsProductionQueueProducersDataSourceModel struct {
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type PagesProjectsDeploymentConfigsProductionR2BucketsDataSourceModel struct {
	Name         types.String `tfsdk:"name" json:"name,computed"`
	Jurisdiction types.String `tfsdk:"jurisdiction" json:"jurisdiction,computed"`
}

type PagesProjectsDeploymentConfigsProductionServicesDataSourceModel struct {
	Environment types.String `tfsdk:"environment" json:"environment,computed"`
	Service     types.String `tfsdk:"service" json:"service,computed"`
	Entrypoint  types.String `tfsdk:"entrypoint" json:"entrypoint,computed"`
}

type PagesProjectsDeploymentConfigsProductionVectorizeBindingsDataSourceModel struct {
	IndexName types.String `tfsdk:"index_name" json:"index_name,computed"`
}

type PagesProjectsLatestDeploymentDataSourceModel struct {
	ID                types.String                                                                            `tfsdk:"id" json:"id,computed"`
	Aliases           customfield.List[types.String]                                                          `tfsdk:"aliases" json:"aliases,computed"`
	BuildConfig       customfield.NestedObject[PagesProjectsLatestDeploymentBuildConfigDataSourceModel]       `tfsdk:"build_config" json:"build_config,computed"`
	CreatedOn         timetypes.RFC3339                                                                       `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	DeploymentTrigger customfield.NestedObject[PagesProjectsLatestDeploymentDeploymentTriggerDataSourceModel] `tfsdk:"deployment_trigger" json:"deployment_trigger,computed"`
	EnvVars           customfield.NestedObjectMap[PagesProjectsLatestDeploymentEnvVarsDataSourceModel]        `tfsdk:"env_vars" json:"env_vars,computed"`
	Environment       types.String                                                                            `tfsdk:"environment" json:"environment,computed"`
	IsSkipped         types.Bool                                                                              `tfsdk:"is_skipped" json:"is_skipped,computed"`
	LatestStage       customfield.NestedObject[PagesProjectsLatestDeploymentLatestStageDataSourceModel]       `tfsdk:"latest_stage" json:"latest_stage,computed"`
	ModifiedOn        timetypes.RFC3339                                                                       `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	ProjectID         types.String                                                                            `tfsdk:"project_id" json:"project_id,computed"`
	ProjectName       types.String                                                                            `tfsdk:"project_name" json:"project_name,computed"`
	ShortID           types.String                                                                            `tfsdk:"short_id" json:"short_id,computed"`
	Source            customfield.NestedObject[PagesProjectsLatestDeploymentSourceDataSourceModel]            `tfsdk:"source" json:"source,computed"`
	Stages            customfield.NestedObjectList[PagesProjectsLatestDeploymentStagesDataSourceModel]        `tfsdk:"stages" json:"stages,computed"`
	URL               types.String                                                                            `tfsdk:"url" json:"url,computed"`
	UsesFunctions     types.Bool                                                                              `tfsdk:"uses_functions" json:"uses_functions,computed"`
}

type PagesProjectsLatestDeploymentBuildConfigDataSourceModel struct {
	WebAnalyticsTag   types.String `tfsdk:"web_analytics_tag" json:"web_analytics_tag,computed"`
	WebAnalyticsToken types.String `tfsdk:"web_analytics_token" json:"web_analytics_token,computed"`
	BuildCaching      types.Bool   `tfsdk:"build_caching" json:"build_caching,computed"`
	BuildCommand      types.String `tfsdk:"build_command" json:"build_command,computed"`
	DestinationDir    types.String `tfsdk:"destination_dir" json:"destination_dir,computed"`
	RootDir           types.String `tfsdk:"root_dir" json:"root_dir,computed"`
}

type PagesProjectsLatestDeploymentDeploymentTriggerDataSourceModel struct {
	Metadata customfield.NestedObject[PagesProjectsLatestDeploymentDeploymentTriggerMetadataDataSourceModel] `tfsdk:"metadata" json:"metadata,computed"`
	Type     types.String                                                                                    `tfsdk:"type" json:"type,computed"`
}

type PagesProjectsLatestDeploymentDeploymentTriggerMetadataDataSourceModel struct {
	Branch        types.String `tfsdk:"branch" json:"branch,computed"`
	CommitDirty   types.Bool   `tfsdk:"commit_dirty" json:"commit_dirty,computed"`
	CommitHash    types.String `tfsdk:"commit_hash" json:"commit_hash,computed"`
	CommitMessage types.String `tfsdk:"commit_message" json:"commit_message,computed"`
}

type PagesProjectsLatestDeploymentEnvVarsDataSourceModel struct {
	Type  types.String `tfsdk:"type" json:"type,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}

type PagesProjectsLatestDeploymentLatestStageDataSourceModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed" format:"date-time"`
	Name      types.String      `tfsdk:"name" json:"name,computed"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed" format:"date-time"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
}

type PagesProjectsLatestDeploymentSourceDataSourceModel struct {
	Config customfield.NestedObject[PagesProjectsLatestDeploymentSourceConfigDataSourceModel] `tfsdk:"config" json:"config,computed"`
	Type   types.String                                                                       `tfsdk:"type" json:"type,computed"`
}

type PagesProjectsLatestDeploymentSourceConfigDataSourceModel struct {
	DeploymentsEnabled           types.Bool                     `tfsdk:"deployments_enabled" json:"deployments_enabled,computed"`
	Owner                        types.String                   `tfsdk:"owner" json:"owner,computed"`
	OwnerID                      types.String                   `tfsdk:"owner_id" json:"owner_id,computed"`
	PathExcludes                 customfield.List[types.String] `tfsdk:"path_excludes" json:"path_excludes,computed"`
	PathIncludes                 customfield.List[types.String] `tfsdk:"path_includes" json:"path_includes,computed"`
	PrCommentsEnabled            types.Bool                     `tfsdk:"pr_comments_enabled" json:"pr_comments_enabled,computed"`
	PreviewBranchExcludes        customfield.List[types.String] `tfsdk:"preview_branch_excludes" json:"preview_branch_excludes,computed"`
	PreviewBranchIncludes        customfield.List[types.String] `tfsdk:"preview_branch_includes" json:"preview_branch_includes,computed"`
	PreviewDeploymentSetting     types.String                   `tfsdk:"preview_deployment_setting" json:"preview_deployment_setting,computed"`
	ProductionBranch             types.String                   `tfsdk:"production_branch" json:"production_branch,computed"`
	ProductionDeploymentsEnabled types.Bool                     `tfsdk:"production_deployments_enabled" json:"production_deployments_enabled,computed"`
	RepoID                       types.String                   `tfsdk:"repo_id" json:"repo_id,computed"`
	RepoName                     types.String                   `tfsdk:"repo_name" json:"repo_name,computed"`
}

type PagesProjectsLatestDeploymentStagesDataSourceModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed" format:"date-time"`
	Name      types.String      `tfsdk:"name" json:"name,computed"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed" format:"date-time"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
}

type PagesProjectsBuildConfigDataSourceModel struct {
	WebAnalyticsTag   types.String `tfsdk:"web_analytics_tag" json:"web_analytics_tag,computed"`
	WebAnalyticsToken types.String `tfsdk:"web_analytics_token" json:"web_analytics_token,computed"`
	BuildCaching      types.Bool   `tfsdk:"build_caching" json:"build_caching,computed"`
	BuildCommand      types.String `tfsdk:"build_command" json:"build_command,computed"`
	DestinationDir    types.String `tfsdk:"destination_dir" json:"destination_dir,computed"`
	RootDir           types.String `tfsdk:"root_dir" json:"root_dir,computed"`
}

type PagesProjectsSourceDataSourceModel struct {
	Config customfield.NestedObject[PagesProjectsSourceConfigDataSourceModel] `tfsdk:"config" json:"config,computed"`
	Type   types.String                                                       `tfsdk:"type" json:"type,computed"`
}

type PagesProjectsSourceConfigDataSourceModel struct {
	DeploymentsEnabled           types.Bool                     `tfsdk:"deployments_enabled" json:"deployments_enabled,computed"`
	Owner                        types.String                   `tfsdk:"owner" json:"owner,computed"`
	OwnerID                      types.String                   `tfsdk:"owner_id" json:"owner_id,computed"`
	PathExcludes                 customfield.List[types.String] `tfsdk:"path_excludes" json:"path_excludes,computed"`
	PathIncludes                 customfield.List[types.String] `tfsdk:"path_includes" json:"path_includes,computed"`
	PrCommentsEnabled            types.Bool                     `tfsdk:"pr_comments_enabled" json:"pr_comments_enabled,computed"`
	PreviewBranchExcludes        customfield.List[types.String] `tfsdk:"preview_branch_excludes" json:"preview_branch_excludes,computed"`
	PreviewBranchIncludes        customfield.List[types.String] `tfsdk:"preview_branch_includes" json:"preview_branch_includes,computed"`
	PreviewDeploymentSetting     types.String                   `tfsdk:"preview_deployment_setting" json:"preview_deployment_setting,computed"`
	ProductionBranch             types.String                   `tfsdk:"production_branch" json:"production_branch,computed"`
	ProductionDeploymentsEnabled types.Bool                     `tfsdk:"production_deployments_enabled" json:"production_deployments_enabled,computed"`
	RepoID                       types.String                   `tfsdk:"repo_id" json:"repo_id,computed"`
	RepoName                     types.String                   `tfsdk:"repo_name" json:"repo_name,computed"`
}
