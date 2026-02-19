// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package authenticated_origin_pulls

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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
var _ resource.ResourceWithConfigure = (*AuthenticatedOriginPullsResource)(nil)
var _ resource.ResourceWithModifyPlan = (*AuthenticatedOriginPullsResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticatedOriginPullsResource)(nil)

func NewResource() resource.Resource {
	return &AuthenticatedOriginPullsResource{}
}

// AuthenticatedOriginPullsResource defines the resource implementation.
type AuthenticatedOriginPullsResource struct {
	client *cloudflare.Client
}

func (r *AuthenticatedOriginPullsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticated_origin_pulls"
}

func (r *AuthenticatedOriginPullsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *AuthenticatedOriginPullsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *AuthenticatedOriginPullsModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Validate config has exactly one item - this resource manages a single hostname association
	if data.Config == nil || len(*data.Config) != 1 {
		resp.Diagnostics.AddError(
			"invalid config",
			"config must contain exactly one hostname association. Create separate resources to manage multiple hostnames.",
		)
		return
	}
	targetHostname := (*data.Config)[0].Hostname.ValueString()
	if targetHostname == "" {
		resp.Diagnostics.AddError("missing hostname", "config[0].hostname must not be empty")
		return
	}

	// Preserve the original config since it has no_refresh tag
	originalConfig := data.Config

	dataBytes, err := data.MarshalJSON()
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	// Use array envelope since the API returns an array of hostname associations
	env := AuthenticatedOriginPullsArrayResultEnvelope{}
	_, err = r.client.OriginTLSClientAuth.Hostnames.Update(
		ctx,
		origin_tls_client_auth.HostnameUpdateParams{
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

	// Find the matching hostname from the array response
	result, err := env.FindByHostname(targetHostname)
	if err != nil {
		resp.Diagnostics.AddError("hostname not found in response", err.Error())
		return
	}

	// Copy computed fields from the API response
	data.Hostname = result.Hostname
	data.ID = result.Hostname
	data.CERTID = result.CERTID
	data.CERTStatus = result.CERTStatus
	data.CERTUpdatedAt = result.CERTUpdatedAt
	data.CERTUploadedOn = result.CERTUploadedOn
	data.Certificate = result.Certificate
	data.CreatedAt = result.CreatedAt
	data.Enabled = result.Enabled
	data.ExpiresOn = result.ExpiresOn
	data.Issuer = result.Issuer
	data.SerialNumber = result.SerialNumber
	data.Signature = result.Signature
	data.Status = result.Status
	data.UpdatedAt = result.UpdatedAt
	// PrivateKey has no_refresh tag and is not returned by API, set to null
	data.PrivateKey = types.StringNull()

	// Restore the original config (has no_refresh tag)
	data.Config = originalConfig

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticatedOriginPullsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *AuthenticatedOriginPullsModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *AuthenticatedOriginPullsModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Validate config has exactly one item - this resource manages a single hostname association
	if data.Config == nil || len(*data.Config) != 1 {
		resp.Diagnostics.AddError(
			"invalid config",
			"config must contain exactly one hostname association. Create separate resources to manage multiple hostnames.",
		)
		return
	}
	targetHostname := (*data.Config)[0].Hostname.ValueString()
	if targetHostname == "" {
		resp.Diagnostics.AddError("missing hostname", "config[0].hostname must not be empty")
		return
	}

	// Preserve the original config since it has no_refresh tag
	originalConfig := data.Config

	dataBytes, err := data.MarshalJSONForUpdate(*state)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	// Use array envelope since the API returns an array of hostname associations
	env := AuthenticatedOriginPullsArrayResultEnvelope{}
	_, err = r.client.OriginTLSClientAuth.Hostnames.Update(
		ctx,
		origin_tls_client_auth.HostnameUpdateParams{
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

	// Find the matching hostname from the array response
	result, err := env.FindByHostname(targetHostname)
	if err != nil {
		resp.Diagnostics.AddError("hostname not found in response", err.Error())
		return
	}

	// Copy computed fields from the API response
	data.Hostname = result.Hostname
	data.ID = result.Hostname
	data.CERTID = result.CERTID
	data.CERTStatus = result.CERTStatus
	data.CERTUpdatedAt = result.CERTUpdatedAt
	data.CERTUploadedOn = result.CERTUploadedOn
	data.Certificate = result.Certificate
	data.CreatedAt = result.CreatedAt
	data.Enabled = result.Enabled
	data.ExpiresOn = result.ExpiresOn
	data.Issuer = result.Issuer
	data.SerialNumber = result.SerialNumber
	data.Signature = result.Signature
	data.Status = result.Status
	data.UpdatedAt = result.UpdatedAt
	// PrivateKey has no_refresh tag and is not returned by API, set to null
	data.PrivateKey = types.StringNull()

	// Restore the original config (has no_refresh tag)
	data.Config = originalConfig

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticatedOriginPullsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *AuthenticatedOriginPullsModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve the original config since it has no_refresh tag
	originalConfig := data.Config

	res := new(http.Response)
	env := AuthenticatedOriginPullsResultEnvelope{*data}
	_, err := r.client.OriginTLSClientAuth.Hostnames.Get(
		ctx,
		data.Hostname.ValueString(),
		origin_tls_client_auth.HostnameGetParams{
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
	data.ID = data.Hostname
	// PrivateKey has no_refresh tag and is not returned by API, set to null
	data.PrivateKey = types.StringNull()

	// Restore the original config (has no_refresh tag)
	data.Config = originalConfig

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticatedOriginPullsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *AuthenticatedOriginPullsModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// To "delete" a hostname association, we send enabled: null to void it.
	// This is the API-documented way to remove hostname AOP associations.
	// See: https://developers.cloudflare.com/ssl/origin-configuration/authenticated-origin-pull/set-up/rollback/
	hostname := data.Hostname.ValueString()
	if hostname == "" {
		// Nothing to delete
		return
	}

	deletePayload := map[string]interface{}{
		"config": []map[string]interface{}{
			{
				"hostname": hostname,
				"cert_id":  data.CERTID.ValueString(), // API requires cert_id
				"enabled":  nil,                       // null voids the association
			},
		},
	}

	dataBytes, err := json.Marshal(deletePayload)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize delete request", err.Error())
		return
	}

	_, err = r.client.OriginTLSClientAuth.Hostnames.Update(
		ctx,
		origin_tls_client_auth.HostnameUpdateParams{
			ZoneID: cloudflare.F(data.ZoneID.ValueString()),
		},
		option.WithRequestBody("application/json", dataBytes),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	// Resource successfully voided - Terraform will remove from state automatically
}

func (r *AuthenticatedOriginPullsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data = new(AuthenticatedOriginPullsModel)

	path_zone_id := ""
	path_hostname := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<zone_id>/<hostname>",
		&path_zone_id,
		&path_hostname,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.ZoneID = types.StringValue(path_zone_id)
	data.Hostname = types.StringValue(path_hostname)

	res := new(http.Response)
	env := AuthenticatedOriginPullsResultEnvelope{*data}
	_, err := r.client.OriginTLSClientAuth.Hostnames.Get(
		ctx,
		path_hostname,
		origin_tls_client_auth.HostnameGetParams{
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
	data.ID = data.Hostname

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticatedOriginPullsResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {
	// No warnings needed - Delete now properly voids the hostname association
}
