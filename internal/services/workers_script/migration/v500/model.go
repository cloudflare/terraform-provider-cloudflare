package v500

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x)
// ============================================================================

// SourceWorkerScriptModel represents the v4 cloudflare_worker_script state (schema_version=0).
// Resource type: cloudflare_worker_script (singular) or cloudflare_workers_script (plural)
//
// Key differences from v5:
// - "name" field instead of "script_name"
// - "module" boolean instead of "main_module"/"body_part" strings
// - 10 separate binding block arrays instead of unified "bindings" list
// - "placement" is an array (single element) instead of an object
// - "tags" and "dispatch_namespace" exist (removed in v5)
type SourceWorkerScriptModel struct {
	ID        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	AccountID types.String `tfsdk:"account_id"`
	Content   types.String `tfsdk:"content"`
	Module    types.Bool   `tfsdk:"module"`
	Tags      types.Set    `tfsdk:"tags"`

	DispatchNamespace types.String `tfsdk:"dispatch_namespace"`

	// V4 binding blocks stored as arrays in SDKv2 state
	PlainTextBinding        []SourcePlainTextBindingModel        `tfsdk:"plain_text_binding"`
	SecretTextBinding       []SourceSecretTextBindingModel       `tfsdk:"secret_text_binding"`
	KVNamespaceBinding      []SourceKVNamespaceBindingModel      `tfsdk:"kv_namespace_binding"`
	WebassemblyBinding      []SourceWebassemblyBindingModel      `tfsdk:"webassembly_binding"`
	ServiceBinding          []SourceServiceBindingModel          `tfsdk:"service_binding"`
	R2BucketBinding         []SourceR2BucketBindingModel         `tfsdk:"r2_bucket_binding"`
	AnalyticsEngineBinding  []SourceAnalyticsEngineBindingModel  `tfsdk:"analytics_engine_binding"`
	QueueBinding            []SourceQueueBindingModel            `tfsdk:"queue_binding"`
	D1DatabaseBinding       []SourceD1DatabaseBindingModel       `tfsdk:"d1_database_binding"`
	HyperdriveConfigBinding []SourceHyperdriveConfigBindingModel `tfsdk:"hyperdrive_config_binding"`

	// V4 placement is a single-element array
	Placement []SourcePlacementModel `tfsdk:"placement"`
}

// V4 binding sub-models

type SourcePlainTextBindingModel struct {
	Name types.String `tfsdk:"name"`
	Text types.String `tfsdk:"text"`
}

type SourceSecretTextBindingModel struct {
	Name types.String `tfsdk:"name"`
	Text types.String `tfsdk:"text"`
}

type SourceKVNamespaceBindingModel struct {
	Name        types.String `tfsdk:"name"`
	NamespaceID types.String `tfsdk:"namespace_id"`
}

type SourceWebassemblyBindingModel struct {
	Name   types.String `tfsdk:"name"`
	Module types.String `tfsdk:"module"` // renamed to "part" in v5
}

type SourceServiceBindingModel struct {
	Name        types.String `tfsdk:"name"`
	Service     types.String `tfsdk:"service"`
	Environment types.String `tfsdk:"environment"`
}

type SourceR2BucketBindingModel struct {
	Name       types.String `tfsdk:"name"`
	BucketName types.String `tfsdk:"bucket_name"`
}

type SourceAnalyticsEngineBindingModel struct {
	Name    types.String `tfsdk:"name"`
	Dataset types.String `tfsdk:"dataset"`
}

type SourceQueueBindingModel struct {
	Binding types.String `tfsdk:"binding"` // renamed to "name" in v5
	Queue   types.String `tfsdk:"queue"`   // renamed to "queue_name" in v5
}

type SourceD1DatabaseBindingModel struct {
	Name       types.String `tfsdk:"name"`
	DatabaseID types.String `tfsdk:"database_id"` // renamed to "id" in v5
}

