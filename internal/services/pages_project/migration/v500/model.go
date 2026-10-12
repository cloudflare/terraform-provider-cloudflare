package v500

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// =============================================================================
// V0 (v4 provider) Source Model Definitions
// =============================================================================

// SourcePagesProjectModelV0 represents the v4 state structure.
// Note: v4 SDKv2 stores blocks as lists, so all block types are slices.
type SourcePagesProjectModelV0 struct {
	ID                types.String                                 `tfsdk:"id"`
	Name              types.String                                 `tfsdk:"name"`
	AccountID         types.String                                 `tfsdk:"account_id"`
	ProductionBranch  types.String                                 `tfsdk:"production_branch"`
	BuildConfig       []SourcePagesProjectBuildConfigModelV0       `tfsdk:"build_config"`
	Source            []SourcePagesProjectSourceModelV0            `tfsdk:"source"`
	DeploymentConfigs []SourcePagesProjectDeploymentConfigsModelV0 `tfsdk:"deployment_configs"`
	CreatedOn         timetypes.RFC3339                            `tfsdk:"created_on"`
	Subdomain         types.String                                 `tfsdk:"subdomain"`
	Domains           types.List                                   `tfsdk:"domains"`
}

type SourcePagesProjectBuildConfigModelV0 struct {
	BuildCaching      types.Bool   `tfsdk:"build_caching"`
	BuildCommand      types.String `tfsdk:"build_command"`
	DestinationDir    types.String `tfsdk:"destination_dir"`
	RootDir           types.String `tfsdk:"root_dir"`
	WebAnalyticsTag   types.String `tfsdk:"web_analytics_tag"`
	WebAnalyticsToken types.String `tfsdk:"web_analytics_token"`
}

type SourcePagesProjectSourceModelV0 struct {
	Type   types.String                            `tfsdk:"type"`
	Config []SourcePagesProjectSourceConfigModelV0 `tfsdk:"config"`
}

type SourcePagesProjectSourceConfigModelV0 struct {
	DeploymentsEnabled          types.Bool   `tfsdk:"deployments_enabled"`
	Owner                       types.String `tfsdk:"owner"`
	PathExcludes                types.List   `tfsdk:"path_excludes"`
	PathIncludes                types.List   `tfsdk:"path_includes"`
	PrCommentsEnabled           types.Bool   `tfsdk:"pr_comments_enabled"`
	PreviewBranchExcludes       types.List   `tfsdk:"preview_branch_excludes"`
	PreviewBranchIncludes       types.List   `tfsdk:"preview_branch_includes"`
	PreviewDeploymentSetting    types.String `tfsdk:"preview_deployment_setting"`
	ProductionBranch            types.String `tfsdk:"production_branch"`
	ProductionDeploymentEnabled types.Bool   `tfsdk:"production_deployment_enabled"` // renamed in v5
	RepoName                    types.String `tfsdk:"repo_name"`
}

type SourcePagesProjectDeploymentConfigsModelV0 struct {
	Preview    []SourcePagesProjectDeploymentConfigModelV0 `tfsdk:"preview"`
	Production []SourcePagesProjectDeploymentConfigModelV0 `tfsdk:"production"`
}

type SourcePagesProjectDeploymentConfigModelV0 struct {
	AlwaysUseLatestCompatibilityDate types.Bool                           `tfsdk:"always_use_latest_compatibility_date"`
	CompatibilityDate                types.String                         `tfsdk:"compatibility_date"`
	CompatibilityFlags               types.List                           `tfsdk:"compatibility_flags"`
	UsageModel                       types.String                         `tfsdk:"usage_model"`
	FailOpen                         types.Bool                           `tfsdk:"fail_open"`
	EnvironmentVariables             types.Map                            `tfsdk:"environment_variables"`
	Secrets                          types.Map                            `tfsdk:"secrets"`
	KVNamespaces                     types.Map                            `tfsdk:"kv_namespaces"`
	D1Databases                      types.Map                            `tfsdk:"d1_databases"`
	R2Buckets                        types.Map                            `tfsdk:"r2_buckets"`
	DurableObjectNamespaces          types.Map                            `tfsdk:"durable_object_namespaces"`
	ServiceBindings                  types.List                           `tfsdk:"service_binding"`
	Placement                        []SourcePagesProjectPlacementModelV0 `tfsdk:"placement"`
}

