// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package argo_tunnel

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
var _ resource.ResourceWithConfigure = (*ArgoTunnelResource)(nil)

type ArgoTunnelResource struct {
	client *cloudflare.Client
}

func NewResource() resource.Resource {
	return &ArgoTunnelResource{}
}

func (r *ArgoTunnelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_argo_tunnel"
}

func (r *ArgoTunnelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ArgoTunnelResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ArgoTunnelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// This resource is deprecated and only exists for schema validation
	resp.Diagnostics.AddError(
		"Resource Deprecated",
		"The cloudflare_argo_tunnel resource has been deprecated. Use cloudflare_zero_trust_tunnel_cloudflared instead.",
	)
}

func (r *ArgoTunnelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ArgoTunnelModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Use the Zero Trust Tunnel API to read the tunnel data
	// This allows existing argo_tunnel resources to be read even though they're deprecated
	res := new(http.Response)
	_, err := r.client.ZeroTrust.Tunnels.Cloudflared.Get(
		ctx,
		data.ID.ValueString(),
		zero_trust.TunnelCloudflaredGetParams{
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

	// For argo_tunnel, we'll keep the existing state since this is a deprecated resource
	// The read operation is mainly for state validation and cleanup
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ArgoTunnelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// This resource is deprecated and only exists for schema validation
	resp.Diagnostics.AddError(
		"Resource Deprecated",
		"The cloudflare_argo_tunnel resource has been deprecated. Use cloudflare_zero_trust_tunnel_cloudflared instead.",
	)
}

func (r *ArgoTunnelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// This resource is deprecated and only exists for schema validation
	resp.Diagnostics.AddError(
		"Resource Deprecated",
		"The cloudflare_argo_tunnel resource has been deprecated. Use cloudflare_zero_trust_tunnel_cloudflared instead.",
	)
}

func (r *ArgoTunnelResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
