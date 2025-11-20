// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_mtls_hostname_settings

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*ZeroTrustAccessMTLSHostnameSettingsResource)(nil)
var _ resource.ResourceWithModifyPlan = (*ZeroTrustAccessMTLSHostnameSettingsResource)(nil)
var _ resource.ResourceWithImportState = (*ZeroTrustAccessMTLSHostnameSettingsResource)(nil)

func NewResource() resource.Resource {
	return &ZeroTrustAccessMTLSHostnameSettingsResource{}
}

// ZeroTrustAccessMTLSHostnameSettingsResource defines the resource implementation.
type ZeroTrustAccessMTLSHostnameSettingsResource struct {
	client *cloudflare.Client
}

func (r *ZeroTrustAccessMTLSHostnameSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zero_trust_access_mtls_hostname_settings"
}

func (r *ZeroTrustAccessMTLSHostnameSettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ZeroTrustAccessMTLSHostnameSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ZeroTrustAccessMTLSHostnameSettingsModel

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
	env := ZeroTrustAccessMTLSHostnameSettingsResultEnvelope{*data}
	params := zero_trust.AccessCertificateSettingUpdateParams{}

	if !data.AccountID.IsNull() {
		params.AccountID = cloudflare.F(data.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
	}

	err = retry.RetryContext(ctx, 2*time.Minute, func() *retry.RetryError {
		_, retryErr := r.client.ZeroTrust.Access.Certificates.Settings.Update(
			ctx,
			params,
			option.WithRequestBody("application/json", dataBytes),
			option.WithResponseBodyInto(&res),
			option.WithMiddleware(logging.Middleware(ctx)),
		)
		if retryErr != nil {
			// Check if it's a conflict error that should be retried
			if res != nil && res.StatusCode == 409 {
				return retry.RetryableError(retryErr)
			}
			return retry.NonRetryableError(retryErr)
		}
		return nil
	})
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

func (r *ZeroTrustAccessMTLSHostnameSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ZeroTrustAccessMTLSHostnameSettingsModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *ZeroTrustAccessMTLSHostnameSettingsModel

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
	env := ZeroTrustAccessMTLSHostnameSettingsResultEnvelope{*data}
	params := zero_trust.AccessCertificateSettingUpdateParams{}

	if !data.AccountID.IsNull() {
		params.AccountID = cloudflare.F(data.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
	}

	err = retry.RetryContext(ctx, 2*time.Minute, func() *retry.RetryError {
		_, retryErr := r.client.ZeroTrust.Access.Certificates.Settings.Update(
			ctx,
			params,
			option.WithRequestBody("application/json", dataBytes),
			option.WithResponseBodyInto(&res),
			option.WithMiddleware(logging.Middleware(ctx)),
		)
		if retryErr != nil {
			// Check if it's a conflict error that should be retried
			if res != nil && res.StatusCode == 409 {
				return retry.RetryableError(retryErr)
			}
			return retry.NonRetryableError(retryErr)
		}
		return nil
	})
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

func (r *ZeroTrustAccessMTLSHostnameSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ZeroTrustAccessMTLSHostnameSettingsModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := ZeroTrustAccessMTLSHostnameSettingsResultEnvelope{*data}
	params := zero_trust.AccessCertificateSettingGetParams{}

	if !data.AccountID.IsNull() {
		params.AccountID = cloudflare.F(data.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
	}

	_, err := r.client.ZeroTrust.Access.Certificates.Settings.Get(
		ctx,
		params,
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

func (r *ZeroTrustAccessMTLSHostnameSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

}

func (r *ZeroTrustAccessMTLSHostnameSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data *ZeroTrustAccessMTLSHostnameSettingsModel = new(ZeroTrustAccessMTLSHostnameSettingsModel)
	importID := req.ID
	
	// Check if it's a zone ID or account ID and set the appropriate field
	if len(importID) == 32 {
		// Account ID
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("account_id"), importID)...)
		data.AccountID = types.StringValue(importID)
	} else {
		// Zone ID
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("zone_id"), importID)...)
		data.ZoneID = types.StringValue(importID)
	}
	
	if resp.Diagnostics.HasError() {
		return
	}

	// Since import just passes the identifier, we need to read the current state
	// from the API. We'll do this by creating a mock ReadRequest and calling Read()
	var fullData *ZeroTrustAccessMTLSHostnameSettingsModel = new(ZeroTrustAccessMTLSHostnameSettingsModel)
	if !data.AccountID.IsNull() {
		fullData.AccountID = data.AccountID
	} else {
		fullData.ZoneID = data.ZoneID
	}
	
	// Set the full data in state first so Read can process it
	resp.Diagnostics.Append(resp.State.Set(ctx, fullData)...)
	if resp.Diagnostics.HasError() {
		return
	}
	
	// Create a Read request and response to fetch the full state
	readReq := resource.ReadRequest{State: resp.State}
	readResp := &resource.ReadResponse{State: resp.State, Diagnostics: resp.Diagnostics}
	
	// Call the Read method to populate the full state
	r.Read(ctx, readReq, readResp)
	
	// Copy back the results
	resp.State = readResp.State
	resp.Diagnostics = readResp.Diagnostics
}

func (r *ZeroTrustAccessMTLSHostnameSettingsResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
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