type SourcePagesProjectPlacementModelV0 struct {
	Mode types.String `tfsdk:"mode"`
}

type SourcePagesProjectServiceBindingModelV0 struct {
	Name        types.String `tfsdk:"name"`
	Service     types.String `tfsdk:"service"`
	Environment types.String `tfsdk:"environment"`
	Entrypoint  types.String `tfsdk:"entrypoint"`
}

// =============================================================================
// Target Model Definitions (v5 provider)
// These are local copies to avoid import cycles with the parent package.
// =============================================================================

// TargetPagesProjectModel represents the v5 state structure.
type TargetPagesProjectModel struct {
	ID                    types.String                                                           `tfsdk:"id"`
	Name                  types.String                                                           `tfsdk:"name"`
	AccountID             types.String                                                           `tfsdk:"account_id"`
	ProductionBranch      types.String                                                           `tfsdk:"production_branch"`
	BuildConfig           customfield.NestedObject[TargetPagesProjectBuildConfigModel]           `tfsdk:"build_config"`
	BuildImageAutoUpgrade customfield.NestedObject[TargetPagesProjectBuildImageAutoUpgradeModel] `tfsdk:"build_image_auto_upgrade"`
	Source                *TargetPagesProjectSourceModel                                         `tfsdk:"source"`
	DeploymentConfigs     customfield.NestedObject[TargetPagesProjectDeploymentConfigsModel]     `tfsdk:"deployment_configs"`
	CreatedOn             timetypes.RFC3339                                                      `tfsdk:"created_on"`
	Framework             types.String                                                           `tfsdk:"framework"`
	FrameworkVersion      types.String                                                           `tfsdk:"framework_version"`
	PreviewScriptName     types.String                                                           `tfsdk:"preview_script_name"`
	ProductionScriptName  types.String                                                           `tfsdk:"production_script_name"`
	Subdomain             types.String                                                           `tfsdk:"subdomain"`
	UsesFunctions         types.Bool                                                             `tfsdk:"uses_functions"`
	Domains               customfield.List[types.String]                                         `tfsdk:"domains"`
	CanonicalDeployment   customfield.NestedObject[TargetPagesProjectCanonicalDeploymentModel]   `tfsdk:"canonical_deployment"`
	LatestDeployment      customfield.NestedObject[TargetPagesProjectLatestDeploymentModel]      `tfsdk:"latest_deployment"`
}

type TargetPagesProjectBuildImageAutoUpgradeModel struct {
	FromVersion types.Int64 `tfsdk:"from_version"`
	ToVersion   types.Int64 `tfsdk:"to_version"`
}

type TargetPagesProjectBuildConfigModel struct {
	BuildCaching      types.Bool   `tfsdk:"build_caching"`
	BuildCommand      types.String `tfsdk:"build_command"`
	DestinationDir    types.String `tfsdk:"destination_dir"`
	RootDir           types.String `tfsdk:"root_dir"`
	WebAnalyticsTag   types.String `tfsdk:"web_analytics_tag"`
	WebAnalyticsToken types.String `tfsdk:"web_analytics_token"`
}

type TargetPagesProjectSourceModel struct {
	Config *TargetPagesProjectSourceConfigModel `tfsdk:"config"`
	Type   types.String                         `tfsdk:"type"`
}

type TargetPagesProjectSourceConfigModel struct {
	DeploymentsEnabled           types.Bool                     `tfsdk:"deployments_enabled"`
	Owner                        types.String                   `tfsdk:"owner"`
	OwnerID                      types.String                   `tfsdk:"owner_id"`
	PathExcludes                 customfield.List[types.String] `tfsdk:"path_excludes"`
	PathIncludes                 customfield.List[types.String] `tfsdk:"path_includes"`
	PrCommentsEnabled            types.Bool                     `tfsdk:"pr_comments_enabled"`
	PreviewBranchExcludes        customfield.List[types.String] `tfsdk:"preview_branch_excludes"`
	PreviewBranchIncludes        customfield.List[types.String] `tfsdk:"preview_branch_includes"`
	PreviewDeploymentSetting     types.String                   `tfsdk:"preview_deployment_setting"`
	ProductionBranch             types.String                   `tfsdk:"production_branch"`
	ProductionDeploymentsEnabled types.Bool                     `tfsdk:"production_deployments_enabled"`
	RepoID                       types.String                   `tfsdk:"repo_id"`
	RepoName                     types.String                   `tfsdk:"repo_name"`
}

