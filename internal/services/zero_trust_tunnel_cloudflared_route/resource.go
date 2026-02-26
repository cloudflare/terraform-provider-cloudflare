// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared_route

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*ZeroTrustTunnelCloudflaredRouteResource)(nil)
var _ resource.ResourceWithModifyPlan = (*ZeroTrustTunnelCloudflaredRouteResource)(nil)
var _ resource.ResourceWithImportState = (*ZeroTrustTunnelCloudflaredRouteResource)(nil)

func NewResource() resource.Resource {
	return &ZeroTrustTunnelCloudflaredRouteResource{}
}

// ZeroTrustTunnelCloudflaredRouteResource defines the resource implementation.
type ZeroTrustTunnelCloudflaredRouteResource struct {
	client *cloudflare.Client
}

func (r *ZeroTrustTunnelCloudflaredRouteResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zero_trust_tunnel_cloudflared_route"
}

func (r *ZeroTrustTunnelCloudflaredRouteResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ZeroTrustTunnelCloudflaredRouteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ZeroTrustTunnelCloudflaredRouteModel

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
	env := ZeroTrustTunnelCloudflaredRouteResultEnvelope{*data}
	_, err = r.client.ZeroTrust.Networks.Routes.New(
		ctx,
		zero_trust.NetworkRouteNewParams{
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

func (r *ZeroTrustTunnelCloudflaredRouteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ZeroTrustTunnelCloudflaredRouteModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *ZeroTrustTunnelCloudflaredRouteModel

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
	env := ZeroTrustTunnelCloudflaredRouteResultEnvelope{*data}
	_, err = r.client.ZeroTrust.Networks.Routes.Edit(
		ctx,
		data.ID.ValueString(),
		zero_trust.NetworkRouteEditParams{
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

func (r *ZeroTrustTunnelCloudflaredRouteResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ZeroTrustTunnelCloudflaredRouteModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// The v4 provider stored the route ID as a network CIDR (e.g. "10.0.0.0/16")
	// or, when a virtual network was set, as an MD5 checksum of "network/vnet_id"
	// (32 lowercase hex chars). The v5 provider requires a UUID (36 chars with
	// hyphens). When we detect a legacy ID format, we look up the route via the
	// List API (filtering by network + optional virtual_network_id) to obtain the
	// real UUID, then update the state ID before proceeding with the normal Read.
	if idStr := data.ID.ValueString(); isLegacyRouteID(idStr) {
		listParams := zero_trust.NetworkRouteListParams{
			AccountID:       cloudflare.F(data.AccountID.ValueString()),
			NetworkSubset:   cloudflare.F(data.Network.ValueString()),
			NetworkSuperset: cloudflare.F(data.Network.ValueString()),
			IsDeleted:       cloudflare.F(false),
		}
		if vnetID := data.VirtualNetworkID.ValueString(); vnetID != "" {
			listParams.VirtualNetworkID = cloudflare.F(vnetID)
		}
		routes, err := r.client.ZeroTrust.Networks.Routes.List(
			ctx,
			listParams,
			option.WithMiddleware(logging.Middleware(ctx)),
		)
		if err != nil {
			resp.Diagnostics.AddError("failed to look up route by network during v4 ID migration", err.Error())
			return
		}
		var found *zero_trust.Teamnet
		for _, route := range routes.Result {
			route := route
			if route.Network == data.Network.ValueString() {
				found = &route
				break
			}
		}
		if found == nil {
			resp.Diagnostics.AddWarning("Resource not found", "Could not locate the tunnel route by network during v4 ID migration; removing from state.")
			resp.State.RemoveResource(ctx)
			return
		}
		data.ID = types.StringValue(found.ID)
	}

	res := new(http.Response)
	env := ZeroTrustTunnelCloudflaredRouteResultEnvelope{*data}
	_, err := r.client.ZeroTrust.Networks.Routes.Get(
		ctx,
		data.ID.ValueString(),
		zero_trust.NetworkRouteGetParams{
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

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustTunnelCloudflaredRouteResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ZeroTrustTunnelCloudflaredRouteModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.ZeroTrust.Networks.Routes.Delete(
		ctx,
		data.ID.ValueString(),
		zero_trust.NetworkRouteDeleteParams{
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

func (r *ZeroTrustTunnelCloudflaredRouteResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data = new(ZeroTrustTunnelCloudflaredRouteModel)

	path_account_id := ""
	path_route_id := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<account_id>/<route_id>",
		&path_account_id,
		&path_route_id,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.AccountID = types.StringValue(path_account_id)
	data.ID = types.StringValue(path_route_id)

	res := new(http.Response)
	env := ZeroTrustTunnelCloudflaredRouteResultEnvelope{*data}
	_, err := r.client.ZeroTrust.Networks.Routes.Get(
		ctx,
		path_route_id,
		zero_trust.NetworkRouteGetParams{
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

func (r *ZeroTrustTunnelCloudflaredRouteResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}

// BEGIN CUSTOM: v4 ID migration helpers

// isLegacyRouteID reports whether id is a v4-format route ID (network CIDR or
// MD5 checksum) rather than a v5 UUID. A UUID is 36 characters in the form
// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx. A CIDR contains "/". A checksum is 32
// lowercase hex characters with no hyphens.
func isLegacyRouteID(id string) bool {
	if id == "" {
		return false
	}
	// CIDR: contains a slash
	if strings.Contains(id, "/") {
		return true
	}
	// MD5 checksum: exactly 32 hex chars, no hyphens
	if len(id) == 32 && !strings.Contains(id, "-") {
		for _, c := range id {
			if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f')) {
				return false
			}
		}
		return true
	}
	return false
}

// END CUSTOM: v4 ID migration helpers
