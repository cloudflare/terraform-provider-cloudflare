// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sso_connector

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*SSOConnectorResource)(nil)
var _ resource.ResourceWithModifyPlan = (*SSOConnectorResource)(nil)
var _ resource.ResourceWithImportState = (*SSOConnectorResource)(nil)

func NewResource() resource.Resource {
	return &SSOConnectorResource{}
}

// SSOConnectorResource defines the resource implementation.
type SSOConnectorResource struct {
	client *cloudflare.Client
}

func (r *SSOConnectorResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sso_connector"
}

func (r *SSOConnectorResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SSOConnectorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	resp.Diagnostics.AddError(
		"SSO Connector not supported",
		"The SSO Connector resource is not currently supported as the SSO service is not available in the cloudflare-go SDK.",
	)
}

func (r *SSOConnectorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError(
		"SSO Connector not supported",
		"The SSO Connector resource is not currently supported as the SSO service is not available in the cloudflare-go SDK.",
	)
}

func (r *SSOConnectorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	resp.Diagnostics.AddError(
		"SSO Connector not supported",
		"The SSO Connector resource is not currently supported as the SSO service is not available in the cloudflare-go SDK.",
	)
}

func (r *SSOConnectorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddError(
		"SSO Connector not supported",
		"The SSO Connector resource is not currently supported as the SSO service is not available in the cloudflare-go SDK.",
	)
}

func (r *SSOConnectorResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError(
		"SSO Connector not supported",
		"The SSO Connector resource is not currently supported as the SSO service is not available in the cloudflare-go SDK.",
	)
}

func (r *SSOConnectorResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
