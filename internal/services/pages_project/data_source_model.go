// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_project

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PagesProjectResultDataSourceEnvelope struct {
	Result PagesProjectDataSourceModel `json:"result,computed"`
}

type PagesProjectResultListDataSourceEnvelope struct {
	Result *[]*PagesProjectDataSourceModel `json:"result,computed"`
}

type PagesProjectDataSourceModel struct {
	AccountID           types.String                                                `tfsdk:"account_id" path:"account_id"`
	ProjectName         types.String                                                `tfsdk:"project_name" path:"project_name"`
	Environment         types.String                                                `tfsdk:"environment" json:"environment"`
	IsSkipped           types.Bool                                                  `tfsdk:"is_skipped" json:"is_skipped"`
	ModifiedOn          timetypes.RFC3339                                           `tfsdk:"modified_on" json:"modified_on"`
	Name                types.String                                                `tfsdk:"name" json:"name"`
	ProductionBranch    types.String                                                `tfsdk:"production_branch" json:"production_branch"`
	ProjectID           types.String                                                `tfsdk:"project_id" json:"project_id"`
	ShortID             types.String                                                `tfsdk:"short_id" json:"short_id"`
	Subdomain           types.String                                                `tfsdk:"subdomain" json:"subdomain"`
	URL                 types.String                                                `tfsdk:"url" json:"url"`
	Aliases             *[]types.String                                             `tfsdk:"aliases" json:"aliases"`
	Domains             *[]types.String                                             `tfsdk:"domains" json:"domains"`
	EnvVars             map[string]jsontypes.Normalized                             `tfsdk:"env_vars" json:"env_vars"`
	CanonicalDeployment *PagesProjectCanonicalDeploymentDataSourceModel             `tfsdk:"canonical_deployment" json:"canonical_deployment"`
	DeploymentConfigs   *PagesProjectDeploymentConfigsDataSourceModel               `tfsdk:"deployment_configs" json:"deployment_configs"`
	DeploymentTrigger   *PagesProjectDeploymentTriggerDataSourceModel               `tfsdk:"deployment_trigger" json:"deployment_trigger"`
	LatestDeployment    *PagesProjectLatestDeploymentDataSourceModel                `tfsdk:"latest_deployment" json:"latest_deployment"`
	LatestStage         *PagesProjectLatestStageDataSourceModel                     `tfsdk:"latest_stage" json:"latest_stage"`
	Stages              *[]*PagesProjectStagesDataSourceModel                       `tfsdk:"stages" json:"stages"`
	CreatedOn           timetypes.RFC3339                                           `tfsdk:"created_on" json:"created_on,computed"`
	ID                  types.String                                                `tfsdk:"id" json:"id,computed"`
	Source              customfield.NestedObject[PagesProjectSourceDataSourceModel] `tfsdk:"source" json:"source,computed"`
	BuildConfig         *PagesProjectBuildConfigDataSourceModel                     `tfsdk:"build_config" json:"build_config"`
	Filter              *PagesProjectFindOneByDataSourceModel                       `tfsdk:"filter"`
}

type PagesProjectCanonicalDeploymentDataSourceModel struct {
	ID                types.String                                                                              `tfsdk:"id" json:"id,computed"`
	Aliases           *[]types.String                                                                           `tfsdk:"aliases" json:"aliases,computed"`
	BuildConfig       *PagesProjectCanonicalDeploymentBuildConfigDataSourceModel                                `tfsdk:"build_config" json:"build_config"`
	CreatedOn         timetypes.RFC3339                                                                         `tfsdk:"created_on" json:"created_on,computed"`
	DeploymentTrigger customfield.NestedObject[PagesProjectCanonicalDeploymentDeploymentTriggerDataSourceModel] `tfsdk:"deployment_trigger" json:"deployment_trigger,computed"`
	EnvVars           map[string]jsontypes.Normalized                                                           `tfsdk:"env_vars" json:"env_vars,computed"`
	Environment       types.String                                                                              `tfsdk:"environment" json:"environment,computed"`
	IsSkipped         types.Bool                                                                                `tfsdk:"is_skipped" json:"is_skipped,computed"`
	LatestStage       customfield.NestedObject[PagesProjectCanonicalDeploymentLatestStageDataSourceModel]       `tfsdk:"latest_stage" json:"latest_stage,computed"`
	ModifiedOn        timetypes.RFC3339                                                                         `tfsdk:"modified_on" json:"modified_on,computed"`
	ProjectID         types.String                                                                              `tfsdk:"project_id" json:"project_id,computed"`
	ProjectName       types.String                                                                              `tfsdk:"project_name" json:"project_name,computed"`
	ShortID           types.String                                                                              `tfsdk:"short_id" json:"short_id,computed"`
	Source            customfield.NestedObject[PagesProjectCanonicalDeploymentSourceDataSourceModel]            `tfsdk:"source" json:"source,computed"`
	Stages            *[]*PagesProjectCanonicalDeploymentStagesDataSourceModel                                  `tfsdk:"stages" json:"stages,computed"`
	URL               types.String                                                                              `tfsdk:"url" json:"url,computed"`
}

