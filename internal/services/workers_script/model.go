// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_script

import (
	"bytes"
	"errors"
	"mime/multipart"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apiform"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersScriptResultEnvelope struct {
	Result WorkersScriptModel `json:"result"`
}

type WorkersScriptModel struct {
	ID                 types.String                                                  `tfsdk:"id" json:"-,computed"`
	ScriptName         types.String                                                  `tfsdk:"script_name" path:"script_name,required"`
	AccountID          types.String                                                  `tfsdk:"account_id" path:"account_id,required"`
	Metadata           *WorkersScriptMetadataModel                                   `tfsdk:"metadata" json:"metadata,required,no_refresh"`
	Files              *[]types.String                                               `tfsdk:"files" json:"files,optional,no_refresh"`
	CompatibilityDate  types.String                                                  `tfsdk:"compatibility_date" json:"compatibility_date,computed,no_refresh"`
	CreatedOn          timetypes.RFC3339                                             `tfsdk:"created_on" json:"created_on,computed,no_refresh" format:"date-time"`
	EntryPoint         types.String                                                  `tfsdk:"entry_point" json:"entry_point,computed,no_refresh"`
	Etag               types.String                                                  `tfsdk:"etag" json:"etag,computed,no_refresh"`
	HasAssets          types.Bool                                                    `tfsdk:"has_assets" json:"has_assets,computed,no_refresh"`
	HasModules         types.Bool                                                    `tfsdk:"has_modules" json:"has_modules,computed,no_refresh"`
	LastDeployedFrom   types.String                                                  `tfsdk:"last_deployed_from" json:"last_deployed_from,computed,no_refresh"`
	Logpush            types.Bool                                                    `tfsdk:"logpush" json:"logpush,computed,no_refresh"`
	MigrationTag       types.String                                                  `tfsdk:"migration_tag" json:"migration_tag,computed,no_refresh"`
	ModifiedOn         timetypes.RFC3339                                             `tfsdk:"modified_on" json:"modified_on,computed,no_refresh" format:"date-time"`
	PlacementMode      types.String                                                  `tfsdk:"placement_mode" json:"placement_mode,computed,no_refresh"`
	PlacementStatus    types.String                                                  `tfsdk:"placement_status" json:"placement_status,computed,no_refresh"`
	StartupTimeMs      types.Int64                                                   `tfsdk:"startup_time_ms" json:"startup_time_ms,computed,no_refresh"`
	Tag                types.String                                                  `tfsdk:"tag" json:"tag,computed,no_refresh"`
	UsageModel         types.String                                                  `tfsdk:"usage_model" json:"usage_model,computed,no_refresh"`
	CompatibilityFlags customfield.Set[types.String]                                 `tfsdk:"compatibility_flags" json:"compatibility_flags,computed,no_refresh"`
	Handlers           customfield.List[types.String]                                `tfsdk:"handlers" json:"handlers,computed,no_refresh"`
	Tags               customfield.Set[types.String]                                 `tfsdk:"tags" json:"tags,computed,no_refresh"`
	NamedHandlers      customfield.NestedObjectList[WorkersScriptNamedHandlersModel] `tfsdk:"named_handlers" json:"named_handlers,computed,no_refresh"`
	Observability      customfield.NestedObject[WorkersScriptObservabilityModel]     `tfsdk:"observability" json:"observability,computed,no_refresh"`
	Placement          customfield.NestedObject[WorkersScriptPlacementModel]         `tfsdk:"placement" json:"placement,computed,no_refresh"`
	TailConsumers      customfield.NestedObjectSet[WorkersScriptTailConsumersModel]  `tfsdk:"tail_consumers" json:"tail_consumers,computed,no_refresh"`
}

