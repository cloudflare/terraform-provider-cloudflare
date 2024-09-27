// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_project

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/pages"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PagesProjectResultDataSourceEnvelope struct {
	Result PagesProjectDataSourceModel `json:"result,computed"`
}

type PagesProjectResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[PagesProjectDataSourceModel] `json:"result,computed"`
}

type PagesProjectDataSourceModel struct {
	AccountID           types.String                                                     `tfsdk:"account_id" path:"account_id,optional"`
	ProjectName         types.String                                                     `tfsdk:"project_name" path:"project_name,computed_optional"`
	Environment         types.String                                                     `tfsdk:"environment" json:"environment,optional"`
	IsSkipped           types.Bool                                                       `tfsdk:"is_skipped" json:"is_skipped,optional"`
	ModifiedOn          timetypes.RFC3339                                                `tfsdk:"modified_on" json:"modified_on,optional" format:"date-time"`
	Name                types.String                                                     `tfsdk:"name" json:"name,optional"`
	ProductionBranch    types.String                                                     `tfsdk:"production_branch" json:"production_branch,optional"`
	ProjectID           types.String                                                     `tfsdk:"project_id" json:"project_id,optional"`
	ShortID             types.String                                                     `tfsdk:"short_id" json:"short_id,optional"`
	Subdomain           types.String                                                     `tfsdk:"subdomain" json:"subdomain,optional"`
	URL                 types.String                                                     `tfsdk:"url" json:"url,optional"`
	Aliases             *[]types.String                                                  `tfsdk:"aliases" json:"aliases,optional"`
	Domains             *[]types.String                                                  `tfsdk:"domains" json:"domains,optional"`
	CanonicalDeployment *PagesProjectCanonicalDeploymentDataSourceModel                  `tfsdk:"canonical_deployment" json:"canonical_deployment,optional"`
	DeploymentConfigs   *PagesProjectDeploymentConfigsDataSourceModel                    `tfsdk:"deployment_configs" json:"deployment_configs,optional"`
	DeploymentTrigger   *PagesProjectDeploymentTriggerDataSourceModel                    `tfsdk:"deployment_trigger" json:"deployment_trigger,optional"`
	EnvVars             map[string]PagesProjectEnvVarsDataSourceModel                    `tfsdk:"env_vars" json:"env_vars,optional"`
	LatestDeployment    *PagesProjectLatestDeploymentDataSourceModel                     `tfsdk:"latest_deployment" json:"latest_deployment,optional"`
	LatestStage         *PagesProjectLatestStageDataSourceModel                          `tfsdk:"latest_stage" json:"latest_stage,optional"`
	Stages              *[]*PagesProjectStagesDataSourceModel                            `tfsdk:"stages" json:"stages,optional"`
	CreatedOn           timetypes.RFC3339                                                `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ID                  types.String                                                     `tfsdk:"id" json:"id,computed"`
	BuildConfig         customfield.NestedObject[PagesProjectBuildConfigDataSourceModel] `tfsdk:"build_config" json:"build_config,computed"`
	Source              customfield.NestedObject[PagesProjectSourceDataSourceModel]      `tfsdk:"source" json:"source,computed"`
	Filter              *PagesProjectFindOneByDataSourceModel                            `tfsdk:"filter"`
}

func (m *PagesProjectDataSourceModel) toReadParams(_ context.Context) (params pages.ProjectGetParams, diags diag.Diagnostics) {
	params = pages.ProjectGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *PagesProjectDataSourceModel) toListParams(_ context.Context) (params pages.ProjectListParams, diags diag.Diagnostics) {
	params = pages.ProjectListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type PagesProjectCanonicalDeploymentDataSourceModel struct {
	ID                types.String                                                                              `tfsdk:"id" json:"id,computed"`
	Aliases           customfield.List[types.String]                                                            `tfsdk:"aliases" json:"aliases,computed"`
	BuildConfig       customfield.NestedObject[PagesProjectCanonicalDeploymentBuildConfigDataSourceModel]       `tfsdk:"build_config" json:"build_config,computed"`
	CreatedOn         timetypes.RFC3339                                                                         `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	DeploymentTrigger customfield.NestedObject[PagesProjectCanonicalDeploymentDeploymentTriggerDataSourceModel] `tfsdk:"deployment_trigger" json:"deployment_trigger,computed"`
	EnvVars           customfield.NestedObjectMap[PagesProjectCanonicalDeploymentEnvVarsDataSourceModel]        `tfsdk:"env_vars" json:"env_vars,computed"`
	Environment       types.String                                                                              `tfsdk:"environment" json:"environment,computed"`
	IsSkipped         types.Bool                                                                                `tfsdk:"is_skipped" json:"is_skipped,computed"`
	LatestStage       customfield.NestedObject[PagesProjectCanonicalDeploymentLatestStageDataSourceModel]       `tfsdk:"latest_stage" json:"latest_stage,computed"`
	ModifiedOn        timetypes.RFC3339                                                                         `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	ProjectID         types.String                                                                              `tfsdk:"project_id" json:"project_id,computed"`
	ProjectName       types.String                                                                              `tfsdk:"project_name" json:"project_name,computed"`
	ShortID           types.String                                                                              `tfsdk:"short_id" json:"short_id,computed"`
	Source            customfield.NestedObject[PagesProjectCanonicalDeploymentSourceDataSourceModel]            `tfsdk:"source" json:"source,computed"`
	Stages            customfield.NestedObjectList[PagesProjectCanonicalDeploymentStagesDataSourceModel]        `tfsdk:"stages" json:"stages,computed"`
	URL               types.String                                                                              `tfsdk:"url" json:"url,computed"`
}

type PagesProjectCanonicalDeploymentBuildConfigDataSourceModel struct {
	BuildCaching      types.Bool   `tfsdk:"build_caching" json:"build_caching,computed"`
	BuildCommand      types.String `tfsdk:"build_command" json:"build_command,computed"`
	DestinationDir    types.String `tfsdk:"destination_dir" json:"destination_dir,computed"`
	RootDir           types.String `tfsdk:"root_dir" json:"root_dir,computed"`
	WebAnalyticsTag   types.String `tfsdk:"web_analytics_tag" json:"web_analytics_tag,computed"`
	WebAnalyticsToken types.String `tfsdk:"web_analytics_token" json:"web_analytics_token,computed"`
}

type PagesProjectCanonicalDeploymentDeploymentTriggerDataSourceModel struct {
	Metadata customfield.NestedObject[PagesProjectCanonicalDeploymentDeploymentTriggerMetadataDataSourceModel] `tfsdk:"metadata" json:"metadata,computed"`
	Type     types.String                                                                                      `tfsdk:"type" json:"type,computed"`
}

type PagesProjectCanonicalDeploymentDeploymentTriggerMetadataDataSourceModel struct {
	Branch        types.String `tfsdk:"branch" json:"branch,computed"`
	CommitHash    types.String `tfsdk:"commit_hash" json:"commit_hash,computed"`
	CommitMessage types.String `tfsdk:"commit_message" json:"commit_message,computed"`
}

type PagesProjectCanonicalDeploymentEnvVarsDataSourceModel struct {
	Value types.String `tfsdk:"value" json:"value,computed"`
	Type  types.String `tfsdk:"type" json:"type,computed"`
}

type PagesProjectCanonicalDeploymentLatestStageDataSourceModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed" format:"date-time"`
	Name      types.String      `tfsdk:"name" json:"name,computed"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed" format:"date-time"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
}

type PagesProjectCanonicalDeploymentSourceDataSourceModel struct {
	Config customfield.NestedObject[PagesProjectCanonicalDeploymentSourceConfigDataSourceModel] `tfsdk:"config" json:"config,computed"`
	Type   types.String                                                                         `tfsdk:"type" json:"type,computed"`
}

type PagesProjectCanonicalDeploymentSourceConfigDataSourceModel struct {
	DeploymentsEnabled           types.Bool                     `tfsdk:"deployments_enabled" json:"deployments_enabled,computed"`
	Owner                        types.String                   `tfsdk:"owner" json:"owner,computed"`
	PathExcludes                 customfield.List[types.String] `tfsdk:"path_excludes" json:"path_excludes,computed"`
	PathIncludes                 customfield.List[types.String] `tfsdk:"path_includes" json:"path_includes,computed"`
	PrCommentsEnabled            types.Bool                     `tfsdk:"pr_comments_enabled" json:"pr_comments_enabled,computed"`
	PreviewBranchExcludes        customfield.List[types.String] `tfsdk:"preview_branch_excludes" json:"preview_branch_excludes,computed"`
	PreviewBranchIncludes        customfield.List[types.String] `tfsdk:"preview_branch_includes" json:"preview_branch_includes,computed"`
	PreviewDeploymentSetting     types.String                   `tfsdk:"preview_deployment_setting" json:"preview_deployment_setting,computed"`
	ProductionBranch             types.String                   `tfsdk:"production_branch" json:"production_branch,computed"`
	ProductionDeploymentsEnabled types.Bool                     `tfsdk:"production_deployments_enabled" json:"production_deployments_enabled,computed"`
	RepoName                     types.String                   `tfsdk:"repo_name" json:"repo_name,computed"`
}

type PagesProjectCanonicalDeploymentStagesDataSourceModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed" format:"date-time"`
	Name      types.String      `tfsdk:"name" json:"name,computed"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed" format:"date-time"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
}

type PagesProjectDeploymentConfigsDataSourceModel struct {
	Preview    customfield.NestedObject[PagesProjectDeploymentConfigsPreviewDataSourceModel]    `tfsdk:"preview" json:"preview,computed"`
	Production customfield.NestedObject[PagesProjectDeploymentConfigsProductionDataSourceModel] `tfsdk:"production" json:"production,computed"`
}

type PagesProjectDeploymentConfigsPreviewDataSourceModel struct {
	AIBindings              customfield.NestedObjectMap[PagesProjectDeploymentConfigsPreviewAIBindingsDataSourceModel]              `tfsdk:"ai_bindings" json:"ai_bindings,computed"`
	AnalyticsEngineDatasets customfield.NestedObjectMap[PagesProjectDeploymentConfigsPreviewAnalyticsEngineDatasetsDataSourceModel] `tfsdk:"analytics_engine_datasets" json:"analytics_engine_datasets,computed"`
	Browsers                customfield.Map[jsontypes.Normalized]                                                                   `tfsdk:"browsers" json:"browsers,computed"`
	CompatibilityDate       types.String                                                                                            `tfsdk:"compatibility_date" json:"compatibility_date,computed"`
	CompatibilityFlags      customfield.List[types.String]                                                                          `tfsdk:"compatibility_flags" json:"compatibility_flags,computed"`
	D1Databases             customfield.NestedObjectMap[PagesProjectDeploymentConfigsPreviewD1DatabasesDataSourceModel]             `tfsdk:"d1_databases" json:"d1_databases,computed"`
	DurableObjectNamespaces customfield.NestedObjectMap[PagesProjectDeploymentConfigsPreviewDurableObjectNamespacesDataSourceModel] `tfsdk:"durable_object_namespaces" json:"durable_object_namespaces,computed"`
	EnvVars                 customfield.NestedObjectMap[PagesProjectDeploymentConfigsPreviewEnvVarsDataSourceModel]                 `tfsdk:"env_vars" json:"env_vars,computed"`
	HyperdriveBindings      customfield.NestedObjectMap[PagesProjectDeploymentConfigsPreviewHyperdriveBindingsDataSourceModel]      `tfsdk:"hyperdrive_bindings" json:"hyperdrive_bindings,computed"`
	KVNamespaces            customfield.NestedObjectMap[PagesProjectDeploymentConfigsPreviewKVNamespacesDataSourceModel]            `tfsdk:"kv_namespaces" json:"kv_namespaces,computed"`
	MTLSCertificates        customfield.NestedObjectMap[PagesProjectDeploymentConfigsPreviewMTLSCertificatesDataSourceModel]        `tfsdk:"mtls_certificates" json:"mtls_certificates,computed"`
	Placement               customfield.NestedObject[PagesProjectDeploymentConfigsPreviewPlacementDataSourceModel]                  `tfsdk:"placement" json:"placement,computed"`
	QueueProducers          customfield.NestedObjectMap[PagesProjectDeploymentConfigsPreviewQueueProducersDataSourceModel]          `tfsdk:"queue_producers" json:"queue_producers,computed"`
	R2Buckets               customfield.NestedObjectMap[PagesProjectDeploymentConfigsPreviewR2BucketsDataSourceModel]               `tfsdk:"r2_buckets" json:"r2_buckets,computed"`
	Services                customfield.NestedObjectMap[PagesProjectDeploymentConfigsPreviewServicesDataSourceModel]                `tfsdk:"services" json:"services,computed"`
	VectorizeBindings       customfield.NestedObjectMap[PagesProjectDeploymentConfigsPreviewVectorizeBindingsDataSourceModel]       `tfsdk:"vectorize_bindings" json:"vectorize_bindings,computed"`
}

type PagesProjectDeploymentConfigsPreviewAIBindingsDataSourceModel struct {
	ProjectID types.String `tfsdk:"project_id" json:"project_id,computed"`
}

type PagesProjectDeploymentConfigsPreviewAnalyticsEngineDatasetsDataSourceModel struct {
	Dataset types.String `tfsdk:"dataset" json:"dataset,computed"`
}

type PagesProjectDeploymentConfigsPreviewD1DatabasesDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type PagesProjectDeploymentConfigsPreviewDurableObjectNamespacesDataSourceModel struct {
	NamespaceID types.String `tfsdk:"namespace_id" json:"namespace_id,computed"`
}

type PagesProjectDeploymentConfigsPreviewEnvVarsDataSourceModel struct {
	Value types.String `tfsdk:"value" json:"value,computed"`
	Type  types.String `tfsdk:"type" json:"type,computed"`
}

type PagesProjectDeploymentConfigsPreviewHyperdriveBindingsDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type PagesProjectDeploymentConfigsPreviewKVNamespacesDataSourceModel struct {
	NamespaceID types.String `tfsdk:"namespace_id" json:"namespace_id,computed"`
}

type PagesProjectDeploymentConfigsPreviewMTLSCertificatesDataSourceModel struct {
	CertificateID types.String `tfsdk:"certificate_id" json:"certificate_id,computed"`
}

type PagesProjectDeploymentConfigsPreviewPlacementDataSourceModel struct {
	Mode types.String `tfsdk:"mode" json:"mode,computed"`
}

type PagesProjectDeploymentConfigsPreviewQueueProducersDataSourceModel struct {
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type PagesProjectDeploymentConfigsPreviewR2BucketsDataSourceModel struct {
	Jurisdiction types.String `tfsdk:"jurisdiction" json:"jurisdiction,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type PagesProjectDeploymentConfigsPreviewServicesDataSourceModel struct {
	Entrypoint  types.String `tfsdk:"entrypoint" json:"entrypoint,computed"`
	Environment types.String `tfsdk:"environment" json:"environment,computed"`
	Service     types.String `tfsdk:"service" json:"service,computed"`
}