type PagesProjectCanonicalDeploymentBuildConfigDataSourceModel struct {
	BuildCaching      types.Bool   `tfsdk:"build_caching" json:"build_caching"`
	BuildCommand      types.String `tfsdk:"build_command" json:"build_command"`
	DestinationDir    types.String `tfsdk:"destination_dir" json:"destination_dir"`
	RootDir           types.String `tfsdk:"root_dir" json:"root_dir"`
	WebAnalyticsTag   types.String `tfsdk:"web_analytics_tag" json:"web_analytics_tag"`
	WebAnalyticsToken types.String `tfsdk:"web_analytics_token" json:"web_analytics_token"`
}

type PagesProjectCanonicalDeploymentDeploymentTriggerDataSourceModel struct {
	Metadata *PagesProjectCanonicalDeploymentDeploymentTriggerMetadataDataSourceModel `tfsdk:"metadata" json:"metadata"`
	Type     types.String                                                             `tfsdk:"type" json:"type,computed"`
}

type PagesProjectCanonicalDeploymentDeploymentTriggerMetadataDataSourceModel struct {
	Branch        types.String `tfsdk:"branch" json:"branch,computed"`
	CommitHash    types.String `tfsdk:"commit_hash" json:"commit_hash,computed"`
	CommitMessage types.String `tfsdk:"commit_message" json:"commit_message,computed"`
}

