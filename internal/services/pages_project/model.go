// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_project

import (
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type PagesProjectResultEnvelope struct {
Result PagesProjectModel `json:"result"`
}

type PagesProjectModel struct {
ID types.String `tfsdk:"id" json:"-,computed"`
Name types.String `tfsdk:"name" json:"name,required"`
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
ProductionBranch types.String `tfsdk:"production_branch" json:"production_branch,optional"`
BuildConfig customfield.NestedObject[PagesProjectBuildConfigModel] `tfsdk:"build_config" json:"build_config,computed_optional"`
DeploymentConfigs customfield.NestedObject[PagesProjectDeploymentConfigsModel] `tfsdk:"deployment_configs" json:"deployment_configs,computed_optional"`
CreatedOn timetypes.RFC3339 `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
Subdomain types.String `tfsdk:"subdomain" json:"subdomain,computed"`
Domains customfield.List[types.String] `tfsdk:"domains" json:"domains,computed"`
CanonicalDeployment customfield.NestedObject[PagesProjectCanonicalDeploymentModel] `tfsdk:"canonical_deployment" json:"canonical_deployment,computed"`
LatestDeployment customfield.NestedObject[PagesProjectLatestDeploymentModel] `tfsdk:"latest_deployment" json:"latest_deployment,computed"`
Source customfield.NestedObject[PagesProjectSourceModel] `tfsdk:"source" json:"source,computed"`
}

func (m PagesProjectModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m PagesProjectModel) MarshalJSONForUpdate(state PagesProjectModel) (data []byte, err error) {
  return apijson.MarshalForPatch(m, state)
}

type PagesProjectBuildConfigModel struct {
BuildCaching types.Bool `tfsdk:"build_caching" json:"build_caching,optional"`
BuildCommand types.String `tfsdk:"build_command" json:"build_command,optional"`
DestinationDir types.String `tfsdk:"destination_dir" json:"destination_dir,optional"`
RootDir types.String `tfsdk:"root_dir" json:"root_dir,optional"`
WebAnalyticsTag types.String `tfsdk:"web_analytics_tag" json:"web_analytics_tag,optional"`
WebAnalyticsToken types.String `tfsdk:"web_analytics_token" json:"web_analytics_token,optional"`
}

type PagesProjectDeploymentConfigsModel struct {
Preview customfield.NestedObject[PagesProjectDeploymentConfigsPreviewModel] `tfsdk:"preview" json:"preview,computed_optional"`
Production customfield.NestedObject[PagesProjectDeploymentConfigsProductionModel] `tfsdk:"production" json:"production,computed_optional"`
}

type PagesProjectDeploymentConfigsPreviewModel struct {
AIBindings customfield.NestedObjectMap[PagesProjectDeploymentConfigsPreviewAIBindingsModel] `tfsdk:"ai_bindings" json:"ai_bindings,computed_optional"`
AnalyticsEngineDatasets customfield.NestedObjectMap[PagesProjectDeploymentConfigsPreviewAnalyticsEngineDatasetsModel] `tfsdk:"analytics_engine_datasets" json:"analytics_engine_datasets,computed_optional"`
Browsers *map[string]PagesProjectDeploymentConfigsPreviewBrowsersModel `tfsdk:"browsers" json:"browsers,optional"`
CompatibilityDate types.String `tfsdk:"compatibility_date" json:"compatibility_date,optional"`
CompatibilityFlags *[]types.String `tfsdk:"compatibility_flags" json:"compatibility_flags,optional"`
D1Databases customfield.NestedObjectMap[PagesProjectDeploymentConfigsPreviewD1DatabasesModel] `tfsdk:"d1_databases" json:"d1_databases,computed_optional"`
DurableObjectNamespaces customfield.NestedObjectMap[PagesProjectDeploymentConfigsPreviewDurableObjectNamespacesModel] `tfsdk:"durable_object_namespaces" json:"durable_object_namespaces,computed_optional"`
EnvVars customfield.NestedObjectMap[PagesProjectDeploymentConfigsPreviewEnvVarsModel] `tfsdk:"env_vars" json:"env_vars,computed_optional"`
HyperdriveBindings customfield.NestedObjectMap[PagesProjectDeploymentConfigsPreviewHyperdriveBindingsModel] `tfsdk:"hyperdrive_bindings" json:"hyperdrive_bindings,computed_optional"`
KVNamespaces customfield.NestedObjectMap[PagesProjectDeploymentConfigsPreviewKVNamespacesModel] `tfsdk:"kv_namespaces" json:"kv_namespaces,computed_optional"`
MTLSCertificates customfield.NestedObjectMap[PagesProjectDeploymentConfigsPreviewMTLSCertificatesModel] `tfsdk:"mtls_certificates" json:"mtls_certificates,computed_optional"`
Placement customfield.NestedObject[PagesProjectDeploymentConfigsPreviewPlacementModel] `tfsdk:"placement" json:"placement,computed_optional"`
QueueProducers customfield.NestedObjectMap[PagesProjectDeploymentConfigsPreviewQueueProducersModel] `tfsdk:"queue_producers" json:"queue_producers,computed_optional"`
R2Buckets customfield.NestedObjectMap[PagesProjectDeploymentConfigsPreviewR2BucketsModel] `tfsdk:"r2_buckets" json:"r2_buckets,computed_optional"`
Services customfield.NestedObjectMap[PagesProjectDeploymentConfigsPreviewServicesModel] `tfsdk:"services" json:"services,computed_optional"`
VectorizeBindings customfield.NestedObjectMap[PagesProjectDeploymentConfigsPreviewVectorizeBindingsModel] `tfsdk:"vectorize_bindings" json:"vectorize_bindings,computed_optional"`
}

type PagesProjectDeploymentConfigsPreviewAIBindingsModel struct {
ProjectID types.String `tfsdk:"project_id" json:"project_id,optional"`
}

type PagesProjectDeploymentConfigsPreviewAnalyticsEngineDatasetsModel struct {
Dataset types.String `tfsdk:"dataset" json:"dataset,optional"`
}

type PagesProjectDeploymentConfigsPreviewBrowsersModel struct {
}

type PagesProjectDeploymentConfigsPreviewD1DatabasesModel struct {
ID types.String `tfsdk:"id" json:"id,optional"`
}

type PagesProjectDeploymentConfigsPreviewDurableObjectNamespacesModel struct {
NamespaceID types.String `tfsdk:"namespace_id" json:"namespace_id,optional"`
}

type PagesProjectDeploymentConfigsPreviewEnvVarsModel struct {
Type types.String `tfsdk:"type" json:"type,required"`
Value types.String `tfsdk:"value" json:"value,required"`
}

type PagesProjectDeploymentConfigsPreviewHyperdriveBindingsModel struct {
ID types.String `tfsdk:"id" json:"id,optional"`
}

type PagesProjectDeploymentConfigsPreviewKVNamespacesModel struct {
NamespaceID types.String `tfsdk:"namespace_id" json:"namespace_id,optional"`
}

type PagesProjectDeploymentConfigsPreviewMTLSCertificatesModel struct {
CertificateID types.String `tfsdk:"certificate_id" json:"certificate_id,optional"`
}

type PagesProjectDeploymentConfigsPreviewPlacementModel struct {
Mode types.String `tfsdk:"mode" json:"mode,optional"`
}

type PagesProjectDeploymentConfigsPreviewQueueProducersModel struct {
Name types.String `tfsdk:"name" json:"name,optional"`
}

type PagesProjectDeploymentConfigsPreviewR2BucketsModel struct {
Jurisdiction types.String `tfsdk:"jurisdiction" json:"jurisdiction,optional"`
Name types.String `tfsdk:"name" json:"name,optional"`
}

type PagesProjectDeploymentConfigsPreviewServicesModel struct {
Entrypoint types.String `tfsdk:"entrypoint" json:"entrypoint,optional"`
Environment types.String `tfsdk:"environment" json:"environment,optional"`
Service types.String `tfsdk:"service" json:"service,optional"`
}

type PagesProjectDeploymentConfigsPreviewVectorizeBindingsModel struct {
IndexName types.String `tfsdk:"index_name" json:"index_name,optional"`
}

type PagesProjectDeploymentConfigsProductionModel struct {
AIBindings customfield.NestedObjectMap[PagesProjectDeploymentConfigsProductionAIBindingsModel] `tfsdk:"ai_bindings" json:"ai_bindings,computed_optional"`
AnalyticsEngineDatasets customfield.NestedObjectMap[PagesProjectDeploymentConfigsProductionAnalyticsEngineDatasetsModel] `tfsdk:"analytics_engine_datasets" json:"analytics_engine_datasets,computed_optional"`
Browsers *map[string]PagesProjectDeploymentConfigsProductionBrowsersModel `tfsdk:"browsers" json:"browsers,optional"`
CompatibilityDate types.String `tfsdk:"compatibility_date" json:"compatibility_date,optional"`
CompatibilityFlags *[]types.String `tfsdk:"compatibility_flags" json:"compatibility_flags,optional"`
D1Databases customfield.NestedObjectMap[PagesProjectDeploymentConfigsProductionD1DatabasesModel] `tfsdk:"d1_databases" json:"d1_databases,computed_optional"`
DurableObjectNamespaces customfield.NestedObjectMap[PagesProjectDeploymentConfigsProductionDurableObjectNamespacesModel] `tfsdk:"durable_object_namespaces" json:"durable_object_namespaces,computed_optional"`
EnvVars customfield.NestedObjectMap[PagesProjectDeploymentConfigsProductionEnvVarsModel] `tfsdk:"env_vars" json:"env_vars,computed_optional"`
HyperdriveBindings customfield.NestedObjectMap[PagesProjectDeploymentConfigsProductionHyperdriveBindingsModel] `tfsdk:"hyperdrive_bindings" json:"hyperdrive_bindings,computed_optional"`
KVNamespaces customfield.NestedObjectMap[PagesProjectDeploymentConfigsProductionKVNamespacesModel] `tfsdk:"kv_namespaces" json:"kv_namespaces,computed_optional"`
MTLSCertificates customfield.NestedObjectMap[PagesProjectDeploymentConfigsProductionMTLSCertificatesModel] `tfsdk:"mtls_certificates" json:"mtls_certificates,computed_optional"`
Placement customfield.NestedObject[PagesProjectDeploymentConfigsProductionPlacementModel] `tfsdk:"placement" json:"placement,computed_optional"`
QueueProducers customfield.NestedObjectMap[PagesProjectDeploymentConfigsProductionQueueProducersModel] `tfsdk:"queue_producers" json:"queue_producers,computed_optional"`
R2Buckets customfield.NestedObjectMap[PagesProjectDeploymentConfigsProductionR2BucketsModel] `tfsdk:"r2_buckets" json:"r2_buckets,computed_optional"`
Services customfield.NestedObjectMap[PagesProjectDeploymentConfigsProductionServicesModel] `tfsdk:"services" json:"services,computed_optional"`
VectorizeBindings customfield.NestedObjectMap[PagesProjectDeploymentConfigsProductionVectorizeBindingsModel] `tfsdk:"vectorize_bindings" json:"vectorize_bindings,computed_optional"`
}

type PagesProjectDeploymentConfigsProductionAIBindingsModel struct {
ProjectID types.String `tfsdk:"project_id" json:"project_id,optional"`
}

type PagesProjectDeploymentConfigsProductionAnalyticsEngineDatasetsModel struct {
Dataset types.String `tfsdk:"dataset" json:"dataset,optional"`
}

type PagesProjectDeploymentConfigsProductionBrowsersModel struct {
}

type PagesProjectDeploymentConfigsProductionD1DatabasesModel struct {
ID types.String `tfsdk:"id" json:"id,optional"`
}

type PagesProjectDeploymentConfigsProductionDurableObjectNamespacesModel struct {
NamespaceID types.String `tfsdk:"namespace_id" json:"namespace_id,optional"`
}

type PagesProjectDeploymentConfigsProductionEnvVarsModel struct {
Type types.String `tfsdk:"type" json:"type,required"`
Value types.String `tfsdk:"value" json:"value,required"`
}

type PagesProjectDeploymentConfigsProductionHyperdriveBindingsModel struct {
ID types.String `tfsdk:"id" json:"id,optional"`
}

type PagesProjectDeploymentConfigsProductionKVNamespacesModel struct {
NamespaceID types.String `tfsdk:"namespace_id" json:"namespace_id,optional"`
}

type PagesProjectDeploymentConfigsProductionMTLSCertificatesModel struct {
CertificateID types.String `tfsdk:"certificate_id" json:"certificate_id,optional"`
}

type PagesProjectDeploymentConfigsProductionPlacementModel struct {
Mode types.String `tfsdk:"mode" json:"mode,optional"`
}

type PagesProjectDeploymentConfigsProductionQueueProducersModel struct {
Name types.String `tfsdk:"name" json:"name,optional"`
}

type PagesProjectDeploymentConfigsProductionR2BucketsModel struct {
Jurisdiction types.String `tfsdk:"jurisdiction" json:"jurisdiction,optional"`
Name types.String `tfsdk:"name" json:"name,optional"`
}

type PagesProjectDeploymentConfigsProductionServicesModel struct {
Entrypoint types.String `tfsdk:"entrypoint" json:"entrypoint,optional"`
Environment types.String `tfsdk:"environment" json:"environment,optional"`
Service types.String `tfsdk:"service" json:"service,optional"`
}

type PagesProjectDeploymentConfigsProductionVectorizeBindingsModel struct {
IndexName types.String `tfsdk:"index_name" json:"index_name,optional"`
}

type PagesProjectCanonicalDeploymentModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
Aliases customfield.List[types.String] `tfsdk:"aliases" json:"aliases,computed"`
BuildConfig customfield.NestedObject[PagesProjectCanonicalDeploymentBuildConfigModel] `tfsdk:"build_config" json:"build_config,computed"`
CreatedOn timetypes.RFC3339 `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
DeploymentTrigger customfield.NestedObject[PagesProjectCanonicalDeploymentDeploymentTriggerModel] `tfsdk:"deployment_trigger" json:"deployment_trigger,computed"`
EnvVars customfield.NestedObjectMap[PagesProjectCanonicalDeploymentEnvVarsModel] `tfsdk:"env_vars" json:"env_vars,computed"`
Environment types.String `tfsdk:"environment" json:"environment,computed"`
IsSkipped types.Bool `tfsdk:"is_skipped" json:"is_skipped,computed"`
LatestStage customfield.NestedObject[PagesProjectCanonicalDeploymentLatestStageModel] `tfsdk:"latest_stage" json:"latest_stage,computed"`
ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
ProjectID types.String `tfsdk:"project_id" json:"project_id,computed"`
ProjectName types.String `tfsdk:"project_name" json:"project_name,computed"`
ShortID types.String `tfsdk:"short_id" json:"short_id,computed"`
Source customfield.NestedObject[PagesProjectCanonicalDeploymentSourceModel] `tfsdk:"source" json:"source,computed"`
Stages customfield.NestedObjectList[PagesProjectCanonicalDeploymentStagesModel] `tfsdk:"stages" json:"stages,computed"`
URL types.String `tfsdk:"url" json:"url,computed"`
}

type PagesProjectCanonicalDeploymentBuildConfigModel struct {
BuildCaching types.Bool `tfsdk:"build_caching" json:"build_caching,computed"`
BuildCommand types.String `tfsdk:"build_command" json:"build_command,computed"`
DestinationDir types.String `tfsdk:"destination_dir" json:"destination_dir,computed"`
RootDir types.String `tfsdk:"root_dir" json:"root_dir,computed"`
WebAnalyticsTag types.String `tfsdk:"web_analytics_tag" json:"web_analytics_tag,computed"`
WebAnalyticsToken types.String `tfsdk:"web_analytics_token" json:"web_analytics_token,computed"`
}

type PagesProjectCanonicalDeploymentDeploymentTriggerModel struct {
Metadata customfield.NestedObject[PagesProjectCanonicalDeploymentDeploymentTriggerMetadataModel] `tfsdk:"metadata" json:"metadata,computed"`
Type types.String `tfsdk:"type" json:"type,computed"`
}

type PagesProjectCanonicalDeploymentDeploymentTriggerMetadataModel struct {
Branch types.String `tfsdk:"branch" json:"branch,computed"`
CommitHash types.String `tfsdk:"commit_hash" json:"commit_hash,computed"`
CommitMessage types.String `tfsdk:"commit_message" json:"commit_message,computed"`
}

type PagesProjectCanonicalDeploymentEnvVarsModel struct {
Type types.String `tfsdk:"type" json:"type,computed"`
Value types.String `tfsdk:"value" json:"value,computed"`
}

type PagesProjectCanonicalDeploymentLatestStageModel struct {
EndedOn timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed" format:"date-time"`
Name types.String `tfsdk:"name" json:"name,computed"`
StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed" format:"date-time"`
Status types.String `tfsdk:"status" json:"status,computed"`
}

type PagesProjectCanonicalDeploymentSourceModel struct {
Config customfield.NestedObject[PagesProjectCanonicalDeploymentSourceConfigModel] `tfsdk:"config" json:"config,computed"`
Type types.String `tfsdk:"type" json:"type,computed"`
}

type PagesProjectCanonicalDeploymentSourceConfigModel struct {
DeploymentsEnabled types.Bool `tfsdk:"deployments_enabled" json:"deployments_enabled,computed"`
Owner types.String `tfsdk:"owner" json:"owner,computed"`
PathExcludes customfield.List[types.String] `tfsdk:"path_excludes" json:"path_excludes,computed"`
PathIncludes customfield.List[types.String] `tfsdk:"path_includes" json:"path_includes,computed"`
PrCommentsEnabled types.Bool `tfsdk:"pr_comments_enabled" json:"pr_comments_enabled,computed"`
PreviewBranchExcludes customfield.List[types.String] `tfsdk:"preview_branch_excludes" json:"preview_branch_excludes,computed"`
PreviewBranchIncludes customfield.List[types.String] `tfsdk:"preview_branch_includes" json:"preview_branch_includes,computed"`
PreviewDeploymentSetting types.String `tfsdk:"preview_deployment_setting" json:"preview_deployment_setting,computed"`
ProductionBranch types.String `tfsdk:"production_branch" json:"production_branch,computed"`
ProductionDeploymentsEnabled types.Bool `tfsdk:"production_deployments_enabled" json:"production_deployments_enabled,computed"`
RepoName types.String `tfsdk:"repo_name" json:"repo_name,computed"`
}

type PagesProjectCanonicalDeploymentStagesModel struct {
EndedOn timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed" format:"date-time"`
Name types.String `tfsdk:"name" json:"name,computed"`
StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed" format:"date-time"`
Status types.String `tfsdk:"status" json:"status,computed"`
}

type PagesProjectLatestDeploymentModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
Aliases customfield.List[types.String] `tfsdk:"aliases" json:"aliases,computed"`
BuildConfig customfield.NestedObject[PagesProjectLatestDeploymentBuildConfigModel] `tfsdk:"build_config" json:"build_config,computed"`
CreatedOn timetypes.RFC3339 `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
DeploymentTrigger customfield.NestedObject[PagesProjectLatestDeploymentDeploymentTriggerModel] `tfsdk:"deployment_trigger" json:"deployment_trigger,computed"`
EnvVars customfield.NestedObjectMap[PagesProjectLatestDeploymentEnvVarsModel] `tfsdk:"env_vars" json:"env_vars,computed"`
Environment types.String `tfsdk:"environment" json:"environment,computed"`
IsSkipped types.Bool `tfsdk:"is_skipped" json:"is_skipped,computed"`
LatestStage customfield.NestedObject[PagesProjectLatestDeploymentLatestStageModel] `tfsdk:"latest_stage" json:"latest_stage,computed"`
ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
ProjectID types.String `tfsdk:"project_id" json:"project_id,computed"`
ProjectName types.String `tfsdk:"project_name" json:"project_name,computed"`
ShortID types.String `tfsdk:"short_id" json:"short_id,computed"`
Source customfield.NestedObject[PagesProjectLatestDeploymentSourceModel] `tfsdk:"source" json:"source,computed"`
Stages customfield.NestedObjectList[PagesProjectLatestDeploymentStagesModel] `tfsdk:"stages" json:"stages,computed"`
URL types.String `tfsdk:"url" json:"url,computed"`
}

type PagesProjectLatestDeploymentBuildConfigModel struct {
BuildCaching types.Bool `tfsdk:"build_caching" json:"build_caching,computed"`
BuildCommand types.String `tfsdk:"build_command" json:"build_command,computed"`
DestinationDir types.String `tfsdk:"destination_dir" json:"destination_dir,computed"`
RootDir types.String `tfsdk:"root_dir" json:"root_dir,computed"`
WebAnalyticsTag types.String `tfsdk:"web_analytics_tag" json:"web_analytics_tag,computed"`
WebAnalyticsToken types.String `tfsdk:"web_analytics_token" json:"web_analytics_token,computed"`
}

type PagesProjectLatestDeploymentDeploymentTriggerModel struct {
Metadata customfield.NestedObject[PagesProjectLatestDeploymentDeploymentTriggerMetadataModel] `tfsdk:"metadata" json:"metadata,computed"`
Type types.String `tfsdk:"type" json:"type,computed"`
}

type PagesProjectLatestDeploymentDeploymentTriggerMetadataModel struct {
Branch types.String `tfsdk:"branch" json:"branch,computed"`
CommitHash types.String `tfsdk:"commit_hash" json:"commit_hash,computed"`
CommitMessage types.String `tfsdk:"commit_message" json:"commit_message,computed"`
}

type PagesProjectLatestDeploymentEnvVarsModel struct {
Type types.String `tfsdk:"type" json:"type,computed"`
Value types.String `tfsdk:"value" json:"value,computed"`
}

type PagesProjectLatestDeploymentLatestStageModel struct {
EndedOn timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed" format:"date-time"`
Name types.String `tfsdk:"name" json:"name,computed"`
StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed" format:"date-time"`
Status types.String `tfsdk:"status" json:"status,computed"`
}

