// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package worker_version

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithUpgradeState = (*WorkerVersionResource)(nil)

func (r *WorkerVersionResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema:   resourceSchemaV0(ctx),
			StateUpgrader: upgradeStateFromV0,
		},
	}
}

func resourceSchemaV0(ctx context.Context) *schema.Schema {
	resourceSchemaLatest := ResourceSchema(ctx)
	resourceSchemaLatest.Version = 0
	resourceSchemaLatest.
		Attributes["assets"].(schema.SingleNestedAttribute).
		Attributes["config"].(schema.SingleNestedAttribute).
		Attributes["run_worker_first"] = schema.ListAttribute{
		Optional:    true,
		CustomType:  customfield.NewListType[types.String](ctx),
		ElementType: types.StringType,
	}
	return &resourceSchemaLatest
}

func upgradeStateFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	var priorStateData resourceModelV0
	resp.Diagnostics.Append(req.State.Get(ctx, &priorStateData)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newStateData := WorkerVersionModel{
		ID:                 priorStateData.ID,
		AccountID:          priorStateData.AccountID,
		WorkerID:           priorStateData.WorkerID,
		CompatibilityDate:  priorStateData.CompatibilityDate,
		MainModule:         priorStateData.MainModule,
		Migrations:         priorStateData.Migrations,
		Modules:            priorStateData.Modules,
		Placement:          priorStateData.Placement,
		UsageModel:         priorStateData.UsageModel,
		CompatibilityFlags: priorStateData.CompatibilityFlags,
		Annotations:        priorStateData.Annotations,
		Bindings:           priorStateData.Bindings,
		Limits:             priorStateData.Limits,
		CreatedOn:          priorStateData.CreatedOn,
		Number:             priorStateData.Number,
		Source:             priorStateData.Source,
		MainScriptBase64:   priorStateData.MainScriptBase64,
		StartupTimeMs:      priorStateData.StartupTimeMs,
	}

	if priorStateData.Assets != nil {
		newStateData.Assets = &WorkerVersionAssetsModel{
			JWT:                 priorStateData.Assets.JWT,
			Directory:           priorStateData.Assets.Directory,
			AssetManifestSHA256: priorStateData.Assets.AssetManifestSHA256,
		}
		if !priorStateData.Assets.Config.IsNull() {
			config, diags := priorStateData.Assets.Config.Value(ctx)
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}
			if config != nil {
				var newConfig WorkerVersionAssetsConfigModel
				newConfig.HTMLHandling = config.HTMLHandling
				newConfig.NotFoundHandling = config.NotFoundHandling
				newConfig.RunWorkerFirst = customfield.RawNormalizedDynamicValueFrom(config.RunWorkerFirst)

				newStateData.Assets.Config, diags = customfield.NewObject(ctx, &newConfig)
				resp.Diagnostics.Append(diags...)
				if resp.Diagnostics.HasError() {
					return
				}
			}
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newStateData)...)
}

type resourceModelV0 struct {
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
	Assets             *resourceModelV0AssetsModel                              `tfsdk:"assets" json:"assets,optional"`
	Bindings           customfield.NestedObjectList[WorkerVersionBindingsModel] `tfsdk:"bindings" json:"bindings,optional"`
	Limits             customfield.NestedObject[WorkerVersionLimitsModel]       `tfsdk:"limits" json:"limits,computed_optional"`
	CreatedOn          timetypes.RFC3339                                        `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Number             types.Int64                                              `tfsdk:"number" json:"number,computed"`
	Source             types.String                                             `tfsdk:"source" json:"source,computed"`
	MainScriptBase64   types.String                                             `tfsdk:"main_script_base64" json:"main_script_base64,computed"`
	StartupTimeMs      types.Int64                                              `tfsdk:"startup_time_ms" json:"startup_time_ms,computed"`
}

type resourceModelV0AssetsModel struct {
	Config              customfield.NestedObject[resourceModelV0AssetsConfigModel] `tfsdk:"config" json:"config,optional"`
	JWT                 types.String                                               `tfsdk:"jwt" json:"jwt,optional"`
	Directory           types.String                                               `tfsdk:"directory" json:"-,optional"`
	AssetManifestSHA256 types.String                                               `tfsdk:"asset_manifest_sha256" json:"-,computed"`
}

type resourceModelV0AssetsConfigModel struct {
	HTMLHandling     types.String                   `tfsdk:"html_handling" json:"html_handling,optional"`
	NotFoundHandling types.String                   `tfsdk:"not_found_handling" json:"not_found_handling,optional"`
	RunWorkerFirst   customfield.List[types.String] `tfsdk:"run_worker_first" json:"run_worker_first,optional"`
}
