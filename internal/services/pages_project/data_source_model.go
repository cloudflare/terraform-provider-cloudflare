// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_project

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PagesProjectResultDataSourceEnvelope struct {
	Result PagesProjectDataSourceModel `json:"result,computed"`
}

type PagesProjectResultListDataSourceEnvelope struct {
	Result *[]*PagesProjectDataSourceModel `json:"result,computed"`
}

type PagesProjectDataSourceModel struct {
	AccountID           types.String                                    `tfsdk:"account_id" path:"account_id"`
	ProjectName         types.String                                    `tfsdk:"project_name" path:"project_name"`
	ID                  types.String                                    `tfsdk:"id" json:"id"`
	BuildConfig         *PagesProjectBuildConfigDataSourceModel         `tfsdk:"build_config" json:"build_config"`
	CanonicalDeployment *PagesProjectCanonicalDeploymentDataSourceModel `tfsdk:"canonical_deployment" json:"canonical_deployment"`
	CreatedOn           types.String                                    `tfsdk:"created_on" json:"created_on"`
	DeploymentConfigs   *PagesProjectDeploymentConfigsDataSourceModel   `tfsdk:"deployment_configs" json:"deployment_configs"`
	Domains             types.String                                    `tfsdk:"domains" json:"domains"`
	LatestDeployment    *PagesProjectLatestDeploymentDataSourceModel    `tfsdk:"latest_deployment" json:"latest_deployment"`
	Name                types.String                                    `tfsdk:"name" json:"name"`
	ProductionBranch    types.String                                    `tfsdk:"production_branch" json:"production_branch"`
	Source              types.String                                    `tfsdk:"source" json:"source"`
	Subdomain           types.String                                    `tfsdk:"subdomain" json:"subdomain"`
	Aliases             types.String                                    `tfsdk:"aliases" json:"aliases"`
	DeploymentTrigger   *PagesProjectDeploymentTriggerDataSourceModel   `tfsdk:"deployment_trigger" json:"deployment_trigger"`
	EnvVars             types.String                                    `tfsdk:"env_vars" json:"env_vars"`
	Environment         types.String                                    `tfsdk:"environment" json:"environment"`
	IsSkipped           types.Bool                                      `tfsdk:"is_skipped" json:"is_skipped"`
	LatestStage         types.String                                    `tfsdk:"latest_stage" json:"latest_stage"`
	ModifiedOn          types.String                                    `tfsdk:"modified_on" json:"modified_on"`
	ProjectID           types.String                                    `tfsdk:"project_id" json:"project_id"`
	ShortID             types.String                                    `tfsdk:"short_id" json:"short_id"`
	Stages              *[]*PagesProjectStagesDataSourceModel           `tfsdk:"stages" json:"stages"`
	URL                 types.String                                    `tfsdk:"url" json:"url"`
	FindOneBy           *PagesProjectFindOneByDataSourceModel           `tfsdk:"find_one_by"`
}

type PagesProjectBuildConfigDataSourceModel struct {
	BuildCaching      types.Bool   `tfsdk:"build_caching" json:"build_caching"`
	BuildCommand      types.String `tfsdk:"build_command" json:"build_command"`
	DestinationDir    types.String `tfsdk:"destination_dir" json:"destination_dir"`
	RootDir           types.String `tfsdk:"root_dir" json:"root_dir"`
	WebAnalyticsTag   types.String `tfsdk:"web_analytics_tag" json:"web_analytics_tag"`
	WebAnalyticsToken types.String `tfsdk:"web_analytics_token" json:"web_analytics_token"`
}