func (r WorkersScriptModel) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		if e := writer.Close(); e != nil {
			err = errors.Join(err, e)
		}
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type WorkersScriptMetadataModel struct {
	Assets             *WorkersScriptMetadataAssetsModel           `tfsdk:"assets" json:"assets,optional"`
	Bindings           *[]*WorkersScriptMetadataBindingsModel      `tfsdk:"bindings" json:"bindings,optional"`
	BodyPart           types.String                                `tfsdk:"body_part" json:"body_part,optional"`
	CompatibilityDate  types.String                                `tfsdk:"compatibility_date" json:"compatibility_date,optional"`
	CompatibilityFlags customfield.Set[types.String]               `tfsdk:"compatibility_flags" json:"compatibility_flags,computed_optional"`
	KeepAssets         types.Bool                                  `tfsdk:"keep_assets" json:"keep_assets,optional"`
	KeepBindings       *[]types.String                             `tfsdk:"keep_bindings" json:"keep_bindings,optional"`
	Limits             *WorkersScriptMetadataLimitsModel           `tfsdk:"limits" json:"limits,optional"`
	Logpush            types.Bool                                  `tfsdk:"logpush" json:"logpush,computed_optional"`
	MainModule         types.String                                `tfsdk:"main_module" json:"main_module,optional"`
	Migrations         *WorkersScriptMetadataMigrationsModel       `tfsdk:"migrations" json:"migrations,optional"`
	Observability      *WorkersScriptMetadataObservabilityModel    `tfsdk:"observability" json:"observability,optional"`
	Placement          *WorkersScriptMetadataPlacementModel        `tfsdk:"placement" json:"placement,optional"`
	Tags               *[]types.String                             `tfsdk:"tags" json:"tags,optional"`
	TailConsumers      *[]*WorkersScriptMetadataTailConsumersModel `tfsdk:"tail_consumers" json:"tail_consumers,optional"`
	UsageModel         types.String                                `tfsdk:"usage_model" json:"usage_model,computed_optional"`
}

type WorkersScriptMetadataAssetsModel struct {
	Config *WorkersScriptMetadataAssetsConfigModel `tfsdk:"config" json:"config,optional"`
	JWT    types.String                            `tfsdk:"jwt" json:"jwt,optional"`
}

type WorkersScriptMetadataAssetsConfigModel struct {
	Headers          types.String    `tfsdk:"headers" json:"_headers,optional"`
	Redirects        types.String    `tfsdk:"redirects" json:"_redirects,optional"`
	HTMLHandling     types.String    `tfsdk:"html_handling" json:"html_handling,optional"`
	NotFoundHandling types.String    `tfsdk:"not_found_handling" json:"not_found_handling,optional"`
	RunWorkerFirst   *[]types.String `tfsdk:"run_worker_first" json:"run_worker_first,optional"`
	ServeDirectly    types.Bool      `tfsdk:"serve_directly" json:"serve_directly,optional"`
}

type WorkersScriptMetadataBindingsModel struct {
	Name                        types.String                                `tfsdk:"name" json:"name,required"`
	Type                        types.String                                `tfsdk:"type" json:"type,required"`
	Dataset                     types.String                                `tfsdk:"dataset" json:"dataset,optional"`
	ID                          types.String                                `tfsdk:"id" json:"id,optional"`
	Part                        types.String                                `tfsdk:"part" json:"part,optional"`
	Namespace                   types.String                                `tfsdk:"namespace" json:"namespace,optional"`
	Outbound                    *WorkersScriptMetadataBindingsOutboundModel `tfsdk:"outbound" json:"outbound,optional"`
	ClassName                   types.String                                `tfsdk:"class_name" json:"class_name,computed_optional"`
	Environment                 types.String                                `tfsdk:"environment" json:"environment,optional"`
	NamespaceID                 types.String                                `tfsdk:"namespace_id" json:"namespace_id,computed_optional"`
	ScriptName                  types.String                                `tfsdk:"script_name" json:"script_name,computed_optional"`
	OldName                     types.String                                `tfsdk:"old_name" json:"old_name,optional"`
	VersionID                   types.String                                `tfsdk:"version_id" json:"version_id,computed_optional"`
	Json                        types.String                                `tfsdk:"json" json:"json,optional"`
	CertificateID               types.String                                `tfsdk:"certificate_id" json:"certificate_id,optional"`
	Text                        types.String                                `tfsdk:"text" json:"text,optional"`
	Pipeline                    types.String                                `tfsdk:"pipeline" json:"pipeline,optional"`
	QueueName                   types.String                                `tfsdk:"queue_name" json:"queue_name,optional"`
	BucketName                  types.String                                `tfsdk:"bucket_name" json:"bucket_name,optional"`
	Jurisdiction                types.String                                `tfsdk:"jurisdiction" json:"jurisdiction,optional"`
	AllowedDestinationAddresses *[]types.String                             `tfsdk:"allowed_destination_addresses" json:"allowed_destination_addresses,optional"`
	AllowedSenderAddresses      *[]types.String                             `tfsdk:"allowed_sender_addresses" json:"allowed_sender_addresses,optional"`
	DestinationAddress          types.String                                `tfsdk:"destination_address" json:"destination_address,optional"`
	Service                     types.String                                `tfsdk:"service" json:"service,optional"`
	IndexName                   types.String                                `tfsdk:"index_name" json:"index_name,optional"`
	SecretName                  types.String                                `tfsdk:"secret_name" json:"secret_name,optional"`
	StoreID                     types.String                                `tfsdk:"store_id" json:"store_id,optional"`
	Algorithm                   jsontypes.Normalized                        `tfsdk:"algorithm" json:"algorithm,optional"`
	Format                      types.String                                `tfsdk:"format" json:"format,optional"`
	Usages                      *[]types.String                             `tfsdk:"usages" json:"usages,optional"`
	KeyBase64                   types.String                                `tfsdk:"key_base64" json:"key_base64,optional"`
	KeyJwk                      jsontypes.Normalized                        `tfsdk:"key_jwk" json:"key_jwk,optional"`
	WorkflowName                types.String                                `tfsdk:"workflow_name" json:"workflow_name,optional"`
}