type TargetPagesProjectDeploymentConfigsModel struct {
	Preview    customfield.NestedObject[TargetPagesProjectDeploymentConfigsPreviewModel]    `tfsdk:"preview"`
	Production customfield.NestedObject[TargetPagesProjectDeploymentConfigsProductionModel] `tfsdk:"production"`
}

type TargetPagesProjectDeploymentConfigsPreviewModel struct {
	AIBindings                       *map[string]TargetPagesProjectDeploymentConfigsPreviewAIBindingsModel              `tfsdk:"ai_bindings"`
	AlwaysUseLatestCompatibilityDate types.Bool                                                                         `tfsdk:"always_use_latest_compatibility_date"`
	AnalyticsEngineDatasets          *map[string]TargetPagesProjectDeploymentConfigsPreviewAnalyticsEngineDatasetsModel `tfsdk:"analytics_engine_datasets"`
	Browsers                         *map[string]TargetPagesProjectDeploymentConfigsPreviewBrowsersModel                `tfsdk:"browsers"`
	BuildImageMajorVersion           types.Int64                                                                        `tfsdk:"build_image_major_version"`
	CompatibilityDate                types.String                                                                       `tfsdk:"compatibility_date"`
	CompatibilityFlags               *[]types.String                                                                    `tfsdk:"compatibility_flags"`
	D1Databases                      *map[string]TargetPagesProjectDeploymentConfigsPreviewD1DatabasesModel             `tfsdk:"d1_databases"`
	DurableObjectNamespaces          *map[string]TargetPagesProjectDeploymentConfigsPreviewDurableObjectNamespacesModel `tfsdk:"durable_object_namespaces"`
	EnvVars                          *map[string]TargetPagesProjectDeploymentConfigsPreviewEnvVarsModel                 `tfsdk:"env_vars"`
	FailOpen                         types.Bool                                                                         `tfsdk:"fail_open"`
	HyperdriveBindings               *map[string]TargetPagesProjectDeploymentConfigsPreviewHyperdriveBindingsModel      `tfsdk:"hyperdrive_bindings"`
	KVNamespaces                     *map[string]TargetPagesProjectDeploymentConfigsPreviewKVNamespacesModel            `tfsdk:"kv_namespaces"`
	Limits                           *TargetPagesProjectDeploymentConfigsPreviewLimitsModel                             `tfsdk:"limits"`
	MTLSCertificates                 *map[string]TargetPagesProjectDeploymentConfigsPreviewMTLSCertificatesModel        `tfsdk:"mtls_certificates"`
	Placement                        *TargetPagesProjectDeploymentConfigsPreviewPlacementModel                          `tfsdk:"placement"`
	QueueProducers                   *map[string]TargetPagesProjectDeploymentConfigsPreviewQueueProducersModel          `tfsdk:"queue_producers"`
	R2Buckets                        *map[string]TargetPagesProjectDeploymentConfigsPreviewR2BucketsModel               `tfsdk:"r2_buckets"`
	Services                         *map[string]TargetPagesProjectDeploymentConfigsPreviewServicesModel                `tfsdk:"services"`
	UsageModel                       types.String                                                                       `tfsdk:"usage_model"`
	VectorizeBindings                *map[string]TargetPagesProjectDeploymentConfigsPreviewVectorizeBindingsModel       `tfsdk:"vectorize_bindings"`
	WranglerConfigHash               types.String                                                                       `tfsdk:"wrangler_config_hash"`
}

type TargetPagesProjectDeploymentConfigsPreviewAIBindingsModel struct {
	ProjectID types.String `tfsdk:"project_id"`
}

type TargetPagesProjectDeploymentConfigsPreviewAnalyticsEngineDatasetsModel struct {
	Dataset types.String `tfsdk:"dataset"`
}

type TargetPagesProjectDeploymentConfigsPreviewBrowsersModel struct {
}

type TargetPagesProjectDeploymentConfigsPreviewD1DatabasesModel struct {
	ID types.String `tfsdk:"id"`
}

