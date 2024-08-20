// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_project

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PagesProjectResultEnvelope struct {
	Result PagesProjectModel `json:"result,computed"`
}

type PagesProjectModel struct {
	ID                  types.String                                                   `tfsdk:"id" json:"-,computed"`
	Name                types.String                                                   `tfsdk:"name" json:"name"`
	AccountID           types.String                                                   `tfsdk:"account_id" path:"account_id"`
	ProductionBranch    types.String                                                   `tfsdk:"production_branch" json:"production_branch"`
	BuildConfig         *PagesProjectBuildConfigModel                                  `tfsdk:"build_config" json:"build_config"`
	DeploymentConfigs   *PagesProjectDeploymentConfigsModel                            `tfsdk:"deployment_configs" json:"deployment_configs"`
	CreatedOn           timetypes.RFC3339                                              `tfsdk:"created_on" json:"created_on,computed"`
	Subdomain           types.String                                                   `tfsdk:"subdomain" json:"subdomain,computed"`
	Domains             *[]types.String                                                `tfsdk:"domains" json:"domains,computed"`
	CanonicalDeployment customfield.NestedObject[PagesProjectCanonicalDeploymentModel] `tfsdk:"canonical_deployment" json:"canonical_deployment,computed"`
	LatestDeployment    customfield.NestedObject[PagesProjectLatestDeploymentModel]    `tfsdk:"latest_deployment" json:"latest_deployment,computed"`
	Source              customfield.NestedObject[PagesProjectSourceModel]              `tfsdk:"source" json:"source,computed"`
}

type PagesProjectBuildConfigModel struct {
	BuildCaching      types.Bool   `tfsdk:"build_caching" json:"build_caching"`
	BuildCommand      types.String `tfsdk:"build_command" json:"build_command"`
	DestinationDir    types.String `tfsdk:"destination_dir" json:"destination_dir"`
	RootDir           types.String `tfsdk:"root_dir" json:"root_dir"`
	WebAnalyticsTag   types.String `tfsdk:"web_analytics_tag" json:"web_analytics_tag"`
	WebAnalyticsToken types.String `tfsdk:"web_analytics_token" json:"web_analytics_token"`
}

type PagesProjectDeploymentConfigsModel struct {
	Preview    *PagesProjectDeploymentConfigsPreviewModel    `tfsdk:"preview" json:"preview"`
	Production *PagesProjectDeploymentConfigsProductionModel `tfsdk:"production" json:"production"`
}

type PagesProjectDeploymentConfigsPreviewModel struct {
	AIBindings              map[string]jsontypes.Normalized                     `tfsdk:"ai_bindings" json:"ai_bindings"`
	AnalyticsEngineDatasets map[string]jsontypes.Normalized                     `tfsdk:"analytics_engine_datasets" json:"analytics_engine_datasets"`
	Browsers                map[string]jsontypes.Normalized                     `tfsdk:"browsers" json:"browsers"`
	CompatibilityDate       types.String                                        `tfsdk:"compatibility_date" json:"compatibility_date"`
	CompatibilityFlags      *[]types.String                                     `tfsdk:"compatibility_flags" json:"compatibility_flags"`
	D1Databases             map[string]jsontypes.Normalized                     `tfsdk:"d1_databases" json:"d1_databases"`
	DurableObjectNamespaces map[string]jsontypes.Normalized                     `tfsdk:"durable_object_namespaces" json:"durable_object_namespaces"`
	EnvVars                 map[string]jsontypes.Normalized                     `tfsdk:"env_vars" json:"env_vars"`
	HyperdriveBindings      map[string]jsontypes.Normalized                     `tfsdk:"hyperdrive_bindings" json:"hyperdrive_bindings"`
	KVNamespaces            map[string]jsontypes.Normalized                     `tfsdk:"kv_namespaces" json:"kv_namespaces"`
	MTLSCertificates        map[string]jsontypes.Normalized                     `tfsdk:"mtls_certificates" json:"mtls_certificates"`
	Placement               *PagesProjectDeploymentConfigsPreviewPlacementModel `tfsdk:"placement" json:"placement"`
	QueueProducers          map[string]jsontypes.Normalized                     `tfsdk:"queue_producers" json:"queue_producers"`
	R2Buckets               map[string]jsontypes.Normalized                     `tfsdk:"r2_buckets" json:"r2_buckets"`
	Services                map[string]jsontypes.Normalized                     `tfsdk:"services" json:"services"`
	VectorizeBindings       map[string]jsontypes.Normalized                     `tfsdk:"vectorize_bindings" json:"vectorize_bindings"`
}

type PagesProjectDeploymentConfigsPreviewPlacementModel struct {
	Mode types.String `tfsdk:"mode" json:"mode"`
}