type PagesProjectDeploymentConfigsPreviewVectorizeBindingsDataSourceModel struct {
	IndexName types.String `tfsdk:"index_name" json:"index_name,computed"`
}

type PagesProjectDeploymentConfigsProductionDataSourceModel struct {
	AIBindings              customfield.NestedObjectMap[PagesProjectDeploymentConfigsProductionAIBindingsDataSourceModel]              `tfsdk:"ai_bindings" json:"ai_bindings,computed"`
	AnalyticsEngineDatasets customfield.NestedObjectMap[PagesProjectDeploymentConfigsProductionAnalyticsEngineDatasetsDataSourceModel] `tfsdk:"analytics_engine_datasets" json:"analytics_engine_datasets,computed"`
	Browsers                customfield.Map[jsontypes.Normalized]                                                                      `tfsdk:"browsers" json:"browsers,computed"`
	CompatibilityDate       types.String                                                                                               `tfsdk:"compatibility_date" json:"compatibility_date,computed"`
	CompatibilityFlags      customfield.List[types.String]                                                                             `tfsdk:"compatibility_flags" json:"compatibility_flags,computed"`
	D1Databases             customfield.NestedObjectMap[PagesProjectDeploymentConfigsProductionD1DatabasesDataSourceModel]             `tfsdk:"d1_databases" json:"d1_databases,computed"`
	DurableObjectNamespaces customfield.NestedObjectMap[PagesProjectDeploymentConfigsProductionDurableObjectNamespacesDataSourceModel] `tfsdk:"durable_object_namespaces" json:"durable_object_namespaces,computed"`
	EnvVars                 customfield.NestedObjectMap[PagesProjectDeploymentConfigsProductionEnvVarsDataSourceModel]                 `tfsdk:"env_vars" json:"env_vars,computed"`
	HyperdriveBindings      customfield.NestedObjectMap[PagesProjectDeploymentConfigsProductionHyperdriveBindingsDataSourceModel]      `tfsdk:"hyperdrive_bindings" json:"hyperdrive_bindings,computed"`
	KVNamespaces            customfield.NestedObjectMap[PagesProjectDeploymentConfigsProductionKVNamespacesDataSourceModel]            `tfsdk:"kv_namespaces" json:"kv_namespaces,computed"`
	MTLSCertificates        customfield.NestedObjectMap[PagesProjectDeploymentConfigsProductionMTLSCertificatesDataSourceModel]        `tfsdk:"mtls_certificates" json:"mtls_certificates,computed"`
	Placement               customfield.NestedObject[PagesProjectDeploymentConfigsProductionPlacementDataSourceModel]                  `tfsdk:"placement" json:"placement,computed"`
	QueueProducers          customfield.NestedObjectMap[PagesProjectDeploymentConfigsProductionQueueProducersDataSourceModel]          `tfsdk:"queue_producers" json:"queue_producers,computed"`
	R2Buckets               customfield.NestedObjectMap[PagesProjectDeploymentConfigsProductionR2BucketsDataSourceModel]               `tfsdk:"r2_buckets" json:"r2_buckets,computed"`
	Services                customfield.NestedObjectMap[PagesProjectDeploymentConfigsProductionServicesDataSourceModel]                `tfsdk:"services" json:"services,computed"`
	VectorizeBindings       customfield.NestedObjectMap[PagesProjectDeploymentConfigsProductionVectorizeBindingsDataSourceModel]       `tfsdk:"vectorize_bindings" json:"vectorize_bindings,computed"`
}