type PagesProjectLatestDeploymentSourceModel struct {
Config customfield.NestedObject[PagesProjectLatestDeploymentSourceConfigModel] `tfsdk:"config" json:"config,computed"`
Type types.String `tfsdk:"type" json:"type,computed"`
}

type PagesProjectLatestDeploymentSourceConfigModel struct {
DeploymentsEnabled types.Bool `tfsdk:"deployments_enabled" json:"deployments_enabled,computed"`
Owner types.String `tfsdk:"owner" json:"owner,computed"`
PathExcludes customfield.List[types.String] `tfsdk:"path_excludes" json:"path_excludes,computed"`
PathIncludes customfield.List[types.String] `tfsdk:"path_includes" json:"path_includes,computed"`
PrCommentsEnabled types.Bool `tfsdk:"pr_comments_enabled" json:"pr_comments_enabled,computed"`
PreviewBranchExcludes customfield.List[types.String] `tfsdk:"preview_branch_excludes" json:"preview_branch_excludes,computed"`
PreviewBranchIncludes customfield.List[types.String] `tfsdk:"preview_branch_includes" json:"preview_branch_includes,computed"`
PreviewDeploymentSetting types.String `tfsdk:"preview_deployment_setting" json:"preview_deployment_setting,computed"`
ProductionBranch types.String `tfsdk:"production_branch" json:"production_branch,computed"`
ProductionDeploymentsEnabled types.Bool `tfsdk:"production_deployments_enabled" json:"production_deployments_enabled,computed"`
RepoName types.String `tfsdk:"repo_name" json:"repo_name,computed"`
}

