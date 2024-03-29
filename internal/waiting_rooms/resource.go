// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_rooms

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/waiting_rooms"
	"github.com/cloudflare/cloudflare-terraform/internal/apijson"
	"github.com/cloudflare/cloudflare-terraform/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &WaitingRoomsResource{}

func NewResource() resource.Resource {
	return &WaitingRoomsResource{}
}

// WaitingRoomsResource defines the resource implementation.
type WaitingRoomsResource struct {
	client *cloudflare.Client
}

func (r *WaitingRoomsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_waiting_rooms"
}

func (r *WaitingRoomsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *WaitingRoomsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *WaitingRoomsModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := apijson.Marshal(data)
	if err != nil {
		resp.Diagnostics.AddError("failed to create resource", err.Error())
		return
	}
	res := new(http.Response)
	env := WaitingRoomsResultEnvelope{*data}
	_, err = r.client.WaitingRooms.New(
		ctx,
		data.ZoneIdentifier.ValueString(),
		waiting_rooms.WaitingRoomNewParams{},
		option.WithRequestBody("application/json", dataBytes),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	bytes, _ := io.ReadAll(res.Body)
	apijson.Unmarshal(bytes, &env)
	data = &env.Result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WaitingRoomsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *WaitingRoomsModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	env := WaitingRoomsResultEnvelope{*data}
	_, err := r.client.WaitingRooms.Get(
		ctx,
		data.ZoneIdentifier.ValueString(),
		data.WaitingRoomID.ValueString(),
		option.WithResponseBodyInto(&env),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	data = &env.Result

	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WaitingRoomsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *WaitingRoomsModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := apijson.Marshal(data)
	if err != nil {
		resp.Diagnostics.AddError("failed to create resource", err.Error())
		return
	}
	res := new(http.Response)
	env := WaitingRoomsResultEnvelope{*data}
	_, err = r.client.WaitingRooms.Update(
		ctx,
		data.ZoneIdentifier.ValueString(),
		data.WaitingRoomID.ValueString(),
		waiting_rooms.WaitingRoomUpdateParams{},
		option.WithRequestBody("application/json", dataBytes),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	bytes, _ := io.ReadAll(res.Body)
	apijson.Unmarshal(bytes, &env)
	data = &env.Result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WaitingRoomsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *WaitingRoomsModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.WaitingRooms.Delete(
		ctx,
		data.ZoneIdentifier.ValueString(),
		data.WaitingRoomID.ValueString(),
		option.WithMiddleware(logging.Middleware(ctx)),
	)

	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