type PagesProjectDeploymentConfigsProductionModel struct {
	AIBindings              map[string]jsontypes.Normalized                        `tfsdk:"ai_bindings" json:"ai_bindings"`
	AnalyticsEngineDatasets map[string]jsontypes.Normalized                        `tfsdk:"analytics_engine_datasets" json:"analytics_engine_datasets"`
	Browsers                map[string]jsontypes.Normalized                        `tfsdk:"browsers" json:"browsers"`
	CompatibilityDate       types.String                                           `tfsdk:"compatibility_date" json:"compatibility_date"`
	CompatibilityFlags      *[]types.String                                        `tfsdk:"compatibility_flags" json:"compatibility_flags"`
	D1Databases             map[string]jsontypes.Normalized                        `tfsdk:"d1_databases" json:"d1_databases"`
	DurableObjectNamespaces map[string]jsontypes.Normalized                        `tfsdk:"durable_object_namespaces" json:"durable_object_namespaces"`
	EnvVars                 map[string]jsontypes.Normalized                        `tfsdk:"env_vars" json:"env_vars"`
	HyperdriveBindings      map[string]jsontypes.Normalized                        `tfsdk:"hyperdrive_bindings" json:"hyperdrive_bindings"`
	KVNamespaces            map[string]jsontypes.Normalized                        `tfsdk:"kv_namespaces" json:"kv_namespaces"`
	MTLSCertificates        map[string]jsontypes.Normalized                        `tfsdk:"mtls_certificates" json:"mtls_certificates"`
	Placement               *PagesProjectDeploymentConfigsProductionPlacementModel `tfsdk:"placement" json:"placement"`
	QueueProducers          map[string]jsontypes.Normalized                        `tfsdk:"queue_producers" json:"queue_producers"`
	R2Buckets               map[string]jsontypes.Normalized                        `tfsdk:"r2_buckets" json:"r2_buckets"`
	Services                map[string]jsontypes.Normalized                        `tfsdk:"services" json:"services"`
	VectorizeBindings       map[string]jsontypes.Normalized                        `tfsdk:"vectorize_bindings" json:"vectorize_bindings"`
}

type PagesProjectDeploymentConfigsProductionPlacementModel struct {
	Mode types.String `tfsdk:"mode" json:"mode"`
}

type PagesProjectCanonicalDeploymentModel struct {
	ID                types.String                                                                    `tfsdk:"id" json:"id,computed"`
	Aliases           *[]types.String                                                                 `tfsdk:"aliases" json:"aliases,computed"`
	BuildConfig       *PagesProjectCanonicalDeploymentBuildConfigModel                                `tfsdk:"build_config" json:"build_config"`
	CreatedOn         timetypes.RFC3339                                                               `tfsdk:"created_on" json:"created_on,computed"`
	DeploymentTrigger customfield.NestedObject[PagesProjectCanonicalDeploymentDeploymentTriggerModel] `tfsdk:"deployment_trigger" json:"deployment_trigger,computed"`
	EnvVars           map[string]jsontypes.Normalized                                                 `tfsdk:"env_vars" json:"env_vars,computed"`
	Environment       types.String                                                                    `tfsdk:"environment" json:"environment,computed"`
	IsSkipped         types.Bool                                                                      `tfsdk:"is_skipped" json:"is_skipped,computed"`
	LatestStage       customfield.NestedObject[PagesProjectCanonicalDeploymentLatestStageModel]       `tfsdk:"latest_stage" json:"latest_stage,computed"`
	ModifiedOn        timetypes.RFC3339                                                               `tfsdk:"modified_on" json:"modified_on,computed"`
	ProjectID         types.String                                                                    `tfsdk:"project_id" json:"project_id,computed"`
	ProjectName       types.String                                                                    `tfsdk:"project_name" json:"project_name,computed"`
	ShortID           types.String                                                                    `tfsdk:"short_id" json:"short_id,computed"`
	Source            customfield.NestedObject[PagesProjectCanonicalDeploymentSourceModel]            `tfsdk:"source" json:"source,computed"`
	Stages            *[]*PagesProjectCanonicalDeploymentStagesModel                                  `tfsdk:"stages" json:"stages,computed"`
	URL               types.String                                                                    `tfsdk:"url" json:"url,computed"`
}

type PagesProjectCanonicalDeploymentBuildConfigModel struct {
	BuildCaching      types.Bool   `tfsdk:"build_caching" json:"build_caching"`
	BuildCommand      types.String `tfsdk:"build_command" json:"build_command"`
	DestinationDir    types.String `tfsdk:"destination_dir" json:"destination_dir"`
	RootDir           types.String `tfsdk:"root_dir" json:"root_dir"`
	WebAnalyticsTag   types.String `tfsdk:"web_analytics_tag" json:"web_analytics_tag"`
	WebAnalyticsToken types.String `tfsdk:"web_analytics_token" json:"web_analytics_token"`
}

type PagesProjectCanonicalDeploymentDeploymentTriggerModel struct {
	Metadata *PagesProjectCanonicalDeploymentDeploymentTriggerMetadataModel `tfsdk:"metadata" json:"metadata"`
	Type     types.String                                                   `tfsdk:"type" json:"type,computed"`
}

