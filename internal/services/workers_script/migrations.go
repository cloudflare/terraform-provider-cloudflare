// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_script

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithUpgradeState = (*WorkersScriptResource)(nil)

func (r *WorkersScriptResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema:   resourceSchemaV0(ctx),
			StateUpgrader: upgradeStateFromV0,
		},
	}
}

// The schema is identical to the schema in schema.go, except the version is 0
// and the assets.config.run_worker_first field is a boolean instead of a
// dynamic value.
func resourceSchemaV0(ctx context.Context) *schema.Schema {
	resourceSchemaLatest := ResourceSchema(ctx)
	resourceSchemaLatest.Version = 0
	resourceSchemaLatest.
		Attributes["assets"].(schema.SingleNestedAttribute).
		Attributes["config"].(schema.SingleNestedAttribute).
		Attributes["run_worker_first"] = schema.BoolAttribute{Optional: true}
	return &resourceSchemaLatest
}

// State upgrade function from version 0 to version 1. This converts
// assets.config.run_worker_first from a boolean value to a dynamic value
// (either boolean or list of strings).
func upgradeStateFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	var priorStateData resourceModelV0
	resp.Diagnostics.Append(req.State.Get(ctx, &priorStateData)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Copy over all state data except for assets unchanged.
	newStateData := WorkersScriptModel{
		ID:               priorStateData.ID,
		ScriptName:       priorStateData.ScriptName,
		AccountID:        priorStateData.AccountID,
		Content:          priorStateData.Content,
		ContentFile:      priorStateData.ContentFile,
		ContentSHA256:    priorStateData.ContentSHA256,
		ContentType:      priorStateData.ContentType,
		CreatedOn:        priorStateData.CreatedOn,
		Etag:             priorStateData.Etag,
		HasAssets:        priorStateData.HasAssets,
		HasModules:       priorStateData.HasModules,
		LastDeployedFrom: priorStateData.LastDeployedFrom,
		MigrationTag:     priorStateData.MigrationTag,
		ModifiedOn:       priorStateData.ModifiedOn,
		StartupTimeMs:    priorStateData.StartupTimeMs,
		Handlers:         priorStateData.Handlers,
		NamedHandlers:    priorStateData.NamedHandlers,
		WorkersScriptMetadataModel: WorkersScriptMetadataModel{
			Bindings:           priorStateData.Bindings,
			BodyPart:           priorStateData.BodyPart,
			CompatibilityDate:  priorStateData.CompatibilityDate,
			CompatibilityFlags: priorStateData.CompatibilityFlags,
			KeepAssets:         priorStateData.KeepAssets,
			KeepBindings:       priorStateData.KeepBindings,
			Limits:             priorStateData.Limits,
			Logpush:            priorStateData.Logpush,
			MainModule:         priorStateData.MainModule,
			Migrations:         priorStateData.Migrations,
			Observability:      priorStateData.Observability,
			Placement:          priorStateData.Placement,
			TailConsumers:      priorStateData.TailConsumers,
			UsageModel:         priorStateData.UsageModel,
		},
	}

	if priorStateData.Assets != nil {
		newStateData.Assets = &WorkersScriptMetadataAssetsModel{
			JWT:                 priorStateData.Assets.JWT,
			Directory:           priorStateData.Assets.Directory,
			AssetManifestSHA256: priorStateData.Assets.AssetManifestSHA256,
		}
		if priorStateData.Assets.Config != nil {
			newStateData.Assets.Config = &WorkersScriptMetadataAssetsConfigModel{
				Headers:          priorStateData.Assets.Config.Headers,
				Redirects:        priorStateData.Assets.Config.Redirects,
				HTMLHandling:     priorStateData.Assets.Config.HTMLHandling,
				NotFoundHandling: priorStateData.Assets.Config.NotFoundHandling,
				ServeDirectly:    priorStateData.Assets.Config.ServeDirectly,
				// Convert run_worker_first from boolean to dynamic value.
				RunWorkerFirst: customfield.RawNormalizedDynamicValueFrom(priorStateData.Assets.Config.RunWorkerFirst),
			}

		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newStateData)...)
}

