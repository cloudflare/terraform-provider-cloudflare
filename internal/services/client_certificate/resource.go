// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package client_certificate

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/client_certificates"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*ClientCertificateResource)(nil)
var _ resource.ResourceWithModifyPlan = (*ClientCertificateResource)(nil)
var _ resource.ResourceWithImportState = (*ClientCertificateResource)(nil)

func NewResource() resource.Resource {
	return &ClientCertificateResource{}
}

// ClientCertificateResource defines the resource implementation.
type ClientCertificateResource struct {
	client *cloudflare.Client
}

func (r *ClientCertificateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_client_certificate"
}

func (r *ClientCertificateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ClientCertificateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ClientCertificateModel

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
	env := ClientCertificateResultEnvelope{*data}
	_, err = r.client.ClientCertificates.New(
		ctx,
		client_certificates.ClientCertificateNewParams{
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

	// Normalize CSR to prevent drift from \r\n vs \n differences
	var configData ClientCertificateModel
	req.Plan.Get(ctx, &configData)

	apiNormalized := strings.ReplaceAll(data.Csr.ValueString(), "\r\n", "\n")
	apiNormalized = strings.TrimRight(apiNormalized, "\n")
	configNormalized := strings.ReplaceAll(configData.Csr.ValueString(), "\r\n", "\n")
	configNormalized = strings.TrimRight(configNormalized, "\n")

	if apiNormalized == configNormalized {
		data.Csr = configData.Csr
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClientCertificateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ClientCertificateModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *ClientCertificateModel

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
	env := ClientCertificateResultEnvelope{*data}
	_, err = r.client.ClientCertificates.Edit(
		ctx,
		data.ID.ValueString(),
		client_certificates.ClientCertificateEditParams{
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

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClientCertificateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ClientCertificateModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := ClientCertificateResultEnvelope{*data}
	_, err := r.client.ClientCertificates.Get(
		ctx,
		data.ID.ValueString(),
		client_certificates.ClientCertificateGetParams{
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

	// Normalize CSR to prevent drift from \r\n vs \n differences
	var stateData ClientCertificateModel
	req.State.Get(ctx, &stateData)

	apiNormalized := strings.ReplaceAll(data.Csr.ValueString(), "\r\n", "\n")
	apiNormalized = strings.TrimRight(apiNormalized, "\n")
	stateNormalized := strings.ReplaceAll(stateData.Csr.ValueString(), "\r\n", "\n")
	stateNormalized = strings.TrimRight(stateNormalized, "\n")

	if apiNormalized == stateNormalized {
		data.Csr = stateData.Csr
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClientCertificateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ClientCertificateModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.ClientCertificates.Delete(
		ctx,
		data.ID.ValueString(),
		client_certificates.ClientCertificateDeleteParams{
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

func (r *ClientCertificateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data = new(ClientCertificateModel)

	path_zone_id := ""
	path_client_certificate_id := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<zone_id>/<client_certificate_id>",
		&path_zone_id,
		&path_client_certificate_id,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.ZoneID = types.StringValue(path_zone_id)
	data.ID = types.StringValue(path_client_certificate_id)

	res := new(http.Response)
	env := ClientCertificateResultEnvelope{*data}
	_, err := r.client.ClientCertificates.Get(
		ctx,
		path_client_certificate_id,
		client_certificates.ClientCertificateGetParams{
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

func (r *ClientCertificateResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
