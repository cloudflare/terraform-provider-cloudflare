// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moq_relay

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/moq"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*MoqRelayResource)(nil)
var _ resource.ResourceWithModifyPlan = (*MoqRelayResource)(nil)
var _ resource.ResourceWithImportState = (*MoqRelayResource)(nil)

func NewResource() resource.Resource {
	return &MoqRelayResource{}
}

// MoqRelayResource defines the resource implementation.
type MoqRelayResource struct {
	client *cloudflare.Client
}

func (r *MoqRelayResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_moq_relay"
}

func (r *MoqRelayResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *MoqRelayResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *MoqRelayModel

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
	env := MoqRelayResultEnvelope{*data}
	_, err = r.client.Moq.Relays.New(
		ctx,
		moq.RelayNewParams{
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
	data.ID = data.UID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MoqRelayResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *MoqRelayModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *MoqRelayModel

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
	env := MoqRelayResultEnvelope{*data}
	_, err = r.client.Moq.Relays.Update(
		ctx,
		data.UID.ValueString(),
		moq.RelayUpdateParams{
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
	data.ID = data.UID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MoqRelayResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *MoqRelayModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := MoqRelayResultEnvelope{*data}
	_, err := r.client.Moq.Relays.Get(
		ctx,
		data.UID.ValueString(),
		moq.RelayGetParams{
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
	data.ID = data.UID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MoqRelayResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *MoqRelayModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.Moq.Relays.Delete(
		ctx,
		data.UID.ValueString(),
		moq.RelayDeleteParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	data.ID = data.UID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MoqRelayResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data = new(MoqRelayModel)

	path_account_id := ""
	path_relay_id := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<account_id>/<relay_id>",
		&path_account_id,
		&path_relay_id,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.AccountID = types.StringValue(path_account_id)
	data.UID = types.StringValue(path_relay_id)

	res := new(http.Response)
	env := MoqRelayResultEnvelope{*data}
	_, err := r.client.Moq.Relays.Get(
		ctx,
		path_relay_id,
		moq.RelayGetParams{
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
	data.ID = data.UID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MoqRelayResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
