// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_script

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/workers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*WorkersScriptResource)(nil)
var _ resource.ResourceWithModifyPlan = (*WorkersScriptResource)(nil)
var _ resource.ResourceWithImportState = (*WorkersScriptResource)(nil)

func NewResource() resource.Resource {
	return &WorkersScriptResource{}
}

// WorkersScriptResource defines the resource implementation.
type WorkersScriptResource struct {
	client *cloudflare.Client
}

func (r *WorkersScriptResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_workers_script"
}

func (r *WorkersScriptResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *WorkersScriptResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *WorkersScriptModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, contentType, err := data.MarshalMultipart()
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize multipart http request", err.Error())
		return
	}
	res := new(http.Response)
	env := WorkersScriptResultEnvelope{*data}
	_, err = r.client.Workers.Scripts.Update(
		ctx,
		data.ScriptName.ValueString(),
		workers.ScriptUpdateParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithRequestBody(contentType, dataBytes),
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
	data.ID = data.ScriptName

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkersScriptResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *WorkersScriptModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *WorkersScriptModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, contentType, err := data.MarshalMultipart()
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize multipart http request", err.Error())
		return
	}
	res := new(http.Response)
	env := WorkersScriptResultEnvelope{*data}
	_, err = r.client.Workers.Scripts.Update(
		ctx,
		data.ScriptName.ValueString(),
		workers.ScriptUpdateParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithRequestBody(contentType, dataBytes),
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
	data.ID = data.ScriptName

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkersScriptResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *WorkersScriptModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	_, err := r.client.Workers.Scripts.Get(
		ctx,
		data.ScriptName.ValueString(),
		workers.ScriptGetParams{
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
	err = apijson.UnmarshalComputed(bytes, &data)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data.ID = data.ScriptName

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkersScriptResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *WorkersScriptModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Workers.Scripts.Delete(
		ctx,
		data.ScriptName.ValueString(),
		workers.ScriptDeleteParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	data.ID = data.ScriptName

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkersScriptResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data *WorkersScriptModel = new(WorkersScriptModel)

	path_account_id := ""
	path_script_name := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<account_id>/<script_name>",
		&path_account_id,
		&path_script_name,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.AccountID = types.StringValue(path_account_id)
	data.ScriptName = types.StringValue(path_script_name)

	res := new(http.Response)
	_, err := r.client.Workers.Scripts.Get(
		ctx,
		path_script_name,
		workers.ScriptGetParams{
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
	err = apijson.Unmarshal(bytes, &data)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data.ID = data.ScriptName

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkersScriptResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
