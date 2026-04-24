// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_custom_domain

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/r2"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*R2CustomDomainResource)(nil)
var _ resource.ResourceWithModifyPlan = (*R2CustomDomainResource)(nil)

func NewResource() resource.Resource {
	return &R2CustomDomainResource{}
}

// R2CustomDomainResource defines the resource implementation.
type R2CustomDomainResource struct {
	client *cloudflare.Client
}

func (r *R2CustomDomainResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_r2_custom_domain"
}

func (r *R2CustomDomainResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *R2CustomDomainResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *R2CustomDomainModel

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
	env := R2CustomDomainResultEnvelope{*data}
	_, err = r.client.R2.Buckets.Domains.Custom.New(
		ctx,
		data.BucketName.ValueString(),
		r2.BucketDomainCustomNewParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithHeader(consts.R2JurisdictionHTTPHeaderName, data.Jurisdiction.ValueString()),
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

	// If the create response is degraded (zone_name is null), do a follow-up GET
	// to retrieve the complete data. This can happen if the SSL4SaaS backend is
	// transiently unavailable during the create operation.
	if data.ZoneName.IsNull() {
		res = new(http.Response)
		env = R2CustomDomainResultEnvelope{*data}
		_, err = r.client.R2.Buckets.Domains.Custom.Get(
			ctx,
			data.BucketName.ValueString(),
			data.Domain.ValueString(),
			r2.BucketDomainCustomGetParams{
				AccountID: cloudflare.F(data.AccountID.ValueString()),
			},
			option.WithHeader(consts.R2JurisdictionHTTPHeaderName, data.Jurisdiction.ValueString()),
			option.WithResponseBodyInto(&res),
			option.WithMiddleware(logging.Middleware(ctx)),
		)
		if err == nil {
			bytes, _ = io.ReadAll(res.Body)
			_ = apijson.Unmarshal(bytes, &env)
			data = &env.Result
		}
		// If the GET also fails or returns degraded data, we'll just use what we have
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *R2CustomDomainResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *R2CustomDomainModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *R2CustomDomainModel

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
	env := R2CustomDomainResultEnvelope{*data}
	_, err = r.client.R2.Buckets.Domains.Custom.Update(
		ctx,
		data.BucketName.ValueString(),
		data.Domain.ValueString(),
		r2.BucketDomainCustomUpdateParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithHeader(consts.R2JurisdictionHTTPHeaderName, data.Jurisdiction.ValueString()),
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

	// If the update response is degraded (zone_name is null), do a follow-up GET
	// to retrieve the complete data. This can happen if the SSL4SaaS backend is
	// transiently unavailable during the update operation.
	if data.ZoneName.IsNull() {
		res = new(http.Response)
		env = R2CustomDomainResultEnvelope{*data}
		_, err = r.client.R2.Buckets.Domains.Custom.Get(
			ctx,
			data.BucketName.ValueString(),
			data.Domain.ValueString(),
			r2.BucketDomainCustomGetParams{
				AccountID: cloudflare.F(data.AccountID.ValueString()),
			},
			option.WithHeader(consts.R2JurisdictionHTTPHeaderName, data.Jurisdiction.ValueString()),
			option.WithResponseBodyInto(&res),
			option.WithMiddleware(logging.Middleware(ctx)),
		)
		if err == nil {
			bytes, _ = io.ReadAll(res.Body)
			_ = apijson.Unmarshal(bytes, &env)
			data = &env.Result
		}
		// If the GET also fails or returns degraded data, we'll just use what we have
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *R2CustomDomainResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *R2CustomDomainModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Snapshot the current state before the API response overwrites it.
	// This allows us to detect and recover from degraded API responses
	// where the SSL4SaaS backend is transiently unavailable.
	previousState := snapshotState(data)

	res := new(http.Response)
	env := R2CustomDomainResultEnvelope{*data}
	_, err := r.client.R2.Buckets.Domains.Custom.Get(
		ctx,
		data.BucketName.ValueString(),
		data.Domain.ValueString(),
		r2.BucketDomainCustomGetParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithHeader(consts.R2JurisdictionHTTPHeaderName, data.Jurisdiction.ValueString()),
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

	// When the API returns a degraded response (status "unknown"/"unknown"),
	// preserve the previous state values for status, zone_name, and min_tls
	// to prevent false drift detection and unnecessary resource replacement.
	preserveStateOnDegradedResponse(ctx, data, previousState)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *R2CustomDomainResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *R2CustomDomainModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.R2.Buckets.Domains.Custom.Delete(
		ctx,
		data.BucketName.ValueString(),
		data.Domain.ValueString(),
		r2.BucketDomainCustomDeleteParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithHeader(consts.R2JurisdictionHTTPHeaderName, data.Jurisdiction.ValueString()),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *R2CustomDomainResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
