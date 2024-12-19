// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_script

import (
	"bytes"
	"mime/multipart"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/jinzhu/copier"
)

type WorkersScriptResultEnvelope struct {
	Result WorkersScriptModel `json:"result"`
}

type WorkersScriptModel struct {
	ID            types.String      `tfsdk:"id" json:"-,computed"`
	ScriptName    types.String      `tfsdk:"script_name" path:"script_name,required"`
	AccountID     types.String      `tfsdk:"account_id" path:"account_id,required"`
	Message       types.String      `tfsdk:"message" json:"message,optional"`
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
	Assets             customfield.NestedObject[WorkersScriptMetadataAssetsModel]            `tfsdk:"assets" json:"assets,computed_optional"`
	Bindings           customfield.NestedObjectList[WorkersScriptMetadataBindingsModel]      `tfsdk:"bindings" json:"bindings,computed_optional"`
	BodyPart           types.String                                                          `tfsdk:"body_part" json:"body_part,optional"`
	CompatibilityDate  types.String                                                          `tfsdk:"compatibility_date" json:"compatibility_date,optional"`
	CompatibilityFlags *[]types.String                                                       `tfsdk:"compatibility_flags" json:"compatibility_flags,optional"`
	KeepAssets         types.Bool                                                            `tfsdk:"keep_assets" json:"keep_assets,optional"`
	KeepBindings       *[]types.String                                                       `tfsdk:"keep_bindings" json:"keep_bindings,optional"`
	Logpush            types.Bool                                                            `tfsdk:"logpush" json:"logpush,optional"`
	MainModule         types.String                                                          `tfsdk:"main_module" json:"main_module,optional"`
	Migrations         customfield.NestedObject[WorkersScriptMetadataMigrationsModel]        `tfsdk:"migrations" json:"migrations,computed_optional"`
	Placement          customfield.NestedObject[WorkersScriptMetadataPlacementModel]         `tfsdk:"placement" json:"placement,computed_optional"`
	Tags               *[]types.String                                                       `tfsdk:"tags" json:"tags,optional"`
	TailConsumers      customfield.NestedObjectList[WorkersScriptMetadataTailConsumersModel] `tfsdk:"tail_consumers" json:"tail_consumers,computed_optional"`
	UsageModel         types.String                                                          `tfsdk:"usage_model" json:"usage_model,optional"`
	VersionTags        *map[string]types.String                                              `tfsdk:"version_tags" json:"version_tags,optional"`
}

type WorkersScriptMetadataAssetsModel struct {
	Config customfield.NestedObject[WorkersScriptMetadataAssetsConfigModel] `tfsdk:"config" json:"config,computed_optional"`
	JWT    types.String                                                     `tfsdk:"jwt" json:"jwt,optional"`
}

type WorkersScriptMetadataAssetsConfigModel struct {
	HTMLHandling     types.String `tfsdk:"html_handling" json:"html_handling,optional"`
	NotFoundHandling types.String `tfsdk:"not_found_handling" json:"not_found_handling,optional"`
	ServeDirectly    types.Bool   `tfsdk:"serve_directly" json:"serve_directly,computed_optional"`
}

type WorkersScriptMetadataBindingsModel struct {
	Name types.String `tfsdk:"name" json:"name,optional"`
	Type types.String `tfsdk:"type" json:"type,optional"`
}

type WorkersScriptMetadataMigrationsModel struct {
	DeletedClasses     *[]types.String                                                                      `tfsdk:"deleted_classes" json:"deleted_classes,optional"`
	NewClasses         *[]types.String                                                                      `tfsdk:"new_classes" json:"new_classes,optional"`
	NewSqliteClasses   *[]types.String                                                                      `tfsdk:"new_sqlite_classes" json:"new_sqlite_classes,optional"`
	NewTag             types.String                                                                         `tfsdk:"new_tag" json:"new_tag,optional"`
	OldTag             types.String                                                                         `tfsdk:"old_tag" json:"old_tag,optional"`
	RenamedClasses     customfield.NestedObjectList[WorkersScriptMetadataMigrationsRenamedClassesModel]     `tfsdk:"renamed_classes" json:"renamed_classes,computed_optional"`
	TransferredClasses customfield.NestedObjectList[WorkersScriptMetadataMigrationsTransferredClassesModel] `tfsdk:"transferred_classes" json:"transferred_classes,computed_optional"`
	Steps              customfield.NestedObjectList[WorkersScriptMetadataMigrationsStepsModel]              `tfsdk:"steps" json:"steps,computed_optional"`
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
	DeletedClasses     *[]types.String                                                                           `tfsdk:"deleted_classes" json:"deleted_classes,optional"`
	NewClasses         *[]types.String                                                                           `tfsdk:"new_classes" json:"new_classes,optional"`
	NewSqliteClasses   *[]types.String                                                                           `tfsdk:"new_sqlite_classes" json:"new_sqlite_classes,optional"`
	RenamedClasses     customfield.NestedObjectList[WorkersScriptMetadataMigrationsStepsRenamedClassesModel]     `tfsdk:"renamed_classes" json:"renamed_classes,computed_optional"`
	TransferredClasses customfield.NestedObjectList[WorkersScriptMetadataMigrationsStepsTransferredClassesModel] `tfsdk:"transferred_classes" json:"transferred_classes,computed_optional"`
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

type WorkersScriptMetadataPlacementModel struct {
	Mode types.String `tfsdk:"mode" json:"mode,optional"`
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
