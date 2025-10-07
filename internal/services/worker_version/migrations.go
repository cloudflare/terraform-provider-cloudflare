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

// The schema is identical to the schema in schema.go, except the version is 0
// and the assets.config.run_worker_first field is a list instead of a
// dynamic value.
func resourceSchemaV0(ctx context.Context) *schema.Schema {
	resourceSchemaLatest := ResourceSchema(ctx)
	s := resourceSchemaLatest
	s.Version = 0

	// Get assets attribute
	assetsAttr := s.Attributes["assets"].(schema.SingleNestedAttribute)
	configAttr := assetsAttr.Attributes["config"].(schema.SingleNestedAttribute)

	// Replace run_worker_first with list version
	configAttrs := make(map[string]schema.Attribute)
	for k, v := range configAttr.Attributes {
		configAttrs[k] = v
	}
	configAttrs["run_worker_first"] = schema.ListAttribute{
		Description: "Contains a list path rules to control routing to either the Worker or assets. Glob (*) and negative (!) rules are supported. Rules must start with either '/' or '!/'. At least one non-negative rule must be provided, and negative rules have higher precedence than non-negative rules.",
		Computed:    true,
		Optional:    true,
		CustomType:  customfield.NewListType[types.String](ctx),
		ElementType: types.StringType,
	}

	// Rebuild config
	configAttr.Attributes = configAttrs

	// Rebuild assets
	assetsAttrs := make(map[string]schema.Attribute)
	for k, v := range assetsAttr.Attributes {
		assetsAttrs[k] = v
	}
	assetsAttrs["config"] = configAttr
	assetsAttr.Attributes = assetsAttrs

	// Rebuild schema
	attrs := make(map[string]schema.Attribute)
	for k, v := range s.Attributes {
		attrs[k] = v
	}
	attrs["assets"] = assetsAttr
	s.Attributes = attrs

	return &s
}

// State upgrade function from version 0 to version 1. This converts
// assets.config.run_worker_first from a list value to a dynamic value
// (either boolean or list of strings).
func upgradeStateFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	// Read the raw state to avoid type issues with nested objects
	type rawAssetsConfigV0 struct {
		HTMLHandling     types.String                   `tfsdk:"html_handling"`
		NotFoundHandling types.String                   `tfsdk:"not_found_handling"`
		RunWorkerFirst   customfield.List[types.String] `tfsdk:"run_worker_first"`
	}

	type rawAssetsV0 struct {
		Config              *rawAssetsConfigV0 `tfsdk:"config"`
		JWT                 types.String       `tfsdk:"jwt"`
		Directory           types.String       `tfsdk:"directory"`
		AssetManifestSHA256 types.String       `tfsdk:"asset_manifest_sha256"`
	}

	var priorStateData struct {
		ID                 types.String                                             `tfsdk:"id"`
		AccountID          types.String                                             `tfsdk:"account_id"`
		WorkerID           types.String                                             `tfsdk:"worker_id"`
		CompatibilityDate  types.String                                             `tfsdk:"compatibility_date"`
		MainModule         types.String                                             `tfsdk:"main_module"`
		Migrations         *WorkerVersionMigrationsModel                            `tfsdk:"migrations"`
		Modules            *[]*WorkerVersionModulesModel                            `tfsdk:"modules"`
		Placement          *WorkerVersionPlacementModel                             `tfsdk:"placement"`
		UsageModel         types.String                                             `tfsdk:"usage_model"`
		CompatibilityFlags customfield.Set[types.String]                            `tfsdk:"compatibility_flags"`
		Annotations        customfield.NestedObject[WorkerVersionAnnotationsModel]  `tfsdk:"annotations"`
		Assets             *rawAssetsV0                                             `tfsdk:"assets"`
		Bindings           customfield.NestedObjectList[WorkerVersionBindingsModel] `tfsdk:"bindings"`
		Limits             customfield.NestedObject[WorkerVersionLimitsModel]       `tfsdk:"limits"`
		CreatedOn          timetypes.RFC3339                                        `tfsdk:"created_on"`
		Number             types.Int64                                              `tfsdk:"number"`
		Source             types.String                                             `tfsdk:"source"`
	}

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
	}

	if priorStateData.Assets != nil {
		newAssets := &WorkerVersionAssetsModel{
			JWT:                 priorStateData.Assets.JWT,
			Directory:           priorStateData.Assets.Directory,
			AssetManifestSHA256: priorStateData.Assets.AssetManifestSHA256,
			Config:              customfield.NullObject[WorkerVersionAssetsConfigModel](ctx),
		}

		if priorStateData.Assets.Config != nil {
			// Convert the list to a dynamic value
			configModel := WorkerVersionAssetsConfigModel{
				HTMLHandling:     priorStateData.Assets.Config.HTMLHandling,
				NotFoundHandling: priorStateData.Assets.Config.NotFoundHandling,
				RunWorkerFirst:   customfield.RawNormalizedDynamicValueFrom(priorStateData.Assets.Config.RunWorkerFirst.ListValue),
			}
			newConfig, diags := customfield.NewObject(ctx, &configModel)
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}
			newAssets.Config = newConfig
		}

		newStateData.Assets = newAssets
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
}

type resourceModelV0AssetsModel struct {
	Config              customfield.NestedObject[resourceModelV0AssetsConfigModel] `tfsdk:"config" json:"config,computed_optional"`
	JWT                 types.String                                               `tfsdk:"jwt" json:"jwt,optional"`
	Directory           types.String                                               `tfsdk:"directory" json:"-"`
	AssetManifestSHA256 types.String                                               `tfsdk:"asset_manifest_sha256" json:"asset_manifest_sha256,computed"`
}

type resourceModelV0AssetsConfigModel struct {
	HTMLHandling     types.String                   `tfsdk:"html_handling" json:"html_handling,computed_optional"`
	NotFoundHandling types.String                   `tfsdk:"not_found_handling" json:"not_found_handling,computed_optional"`
	RunWorkerFirst   customfield.List[types.String] `tfsdk:"run_worker_first" json:"run_worker_first,optional"`
}
