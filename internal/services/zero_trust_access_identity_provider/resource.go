// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_identity_provider

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*ZeroTrustAccessIdentityProviderResource)(nil)
var _ resource.ResourceWithModifyPlan = (*ZeroTrustAccessIdentityProviderResource)(nil)
var _ resource.ResourceWithImportState = (*ZeroTrustAccessIdentityProviderResource)(nil)

func NewResource() resource.Resource {
	return &ZeroTrustAccessIdentityProviderResource{}
}

// ZeroTrustAccessIdentityProviderResource defines the resource implementation.
type ZeroTrustAccessIdentityProviderResource struct {
	client *cloudflare.Client
}

func (r *ZeroTrustAccessIdentityProviderResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zero_trust_access_identity_provider"
}

func (r *ZeroTrustAccessIdentityProviderResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ZeroTrustAccessIdentityProviderResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ZeroTrustAccessIdentityProviderModel

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
	env := ZeroTrustAccessIdentityProviderResultEnvelope{*data}
	params := zero_trust.IdentityProviderNewParams{}

	if !data.AccountID.IsNull() {
		params.AccountID = cloudflare.F(data.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
	}

	_, err = r.client.ZeroTrust.IdentityProviders.New(
		ctx,
		params,
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

func (r *ZeroTrustAccessIdentityProviderResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ZeroTrustAccessIdentityProviderModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *ZeroTrustAccessIdentityProviderModel

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
	env := ZeroTrustAccessIdentityProviderResultEnvelope{*data}
	params := zero_trust.IdentityProviderUpdateParams{}

	if !data.AccountID.IsNull() {
		params.AccountID = cloudflare.F(data.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
	}

	_, err = r.client.ZeroTrust.IdentityProviders.Update(
		ctx,
		data.ID.ValueString(),
		params,
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

func (r *ZeroTrustAccessIdentityProviderResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ZeroTrustAccessIdentityProviderModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := ZeroTrustAccessIdentityProviderResultEnvelope{*data}
	params := zero_trust.IdentityProviderGetParams{}

	if !data.AccountID.IsNull() {
		params.AccountID = cloudflare.F(data.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
	}

	_, err := r.client.ZeroTrust.IdentityProviders.Get(
		ctx,
		data.ID.ValueString(),
		params,
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

func (r *ZeroTrustAccessIdentityProviderResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ZeroTrustAccessIdentityProviderModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	params := zero_trust.IdentityProviderDeleteParams{}

	if !data.AccountID.IsNull() {
		params.AccountID = cloudflare.F(data.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
	}

	_, err := r.client.ZeroTrust.IdentityProviders.Delete(
		ctx,
		data.ID.ValueString(),
		params,
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustAccessIdentityProviderResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data *ZeroTrustAccessIdentityProviderModel
	params := zero_trust.IdentityProviderGetParams{}

	path_accounts_or_zones, path_account_id_or_zone_id := "", ""
	path_identity_provider_id := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<{accounts|zones}/{account_id|zone_id}>/<identity_provider_id>",
		&path_accounts_or_zones,
		&path_account_id_or_zone_id,
		&path_identity_provider_id,
	)
	resp.Diagnostics.Append(diags...)
	switch path_accounts_or_zones {
	case "accounts":
		params.AccountID = cloudflare.F(path_account_id_or_zone_id)
	case "zones":
		params.ZoneID = cloudflare.F(path_account_id_or_zone_id)
	default:
		resp.Diagnostics.AddError("invalid discriminator segment - <{accounts|zones}/{account_id|zone_id}>", "expected discriminator to be one of {accounts|zones}")
	}
	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := ZeroTrustAccessIdentityProviderResultEnvelope{*data}
	_, err := r.client.ZeroTrust.IdentityProviders.Get(
		ctx,
		path_identity_provider_id,
		params,
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

func (r *ZeroTrustAccessIdentityProviderResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
