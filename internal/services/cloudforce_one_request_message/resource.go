// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_request_message

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/cloudforce_one"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*CloudforceOneRequestMessageResource)(nil)
var _ resource.ResourceWithModifyPlan = (*CloudforceOneRequestMessageResource)(nil)
var _ resource.ResourceWithImportState = (*CloudforceOneRequestMessageResource)(nil)

func NewResource() resource.Resource {
	return &CloudforceOneRequestMessageResource{}
}

// CloudforceOneRequestMessageResource defines the resource implementation.
type CloudforceOneRequestMessageResource struct {
	client *cloudflare.Client
}

func (r *CloudforceOneRequestMessageResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloudforce_one_request_message"
}

func (r *CloudforceOneRequestMessageResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *CloudforceOneRequestMessageResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *CloudforceOneRequestMessageModel

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
	env := CloudforceOneRequestMessageResultEnvelope{*data}
	_, err = r.client.CloudforceOne.Requests.Message.New(
		ctx,
		data.AccountIdentifier.ValueString(),
		data.RequestIdentifier.ValueString(),
		cloudforce_one.RequestMessageNewParams{},
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

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudforceOneRequestMessageResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *CloudforceOneRequestMessageModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *CloudforceOneRequestMessageModel

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
	env := CloudforceOneRequestMessageResultEnvelope{*data}
	_, err = r.client.CloudforceOne.Requests.Message.Update(
		ctx,
		data.AccountIdentifier.ValueString(),
		data.RequestIdentifier.ValueString(),
		data.ID.ValueInt64(),
		cloudforce_one.RequestMessageUpdateParams{},
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

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudforceOneRequestMessageResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *CloudforceOneRequestMessageModel

// 	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

// 	if resp.Diagnostics.HasError() {
// 		return
// 	}

// 	res := new(http.Response)
// 	env := CloudforceOneRequestMessageResultEnvelope{*data}
// 	_, err := r.client.CloudforceOne.Requests.Message.Get(
// 		ctx,
// 		data.AccountIdentifier.ValueString(),
// 		data.ID.ValueInt64(),
// 		cloudforce_one.RequestMessageGetParams{},
// 		option.WithResponseBodyInto(&res),
// 		option.WithMiddleware(logging.Middleware(ctx)),
// 	)
// 	if res != nil && res.StatusCode == 404 {
// 		resp.Diagnostics.AddWarning("Resource not found", "The resource was not found on the server and will be removed from state.")
// 		resp.State.RemoveResource(ctx)
// 		return
// 	}
// 	if err != nil {
// 		resp.Diagnostics.AddError("failed to make http request", err.Error())
// 		return
// 	}
// 	bytes, _ := io.ReadAll(res.Body)
// 	err = apijson.UnmarshalComputed(bytes, &env)
// 	if err != nil {
// 		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
// 		return
// 	}
// 	data = &env.Result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudforceOneRequestMessageResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *CloudforceOneRequestMessageModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.CloudforceOne.Requests.Message.Delete(
		ctx,
		data.AccountIdentifier.ValueString(),
		data.RequestIdentifier.ValueString(),
		data.ID.ValueInt64(),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudforceOneRequestMessageResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data *CloudforceOneRequestMessageModel = new(CloudforceOneRequestMessageModel)

	path_account_identifier := ""
	path_request_identifier := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<account_identifier>/<request_identifier>",
		&path_account_identifier,
		&path_request_identifier,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.AccountIdentifier = types.StringValue(path_account_identifier)
	data.RequestIdentifier = types.StringValue(path_request_identifier)

	res := new(http.Response)
	env := CloudforceOneRequestMessageResultEnvelope{*data}
	_, err := r.client.CloudforceOne.Requests.Message.Get(
		ctx,
		path_account_identifier,
		path_request_identifier,
		cloudforce_one.RequestMessageGetParams{},
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

func (r *CloudforceOneRequestMessageResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
