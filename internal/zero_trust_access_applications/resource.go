// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_applications

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/cloudflare/cloudflare-terraform/internal/apijson"
	"github.com/cloudflare/cloudflare-terraform/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ZeroTrustAccessApplicationsResource{}

func NewResource() resource.Resource {
	return &ZeroTrustAccessApplicationsResource{}
}

// ZeroTrustAccessApplicationsResource defines the resource implementation.
type ZeroTrustAccessApplicationsResource struct {
	client *cloudflare.Client
}

func (r *ZeroTrustAccessApplicationsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zero_trust_access_applications"
}

func (r *ZeroTrustAccessApplicationsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ZeroTrustAccessApplicationsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ZeroTrustAccessApplicationsModel

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
	env := ZeroTrustAccessApplicationsResultEnvelope{*data}
	_, err = r.client.ZeroTrust.Access.Applications.New(
		ctx,
		zero_trust.AccessApplicationNewParamsSelfHostedApplication{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
			ZoneID:    cloudflare.F(data.ZoneID.ValueString()),
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
	apijson.Unmarshal(bytes, &env)
	data = &env.Result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustAccessApplicationsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ZeroTrustAccessApplicationsModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	env := ZeroTrustAccessApplicationsResultEnvelope{*data}
	_, err := r.client.ZeroTrust.Access.Applications.Get(
		ctx,
		data.AppID.ValueString(),
		zero_trust.AccessApplicationGetParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
			ZoneID:    cloudflare.F(data.ZoneID.ValueString()),
		},
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

func (r *ZeroTrustAccessApplicationsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ZeroTrustAccessApplicationsModel

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
	env := ZeroTrustAccessApplicationsResultEnvelope{*data}
	_, err = r.client.ZeroTrust.Access.Applications.Update(
		ctx,
		data.AppID.ValueString(),
		zero_trust.AccessApplicationUpdateParamsSelfHostedApplication{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
			ZoneID:    cloudflare.F(data.ZoneID.ValueString()),
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
	apijson.Unmarshal(bytes, &env)
	data = &env.Result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustAccessApplicationsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ZeroTrustAccessApplicationsModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.ZeroTrust.Access.Applications.Delete(
		ctx,
		data.AppID.ValueString(),
		zero_trust.AccessApplicationDeleteParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
			ZoneID:    cloudflare.F(data.ZoneID.ValueString()),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)

	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