type PagesProjectCanonicalDeploymentLatestStageDataSourceModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed"`
	Name      types.String      `tfsdk:"name" json:"name"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
}

type PagesProjectCanonicalDeploymentSourceDataSourceModel struct {
	Config *PagesProjectCanonicalDeploymentSourceConfigDataSourceModel `tfsdk:"config" json:"config"`
	Type   types.String                                                `tfsdk:"type" json:"type"`
}

type PagesProjectCanonicalDeploymentSourceConfigDataSourceModel struct {
	DeploymentsEnabled           types.Bool      `tfsdk:"deployments_enabled" json:"deployments_enabled"`
	Owner                        types.String    `tfsdk:"owner" json:"owner"`
	PathExcludes                 *[]types.String `tfsdk:"path_excludes" json:"path_excludes"`
	PathIncludes                 *[]types.String `tfsdk:"path_includes" json:"path_includes"`
	PrCommentsEnabled            types.Bool      `tfsdk:"pr_comments_enabled" json:"pr_comments_enabled"`
	PreviewBranchExcludes        *[]types.String `tfsdk:"preview_branch_excludes" json:"preview_branch_excludes"`
	PreviewBranchIncludes        *[]types.String `tfsdk:"preview_branch_includes" json:"preview_branch_includes"`
	PreviewDeploymentSetting     types.String    `tfsdk:"preview_deployment_setting" json:"preview_deployment_setting"`
	ProductionBranch             types.String    `tfsdk:"production_branch" json:"production_branch"`
	ProductionDeploymentsEnabled types.Bool      `tfsdk:"production_deployments_enabled" json:"production_deployments_enabled"`
	RepoName                     types.String    `tfsdk:"repo_name" json:"repo_name"`
}

type PagesProjectCanonicalDeploymentStagesDataSourceModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed"`
	Name      types.String      `tfsdk:"name" json:"name"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
}

type PagesProjectDeploymentConfigsDataSourceModel struct {
	Preview    *PagesProjectDeploymentConfigsPreviewDataSourceModel    `tfsdk:"preview" json:"preview"`
	Production *PagesProjectDeploymentConfigsProductionDataSourceModel `tfsdk:"production" json:"production"`
}

type PagesProjectDeploymentConfigsPreviewDataSourceModel struct {
	AIBindings              map[string]jsontypes.Normalized                               `tfsdk:"ai_bindings" json:"ai_bindings"`
	AnalyticsEngineDatasets map[string]jsontypes.Normalized                               `tfsdk:"analytics_engine_datasets" json:"analytics_engine_datasets"`
	Browsers                map[string]jsontypes.Normalized                               `tfsdk:"browsers" json:"browsers"`
	CompatibilityDate       types.String                                                  `tfsdk:"compatibility_date" json:"compatibility_date"`
	CompatibilityFlags      *[]types.String                                               `tfsdk:"compatibility_flags" json:"compatibility_flags"`
	D1Databases             map[string]jsontypes.Normalized                               `tfsdk:"d1_databases" json:"d1_databases"`
	DurableObjectNamespaces map[string]jsontypes.Normalized                               `tfsdk:"durable_object_namespaces" json:"durable_object_namespaces"`
	EnvVars                 map[string]jsontypes.Normalized                               `tfsdk:"env_vars" json:"env_vars"`
	HyperdriveBindings      map[string]jsontypes.Normalized                               `tfsdk:"hyperdrive_bindings" json:"hyperdrive_bindings"`
	KVNamespaces            map[string]jsontypes.Normalized                               `tfsdk:"kv_namespaces" json:"kv_namespaces"`
	MTLSCertificates        map[string]jsontypes.Normalized                               `tfsdk:"mtls_certificates" json:"mtls_certificates"`
	Placement               *PagesProjectDeploymentConfigsPreviewPlacementDataSourceModel `tfsdk:"placement" json:"placement"`
	QueueProducers          map[string]jsontypes.Normalized                               `tfsdk:"queue_producers" json:"queue_producers"`
	R2Buckets               map[string]jsontypes.Normalized                               `tfsdk:"r2_buckets" json:"r2_buckets"`
	Services                map[string]jsontypes.Normalized                               `tfsdk:"services" json:"services"`
	VectorizeBindings       map[string]jsontypes.Normalized                               `tfsdk:"vectorize_bindings" json:"vectorize_bindings"`
}

type PagesProjectDeploymentConfigsPreviewPlacementDataSourceModel struct {
	Mode types.String `tfsdk:"mode" json:"mode"`
}

type PagesProjectDeploymentConfigsProductionDataSourceModel struct {
	AIBindings              map[string]jsontypes.Normalized                                  `tfsdk:"ai_bindings" json:"ai_bindings"`
	AnalyticsEngineDatasets map[string]jsontypes.Normalized                                  `tfsdk:"analytics_engine_datasets" json:"analytics_engine_datasets"`
	Browsers                map[string]jsontypes.Normalized                                  `tfsdk:"browsers" json:"browsers"`
	CompatibilityDate       types.String                                                     `tfsdk:"compatibility_date" json:"compatibility_date"`
	CompatibilityFlags      *[]types.String                                                  `tfsdk:"compatibility_flags" json:"compatibility_flags"`
	D1Databases             map[string]jsontypes.Normalized                                  `tfsdk:"d1_databases" json:"d1_databases"`
	DurableObjectNamespaces map[string]jsontypes.Normalized                                  `tfsdk:"durable_object_namespaces" json:"durable_object_namespaces"`
	EnvVars                 map[string]jsontypes.Normalized                                  `tfsdk:"env_vars" json:"env_vars"`
	HyperdriveBindings      map[string]jsontypes.Normalized                                  `tfsdk:"hyperdrive_bindings" json:"hyperdrive_bindings"`
	KVNamespaces            map[string]jsontypes.Normalized                                  `tfsdk:"kv_namespaces" json:"kv_namespaces"`
	MTLSCertificates        map[string]jsontypes.Normalized                                  `tfsdk:"mtls_certificates" json:"mtls_certificates"`
	Placement               *PagesProjectDeploymentConfigsProductionPlacementDataSourceModel `tfsdk:"placement" json:"placement"`
	QueueProducers          map[string]jsontypes.Normalized                                  `tfsdk:"queue_producers" json:"queue_producers"`
	R2Buckets               map[string]jsontypes.Normalized                                  `tfsdk:"r2_buckets" json:"r2_buckets"`
	Services                map[string]jsontypes.Normalized                                  `tfsdk:"services" json:"services"`
	VectorizeBindings       map[string]jsontypes.Normalized                                  `tfsdk:"vectorize_bindings" json:"vectorize_bindings"`
}

type PagesProjectDeploymentConfigsProductionPlacementDataSourceModel struct {
	Mode types.String `tfsdk:"mode" json:"mode"`
}

type PagesProjectDeploymentTriggerDataSourceModel struct {
	Metadata *PagesProjectDeploymentTriggerMetadataDataSourceModel `tfsdk:"metadata" json:"metadata"`
	Type     types.String                                          `tfsdk:"type" json:"type,computed"`
}

type PagesProjectDeploymentTriggerMetadataDataSourceModel struct {
	Branch        types.String `tfsdk:"branch" json:"branch,computed"`
	CommitHash    types.String `tfsdk:"commit_hash" json:"commit_hash,computed"`
	CommitMessage types.String `tfsdk:"commit_message" json:"commit_message,computed"`
}

type PagesProjectLatestDeploymentDataSourceModel struct {
	ID                types.String                                                                           `tfsdk:"id" json:"id,computed"`
	Aliases           *[]types.String                                                                        `tfsdk:"aliases" json:"aliases,computed"`
	BuildConfig       *PagesProjectLatestDeploymentBuildConfigDataSourceModel                                `tfsdk:"build_config" json:"build_config"`
	CreatedOn         timetypes.RFC3339                                                                      `tfsdk:"created_on" json:"created_on,computed"`
	DeploymentTrigger customfield.NestedObject[PagesProjectLatestDeploymentDeploymentTriggerDataSourceModel] `tfsdk:"deployment_trigger" json:"deployment_trigger,computed"`
	EnvVars           map[string]jsontypes.Normalized                                                        `tfsdk:"env_vars" json:"env_vars,computed"`
	Environment       types.String                                                                           `tfsdk:"environment" json:"environment,computed"`
	IsSkipped         types.Bool                                                                             `tfsdk:"is_skipped" json:"is_skipped,computed"`
	LatestStage       customfield.NestedObject[PagesProjectLatestDeploymentLatestStageDataSourceModel]       `tfsdk:"latest_stage" json:"latest_stage,computed"`
	ModifiedOn        timetypes.RFC3339                                                                      `tfsdk:"modified_on" json:"modified_on,computed"`
	ProjectID         types.String                                                                           `tfsdk:"project_id" json:"project_id,computed"`
	ProjectName       types.String                                                                           `tfsdk:"project_name" json:"project_name,computed"`
	ShortID           types.String                                                                           `tfsdk:"short_id" json:"short_id,computed"`
	Source            customfield.NestedObject[PagesProjectLatestDeploymentSourceDataSourceModel]            `tfsdk:"source" json:"source,computed"`
	Stages            *[]*PagesProjectLatestDeploymentStagesDataSourceModel                                  `tfsdk:"stages" json:"stages,computed"`
	URL               types.String                                                                           `tfsdk:"url" json:"url,computed"`
}

type PagesProjectLatestDeploymentBuildConfigDataSourceModel struct {
	BuildCaching      types.Bool   `tfsdk:"build_caching" json:"build_caching"`
	BuildCommand      types.String `tfsdk:"build_command" json:"build_command"`
	DestinationDir    types.String `tfsdk:"destination_dir" json:"destination_dir"`
	RootDir           types.String `tfsdk:"root_dir" json:"root_dir"`
	WebAnalyticsTag   types.String `tfsdk:"web_analytics_tag" json:"web_analytics_tag"`
	WebAnalyticsToken types.String `tfsdk:"web_analytics_token" json:"web_analytics_token"`
}

type PagesProjectLatestDeploymentDeploymentTriggerDataSourceModel struct {
	Metadata *PagesProjectLatestDeploymentDeploymentTriggerMetadataDataSourceModel `tfsdk:"metadata" json:"metadata"`
	Type     types.String                                                          `tfsdk:"type" json:"type,computed"`
}

type PagesProjectLatestDeploymentDeploymentTriggerMetadataDataSourceModel struct {
	Branch        types.String `tfsdk:"branch" json:"branch,computed"`
	CommitHash    types.String `tfsdk:"commit_hash" json:"commit_hash,computed"`
	CommitMessage types.String `tfsdk:"commit_message" json:"commit_message,computed"`
}

type PagesProjectLatestDeploymentLatestStageDataSourceModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed"`
	Name      types.String      `tfsdk:"name" json:"name"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
}

type PagesProjectLatestDeploymentSourceDataSourceModel struct {
	Config *PagesProjectLatestDeploymentSourceConfigDataSourceModel `tfsdk:"config" json:"config"`
	Type   types.String                                             `tfsdk:"type" json:"type"`
}

type PagesProjectLatestDeploymentSourceConfigDataSourceModel struct {
	DeploymentsEnabled           types.Bool      `tfsdk:"deployments_enabled" json:"deployments_enabled"`
	Owner                        types.String    `tfsdk:"owner" json:"owner"`
	PathExcludes                 *[]types.String `tfsdk:"path_excludes" json:"path_excludes"`
	PathIncludes                 *[]types.String `tfsdk:"path_includes" json:"path_includes"`
	PrCommentsEnabled            types.Bool      `tfsdk:"pr_comments_enabled" json:"pr_comments_enabled"`
	PreviewBranchExcludes        *[]types.String `tfsdk:"preview_branch_excludes" json:"preview_branch_excludes"`
	PreviewBranchIncludes        *[]types.String `tfsdk:"preview_branch_includes" json:"preview_branch_includes"`
	PreviewDeploymentSetting     types.String    `tfsdk:"preview_deployment_setting" json:"preview_deployment_setting"`
	ProductionBranch             types.String    `tfsdk:"production_branch" json:"production_branch"`
	ProductionDeploymentsEnabled types.Bool      `tfsdk:"production_deployments_enabled" json:"production_deployments_enabled"`
	RepoName                     types.String    `tfsdk:"repo_name" json:"repo_name"`
}

type PagesProjectLatestDeploymentStagesDataSourceModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed"`
	Name      types.String      `tfsdk:"name" json:"name"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
}

type PagesProjectLatestStageDataSourceModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed"`
	Name      types.String      `tfsdk:"name" json:"name"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
}

type PagesProjectStagesDataSourceModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed"`
	Name      types.String      `tfsdk:"name" json:"name"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
}

type PagesProjectSourceDataSourceModel struct {
	Config *PagesProjectSourceConfigDataSourceModel `tfsdk:"config" json:"config"`
	Type   types.String                             `tfsdk:"type" json:"type"`
}

type PagesProjectSourceConfigDataSourceModel struct {
	DeploymentsEnabled           types.Bool      `tfsdk:"deployments_enabled" json:"deployments_enabled"`
	Owner                        types.String    `tfsdk:"owner" json:"owner"`
	PathExcludes                 *[]types.String `tfsdk:"path_excludes" json:"path_excludes"`
	PathIncludes                 *[]types.String `tfsdk:"path_includes" json:"path_includes"`
	PrCommentsEnabled            types.Bool      `tfsdk:"pr_comments_enabled" json:"pr_comments_enabled"`
	PreviewBranchExcludes        *[]types.String `tfsdk:"preview_branch_excludes" json:"preview_branch_excludes"`
	PreviewBranchIncludes        *[]types.String `tfsdk:"preview_branch_includes" json:"preview_branch_includes"`
	PreviewDeploymentSetting     types.String    `tfsdk:"preview_deployment_setting" json:"preview_deployment_setting"`
	ProductionBranch             types.String    `tfsdk:"production_branch" json:"production_branch"`
	ProductionDeploymentsEnabled types.Bool      `tfsdk:"production_deployments_enabled" json:"production_deployments_enabled"`
	RepoName                     types.String    `tfsdk:"repo_name" json:"repo_name"`
}

type PagesProjectBuildConfigDataSourceModel struct {
	BuildCaching      types.Bool   `tfsdk:"build_caching" json:"build_caching"`
	BuildCommand      types.String `tfsdk:"build_command" json:"build_command"`
	DestinationDir    types.String `tfsdk:"destination_dir" json:"destination_dir"`
	RootDir           types.String `tfsdk:"root_dir" json:"root_dir"`
	WebAnalyticsTag   types.String `tfsdk:"web_analytics_tag" json:"web_analytics_tag"`
	WebAnalyticsToken types.String `tfsdk:"web_analytics_token" json:"web_analytics_token"`
}

type PagesProjectFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