type PagesProjectLatestDeploymentStagesModel struct {
EndedOn timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed" format:"date-time"`
Name types.String `tfsdk:"name" json:"name,computed"`
StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed" format:"date-time"`
Status types.String `tfsdk:"status" json:"status,computed"`
}

type PagesProjectSourceModel struct {
Config customfield.NestedObject[PagesProjectSourceConfigModel] `tfsdk:"config" json:"config,computed"`
Type types.String `tfsdk:"type" json:"type,computed"`
}

type PagesProjectSourceConfigModel struct {
DeploymentsEnabled types.Bool `tfsdk:"deployments_enabled" json:"deployments_enabled,computed"`
Owner types.String `tfsdk:"owner" json:"owner,computed"`
PathExcludes customfield.List[types.String] `tfsdk:"path_excludes" json:"path_excludes,computed"`
PathIncludes customfield.List[types.String] `tfsdk:"path_includes" json:"path_includes,computed"`
PrCommentsEnabled types.Bool `tfsdk:"pr_comments_enabled" json:"pr_comments_enabled,computed"`
PreviewBranchExcludes customfield.List[types.String] `tfsdk:"preview_branch_excludes" json:"preview_branch_excludes,computed"`
PreviewBranchIncludes customfield.List[types.String] `tfsdk:"preview_branch_includes" json:"preview_branch_includes,computed"`
PreviewDeploymentSetting types.String `tfsdk:"preview_deployment_setting" json:"preview_deployment_setting,computed"`
ProductionBranch types.String `tfsdk:"production_branch" json:"production_branch,computed"`
ProductionDeploymentsEnabled types.Bool `tfsdk:"production_deployments_enabled" json:"production_deployments_enabled,computed"`
RepoName types.String `tfsdk:"repo_name" json:"repo_name,computed"`
}
