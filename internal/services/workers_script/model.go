// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_script

import (
	"bytes"
	"mime/multipart"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apiform"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersScriptResultEnvelope struct {
	Result WorkersScriptModel `json:"result"`
}

type WorkersScriptModel struct {
	ID              types.String                                                  `tfsdk:"id" json:"-,computed"`
	ScriptName      types.String                                                  `tfsdk:"script_name" path:"script_name,required"`
	AccountID       types.String                                                  `tfsdk:"account_id" path:"account_id,required"`
	Metadata        *WorkersScriptMetadataModel                                   `tfsdk:"metadata" json:"metadata,required"`
	CreatedOn       timetypes.RFC3339                                             `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Etag            types.String                                                  `tfsdk:"etag" json:"etag,computed"`
	HasAssets       types.Bool                                                    `tfsdk:"has_assets" json:"has_assets,computed"`
	HasModules      types.Bool                                                    `tfsdk:"has_modules" json:"has_modules,computed"`
	Logpush         types.Bool                                                    `tfsdk:"logpush" json:"logpush,computed"`
	ModifiedOn      timetypes.RFC3339                                             `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	PlacementMode   types.String                                                  `tfsdk:"placement_mode" json:"placement_mode,computed"`
	PlacementStatus types.String                                                  `tfsdk:"placement_status" json:"placement_status,computed"`
	StartupTimeMs   types.Int64                                                   `tfsdk:"startup_time_ms" json:"startup_time_ms,computed"`
	UsageModel      types.String                                                  `tfsdk:"usage_model" json:"usage_model,computed"`
	Placement       customfield.NestedObject[WorkersScriptPlacementModel]         `tfsdk:"placement" json:"placement,computed"`
	TailConsumers   customfield.NestedObjectList[WorkersScriptTailConsumersModel] `tfsdk:"tail_consumers" json:"tail_consumers,computed"`
}

func (r WorkersScriptModel) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
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
	CompatibilityFlags *[]types.String                             `tfsdk:"compatibility_flags" json:"compatibility_flags,optional"`
	KeepAssets         types.Bool                                  `tfsdk:"keep_assets" json:"keep_assets,optional"`
	KeepBindings       *[]types.String                             `tfsdk:"keep_bindings" json:"keep_bindings,optional"`
	Logpush            types.Bool                                  `tfsdk:"logpush" json:"logpush,optional"`
	MainModule         types.String                                `tfsdk:"main_module" json:"main_module,optional"`
	Migrations         *WorkersScriptMetadataMigrationsModel       `tfsdk:"migrations" json:"migrations,optional"`
	Observability      *WorkersScriptMetadataObservabilityModel    `tfsdk:"observability" json:"observability,optional"`
	Placement          *WorkersScriptMetadataPlacementModel        `tfsdk:"placement" json:"placement,optional"`
	Tags               *[]types.String                             `tfsdk:"tags" json:"tags,optional"`
	TailConsumers      *[]*WorkersScriptMetadataTailConsumersModel `tfsdk:"tail_consumers" json:"tail_consumers,optional"`
	UsageModel         types.String                                `tfsdk:"usage_model" json:"usage_model,optional"`
}

type WorkersScriptMetadataAssetsModel struct {
	Config *WorkersScriptMetadataAssetsConfigModel `tfsdk:"config" json:"config,optional"`
	JWT    types.String                            `tfsdk:"jwt" json:"jwt,optional"`
}

type WorkersScriptMetadataAssetsConfigModel struct {
	HTMLHandling     types.String `tfsdk:"html_handling" json:"html_handling,optional"`
	NotFoundHandling types.String `tfsdk:"not_found_handling" json:"not_found_handling,optional"`
	ServeDirectly    types.Bool   `tfsdk:"serve_directly" json:"serve_directly,computed_optional"`
}

type WorkersScriptMetadataBindingsModel struct {
	Name          types.String                                `tfsdk:"name" json:"name,required"`
	Type          types.String                                `tfsdk:"type" json:"type,required"`
	Dataset       types.String                                `tfsdk:"dataset" json:"dataset,optional"`
	ID            types.String                                `tfsdk:"id" json:"id,optional"`
	Namespace     types.String                                `tfsdk:"namespace" json:"namespace,optional"`
	Outbound      *WorkersScriptMetadataBindingsOutboundModel `tfsdk:"outbound" json:"outbound,optional"`
	ClassName     types.String                                `tfsdk:"class_name" json:"class_name,optional"`
	Environment   types.String                                `tfsdk:"environment" json:"environment,optional"`
	NamespaceID   types.String                                `tfsdk:"namespace_id" json:"namespace_id,optional"`
	ScriptName    types.String                                `tfsdk:"script_name" json:"script_name,optional"`
	JSON          types.String                                `tfsdk:"json" json:"json,optional"`
	CertificateID types.String                                `tfsdk:"certificate_id" json:"certificate_id,optional"`
	Text          types.String                                `tfsdk:"text" json:"text,optional"`
	QueueName     types.String                                `tfsdk:"queue_name" json:"queue_name,optional"`
	BucketName    types.String                                `tfsdk:"bucket_name" json:"bucket_name,optional"`
	Service       types.String                                `tfsdk:"service" json:"service,optional"`
	IndexName     types.String                                `tfsdk:"index_name" json:"index_name,optional"`
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
	Enabled          types.Bool    `tfsdk:"enabled" json:"enabled,required"`
	HeadSamplingRate types.Float64 `tfsdk:"head_sampling_rate" json:"head_sampling_rate,optional"`
}

type WorkersScriptMetadataPlacementModel struct {
	Mode   types.String `tfsdk:"mode" json:"mode,optional"`
	Status types.String `tfsdk:"status" json:"status,computed"`
}

type WorkersScriptMetadataTailConsumersModel struct {
	Service     types.String `tfsdk:"service" json:"service,required"`
	Environment types.String `tfsdk:"environment" json:"environment,optional"`
	Namespace   types.String `tfsdk:"namespace" json:"namespace,optional"`
}

type WorkersScriptPlacementModel struct {
	Mode   types.String `tfsdk:"mode" json:"mode,computed"`
	Status types.String `tfsdk:"status" json:"status,computed"`
}

type WorkersScriptTailConsumersModel struct {
	Service     types.String `tfsdk:"service" json:"service,computed"`
	Environment types.String `tfsdk:"environment" json:"environment,computed"`
	Namespace   types.String `tfsdk:"namespace" json:"namespace,computed"`
}
