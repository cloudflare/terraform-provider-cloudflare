// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package argo_tiered_caching

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithUpgradeState = (*ArgoTieredCachingResource)(nil)
var _ resource.ResourceWithMoveState = (*ArgoTieredCachingResource)(nil)

func (r *ArgoTieredCachingResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}

// MoveState implements resource.ResourceWithMoveState
// This allows moving from cloudflare_tiered_cache resources with cache_type="generic"
// to cloudflare_argo_tiered_caching resources
func (r *ArgoTieredCachingResource) MoveState(ctx context.Context) []resource.StateMover {
	return []resource.StateMover{
		{
			// Define the source schema from v4 cloudflare_tiered_cache
			SourceSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"zone_id": schema.StringAttribute{
						Required: true,
					},
					"cache_type": schema.StringAttribute{
						Required: true,
					},
				},
			},
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				// Only handle moves from cloudflare_tiered_cache
				if req.SourceTypeName != "cloudflare_tiered_cache" {
					// Skip - not the source resource type we handle
					return
				}

				// Get the source state data
				var sourceState struct {
					ID        types.String `tfsdk:"id"`
					ZoneID    types.String `tfsdk:"zone_id"`
					CacheType types.String `tfsdk:"cache_type"`
				}

				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				// Only handle cache_type="generic" moves
				if sourceState.CacheType.ValueString() != "generic" {
					// Skip - not the cache_type we handle
					return
				}

				// Create the target state for argo_tiered_caching
				targetState := ArgoTieredCachingModel{
					ID:     sourceState.ID,
					ZoneID: sourceState.ZoneID,
					Value:  types.StringValue("on"), // generic maps to "on"
				}

				// Set the target state
				resp.Diagnostics.Append(resp.TargetState.Set(ctx, targetState)...)
			},
		},
		{
			// Also handle v5 tiered_cache resources (after Grit renames cache_type to value)
			SourceSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"zone_id": schema.StringAttribute{
						Required: true,
					},
					"value": schema.StringAttribute{
						Required: true,
					},
				},
			},
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				// Only handle moves from cloudflare_tiered_cache
				if req.SourceTypeName != "cloudflare_tiered_cache" {
					// Skip - not the source resource type we handle
					return
				}

				// Get the source state data
				var sourceState struct {
					ID     types.String `tfsdk:"id"`
					ZoneID types.String `tfsdk:"zone_id"`
					Value  types.String `tfsdk:"value"`
				}

				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				// Only handle value="generic" moves (after Grit transformation)
				if sourceState.Value.ValueString() != "generic" {
					// Skip - not the value we handle
					return
				}

				// Create the target state for argo_tiered_caching
				targetState := ArgoTieredCachingModel{
					ID:     sourceState.ID,
					ZoneID: sourceState.ZoneID,
					Value:  types.StringValue("on"), // generic maps to "on"
				}

				// Set the target state
				resp.Diagnostics.Append(resp.TargetState.Set(ctx, targetState)...)
			},
		},
	}
}