type WorkersScriptMetadataBindingsOutboundModel struct {
	Params *[]types.String                                   `tfsdk:"params" json:"params,optional"`
	Worker *WorkersScriptMetadataBindingsOutboundWorkerModel `tfsdk:"worker" json:"worker,optional"`
}

type WorkersScriptMetadataBindingsOutboundWorkerModel struct {
	Environment types.String `tfsdk:"environment" json:"environment,optional"`
	Service     types.String `tfsdk:"service" json:"service,optional"`
}

type WorkersScriptMetadataLimitsModel struct {
	CPUMs types.Int64 `tfsdk:"cpu_ms" json:"cpu_ms,optional"`
}

type WorkersScriptMetadataMigrationsModel struct {
	DeletedClasses     *[]types.String                                            `tfsdk:"deleted_classes" json:"deleted_classes,optional"`
	NewClasses         *[]types.String                                            `tfsdk:"new_classes" json:"new_classes,optional"`
	NewSqliteClasses   *[]types.String                                            `tfsdk:"new_sqlite_classes" json:"new_sqlite_classes,optional"`
	NewTag             types.String                                               `tfsdk:"new_tag" json:"new_tag,optional"`
	OldTag             types.String                                               `tfsdk:"old_tag" json:"old_tag,optional"`
	RenamedClasses     *[]*WorkersScriptMetadataMigrationsRenamedClassesModel     `tfsdk:"renamed_classes" json:"renamed_classes,optional"`
	TransferredClasses *[]*WorkersScriptMetadataMigrationsTransferredClassesModel `tfsdk:"transferred_classes" json:"transferred_classes,optional"`
	Steps              *[]*WorkersScriptMetadataMigrationsStepsModel              `tfsdk:"steps" json:"steps,optional"`
}

type WorkersScriptMetadataMigrationsRenamedClassesModel struct {
	From types.String `tfsdk:"from" json:"from,optional"`
	To   types.String `tfsdk:"to" json:"to,optional"`
}

type WorkersScriptMetadataMigrationsTransferredClassesModel struct {
	From       types.String `tfsdk:"from" json:"from,optional"`
	FromScript types.String `tfsdk:"from_script" json:"from_script,optional"`
	To         types.String `tfsdk:"to" json:"to,optional"`
}

type WorkersScriptMetadataMigrationsStepsModel struct {
	DeletedClasses     *[]types.String                                                 `tfsdk:"deleted_classes" json:"deleted_classes,optional"`
	NewClasses         *[]types.String                                                 `tfsdk:"new_classes" json:"new_classes,optional"`
	NewSqliteClasses   *[]types.String                                                 `tfsdk:"new_sqlite_classes" json:"new_sqlite_classes,optional"`
	RenamedClasses     *[]*WorkersScriptMetadataMigrationsStepsRenamedClassesModel     `tfsdk:"renamed_classes" json:"renamed_classes,optional"`
	TransferredClasses *[]*WorkersScriptMetadataMigrationsStepsTransferredClassesModel `tfsdk:"transferred_classes" json:"transferred_classes,optional"`
}

type WorkersScriptMetadataMigrationsStepsRenamedClassesModel struct {
	From types.String `tfsdk:"from" json:"from,optional"`
	To   types.String `tfsdk:"to" json:"to,optional"`
}

type WorkersScriptMetadataMigrationsStepsTransferredClassesModel struct {
	From       types.String `tfsdk:"from" json:"from,optional"`
	FromScript types.String `tfsdk:"from_script" json:"from_script,optional"`
	To         types.String `tfsdk:"to" json:"to,optional"`
}

type WorkersScriptMetadataObservabilityModel struct {
	Enabled          types.Bool                                   `tfsdk:"enabled" json:"enabled,required"`
	HeadSamplingRate types.Float64                                `tfsdk:"head_sampling_rate" json:"head_sampling_rate,optional"`
	Logs             *WorkersScriptMetadataObservabilityLogsModel `tfsdk:"logs" json:"logs,optional"`
}

