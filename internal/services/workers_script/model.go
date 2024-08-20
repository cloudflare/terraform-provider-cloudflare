// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_script

import (
	"bytes"
	"mime/multipart"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apiform"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersScriptResultEnvelope struct {
	Result WorkersScriptModel `json:"result,computed"`
}

type WorkersScriptModel struct {
	ID            types.String                        `tfsdk:"id" json:"-,computed"`
	ScriptName    types.String                        `tfsdk:"script_name" path:"script_name"`
	AccountID     types.String                        `tfsdk:"account_id" path:"account_id"`
	Message       types.String                        `tfsdk:"message" json:"message"`
	AnyPartName   *[]types.String                     `tfsdk:"any_part_name" json:"<any part name>"`
	Metadata      *WorkersScriptMetadataModel         `tfsdk:"metadata" json:"metadata"`
	CreatedOn     timetypes.RFC3339                   `tfsdk:"created_on" json:"created_on,computed"`
	Etag          types.String                        `tfsdk:"etag" json:"etag,computed"`
	Logpush       types.Bool                          `tfsdk:"logpush" json:"logpush,computed"`
	ModifiedOn    timetypes.RFC3339                   `tfsdk:"modified_on" json:"modified_on,computed"`
	PlacementMode types.String                        `tfsdk:"placement_mode" json:"placement_mode,computed"`
	StartupTimeMs types.Int64                         `tfsdk:"startup_time_ms" json:"startup_time_ms,computed"`
	UsageModel    types.String                        `tfsdk:"usage_model" json:"usage_model,computed"`
	TailConsumers *[]*WorkersScriptTailConsumersModel `tfsdk:"tail_consumers" json:"tail_consumers,computed"`
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
	Bindings           *[]*WorkersScriptMetadataBindingsModel      `tfsdk:"bindings" json:"bindings"`
	BodyPart           types.String                                `tfsdk:"body_part" json:"body_part"`
	CompatibilityDate  types.String                                `tfsdk:"compatibility_date" json:"compatibility_date"`
	CompatibilityFlags *[]types.String                             `tfsdk:"compatibility_flags" json:"compatibility_flags"`
	KeepBindings       *[]types.String                             `tfsdk:"keep_bindings" json:"keep_bindings"`
	Logpush            types.Bool                                  `tfsdk:"logpush" json:"logpush"`
	MainModule         types.String                                `tfsdk:"main_module" json:"main_module"`
	Migrations         *WorkersScriptMetadataMigrationsModel       `tfsdk:"migrations" json:"migrations"`
	Placement          *WorkersScriptMetadataPlacementModel        `tfsdk:"placement" json:"placement"`
	Tags               *[]types.String                             `tfsdk:"tags" json:"tags"`
	TailConsumers      *[]*WorkersScriptMetadataTailConsumersModel `tfsdk:"tail_consumers" json:"tail_consumers"`
	UsageModel         types.String                                `tfsdk:"usage_model" json:"usage_model"`
	VersionTags        map[string]types.String                     `tfsdk:"version_tags" json:"version_tags"`
}

type WorkersScriptMetadataBindingsModel struct {
	Name types.String `tfsdk:"name" json:"name"`
	Type types.String `tfsdk:"type" json:"type"`
}

type WorkersScriptMetadataMigrationsModel struct {
	DeletedClasses     *[]types.String                                            `tfsdk:"deleted_classes" json:"deleted_classes"`
	NewClasses         *[]types.String                                            `tfsdk:"new_classes" json:"new_classes"`
	NewTag             types.String                                               `tfsdk:"new_tag" json:"new_tag"`
	OldTag             types.String                                               `tfsdk:"old_tag" json:"old_tag"`
	RenamedClasses     *[]*WorkersScriptMetadataMigrationsRenamedClassesModel     `tfsdk:"renamed_classes" json:"renamed_classes"`
	TransferredClasses *[]*WorkersScriptMetadataMigrationsTransferredClassesModel `tfsdk:"transferred_classes" json:"transferred_classes"`
	Steps              *[]*WorkersScriptMetadataMigrationsStepsModel              `tfsdk:"steps" json:"steps"`
}

type WorkersScriptMetadataMigrationsRenamedClassesModel struct {
	From types.String `tfsdk:"from" json:"from"`
	To   types.String `tfsdk:"to" json:"to"`
}

type WorkersScriptMetadataMigrationsTransferredClassesModel struct {
	From       types.String `tfsdk:"from" json:"from"`
	FromScript types.String `tfsdk:"from_script" json:"from_script"`
	To         types.String `tfsdk:"to" json:"to"`
}

type WorkersScriptMetadataMigrationsStepsModel struct {
	DeletedClasses     *[]types.String                                                 `tfsdk:"deleted_classes" json:"deleted_classes"`
	NewClasses         *[]types.String                                                 `tfsdk:"new_classes" json:"new_classes"`
	RenamedClasses     *[]*WorkersScriptMetadataMigrationsStepsRenamedClassesModel     `tfsdk:"renamed_classes" json:"renamed_classes"`
	TransferredClasses *[]*WorkersScriptMetadataMigrationsStepsTransferredClassesModel `tfsdk:"transferred_classes" json:"transferred_classes"`
}

type WorkersScriptMetadataMigrationsStepsRenamedClassesModel struct {
	From types.String `tfsdk:"from" json:"from"`
	To   types.String `tfsdk:"to" json:"to"`
}

type WorkersScriptMetadataMigrationsStepsTransferredClassesModel struct {
	From       types.String `tfsdk:"from" json:"from"`
	FromScript types.String `tfsdk:"from_script" json:"from_script"`
	To         types.String `tfsdk:"to" json:"to"`
}

type WorkersScriptMetadataPlacementModel struct {
	Mode types.String `tfsdk:"mode" json:"mode"`
}

type WorkersScriptMetadataTailConsumersModel struct {
	Service     types.String `tfsdk:"service" json:"service"`
	Environment types.String `tfsdk:"environment" json:"environment"`
	Namespace   types.String `tfsdk:"namespace" json:"namespace"`
}

type WorkersScriptTailConsumersModel struct {
	Service     types.String `tfsdk:"service" json:"service"`
	Environment types.String `tfsdk:"environment" json:"environment"`
	Namespace   types.String `tfsdk:"namespace" json:"namespace"`
}
