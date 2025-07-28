package fallback_domain

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
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*FallbackDomainResource)(nil)

// FallbackDomainResource defines the deprecated resource implementation.
type FallbackDomainResource struct {
	client *cloudflare.Client
}

// NewResource returns a new instance of the resource.
func NewResource() resource.Resource {
	return &FallbackDomainResource{}
}

func (r *FallbackDomainResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fallback_domain"
}

func (r *FallbackDomainResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *FallbackDomainResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

// No custom validators required.
func (r *FallbackDomainResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}

func (r *FallbackDomainResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	resp.Diagnostics.AddError(
		"Resource Deprecated",
		"The cloudflare_fallback_domain resource has been deprecated and renamed to cloudflare_zero_trust_device_default_profile_local_domain_fallback in v5. Please update your configuration.",
	)
}

func (r *FallbackDomainResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError(
		"Resource Deprecated",
		"The cloudflare_fallback_domain resource has been deprecated and renamed to cloudflare_zero_trust_device_default_profile_local_domain_fallback in v5. Please update your configuration.",
	)
}

func (r *FallbackDomainResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddError(
		"Resource Deprecated",
		"The cloudflare_fallback_domain resource has been deprecated and renamed to cloudflare_zero_trust_device_default_profile_local_domain_fallback in v5. Please update your configuration.",
	)
}

func (r *FallbackDomainResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *FallbackDomainModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := FallbackDomainResultEnvelope{data.Domains}
	_, err := r.client.ZeroTrust.Devices.Policies.Default.FallbackDomains.Get(
		ctx,
		zero_trust.DevicePolicyDefaultFallbackDomainGetParams{
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
	data.Domains = env.Result
	data.ID = data.AccountID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FallbackDomainResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data *FallbackDomainModel = new(FallbackDomainModel)

	// For deprecated resources, we'll use the account_id as the import ID
	data.AccountID = types.StringValue(req.ID)

	res := new(http.Response)
	env := FallbackDomainResultEnvelope{data.Domains}
	_, err := r.client.ZeroTrust.Devices.Policies.Default.FallbackDomains.Get(
		ctx,
		zero_trust.DevicePolicyDefaultFallbackDomainGetParams{
			AccountID: cloudflare.F(req.ID),
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
	data.Domains = env.Result
	data.ID = data.AccountID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
