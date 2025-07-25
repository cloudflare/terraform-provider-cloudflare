// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*AccessPolicyResource)(nil)

type AccessPolicyResource struct {
	client *cloudflare.Client
}

func NewResource() resource.Resource {
	return &AccessPolicyResource{}
}

func (r *AccessPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_access_policy"
}

func (r *AccessPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *AccessPolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*cloudflare.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *cloudflare.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = client
}

func (r *AccessPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	resp.Diagnostics.AddError(
		"Resource Deprecated",
		"The cloudflare_access_policy resource has been deprecated. Please use cloudflare_zero_trust_access_policy instead.",
	)
}

func (r *AccessPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *AccessPolicyModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Use the Zero Trust Access Policy API to read the policy data
	// This allows existing access_policy resources to be read even though they're deprecated
	res := new(http.Response)
	_, err := r.client.ZeroTrust.Access.Policies.Get(
		ctx,
		data.ID.ValueString(),
		zero_trust.AccessPolicyGetParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if res != nil && res.StatusCode == 404 {
		resp.Diagnostics.AddWarning(
			"Resource Not Found",
			"The access policy was not found and will be removed from state.",
		)
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Access Policy",
			fmt.Sprintf("Could not read access policy: %s", err),
		)
		return
	}

	// Keep the resource in state for backward compatibility
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r *AccessPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError(
		"Resource Deprecated",
		"The cloudflare_access_policy resource has been deprecated. Please use cloudflare_zero_trust_access_policy instead.",
	)
}

func (r *AccessPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddError(
		"Resource Deprecated",
		"The cloudflare_access_policy resource has been deprecated. Please use cloudflare_zero_trust_access_policy instead.",
	)
}
