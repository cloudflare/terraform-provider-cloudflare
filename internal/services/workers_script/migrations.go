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
			// Handle v4 -> v5 migration: name -> script_name
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"account_id": schema.StringAttribute{
						Required: true,
					},
					"name": schema.StringAttribute{
						Required: true,
					},
					"script_name": schema.StringAttribute{
						Optional: true,
					},
					"content": schema.StringAttribute{
						Optional: true,
					},
					"content_file": schema.StringAttribute{
						Optional: true,
					},
					"content_sha256": schema.StringAttribute{
						Optional: true,
					},
					"content_type": schema.StringAttribute{
						Optional: true,
					},
					"compatibility_date": schema.StringAttribute{
						Optional: true,
					},
					"compatibility_flags": schema.SetAttribute{
						Optional:    true,
						ElementType: types.StringType,
					},
					"created_on": schema.StringAttribute{
						Computed: true,
					},
					"etag": schema.StringAttribute{
						Computed: true,
					},
					"modified_on": schema.StringAttribute{
						Computed: true,
					},
					"has_assets": schema.BoolAttribute{
						Computed: true,
					},
					"has_modules": schema.BoolAttribute{
						Computed: true,
					},
					"startup_time_ms": schema.Float64Attribute{
						Computed: true,
					},
					"logpush": schema.BoolAttribute{
						Optional: true,
					},
					"usage_model": schema.StringAttribute{
						Optional: true,
					},
				},
			},
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var priorStateData struct {
					ID                  types.String  `tfsdk:"id"`
					AccountID           types.String  `tfsdk:"account_id"`
					Name                types.String  `tfsdk:"name"`
					ScriptName          types.String  `tfsdk:"script_name"`
					Content             types.String  `tfsdk:"content"`
					ContentFile         types.String  `tfsdk:"content_file"`
					ContentSHA256       types.String  `tfsdk:"content_sha256"`
					ContentType         types.String  `tfsdk:"content_type"`
					CompatibilityDate   types.String  `tfsdk:"compatibility_date"`
					CompatibilityFlags  types.Set     `tfsdk:"compatibility_flags"`
					CreatedOn           types.String  `tfsdk:"created_on"`
					Etag                types.String  `tfsdk:"etag"`
					ModifiedOn          types.String  `tfsdk:"modified_on"`
					HasAssets           types.Bool    `tfsdk:"has_assets"`
					HasModules          types.Bool    `tfsdk:"has_modules"`
					StartupTimeMs       types.Float64 `tfsdk:"startup_time_ms"`
					Logpush             types.Bool    `tfsdk:"logpush"`
					UsageModel          types.String  `tfsdk:"usage_model"`
				}

				resp.Diagnostics.Append(req.State.Get(ctx, &priorStateData)...)
				if resp.Diagnostics.HasError() {
					return
				}

				// Initialize new state
				var newState = WorkersScriptModel{
					ID:            priorStateData.ID,
					AccountID:     priorStateData.AccountID,
					ScriptName:    priorStateData.ScriptName,
					Content:       priorStateData.Content,
					ContentFile:   priorStateData.ContentFile,
					ContentSHA256: priorStateData.ContentSHA256,
					ContentType:   priorStateData.ContentType,
					Etag:          priorStateData.Etag,
					HasAssets:     priorStateData.HasAssets,
					HasModules:    priorStateData.HasModules,
					WorkersScriptMetadataModel: WorkersScriptMetadataModel{
						CompatibilityDate:  priorStateData.CompatibilityDate,
						CompatibilityFlags: customfield.NullSet[types.String](ctx),
						Logpush:            priorStateData.Logpush,
						UsageModel:         priorStateData.UsageModel,
					},
				}

				// Convert RFC3339 strings to timetypes.RFC3339
				if !priorStateData.CreatedOn.IsNull() {
					newState.CreatedOn = timetypes.NewRFC3339ValueMust(priorStateData.CreatedOn.ValueString())
				}
				if !priorStateData.ModifiedOn.IsNull() {
					newState.ModifiedOn = timetypes.NewRFC3339ValueMust(priorStateData.ModifiedOn.ValueString())
				}

				// Convert Float64 to Int64 for startup time
				if !priorStateData.StartupTimeMs.IsNull() {
					newState.StartupTimeMs = types.Int64Value(int64(priorStateData.StartupTimeMs.ValueFloat64()))
				}

				// Handle compatibility flags set conversion if needed
				if !priorStateData.CompatibilityFlags.IsNull() {
					flagsSet, diags := customfield.NewSet[types.String](ctx, priorStateData.CompatibilityFlags.Elements())
					if !diags.HasError() {
						newState.WorkersScriptMetadataModel.CompatibilityFlags = flagsSet
					}
				}

				// If script_name is null but name exists, migrate it
				if newState.ScriptName.IsNull() && !priorStateData.Name.IsNull() {
					newState.ScriptName = priorStateData.Name
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
			},
		},
	}
}
