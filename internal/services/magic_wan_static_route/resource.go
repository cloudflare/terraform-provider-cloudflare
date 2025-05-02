// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_wan_static_route

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/magic_transit"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*MagicWANStaticRouteResource)(nil)
var _ resource.ResourceWithModifyPlan = (*MagicWANStaticRouteResource)(nil)
var _ resource.ResourceWithImportState = (*MagicWANStaticRouteResource)(nil)

func NewResource() resource.Resource {
	return &MagicWANStaticRouteResource{}
}

// MagicWANStaticRouteResource defines the resource implementation.
type MagicWANStaticRouteResource struct {
	client *cloudflare.Client
}

func (r *MagicWANStaticRouteResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_magic_wan_static_route"
}

func (r *MagicWANStaticRouteResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *MagicWANStaticRouteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *MagicWANStaticRouteModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := data.MarshalJSON()

	// Workaround to wrap the route into an array, which is what the API expects.
	dataBytes = append([]byte("["), dataBytes...)
	dataBytes = append(dataBytes, byte(']'))

	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := MagicWANStaticRouteResultEnvelope{*data}
	_, err = r.client.MagicTransit.Routes.New(
		ctx,
		magic_transit.RouteNewParams{
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

func (r *MagicWANStaticRouteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *MagicWANStaticRouteModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *MagicWANStaticRouteModel

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
	env := MagicWANStaticRouteResultEnvelope{*data}
	_, err = r.client.MagicTransit.Routes.Update(
		ctx,
		data.ID.ValueString(),
		magic_transit.RouteUpdateParams{
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

func (r *MagicWANStaticRouteResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *MagicWANStaticRouteModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := MagicWANStaticRouteResultEnvelope{*data}
	_, err := r.client.MagicTransit.Routes.Get(
		ctx,
		data.ID.ValueString(),
		magic_transit.RouteGetParams{
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

func (r *MagicWANStaticRouteResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *MagicWANStaticRouteModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.MagicTransit.Routes.Delete(
		ctx,
		data.ID.ValueString(),
		magic_transit.RouteDeleteParams{
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

func (r *MagicWANStaticRouteResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data *MagicWANStaticRouteModel = new(MagicWANStaticRouteModel)

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
	env := MagicWANStaticRouteResultEnvelope{*data}
	_, err := r.client.MagicTransit.Routes.Get(
		ctx,
		path_route_id,
		magic_transit.RouteGetParams{
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

func (r *MagicWANStaticRouteResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