type resourceModelV0 struct {
	ID               types.String                                                  `tfsdk:"id" json:"-,computed"`
	ScriptName       types.String                                                  `tfsdk:"script_name" path:"script_name,required"`
	AccountID        types.String                                                  `tfsdk:"account_id" path:"account_id,required"`
	Content          types.String                                                  `tfsdk:"content" json:"-"`
	ContentFile      types.String                                                  `tfsdk:"content_file" json:"-"`
	ContentSHA256    types.String                                                  `tfsdk:"content_sha256" json:"-"`
	ContentType      types.String                                                  `tfsdk:"content_type" json:"-"`
	CreatedOn        timetypes.RFC3339                                             `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Etag             types.String                                                  `tfsdk:"etag" json:"etag,computed"`
	HasAssets        types.Bool                                                    `tfsdk:"has_assets" json:"has_assets,computed"`
	HasModules       types.Bool                                                    `tfsdk:"has_modules" json:"has_modules,computed"`
	LastDeployedFrom types.String                                                  `tfsdk:"last_deployed_from" json:"last_deployed_from,computed"`
	MigrationTag     types.String                                                  `tfsdk:"migration_tag" json:"migration_tag,computed"`
	ModifiedOn       timetypes.RFC3339                                             `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	StartupTimeMs    types.Int64                                                   `tfsdk:"startup_time_ms" json:"startup_time_ms,computed"`
	Handlers         customfield.List[types.String]                                `tfsdk:"handlers" json:"handlers,computed"`
	NamedHandlers    customfield.NestedObjectList[WorkersScriptNamedHandlersModel] `tfsdk:"named_handlers" json:"named_handlers,computed"`

	// Embedded metadata properties
	Bindings           customfield.NestedObjectList[WorkersScriptMetadataBindingsModel]     `tfsdk:"bindings" json:"bindings,computed_optional"`
	BodyPart           types.String                                                         `tfsdk:"body_part" json:"body_part,optional"`
	CompatibilityDate  types.String                                                         `tfsdk:"compatibility_date" json:"compatibility_date,computed_optional"`
	CompatibilityFlags customfield.Set[types.String]                                        `tfsdk:"compatibility_flags" json:"compatibility_flags,computed_optional"`
	KeepAssets         types.Bool                                                           `tfsdk:"keep_assets" json:"keep_assets,optional"`
	KeepBindings       *[]types.String                                                      `tfsdk:"keep_bindings" json:"keep_bindings,optional"`
	Limits             *WorkersScriptMetadataLimitsModel                                    `tfsdk:"limits" json:"limits,optional"`
	Logpush            types.Bool                                                           `tfsdk:"logpush" json:"logpush,computed_optional"`
	MainModule         types.String                                                         `tfsdk:"main_module" json:"main_module,optional"`
	Migrations         customfield.NestedObject[WorkersScriptMetadataMigrationsModel]       `tfsdk:"migrations" json:"migrations,optional"`
	Observability      *WorkersScriptMetadataObservabilityModel                             `tfsdk:"observability" json:"observability,optional"`
	Placement          customfield.NestedObject[WorkersScriptMetadataPlacementModel]        `tfsdk:"placement" json:"placement,computed_optional"`
	TailConsumers      customfield.NestedObjectSet[WorkersScriptMetadataTailConsumersModel] `tfsdk:"tail_consumers" json:"tail_consumers,computed_optional"`
	UsageModel         types.String                                                         `tfsdk:"usage_model" json:"usage_model,computed_optional"`

	// Old assets type definition
	Assets *struct {
		Config *struct {
			Headers          types.String `tfsdk:"headers" json:"_headers,optional"`
			Redirects        types.String `tfsdk:"redirects" json:"_redirects,optional"`
			HTMLHandling     types.String `tfsdk:"html_handling" json:"html_handling,optional"`
			NotFoundHandling types.String `tfsdk:"not_found_handling" json:"not_found_handling,optional"`
			RunWorkerFirst   types.Bool   `tfsdk:"run_worker_first" json:"run_worker_first,optional"`
			ServeDirectly    types.Bool   `tfsdk:"serve_directly" json:"serve_directly,optional"`
		} `tfsdk:"config" json:"config,optional"`
		JWT                 types.String `tfsdk:"jwt" json:"jwt,optional"`
		Directory           types.String `tfsdk:"directory" json:"-,optional"`
		AssetManifestSHA256 types.String `tfsdk:"asset_manifest_sha256" json:"-,computed"`
	} `tfsdk:"assets" json:"assets,optional"`
}
