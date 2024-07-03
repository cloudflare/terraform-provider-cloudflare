// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_project

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PagesProjectResultEnvelope struct {
	Result PagesProjectModel `json:"result,computed"`
}

type PagesProjectModel struct {
	AccountID           types.String                          `tfsdk:"account_id" path:"account_id"`
	ProjectName         types.String                          `tfsdk:"project_name" path:"project_name"`
	BuildConfig         *PagesProjectBuildConfigModel         `tfsdk:"build_config" json:"build_config"`
	CanonicalDeployment *PagesProjectCanonicalDeploymentModel `tfsdk:"canonical_deployment" json:"canonical_deployment"`
	DeploymentConfigs   *PagesProjectDeploymentConfigsModel   `tfsdk:"deployment_configs" json:"deployment_configs"`
	LatestDeployment    *PagesProjectLatestDeploymentModel    `tfsdk:"latest_deployment" json:"latest_deployment"`
	Name                types.String                          `tfsdk:"name" json:"name"`
	ProductionBranch    types.String                          `tfsdk:"production_branch" json:"production_branch"`
	ID                  types.String                          `tfsdk:"id" json:"id,computed"`
	CreatedOn           types.String                          `tfsdk:"created_on" json:"created_on,computed"`
	Domains             *[]types.String                       `tfsdk:"domains" json:"domains,computed"`
	Source              types.String                          `tfsdk:"source" json:"source,computed"`
	Subdomain           types.String                          `tfsdk:"subdomain" json:"subdomain,computed"`
}

type PagesProjectBuildConfigModel struct {
	BuildCaching      types.Bool   `tfsdk:"build_caching" json:"build_caching"`
	BuildCommand      types.String `tfsdk:"build_command" json:"build_command"`
	DestinationDir    types.String `tfsdk:"destination_dir" json:"destination_dir"`
	RootDir           types.String `tfsdk:"root_dir" json:"root_dir"`
	WebAnalyticsTag   types.String `tfsdk:"web_analytics_tag" json:"web_analytics_tag"`
	WebAnalyticsToken types.String `tfsdk:"web_analytics_token" json:"web_analytics_token"`
}

