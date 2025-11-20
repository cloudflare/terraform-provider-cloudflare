// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_certificate

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*ZeroTrustGatewayCertificateResource)(nil)
var _ resource.ResourceWithModifyPlan = (*ZeroTrustGatewayCertificateResource)(nil)
var _ resource.ResourceWithImportState = (*ZeroTrustGatewayCertificateResource)(nil)

func NewResource() resource.Resource {
	return &ZeroTrustGatewayCertificateResource{}
}

// ZeroTrustGatewayCertificateResource defines the resource implementation.
type ZeroTrustGatewayCertificateResource struct {
	client *cloudflare.Client
}

func (r *ZeroTrustGatewayCertificateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zero_trust_gateway_certificate"
}

func (r *ZeroTrustGatewayCertificateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ZeroTrustGatewayCertificateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ZeroTrustGatewayCertificateModel

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
	env := ZeroTrustGatewayCertificateResultEnvelope{*data}
	_, err = r.client.ZeroTrust.Gateway.Certificates.New(
		ctx,
		zero_trust.GatewayCertificateNewParams{
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

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustGatewayCertificateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan *ZeroTrustGatewayCertificateModel
	var state *ZeroTrustGatewayCertificateModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// First, get the current certificate resource to check its status
	res := new(http.Response)
	env := ZeroTrustGatewayCertificateResultEnvelope{*state}
	_, err := r.client.ZeroTrust.Gateway.Certificates.Get(
		ctx,
		state.ID.ValueString(),
		zero_trust.GatewayCertificateGetParams{
			AccountID: cloudflare.F(state.AccountID.ValueString()),
		},
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to get certificate", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.Unmarshal(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize certificate", err.Error())
		return
	}
	
	// Update state with current certificate data
	*state = env.Result
	
	// Check if activate field has changed and handle activation/deactivation
	if !plan.Activate.IsNull() {
		if plan.Activate.ValueBool() {
			// Activate the certificate
			_, err := r.client.ZeroTrust.Gateway.Certificates.Activate(
				ctx,
				state.ID.ValueString(),
				zero_trust.GatewayCertificateActivateParams{
					AccountID: cloudflare.F(state.AccountID.ValueString()),
					Body:      struct{}{}, // Empty body as required by the API
				},
				option.WithMiddleware(logging.Middleware(ctx)),
			)
			if err != nil {
				resp.Diagnostics.AddError("failed to activate certificate", err.Error())
				return
			}
		} else {
			// Deactivate the certificate
			_, err := r.client.ZeroTrust.Gateway.Certificates.Deactivate(
				ctx,
				state.ID.ValueString(),
				zero_trust.GatewayCertificateDeactivateParams{
					AccountID: cloudflare.F(state.AccountID.ValueString()),
					Body:      struct{}{}, // Empty body as required by the API
				},
				option.WithMiddleware(logging.Middleware(ctx)),
			)
			if err != nil {
				resp.Diagnostics.AddError("failed to deactivate certificate", err.Error())
				return
			}
		}
		
		// After activation/deactivation, get the updated certificate status
		res = new(http.Response)
		env = ZeroTrustGatewayCertificateResultEnvelope{*state}
		_, err = r.client.ZeroTrust.Gateway.Certificates.Get(
			ctx,
			state.ID.ValueString(),
			zero_trust.GatewayCertificateGetParams{
				AccountID: cloudflare.F(state.AccountID.ValueString()),
			},
			option.WithResponseBodyInto(&res),
			option.WithMiddleware(logging.Middleware(ctx)),
		)
		if err != nil {
			resp.Diagnostics.AddError("failed to get updated certificate", err.Error())
			return
		}
		bytes, _ = io.ReadAll(res.Body)
		err = apijson.Unmarshal(bytes, &env)
		if err != nil {
			resp.Diagnostics.AddError("failed to deserialize updated certificate", err.Error())
			return
		}
		*state = env.Result
	}
	
	// Set the activate field from the plan to maintain user's intended state
	state.Activate = plan.Activate

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ZeroTrustGatewayCertificateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ZeroTrustGatewayCertificateModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve the activate field from current state since it's Terraform-only
	currentActivate := data.Activate

	res := new(http.Response)
	env := ZeroTrustGatewayCertificateResultEnvelope{*data}
	_, err := r.client.ZeroTrust.Gateway.Certificates.Get(
		ctx,
		data.ID.ValueString(),
		zero_trust.GatewayCertificateGetParams{
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

	// Restore the activate field since it's not returned by the API
	data.Activate = currentActivate

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustGatewayCertificateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ZeroTrustGatewayCertificateModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.ZeroTrust.Gateway.Certificates.Delete(
		ctx,
		data.ID.ValueString(),
		zero_trust.GatewayCertificateDeleteParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustGatewayCertificateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data *ZeroTrustGatewayCertificateModel = new(ZeroTrustGatewayCertificateModel)

	path_account_id := ""
	path_certificate_id := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<account_id>/<certificate_id>",
		&path_account_id,
		&path_certificate_id,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.AccountID = types.StringValue(path_account_id)
	data.ID = types.StringValue(path_certificate_id)

	res := new(http.Response)
	env := ZeroTrustGatewayCertificateResultEnvelope{*data}
	_, err := r.client.ZeroTrust.Gateway.Certificates.Get(
		ctx,
		path_certificate_id,
		zero_trust.GatewayCertificateGetParams{
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

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustGatewayCertificateResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