type PagesProjectDeploymentConfigsProductionAIBindingsDataSourceModel struct {
	ProjectID types.String `tfsdk:"project_id" json:"project_id,computed"`
}

type PagesProjectDeploymentConfigsProductionAnalyticsEngineDatasetsDataSourceModel struct {
	Dataset types.String `tfsdk:"dataset" json:"dataset,computed"`
}

type PagesProjectDeploymentConfigsProductionD1DatabasesDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type PagesProjectDeploymentConfigsProductionDurableObjectNamespacesDataSourceModel struct {
	NamespaceID types.String `tfsdk:"namespace_id" json:"namespace_id,computed"`
}

type PagesProjectDeploymentConfigsProductionEnvVarsDataSourceModel struct {
	Value types.String `tfsdk:"value" json:"value,computed"`
	Type  types.String `tfsdk:"type" json:"type,computed"`
}

type PagesProjectDeploymentConfigsProductionHyperdriveBindingsDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type PagesProjectDeploymentConfigsProductionKVNamespacesDataSourceModel struct {
	NamespaceID types.String `tfsdk:"namespace_id" json:"namespace_id,computed"`
}

type PagesProjectDeploymentConfigsProductionMTLSCertificatesDataSourceModel struct {
	CertificateID types.String `tfsdk:"certificate_id" json:"certificate_id,computed"`
}

