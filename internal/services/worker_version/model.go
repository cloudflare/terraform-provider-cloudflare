// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package worker_version

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkerVersionResultEnvelope struct {
	Result WorkerVersionModel `json:"result"`
}

type WorkerVersionModel struct {
	ID                 types.String                                             `tfsdk:"id" json:"id,computed"`
	AccountID          types.String                                             `tfsdk:"account_id" path:"account_id,required"`
	WorkerID           types.String                                             `tfsdk:"worker_id" path:"worker_id,required"`
	CompatibilityDate  types.String                                             `tfsdk:"compatibility_date" json:"compatibility_date,optional"`
	MainModule         types.String                                             `tfsdk:"main_module" json:"main_module,optional"`
	Migrations         *WorkerVersionMigrationsModel                            `tfsdk:"migrations" json:"migrations,optional"`
	Modules            *[]*WorkerVersionModulesModel                            `tfsdk:"modules" json:"modules,optional"`
	Placement          *WorkerVersionPlacementModel                             `tfsdk:"placement" json:"placement,optional"`
	UsageModel         types.String                                             `tfsdk:"usage_model" json:"usage_model,computed_optional"`
	CompatibilityFlags customfield.Set[types.String]                            `tfsdk:"compatibility_flags" json:"compatibility_flags,computed_optional"`
	Annotations        customfield.NestedObject[WorkerVersionAnnotationsModel]  `tfsdk:"annotations" json:"annotations,computed_optional"`
	Assets             *WorkerVersionAssetsModel                                `tfsdk:"assets" json:"assets,optional"`
	Bindings           customfield.NestedObjectList[WorkerVersionBindingsModel] `tfsdk:"bindings" json:"bindings,optional"`
	Limits             customfield.NestedObject[WorkerVersionLimitsModel]       `tfsdk:"limits" json:"limits,computed_optional"`
	CreatedOn          timetypes.RFC3339                                        `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Number             types.Int64                                              `tfsdk:"number" json:"number,computed"`
	Source             types.String                                             `tfsdk:"source" json:"source,computed"`
	MainScriptBase64   types.String                                             `tfsdk:"main_script_base64" json:"main_script_base64,computed"`
	StartupTimeMs      types.Int64                                              `tfsdk:"startup_time_ms" json:"startup_time_ms,computed"`
}

func (m WorkerVersionModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m WorkerVersionModel) MarshalJSONForUpdate(state WorkerVersionModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type WorkerVersionMigrationsModel struct {
	DeletedClasses     *[]types.String                                    `tfsdk:"deleted_classes" json:"deleted_classes,optional"`
	NewClasses         *[]types.String                                    `tfsdk:"new_classes" json:"new_classes,optional"`
	NewSqliteClasses   *[]types.String                                    `tfsdk:"new_sqlite_classes" json:"new_sqlite_classes,optional"`
	NewTag             types.String                                       `tfsdk:"new_tag" json:"new_tag,optional"`
	OldTag             types.String                                       `tfsdk:"old_tag" json:"old_tag,optional"`
	RenamedClasses     *[]*WorkerVersionMigrationsRenamedClassesModel     `tfsdk:"renamed_classes" json:"renamed_classes,optional"`
	TransferredClasses *[]*WorkerVersionMigrationsTransferredClassesModel `tfsdk:"transferred_classes" json:"transferred_classes,optional"`
	Steps              *[]*WorkerVersionMigrationsStepsModel              `tfsdk:"steps" json:"steps,optional"`
}

type WorkerVersionMigrationsRenamedClassesModel struct {
	From types.String `tfsdk:"from" json:"from,optional"`
	To   types.String `tfsdk:"to" json:"to,optional"`
}

type WorkerVersionMigrationsTransferredClassesModel struct {
	From       types.String `tfsdk:"from" json:"from,optional"`
	FromScript types.String `tfsdk:"from_script" json:"from_script,optional"`
	To         types.String `tfsdk:"to" json:"to,optional"`
}

type WorkerVersionMigrationsStepsModel struct {
	DeletedClasses     *[]types.String                                         `tfsdk:"deleted_classes" json:"deleted_classes,optional"`
	NewClasses         *[]types.String                                         `tfsdk:"new_classes" json:"new_classes,optional"`
	NewSqliteClasses   *[]types.String                                         `tfsdk:"new_sqlite_classes" json:"new_sqlite_classes,optional"`
	RenamedClasses     *[]*WorkerVersionMigrationsStepsRenamedClassesModel     `tfsdk:"renamed_classes" json:"renamed_classes,optional"`
	TransferredClasses *[]*WorkerVersionMigrationsStepsTransferredClassesModel `tfsdk:"transferred_classes" json:"transferred_classes,optional"`
}

type WorkerVersionMigrationsStepsRenamedClassesModel struct {
	From types.String `tfsdk:"from" json:"from,optional"`
	To   types.String `tfsdk:"to" json:"to,optional"`
}

type WorkerVersionMigrationsStepsTransferredClassesModel struct {
	From       types.String `tfsdk:"from" json:"from,optional"`
	FromScript types.String `tfsdk:"from_script" json:"from_script,optional"`
	To         types.String `tfsdk:"to" json:"to,optional"`
}

type WorkerVersionModulesModel struct {
	ContentBase64 types.String `tfsdk:"content_base64" json:"content_base64,optional"`
	ContentType   types.String `tfsdk:"content_type" json:"content_type,required"`
	Name          types.String `tfsdk:"name" json:"name,required"`
	ContentFile   types.String `tfsdk:"content_file" json:"-,optional"`
	ContentSHA256 types.String `tfsdk:"content_sha256" json:"-,computed"`
}

type WorkerVersionPlacementModel struct {
	Mode     types.String                          `tfsdk:"mode" json:"mode,optional"`
	Region   types.String                          `tfsdk:"region" json:"region,optional"`
	Hostname types.String                          `tfsdk:"hostname" json:"hostname,optional"`
	Host     types.String                          `tfsdk:"host" json:"host,optional"`
	Target   *[]*WorkerVersionPlacementTargetModel `tfsdk:"target" json:"target,optional"`
}

type WorkerVersionPlacementTargetModel struct {
	Region   types.String `tfsdk:"region" json:"region,optional"`
	Hostname types.String `tfsdk:"hostname" json:"hostname,optional"`
	Host     types.String `tfsdk:"host" json:"host,optional"`
}

type WorkerVersionAnnotationsModel struct {
	WorkersMessage     types.String `tfsdk:"workers_message" json:"workers/message,optional"`
	WorkersTag         types.String `tfsdk:"workers_tag" json:"workers/tag,optional"`
	WorkersTriggeredBy types.String `tfsdk:"workers_triggered_by" json:"workers/triggered_by,computed"`
}

type WorkerVersionAssetsModel struct {
	Config              customfield.NestedObject[WorkerVersionAssetsConfigModel] `tfsdk:"config" json:"config,optional"`
	JWT                 types.String                                             `tfsdk:"jwt" json:"jwt,optional"`
	Directory           types.String                                             `tfsdk:"directory" json:"-,optional"`
	AssetManifestSHA256 types.String                                             `tfsdk:"asset_manifest_sha256" json:"-,computed"`
}

type WorkerVersionAssetsConfigModel struct {
	HTMLHandling     types.String                       `tfsdk:"html_handling" json:"html_handling,optional"`
	NotFoundHandling types.String                       `tfsdk:"not_found_handling" json:"not_found_handling,optional"`
	RunWorkerFirst   customfield.NormalizedDynamicValue `tfsdk:"run_worker_first" json:"run_worker_first,optional"`
}

type WorkerVersionBindingsModel struct {
	Name                        types.String                        `tfsdk:"name" json:"name,required"`
	Type                        types.String                        `tfsdk:"type" json:"type,required"`
	Dataset                     types.String                        `tfsdk:"dataset" json:"dataset,optional"`
	ID                          types.String                        `tfsdk:"id" json:"id,optional"`
	Part                        types.String                        `tfsdk:"part" json:"part,optional"`
	Namespace                   types.String                        `tfsdk:"namespace" json:"namespace,optional"`
	Outbound                    *WorkerVersionBindingsOutboundModel `tfsdk:"outbound" json:"outbound,optional"`
	ClassName                   types.String                        `tfsdk:"class_name" json:"class_name,computed_optional"`
	Environment                 types.String                        `tfsdk:"environment" json:"environment,optional"`
	NamespaceID                 types.String                        `tfsdk:"namespace_id" json:"namespace_id,computed_optional"`
	ScriptName                  types.String                        `tfsdk:"script_name" json:"script_name,computed_optional"`
	OldName                     types.String                        `tfsdk:"old_name" json:"old_name,optional"`
	VersionID                   types.String                        `tfsdk:"version_id" json:"version_id,optional"`
	Json                        types.String                        `tfsdk:"json" json:"json,optional"`
	CertificateID               types.String                        `tfsdk:"certificate_id" json:"certificate_id,optional"`
	Text                        types.String                        `tfsdk:"text" json:"text,optional"`
	Pipeline                    types.String                        `tfsdk:"pipeline" json:"pipeline,optional"`
	QueueName                   types.String                        `tfsdk:"queue_name" json:"queue_name,optional"`
	BucketName                  types.String                        `tfsdk:"bucket_name" json:"bucket_name,optional"`
	Jurisdiction                types.String                        `tfsdk:"jurisdiction" json:"jurisdiction,optional"`
	AllowedDestinationAddresses *[]types.String                     `tfsdk:"allowed_destination_addresses" json:"allowed_destination_addresses,optional"`
	AllowedSenderAddresses      *[]types.String                     `tfsdk:"allowed_sender_addresses" json:"allowed_sender_addresses,optional"`
	DestinationAddress          types.String                        `tfsdk:"destination_address" json:"destination_address,optional"`
	Service                     types.String                        `tfsdk:"service" json:"service,optional"`
	IndexName                   types.String                        `tfsdk:"index_name" json:"index_name,optional"`
	SecretName                  types.String                        `tfsdk:"secret_name" json:"secret_name,optional"`
	StoreID                     types.String                        `tfsdk:"store_id" json:"store_id,optional"`
	Algorithm                   jsontypes.Normalized                `tfsdk:"algorithm" json:"algorithm,optional"`
	Format                      types.String                        `tfsdk:"format" json:"format,optional"`
	Usages                      customfield.Set[types.String]       `tfsdk:"usages" json:"usages,optional"`
	KeyBase64                   types.String                        `tfsdk:"key_base64" json:"key_base64,optional"`
	KeyJwk                      jsontypes.Normalized                `tfsdk:"key_jwk" json:"key_jwk,optional"`
	WorkflowName                types.String                        `tfsdk:"workflow_name" json:"workflow_name,optional"`
}

type WorkerVersionBindingsOutboundModel struct {
	Params *[]types.String                           `tfsdk:"params" json:"params,optional"`
	Worker *WorkerVersionBindingsOutboundWorkerModel `tfsdk:"worker" json:"worker,optional"`
}

type WorkerVersionBindingsOutboundWorkerModel struct {
	Environment types.String `tfsdk:"environment" json:"environment,optional"`
	Service     types.String `tfsdk:"service" json:"service,optional"`
}

type WorkerVersionLimitsModel struct {
	CPUMs types.Int64 `tfsdk:"cpu_ms" json:"cpu_ms,required"`
}
