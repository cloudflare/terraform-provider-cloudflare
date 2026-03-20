// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/ai_gateway"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigure = (*AIGatewayResource)(nil)
var _ resource.ResourceWithImportState = (*AIGatewayResource)(nil)

func NewResource() resource.Resource {
	return &AIGatewayResource{}
}

type AIGatewayResource struct {
	client *cloudflare.Client
}

func (r *AIGatewayResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ai_gateway"
}

func (r *AIGatewayResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *AIGatewayResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *AIGatewayModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	result, err := r.client.AIGateway.New(
		ctx,
		ai_gateway.AIGatewayNewParams{
			AccountID:               cloudflare.F(data.AccountID.ValueString()),
			ID:                      cloudflare.F(data.ID.ValueString()),
			CacheInvalidateOnUpdate: cloudflare.F(data.CacheInvalidateOnUpdate.ValueBool()),
			CacheTTL:                cloudflare.F(data.CacheTTL.ValueInt64()),
			CollectLogs:             cloudflare.F(data.CollectLogs.ValueBool()),
			RateLimitingInterval:    cloudflare.F(data.RateLimitingInterval.ValueInt64()),
			RateLimitingLimit:       cloudflare.F(data.RateLimitingLimit.ValueInt64()),
			RateLimitingTechnique:   cloudflare.F(ai_gateway.AIGatewayNewParamsRateLimitingTechnique(data.RateLimitingTechnique.ValueString())),
			Authentication:          cloudflare.F(data.Authentication.ValueBool()),
			IsDefault:               cloudflare.F(data.IsDefault.ValueBool()),
			LogManagement:           cloudflare.F(data.LogManagement.ValueInt64()),
			LogManagementStrategy:   cloudflare.F(ai_gateway.AIGatewayNewParamsLogManagementStrategy(data.LogManagementStrategy.ValueString())),
			Logpush:                 cloudflare.F(data.Logpush.ValueBool()),
			LogpushPublicKey:        cloudflare.F(data.LogpushPublicKey.ValueString()),
			Zdr:                     cloudflare.F(data.ZDR.ValueBool()),
		},
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to create AI Gateway", err.Error())
		return
	}

	data.ID = types.StringValue(result.ID)
	data.AccountID = types.StringValue(result.AccountID)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AIGatewayResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *AIGatewayModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *AIGatewayModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	result, err := r.client.AIGateway.Update(
		ctx,
		state.ID.ValueString(),
		ai_gateway.AIGatewayUpdateParams{
			AccountID:               cloudflare.F(data.AccountID.ValueString()),
			CacheInvalidateOnUpdate: cloudflare.F(data.CacheInvalidateOnUpdate.ValueBool()),
			CacheTTL:                cloudflare.F(data.CacheTTL.ValueInt64()),
			CollectLogs:             cloudflare.F(data.CollectLogs.ValueBool()),
			RateLimitingInterval:    cloudflare.F(data.RateLimitingInterval.ValueInt64()),
			RateLimitingLimit:       cloudflare.F(data.RateLimitingLimit.ValueInt64()),
			RateLimitingTechnique:   cloudflare.F(ai_gateway.AIGatewayUpdateParamsRateLimitingTechnique(data.RateLimitingTechnique.ValueString())),
			Authentication:          cloudflare.F(data.Authentication.ValueBool()),
			IsDefault:               cloudflare.F(data.IsDefault.ValueBool()),
			LogManagement:           cloudflare.F(data.LogManagement.ValueInt64()),
			LogManagementStrategy:   cloudflare.F(ai_gateway.AIGatewayUpdateParamsLogManagementStrategy(data.LogManagementStrategy.ValueString())),
			Logpush:                 cloudflare.F(data.Logpush.ValueBool()),
			LogpushPublicKey:        cloudflare.F(data.LogpushPublicKey.ValueString()),
			Zdr:                     cloudflare.F(data.ZDR.ValueBool()),
		},
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to update AI Gateway", err.Error())
		return
	}

	data.ID = types.StringValue(result.ID)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AIGatewayResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *AIGatewayModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	result, err := r.client.AIGateway.Get(
		ctx,
		data.ID.ValueString(),
		ai_gateway.AIGatewayGetParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to get AI Gateway", err.Error())
		return
	}

	data.ID = types.StringValue(result.ID)
	data.AccountID = types.StringValue(result.AccountID)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AIGatewayResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *AIGatewayModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.AIGateway.Delete(
		ctx,
		data.ID.ValueString(),
		ai_gateway.AIGatewayDeleteParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to delete AI Gateway", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AIGatewayResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data = new(AIGatewayModel)

	path_account_id := ""
	path_id := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<account_id>/<id>",
		&path_account_id,
		&path_id,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.AccountID = types.StringValue(path_account_id)
	data.ID = types.StringValue(path_id)

	result, err := r.client.AIGateway.Get(
		ctx,
		path_id,
		ai_gateway.AIGatewayGetParams{
			AccountID: cloudflare.F(path_account_id),
		},
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to get AI Gateway", err.Error())
		return
	}

	data.ID = types.StringValue(result.ID)
	data.AccountID = types.StringValue(result.AccountID)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
