// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_script

import (
	"bytes"
	"mime/multipart"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/jinzhu/copier"
)

type WorkersServiceResultEnvelope struct {
	Result WorkersServiceModel `json:"result"`
}

type WorkersScriptResultEnvelope struct {
	Result WorkersScriptModel `json:"result"`
}

type WorkersScriptMetadataResultEnvelope struct {
	Result WorkersScriptMetadataModel `json:"result"`
}

type WorkersServiceModel struct {
	DefaultEnvironment WorkersEnvironmentModel `json:"default_environment"`
}

type WorkersEnvironmentModel struct {
	Script WorkersScriptModel `json:"script"`
}

type WorkersScriptModel struct {
	ID            types.String      `tfsdk:"id" json:"-,computed"`
	ScriptName    types.String      `tfsdk:"script_name" path:"script_name,required"`
	AccountID     types.String      `tfsdk:"account_id" path:"account_id,required"`
	Content       types.String      `tfsdk:"content" json:"content,required"`
	CreatedOn     timetypes.RFC3339 `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Etag          types.String      `tfsdk:"etag" json:"etag,computed"`
	HasAssets     types.Bool        `tfsdk:"has_assets" json:"has_assets,computed"`
	HasModules    types.Bool        `tfsdk:"has_modules" json:"has_modules,computed"`
	ModifiedOn    timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	StartupTimeMs types.Int64       `tfsdk:"startup_time_ms" json:"startup_time_ms,computed"`

	WorkersScriptMetadataModel
}

func (r WorkersScriptModel) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	var metadata WorkersScriptMetadataModel
	workerBody := bytes.NewReader([]byte(r.Content.ValueString()))

	if r.MainModule.ValueString() != "" {
		mainModuleName := r.MainModule.ValueString()
		writeFileBytes(mainModuleName, mainModuleName, "application/javascript+module", workerBody, writer)
	} else {
		writeFileBytes("script", "script", "application/javascript", workerBody, writer)
		r.BodyPart = types.StringValue("script")
	}

	topLevelMetadata := r.WorkersScriptMetadataModel
	copier.Copy(&metadata, &topLevelMetadata)

	payload, _ := apijson.Marshal(metadata)
	metadataContent := bytes.NewReader(payload)
	writeFileBytes("metadata", "", "application/json", metadataContent, writer)

	err = writer.Close()
	if err != nil {
		return nil, "", err
	}

	return buf.Bytes(), writer.FormDataContentType(), nil
}

type WorkersScriptMetadataModel struct {
	Assets             *WorkersScriptMetadataAssetsModel                                `tfsdk:"assets" json:"assets,optional"`
	Bindings           customfield.NestedObjectList[WorkersScriptMetadataBindingsModel] `tfsdk:"bindings" json:"bindings,computed_optional"`
	BodyPart           types.String                                                     `tfsdk:"body_part" json:"body_part,optional"`
	CompatibilityDate  types.String                                                     `tfsdk:"compatibility_date" json:"compatibility_date,computed_optional"`
	CompatibilityFlags customfield.List[types.String]                                   `tfsdk:"compatibility_flags" json:"compatibility_flags,computed_optional"`
	KeepAssets         types.Bool                                                       `tfsdk:"keep_assets" json:"keep_assets,optional"`
	KeepBindings       *[]types.String                                                  `tfsdk:"keep_bindings" json:"keep_bindings,optional"`
	Logpush            types.Bool                                                       `tfsdk:"logpush" json:"logpush,computed_optional"`
	MainModule         types.String                                                     `tfsdk:"main_module" json:"main_module,optional"`
	Migrations         customfield.NestedObject[WorkersScriptMetadataMigrationsModel]   `tfsdk:"migrations" json:"migrations,computed_optional"`
	Observability      *WorkersScriptMetadataObservabilityModel                         `tfsdk:"observability" json:"observability,optional"`
	Placement          customfield.NestedObject[WorkersScriptMetadataPlacementModel]    `tfsdk:"placement" json:"placement,computed_optional"`
	// Tags               *[]types.String                                                       `tfsdk:"tags" json:"tags,optional"`
	TailConsumers customfield.NestedObjectList[WorkersScriptMetadataTailConsumersModel] `tfsdk:"tail_consumers" json:"tail_consumers,computed_optional"`
	UsageModel    types.String                                                          `tfsdk:"usage_model" json:"usage_model,computed_optional"`
}

type WorkersScriptMetadataAssetsModel struct {
	Config *WorkersScriptMetadataAssetsConfigModel `tfsdk:"config" json:"config,optional"`
	JWT    types.String                            `tfsdk:"jwt" json:"jwt,optional"`
}

type WorkersScriptMetadataAssetsConfigModel struct {
	Headers          types.String `tfsdk:"_headers" json:"_headers,optional"`
	Redirects        types.String `tfsdk:"_redirects" json:"_redirects,optional"`
	HTMLHandling     types.String `tfsdk:"html_handling" json:"html_handling,optional"`
	NotFoundHandling types.String `tfsdk:"not_found_handling" json:"not_found_handling,optional"`
	RunWorkerFirst   types.Bool   `tfsdk:"run_worker_first" json:"run_worker_first,optional"`
	ServeDirectly    types.Bool   `tfsdk:"serve_directly" json:"serve_directly,optional"`
}

type WorkersScriptMetadataBindingsModel struct {
	Name          types.String                                `tfsdk:"name" json:"name,required"`
	Type          types.String                                `tfsdk:"type" json:"type,required"`
	Dataset       types.String                                `tfsdk:"dataset" json:"dataset,optional"`
	ID            types.String                                `tfsdk:"id" json:"id,optional"`
	Namespace     types.String                                `tfsdk:"namespace" json:"namespace,optional"`
	Outbound      *WorkersScriptMetadataBindingsOutboundModel `tfsdk:"outbound" json:"outbound,optional"`
	ClassName     types.String                                `tfsdk:"class_name" json:"class_name,computed_optional"`
	Environment   types.String                                `tfsdk:"environment" json:"environment,optional"`
	NamespaceID   types.String                                `tfsdk:"namespace_id" json:"namespace_id,computed_optional"`
	ScriptName    types.String                                `tfsdk:"script_name" json:"script_name,computed_optional"`
	Json          types.String                                `tfsdk:"json" json:"json,optional"`
	CertificateID types.String                                `tfsdk:"certificate_id" json:"certificate_id,optional"`
	Text          types.String                                `tfsdk:"text" json:"text,optional"`
	Pipeline      types.String                                `tfsdk:"pipeline" json:"pipeline,optional"`
	QueueName     types.String                                `tfsdk:"queue_name" json:"queue_name,optional"`
	BucketName    types.String                                `tfsdk:"bucket_name" json:"bucket_name,optional"`
	Service       types.String                                `tfsdk:"service" json:"service,optional"`
	IndexName     types.String                                `tfsdk:"index_name" json:"index_name,optional"`
	SecretName    types.String                                `tfsdk:"secret_name" json:"secret_name,optional"`
	StoreID       types.String                                `tfsdk:"store_id" json:"store_id,optional"`
	Algorithm     jsontypes.Normalized                        `tfsdk:"algorithm" json:"algorithm,optional"`
	Format        types.String                                `tfsdk:"format" json:"format,optional"`
	Usages        *[]types.String                             `tfsdk:"usages" json:"usages,optional"`
	KeyBase64     types.String                                `tfsdk:"key_base64" json:"key_base64,optional"`
	KeyJwk        jsontypes.Normalized                        `tfsdk:"key_jwk" json:"key_jwk,optional"`
	WorkflowName  types.String                                `tfsdk:"workflow_name" json:"workflow_name,optional"`
}

type WorkersScriptMetadataBindingsOutboundModel struct {
	Params *[]types.String                                   `tfsdk:"params" json:"params,optional"`
	Worker *WorkersScriptMetadataBindingsOutboundWorkerModel `tfsdk:"worker" json:"worker,optional"`
}

type WorkersScriptMetadataBindingsOutboundWorkerModel struct {
	Environment types.String `tfsdk:"environment" json:"environment,optional"`
	Service     types.String `tfsdk:"service" json:"service,optional"`
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
	Enabled          types.Bool    `tfsdk:"enabled" json:"enabled,required"`
	InvocationLogs   types.Bool    `tfsdk:"invocation_logs" json:"invocation_logs,required"`
	HeadSamplingRate types.Float64 `tfsdk:"head_sampling_rate" json:"head_sampling_rate,optional"`
}

type WorkersScriptMetadataPlacementModel struct {
	LastAnalyzedAt timetypes.RFC3339 `tfsdk:"last_analyzed_at" json:"last_analyzed_at,computed" format:"date-time"`
	Mode           types.String      `tfsdk:"mode" json:"mode,optional"`
	Status         types.String      `tfsdk:"status" json:"status,computed"`
}

type WorkersScriptMetadataTailConsumersModel struct {
	Service     types.String `tfsdk:"service" json:"service,required"`
	Environment types.String `tfsdk:"environment" json:"environment,optional"`
	Namespace   types.String `tfsdk:"namespace" json:"namespace,optional"`
}

type WorkersScriptTailConsumersModel struct {
	Service     types.String `tfsdk:"service" json:"service,computed"`
	Environment types.String `tfsdk:"environment" json:"environment,computed"`
	Namespace   types.String `tfsdk:"namespace" json:"namespace,computed"`
}
