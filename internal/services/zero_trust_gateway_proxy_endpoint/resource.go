// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_proxy_endpoint

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*ZeroTrustGatewayProxyEndpointResource)(nil)
var _ resource.ResourceWithModifyPlan = (*ZeroTrustGatewayProxyEndpointResource)(nil)
var _ resource.ResourceWithImportState = (*ZeroTrustGatewayProxyEndpointResource)(nil)

func NewResource() resource.Resource {
	return &ZeroTrustGatewayProxyEndpointResource{}
}

// ZeroTrustGatewayProxyEndpointResource defines the resource implementation.
type ZeroTrustGatewayProxyEndpointResource struct {
	client *cloudflare.Client
}

func (r *ZeroTrustGatewayProxyEndpointResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zero_trust_gateway_proxy_endpoint"
}

func (r *ZeroTrustGatewayProxyEndpointResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ZeroTrustGatewayProxyEndpointResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ZeroTrustGatewayProxyEndpointModel

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
	env := ZeroTrustGatewayProxyEndpointResultEnvelope{*data}
	_, err = r.client.ZeroTrust.Gateway.ProxyEndpoints.New(
		ctx,
		zero_trust.GatewayProxyEndpointNewParams{
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

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustGatewayProxyEndpointResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ZeroTrustGatewayProxyEndpointModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *ZeroTrustGatewayProxyEndpointModel

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
	env := ZeroTrustGatewayProxyEndpointResultEnvelope{*data}
	_, err = r.client.ZeroTrust.Gateway.ProxyEndpoints.Edit(
		ctx,
		data.ID.ValueString(),
		zero_trust.GatewayProxyEndpointEditParams{
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

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustGatewayProxyEndpointResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ZeroTrustGatewayProxyEndpointModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := ZeroTrustGatewayProxyEndpointResultEnvelope{*data}
	_, err := r.client.ZeroTrust.Gateway.ProxyEndpoints.Get(
		ctx,
		data.ID.ValueString(),
		zero_trust.GatewayProxyEndpointGetParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
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

func (r *ZeroTrustGatewayProxyEndpointResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ZeroTrustGatewayProxyEndpointModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.ZeroTrust.Gateway.ProxyEndpoints.Delete(
		ctx,
		data.ID.ValueString(),
		zero_trust.GatewayProxyEndpointDeleteParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustGatewayProxyEndpointResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data *ZeroTrustGatewayProxyEndpointModel

	path_account_id := ""
	path_proxy_endpoint_id := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<account_id>/<proxy_endpoint_id>",
		&path_account_id,
		&path_proxy_endpoint_id,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := ZeroTrustGatewayProxyEndpointResultEnvelope{*data}
	_, err := r.client.ZeroTrust.Gateway.ProxyEndpoints.Get(
		ctx,
		path_proxy_endpoint_id,
		zero_trust.GatewayProxyEndpointGetParams{
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
	err = apijson.UnmarshalComputed(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustGatewayProxyEndpointResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
