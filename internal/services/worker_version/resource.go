// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package worker_version

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/workers"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*WorkerVersionResource)(nil)
var _ resource.ResourceWithModifyPlan = (*WorkerVersionResource)(nil)
var _ resource.ResourceWithImportState = (*WorkerVersionResource)(nil)

func NewResource() resource.Resource {
	return &WorkerVersionResource{}
}

// WorkerVersionResource defines the resource implementation.
type WorkerVersionResource struct {
	client *cloudflare.Client
}

func (r *WorkerVersionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_worker_version"
}

func (r *WorkerVersionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*cloudflare.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *cloudflare.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *WorkerVersionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *WorkerVersionModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save planned placement before API call; versions endpoint does not accept
	// or return placement — it must be applied separately via the settings API.
	planPlacement := data.Placement

	var assets *WorkerVersionAssetsModel
	if data.Assets != nil {
		assets = &WorkerVersionAssetsModel{
			Config:              data.Assets.Config,
			JWT:                 data.Assets.JWT,
			Directory:           data.Assets.Directory,
			AssetManifestSHA256: data.Assets.AssetManifestSHA256,
		}
	}
	err := handleAssets(ctx, r.client, data)
	if err != nil {
		resp.Diagnostics.AddError("failed to upload assets", err.Error())
		return
	}

	var planModules *[]*WorkerVersionModulesModel
	if data.Modules != nil {
		planModules = data.Modules

		copied := make([]*WorkerVersionModulesModel, len(*data.Modules))
		for i, mod := range *data.Modules {
			modCopy := *mod
			copied[i] = &modCopy

			if !mod.ContentFile.IsNull() {
				content, err := readFile(mod.ContentFile.ValueString())
				if err != nil {
					resp.Diagnostics.AddError("Error reading file", err.Error())
					return
				}
				copied[i].ContentBase64 = types.StringValue(base64.StdEncoding.EncodeToString([]byte(content)))
			}
		}
		data.Modules = &copied
	}

	// Bindings as ordered in the plan. Terraform expects bindings written to
	// state to appear in the same order as the plan.
	planBindings := data.Bindings

	var diags diag.Diagnostics
	// Reorder plan bindings to be sorted in ascending order by name, which
	// matches the order that the API returns them. This is important for
	// apijson.UnmarshalComputed to work correctly. If the unmarshal target
	// doesn't match the order that the API returns the bindings, the unmarshal
	// operation will assign computed properties to the wrong bindings.
	data.Bindings, diags = SortBindingsByName(ctx, planBindings)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := data.MarshalJSON()
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := WorkerVersionResultEnvelope{*data}
	_, err = r.client.Workers.Beta.Workers.Versions.New(
		ctx,
		data.WorkerID.ValueString(),
		workers.BetaWorkerVersionNewParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithRequestBody("application/json", dataBytes),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.UnmarshalComputed(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result

	// The API returns database_id for D1 bindings but not for other types.
	// Null out unknown database_id values to prevent "unknown after apply" errors.
	if !data.Bindings.IsNull() && !data.Bindings.IsUnknown() {
		var bindingsList []WorkerVersionBindingsModel
		diags = data.Bindings.ElementsAs(ctx, &bindingsList, true)
		resp.Diagnostics.Append(diags...)
		for i := range bindingsList {
			if bindingsList[i].DatabaseID.IsUnknown() {
				bindingsList[i].DatabaseID = types.StringNull()
			}
		}
		data.Bindings, diags = customfield.NewObjectList(ctx, bindingsList)
		resp.Diagnostics.Append(diags...)
	}

	if data.Modules != nil && planModules != nil {
		apiModuleNameMap := make(map[string]*WorkerVersionModulesModel)
		for _, mod := range *data.Modules {
			apiModuleNameMap[mod.Name.ValueString()] = mod
		}

		for _, planMod := range *planModules {
			if apiMod, ok := apiModuleNameMap[planMod.Name.ValueString()]; ok {
				contentBase64 := apiMod.ContentBase64.ValueString()
				content, err := base64.StdEncoding.DecodeString(contentBase64)
				if err != nil {
					resp.Diagnostics.AddError("Create Error", err.Error())
					return
				}
				contentSHA256, err := calculateStringHash(string(content))
				if err != nil {
					resp.Diagnostics.AddError("Create Error", err.Error())
					return
				}
				planMod.ContentSHA256 = types.StringValue(contentSHA256)
			}
		}
	}
	data.Modules = planModules

	if assets != nil && data.Assets != nil {
		assets.Config = data.Assets.Config
	}

	data.Assets = assets
	// Finally, reorder refreshed bindings to match the plan, now that computed
	// properties have been filled in.
	data.Bindings, diags = SortRefreshedBindingsToMatchPrevious(
		ctx,
		data.Bindings,
		planBindings,
	)
	resp.Diagnostics.Append(diags...)

	// Restore planned placement and apply it via the script settings endpoint.
	// The versions API does not accept or return placement; the correct endpoint
	// is PATCH /accounts/{id}/workers/scripts/{name}/settings.
	data.Placement = planPlacement
	if settingsErr := r.applyPlacementSettings(ctx, data.AccountID.ValueString(), data.WorkerID.ValueString(), planPlacement); settingsErr != nil {
		resp.Diagnostics.AddError("failed to apply placement settings", settingsErr.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// This resource is immutable at the API level, but can be updated "in-place" if
// the only changes are to provider-only attributes (namely the content_file
// module attribute). Allowing "in-place" updates to these attributes makes it
// possible to import this resource without destroying and re-creating it in the
// process.
func (r *WorkerVersionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan *WorkerVersionModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Computed properties are marked as unknown in the plan and can't be copied
	// to state. The modules attribute is the only attribute that can be updated
	// in-place, so we only copy that attribute to state.
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("modules"), plan.Modules)...)
}

func (r *WorkerVersionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *WorkerVersionModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	assets := data.Assets
	var stateModules *[]*WorkerVersionModulesModel
	if data.Modules != nil {
		copied := make([]*WorkerVersionModulesModel, len(*data.Modules))
		for i, mod := range *data.Modules {
			modCopy := *mod
			copied[i] = &modCopy
		}
		stateModules = &copied
	}

	res := new(http.Response)
	env := WorkerVersionResultEnvelope{*data}
	_, err := r.client.Workers.Beta.Workers.Versions.Get(
		ctx,
		data.WorkerID.ValueString(),
		data.ID.ValueString(),
		workers.BetaWorkerVersionGetParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
			Include:   cloudflare.F(workers.BetaWorkerVersionGetParamsIncludeModules),
		},
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if res != nil && res.StatusCode == 404 {
		resp.Diagnostics.AddWarning("Resource not found", "The resource was not found on the server and will be removed from state.")
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.Unmarshal(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result
	data.Assets = assets

	// The API returns database_id for D1 bindings but not for other types.
	// Null out unknown database_id values to prevent drift.
	if !data.Bindings.IsNull() && !data.Bindings.IsUnknown() {
		var readBindingsList []WorkerVersionBindingsModel
		readDiags := data.Bindings.ElementsAs(ctx, &readBindingsList, true)
		resp.Diagnostics.Append(readDiags...)
		for i := range readBindingsList {
			if readBindingsList[i].DatabaseID.IsUnknown() {
				readBindingsList[i].DatabaseID = types.StringNull()
			}
		}
		data.Bindings, readDiags = customfield.NewObjectList(ctx, readBindingsList)
		resp.Diagnostics.Append(readDiags...)
	}

	apiModuleNameMap := make(map[string]*WorkerVersionModulesModel)
	if data.Modules != nil {
		for _, mod := range *data.Modules {
			apiModuleNameMap[mod.Name.ValueString()] = mod
		}
	}

	if stateModules != nil {
		for _, stateMod := range *stateModules {
			if apiMod, ok := apiModuleNameMap[stateMod.Name.ValueString()]; ok {
				contentBase64 := apiMod.ContentBase64.ValueString()
				content, err := base64.StdEncoding.DecodeString(contentBase64)
				if err != nil {
					resp.Diagnostics.AddError("Refresh Error", err.Error())
					return
				}
				contentSHA256, err := calculateStringHash(string(content))
				if err != nil {
					resp.Diagnostics.AddError("Refresh Error", err.Error())
					return
				}
				stateMod.ContentSHA256 = types.StringValue(contentSHA256)

				if stateMod.ContentBase64.IsNull() || stateMod.ContentBase64.IsUnknown() {
					// content_file was used, keep it as is
				} else {
					// content_base64 was used, update it from API
					stateMod.ContentBase64 = apiMod.ContentBase64
				}
			}
		}
		data.Modules = stateModules
	}

	// restore any secret_text `text` values from state since they aren't returned by the API
	var state *WorkerVersionModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	var diags diag.Diagnostics
	data.Bindings, diags = UpdateSecretTextsFromState(
		ctx,
		data.Bindings,
		state.Bindings,
	)
	resp.Diagnostics.Append(diags...)
	data.Bindings, diags = SortRefreshedBindingsToMatchPrevious(
		ctx,
		data.Bindings,
		state.Bindings,
	)
	resp.Diagnostics.Append(diags...)

	// Read placement from the script settings endpoint — the versions API does
	// not return placement since it is a script-level (not version-level) setting.
	settings, settingsErr := r.client.Workers.Scripts.ScriptAndVersionSettings.Get(
		ctx,
		data.WorkerID.ValueString(),
		workers.ScriptScriptAndVersionSettingGetParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if settingsErr != nil {
		resp.Diagnostics.AddError("failed to read script settings", settingsErr.Error())
		return
	}
	data.Placement = placementFromSettings(settings.Placement)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkerVersionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

}

func (r *WorkerVersionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data = new(WorkerVersionModel)

	path_account_id := ""
	path_worker_id := ""
	path_version_id := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<account_id>/<worker_id>/<version_id>",
		&path_account_id,
		&path_worker_id,
		&path_version_id,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.AccountID = types.StringValue(path_account_id)
	data.WorkerID = types.StringValue(path_worker_id)
	data.ID = types.StringValue(path_version_id)

	res := new(http.Response)
	env := WorkerVersionResultEnvelope{*data}
	_, err := r.client.Workers.Beta.Workers.Versions.Get(
		ctx,
		path_worker_id,
		path_version_id,
		workers.BetaWorkerVersionGetParams{
			AccountID: cloudflare.F(path_account_id),
			Include:   cloudflare.F(workers.BetaWorkerVersionGetParamsIncludeModules),
		},
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.Unmarshal(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result

	settings, settingsErr := r.client.Workers.Scripts.ScriptAndVersionSettings.Get(
		ctx,
		path_worker_id,
		workers.ScriptScriptAndVersionSettingGetParams{
			AccountID: cloudflare.F(path_account_id),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if settingsErr != nil {
		resp.Diagnostics.AddError("failed to read script settings", settingsErr.Error())
		return
	}
	data.Placement = placementFromSettings(settings.Placement)

	if data.Modules != nil {
		for _, mod := range *data.Modules {
			contentBase64 := mod.ContentBase64.ValueString()
			content, err := base64.StdEncoding.DecodeString(contentBase64)
			if err != nil {
				resp.Diagnostics.AddError("Import Error", err.Error())
				return
			}
			contentSHA256, err := calculateStringHash(string(content))
			if err != nil {
				resp.Diagnostics.AddError("Import Error", err.Error())
				return
			}
			mod.ContentSHA256 = types.StringValue(contentSHA256)
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkerVersionResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}

// applyPlacementSettings sends a PATCH to the script-and-version-settings endpoint
// to set or clear the placement configuration. The versions API endpoint does not
// accept placement, so this must be called after version creation.
// Only one placement variant (region, hostname, host, or mode) is sent per call;
// the Cloudflare placement field is a union type. The target field is not yet supported.
func (r *WorkerVersionResource) applyPlacementSettings(ctx context.Context, accountID, scriptName string, placement *WorkerVersionPlacementModel) error {
	settings := workers.ScriptScriptAndVersionSettingEditParamsSettings{
		Placement: cloudflare.Null[workers.ScriptScriptAndVersionSettingEditParamsSettingsPlacementUnion](),
	}

	if placement != nil {
		// TODO: add Target support once the SDK exposes a union variant for it.
		switch {
		case !placement.Region.IsNull() && !placement.Region.IsUnknown() && placement.Region.ValueString() != "":
			settings.Placement = cloudflare.F[workers.ScriptScriptAndVersionSettingEditParamsSettingsPlacementUnion](
				workers.ScriptScriptAndVersionSettingEditParamsSettingsPlacementRegion{
					Region: cloudflare.F(placement.Region.ValueString()),
				},
			)
		case !placement.Hostname.IsNull() && !placement.Hostname.IsUnknown() && placement.Hostname.ValueString() != "":
			settings.Placement = cloudflare.F[workers.ScriptScriptAndVersionSettingEditParamsSettingsPlacementUnion](
				workers.ScriptScriptAndVersionSettingEditParamsSettingsPlacementHostname{
					Hostname: cloudflare.F(placement.Hostname.ValueString()),
				},
			)
		case !placement.Host.IsNull() && !placement.Host.IsUnknown() && placement.Host.ValueString() != "":
			settings.Placement = cloudflare.F[workers.ScriptScriptAndVersionSettingEditParamsSettingsPlacementUnion](
				workers.ScriptScriptAndVersionSettingEditParamsSettingsPlacementHost{
					Host: cloudflare.F(placement.Host.ValueString()),
				},
			)
		case !placement.Mode.IsNull() && !placement.Mode.IsUnknown() && placement.Mode.ValueString() != "":
			settings.Placement = cloudflare.F[workers.ScriptScriptAndVersionSettingEditParamsSettingsPlacementUnion](
				workers.ScriptScriptAndVersionSettingEditParamsSettingsPlacementMode{
					Mode: cloudflare.F(workers.ScriptScriptAndVersionSettingEditParamsSettingsPlacementModeMode(placement.Mode.ValueString())),
				},
			)
		}
	}

	_, err := r.client.Workers.Scripts.ScriptAndVersionSettings.Edit(
		ctx,
		scriptName,
		workers.ScriptScriptAndVersionSettingEditParams{
			AccountID: cloudflare.F(accountID),
			Settings:  cloudflare.F(settings),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	return err
}

// placementFromSettings converts the placement field from a script settings
// response into the WorkerVersionPlacementModel used in state. Returns nil
// when the API reports no active placement configuration.
// Note: the target field is not yet read back (see applyPlacementSettings TODO).
func placementFromSettings(p workers.ScriptScriptAndVersionSettingGetResponsePlacement) *WorkerVersionPlacementModel {
	if p.Region == "" && string(p.Mode) == "" && p.Hostname == "" && p.Host == "" {
		return nil
	}
	m := &WorkerVersionPlacementModel{
		Region:   types.StringNull(),
		Mode:     types.StringNull(),
		Hostname: types.StringNull(),
		Host:     types.StringNull(),
		Target:   nil,
	}
	if p.Region != "" {
		m.Region = types.StringValue(p.Region)
	}
	if string(p.Mode) != "" {
		m.Mode = types.StringValue(string(p.Mode))
	}
	if p.Hostname != "" {
		m.Hostname = types.StringValue(p.Hostname)
	}
	if p.Host != "" {
		m.Host = types.StringValue(p.Host)
	}
	return m
}