type PagesProjectDeploymentConfigsProductionPlacementDataSourceModel struct {
	Mode types.String `tfsdk:"mode" json:"mode,computed"`
}

type PagesProjectDeploymentConfigsProductionQueueProducersDataSourceModel struct {
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type PagesProjectDeploymentConfigsProductionR2BucketsDataSourceModel struct {
	Jurisdiction types.String `tfsdk:"jurisdiction" json:"jurisdiction,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type PagesProjectDeploymentConfigsProductionServicesDataSourceModel struct {
	Entrypoint  types.String `tfsdk:"entrypoint" json:"entrypoint,computed"`
	Environment types.String `tfsdk:"environment" json:"environment,computed"`
	Service     types.String `tfsdk:"service" json:"service,computed"`
}

type PagesProjectDeploymentConfigsProductionVectorizeBindingsDataSourceModel struct {
	IndexName types.String `tfsdk:"index_name" json:"index_name,computed"`
}

type PagesProjectDeploymentTriggerDataSourceModel struct {
	Metadata customfield.NestedObject[PagesProjectDeploymentTriggerMetadataDataSourceModel] `tfsdk:"metadata" json:"metadata,computed"`
	Type     types.String                                                                   `tfsdk:"type" json:"type,computed"`
}

type PagesProjectDeploymentTriggerMetadataDataSourceModel struct {
	Branch        types.String `tfsdk:"branch" json:"branch,computed"`
	CommitHash    types.String `tfsdk:"commit_hash" json:"commit_hash,computed"`
	CommitMessage types.String `tfsdk:"commit_message" json:"commit_message,computed"`
}

type PagesProjectEnvVarsDataSourceModel struct {
	Value types.String `tfsdk:"value" json:"value,computed"`
	Type  types.String `tfsdk:"type" json:"type,computed"`
}

type PagesProjectLatestDeploymentDataSourceModel struct {
	ID                types.String                                                                           `tfsdk:"id" json:"id,computed"`
	Aliases           customfield.List[types.String]                                                         `tfsdk:"aliases" json:"aliases,computed"`
	BuildConfig       customfield.NestedObject[PagesProjectLatestDeploymentBuildConfigDataSourceModel]       `tfsdk:"build_config" json:"build_config,computed"`
	CreatedOn         timetypes.RFC3339                                                                      `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	DeploymentTrigger customfield.NestedObject[PagesProjectLatestDeploymentDeploymentTriggerDataSourceModel] `tfsdk:"deployment_trigger" json:"deployment_trigger,computed"`
	EnvVars           customfield.NestedObjectMap[PagesProjectLatestDeploymentEnvVarsDataSourceModel]        `tfsdk:"env_vars" json:"env_vars,computed"`
	Environment       types.String                                                                           `tfsdk:"environment" json:"environment,computed"`
	IsSkipped         types.Bool                                                                             `tfsdk:"is_skipped" json:"is_skipped,computed"`
	LatestStage       customfield.NestedObject[PagesProjectLatestDeploymentLatestStageDataSourceModel]       `tfsdk:"latest_stage" json:"latest_stage,computed"`
	ModifiedOn        timetypes.RFC3339                                                                      `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	ProjectID         types.String                                                                           `tfsdk:"project_id" json:"project_id,computed"`
	ProjectName       types.String                                                                           `tfsdk:"project_name" json:"project_name,computed"`
	ShortID           types.String                                                                           `tfsdk:"short_id" json:"short_id,computed"`
	Source            customfield.NestedObject[PagesProjectLatestDeploymentSourceDataSourceModel]            `tfsdk:"source" json:"source,computed"`
	Stages            customfield.NestedObjectList[PagesProjectLatestDeploymentStagesDataSourceModel]        `tfsdk:"stages" json:"stages,computed"`
	URL               types.String                                                                           `tfsdk:"url" json:"url,computed"`
}

type PagesProjectLatestDeploymentBuildConfigDataSourceModel struct {
	BuildCaching      types.Bool   `tfsdk:"build_caching" json:"build_caching,computed"`
	BuildCommand      types.String `tfsdk:"build_command" json:"build_command,computed"`
	DestinationDir    types.String `tfsdk:"destination_dir" json:"destination_dir,computed"`
	RootDir           types.String `tfsdk:"root_dir" json:"root_dir,computed"`
	WebAnalyticsTag   types.String `tfsdk:"web_analytics_tag" json:"web_analytics_tag,computed"`
	WebAnalyticsToken types.String `tfsdk:"web_analytics_token" json:"web_analytics_token,computed"`
}

type PagesProjectLatestDeploymentDeploymentTriggerDataSourceModel struct {
	Metadata customfield.NestedObject[PagesProjectLatestDeploymentDeploymentTriggerMetadataDataSourceModel] `tfsdk:"metadata" json:"metadata,computed"`
	Type     types.String                                                                                   `tfsdk:"type" json:"type,computed"`
}

type PagesProjectLatestDeploymentDeploymentTriggerMetadataDataSourceModel struct {
	Branch        types.String `tfsdk:"branch" json:"branch,computed"`
	CommitHash    types.String `tfsdk:"commit_hash" json:"commit_hash,computed"`
	CommitMessage types.String `tfsdk:"commit_message" json:"commit_message,computed"`
}

type PagesProjectLatestDeploymentEnvVarsDataSourceModel struct {
	Value types.String `tfsdk:"value" json:"value,computed"`
	Type  types.String `tfsdk:"type" json:"type,computed"`
}

type PagesProjectLatestDeploymentLatestStageDataSourceModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed" format:"date-time"`
	Name      types.String      `tfsdk:"name" json:"name,computed"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed" format:"date-time"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
}

type PagesProjectLatestDeploymentSourceDataSourceModel struct {
	Config customfield.NestedObject[PagesProjectLatestDeploymentSourceConfigDataSourceModel] `tfsdk:"config" json:"config,computed"`
	Type   types.String                                                                      `tfsdk:"type" json:"type,computed"`
}

type PagesProjectLatestDeploymentSourceConfigDataSourceModel struct {
	DeploymentsEnabled           types.Bool                     `tfsdk:"deployments_enabled" json:"deployments_enabled,computed"`
	Owner                        types.String                   `tfsdk:"owner" json:"owner,computed"`
	PathExcludes                 customfield.List[types.String] `tfsdk:"path_excludes" json:"path_excludes,computed"`
	PathIncludes                 customfield.List[types.String] `tfsdk:"path_includes" json:"path_includes,computed"`
	PrCommentsEnabled            types.Bool                     `tfsdk:"pr_comments_enabled" json:"pr_comments_enabled,computed"`
	PreviewBranchExcludes        customfield.List[types.String] `tfsdk:"preview_branch_excludes" json:"preview_branch_excludes,computed"`
	PreviewBranchIncludes        customfield.List[types.String] `tfsdk:"preview_branch_includes" json:"preview_branch_includes,computed"`
	PreviewDeploymentSetting     types.String                   `tfsdk:"preview_deployment_setting" json:"preview_deployment_setting,computed"`
	ProductionBranch             types.String                   `tfsdk:"production_branch" json:"production_branch,computed"`
	ProductionDeploymentsEnabled types.Bool                     `tfsdk:"production_deployments_enabled" json:"production_deployments_enabled,computed"`
	RepoName                     types.String                   `tfsdk:"repo_name" json:"repo_name,computed"`
}

type PagesProjectLatestDeploymentStagesDataSourceModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed" format:"date-time"`
	Name      types.String      `tfsdk:"name" json:"name,computed"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed" format:"date-time"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
}

type PagesProjectLatestStageDataSourceModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed" format:"date-time"`
	Name      types.String      `tfsdk:"name" json:"name,computed"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed" format:"date-time"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
}

type PagesProjectStagesDataSourceModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed" format:"date-time"`
	Name      types.String      `tfsdk:"name" json:"name,computed"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed" format:"date-time"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
}

type PagesProjectBuildConfigDataSourceModel struct {
	BuildCaching      types.Bool   `tfsdk:"build_caching" json:"build_caching,computed"`
	BuildCommand      types.String `tfsdk:"build_command" json:"build_command,computed"`
	DestinationDir    types.String `tfsdk:"destination_dir" json:"destination_dir,computed"`
	RootDir           types.String `tfsdk:"root_dir" json:"root_dir,computed"`
	WebAnalyticsTag   types.String `tfsdk:"web_analytics_tag" json:"web_analytics_tag,computed"`
	WebAnalyticsToken types.String `tfsdk:"web_analytics_token" json:"web_analytics_token,computed"`
}

type PagesProjectSourceDataSourceModel struct {
	Config customfield.NestedObject[PagesProjectSourceConfigDataSourceModel] `tfsdk:"config" json:"config,computed"`
	Type   types.String                                                      `tfsdk:"type" json:"type,computed"`
}

type PagesProjectSourceConfigDataSourceModel struct {
	DeploymentsEnabled           types.Bool                     `tfsdk:"deployments_enabled" json:"deployments_enabled,computed"`
	Owner                        types.String                   `tfsdk:"owner" json:"owner,computed"`
	PathExcludes                 customfield.List[types.String] `tfsdk:"path_excludes" json:"path_excludes,computed"`
	PathIncludes                 customfield.List[types.String] `tfsdk:"path_includes" json:"path_includes,computed"`
	PrCommentsEnabled            types.Bool                     `tfsdk:"pr_comments_enabled" json:"pr_comments_enabled,computed"`
	PreviewBranchExcludes        customfield.List[types.String] `tfsdk:"preview_branch_excludes" json:"preview_branch_excludes,computed"`
	PreviewBranchIncludes        customfield.List[types.String] `tfsdk:"preview_branch_includes" json:"preview_branch_includes,computed"`
	PreviewDeploymentSetting     types.String                   `tfsdk:"preview_deployment_setting" json:"preview_deployment_setting,computed"`
	ProductionBranch             types.String                   `tfsdk:"production_branch" json:"production_branch,computed"`
	ProductionDeploymentsEnabled types.Bool                     `tfsdk:"production_deployments_enabled" json:"production_deployments_enabled,computed"`
	RepoName                     types.String                   `tfsdk:"repo_name" json:"repo_name,computed"`
}

type PagesProjectFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
}
