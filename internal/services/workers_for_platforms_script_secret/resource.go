// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms_script_secret

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/workers_for_platforms"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*WorkersForPlatformsScriptSecretResource)(nil)
var _ resource.ResourceWithModifyPlan = (*WorkersForPlatformsScriptSecretResource)(nil)
var _ resource.ResourceWithImportState = (*WorkersForPlatformsScriptSecretResource)(nil)

func NewResource() resource.Resource {
	return &WorkersForPlatformsScriptSecretResource{}
}

// WorkersForPlatformsScriptSecretResource defines the resource implementation.
type WorkersForPlatformsScriptSecretResource struct {
	client *cloudflare.Client
}

func (r *WorkersForPlatformsScriptSecretResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_workers_for_platforms_script_secret"
}

func (r *WorkersForPlatformsScriptSecretResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *WorkersForPlatformsScriptSecretResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *WorkersForPlatformsScriptSecretModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := data.MarshalJSON()
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := WorkersForPlatformsScriptSecretResultEnvelope{*data}
	_, err = r.client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.Secrets.Update(
		ctx,
		data.DispatchNamespace.ValueString(),
		data.ScriptName.ValueString(),
		workers_for_platforms.DispatchNamespaceScriptSecretUpdateParams{
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
	data.ID = data.Name

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkersForPlatformsScriptSecretResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *WorkersForPlatformsScriptSecretModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *WorkersForPlatformsScriptSecretModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := data.MarshalJSONForUpdate(*state)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := WorkersForPlatformsScriptSecretResultEnvelope{*data}
	_, err = r.client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.Secrets.Update(
		ctx,
		data.DispatchNamespace.ValueString(),
		data.ScriptName.ValueString(),
		workers_for_platforms.DispatchNamespaceScriptSecretUpdateParams{
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
	data.ID = data.Name

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkersForPlatformsScriptSecretResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *WorkersForPlatformsScriptSecretModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := WorkersForPlatformsScriptSecretResultEnvelope{*data}
	_, err := r.client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.Secrets.Get(
		ctx,
		data.DispatchNamespace.ValueString(),
		data.ScriptName.ValueString(),
		data.Name.ValueString(),
		workers_for_platforms.DispatchNamespaceScriptSecretGetParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
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
	data.ID = data.Name

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkersForPlatformsScriptSecretResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *WorkersForPlatformsScriptSecretModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.Secrets.Delete(
		ctx,
		data.DispatchNamespace.ValueString(),
		data.ScriptName.ValueString(),
		data.Name.ValueString(),
		workers_for_platforms.DispatchNamespaceScriptSecretDeleteParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	data.ID = data.Name

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkersForPlatformsScriptSecretResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data *WorkersForPlatformsScriptSecretModel = new(WorkersForPlatformsScriptSecretModel)

	path_account_id := ""
	path_dispatch_namespace := ""
	path_script_name := ""
	path_secret_name := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<account_id>/<dispatch_namespace>/<script_name>/<secret_name>",
		&path_account_id,
		&path_dispatch_namespace,
		&path_script_name,
		&path_secret_name,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.AccountID = types.StringValue(path_account_id)
	data.DispatchNamespace = types.StringValue(path_dispatch_namespace)
	data.ScriptName = types.StringValue(path_script_name)
	data.Name = types.StringValue(path_secret_name)

	res := new(http.Response)
	env := WorkersForPlatformsScriptSecretResultEnvelope{*data}
	_, err := r.client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.Secrets.Get(
		ctx,
		path_dispatch_namespace,
		path_script_name,
		path_secret_name,
		workers_for_platforms.DispatchNamespaceScriptSecretGetParams{
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
	data.ID = data.Name

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkersForPlatformsScriptSecretResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