type SourceHyperdriveConfigBindingModel struct {
	Binding types.String `tfsdk:"binding"` // renamed to "name" in v5
	ID      types.String `tfsdk:"id"`
}

type SourcePlacementModel struct {
	Mode types.String `tfsdk:"mode"`
}

// ============================================================================
// Target Models (Current Provider - v5.x+)
// ============================================================================

// TargetWorkersScriptModel represents the v5 cloudflare_workers_script state.
// Must match the current schema exactly for resp.TargetState.Set() to work.
type TargetWorkersScriptModel struct {
	ID               types.String                                           `tfsdk:"id"`
	ScriptName       types.String                                           `tfsdk:"script_name"`
	AccountID        types.String                                           `tfsdk:"account_id"`
	Content          types.String                                           `tfsdk:"content"`
	ContentFile      types.String                                           `tfsdk:"content_file"`
	ContentSHA256    types.String                                           `tfsdk:"content_sha256"`
	ContentType      types.String                                           `tfsdk:"content_type"`
	CreatedOn        timetypes.RFC3339                                      `tfsdk:"created_on"`
	Etag             types.String                                           `tfsdk:"etag"`
	HasAssets        types.Bool                                             `tfsdk:"has_assets"`
	HasModules       types.Bool                                             `tfsdk:"has_modules"`
	LastDeployedFrom types.String                                           `tfsdk:"last_deployed_from"`
	MigrationTag     types.String                                           `tfsdk:"migration_tag"`
	ModifiedOn       timetypes.RFC3339                                      `tfsdk:"modified_on"`
	PlacementMode    types.String                                           `tfsdk:"placement_mode"`
	PlacementStatus  types.String                                           `tfsdk:"placement_status"`
	StartupTimeMs    types.Int64                                            `tfsdk:"startup_time_ms"`
	Handlers         customfield.List[types.String]                         `tfsdk:"handlers"`
	NamedHandlers    customfield.NestedObjectList[TargetNamedHandlersModel] `tfsdk:"named_handlers"`

	// Metadata fields (embedded in WorkersScriptModel)
	Assets             *TargetAssetsModel                                    `tfsdk:"assets"`
	Bindings           customfield.NestedObjectList[TargetBindingsModel]     `tfsdk:"bindings"`
	BodyPart           types.String                                          `tfsdk:"body_part"`
	CompatibilityDate  types.String                                          `tfsdk:"compatibility_date"`
	CompatibilityFlags customfield.Set[types.String]                         `tfsdk:"compatibility_flags"`
	KeepAssets         types.Bool                                            `tfsdk:"keep_assets"`
	KeepBindings       *[]types.String                                       `tfsdk:"keep_bindings"`
	Limits             *TargetLimitsModel                                    `tfsdk:"limits"`
	Logpush            types.Bool                                            `tfsdk:"logpush"`
	MainModule         types.String                                          `tfsdk:"main_module"`
	Migrations         customfield.NestedObject[TargetMigrationsModel]       `tfsdk:"migrations"`
	Observability      *TargetObservabilityModel                             `tfsdk:"observability"`
	Placement          customfield.NestedObject[TargetPlacementModel]        `tfsdk:"placement"`
	TailConsumers      customfield.NestedObjectSet[TargetTailConsumersModel] `tfsdk:"tail_consumers"`
	UsageModel         types.String                                          `tfsdk:"usage_model"`
}

