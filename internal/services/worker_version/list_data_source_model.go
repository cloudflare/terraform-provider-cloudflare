// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package worker_version

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/workers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkerVersionsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[WorkerVersionsResultDataSourceModel] `json:"result,computed"`
}

type WorkerVersionsDataSourceModel struct {
	AccountID types.String                                                      `tfsdk:"account_id" path:"account_id,required"`
	WorkerID  types.String                                                      `tfsdk:"worker_id" path:"worker_id,required"`
	MaxItems  types.Int64                                                       `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[WorkerVersionsResultDataSourceModel] `tfsdk:"result"`
}

func (m *WorkerVersionsDataSourceModel) toListParams(_ context.Context) (params workers.BetaWorkerVersionListParams, diags diag.Diagnostics) {
	params = workers.BetaWorkerVersionListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type WorkerVersionsResultDataSourceModel struct {
	ID                 types.String                                                        `tfsdk:"id" json:"id,computed"`
	CreatedOn          timetypes.RFC3339                                                   `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Number             types.Int64                                                         `tfsdk:"number" json:"number,computed"`
	Annotations        customfield.NestedObject[WorkerVersionsAnnotationsDataSourceModel]  `tfsdk:"annotations" json:"annotations,computed"`
	Assets             customfield.NestedObject[WorkerVersionsAssetsDataSourceModel]       `tfsdk:"assets" json:"assets,computed"`
	Bindings           customfield.NestedObjectList[WorkerVersionsBindingsDataSourceModel] `tfsdk:"bindings" json:"bindings,computed"`
	CompatibilityDate  types.String                                                        `tfsdk:"compatibility_date" json:"compatibility_date,computed"`
	CompatibilityFlags customfield.Set[types.String]                                       `tfsdk:"compatibility_flags" json:"compatibility_flags,computed"`
	Limits             customfield.NestedObject[WorkerVersionsLimitsDataSourceModel]       `tfsdk:"limits" json:"limits,computed"`
	MainModule         types.String                                                        `tfsdk:"main_module" json:"main_module,computed"`
	Migrations         customfield.NestedObject[WorkerVersionsMigrationsDataSourceModel]   `tfsdk:"migrations" json:"migrations,computed"`
	Modules            customfield.NestedObjectSet[WorkerVersionsModulesDataSourceModel]   `tfsdk:"modules" json:"modules,computed"`
	Placement          customfield.NestedObject[WorkerVersionsPlacementDataSourceModel]    `tfsdk:"placement" json:"placement,computed"`
	Source             types.String                                                        `tfsdk:"source" json:"source,computed"`
	StartupTimeMs      types.Int64                                                         `tfsdk:"startup_time_ms" json:"startup_time_ms,computed"`
	UsageModel         types.String                                                        `tfsdk:"usage_model" json:"usage_model,computed"`
}

type WorkerVersionsAnnotationsDataSourceModel struct {
	WorkersMessage     types.String `tfsdk:"workers_message" json:"workers/message,computed"`
	WorkersTag         types.String `tfsdk:"workers_tag" json:"workers/tag,computed"`
	WorkersTriggeredBy types.String `tfsdk:"workers_triggered_by" json:"workers/triggered_by,computed"`
}

type WorkerVersionsAssetsDataSourceModel struct {
	Config customfield.NestedObject[WorkerVersionsAssetsConfigDataSourceModel] `tfsdk:"config" json:"config,computed"`
	JWT    types.String                                                        `tfsdk:"jwt" json:"jwt,computed"`
}

type WorkerVersionsAssetsConfigDataSourceModel struct {
	HTMLHandling     types.String                   `tfsdk:"html_handling" json:"html_handling,computed"`
	NotFoundHandling types.String                   `tfsdk:"not_found_handling" json:"not_found_handling,computed"`
	RunWorkerFirst   customfield.List[types.String] `tfsdk:"run_worker_first" json:"run_worker_first,computed"`
}