type PagesProjectCanonicalDeploymentDeploymentTriggerMetadataModel struct {
	Branch        types.String `tfsdk:"branch" json:"branch,computed"`
	CommitHash    types.String `tfsdk:"commit_hash" json:"commit_hash,computed"`
	CommitMessage types.String `tfsdk:"commit_message" json:"commit_message,computed"`
}

type PagesProjectCanonicalDeploymentLatestStageModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed"`
	Name      types.String      `tfsdk:"name" json:"name"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
}

type PagesProjectCanonicalDeploymentSourceModel struct {
	Config *PagesProjectCanonicalDeploymentSourceConfigModel `tfsdk:"config" json:"config"`
	Type   types.String                                      `tfsdk:"type" json:"type"`
}

type PagesProjectCanonicalDeploymentSourceConfigModel struct {
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

type PagesProjectCanonicalDeploymentStagesModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed"`
	Name      types.String      `tfsdk:"name" json:"name"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
}

type PagesProjectLatestDeploymentModel struct {
	ID                types.String                                                                 `tfsdk:"id" json:"id,computed"`
	Aliases           *[]types.String                                                              `tfsdk:"aliases" json:"aliases,computed"`
	BuildConfig       *PagesProjectLatestDeploymentBuildConfigModel                                `tfsdk:"build_config" json:"build_config"`
	CreatedOn         timetypes.RFC3339                                                            `tfsdk:"created_on" json:"created_on,computed"`
	DeploymentTrigger customfield.NestedObject[PagesProjectLatestDeploymentDeploymentTriggerModel] `tfsdk:"deployment_trigger" json:"deployment_trigger,computed"`
	EnvVars           map[string]jsontypes.Normalized                                              `tfsdk:"env_vars" json:"env_vars,computed"`
	Environment       types.String                                                                 `tfsdk:"environment" json:"environment,computed"`
	IsSkipped         types.Bool                                                                   `tfsdk:"is_skipped" json:"is_skipped,computed"`
	LatestStage       customfield.NestedObject[PagesProjectLatestDeploymentLatestStageModel]       `tfsdk:"latest_stage" json:"latest_stage,computed"`
	ModifiedOn        timetypes.RFC3339                                                            `tfsdk:"modified_on" json:"modified_on,computed"`
	ProjectID         types.String                                                                 `tfsdk:"project_id" json:"project_id,computed"`
	ProjectName       types.String                                                                 `tfsdk:"project_name" json:"project_name,computed"`
	ShortID           types.String                                                                 `tfsdk:"short_id" json:"short_id,computed"`
	Source            customfield.NestedObject[PagesProjectLatestDeploymentSourceModel]            `tfsdk:"source" json:"source,computed"`
	Stages            *[]*PagesProjectLatestDeploymentStagesModel                                  `tfsdk:"stages" json:"stages,computed"`
	URL               types.String                                                                 `tfsdk:"url" json:"url,computed"`
}

type PagesProjectLatestDeploymentBuildConfigModel struct {
	BuildCaching      types.Bool   `tfsdk:"build_caching" json:"build_caching"`
	BuildCommand      types.String `tfsdk:"build_command" json:"build_command"`
	DestinationDir    types.String `tfsdk:"destination_dir" json:"destination_dir"`
	RootDir           types.String `tfsdk:"root_dir" json:"root_dir"`
	WebAnalyticsTag   types.String `tfsdk:"web_analytics_tag" json:"web_analytics_tag"`
	WebAnalyticsToken types.String `tfsdk:"web_analytics_token" json:"web_analytics_token"`
}

type PagesProjectLatestDeploymentDeploymentTriggerModel struct {
	Metadata *PagesProjectLatestDeploymentDeploymentTriggerMetadataModel `tfsdk:"metadata" json:"metadata"`
	Type     types.String                                                `tfsdk:"type" json:"type,computed"`
}

type PagesProjectLatestDeploymentDeploymentTriggerMetadataModel struct {
	Branch        types.String `tfsdk:"branch" json:"branch,computed"`
	CommitHash    types.String `tfsdk:"commit_hash" json:"commit_hash,computed"`
	CommitMessage types.String `tfsdk:"commit_message" json:"commit_message,computed"`
}

type PagesProjectLatestDeploymentLatestStageModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed"`
	Name      types.String      `tfsdk:"name" json:"name"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
}

type PagesProjectLatestDeploymentSourceModel struct {
	Config *PagesProjectLatestDeploymentSourceConfigModel `tfsdk:"config" json:"config"`
	Type   types.String                                   `tfsdk:"type" json:"type"`
}

type PagesProjectLatestDeploymentSourceConfigModel struct {
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

type PagesProjectLatestDeploymentStagesModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed"`
	Name      types.String      `tfsdk:"name" json:"name"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
}

type PagesProjectSourceModel struct {
	Config *PagesProjectSourceConfigModel `tfsdk:"config" json:"config"`
	Type   types.String                   `tfsdk:"type" json:"type"`
}

type PagesProjectSourceConfigModel struct {
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