// TargetBindingsModel mirrors WorkersScriptMetadataBindingsModel.
// All fields except name and type are Optional — most will be null for any given binding.
type TargetBindingsModel struct {
	Name                        types.String                  `tfsdk:"name"`
	Type                        types.String                  `tfsdk:"type"`
	InstanceName                types.String                  `tfsdk:"instance_name"`
	Dataset                     types.String                  `tfsdk:"dataset"`
	ID                          types.String                  `tfsdk:"id"`
	Outbound                    *TargetBindingsOutboundModel  `tfsdk:"outbound"`
	ClassName                   types.String                  `tfsdk:"class_name"`
	NamespaceID                 types.String                  `tfsdk:"namespace_id"`
	ScriptName                  types.String                  `tfsdk:"script_name"`
	Json                        types.String                  `tfsdk:"json"`
	CertificateID               types.String                  `tfsdk:"certificate_id"`
	Text                        types.String                  `tfsdk:"text"`
	Pipeline                    types.String                  `tfsdk:"pipeline"`
	QueueName                   types.String                  `tfsdk:"queue_name"`
	Simple                      *TargetBindingsSimpleModel    `tfsdk:"simple"`
	BucketName                  types.String                  `tfsdk:"bucket_name"`
	Jurisdiction                types.String                  `tfsdk:"jurisdiction"`
	IndexName                   types.String                  `tfsdk:"index_name"`
	SecretName                  types.String                  `tfsdk:"secret_name"`
	StoreID                     types.String                  `tfsdk:"store_id"`
	Algorithm                   jsontypes.Normalized          `tfsdk:"algorithm"`
	Format                      types.String                  `tfsdk:"format"`
	Usages                      customfield.Set[types.String] `tfsdk:"usages"`
	KeyBase64                   types.String                  `tfsdk:"key_base64"`
	KeyJwk                      jsontypes.Normalized          `tfsdk:"key_jwk"`
	WorkflowName                types.String                  `tfsdk:"workflow_name"`
	VersionID                   types.String                  `tfsdk:"version_id"`
	Part                        types.String                  `tfsdk:"part"`
	Namespace                   types.String                  `tfsdk:"namespace"`
	Environment                 types.String                  `tfsdk:"environment"`
	OldName                     types.String                  `tfsdk:"old_name"`
	AllowedDestinationAddresses *[]types.String               `tfsdk:"allowed_destination_addresses"`
	AllowedSenderAddresses      *[]types.String               `tfsdk:"allowed_sender_addresses"`
	DestinationAddress          types.String                  `tfsdk:"destination_address"`
	Service                     types.String                  `tfsdk:"service"`
	DispatchNamespace           types.String                  `tfsdk:"dispatch_namespace"`
	Entrypoint                  types.String                  `tfsdk:"entrypoint"`
	ServiceID                   types.String                  `tfsdk:"service_id"`
	NetworkID                   types.String                  `tfsdk:"network_id"`
	TunnelID                    types.String                  `tfsdk:"tunnel_id"`
}

type TargetBindingsOutboundModel struct {
	Params *[]types.String                    `tfsdk:"params"`
	Worker *TargetBindingsOutboundWorkerModel `tfsdk:"worker"`
}

type TargetBindingsOutboundWorkerModel struct {
	Environment types.String `tfsdk:"environment"`
	Service     types.String `tfsdk:"service"`
}

type TargetBindingsSimpleModel struct {
	Limit  types.Float64 `tfsdk:"limit"`
	Period types.Int64   `tfsdk:"period"`
}

type TargetAssetsModel struct {
	Config              *TargetAssetsConfigModel `tfsdk:"config"`
	JWT                 types.String             `tfsdk:"jwt"`
	Directory           types.String             `tfsdk:"directory"`
	AssetManifestSHA256 types.String             `tfsdk:"asset_manifest_sha256"`
}

type TargetAssetsConfigModel struct {
	Headers          types.String                       `tfsdk:"headers"`
	Redirects        types.String                       `tfsdk:"redirects"`
	HTMLHandling     types.String                       `tfsdk:"html_handling"`
	NotFoundHandling types.String                       `tfsdk:"not_found_handling"`
	RunWorkerFirst   customfield.NormalizedDynamicValue `tfsdk:"run_worker_first"`
	ServeDirectly    types.Bool                         `tfsdk:"serve_directly"`
}

