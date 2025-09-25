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

	modules := data.Modules
	if modules != nil {
		for _, mod := range *data.Modules {
			content, err := readFile(mod.ContentFile.ValueString())
			if err != nil {
				resp.Diagnostics.AddError("Error reading file", err.Error())
			}
			mod.ContentBase64 = types.StringValue(base64.StdEncoding.EncodeToString([]byte(content)))
		}
	}

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
	data.Modules = modules

	if assets != nil && data.Assets != nil {
		assets.Config = data.Assets.Config
	}

	data.Assets = assets

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkerVersionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Update is not supported for this resource
}

func (r *WorkerVersionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *WorkerVersionModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	assets := data.Assets // "assets" is not returned by the API, so preserve its state value
	stateModules := data.Modules

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

	// Refresh content_sha256 on each module
	moduleNameMap := make(map[string]*WorkerVersionModulesModel)
	if stateModules != nil {
		for _, mod := range *stateModules {
			moduleNameMap[mod.Name.ValueString()] = mod
		}
	}
	if data.Modules != nil {
		for _, mod := range *data.Modules {
			contentBase64 := mod.ContentBase64.ValueString()
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

			mod.ContentSHA256 = types.StringValue(contentSHA256)
			if stateMod, ok := moduleNameMap[mod.Name.ValueString()]; ok {
				mod.ContentFile = stateMod.ContentFile
			}
		}
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

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkerVersionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

}

func (r *WorkerVersionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data *WorkerVersionModel = new(WorkerVersionModel)

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

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkerVersionResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
