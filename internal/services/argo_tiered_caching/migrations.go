// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package argo_tiered_caching

import (
	"context"
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/argo_tiered_caching/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.ResourceWithUpgradeState = (*ArgoTieredCachingResource)(nil)
var _ resource.ResourceWithMoveState = (*ArgoTieredCachingResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
//
// This handles two upgrade paths:
// 1. v4 state (schema_version=0) → v5 (version=500): Full transformation from cloudflare_argo
// 2. v5 state (version=1) → v5 (version=500): No-op upgrade
//
// The separation of schema versions (v4=0, v5=1/500) eliminates the need for
// dual-format detection that was required in earlier implementations.
func (r *ArgoTieredCachingResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)

	sourceSchema := v500.SourceCloudflareArgoSchema()

	return map[int64]resource.StateUpgrader{
		// Handle state from v4 SDKv2 provider cloudflare_argo (schema_version=0)
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeArgoToTieredCaching,
		},

		// Handle state from v5 Plugin Framework provider with version=1
		// This is a no-op upgrade that just bumps the version to 500
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV5TieredCaching,
		},
	}
}

// tieredCacheSourceSchema defines the source schema for moves from cloudflare_tiered_cache
// This represents the v4 cloudflare_tiered_cache schema that we're moving FROM
var tieredCacheSourceSchema = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:      true,
			PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
		},
		"zone_id": schema.StringAttribute{
			Required:      true,
			PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
		},
		"cache_type": schema.StringAttribute{
			Required: true,
		},
		"editable": schema.BoolAttribute{
			Computed: true,
		},
		"modified_on": schema.StringAttribute{
			Computed:   true,
			CustomType: timetypes.RFC3339Type{},
		},
	},
}

// MoveState implements ResourceWithMoveState interface
// This enables moving state from TWO different source resources:
// 1. cloudflare_tiered_cache (with cache_type="generic") → cloudflare_argo_tiered_caching (with value="on")
// 2. cloudflare_argo (with tiered_caching attribute) → cloudflare_argo_tiered_caching (with value from tiered_caching)
func (r *ArgoTieredCachingResource) MoveState(ctx context.Context) []resource.StateMover {
	argoSourceSchema := v500.SourceCloudflareArgoSchema()

	return []resource.StateMover{
		// StateMover 1: cloudflare_argo → cloudflare_argo_tiered_caching
		{
			SourceSchema: &argoSourceSchema,
			StateMover:   v500.MoveArgoToTieredCaching,
		},
		// StateMover 2: cloudflare_tiered_cache → cloudflare_argo_tiered_caching
		{
			SourceSchema: &tieredCacheSourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				tflog.Info(ctx, "Starting state move from cloudflare_tiered_cache to cloudflare_argo_tiered_caching")

				// Define source state structure (from cloudflare_tiered_cache)
				var sourceState struct {
					ID         types.String      `tfsdk:"id"`
					ZoneID     types.String      `tfsdk:"zone_id"`
					CacheType  types.String      `tfsdk:"cache_type"`
					Editable   types.Bool        `tfsdk:"editable"`
					ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on"`
				}

				// Get the source state
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					tflog.Error(ctx, "Failed to get source state during move")
					return
				}

				tflog.Debug(ctx, "Source state retrieved", map[string]interface{}{
					"zone_id":    sourceState.ZoneID.ValueString(),
					"cache_type": sourceState.CacheType.ValueString(),
					"id":         sourceState.ID.ValueString(),
				})

				// Validate that this is a generic tiered cache (the only type we should move)
				if !sourceState.CacheType.IsNull() && sourceState.CacheType.ValueString() != "generic" {
					resp.Diagnostics.AddError(
						"Invalid State Move",
						fmt.Sprintf("Cannot move cloudflare_tiered_cache with cache_type='%s' to cloudflare_argo_tiered_caching. Only cache_type='generic' should be moved to this resource type.", sourceState.CacheType.ValueString()),
					)
					return
				}

				// Create the target state (for cloudflare_argo_tiered_caching)
				targetState := ArgoTieredCachingModel{
					ID:       sourceState.ID,       // Preserve ID
					ZoneID:   sourceState.ZoneID,   // Preserve zone_id
					Editable: sourceState.Editable, // Preserve editable
					// Transform cache_type="generic" to value="on"
					Value: types.StringValue("on"),
				}

				// Handle ModifiedOn (might be null in some states)
				if !sourceState.ModifiedOn.IsNull() {
					targetState.ModifiedOn = sourceState.ModifiedOn
				}

				tflog.Info(ctx, "State move transformation completed", map[string]interface{}{
					"source_cache_type": sourceState.CacheType.ValueString(),
					"target_value":      "on",
					"zone_id":           targetState.ZoneID.ValueString(),
				})

				// Set the target state
				resp.Diagnostics.Append(resp.TargetState.Set(ctx, &targetState)...)
				if resp.Diagnostics.HasError() {
					tflog.Error(ctx, "Failed to set target state during move")
					return
				}

				tflog.Info(ctx, "State move from cloudflare_tiered_cache to cloudflare_argo_tiered_caching completed successfully")
			},
		},
	}
}