type TargetLimitsModel struct {
	CPUMs types.Int64 `tfsdk:"cpu_ms"`
}

type TargetMigrationsModel struct {
	DeletedClasses     *[]types.String                      `tfsdk:"deleted_classes"`
	NewClasses         *[]types.String                      `tfsdk:"new_classes"`
	NewSqliteClasses   *[]types.String                      `tfsdk:"new_sqlite_classes"`
	NewTag             types.String                         `tfsdk:"new_tag"`
	OldTag             types.String                         `tfsdk:"old_tag"`
	RenamedClasses     *[]*TargetMigrationsRenamedModel     `tfsdk:"renamed_classes"`
	TransferredClasses *[]*TargetMigrationsTransferredModel `tfsdk:"transferred_classes"`
	Steps              *[]*TargetMigrationsStepsModel       `tfsdk:"steps"`
}

type TargetMigrationsRenamedModel struct {
	From types.String `tfsdk:"from"`
	To   types.String `tfsdk:"to"`
}

type TargetMigrationsTransferredModel struct {
	From       types.String `tfsdk:"from"`
	FromScript types.String `tfsdk:"from_script"`
	To         types.String `tfsdk:"to"`
}

type TargetMigrationsStepsModel struct {
	DeletedClasses     *[]types.String                      `tfsdk:"deleted_classes"`
	NewClasses         *[]types.String                      `tfsdk:"new_classes"`
	NewSqliteClasses   *[]types.String                      `tfsdk:"new_sqlite_classes"`
	RenamedClasses     *[]*TargetMigrationsRenamedModel     `tfsdk:"renamed_classes"`
	TransferredClasses *[]*TargetMigrationsTransferredModel `tfsdk:"transferred_classes"`
}

type TargetObservabilityModel struct {
	Enabled          types.Bool                      `tfsdk:"enabled"`
	HeadSamplingRate types.Float64                   `tfsdk:"head_sampling_rate"`
	Logs             *TargetObservabilityLogsModel   `tfsdk:"logs"`
	Traces           *TargetObservabilityTracesModel `tfsdk:"traces"`
}

type TargetObservabilityLogsModel struct {
	Enabled          types.Bool      `tfsdk:"enabled"`
	InvocationLogs   types.Bool      `tfsdk:"invocation_logs"`
	Destinations     *[]types.String `tfsdk:"destinations"`
	HeadSamplingRate types.Float64   `tfsdk:"head_sampling_rate"`
	Persist          types.Bool      `tfsdk:"persist"`
}

type TargetObservabilityTracesModel struct {
	Destinations     *[]types.String `tfsdk:"destinations"`
	Enabled          types.Bool      `tfsdk:"enabled"`
	HeadSamplingRate types.Float64   `tfsdk:"head_sampling_rate"`
	Persist          types.Bool      `tfsdk:"persist"`
}

type TargetPlacementModel struct {
	Mode           types.String                   `tfsdk:"mode"`
	LastAnalyzedAt timetypes.RFC3339              `tfsdk:"last_analyzed_at"`
	Status         types.String                   `tfsdk:"status"`
	Region         types.String                   `tfsdk:"region"`
	Hostname       types.String                   `tfsdk:"hostname"`
	Host           types.String                   `tfsdk:"host"`
	Target         *[]*TargetPlacementTargetModel `tfsdk:"target"`
}

type TargetPlacementTargetModel struct {
	Region   types.String `tfsdk:"region"`
	Hostname types.String `tfsdk:"hostname"`
	Host     types.String `tfsdk:"host"`
}

type TargetTailConsumersModel struct {
	Service     types.String `tfsdk:"service"`
	Environment types.String `tfsdk:"environment"`
	Namespace   types.String `tfsdk:"namespace"`
}

type TargetNamedHandlersModel struct {
	Handlers customfield.List[types.String] `tfsdk:"handlers"`
	Name     types.String                   `tfsdk:"name"`
}
