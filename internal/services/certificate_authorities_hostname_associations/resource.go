// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package certificate_authorities_hostname_associations

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/certificate_authorities"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*CertificateAuthoritiesHostnameAssociationsResource)(nil)
var _ resource.ResourceWithModifyPlan = (*CertificateAuthoritiesHostnameAssociationsResource)(nil)
var _ resource.ResourceWithImportState = (*CertificateAuthoritiesHostnameAssociationsResource)(nil)

func NewResource() resource.Resource {
	return &CertificateAuthoritiesHostnameAssociationsResource{}
}

// CertificateAuthoritiesHostnameAssociationsResource defines the resource implementation.
type CertificateAuthoritiesHostnameAssociationsResource struct {
	client *cloudflare.Client
}

func (r *CertificateAuthoritiesHostnameAssociationsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_certificate_authorities_hostname_associations"
}

func (r *CertificateAuthoritiesHostnameAssociationsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *CertificateAuthoritiesHostnameAssociationsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *CertificateAuthoritiesHostnameAssociationsModel

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
	env := CertificateAuthoritiesHostnameAssociationsResultEnvelope{*data}
	_, err = r.client.CertificateAuthorities.HostnameAssociations.Update(
		ctx,
		certificate_authorities.HostnameAssociationUpdateParams{
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
	data.ID = data.ZoneID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CertificateAuthoritiesHostnameAssociationsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *CertificateAuthoritiesHostnameAssociationsModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *CertificateAuthoritiesHostnameAssociationsModel

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
	env := CertificateAuthoritiesHostnameAssociationsResultEnvelope{*data}
	_, err = r.client.CertificateAuthorities.HostnameAssociations.Update(
		ctx,
		certificate_authorities.HostnameAssociationUpdateParams{
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
	data.ID = data.ZoneID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CertificateAuthoritiesHostnameAssociationsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *CertificateAuthoritiesHostnameAssociationsModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := CertificateAuthoritiesHostnameAssociationsResultEnvelope{*data}
	getParams := certificate_authorities.HostnameAssociationGetParams{
		ZoneID: cloudflare.F(data.ZoneID.ValueString()),
	}
	if !data.MTLSCertificateID.IsNull() && !data.MTLSCertificateID.IsUnknown() {
		getParams.MTLSCertificateID = cloudflare.F(data.MTLSCertificateID.ValueString())
	}
	_, err := r.client.CertificateAuthorities.HostnameAssociations.Get(
		ctx,
		getParams,
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
	data.ID = data.ZoneID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CertificateAuthoritiesHostnameAssociationsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

}

func (r *CertificateAuthoritiesHostnameAssociationsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data = new(CertificateAuthoritiesHostnameAssociationsModel)

	// Support import formats:
	//   <zone_id>                        - for managed CA associations
	//   <zone_id>/<mtls_certificate_id>  - for mTLS certificate associations
	parts := strings.Split(req.ID, "/")
	if len(parts) < 1 || len(parts) > 2 {
		resp.Diagnostics.AddError(
			"invalid import ID",
			fmt.Sprintf("expected format \"<zone_id>\" or \"<zone_id>/<mtls_certificate_id>\", got %q", req.ID),
		)
		return
	}

	data.ZoneID = types.StringValue(parts[0])
	getParams := certificate_authorities.HostnameAssociationGetParams{
		ZoneID: cloudflare.F(parts[0]),
	}
	if len(parts) == 2 && parts[1] != "" {
		data.MTLSCertificateID = types.StringValue(parts[1])
		getParams.MTLSCertificateID = cloudflare.F(parts[1])
	}

	res := new(http.Response)
	env := CertificateAuthoritiesHostnameAssociationsResultEnvelope{*data}
	_, err := r.client.CertificateAuthorities.HostnameAssociations.Get(
		ctx,
		getParams,
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
	data.ID = data.ZoneID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CertificateAuthoritiesHostnameAssociationsResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
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