type PagesProjectCanonicalDeploymentModel struct {
	ID          types.String                                   `tfsdk:"id" json:"id,computed"`
	Aliases     *[]types.String                                `tfsdk:"aliases" json:"aliases,computed"`
	BuildConfig types.String                                   `tfsdk:"build_config" json:"build_config,computed"`
	CreatedOn   types.String                                   `tfsdk:"created_on" json:"created_on,computed"`
	EnvVars     types.String                                   `tfsdk:"env_vars" json:"env_vars,computed"`
	Environment types.String                                   `tfsdk:"environment" json:"environment,computed"`
	IsSkipped   types.Bool                                     `tfsdk:"is_skipped" json:"is_skipped,computed"`
	LatestStage types.String                                   `tfsdk:"latest_stage" json:"latest_stage,computed"`
	ModifiedOn  types.String                                   `tfsdk:"modified_on" json:"modified_on,computed"`
	ProjectID   types.String                                   `tfsdk:"project_id" json:"project_id,computed"`
	ProjectName types.String                                   `tfsdk:"project_name" json:"project_name,computed"`
	ShortID     types.String                                   `tfsdk:"short_id" json:"short_id,computed"`
	Source      types.String                                   `tfsdk:"source" json:"source,computed"`
	Stages      *[]*PagesProjectCanonicalDeploymentStagesModel `tfsdk:"stages" json:"stages,computed"`
	URL         types.String                                   `tfsdk:"url" json:"url,computed"`
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

type PagesProjectCanonicalDeploymentStagesModel struct {
	EndedOn   types.String `tfsdk:"ended_on" json:"ended_on,computed"`
	Name      types.String `tfsdk:"name" json:"name"`
	StartedOn types.String `tfsdk:"started_on" json:"started_on,computed"`
	Status    types.String `tfsdk:"status" json:"status,computed"`
}

type PagesProjectDeploymentConfigsModel struct {
	Preview    *PagesProjectDeploymentConfigsPreviewModel    `tfsdk:"preview" json:"preview"`
	Production *PagesProjectDeploymentConfigsProductionModel `tfsdk:"production" json:"production"`
}

type PagesProjectDeploymentConfigsPreviewModel struct {
	AIBindings              *PagesProjectDeploymentConfigsPreviewAIBindingsModel              `tfsdk:"ai_bindings" json:"ai_bindings"`
	AnalyticsEngineDatasets *PagesProjectDeploymentConfigsPreviewAnalyticsEngineDatasetsModel `tfsdk:"analytics_engine_datasets" json:"analytics_engine_datasets"`
	Browsers                *PagesProjectDeploymentConfigsPreviewBrowsersModel                `tfsdk:"browsers" json:"browsers"`
	CompatibilityDate       types.String                                                      `tfsdk:"compatibility_date" json:"compatibility_date"`
	CompatibilityFlags      *[]types.String                                                   `tfsdk:"compatibility_flags" json:"compatibility_flags"`
	D1Databases             *PagesProjectDeploymentConfigsPreviewD1DatabasesModel             `tfsdk:"d1_databases" json:"d1_databases"`
	DurableObjectNamespaces *PagesProjectDeploymentConfigsPreviewDurableObjectNamespacesModel `tfsdk:"durable_object_namespaces" json:"durable_object_namespaces"`
	EnvVars                 *PagesProjectDeploymentConfigsPreviewEnvVarsModel                 `tfsdk:"env_vars" json:"env_vars"`
	HyperdriveBindings      *PagesProjectDeploymentConfigsPreviewHyperdriveBindingsModel      `tfsdk:"hyperdrive_bindings" json:"hyperdrive_bindings"`
	KVNamespaces            *PagesProjectDeploymentConfigsPreviewKVNamespacesModel            `tfsdk:"kv_namespaces" json:"kv_namespaces"`
	MTLSCertificates        *PagesProjectDeploymentConfigsPreviewMTLSCertificatesModel        `tfsdk:"mtls_certificates" json:"mtls_certificates"`
	Placement               *PagesProjectDeploymentConfigsPreviewPlacementModel               `tfsdk:"placement" json:"placement"`
	QueueProducers          *PagesProjectDeploymentConfigsPreviewQueueProducersModel          `tfsdk:"queue_producers" json:"queue_producers"`
	R2Buckets               *PagesProjectDeploymentConfigsPreviewR2BucketsModel               `tfsdk:"r2_buckets" json:"r2_buckets"`
	Services                *PagesProjectDeploymentConfigsPreviewServicesModel                `tfsdk:"services" json:"services"`
	VectorizeBindings       *PagesProjectDeploymentConfigsPreviewVectorizeBindingsModel       `tfsdk:"vectorize_bindings" json:"vectorize_bindings"`
}

type PagesProjectDeploymentConfigsPreviewAIBindingsModel struct {
	AIBinding *PagesProjectDeploymentConfigsPreviewAIBindingsAIBindingModel `tfsdk:"ai_binding" json:"AI_BINDING"`
}

type PagesProjectDeploymentConfigsPreviewAIBindingsAIBindingModel struct {
	ProjectID types.String `tfsdk:"project_id" json:"project_id"`
}

type PagesProjectDeploymentConfigsPreviewAnalyticsEngineDatasetsModel struct {
	AnalyticsEngineBinding *PagesProjectDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBindingModel `tfsdk:"analytics_engine_binding" json:"ANALYTICS_ENGINE_BINDING"`
}

type PagesProjectDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBindingModel struct {
	Dataset types.String `tfsdk:"dataset" json:"dataset"`
}

type PagesProjectDeploymentConfigsPreviewBrowsersModel struct {
	Browser types.String `tfsdk:"browser" json:"BROWSER"`
}

type PagesProjectDeploymentConfigsPreviewD1DatabasesModel struct {
	D1Binding *PagesProjectDeploymentConfigsPreviewD1DatabasesD1BindingModel `tfsdk:"d1_binding" json:"D1_BINDING"`
}

type PagesProjectDeploymentConfigsPreviewD1DatabasesD1BindingModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type PagesProjectDeploymentConfigsPreviewDurableObjectNamespacesModel struct {
	DoBinding *PagesProjectDeploymentConfigsPreviewDurableObjectNamespacesDoBindingModel `tfsdk:"do_binding" json:"DO_BINDING"`
}

type PagesProjectDeploymentConfigsPreviewDurableObjectNamespacesDoBindingModel struct {
	NamespaceID types.String `tfsdk:"namespace_id" json:"namespace_id"`
}

type PagesProjectDeploymentConfigsPreviewEnvVarsModel struct {
	EnvironmentVariable *PagesProjectDeploymentConfigsPreviewEnvVarsEnvironmentVariableModel `tfsdk:"environment_variable" json:"ENVIRONMENT_VARIABLE"`
}

type PagesProjectDeploymentConfigsPreviewEnvVarsEnvironmentVariableModel struct {
	Type  types.String `tfsdk:"type" json:"type"`
	Value types.String `tfsdk:"value" json:"value"`
}

type PagesProjectDeploymentConfigsPreviewHyperdriveBindingsModel struct {
	Hyperdrive *PagesProjectDeploymentConfigsPreviewHyperdriveBindingsHyperdriveModel `tfsdk:"hyperdrive" json:"HYPERDRIVE"`
}

type PagesProjectDeploymentConfigsPreviewHyperdriveBindingsHyperdriveModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type PagesProjectDeploymentConfigsPreviewKVNamespacesModel struct {
	KVBinding *PagesProjectDeploymentConfigsPreviewKVNamespacesKVBindingModel `tfsdk:"kv_binding" json:"KV_BINDING"`
}

type PagesProjectDeploymentConfigsPreviewKVNamespacesKVBindingModel struct {
	NamespaceID types.String `tfsdk:"namespace_id" json:"namespace_id"`
}

type PagesProjectDeploymentConfigsPreviewMTLSCertificatesModel struct {
	MTLS *PagesProjectDeploymentConfigsPreviewMTLSCertificatesMTLSModel `tfsdk:"mtls" json:"MTLS"`
}

type PagesProjectDeploymentConfigsPreviewMTLSCertificatesMTLSModel struct {
	CertificateID types.String `tfsdk:"certificate_id" json:"certificate_id"`
}

type PagesProjectDeploymentConfigsPreviewPlacementModel struct {
	Mode types.String `tfsdk:"mode" json:"mode"`
}

type PagesProjectDeploymentConfigsPreviewQueueProducersModel struct {
	QueueProducerBinding *PagesProjectDeploymentConfigsPreviewQueueProducersQueueProducerBindingModel `tfsdk:"queue_producer_binding" json:"QUEUE_PRODUCER_BINDING"`
}

type PagesProjectDeploymentConfigsPreviewQueueProducersQueueProducerBindingModel struct {
	Name types.String `tfsdk:"name" json:"name"`
}

type PagesProjectDeploymentConfigsPreviewR2BucketsModel struct {
	R2Binding *PagesProjectDeploymentConfigsPreviewR2BucketsR2BindingModel `tfsdk:"r2_binding" json:"R2_BINDING"`
}

type PagesProjectDeploymentConfigsPreviewR2BucketsR2BindingModel struct {
	Jurisdiction types.String `tfsdk:"jurisdiction" json:"jurisdiction"`
	Name         types.String `tfsdk:"name" json:"name"`
}

type PagesProjectDeploymentConfigsPreviewServicesModel struct {
	ServiceBinding *PagesProjectDeploymentConfigsPreviewServicesServiceBindingModel `tfsdk:"service_binding" json:"SERVICE_BINDING"`
}

type PagesProjectDeploymentConfigsPreviewServicesServiceBindingModel struct {
	Entrypoint  types.String `tfsdk:"entrypoint" json:"entrypoint"`
	Environment types.String `tfsdk:"environment" json:"environment"`
	Service     types.String `tfsdk:"service" json:"service"`
}

type PagesProjectDeploymentConfigsPreviewVectorizeBindingsModel struct {
	Vectorize *PagesProjectDeploymentConfigsPreviewVectorizeBindingsVectorizeModel `tfsdk:"vectorize" json:"VECTORIZE"`
}

type PagesProjectDeploymentConfigsPreviewVectorizeBindingsVectorizeModel struct {
	IndexName types.String `tfsdk:"index_name" json:"index_name"`
}

type PagesProjectDeploymentConfigsProductionModel struct {
	AIBindings              *PagesProjectDeploymentConfigsProductionAIBindingsModel              `tfsdk:"ai_bindings" json:"ai_bindings"`
	AnalyticsEngineDatasets *PagesProjectDeploymentConfigsProductionAnalyticsEngineDatasetsModel `tfsdk:"analytics_engine_datasets" json:"analytics_engine_datasets"`
	Browsers                *PagesProjectDeploymentConfigsProductionBrowsersModel                `tfsdk:"browsers" json:"browsers"`
	CompatibilityDate       types.String                                                         `tfsdk:"compatibility_date" json:"compatibility_date"`
	CompatibilityFlags      *[]types.String                                                      `tfsdk:"compatibility_flags" json:"compatibility_flags"`
	D1Databases             *PagesProjectDeploymentConfigsProductionD1DatabasesModel             `tfsdk:"d1_databases" json:"d1_databases"`
	DurableObjectNamespaces *PagesProjectDeploymentConfigsProductionDurableObjectNamespacesModel `tfsdk:"durable_object_namespaces" json:"durable_object_namespaces"`
	EnvVars                 *PagesProjectDeploymentConfigsProductionEnvVarsModel                 `tfsdk:"env_vars" json:"env_vars"`
	HyperdriveBindings      *PagesProjectDeploymentConfigsProductionHyperdriveBindingsModel      `tfsdk:"hyperdrive_bindings" json:"hyperdrive_bindings"`
	KVNamespaces            *PagesProjectDeploymentConfigsProductionKVNamespacesModel            `tfsdk:"kv_namespaces" json:"kv_namespaces"`
	MTLSCertificates        *PagesProjectDeploymentConfigsProductionMTLSCertificatesModel        `tfsdk:"mtls_certificates" json:"mtls_certificates"`
	Placement               *PagesProjectDeploymentConfigsProductionPlacementModel               `tfsdk:"placement" json:"placement"`
	QueueProducers          *PagesProjectDeploymentConfigsProductionQueueProducersModel          `tfsdk:"queue_producers" json:"queue_producers"`
	R2Buckets               *PagesProjectDeploymentConfigsProductionR2BucketsModel               `tfsdk:"r2_buckets" json:"r2_buckets"`
	Services                *PagesProjectDeploymentConfigsProductionServicesModel                `tfsdk:"services" json:"services"`
	VectorizeBindings       *PagesProjectDeploymentConfigsProductionVectorizeBindingsModel       `tfsdk:"vectorize_bindings" json:"vectorize_bindings"`
}

type PagesProjectDeploymentConfigsProductionAIBindingsModel struct {
	AIBinding *PagesProjectDeploymentConfigsProductionAIBindingsAIBindingModel `tfsdk:"ai_binding" json:"AI_BINDING"`
}

type PagesProjectDeploymentConfigsProductionAIBindingsAIBindingModel struct {
	ProjectID types.String `tfsdk:"project_id" json:"project_id"`
}

type PagesProjectDeploymentConfigsProductionAnalyticsEngineDatasetsModel struct {
	AnalyticsEngineBinding *PagesProjectDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBindingModel `tfsdk:"analytics_engine_binding" json:"ANALYTICS_ENGINE_BINDING"`
}

type PagesProjectDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBindingModel struct {
	Dataset types.String `tfsdk:"dataset" json:"dataset"`
}

type PagesProjectDeploymentConfigsProductionBrowsersModel struct {
	Browser types.String `tfsdk:"browser" json:"BROWSER"`
}

type PagesProjectDeploymentConfigsProductionD1DatabasesModel struct {
	D1Binding *PagesProjectDeploymentConfigsProductionD1DatabasesD1BindingModel `tfsdk:"d1_binding" json:"D1_BINDING"`
}

type PagesProjectDeploymentConfigsProductionD1DatabasesD1BindingModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type PagesProjectDeploymentConfigsProductionDurableObjectNamespacesModel struct {
	DoBinding *PagesProjectDeploymentConfigsProductionDurableObjectNamespacesDoBindingModel `tfsdk:"do_binding" json:"DO_BINDING"`
}

type PagesProjectDeploymentConfigsProductionDurableObjectNamespacesDoBindingModel struct {
	NamespaceID types.String `tfsdk:"namespace_id" json:"namespace_id"`
}

type PagesProjectDeploymentConfigsProductionEnvVarsModel struct {
	EnvironmentVariable *PagesProjectDeploymentConfigsProductionEnvVarsEnvironmentVariableModel `tfsdk:"environment_variable" json:"ENVIRONMENT_VARIABLE"`
}

type PagesProjectDeploymentConfigsProductionEnvVarsEnvironmentVariableModel struct {
	Type  types.String `tfsdk:"type" json:"type"`
	Value types.String `tfsdk:"value" json:"value"`
}

type PagesProjectDeploymentConfigsProductionHyperdriveBindingsModel struct {
	Hyperdrive *PagesProjectDeploymentConfigsProductionHyperdriveBindingsHyperdriveModel `tfsdk:"hyperdrive" json:"HYPERDRIVE"`
}

type PagesProjectDeploymentConfigsProductionHyperdriveBindingsHyperdriveModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type PagesProjectDeploymentConfigsProductionKVNamespacesModel struct {
	KVBinding *PagesProjectDeploymentConfigsProductionKVNamespacesKVBindingModel `tfsdk:"kv_binding" json:"KV_BINDING"`
}

type PagesProjectDeploymentConfigsProductionKVNamespacesKVBindingModel struct {
	NamespaceID types.String `tfsdk:"namespace_id" json:"namespace_id"`
}

type PagesProjectDeploymentConfigsProductionMTLSCertificatesModel struct {
	MTLS *PagesProjectDeploymentConfigsProductionMTLSCertificatesMTLSModel `tfsdk:"mtls" json:"MTLS"`
}

type PagesProjectDeploymentConfigsProductionMTLSCertificatesMTLSModel struct {
	CertificateID types.String `tfsdk:"certificate_id" json:"certificate_id"`
}

type PagesProjectDeploymentConfigsProductionPlacementModel struct {
	Mode types.String `tfsdk:"mode" json:"mode"`
}

type PagesProjectDeploymentConfigsProductionQueueProducersModel struct {
	QueueProducerBinding *PagesProjectDeploymentConfigsProductionQueueProducersQueueProducerBindingModel `tfsdk:"queue_producer_binding" json:"QUEUE_PRODUCER_BINDING"`
}

type PagesProjectDeploymentConfigsProductionQueueProducersQueueProducerBindingModel struct {
	Name types.String `tfsdk:"name" json:"name"`
}

type PagesProjectDeploymentConfigsProductionR2BucketsModel struct {
	R2Binding *PagesProjectDeploymentConfigsProductionR2BucketsR2BindingModel `tfsdk:"r2_binding" json:"R2_BINDING"`
}

type PagesProjectDeploymentConfigsProductionR2BucketsR2BindingModel struct {
	Jurisdiction types.String `tfsdk:"jurisdiction" json:"jurisdiction"`
	Name         types.String `tfsdk:"name" json:"name"`
}

type PagesProjectDeploymentConfigsProductionServicesModel struct {
	ServiceBinding *PagesProjectDeploymentConfigsProductionServicesServiceBindingModel `tfsdk:"service_binding" json:"SERVICE_BINDING"`
}

type PagesProjectDeploymentConfigsProductionServicesServiceBindingModel struct {
	Entrypoint  types.String `tfsdk:"entrypoint" json:"entrypoint"`
	Environment types.String `tfsdk:"environment" json:"environment"`
	Service     types.String `tfsdk:"service" json:"service"`
}

type PagesProjectDeploymentConfigsProductionVectorizeBindingsModel struct {
	Vectorize *PagesProjectDeploymentConfigsProductionVectorizeBindingsVectorizeModel `tfsdk:"vectorize" json:"VECTORIZE"`
}

type PagesProjectDeploymentConfigsProductionVectorizeBindingsVectorizeModel struct {
	IndexName types.String `tfsdk:"index_name" json:"index_name"`
}

type PagesProjectLatestDeploymentModel struct {
	ID          types.String                                `tfsdk:"id" json:"id,computed"`
	Aliases     *[]types.String                             `tfsdk:"aliases" json:"aliases,computed"`
	BuildConfig types.String                                `tfsdk:"build_config" json:"build_config,computed"`
	CreatedOn   types.String                                `tfsdk:"created_on" json:"created_on,computed"`
	EnvVars     types.String                                `tfsdk:"env_vars" json:"env_vars,computed"`
	Environment types.String                                `tfsdk:"environment" json:"environment,computed"`
	IsSkipped   types.Bool                                  `tfsdk:"is_skipped" json:"is_skipped,computed"`
	LatestStage types.String                                `tfsdk:"latest_stage" json:"latest_stage,computed"`
	ModifiedOn  types.String                                `tfsdk:"modified_on" json:"modified_on,computed"`
	ProjectID   types.String                                `tfsdk:"project_id" json:"project_id,computed"`
	ProjectName types.String                                `tfsdk:"project_name" json:"project_name,computed"`
	ShortID     types.String                                `tfsdk:"short_id" json:"short_id,computed"`
	Source      types.String                                `tfsdk:"source" json:"source,computed"`
	Stages      *[]*PagesProjectLatestDeploymentStagesModel `tfsdk:"stages" json:"stages,computed"`
	URL         types.String                                `tfsdk:"url" json:"url,computed"`
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

type PagesProjectLatestDeploymentStagesModel struct {
	EndedOn   types.String `tfsdk:"ended_on" json:"ended_on,computed"`
	Name      types.String `tfsdk:"name" json:"name"`
	StartedOn types.String `tfsdk:"started_on" json:"started_on,computed"`
	Status    types.String `tfsdk:"status" json:"status,computed"`
}
