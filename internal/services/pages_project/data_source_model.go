// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_project

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/pages"
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
	AccountID           types.String                                                `tfsdk:"account_id" path:"account_id"`
	ProjectName         types.String                                                `tfsdk:"project_name" path:"project_name,computed_optional"`
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
	BuildConfig         *PagesProjectBuildConfigDataSourceModel                     `tfsdk:"build_config" json:"build_config,computed_optional"`
	Filter              *PagesProjectFindOneByDataSourceModel                       `tfsdk:"filter"`
}

func (m *PagesProjectDataSourceModel) toReadParams() (params pages.ProjectGetParams, diags diag.Diagnostics) {
	params = pages.ProjectGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *PagesProjectDataSourceModel) toListParams() (params pages.ProjectListParams, diags diag.Diagnostics) {
	params = pages.ProjectListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type PagesProjectCanonicalDeploymentDataSourceModel struct {
	ID                types.String                                                                              `tfsdk:"id" json:"id,computed"`
	Aliases           types.List                                                                                `tfsdk:"aliases" json:"aliases,computed"`
	BuildConfig       *PagesProjectCanonicalDeploymentBuildConfigDataSourceModel                                `tfsdk:"build_config" json:"build_config,computed_optional"`
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
	Stages            customfield.NestedObjectList[PagesProjectCanonicalDeploymentStagesDataSourceModel]        `tfsdk:"stages" json:"stages,computed"`
	URL               types.String                                                                              `tfsdk:"url" json:"url,computed"`
}

type PagesProjectCanonicalDeploymentBuildConfigDataSourceModel struct {
	BuildCaching      types.Bool   `tfsdk:"build_caching" json:"build_caching,computed_optional"`
	BuildCommand      types.String `tfsdk:"build_command" json:"build_command,computed_optional"`
	DestinationDir    types.String `tfsdk:"destination_dir" json:"destination_dir,computed_optional"`
	RootDir           types.String `tfsdk:"root_dir" json:"root_dir,computed_optional"`
	WebAnalyticsTag   types.String `tfsdk:"web_analytics_tag" json:"web_analytics_tag,computed_optional"`
	WebAnalyticsToken types.String `tfsdk:"web_analytics_token" json:"web_analytics_token,computed_optional"`
}

type PagesProjectCanonicalDeploymentDeploymentTriggerDataSourceModel struct {
	Metadata *PagesProjectCanonicalDeploymentDeploymentTriggerMetadataDataSourceModel `tfsdk:"metadata" json:"metadata,computed_optional"`
	Type     types.String                                                             `tfsdk:"type" json:"type,computed"`
}

type PagesProjectCanonicalDeploymentDeploymentTriggerMetadataDataSourceModel struct {
	Branch        types.String `tfsdk:"branch" json:"branch,computed"`
	CommitHash    types.String `tfsdk:"commit_hash" json:"commit_hash,computed"`
	CommitMessage types.String `tfsdk:"commit_message" json:"commit_message,computed"`
}

type PagesProjectCanonicalDeploymentLatestStageDataSourceModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed"`
	Name      types.String      `tfsdk:"name" json:"name,computed_optional"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
}

type PagesProjectCanonicalDeploymentSourceDataSourceModel struct {
	Config *PagesProjectCanonicalDeploymentSourceConfigDataSourceModel `tfsdk:"config" json:"config,computed_optional"`
	Type   types.String                                                `tfsdk:"type" json:"type,computed_optional"`
}

type PagesProjectCanonicalDeploymentSourceConfigDataSourceModel struct {
	DeploymentsEnabled           types.Bool      `tfsdk:"deployments_enabled" json:"deployments_enabled,computed_optional"`
	Owner                        types.String    `tfsdk:"owner" json:"owner,computed_optional"`
	PathExcludes                 *[]types.String `tfsdk:"path_excludes" json:"path_excludes,computed_optional"`
	PathIncludes                 *[]types.String `tfsdk:"path_includes" json:"path_includes,computed_optional"`
	PrCommentsEnabled            types.Bool      `tfsdk:"pr_comments_enabled" json:"pr_comments_enabled,computed_optional"`
	PreviewBranchExcludes        *[]types.String `tfsdk:"preview_branch_excludes" json:"preview_branch_excludes,computed_optional"`
	PreviewBranchIncludes        *[]types.String `tfsdk:"preview_branch_includes" json:"preview_branch_includes,computed_optional"`
	PreviewDeploymentSetting     types.String    `tfsdk:"preview_deployment_setting" json:"preview_deployment_setting,computed_optional"`
	ProductionBranch             types.String    `tfsdk:"production_branch" json:"production_branch,computed_optional"`
	ProductionDeploymentsEnabled types.Bool      `tfsdk:"production_deployments_enabled" json:"production_deployments_enabled,computed_optional"`
	RepoName                     types.String    `tfsdk:"repo_name" json:"repo_name,computed_optional"`
}

type PagesProjectCanonicalDeploymentStagesDataSourceModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed"`
	Name      types.String      `tfsdk:"name" json:"name,computed_optional"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
}

type PagesProjectDeploymentConfigsDataSourceModel struct {
	Preview    *PagesProjectDeploymentConfigsPreviewDataSourceModel    `tfsdk:"preview" json:"preview,computed_optional"`
	Production *PagesProjectDeploymentConfigsProductionDataSourceModel `tfsdk:"production" json:"production,computed_optional"`
}

type PagesProjectDeploymentConfigsPreviewDataSourceModel struct {
	AIBindings              map[string]jsontypes.Normalized                               `tfsdk:"ai_bindings" json:"ai_bindings,computed_optional"`
	AnalyticsEngineDatasets map[string]jsontypes.Normalized                               `tfsdk:"analytics_engine_datasets" json:"analytics_engine_datasets,computed_optional"`
	Browsers                map[string]jsontypes.Normalized                               `tfsdk:"browsers" json:"browsers,computed_optional"`
	CompatibilityDate       types.String                                                  `tfsdk:"compatibility_date" json:"compatibility_date,computed_optional"`
	CompatibilityFlags      *[]types.String                                               `tfsdk:"compatibility_flags" json:"compatibility_flags,computed_optional"`
	D1Databases             map[string]jsontypes.Normalized                               `tfsdk:"d1_databases" json:"d1_databases,computed_optional"`
	DurableObjectNamespaces map[string]jsontypes.Normalized                               `tfsdk:"durable_object_namespaces" json:"durable_object_namespaces,computed_optional"`
	EnvVars                 map[string]jsontypes.Normalized                               `tfsdk:"env_vars" json:"env_vars,computed_optional"`
	HyperdriveBindings      map[string]jsontypes.Normalized                               `tfsdk:"hyperdrive_bindings" json:"hyperdrive_bindings,computed_optional"`
	KVNamespaces            map[string]jsontypes.Normalized                               `tfsdk:"kv_namespaces" json:"kv_namespaces,computed_optional"`
	MTLSCertificates        map[string]jsontypes.Normalized                               `tfsdk:"mtls_certificates" json:"mtls_certificates,computed_optional"`
	Placement               *PagesProjectDeploymentConfigsPreviewPlacementDataSourceModel `tfsdk:"placement" json:"placement,computed_optional"`
	QueueProducers          map[string]jsontypes.Normalized                               `tfsdk:"queue_producers" json:"queue_producers,computed_optional"`
	R2Buckets               map[string]jsontypes.Normalized                               `tfsdk:"r2_buckets" json:"r2_buckets,computed_optional"`
	Services                map[string]jsontypes.Normalized                               `tfsdk:"services" json:"services,computed_optional"`
	VectorizeBindings       map[string]jsontypes.Normalized                               `tfsdk:"vectorize_bindings" json:"vectorize_bindings,computed_optional"`
}

type PagesProjectDeploymentConfigsPreviewPlacementDataSourceModel struct {
	Mode types.String `tfsdk:"mode" json:"mode,computed_optional"`
}

type PagesProjectDeploymentConfigsProductionDataSourceModel struct {
	AIBindings              map[string]jsontypes.Normalized                                  `tfsdk:"ai_bindings" json:"ai_bindings,computed_optional"`
	AnalyticsEngineDatasets map[string]jsontypes.Normalized                                  `tfsdk:"analytics_engine_datasets" json:"analytics_engine_datasets,computed_optional"`
	Browsers                map[string]jsontypes.Normalized                                  `tfsdk:"browsers" json:"browsers,computed_optional"`
	CompatibilityDate       types.String                                                     `tfsdk:"compatibility_date" json:"compatibility_date,computed_optional"`
	CompatibilityFlags      *[]types.String                                                  `tfsdk:"compatibility_flags" json:"compatibility_flags,computed_optional"`
	D1Databases             map[string]jsontypes.Normalized                                  `tfsdk:"d1_databases" json:"d1_databases,computed_optional"`
	DurableObjectNamespaces map[string]jsontypes.Normalized                                  `tfsdk:"durable_object_namespaces" json:"durable_object_namespaces,computed_optional"`
	EnvVars                 map[string]jsontypes.Normalized                                  `tfsdk:"env_vars" json:"env_vars,computed_optional"`
	HyperdriveBindings      map[string]jsontypes.Normalized                                  `tfsdk:"hyperdrive_bindings" json:"hyperdrive_bindings,computed_optional"`
	KVNamespaces            map[string]jsontypes.Normalized                                  `tfsdk:"kv_namespaces" json:"kv_namespaces,computed_optional"`
	MTLSCertificates        map[string]jsontypes.Normalized                                  `tfsdk:"mtls_certificates" json:"mtls_certificates,computed_optional"`
	Placement               *PagesProjectDeploymentConfigsProductionPlacementDataSourceModel `tfsdk:"placement" json:"placement,computed_optional"`
	QueueProducers          map[string]jsontypes.Normalized                                  `tfsdk:"queue_producers" json:"queue_producers,computed_optional"`
	R2Buckets               map[string]jsontypes.Normalized                                  `tfsdk:"r2_buckets" json:"r2_buckets,computed_optional"`
	Services                map[string]jsontypes.Normalized                                  `tfsdk:"services" json:"services,computed_optional"`
	VectorizeBindings       map[string]jsontypes.Normalized                                  `tfsdk:"vectorize_bindings" json:"vectorize_bindings,computed_optional"`
}

type PagesProjectDeploymentConfigsProductionPlacementDataSourceModel struct {
	Mode types.String `tfsdk:"mode" json:"mode,computed_optional"`
}

type PagesProjectDeploymentTriggerDataSourceModel struct {
	Metadata *PagesProjectDeploymentTriggerMetadataDataSourceModel `tfsdk:"metadata" json:"metadata,computed_optional"`
	Type     types.String                                          `tfsdk:"type" json:"type,computed"`
}

type PagesProjectDeploymentTriggerMetadataDataSourceModel struct {
	Branch        types.String `tfsdk:"branch" json:"branch,computed"`
	CommitHash    types.String `tfsdk:"commit_hash" json:"commit_hash,computed"`
	CommitMessage types.String `tfsdk:"commit_message" json:"commit_message,computed"`
}

type PagesProjectLatestDeploymentDataSourceModel struct {
	ID                types.String                                                                           `tfsdk:"id" json:"id,computed"`
	Aliases           types.List                                                                             `tfsdk:"aliases" json:"aliases,computed"`
	BuildConfig       *PagesProjectLatestDeploymentBuildConfigDataSourceModel                                `tfsdk:"build_config" json:"build_config,computed_optional"`
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
	Stages            customfield.NestedObjectList[PagesProjectLatestDeploymentStagesDataSourceModel]        `tfsdk:"stages" json:"stages,computed"`
	URL               types.String                                                                           `tfsdk:"url" json:"url,computed"`
}

type PagesProjectLatestDeploymentBuildConfigDataSourceModel struct {
	BuildCaching      types.Bool   `tfsdk:"build_caching" json:"build_caching,computed_optional"`
	BuildCommand      types.String `tfsdk:"build_command" json:"build_command,computed_optional"`
	DestinationDir    types.String `tfsdk:"destination_dir" json:"destination_dir,computed_optional"`
	RootDir           types.String `tfsdk:"root_dir" json:"root_dir,computed_optional"`
	WebAnalyticsTag   types.String `tfsdk:"web_analytics_tag" json:"web_analytics_tag,computed_optional"`
	WebAnalyticsToken types.String `tfsdk:"web_analytics_token" json:"web_analytics_token,computed_optional"`
}

type PagesProjectLatestDeploymentDeploymentTriggerDataSourceModel struct {
	Metadata *PagesProjectLatestDeploymentDeploymentTriggerMetadataDataSourceModel `tfsdk:"metadata" json:"metadata,computed_optional"`
	Type     types.String                                                          `tfsdk:"type" json:"type,computed"`
}

type PagesProjectLatestDeploymentDeploymentTriggerMetadataDataSourceModel struct {
	Branch        types.String `tfsdk:"branch" json:"branch,computed"`
	CommitHash    types.String `tfsdk:"commit_hash" json:"commit_hash,computed"`
	CommitMessage types.String `tfsdk:"commit_message" json:"commit_message,computed"`
}

type PagesProjectLatestDeploymentLatestStageDataSourceModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed"`
	Name      types.String      `tfsdk:"name" json:"name,computed_optional"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
}

type PagesProjectLatestDeploymentSourceDataSourceModel struct {
	Config *PagesProjectLatestDeploymentSourceConfigDataSourceModel `tfsdk:"config" json:"config,computed_optional"`
	Type   types.String                                             `tfsdk:"type" json:"type,computed_optional"`
}

type PagesProjectLatestDeploymentSourceConfigDataSourceModel struct {
	DeploymentsEnabled           types.Bool      `tfsdk:"deployments_enabled" json:"deployments_enabled,computed_optional"`
	Owner                        types.String    `tfsdk:"owner" json:"owner,computed_optional"`
	PathExcludes                 *[]types.String `tfsdk:"path_excludes" json:"path_excludes,computed_optional"`
	PathIncludes                 *[]types.String `tfsdk:"path_includes" json:"path_includes,computed_optional"`
	PrCommentsEnabled            types.Bool      `tfsdk:"pr_comments_enabled" json:"pr_comments_enabled,computed_optional"`
	PreviewBranchExcludes        *[]types.String `tfsdk:"preview_branch_excludes" json:"preview_branch_excludes,computed_optional"`
	PreviewBranchIncludes        *[]types.String `tfsdk:"preview_branch_includes" json:"preview_branch_includes,computed_optional"`
	PreviewDeploymentSetting     types.String    `tfsdk:"preview_deployment_setting" json:"preview_deployment_setting,computed_optional"`
	ProductionBranch             types.String    `tfsdk:"production_branch" json:"production_branch,computed_optional"`
	ProductionDeploymentsEnabled types.Bool      `tfsdk:"production_deployments_enabled" json:"production_deployments_enabled,computed_optional"`
	RepoName                     types.String    `tfsdk:"repo_name" json:"repo_name,computed_optional"`
}

type PagesProjectLatestDeploymentStagesDataSourceModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed"`
	Name      types.String      `tfsdk:"name" json:"name,computed_optional"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
}

type PagesProjectLatestStageDataSourceModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed"`
	Name      types.String      `tfsdk:"name" json:"name,computed_optional"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
}

type PagesProjectStagesDataSourceModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed"`
	Name      types.String      `tfsdk:"name" json:"name,computed_optional"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
}

type PagesProjectSourceDataSourceModel struct {
	Config *PagesProjectSourceConfigDataSourceModel `tfsdk:"config" json:"config,computed_optional"`
	Type   types.String                             `tfsdk:"type" json:"type,computed_optional"`
}

type PagesProjectSourceConfigDataSourceModel struct {
	DeploymentsEnabled           types.Bool      `tfsdk:"deployments_enabled" json:"deployments_enabled,computed_optional"`
	Owner                        types.String    `tfsdk:"owner" json:"owner,computed_optional"`
	PathExcludes                 *[]types.String `tfsdk:"path_excludes" json:"path_excludes,computed_optional"`
	PathIncludes                 *[]types.String `tfsdk:"path_includes" json:"path_includes,computed_optional"`
	PrCommentsEnabled            types.Bool      `tfsdk:"pr_comments_enabled" json:"pr_comments_enabled,computed_optional"`
	PreviewBranchExcludes        *[]types.String `tfsdk:"preview_branch_excludes" json:"preview_branch_excludes,computed_optional"`
	PreviewBranchIncludes        *[]types.String `tfsdk:"preview_branch_includes" json:"preview_branch_includes,computed_optional"`
	PreviewDeploymentSetting     types.String    `tfsdk:"preview_deployment_setting" json:"preview_deployment_setting,computed_optional"`
	ProductionBranch             types.String    `tfsdk:"production_branch" json:"production_branch,computed_optional"`
	ProductionDeploymentsEnabled types.Bool      `tfsdk:"production_deployments_enabled" json:"production_deployments_enabled,computed_optional"`
	RepoName                     types.String    `tfsdk:"repo_name" json:"repo_name,computed_optional"`
}

type PagesProjectBuildConfigDataSourceModel struct {
	BuildCaching      types.Bool   `tfsdk:"build_caching" json:"build_caching,computed_optional"`
	BuildCommand      types.String `tfsdk:"build_command" json:"build_command,computed_optional"`
	DestinationDir    types.String `tfsdk:"destination_dir" json:"destination_dir,computed_optional"`
	RootDir           types.String `tfsdk:"root_dir" json:"root_dir,computed_optional"`
	WebAnalyticsTag   types.String `tfsdk:"web_analytics_tag" json:"web_analytics_tag,computed_optional"`
	WebAnalyticsToken types.String `tfsdk:"web_analytics_token" json:"web_analytics_token,computed_optional"`
}

type PagesProjectFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
