// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package snippets

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*SnippetsResource)(nil)
var _ resource.ResourceWithModifyPlan = (*SnippetsResource)(nil)

func NewResource() resource.Resource {
	return &SnippetsResource{}
}

// SnippetsResource defines the resource implementation.
type SnippetsResource struct {
	client *cloudflare.Client
}

func (r *SnippetsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snippets"
}

func (r *SnippetsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

var ErrNonfunctionalResource = fmt.Errorf("use 'cloudflare_snippet' instead of 'cloudflare_snippets'")

func (r *SnippetsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	resp.Diagnostics.AddError("resource is non-functional", ErrNonfunctionalResource.Error())
}

func (r *SnippetsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("resource is non-functional", ErrNonfunctionalResource.Error())
}

func (r *SnippetsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	resp.Diagnostics.AddError("resource is non-functional", ErrNonfunctionalResource.Error())
}

func (r *SnippetsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddError("resource is non-functional", ErrNonfunctionalResource.Error())
}

func (r *SnippetsResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