type WorkerVersionsBindingsDataSourceModel struct {
	Name                        types.String                                                            `tfsdk:"name" json:"name,computed"`
	Type                        types.String                                                            `tfsdk:"type" json:"type,computed"`
	Dataset                     types.String                                                            `tfsdk:"dataset" json:"dataset,computed"`
	ID                          types.String                                                            `tfsdk:"id" json:"id,computed"`
	Part                        types.String                                                            `tfsdk:"part" json:"part,computed"`
	Namespace                   types.String                                                            `tfsdk:"namespace" json:"namespace,computed"`
	Outbound                    customfield.NestedObject[WorkerVersionsBindingsOutboundDataSourceModel] `tfsdk:"outbound" json:"outbound,computed"`
	ClassName                   types.String                                                            `tfsdk:"class_name" json:"class_name,computed"`
	Environment                 types.String                                                            `tfsdk:"environment" json:"environment,computed"`
	NamespaceID                 types.String                                                            `tfsdk:"namespace_id" json:"namespace_id,computed"`
	ScriptName                  types.String                                                            `tfsdk:"script_name" json:"script_name,computed"`
	OldName                     types.String                                                            `tfsdk:"old_name" json:"old_name,computed"`
	VersionID                   types.String                                                            `tfsdk:"version_id" json:"version_id,computed"`
	Json                        types.String                                                            `tfsdk:"json" json:"json,computed"`
	CertificateID               types.String                                                            `tfsdk:"certificate_id" json:"certificate_id,computed"`
	Text                        types.String                                                            `tfsdk:"text" json:"text,computed"`
	Pipeline                    types.String                                                            `tfsdk:"pipeline" json:"pipeline,computed"`
	QueueName                   types.String                                                            `tfsdk:"queue_name" json:"queue_name,computed"`
	BucketName                  types.String                                                            `tfsdk:"bucket_name" json:"bucket_name,computed"`
	Jurisdiction                types.String                                                            `tfsdk:"jurisdiction" json:"jurisdiction,computed"`
	AllowedDestinationAddresses customfield.List[types.String]                                          `tfsdk:"allowed_destination_addresses" json:"allowed_destination_addresses,computed"`
	AllowedSenderAddresses      customfield.List[types.String]                                          `tfsdk:"allowed_sender_addresses" json:"allowed_sender_addresses,computed"`
	DestinationAddress          types.String                                                            `tfsdk:"destination_address" json:"destination_address,computed"`
	Service                     types.String                                                            `tfsdk:"service" json:"service,computed"`
	IndexName                   types.String                                                            `tfsdk:"index_name" json:"index_name,computed"`
	SecretName                  types.String                                                            `tfsdk:"secret_name" json:"secret_name,computed"`
	StoreID                     types.String                                                            `tfsdk:"store_id" json:"store_id,computed"`
	Algorithm                   jsontypes.Normalized                                                    `tfsdk:"algorithm" json:"algorithm,computed"`
	Format                      types.String                                                            `tfsdk:"format" json:"format,computed"`
	Usages                      customfield.Set[types.String]                                           `tfsdk:"usages" json:"usages,computed"`
	KeyBase64                   types.String                                                            `tfsdk:"key_base64" json:"key_base64,computed"`
	KeyJwk                      jsontypes.Normalized                                                    `tfsdk:"key_jwk" json:"key_jwk,computed"`
	WorkflowName                types.String                                                            `tfsdk:"workflow_name" json:"workflow_name,computed"`
}

type WorkerVersionsBindingsOutboundDataSourceModel struct {
	Params customfield.List[types.String]                                                `tfsdk:"params" json:"params,computed"`
	Worker customfield.NestedObject[WorkerVersionsBindingsOutboundWorkerDataSourceModel] `tfsdk:"worker" json:"worker,computed"`
}

type WorkerVersionsBindingsOutboundWorkerDataSourceModel struct {
	Environment types.String `tfsdk:"environment" json:"environment,computed"`
	Service     types.String `tfsdk:"service" json:"service,computed"`
}

type WorkerVersionsLimitsDataSourceModel struct {
	CPUMs types.Int64 `tfsdk:"cpu_ms" json:"cpu_ms,computed"`
}

