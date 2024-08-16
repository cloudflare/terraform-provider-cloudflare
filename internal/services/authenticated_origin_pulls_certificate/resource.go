// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package authenticated_origin_pulls_certificate

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/origin_tls_client_auth"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = &AuthenticatedOriginPullsCertificateResource{}
var _ resource.ResourceWithModifyPlan = &AuthenticatedOriginPullsCertificateResource{}
var _ resource.ResourceWithImportState = &AuthenticatedOriginPullsCertificateResource{}

func NewResource() resource.Resource {
	return &AuthenticatedOriginPullsCertificateResource{}
}

// AuthenticatedOriginPullsCertificateResource defines the resource implementation.
type AuthenticatedOriginPullsCertificateResource struct {
	client *cloudflare.Client
}

func (r *AuthenticatedOriginPullsCertificateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticated_origin_pulls_certificate"
}

func (r *AuthenticatedOriginPullsCertificateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *AuthenticatedOriginPullsCertificateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *AuthenticatedOriginPullsCertificateModel

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
	env := AuthenticatedOriginPullsCertificateResultEnvelope{*data}
	_, err = r.client.OriginTLSClientAuth.New(
		ctx,
		origin_tls_client_auth.OriginTLSClientAuthNewParams{
			ZoneID: cloudflare.F(data.ZoneID.ValueString()),
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

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticatedOriginPullsCertificateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *AuthenticatedOriginPullsCertificateModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *AuthenticatedOriginPullsCertificateModel

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
	env := AuthenticatedOriginPullsCertificateResultEnvelope{*data}
	_, err = r.client.OriginTLSClientAuth.New(
		ctx,
		origin_tls_client_auth.OriginTLSClientAuthNewParams{
			ZoneID: cloudflare.F(data.ID.ValueString()),
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

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticatedOriginPullsCertificateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *AuthenticatedOriginPullsCertificateModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := AuthenticatedOriginPullsCertificateResultEnvelope{*data}
	_, err := r.client.OriginTLSClientAuth.Get(
		ctx,
		data.ID.ValueString(),
		origin_tls_client_auth.OriginTLSClientAuthGetParams{
			ZoneID: cloudflare.F(data.ZoneID.ValueString()),
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

func (r *AuthenticatedOriginPullsCertificateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *AuthenticatedOriginPullsCertificateModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.OriginTLSClientAuth.Delete(
		ctx,
		data.ID.ValueString(),
		origin_tls_client_auth.OriginTLSClientAuthDeleteParams{
			ZoneID: cloudflare.F(data.ZoneID.ValueString()),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticatedOriginPullsCertificateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data *AuthenticatedOriginPullsCertificateModel

	path := strings.Split(req.ID, "/")
	if len(path) != 2 {
		resp.Diagnostics.AddError("Invalid ID", "expected urlencoded segments <zone_id>/<certificate_id>")
		return
	}
	path_zone_id, err := url.PathUnescape(path[0])
	if err != nil {
		resp.Diagnostics.AddError("invalid urlencoded segment - <zone_id>", fmt.Sprintf("%s -> %q", err.Error(), path[0]))
	}
	path_certificate_id, err := url.PathUnescape(path[1])
	if err != nil {
		resp.Diagnostics.AddError("invalid urlencoded segment - <certificate_id>", fmt.Sprintf("%s -> %q", err.Error(), path[1]))
	}
	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := AuthenticatedOriginPullsCertificateResultEnvelope{*data}
	_, err = r.client.OriginTLSClientAuth.Get(
		ctx,
		path_certificate_id,
		origin_tls_client_auth.OriginTLSClientAuthGetParams{
			ZoneID: cloudflare.F(path_zone_id),
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

func (r *AuthenticatedOriginPullsCertificateResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
