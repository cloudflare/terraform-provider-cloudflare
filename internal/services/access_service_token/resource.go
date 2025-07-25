// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_service_token

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
var _ resource.ResourceWithConfigure = (*AccessServiceTokenResource)(nil)

// AccessServiceTokenResource defines the deprecated resource implementation.
type AccessServiceTokenResource struct {
	client *cloudflare.Client
}

// NewResource returns a new instance of the resource.
func NewResource() resource.Resource {
	return &AccessServiceTokenResource{}
}

func (r *AccessServiceTokenResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_access_service_token"
}

func (r *AccessServiceTokenResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *AccessServiceTokenResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// Create returns an explicit error as the resource is deprecated.
func (r *AccessServiceTokenResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	resp.Diagnostics.AddError(
		"Resource Deprecated",
		"The cloudflare_access_service_token resource has been renamed to cloudflare_zero_trust_access_service_token and is deprecated. Use the new resource instead.",
	)
}

// Read fetches existing state using the current Zero Trust API so that removed blocks can still be validated.
func (r *AccessServiceTokenResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *AccessServiceTokenModel

	secret := types.StringNull()
	req.State.GetAttribute(ctx, path.Root("client_secret"), &secret)

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := AccessServiceTokenResultEnvelope{*data}
	params := zero_trust.AccessServiceTokenGetParams{}

	if !data.AccountID.IsNull() {
		params.AccountID = cloudflare.F(data.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
	}

	_, err := r.client.ZeroTrust.Access.ServiceTokens.Get(
		ctx,
		data.ID.ValueString(),
		params,
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if res != nil && res.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddWarning("Resource not found", "The resource was not found on the server and will be removed from state.")
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	bytes, _ := io.ReadAll(res.Body)
	if err := apijson.Unmarshal(bytes, &env); err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}

	data = &env.Result
	data.ClientSecret = secret

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update returns an explicit error as the resource is deprecated.
func (r *AccessServiceTokenResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError(
		"Resource Deprecated",
		"The cloudflare_access_service_token resource has been renamed to cloudflare_zero_trust_access_service_token and is deprecated. Use the new resource instead.",
	)
}

// Delete returns an explicit error as the resource is deprecated.
func (r *AccessServiceTokenResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddError(
		"Resource Deprecated",
		"The cloudflare_access_service_token resource has been renamed to cloudflare_zero_trust_access_service_token and is deprecated. Use the new resource instead.",
	)
}

func (r *AccessServiceTokenResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