type WorkerVersionsMigrationsDataSourceModel struct {
	DeletedClasses     customfield.List[types.String]                                                          `tfsdk:"deleted_classes" json:"deleted_classes,computed"`
	NewClasses         customfield.List[types.String]                                                          `tfsdk:"new_classes" json:"new_classes,computed"`
	NewSqliteClasses   customfield.List[types.String]                                                          `tfsdk:"new_sqlite_classes" json:"new_sqlite_classes,computed"`
	NewTag             types.String                                                                            `tfsdk:"new_tag" json:"new_tag,computed"`
	OldTag             types.String                                                                            `tfsdk:"old_tag" json:"old_tag,computed"`
	RenamedClasses     customfield.NestedObjectList[WorkerVersionsMigrationsRenamedClassesDataSourceModel]     `tfsdk:"renamed_classes" json:"renamed_classes,computed"`
	TransferredClasses customfield.NestedObjectList[WorkerVersionsMigrationsTransferredClassesDataSourceModel] `tfsdk:"transferred_classes" json:"transferred_classes,computed"`
	Steps              customfield.NestedObjectList[WorkerVersionsMigrationsStepsDataSourceModel]              `tfsdk:"steps" json:"steps,computed"`
}

type WorkerVersionsMigrationsRenamedClassesDataSourceModel struct {
	From types.String `tfsdk:"from" json:"from,computed"`
	To   types.String `tfsdk:"to" json:"to,computed"`
}

type WorkerVersionsMigrationsTransferredClassesDataSourceModel struct {
	From       types.String `tfsdk:"from" json:"from,computed"`
	FromScript types.String `tfsdk:"from_script" json:"from_script,computed"`
	To         types.String `tfsdk:"to" json:"to,computed"`
}

type WorkerVersionsMigrationsStepsDataSourceModel struct {
	DeletedClasses     customfield.List[types.String]                                                               `tfsdk:"deleted_classes" json:"deleted_classes,computed"`
	NewClasses         customfield.List[types.String]                                                               `tfsdk:"new_classes" json:"new_classes,computed"`
	NewSqliteClasses   customfield.List[types.String]                                                               `tfsdk:"new_sqlite_classes" json:"new_sqlite_classes,computed"`
	RenamedClasses     customfield.NestedObjectList[WorkerVersionsMigrationsStepsRenamedClassesDataSourceModel]     `tfsdk:"renamed_classes" json:"renamed_classes,computed"`
	TransferredClasses customfield.NestedObjectList[WorkerVersionsMigrationsStepsTransferredClassesDataSourceModel] `tfsdk:"transferred_classes" json:"transferred_classes,computed"`
}

type WorkerVersionsMigrationsStepsRenamedClassesDataSourceModel struct {
	From types.String `tfsdk:"from" json:"from,computed"`
	To   types.String `tfsdk:"to" json:"to,computed"`
}

type WorkerVersionsMigrationsStepsTransferredClassesDataSourceModel struct {
	From       types.String `tfsdk:"from" json:"from,computed"`
	FromScript types.String `tfsdk:"from_script" json:"from_script,computed"`
	To         types.String `tfsdk:"to" json:"to,computed"`
}

type WorkerVersionsModulesDataSourceModel struct {
	ContentBase64 types.String `tfsdk:"content_base64" json:"content_base64,computed"`
	ContentType   types.String `tfsdk:"content_type" json:"content_type,computed"`
	Name          types.String `tfsdk:"name" json:"name,computed"`
}

type WorkerVersionsPlacementDataSourceModel struct {
	Mode     types.String                                                               `tfsdk:"mode" json:"mode,computed"`
	Region   types.String                                                               `tfsdk:"region" json:"region,computed"`
	Hostname types.String                                                               `tfsdk:"hostname" json:"hostname,computed"`
	Host     types.String                                                               `tfsdk:"host" json:"host,computed"`
	Target   customfield.NestedObjectList[WorkerVersionsPlacementTargetDataSourceModel] `tfsdk:"target" json:"target,computed"`
}

type WorkerVersionsPlacementTargetDataSourceModel struct {
	Region   types.String `tfsdk:"region" json:"region,computed"`
	Hostname types.String `tfsdk:"hostname" json:"hostname,computed"`
	Host     types.String `tfsdk:"host" json:"host,computed"`
}
