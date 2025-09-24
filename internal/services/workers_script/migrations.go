// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_script

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithUpgradeState = (*WorkersScriptResource)(nil)

func (r *WorkersScriptResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		// State upgrade from schema version 0 to 1 
		// Migrates run_worker_first from boolean to dynamic type
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"script_name": schema.StringAttribute{
						Required: true,
					},
					"account_id": schema.StringAttribute{
						Required: true,
					},
					"assets": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"config": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"run_worker_first": schema.BoolAttribute{
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var priorStateData struct {
					Assets types.Object `tfsdk:"assets"`
				}

				resp.Diagnostics.Append(req.State.Get(ctx, &priorStateData)...)
				if resp.Diagnostics.HasError() {
					return
				}

				// If assets is null or unknown, no migration needed
				if priorStateData.Assets.IsNull() || priorStateData.Assets.IsUnknown() {
					resp.Diagnostics.Append(req.State.Set(ctx, req.State)...)
					return
				}

				// Extract assets.config.run_worker_first
				assetsAttrs := priorStateData.Assets.Attributes()
				if config, exists := assetsAttrs["config"]; exists && !config.IsNull() && !config.IsUnknown() {
					configObj := config.(types.Object)
					configAttrs := configObj.Attributes()
					
					if runWorkerFirst, exists := configAttrs["run_worker_first"]; exists && !runWorkerFirst.IsNull() && !runWorkerFirst.IsUnknown() {
						// Convert boolean to dynamic format
						boolVal := runWorkerFirst.(types.Bool)
						
						// Create new dynamic value with proper structure
						dynamicValue := types.DynamicValue(types.BoolValue(boolVal.ValueBool()))
						
						// Update the config attributes
						newConfigAttrs := make(map[string]attr.Value)
						for k, v := range configAttrs {
							if k == "run_worker_first" {
								newConfigAttrs[k] = dynamicValue
							} else {
								newConfigAttrs[k] = v
							}
						}
						
						// Create new config object
						newConfig, diags := types.ObjectValue(configObj.AttributeTypes(ctx), newConfigAttrs)
						resp.Diagnostics.Append(diags...)
						if resp.Diagnostics.HasError() {
							return
						}
						
						// Update the assets attributes
						newAssetsAttrs := make(map[string]attr.Value)
						for k, v := range assetsAttrs {
							if k == "config" {
								newAssetsAttrs[k] = newConfig
							} else {
								newAssetsAttrs[k] = v
							}
						}
						
						// Create new assets object
						newAssets, diags := types.ObjectValue(priorStateData.Assets.AttributeTypes(ctx), newAssetsAttrs)
						resp.Diagnostics.Append(diags...)
						if resp.Diagnostics.HasError() {
							return
						}
						
						// Set the new state with migrated data
						var newStateData struct {
							Assets types.Object `tfsdk:"assets"`
						}
						newStateData.Assets = newAssets
						
						resp.Diagnostics.Append(req.State.SetAttribute(ctx, path.Root("assets"), newAssets)...)
						return
					}
				}
				
				// No migration needed, copy state as-is
				resp.Diagnostics.Append(req.State.Set(ctx, req.State)...)
			},
		},
	}
}
