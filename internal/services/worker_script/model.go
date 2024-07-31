// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package worker_script

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkerScriptResultEnvelope struct {
	Result WorkerScriptModel `json:"result,computed"`
}

type WorkerScriptModel struct {
	ID            types.String                       `tfsdk:"id" json:"-,computed"`
	AccountID     types.String                       `tfsdk:"account_id" path:"account_id"`
	ScriptName    types.String                       `tfsdk:"script_name" path:"script_name"`
	AnyPartName   *[]types.String                    `tfsdk:"any_part_name" json:"<any part name>"`
	Metadata      *WorkerScriptMetadataModel         `tfsdk:"metadata" json:"metadata"`
	Message       types.String                       `tfsdk:"message" json:"message"`
	CreatedOn     timetypes.RFC3339                  `tfsdk:"created_on" json:"created_on,computed"`
	Etag          types.String                       `tfsdk:"etag" json:"etag,computed"`
	Logpush       types.Bool                         `tfsdk:"logpush" json:"logpush,computed"`
	ModifiedOn    timetypes.RFC3339                  `tfsdk:"modified_on" json:"modified_on,computed"`
	PlacementMode types.String                       `tfsdk:"placement_mode" json:"placement_mode,computed"`
	StartupTimeMs types.Int64                        `tfsdk:"startup_time_ms" json:"startup_time_ms,computed"`
	TailConsumers *[]*WorkerScriptTailConsumersModel `tfsdk:"tail_consumers" json:"tail_consumers,computed"`
	UsageModel    types.String                       `tfsdk:"usage_model" json:"usage_model,computed"`
}

type WorkerScriptMetadataModel struct {
	Bindings           *[]jsontypes.Normalized                    `tfsdk:"bindings" json:"bindings"`
	BodyPart           types.String                               `tfsdk:"body_part" json:"body_part"`
	CompatibilityDate  types.String                               `tfsdk:"compatibility_date" json:"compatibility_date"`
	CompatibilityFlags *[]types.String                            `tfsdk:"compatibility_flags" json:"compatibility_flags"`
	KeepBindings       *[]types.String                            `tfsdk:"keep_bindings" json:"keep_bindings"`
	Logpush            types.Bool                                 `tfsdk:"logpush" json:"logpush"`
	MainModule         types.String                               `tfsdk:"main_module" json:"main_module"`
	Migrations         *WorkerScriptMetadataMigrationsModel       `tfsdk:"migrations" json:"migrations"`
	Placement          *WorkerScriptMetadataPlacementModel        `tfsdk:"placement" json:"placement"`
	Tags               *[]types.String                            `tfsdk:"tags" json:"tags"`
	TailConsumers      *[]*WorkerScriptMetadataTailConsumersModel `tfsdk:"tail_consumers" json:"tail_consumers"`
	UsageModel         types.String                               `tfsdk:"usage_model" json:"usage_model"`
	VersionTags        jsontypes.Normalized                       `tfsdk:"version_tags" json:"version_tags"`
}

type WorkerScriptMetadataMigrationsModel struct {
	DeletedClasses     *[]types.String                                           `tfsdk:"deleted_classes" json:"deleted_classes"`
	NewClasses         *[]types.String                                           `tfsdk:"new_classes" json:"new_classes"`
	NewTag             types.String                                              `tfsdk:"new_tag" json:"new_tag"`
	OldTag             types.String                                              `tfsdk:"old_tag" json:"old_tag"`
	RenamedClasses     *[]*WorkerScriptMetadataMigrationsRenamedClassesModel     `tfsdk:"renamed_classes" json:"renamed_classes"`
	TransferredClasses *[]*WorkerScriptMetadataMigrationsTransferredClassesModel `tfsdk:"transferred_classes" json:"transferred_classes"`
	Steps              *[]*WorkerScriptMetadataMigrationsStepsModel              `tfsdk:"steps" json:"steps"`
}

type WorkerScriptMetadataMigrationsRenamedClassesModel struct {
	From types.String `tfsdk:"from" json:"from"`
	To   types.String `tfsdk:"to" json:"to"`
}

type WorkerScriptMetadataMigrationsTransferredClassesModel struct {
	From       types.String `tfsdk:"from" json:"from"`
	FromScript types.String `tfsdk:"from_script" json:"from_script"`
	To         types.String `tfsdk:"to" json:"to"`
}

type WorkerScriptMetadataMigrationsStepsModel struct {
	DeletedClasses     *[]types.String                                                `tfsdk:"deleted_classes" json:"deleted_classes"`
	NewClasses         *[]types.String                                                `tfsdk:"new_classes" json:"new_classes"`
	RenamedClasses     *[]*WorkerScriptMetadataMigrationsStepsRenamedClassesModel     `tfsdk:"renamed_classes" json:"renamed_classes"`
	TransferredClasses *[]*WorkerScriptMetadataMigrationsStepsTransferredClassesModel `tfsdk:"transferred_classes" json:"transferred_classes"`
}

type WorkerScriptMetadataMigrationsStepsRenamedClassesModel struct {
	From types.String `tfsdk:"from" json:"from"`
	To   types.String `tfsdk:"to" json:"to"`
}

type WorkerScriptMetadataMigrationsStepsTransferredClassesModel struct {
	From       types.String `tfsdk:"from" json:"from"`
	FromScript types.String `tfsdk:"from_script" json:"from_script"`
	To         types.String `tfsdk:"to" json:"to"`
}

type WorkerScriptMetadataPlacementModel struct {
	Mode types.String `tfsdk:"mode" json:"mode"`
}

type WorkerScriptMetadataTailConsumersModel struct {
	Service     types.String `tfsdk:"service" json:"service"`
	Environment types.String `tfsdk:"environment" json:"environment"`
	Namespace   types.String `tfsdk:"namespace" json:"namespace"`
}

type WorkerScriptTailConsumersModel struct {
	Service     types.String `tfsdk:"service" json:"service"`
	Environment types.String `tfsdk:"environment" json:"environment"`
	Namespace   types.String `tfsdk:"namespace" json:"namespace"`
}
