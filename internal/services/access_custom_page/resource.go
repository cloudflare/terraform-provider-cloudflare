// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_custom_page

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = &AccessCustomPageResource{}

func NewResource() resource.Resource {
	return &AccessCustomPageResource{}
}

// AccessCustomPageResource defines the resource implementation.
type AccessCustomPageResource struct {
	client *cloudflare.Client
}

func (r *AccessCustomPageResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_access_custom_page"
}

func (r *AccessCustomPageResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *AccessCustomPageResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *AccessCustomPageModel

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
	env := AccessCustomPageResultEnvelope{*data}
	_, err = r.client.ZeroTrust.Access.CustomPages.New(
		ctx,
		zero_trust.AccessCustomPageNewParams{
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
	data.ID = data.UID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AccessCustomPageResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *AccessCustomPageModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := AccessCustomPageResultEnvelope{*data}
	_, err := r.client.ZeroTrust.Access.CustomPages.Get(
		ctx,
		data.UID.ValueString(),
		zero_trust.AccessCustomPageGetParams{
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
	data.ID = data.UID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AccessCustomPageResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *AccessCustomPageModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *AccessCustomPageModel

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
	env := AccessCustomPageResultEnvelope{*data}
	_, err = r.client.ZeroTrust.Access.CustomPages.Update(
		ctx,
		data.UID.ValueString(),
		zero_trust.AccessCustomPageUpdateParams{
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
	data.ID = data.UID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AccessCustomPageResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *AccessCustomPageModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.ZeroTrust.Access.CustomPages.Delete(
		ctx,
		data.UID.ValueString(),
		zero_trust.AccessCustomPageDeleteParams{
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
