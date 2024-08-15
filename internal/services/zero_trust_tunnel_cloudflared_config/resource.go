// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared_config

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = &ZeroTrustTunnelCloudflaredConfigResource{}
var _ resource.ResourceWithModifyPlan = &ZeroTrustTunnelCloudflaredConfigResource{}
var _ resource.ResourceWithImportState = &ZeroTrustTunnelCloudflaredConfigResource{}

func NewResource() resource.Resource {
	return &ZeroTrustTunnelCloudflaredConfigResource{}
}

// ZeroTrustTunnelCloudflaredConfigResource defines the resource implementation.
type ZeroTrustTunnelCloudflaredConfigResource struct {
	client *cloudflare.Client
}

func (r *ZeroTrustTunnelCloudflaredConfigResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zero_trust_tunnel_cloudflared_config"
}

func (r *ZeroTrustTunnelCloudflaredConfigResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ZeroTrustTunnelCloudflaredConfigResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ZeroTrustTunnelCloudflaredConfigModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := apijson.Marshal(data)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := ZeroTrustTunnelCloudflaredConfigResultEnvelope{*data}
	_, err = r.client.ZeroTrust.Tunnels.Configurations.Update(
		ctx,
		data.TunnelID.ValueString(),
		zero_trust.TunnelConfigurationUpdateParams{
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
	err = apijson.Unmarshal(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result
	data.ID = data.TunnelID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustTunnelCloudflaredConfigResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ZeroTrustTunnelCloudflaredConfigModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := ZeroTrustTunnelCloudflaredConfigResultEnvelope{*data}
	_, err := r.client.ZeroTrust.Tunnels.Configurations.Get(
		ctx,
		data.TunnelID.ValueString(),
		zero_trust.TunnelConfigurationGetParams{
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
	err = apijson.Unmarshal(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result
	data.ID = data.TunnelID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustTunnelCloudflaredConfigResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data *ZeroTrustTunnelCloudflaredConfigModel

	path := strings.Split(req.ID, "/")
	if len(path) != 2 {
		resp.Diagnostics.AddError("Invalid ID", "expected urlencoded segments <account_id>/<tunnel_id>")
		return
	}
	path_account_id, err := url.PathUnescape(path[0])
	if err != nil {
		resp.Diagnostics.AddError("invalid urlencoded segment - <account_id>", fmt.Sprintf("%s -> %q", err.Error(), path[0]))
	}
	path_tunnel_id, err := url.PathUnescape(path[1])
	if err != nil {
		resp.Diagnostics.AddError("invalid urlencoded segment - <tunnel_id>", fmt.Sprintf("%s -> %q", err.Error(), path[1]))
	}
	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := ZeroTrustTunnelCloudflaredConfigResultEnvelope{*data}
	_, err = r.client.ZeroTrust.Tunnels.Configurations.Get(
		ctx,
		path_tunnel_id,
		zero_trust.TunnelConfigurationGetParams{
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
	data.ID = data.TunnelID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustTunnelCloudflaredConfigResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ZeroTrustTunnelCloudflaredConfigModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *ZeroTrustTunnelCloudflaredConfigModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := apijson.MarshalForUpdate(data, state)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := ZeroTrustTunnelCloudflaredConfigResultEnvelope{*data}
	_, err = r.client.ZeroTrust.Tunnels.Configurations.Update(
		ctx,
		data.TunnelID.ValueString(),
		zero_trust.TunnelConfigurationUpdateParams{
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
	err = apijson.Unmarshal(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result
	data.ID = data.TunnelID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustTunnelCloudflaredConfigResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

}

func (r *ZeroTrustTunnelCloudflaredConfigResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.State.Raw.IsNull() {
		resp.Diagnostics.AddWarning(
			"Resource Destruction Considerations",
			"This resource cannot be destroyed from Terraform. If you create this resource, it will be "+
				"present in the API until manually deleted.",
		)
	}
	if req.Plan.Raw.IsNull() {
		resp.Diagnostics.AddWarning(
			"Resource Destruction Considerations",
			"Applying this resource destruction will remove the resource from the Terraform state "+
				"but will not change it in the API. If you would like to destroy or reset this resource "+
				"in the API, refer to the documentation for how to do it manually.",
		)
	}
}