type PagesProjectCanonicalDeploymentDataSourceModel struct {
	ID          types.String                                             `tfsdk:"id" json:"id,computed"`
	Aliases     *[]types.String                                          `tfsdk:"aliases" json:"aliases,computed"`
	BuildConfig types.String                                             `tfsdk:"build_config" json:"build_config,computed"`
	CreatedOn   types.String                                             `tfsdk:"created_on" json:"created_on,computed"`
	EnvVars     types.String                                             `tfsdk:"env_vars" json:"env_vars,computed"`
	Environment types.String                                             `tfsdk:"environment" json:"environment,computed"`
	IsSkipped   types.Bool                                               `tfsdk:"is_skipped" json:"is_skipped,computed"`
	LatestStage types.String                                             `tfsdk:"latest_stage" json:"latest_stage,computed"`
	ModifiedOn  types.String                                             `tfsdk:"modified_on" json:"modified_on,computed"`
	ProjectID   types.String                                             `tfsdk:"project_id" json:"project_id,computed"`
	ProjectName types.String                                             `tfsdk:"project_name" json:"project_name,computed"`
	ShortID     types.String                                             `tfsdk:"short_id" json:"short_id,computed"`
	Source      types.String                                             `tfsdk:"source" json:"source,computed"`
	Stages      *[]*PagesProjectCanonicalDeploymentStagesDataSourceModel `tfsdk:"stages" json:"stages,computed"`
	URL         types.String                                             `tfsdk:"url" json:"url,computed"`
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

type PagesProjectCanonicalDeploymentStagesDataSourceModel struct {
	EndedOn   types.String `tfsdk:"ended_on" json:"ended_on,computed"`
	Name      types.String `tfsdk:"name" json:"name"`
	StartedOn types.String `tfsdk:"started_on" json:"started_on,computed"`
	Status    types.String `tfsdk:"status" json:"status,computed"`
}

type PagesProjectDeploymentConfigsDataSourceModel struct {
	Preview    *PagesProjectDeploymentConfigsPreviewDataSourceModel    `tfsdk:"preview" json:"preview"`
	Production *PagesProjectDeploymentConfigsProductionDataSourceModel `tfsdk:"production" json:"production"`
}

type PagesProjectDeploymentConfigsPreviewDataSourceModel struct {
	AIBindings              *PagesProjectDeploymentConfigsPreviewAIBindingsDataSourceModel              `tfsdk:"ai_bindings" json:"ai_bindings"`
	AnalyticsEngineDatasets *PagesProjectDeploymentConfigsPreviewAnalyticsEngineDatasetsDataSourceModel `tfsdk:"analytics_engine_datasets" json:"analytics_engine_datasets"`
	Browsers                *PagesProjectDeploymentConfigsPreviewBrowsersDataSourceModel                `tfsdk:"browsers" json:"browsers"`
	CompatibilityDate       types.String                                                                `tfsdk:"compatibility_date" json:"compatibility_date"`
	CompatibilityFlags      *[]types.String                                                             `tfsdk:"compatibility_flags" json:"compatibility_flags"`
	D1Databases             *PagesProjectDeploymentConfigsPreviewD1DatabasesDataSourceModel             `tfsdk:"d1_databases" json:"d1_databases"`
	DurableObjectNamespaces *PagesProjectDeploymentConfigsPreviewDurableObjectNamespacesDataSourceModel `tfsdk:"durable_object_namespaces" json:"durable_object_namespaces"`
	EnvVars                 *PagesProjectDeploymentConfigsPreviewEnvVarsDataSourceModel                 `tfsdk:"env_vars" json:"env_vars"`
	HyperdriveBindings      *PagesProjectDeploymentConfigsPreviewHyperdriveBindingsDataSourceModel      `tfsdk:"hyperdrive_bindings" json:"hyperdrive_bindings"`
	KVNamespaces            *PagesProjectDeploymentConfigsPreviewKVNamespacesDataSourceModel            `tfsdk:"kv_namespaces" json:"kv_namespaces"`
	MTLSCertificates        *PagesProjectDeploymentConfigsPreviewMTLSCertificatesDataSourceModel        `tfsdk:"mtls_certificates" json:"mtls_certificates"`
	Placement               *PagesProjectDeploymentConfigsPreviewPlacementDataSourceModel               `tfsdk:"placement" json:"placement"`
	QueueProducers          *PagesProjectDeploymentConfigsPreviewQueueProducersDataSourceModel          `tfsdk:"queue_producers" json:"queue_producers"`
	R2Buckets               *PagesProjectDeploymentConfigsPreviewR2BucketsDataSourceModel               `tfsdk:"r2_buckets" json:"r2_buckets"`
	Services                *PagesProjectDeploymentConfigsPreviewServicesDataSourceModel                `tfsdk:"services" json:"services"`
	VectorizeBindings       *PagesProjectDeploymentConfigsPreviewVectorizeBindingsDataSourceModel       `tfsdk:"vectorize_bindings" json:"vectorize_bindings"`
}

type PagesProjectDeploymentConfigsPreviewAIBindingsDataSourceModel struct {
	AIBinding *PagesProjectDeploymentConfigsPreviewAIBindingsAIBindingDataSourceModel `tfsdk:"ai_binding" json:"AI_BINDING"`
}

type PagesProjectDeploymentConfigsPreviewAIBindingsAIBindingDataSourceModel struct {
	ProjectID types.String `tfsdk:"project_id" json:"project_id"`
}

type PagesProjectDeploymentConfigsPreviewAnalyticsEngineDatasetsDataSourceModel struct {
	AnalyticsEngineBinding *PagesProjectDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBindingDataSourceModel `tfsdk:"analytics_engine_binding" json:"ANALYTICS_ENGINE_BINDING"`
}

type PagesProjectDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBindingDataSourceModel struct {
	Dataset types.String `tfsdk:"dataset" json:"dataset"`
}

type PagesProjectDeploymentConfigsPreviewBrowsersDataSourceModel struct {
	Browser types.String `tfsdk:"browser" json:"BROWSER"`
}

type PagesProjectDeploymentConfigsPreviewD1DatabasesDataSourceModel struct {
	D1Binding *PagesProjectDeploymentConfigsPreviewD1DatabasesD1BindingDataSourceModel `tfsdk:"d1_binding" json:"D1_BINDING"`
}

type PagesProjectDeploymentConfigsPreviewD1DatabasesD1BindingDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type PagesProjectDeploymentConfigsPreviewDurableObjectNamespacesDataSourceModel struct {
	DoBinding *PagesProjectDeploymentConfigsPreviewDurableObjectNamespacesDoBindingDataSourceModel `tfsdk:"do_binding" json:"DO_BINDING"`
}

type PagesProjectDeploymentConfigsPreviewDurableObjectNamespacesDoBindingDataSourceModel struct {
	NamespaceID types.String `tfsdk:"namespace_id" json:"namespace_id"`
}

type PagesProjectDeploymentConfigsPreviewEnvVarsDataSourceModel struct {
	EnvironmentVariable *PagesProjectDeploymentConfigsPreviewEnvVarsEnvironmentVariableDataSourceModel `tfsdk:"environment_variable" json:"ENVIRONMENT_VARIABLE"`
}

type PagesProjectDeploymentConfigsPreviewEnvVarsEnvironmentVariableDataSourceModel struct {
	Type  types.String `tfsdk:"type" json:"type"`
	Value types.String `tfsdk:"value" json:"value"`
}

type PagesProjectDeploymentConfigsPreviewHyperdriveBindingsDataSourceModel struct {
	Hyperdrive *PagesProjectDeploymentConfigsPreviewHyperdriveBindingsHyperdriveDataSourceModel `tfsdk:"hyperdrive" json:"HYPERDRIVE"`
}

type PagesProjectDeploymentConfigsPreviewHyperdriveBindingsHyperdriveDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type PagesProjectDeploymentConfigsPreviewKVNamespacesDataSourceModel struct {
	KVBinding *PagesProjectDeploymentConfigsPreviewKVNamespacesKVBindingDataSourceModel `tfsdk:"kv_binding" json:"KV_BINDING"`
}

type PagesProjectDeploymentConfigsPreviewKVNamespacesKVBindingDataSourceModel struct {
	NamespaceID types.String `tfsdk:"namespace_id" json:"namespace_id"`
}

type PagesProjectDeploymentConfigsPreviewMTLSCertificatesDataSourceModel struct {
	MTLS *PagesProjectDeploymentConfigsPreviewMTLSCertificatesMTLSDataSourceModel `tfsdk:"mtls" json:"MTLS"`
}

type PagesProjectDeploymentConfigsPreviewMTLSCertificatesMTLSDataSourceModel struct {
	CertificateID types.String `tfsdk:"certificate_id" json:"certificate_id"`
}

type PagesProjectDeploymentConfigsPreviewPlacementDataSourceModel struct {
	Mode types.String `tfsdk:"mode" json:"mode"`
}

type PagesProjectDeploymentConfigsPreviewQueueProducersDataSourceModel struct {
	QueueProducerBinding *PagesProjectDeploymentConfigsPreviewQueueProducersQueueProducerBindingDataSourceModel `tfsdk:"queue_producer_binding" json:"QUEUE_PRODUCER_BINDING"`
}

type PagesProjectDeploymentConfigsPreviewQueueProducersQueueProducerBindingDataSourceModel struct {
	Name types.String `tfsdk:"name" json:"name"`
}

type PagesProjectDeploymentConfigsPreviewR2BucketsDataSourceModel struct {
	R2Binding *PagesProjectDeploymentConfigsPreviewR2BucketsR2BindingDataSourceModel `tfsdk:"r2_binding" json:"R2_BINDING"`
}

type PagesProjectDeploymentConfigsPreviewR2BucketsR2BindingDataSourceModel struct {
	Jurisdiction types.String `tfsdk:"jurisdiction" json:"jurisdiction"`
	Name         types.String `tfsdk:"name" json:"name"`
}

type PagesProjectDeploymentConfigsPreviewServicesDataSourceModel struct {
	ServiceBinding *PagesProjectDeploymentConfigsPreviewServicesServiceBindingDataSourceModel `tfsdk:"service_binding" json:"SERVICE_BINDING"`
}

type PagesProjectDeploymentConfigsPreviewServicesServiceBindingDataSourceModel struct {
	Entrypoint  types.String `tfsdk:"entrypoint" json:"entrypoint"`
	Environment types.String `tfsdk:"environment" json:"environment"`
	Service     types.String `tfsdk:"service" json:"service"`
}

type PagesProjectDeploymentConfigsPreviewVectorizeBindingsDataSourceModel struct {
	Vectorize *PagesProjectDeploymentConfigsPreviewVectorizeBindingsVectorizeDataSourceModel `tfsdk:"vectorize" json:"VECTORIZE"`
}

type PagesProjectDeploymentConfigsPreviewVectorizeBindingsVectorizeDataSourceModel struct {
	IndexName types.String `tfsdk:"index_name" json:"index_name"`
}

type PagesProjectDeploymentConfigsProductionDataSourceModel struct {
	AIBindings              *PagesProjectDeploymentConfigsProductionAIBindingsDataSourceModel              `tfsdk:"ai_bindings" json:"ai_bindings"`
	AnalyticsEngineDatasets *PagesProjectDeploymentConfigsProductionAnalyticsEngineDatasetsDataSourceModel `tfsdk:"analytics_engine_datasets" json:"analytics_engine_datasets"`
	Browsers                *PagesProjectDeploymentConfigsProductionBrowsersDataSourceModel                `tfsdk:"browsers" json:"browsers"`
	CompatibilityDate       types.String                                                                   `tfsdk:"compatibility_date" json:"compatibility_date"`
	CompatibilityFlags      *[]types.String                                                                `tfsdk:"compatibility_flags" json:"compatibility_flags"`
	D1Databases             *PagesProjectDeploymentConfigsProductionD1DatabasesDataSourceModel             `tfsdk:"d1_databases" json:"d1_databases"`
	DurableObjectNamespaces *PagesProjectDeploymentConfigsProductionDurableObjectNamespacesDataSourceModel `tfsdk:"durable_object_namespaces" json:"durable_object_namespaces"`
	EnvVars                 *PagesProjectDeploymentConfigsProductionEnvVarsDataSourceModel                 `tfsdk:"env_vars" json:"env_vars"`
	HyperdriveBindings      *PagesProjectDeploymentConfigsProductionHyperdriveBindingsDataSourceModel      `tfsdk:"hyperdrive_bindings" json:"hyperdrive_bindings"`
	KVNamespaces            *PagesProjectDeploymentConfigsProductionKVNamespacesDataSourceModel            `tfsdk:"kv_namespaces" json:"kv_namespaces"`
	MTLSCertificates        *PagesProjectDeploymentConfigsProductionMTLSCertificatesDataSourceModel        `tfsdk:"mtls_certificates" json:"mtls_certificates"`
	Placement               *PagesProjectDeploymentConfigsProductionPlacementDataSourceModel               `tfsdk:"placement" json:"placement"`
	QueueProducers          *PagesProjectDeploymentConfigsProductionQueueProducersDataSourceModel          `tfsdk:"queue_producers" json:"queue_producers"`
	R2Buckets               *PagesProjectDeploymentConfigsProductionR2BucketsDataSourceModel               `tfsdk:"r2_buckets" json:"r2_buckets"`
	Services                *PagesProjectDeploymentConfigsProductionServicesDataSourceModel                `tfsdk:"services" json:"services"`
	VectorizeBindings       *PagesProjectDeploymentConfigsProductionVectorizeBindingsDataSourceModel       `tfsdk:"vectorize_bindings" json:"vectorize_bindings"`
}

type PagesProjectDeploymentConfigsProductionAIBindingsDataSourceModel struct {
	AIBinding *PagesProjectDeploymentConfigsProductionAIBindingsAIBindingDataSourceModel `tfsdk:"ai_binding" json:"AI_BINDING"`
}

type PagesProjectDeploymentConfigsProductionAIBindingsAIBindingDataSourceModel struct {
	ProjectID types.String `tfsdk:"project_id" json:"project_id"`
}

type PagesProjectDeploymentConfigsProductionAnalyticsEngineDatasetsDataSourceModel struct {
	AnalyticsEngineBinding *PagesProjectDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBindingDataSourceModel `tfsdk:"analytics_engine_binding" json:"ANALYTICS_ENGINE_BINDING"`
}

type PagesProjectDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBindingDataSourceModel struct {
	Dataset types.String `tfsdk:"dataset" json:"dataset"`
}

type PagesProjectDeploymentConfigsProductionBrowsersDataSourceModel struct {
	Browser types.String `tfsdk:"browser" json:"BROWSER"`
}

type PagesProjectDeploymentConfigsProductionD1DatabasesDataSourceModel struct {
	D1Binding *PagesProjectDeploymentConfigsProductionD1DatabasesD1BindingDataSourceModel `tfsdk:"d1_binding" json:"D1_BINDING"`
}

type PagesProjectDeploymentConfigsProductionD1DatabasesD1BindingDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type PagesProjectDeploymentConfigsProductionDurableObjectNamespacesDataSourceModel struct {
	DoBinding *PagesProjectDeploymentConfigsProductionDurableObjectNamespacesDoBindingDataSourceModel `tfsdk:"do_binding" json:"DO_BINDING"`
}

type PagesProjectDeploymentConfigsProductionDurableObjectNamespacesDoBindingDataSourceModel struct {
	NamespaceID types.String `tfsdk:"namespace_id" json:"namespace_id"`
}

type PagesProjectDeploymentConfigsProductionEnvVarsDataSourceModel struct {
	EnvironmentVariable *PagesProjectDeploymentConfigsProductionEnvVarsEnvironmentVariableDataSourceModel `tfsdk:"environment_variable" json:"ENVIRONMENT_VARIABLE"`
}

type PagesProjectDeploymentConfigsProductionEnvVarsEnvironmentVariableDataSourceModel struct {
	Type  types.String `tfsdk:"type" json:"type"`
	Value types.String `tfsdk:"value" json:"value"`
}

type PagesProjectDeploymentConfigsProductionHyperdriveBindingsDataSourceModel struct {
	Hyperdrive *PagesProjectDeploymentConfigsProductionHyperdriveBindingsHyperdriveDataSourceModel `tfsdk:"hyperdrive" json:"HYPERDRIVE"`
}

type PagesProjectDeploymentConfigsProductionHyperdriveBindingsHyperdriveDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type PagesProjectDeploymentConfigsProductionKVNamespacesDataSourceModel struct {
	KVBinding *PagesProjectDeploymentConfigsProductionKVNamespacesKVBindingDataSourceModel `tfsdk:"kv_binding" json:"KV_BINDING"`
}

type PagesProjectDeploymentConfigsProductionKVNamespacesKVBindingDataSourceModel struct {
	NamespaceID types.String `tfsdk:"namespace_id" json:"namespace_id"`
}

type PagesProjectDeploymentConfigsProductionMTLSCertificatesDataSourceModel struct {
	MTLS *PagesProjectDeploymentConfigsProductionMTLSCertificatesMTLSDataSourceModel `tfsdk:"mtls" json:"MTLS"`
}

type PagesProjectDeploymentConfigsProductionMTLSCertificatesMTLSDataSourceModel struct {
	CertificateID types.String `tfsdk:"certificate_id" json:"certificate_id"`
}

type PagesProjectDeploymentConfigsProductionPlacementDataSourceModel struct {
	Mode types.String `tfsdk:"mode" json:"mode"`
}

type PagesProjectDeploymentConfigsProductionQueueProducersDataSourceModel struct {
	QueueProducerBinding *PagesProjectDeploymentConfigsProductionQueueProducersQueueProducerBindingDataSourceModel `tfsdk:"queue_producer_binding" json:"QUEUE_PRODUCER_BINDING"`
}

type PagesProjectDeploymentConfigsProductionQueueProducersQueueProducerBindingDataSourceModel struct {
	Name types.String `tfsdk:"name" json:"name"`
}

type PagesProjectDeploymentConfigsProductionR2BucketsDataSourceModel struct {
	R2Binding *PagesProjectDeploymentConfigsProductionR2BucketsR2BindingDataSourceModel `tfsdk:"r2_binding" json:"R2_BINDING"`
}

type PagesProjectDeploymentConfigsProductionR2BucketsR2BindingDataSourceModel struct {
	Jurisdiction types.String `tfsdk:"jurisdiction" json:"jurisdiction"`
	Name         types.String `tfsdk:"name" json:"name"`
}

type PagesProjectDeploymentConfigsProductionServicesDataSourceModel struct {
	ServiceBinding *PagesProjectDeploymentConfigsProductionServicesServiceBindingDataSourceModel `tfsdk:"service_binding" json:"SERVICE_BINDING"`
}

type PagesProjectDeploymentConfigsProductionServicesServiceBindingDataSourceModel struct {
	Entrypoint  types.String `tfsdk:"entrypoint" json:"entrypoint"`
	Environment types.String `tfsdk:"environment" json:"environment"`
	Service     types.String `tfsdk:"service" json:"service"`
}

type PagesProjectDeploymentConfigsProductionVectorizeBindingsDataSourceModel struct {
	Vectorize *PagesProjectDeploymentConfigsProductionVectorizeBindingsVectorizeDataSourceModel `tfsdk:"vectorize" json:"VECTORIZE"`
}

type PagesProjectDeploymentConfigsProductionVectorizeBindingsVectorizeDataSourceModel struct {
	IndexName types.String `tfsdk:"index_name" json:"index_name"`
}

type PagesProjectLatestDeploymentDataSourceModel struct {
	ID          types.String                                          `tfsdk:"id" json:"id,computed"`
	Aliases     *[]types.String                                       `tfsdk:"aliases" json:"aliases,computed"`
	BuildConfig types.String                                          `tfsdk:"build_config" json:"build_config,computed"`
	CreatedOn   types.String                                          `tfsdk:"created_on" json:"created_on,computed"`
	EnvVars     types.String                                          `tfsdk:"env_vars" json:"env_vars,computed"`
	Environment types.String                                          `tfsdk:"environment" json:"environment,computed"`
	IsSkipped   types.Bool                                            `tfsdk:"is_skipped" json:"is_skipped,computed"`
	LatestStage types.String                                          `tfsdk:"latest_stage" json:"latest_stage,computed"`
	ModifiedOn  types.String                                          `tfsdk:"modified_on" json:"modified_on,computed"`
	ProjectID   types.String                                          `tfsdk:"project_id" json:"project_id,computed"`
	ProjectName types.String                                          `tfsdk:"project_name" json:"project_name,computed"`
	ShortID     types.String                                          `tfsdk:"short_id" json:"short_id,computed"`
	Source      types.String                                          `tfsdk:"source" json:"source,computed"`
	Stages      *[]*PagesProjectLatestDeploymentStagesDataSourceModel `tfsdk:"stages" json:"stages,computed"`
	URL         types.String                                          `tfsdk:"url" json:"url,computed"`
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

type PagesProjectLatestDeploymentStagesDataSourceModel struct {
	EndedOn   types.String `tfsdk:"ended_on" json:"ended_on,computed"`
	Name      types.String `tfsdk:"name" json:"name"`
	StartedOn types.String `tfsdk:"started_on" json:"started_on,computed"`
	Status    types.String `tfsdk:"status" json:"status,computed"`
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

type PagesProjectStagesDataSourceModel struct {
	EndedOn   types.String `tfsdk:"ended_on" json:"ended_on,computed"`
	Name      types.String `tfsdk:"name" json:"name"`
	StartedOn types.String `tfsdk:"started_on" json:"started_on,computed"`
	Status    types.String `tfsdk:"status" json:"status,computed"`
}

type PagesProjectFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