type TargetPagesProjectDeploymentConfigsPreviewDurableObjectNamespacesModel struct {
	NamespaceID types.String `tfsdk:"namespace_id"`
}

type TargetPagesProjectDeploymentConfigsPreviewEnvVarsModel struct {
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

type TargetPagesProjectDeploymentConfigsPreviewKVNamespacesModel struct {
	NamespaceID types.String `tfsdk:"namespace_id"`
}

type TargetPagesProjectDeploymentConfigsPreviewHyperdriveBindingsModel struct {
	ID types.String `tfsdk:"id"`
}

type TargetPagesProjectDeploymentConfigsPreviewLimitsModel struct {
	CPUMs types.Int64 `tfsdk:"cpu_ms"`
}

type TargetPagesProjectDeploymentConfigsPreviewMTLSCertificatesModel struct {
	CertificateID types.String `tfsdk:"certificate_id"`
}

type TargetPagesProjectDeploymentConfigsPreviewPlacementModel struct {
	Mode types.String `tfsdk:"mode"`
}

type TargetPagesProjectDeploymentConfigsPreviewQueueProducersModel struct {
	Name types.String `tfsdk:"name"`
}

type TargetPagesProjectDeploymentConfigsPreviewR2BucketsModel struct {
	Name         types.String `tfsdk:"name"`
	Jurisdiction types.String `tfsdk:"jurisdiction"`
}

type TargetPagesProjectDeploymentConfigsPreviewServicesModel struct {
	Service     types.String `tfsdk:"service"`
	Entrypoint  types.String `tfsdk:"entrypoint"`
	Environment types.String `tfsdk:"environment"`
}

type TargetPagesProjectDeploymentConfigsPreviewVectorizeBindingsModel struct {
	IndexName types.String `tfsdk:"index_name"`
}

type TargetPagesProjectDeploymentConfigsProductionModel struct {
	AIBindings                       *map[string]TargetPagesProjectDeploymentConfigsProductionAIBindingsModel              `tfsdk:"ai_bindings"`
	AlwaysUseLatestCompatibilityDate types.Bool                                                                            `tfsdk:"always_use_latest_compatibility_date"`
	AnalyticsEngineDatasets          *map[string]TargetPagesProjectDeploymentConfigsProductionAnalyticsEngineDatasetsModel `tfsdk:"analytics_engine_datasets"`
	Browsers                         *map[string]TargetPagesProjectDeploymentConfigsProductionBrowsersModel                `tfsdk:"browsers"`
	BuildImageMajorVersion           types.Int64                                                                           `tfsdk:"build_image_major_version"`
	CompatibilityDate                types.String                                                                          `tfsdk:"compatibility_date"`
	CompatibilityFlags               *[]types.String                                                                       `tfsdk:"compatibility_flags"`
	D1Databases                      *map[string]TargetPagesProjectDeploymentConfigsProductionD1DatabasesModel             `tfsdk:"d1_databases"`
	DurableObjectNamespaces          *map[string]TargetPagesProjectDeploymentConfigsProductionDurableObjectNamespacesModel `tfsdk:"durable_object_namespaces"`
	EnvVars                          *map[string]TargetPagesProjectDeploymentConfigsProductionEnvVarsModel                 `tfsdk:"env_vars"`
	FailOpen                         types.Bool                                                                            `tfsdk:"fail_open"`
	HyperdriveBindings               *map[string]TargetPagesProjectDeploymentConfigsProductionHyperdriveBindingsModel      `tfsdk:"hyperdrive_bindings"`
	KVNamespaces                     *map[string]TargetPagesProjectDeploymentConfigsProductionKVNamespacesModel            `tfsdk:"kv_namespaces"`
	Limits                           *TargetPagesProjectDeploymentConfigsProductionLimitsModel                             `tfsdk:"limits"`
	MTLSCertificates                 *map[string]TargetPagesProjectDeploymentConfigsProductionMTLSCertificatesModel        `tfsdk:"mtls_certificates"`
	Placement                        *TargetPagesProjectDeploymentConfigsProductionPlacementModel                          `tfsdk:"placement"`
	QueueProducers                   *map[string]TargetPagesProjectDeploymentConfigsProductionQueueProducersModel          `tfsdk:"queue_producers"`
	R2Buckets                        *map[string]TargetPagesProjectDeploymentConfigsProductionR2BucketsModel               `tfsdk:"r2_buckets"`
	Services                         *map[string]TargetPagesProjectDeploymentConfigsProductionServicesModel                `tfsdk:"services"`
	UsageModel                       types.String                                                                          `tfsdk:"usage_model"`
	VectorizeBindings                *map[string]TargetPagesProjectDeploymentConfigsProductionVectorizeBindingsModel       `tfsdk:"vectorize_bindings"`
	WranglerConfigHash               types.String                                                                          `tfsdk:"wrangler_config_hash"`
}

type TargetPagesProjectDeploymentConfigsProductionAIBindingsModel struct {
	ProjectID types.String `tfsdk:"project_id"`
}

type TargetPagesProjectDeploymentConfigsProductionAnalyticsEngineDatasetsModel struct {
	Dataset types.String `tfsdk:"dataset"`
}

type TargetPagesProjectDeploymentConfigsProductionBrowsersModel struct {
}

type TargetPagesProjectDeploymentConfigsProductionD1DatabasesModel struct {
	ID types.String `tfsdk:"id"`
}

type TargetPagesProjectDeploymentConfigsProductionDurableObjectNamespacesModel struct {
	NamespaceID types.String `tfsdk:"namespace_id"`
}

type TargetPagesProjectDeploymentConfigsProductionEnvVarsModel struct {
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

type TargetPagesProjectDeploymentConfigsProductionKVNamespacesModel struct {
	NamespaceID types.String `tfsdk:"namespace_id"`
}

type TargetPagesProjectDeploymentConfigsProductionHyperdriveBindingsModel struct {
	ID types.String `tfsdk:"id"`
}

type TargetPagesProjectDeploymentConfigsProductionLimitsModel struct {
	CPUMs types.Int64 `tfsdk:"cpu_ms"`
}

type TargetPagesProjectDeploymentConfigsProductionMTLSCertificatesModel struct {
	CertificateID types.String `tfsdk:"certificate_id"`
}

type TargetPagesProjectDeploymentConfigsProductionPlacementModel struct {
	Mode types.String `tfsdk:"mode"`
}

type TargetPagesProjectDeploymentConfigsProductionQueueProducersModel struct {
	Name types.String `tfsdk:"name"`
}

type TargetPagesProjectDeploymentConfigsProductionR2BucketsModel struct {
	Name         types.String `tfsdk:"name"`
	Jurisdiction types.String `tfsdk:"jurisdiction"`
}

type TargetPagesProjectDeploymentConfigsProductionServicesModel struct {
	Service     types.String `tfsdk:"service"`
	Entrypoint  types.String `tfsdk:"entrypoint"`
	Environment types.String `tfsdk:"environment"`
}

type TargetPagesProjectDeploymentConfigsProductionVectorizeBindingsModel struct {
	IndexName types.String `tfsdk:"index_name"`
}

// =============================================================================
// Canonical and Latest Deployment Models (computed, read-only)
// =============================================================================

type TargetPagesProjectCanonicalDeploymentModel struct {
	ID                types.String                                                                          `tfsdk:"id"`
	Aliases           customfield.List[types.String]                                                        `tfsdk:"aliases"`
	BuildConfig       customfield.NestedObject[TargetPagesProjectCanonicalDeploymentBuildConfigModel]       `tfsdk:"build_config"`
	CreatedOn         timetypes.RFC3339                                                                     `tfsdk:"created_on"`
	DeploymentTrigger customfield.NestedObject[TargetPagesProjectCanonicalDeploymentDeploymentTriggerModel] `tfsdk:"deployment_trigger"`
	EnvVars           customfield.NestedObjectMap[TargetPagesProjectCanonicalDeploymentEnvVarsModel]        `tfsdk:"env_vars"`
	Environment       types.String                                                                          `tfsdk:"environment"`
	IsSkipped         types.Bool                                                                            `tfsdk:"is_skipped"`
	LatestStage       customfield.NestedObject[TargetPagesProjectCanonicalDeploymentLatestStageModel]       `tfsdk:"latest_stage"`
	ModifiedOn        timetypes.RFC3339                                                                     `tfsdk:"modified_on"`
	ProjectID         types.String                                                                          `tfsdk:"project_id"`
	ProjectName       types.String                                                                          `tfsdk:"project_name"`
	ShortID           types.String                                                                          `tfsdk:"short_id"`
	Source            customfield.NestedObject[TargetPagesProjectCanonicalDeploymentSourceModel]            `tfsdk:"source"`
	Stages            customfield.NestedObjectList[TargetPagesProjectCanonicalDeploymentStagesModel]        `tfsdk:"stages"`
	URL               types.String                                                                          `tfsdk:"url"`
	UsesFunctions     types.Bool                                                                            `tfsdk:"uses_functions"`
}

type TargetPagesProjectCanonicalDeploymentBuildConfigModel struct {
	WebAnalyticsTag   types.String `tfsdk:"web_analytics_tag"`
	WebAnalyticsToken types.String `tfsdk:"web_analytics_token"`
	BuildCaching      types.Bool   `tfsdk:"build_caching"`
	BuildCommand      types.String `tfsdk:"build_command"`
	DestinationDir    types.String `tfsdk:"destination_dir"`
	RootDir           types.String `tfsdk:"root_dir"`
}

type TargetPagesProjectCanonicalDeploymentDeploymentTriggerModel struct {
	Metadata customfield.NestedObject[TargetPagesProjectCanonicalDeploymentDeploymentTriggerMetadataModel] `tfsdk:"metadata"`
	Type     types.String                                                                                  `tfsdk:"type"`
}

type TargetPagesProjectCanonicalDeploymentDeploymentTriggerMetadataModel struct {
	Branch        types.String `tfsdk:"branch"`
	CommitDirty   types.Bool   `tfsdk:"commit_dirty"`
	CommitHash    types.String `tfsdk:"commit_hash"`
	CommitMessage types.String `tfsdk:"commit_message"`
}

type TargetPagesProjectCanonicalDeploymentEnvVarsModel struct {
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

type TargetPagesProjectCanonicalDeploymentLatestStageModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on"`
	Name      types.String      `tfsdk:"name"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on"`
	Status    types.String      `tfsdk:"status"`
}

type TargetPagesProjectCanonicalDeploymentSourceModel struct {
	Config customfield.NestedObject[TargetPagesProjectCanonicalDeploymentSourceConfigModel] `tfsdk:"config"`
	Type   types.String                                                                     `tfsdk:"type"`
}

type TargetPagesProjectCanonicalDeploymentSourceConfigModel struct {
	DeploymentsEnabled           types.Bool                     `tfsdk:"deployments_enabled"`
	Owner                        types.String                   `tfsdk:"owner"`
	OwnerID                      types.String                   `tfsdk:"owner_id"`
	PathExcludes                 customfield.List[types.String] `tfsdk:"path_excludes"`
	PathIncludes                 customfield.List[types.String] `tfsdk:"path_includes"`
	PrCommentsEnabled            types.Bool                     `tfsdk:"pr_comments_enabled"`
	PreviewBranchExcludes        customfield.List[types.String] `tfsdk:"preview_branch_excludes"`
	PreviewBranchIncludes        customfield.List[types.String] `tfsdk:"preview_branch_includes"`
	PreviewDeploymentSetting     types.String                   `tfsdk:"preview_deployment_setting"`
	ProductionBranch             types.String                   `tfsdk:"production_branch"`
	ProductionDeploymentsEnabled types.Bool                     `tfsdk:"production_deployments_enabled"`
	RepoID                       types.String                   `tfsdk:"repo_id"`
	RepoName                     types.String                   `tfsdk:"repo_name"`
}

type TargetPagesProjectCanonicalDeploymentStagesModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on"`
	Name      types.String      `tfsdk:"name"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on"`
	Status    types.String      `tfsdk:"status"`
}

type TargetPagesProjectLatestDeploymentModel struct {
	ID                types.String                                                                       `tfsdk:"id"`
	Aliases           customfield.List[types.String]                                                     `tfsdk:"aliases"`
	BuildConfig       customfield.NestedObject[TargetPagesProjectLatestDeploymentBuildConfigModel]       `tfsdk:"build_config"`
	CreatedOn         timetypes.RFC3339                                                                  `tfsdk:"created_on"`
	DeploymentTrigger customfield.NestedObject[TargetPagesProjectLatestDeploymentDeploymentTriggerModel] `tfsdk:"deployment_trigger"`
	EnvVars           customfield.NestedObjectMap[TargetPagesProjectLatestDeploymentEnvVarsModel]        `tfsdk:"env_vars"`
	Environment       types.String                                                                       `tfsdk:"environment"`
	IsSkipped         types.Bool                                                                         `tfsdk:"is_skipped"`
	LatestStage       customfield.NestedObject[TargetPagesProjectLatestDeploymentLatestStageModel]       `tfsdk:"latest_stage"`
	ModifiedOn        timetypes.RFC3339                                                                  `tfsdk:"modified_on"`
	ProjectID         types.String                                                                       `tfsdk:"project_id"`
	ProjectName       types.String                                                                       `tfsdk:"project_name"`
	ShortID           types.String                                                                       `tfsdk:"short_id"`
	Source            customfield.NestedObject[TargetPagesProjectLatestDeploymentSourceModel]            `tfsdk:"source"`
	Stages            customfield.NestedObjectList[TargetPagesProjectLatestDeploymentStagesModel]        `tfsdk:"stages"`
	URL               types.String                                                                       `tfsdk:"url"`
	UsesFunctions     types.Bool                                                                         `tfsdk:"uses_functions"`
}

type TargetPagesProjectLatestDeploymentBuildConfigModel struct {
	WebAnalyticsTag   types.String `tfsdk:"web_analytics_tag"`
	WebAnalyticsToken types.String `tfsdk:"web_analytics_token"`
	BuildCaching      types.Bool   `tfsdk:"build_caching"`
	BuildCommand      types.String `tfsdk:"build_command"`
	DestinationDir    types.String `tfsdk:"destination_dir"`
	RootDir           types.String `tfsdk:"root_dir"`
}

type TargetPagesProjectLatestDeploymentDeploymentTriggerModel struct {
	Metadata customfield.NestedObject[TargetPagesProjectLatestDeploymentDeploymentTriggerMetadataModel] `tfsdk:"metadata"`
	Type     types.String                                                                               `tfsdk:"type"`
}

type TargetPagesProjectLatestDeploymentDeploymentTriggerMetadataModel struct {
	Branch        types.String `tfsdk:"branch"`
	CommitDirty   types.Bool   `tfsdk:"commit_dirty"`
	CommitHash    types.String `tfsdk:"commit_hash"`
	CommitMessage types.String `tfsdk:"commit_message"`
}

type TargetPagesProjectLatestDeploymentEnvVarsModel struct {
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

type TargetPagesProjectLatestDeploymentLatestStageModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on"`
	Name      types.String      `tfsdk:"name"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on"`
	Status    types.String      `tfsdk:"status"`
}

type TargetPagesProjectLatestDeploymentSourceModel struct {
	Config customfield.NestedObject[TargetPagesProjectLatestDeploymentSourceConfigModel] `tfsdk:"config"`
	Type   types.String                                                                  `tfsdk:"type"`
}

type TargetPagesProjectLatestDeploymentSourceConfigModel struct {
	DeploymentsEnabled           types.Bool                     `tfsdk:"deployments_enabled"`
	Owner                        types.String                   `tfsdk:"owner"`
	OwnerID                      types.String                   `tfsdk:"owner_id"`
	PathExcludes                 customfield.List[types.String] `tfsdk:"path_excludes"`
	PathIncludes                 customfield.List[types.String] `tfsdk:"path_includes"`
	PrCommentsEnabled            types.Bool                     `tfsdk:"pr_comments_enabled"`
	PreviewBranchExcludes        customfield.List[types.String] `tfsdk:"preview_branch_excludes"`
	PreviewBranchIncludes        customfield.List[types.String] `tfsdk:"preview_branch_includes"`
	PreviewDeploymentSetting     types.String                   `tfsdk:"preview_deployment_setting"`
	ProductionBranch             types.String                   `tfsdk:"production_branch"`
	ProductionDeploymentsEnabled types.Bool                     `tfsdk:"production_deployments_enabled"`
	RepoID                       types.String                   `tfsdk:"repo_id"`
	RepoName                     types.String                   `tfsdk:"repo_name"`
}

type TargetPagesProjectLatestDeploymentStagesModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on"`
	Name      types.String      `tfsdk:"name"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on"`
	Status    types.String      `tfsdk:"status"`
}
