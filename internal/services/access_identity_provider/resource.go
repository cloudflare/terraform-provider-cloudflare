// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_identity_provider

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*AccessIdentityProviderResource)(nil)

// AccessIdentityProviderResource defines the deprecated resource implementation.
type AccessIdentityProviderResource struct {
	client *cloudflare.Client
}

// NewResource returns a new instance of the resource.
func NewResource() resource.Resource {
	return &AccessIdentityProviderResource{}
}

func (r *AccessIdentityProviderResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_access_identity_provider"
}

func (r *AccessIdentityProviderResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*cloudflare.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *cloudflare.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

// Schema implementation for the resource interface.
func (r *AccessIdentityProviderResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

// No custom validators required.
func (r *AccessIdentityProviderResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}

func (r *AccessIdentityProviderResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	resp.Diagnostics.AddError(
		"Resource Deprecated",
		"The cloudflare_access_identity_provider resource has been deprecated and renamed to cloudflare_zero_trust_access_identity_provider in v5. Please update your configuration.",
	)
}

func (r *AccessIdentityProviderResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError(
		"Resource Deprecated",
		"The cloudflare_access_identity_provider resource has been deprecated and renamed to cloudflare_zero_trust_access_identity_provider in v5. Please update your configuration.",
	)
}

func (r *AccessIdentityProviderResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddError(
		"Resource Deprecated",
		"The cloudflare_access_identity_provider resource has been deprecated and renamed to cloudflare_zero_trust_access_identity_provider in v5. Please update your configuration.",
	)
}

func (r *AccessIdentityProviderResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *AccessIdentityProviderModel
	secret := types.StringNull()
	req.State.GetAttribute(ctx, path.Root("scim_config").AtName("secret"), &secret)

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := AccessIdentityProviderResultEnvelope{*data}
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

	// Preserve the secret from state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("scim_config").AtName("secret"), secret)...)
}

func (r *AccessIdentityProviderResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data *AccessIdentityProviderModel = new(AccessIdentityProviderModel)

	// For deprecated resources, we'll use the ID as the import ID
	data.ID = types.StringValue(req.ID)

	res := new(http.Response)
	env := AccessIdentityProviderResultEnvelope{*data}
	params := zero_trust.IdentityProviderGetParams{}

	// Try to determine account_id or zone_id from the import ID
	// This is a simplified approach - in practice, you might need to parse the ID differently
	if !data.AccountID.IsNull() {
		params.AccountID = cloudflare.F(data.AccountID.ValueString())
	} else if !data.ZoneID.IsNull() {
		params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
	}

	_, err := r.client.ZeroTrust.IdentityProviders.Get(
		ctx,
		req.ID,
		params,
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
