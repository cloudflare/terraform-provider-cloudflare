// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package worker_version

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/workers"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
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
