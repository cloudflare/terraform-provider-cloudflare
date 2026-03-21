// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway

import (
	"context"
	"fmt"
	"time"

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

	// Build params with all fields explicitly set
	// Use Null for fields that should be omitted (not sent to API)
	params := ai_gateway.AIGatewayNewParams{
		AccountID:               cloudflare.F(data.AccountID.ValueString()),
		ID:                      cloudflare.F(data.ID.ValueString()),
		CacheInvalidateOnUpdate: cloudflare.F(data.CacheInvalidateOnUpdate.ValueBool()),
		CacheTTL:                cloudflare.F(data.CacheTTL.ValueInt64()),
		CollectLogs:             cloudflare.F(data.CollectLogs.ValueBool()),
		RateLimitingInterval:    cloudflare.F(data.RateLimitingInterval.ValueInt64()),
		RateLimitingLimit:       cloudflare.F(data.RateLimitingLimit.ValueInt64()),
		RateLimitingTechnique:   cloudflare.F(ai_gateway.AIGatewayNewParamsRateLimitingTechnique(data.RateLimitingTechnique.ValueString())),
		Authentication:          cloudflare.F(data.Authentication.ValueBool()),
	}

	// Only set optional fields if they have meaningful values
	if data.IsDefault.ValueBool() {
		params.IsDefault = cloudflare.F(data.IsDefault.ValueBool())
	}
	if data.LogManagement.ValueInt64() >= 10000 {
		params.LogManagement = cloudflare.F(data.LogManagement.ValueInt64())
	}
	if data.LogManagementStrategy.ValueString() != "" {
		params.LogManagementStrategy = cloudflare.F(ai_gateway.AIGatewayNewParamsLogManagementStrategy(data.LogManagementStrategy.ValueString()))
	}
	if data.Logpush.ValueBool() {
		params.Logpush = cloudflare.F(data.Logpush.ValueBool())
		if len(data.LogpushPublicKey.ValueString()) >= 16 {
			params.LogpushPublicKey = cloudflare.F(data.LogpushPublicKey.ValueString())
		}
	}
	if data.ZDR.ValueBool() {
		params.Zdr = cloudflare.F(data.ZDR.ValueBool())
	}

	result, err := r.client.AIGateway.New(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("failed to create AI Gateway", err.Error())
		return
	}

	data.ID = types.StringValue(result.ID)
	if result.AccountID != "" {
		data.AccountID = types.StringValue(result.AccountID)
	}
	if !result.CreatedAt.IsZero() {
		data.CreatedAt = types.StringValue(result.CreatedAt.Format(time.RFC3339))
	}
	if !result.ModifiedAt.IsZero() {
		data.ModifiedAt = types.StringValue(result.ModifiedAt.Format(time.RFC3339))
	}
	// Populate all fields from API response
	data.CacheInvalidateOnUpdate = types.BoolValue(result.CacheInvalidateOnUpdate)
	data.CacheTTL = types.Int64Value(result.CacheTTL)
	data.CollectLogs = types.BoolValue(result.CollectLogs)
	data.RateLimitingInterval = types.Int64Value(result.RateLimitingInterval)
	data.RateLimitingLimit = types.Int64Value(result.RateLimitingLimit)
	data.RateLimitingTechnique = types.StringValue(string(result.RateLimitingTechnique))
	data.Authentication = types.BoolValue(result.Authentication)
	data.IsDefault = types.BoolValue(result.IsDefault)
	data.LogManagement = types.Int64Value(result.LogManagement)
	data.LogManagementStrategy = types.StringValue(string(result.LogManagementStrategy))
	data.Logpush = types.BoolValue(result.Logpush)
	data.LogpushPublicKey = types.StringValue(result.LogpushPublicKey)
	data.StoreID = types.StringValue(result.StoreID)
	data.ZDR = types.BoolValue(result.Zdr)

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

	// Build params with all fields explicitly set (required fields)
	params := ai_gateway.AIGatewayUpdateParams{
		AccountID:               cloudflare.F(data.AccountID.ValueString()),
		CacheInvalidateOnUpdate: cloudflare.F(data.CacheInvalidateOnUpdate.ValueBool()),
		CacheTTL:                cloudflare.F(data.CacheTTL.ValueInt64()),
		CollectLogs:             cloudflare.F(data.CollectLogs.ValueBool()),
		RateLimitingInterval:    cloudflare.F(data.RateLimitingInterval.ValueInt64()),
		RateLimitingLimit:       cloudflare.F(data.RateLimitingLimit.ValueInt64()),
		RateLimitingTechnique:   cloudflare.F(ai_gateway.AIGatewayUpdateParamsRateLimitingTechnique(data.RateLimitingTechnique.ValueString())),
		Authentication:          cloudflare.F(data.Authentication.ValueBool()),
	}

	// Only set optional fields if they have meaningful values
	if data.IsDefault.ValueBool() {
		params.IsDefault = cloudflare.F(data.IsDefault.ValueBool())
	}
	if data.LogManagement.ValueInt64() >= 10000 {
		params.LogManagement = cloudflare.F(data.LogManagement.ValueInt64())
	}
	if data.LogManagementStrategy.ValueString() != "" {
		params.LogManagementStrategy = cloudflare.F(ai_gateway.AIGatewayUpdateParamsLogManagementStrategy(data.LogManagementStrategy.ValueString()))
	}
	if data.Logpush.ValueBool() {
		params.Logpush = cloudflare.F(data.Logpush.ValueBool())
		if len(data.LogpushPublicKey.ValueString()) >= 16 {
			params.LogpushPublicKey = cloudflare.F(data.LogpushPublicKey.ValueString())
		}
	}
	if data.ZDR.ValueBool() {
		params.Zdr = cloudflare.F(data.ZDR.ValueBool())
	}

	result, err := r.client.AIGateway.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("failed to update AI Gateway", err.Error())
		return
	}

	// Populate all fields from API response
	data.ID = types.StringValue(result.ID)
	if result.AccountID != "" {
		data.AccountID = types.StringValue(result.AccountID)
	}
	if !result.CreatedAt.IsZero() {
		data.CreatedAt = types.StringValue(result.CreatedAt.Format(time.RFC3339))
	}
	if !result.ModifiedAt.IsZero() {
		data.ModifiedAt = types.StringValue(result.ModifiedAt.Format(time.RFC3339))
	}
	data.CacheInvalidateOnUpdate = types.BoolValue(result.CacheInvalidateOnUpdate)
	data.CacheTTL = types.Int64Value(result.CacheTTL)
	data.CollectLogs = types.BoolValue(result.CollectLogs)
	data.RateLimitingInterval = types.Int64Value(result.RateLimitingInterval)
	data.RateLimitingLimit = types.Int64Value(result.RateLimitingLimit)
	data.RateLimitingTechnique = types.StringValue(string(result.RateLimitingTechnique))
	data.Authentication = types.BoolValue(result.Authentication)
	data.IsDefault = types.BoolValue(result.IsDefault)
	data.LogManagement = types.Int64Value(result.LogManagement)
	data.LogManagementStrategy = types.StringValue(string(result.LogManagementStrategy))
	data.Logpush = types.BoolValue(result.Logpush)
	data.LogpushPublicKey = types.StringValue(result.LogpushPublicKey)
	data.StoreID = types.StringValue(result.StoreID)
	data.ZDR = types.BoolValue(result.Zdr)

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
	// Only update account_id if API returns a valid value (some APIs return null for account_id)
	if result.AccountID != "" {
		data.AccountID = types.StringValue(result.AccountID)
	}
	if !result.CreatedAt.IsZero() {
		data.CreatedAt = types.StringValue(result.CreatedAt.Format(time.RFC3339))
	}
	if !result.ModifiedAt.IsZero() {
		data.ModifiedAt = types.StringValue(result.ModifiedAt.Format(time.RFC3339))
	}
	// Populate other fields from API response
	data.CacheInvalidateOnUpdate = types.BoolValue(result.CacheInvalidateOnUpdate)
	data.CacheTTL = types.Int64Value(result.CacheTTL)
	data.CollectLogs = types.BoolValue(result.CollectLogs)
	data.RateLimitingInterval = types.Int64Value(result.RateLimitingInterval)
	data.RateLimitingLimit = types.Int64Value(result.RateLimitingLimit)
	data.RateLimitingTechnique = types.StringValue(string(result.RateLimitingTechnique))
	data.Authentication = types.BoolValue(result.Authentication)
	data.IsDefault = types.BoolValue(result.IsDefault)
	data.LogManagement = types.Int64Value(result.LogManagement)
	data.LogManagementStrategy = types.StringValue(string(result.LogManagementStrategy))
	data.Logpush = types.BoolValue(result.Logpush)
	data.LogpushPublicKey = types.StringValue(result.LogpushPublicKey)
	data.StoreID = types.StringValue(result.StoreID)
	data.ZDR = types.BoolValue(result.Zdr)

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
	if !result.CreatedAt.IsZero() {
		data.CreatedAt = types.StringValue(result.CreatedAt.Format(time.RFC3339))
	}
	if !result.ModifiedAt.IsZero() {
		data.ModifiedAt = types.StringValue(result.ModifiedAt.Format(time.RFC3339))
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