type WorkersScriptMetadataObservabilityLogsModel struct {
	Enabled          types.Bool      `tfsdk:"enabled" json:"enabled,required"`
	InvocationLogs   types.Bool      `tfsdk:"invocation_logs" json:"invocation_logs,required"`
	Destinations     *[]types.String `tfsdk:"destinations" json:"destinations,optional"`
	HeadSamplingRate types.Float64   `tfsdk:"head_sampling_rate" json:"head_sampling_rate,optional"`
	Persist          types.Bool      `tfsdk:"persist" json:"persist,computed_optional"`
}

type WorkersScriptMetadataPlacementModel struct {
	Mode           types.String                                  `tfsdk:"mode" json:"mode,optional"`
	LastAnalyzedAt timetypes.RFC3339                             `tfsdk:"last_analyzed_at" json:"last_analyzed_at,computed" format:"date-time"`
	Status         types.String                                  `tfsdk:"status" json:"status,computed"`
	Region         types.String                                  `tfsdk:"region" json:"region,optional"`
	Hostname       types.String                                  `tfsdk:"hostname" json:"hostname,optional"`
	Host           types.String                                  `tfsdk:"host" json:"host,optional"`
	Target         *[]*WorkersScriptMetadataPlacementTargetModel `tfsdk:"target" json:"target,optional"`
}

type WorkersScriptMetadataPlacementTargetModel struct {
	Region   types.String `tfsdk:"region" json:"region,optional"`
	Hostname types.String `tfsdk:"hostname" json:"hostname,optional"`
	Host     types.String `tfsdk:"host" json:"host,optional"`
}

type WorkersScriptMetadataTailConsumersModel struct {
	Service     types.String `tfsdk:"service" json:"service,required"`
	Environment types.String `tfsdk:"environment" json:"environment,optional"`
	Namespace   types.String `tfsdk:"namespace" json:"namespace,optional"`
}

type WorkersScriptNamedHandlersModel struct {
	Handlers customfield.List[types.String] `tfsdk:"handlers" json:"handlers,computed"`
	Name     types.String                   `tfsdk:"name" json:"name,computed"`
}

type WorkersScriptObservabilityModel struct {
	Enabled          types.Bool                                                    `tfsdk:"enabled" json:"enabled,computed"`
	HeadSamplingRate types.Float64                                                 `tfsdk:"head_sampling_rate" json:"head_sampling_rate,computed"`
	Logs             customfield.NestedObject[WorkersScriptObservabilityLogsModel] `tfsdk:"logs" json:"logs,computed"`
}

type WorkersScriptObservabilityLogsModel struct {
	Enabled          types.Bool                     `tfsdk:"enabled" json:"enabled,computed"`
	InvocationLogs   types.Bool                     `tfsdk:"invocation_logs" json:"invocation_logs,computed"`
	Destinations     customfield.List[types.String] `tfsdk:"destinations" json:"destinations,computed"`
	HeadSamplingRate types.Float64                  `tfsdk:"head_sampling_rate" json:"head_sampling_rate,computed"`
	Persist          types.Bool                     `tfsdk:"persist" json:"persist,computed"`
}

type WorkersScriptPlacementModel struct {
	Mode           types.String                                                    `tfsdk:"mode" json:"mode,computed"`
	LastAnalyzedAt timetypes.RFC3339                                               `tfsdk:"last_analyzed_at" json:"last_analyzed_at,computed" format:"date-time"`
	Status         types.String                                                    `tfsdk:"status" json:"status,computed"`
	Region         types.String                                                    `tfsdk:"region" json:"region,computed"`
	Hostname       types.String                                                    `tfsdk:"hostname" json:"hostname,computed"`
	Host           types.String                                                    `tfsdk:"host" json:"host,computed"`
	Target         customfield.NestedObjectList[WorkersScriptPlacementTargetModel] `tfsdk:"target" json:"target,computed"`
}

type WorkersScriptPlacementTargetModel struct {
	Region   types.String `tfsdk:"region" json:"region,computed"`
	Hostname types.String `tfsdk:"hostname" json:"hostname,computed"`
	Host     types.String `tfsdk:"host" json:"host,computed"`
}

type WorkersScriptTailConsumersModel struct {
	Service     types.String `tfsdk:"service" json:"service,computed"`
	Environment types.String `tfsdk:"environment" json:"environment,computed"`
	Namespace   types.String `tfsdk:"namespace" json:"namespace,computed"`
}
