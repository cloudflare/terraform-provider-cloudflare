// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package authenticated_origin_pulls_certificate

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/origin_tls_client_auth"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*AuthenticatedOriginPullsCertificateResource)(nil)
var _ resource.ResourceWithModifyPlan = (*AuthenticatedOriginPullsCertificateResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticatedOriginPullsCertificateResource)(nil)

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

	privateKey := data.PrivateKey
	dataBytes, err := data.MarshalJSON()
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := AuthenticatedOriginPullsCertificateResultEnvelope{*data}
	_, err = r.client.OriginTLSClientAuth.ZoneCertificates.New(
		ctx,
		origin_tls_client_auth.ZoneCertificateNewParams{
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
	err = apijson.UnmarshalComputed(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result
	data.CertificateID = data.ID
	data.PrivateKey = privateKey

	// Get the original config value to preserve its format
	var configData AuthenticatedOriginPullsCertificateModel
	req.Config.Get(ctx, &configData)

	apiNormalized := strings.TrimRight(data.Certificate.ValueString(), "\n")
	configNormalized := strings.TrimRight(configData.Certificate.ValueString(), "\n")

	// If they match, use the config format
	if apiNormalized == configNormalized {
		data.Certificate = configData.Certificate
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticatedOriginPullsCertificateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Update is not supported for this resource
}

func (r *AuthenticatedOriginPullsCertificateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *AuthenticatedOriginPullsCertificateModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	privateKey := data.PrivateKey
	res := new(http.Response)
	env := AuthenticatedOriginPullsCertificateResultEnvelope{*data}
	_, err := r.client.OriginTLSClientAuth.ZoneCertificates.Get(
		ctx,
		data.ID.ValueString(),
		origin_tls_client_auth.ZoneCertificateGetParams{
			ZoneID: cloudflare.F(data.ZoneID.ValueString()),
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
	data.CertificateID = data.ID
	data.PrivateKey = privateKey

	// Keep the original state format if they're semantically equal
	var stateData AuthenticatedOriginPullsCertificateModel
	req.State.Get(ctx, &stateData)

	apiNormalized := strings.TrimRight(data.Certificate.ValueString(), "\n")
	stateNormalized := strings.TrimRight(stateData.Certificate.ValueString(), "\n")

	if apiNormalized == stateNormalized {
		data.Certificate = stateData.Certificate
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticatedOriginPullsCertificateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *AuthenticatedOriginPullsCertificateModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.OriginTLSClientAuth.ZoneCertificates.Delete(
		ctx,
		data.ID.ValueString(),
		origin_tls_client_auth.ZoneCertificateDeleteParams{
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
	var data = new(AuthenticatedOriginPullsCertificateModel)

	path_zone_id := ""
	path_certificate_id := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<zone_id>/<certificate_id>",
		&path_zone_id,
		&path_certificate_id,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.ZoneID = types.StringValue(path_zone_id)
	data.ID = types.StringValue(path_certificate_id)

	res := new(http.Response)
	env := AuthenticatedOriginPullsCertificateResultEnvelope{*data}
	_, err := r.client.OriginTLSClientAuth.ZoneCertificates.Get(
		ctx,
		path_certificate_id,
		origin_tls_client_auth.ZoneCertificateGetParams{
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
