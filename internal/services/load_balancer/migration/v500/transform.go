package v500

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts source (v4 SDKv2) state to target (v5 Framework) state.
//
// This handles all the transformations needed for load_balancer migration:
// - Field renames: default_pool_ids → default_pools, fallback_pool_id → fallback_pool
// - Type conversions: Int64 → Float64 for ttl, session_affinity_ttl, drain_duration
// - Structure changes: Arrays → single objects, Sets → maps
// - Nested transformations: Rules with overrides
func Transform(ctx context.Context, source SourceCloudflareLoadBalancerModel) (*TargetLoadBalancerModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Step 1: Validate required fields
	if source.ZoneID.IsNull() || source.ZoneID.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"zone_id is required for load_balancer migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}
	if source.Name.IsNull() || source.Name.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"name is required for load_balancer migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}
	if source.FallbackPoolId.IsNull() || source.FallbackPoolId.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"fallback_pool_id is required for load_balancer migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	// Step 2: Initialize target with direct copies
	target := &TargetLoadBalancerModel{
		ID:              source.ID,
		ZoneID:          source.ZoneID,
		Name:            source.Name,
		Description:     source.Description,
		SessionAffinity: source.SessionAffinity,
		Proxied:         source.Proxied,
		Enabled:         source.Enabled,
		SteeringPolicy:  source.SteeringPolicy,
		CreatedOn:       source.CreatedOn,
		ModifiedOn:      source.ModifiedOn,
		ZoneName:        types.StringNull(), // Computed, will refresh from API
	}

	// Step 3: Handle field renames
	target.FallbackPool = source.FallbackPoolId // fallback_pool_id → fallback_pool

	// Convert default_pool_ids (List) → default_pools (pointer to slice)
	if !source.DefaultPoolIds.IsNull() && !source.DefaultPoolIds.IsUnknown() {
		var poolIDs []types.String
		diags.Append(source.DefaultPoolIds.ElementsAs(ctx, &poolIDs, false)...)
		if !diags.HasError() {
			target.DefaultPools = &poolIDs
		}
	}

	// Step 4: Handle type conversions (Int64 → Float64)
	target.TTL = convertInt64ToFloat64Nullable(source.TTL)
	target.SessionAffinityTTL = convertInt64ToFloat64Nullable(source.SessionAffinityTTL)

	// Step 5: Transform single-object fields (Array[0] → single object)
	// session_affinity_attributes
	if len(source.SessionAffinityAttributes) > 0 {
		target.SessionAffinityAttributes = transformSessionAffinityAttributes(ctx, source.SessionAffinityAttributes[0])
	} else {
		target.SessionAffinityAttributes = customfield.NullObject[TargetSessionAffinityAttributesModel](ctx)
	}

	// adaptive_routing
	if len(source.AdaptiveRouting) > 0 {
		target.AdaptiveRouting = transformAdaptiveRouting(ctx, source.AdaptiveRouting[0])
	} else {
		target.AdaptiveRouting = customfield.NullObject[TargetAdaptiveRoutingModel](ctx)
	}

	// location_strategy
	if len(source.LocationStrategy) > 0 {
		target.LocationStrategy = transformLocationStrategy(ctx, source.LocationStrategy[0])
	} else {
		target.LocationStrategy = customfield.NullObject[TargetLocationStrategyModel](ctx)
	}

	// random_steering
	if len(source.RandomSteering) > 0 {
		target.RandomSteering = transformRandomSteering(ctx, source.RandomSteering[0])
	} else {
		target.RandomSteering = customfield.NullObject[TargetRandomSteeringModel](ctx)
	}

	// Step 6: Transform pool blocks (Set of blocks → Map)
	target.RegionPools = transformPoolsToMap(ctx, source.RegionPools, "region")
	target.POPPools = transformPoolsToMap(ctx, source.PopPools, "pop")
	target.CountryPools = transformPoolsToMap(ctx, source.CountryPools, "country")

	// Step 7: Transform rules (with nested overrides)
	if len(source.Rules) > 0 {
		rules, ruleDiags := transformRules(ctx, source.Rules)
		diags.Append(ruleDiags...)
		target.Rules = rules
	} else {
		target.Rules = customfield.NullObjectList[TargetRulesModel](ctx)
	}

	// Step 8: Set networks to null (new field in v5, will refresh from API)
	target.Networks = customfield.NullList[types.String](ctx)

	return target, diags
}

// transformSessionAffinityAttributes converts v4 session_affinity_attributes to v5 format.
func transformSessionAffinityAttributes(ctx context.Context, source SourceSessionAffinityAttributesModel) customfield.NestedObject[TargetSessionAffinityAttributesModel] {
	target := TargetSessionAffinityAttributesModel{
		Samesite:             source.Samesite,
		Secure:               source.Secure,
		DrainDuration:        convertInt64ToFloat64Nullable(source.DrainDuration),
		ZeroDowntimeFailover: source.ZeroDowntimeFailover,
		RequireAllHeaders:    source.RequireAllHeaders,
	}

	// Convert headers list
	if !source.Headers.IsNull() && !source.Headers.IsUnknown() {
		var headers []types.String
		source.Headers.ElementsAs(ctx, &headers, false)
		target.Headers = &headers
	}

	return customfield.NewObjectMust(ctx, &target)
}

// transformAdaptiveRouting converts v4 adaptive_routing to v5 format.
func transformAdaptiveRouting(ctx context.Context, source SourceAdaptiveRoutingModel) customfield.NestedObject[TargetAdaptiveRoutingModel] {
	target := TargetAdaptiveRoutingModel{
		FailoverAcrossPools: source.FailoverAcrossPools,
	}
	return customfield.NewObjectMust(ctx, &target)
}

// transformLocationStrategy converts v4 location_strategy to v5 format.
func transformLocationStrategy(ctx context.Context, source SourceLocationStrategyModel) customfield.NestedObject[TargetLocationStrategyModel] {
	target := TargetLocationStrategyModel{
		PreferECS: source.PreferECS,
		Mode:      source.Mode,
	}
	return customfield.NewObjectMust(ctx, &target)
}

// transformRandomSteering converts v4 random_steering to v5 format.
func transformRandomSteering(ctx context.Context, source SourceRandomSteeringModel) customfield.NestedObject[TargetRandomSteeringModel] {
	target := TargetRandomSteeringModel{
		DefaultWeight: source.DefaultWeight,
	}

	// Convert pool_weights map
	if !source.PoolWeights.IsNull() && !source.PoolWeights.IsUnknown() {
		var poolWeights map[string]types.Float64
		source.PoolWeights.ElementsAs(ctx, &poolWeights, false)
		target.PoolWeights = &poolWeights
	}

	return customfield.NewObjectMust(ctx, &target)
}

// transformPoolsToMap converts a set of pool blocks to a map.
// Used for region_pools, pop_pools, and country_pools.
//
// v4: [{region: "WNAM", pool_ids: ["pool1"]}, {region: "ENAM", pool_ids: ["pool2"]}]
// v5: {"WNAM": ["pool1"], "ENAM": ["pool2"]}
func transformPoolsToMap(ctx context.Context, sourceBlocks interface{}, keyField string) customfield.Map[customfield.List[types.String]] {
	var pools map[string]customfield.List[types.String]

	switch keyField {
	case "region":
		if regionPools, ok := sourceBlocks.([]SourceRegionPoolModel); ok && len(regionPools) > 0 {
			pools = make(map[string]customfield.List[types.String])
			for _, block := range regionPools {
				if !block.Region.IsNull() && !block.Region.IsUnknown() {
					key := block.Region.ValueString()
					var poolIDs []types.String
					block.PoolIds.ElementsAs(ctx, &poolIDs, false)
					// Convert []types.String to []attr.Value
					poolIDValues := make([]attr.Value, len(poolIDs))
					for i, id := range poolIDs {
						poolIDValues[i] = id
					}
					pools[key] = customfield.NewListMust[types.String](ctx, poolIDValues)
				}
			}
		}
	case "pop":
		if popPools, ok := sourceBlocks.([]SourcePopPoolModel); ok && len(popPools) > 0 {
			pools = make(map[string]customfield.List[types.String])
			for _, block := range popPools {
				if !block.Pop.IsNull() && !block.Pop.IsUnknown() {
					key := block.Pop.ValueString()
					var poolIDs []types.String
					block.PoolIds.ElementsAs(ctx, &poolIDs, false)
					// Convert []types.String to []attr.Value
					poolIDValues := make([]attr.Value, len(poolIDs))
					for i, id := range poolIDs {
						poolIDValues[i] = id
					}
					pools[key] = customfield.NewListMust[types.String](ctx, poolIDValues)
				}
			}
		}
	case "country":
		if countryPools, ok := sourceBlocks.([]SourceCountryPoolModel); ok && len(countryPools) > 0 {
			pools = make(map[string]customfield.List[types.String])
			for _, block := range countryPools {
				if !block.Country.IsNull() && !block.Country.IsUnknown() {
					key := block.Country.ValueString()
					var poolIDs []types.String
					block.PoolIds.ElementsAs(ctx, &poolIDs, false)
					// Convert []types.String to []attr.Value
					poolIDValues := make([]attr.Value, len(poolIDs))
					for i, id := range poolIDs {
						poolIDValues[i] = id
					}
					pools[key] = customfield.NewListMust[types.String](ctx, poolIDValues)
				}
			}
		}
	}

	if len(pools) > 0 {
		return customfield.NewMapMust(ctx, pools)
	}
	return customfield.NullMap[customfield.List[types.String]](ctx)
}

// transformRules converts v4 rules to v5 format (with nested overrides).
func transformRules(ctx context.Context, sourceRules []SourceRulesModel) (customfield.NestedObjectList[TargetRulesModel], diag.Diagnostics) {
	var diags diag.Diagnostics
	targetRules := make([]TargetRulesModel, 0, len(sourceRules))

	for _, sourceRule := range sourceRules {
		targetRule := TargetRulesModel{
			Name:       sourceRule.Name,
			Priority:   sourceRule.Priority,
			Disabled:   sourceRule.Disabled,
			Condition:  sourceRule.Condition,
			Terminates: sourceRule.Terminates,
		}

		// Transform overrides (Array[0] → single object)
		if len(sourceRule.Overrides) > 0 {
			targetRule.Overrides = transformRulesOverrides(ctx, sourceRule.Overrides[0])
		} else {
			targetRule.Overrides = customfield.NullObject[TargetRulesOverridesModel](ctx)
		}

		// Transform fixed_response (Array[0] → single object)
		if len(sourceRule.FixedResponse) > 0 {
			fixed := &TargetFixedResponseModel{
				MessageBody: sourceRule.FixedResponse[0].MessageBody,
				StatusCode:  sourceRule.FixedResponse[0].StatusCode,
				ContentType: sourceRule.FixedResponse[0].ContentType,
				Location:    sourceRule.FixedResponse[0].Location,
			}
			targetRule.FixedResponse = fixed
		}

		targetRules = append(targetRules, targetRule)
	}

	return customfield.NewObjectListMust(ctx, targetRules), diags
}

// transformRulesOverrides converts v4 rules.overrides to v5 format.
// This mirrors the top-level transformations (same nested structure changes).
func transformRulesOverrides(ctx context.Context, source SourceRulesOverridesModel) customfield.NestedObject[TargetRulesOverridesModel] {
	target := TargetRulesOverridesModel{
		SessionAffinity:    source.SessionAffinity,
		SessionAffinityTTL: convertInt64ToFloat64Nullable(source.SessionAffinityTTL),
		TTL:                convertInt64ToFloat64Nullable(source.TTL),
		SteeringPolicy:     source.SteeringPolicy,
		FallbackPool:       source.FallbackPool,
	}

	// Convert default_pools
	if !source.DefaultPools.IsNull() && !source.DefaultPools.IsUnknown() {
		var poolIDs []types.String
		source.DefaultPools.ElementsAs(ctx, &poolIDs, false)
		target.DefaultPools = &poolIDs
	}

	// Transform session_affinity_attributes (Array[0] → single object)
	if len(source.SessionAffinityAttributes) > 0 {
		target.SessionAffinityAttributes = transformSessionAffinityAttributes(ctx, source.SessionAffinityAttributes[0])
	} else {
		target.SessionAffinityAttributes = customfield.NullObject[TargetSessionAffinityAttributesModel](ctx)
	}

	// Transform adaptive_routing (Array[0] → single object)
	if len(source.AdaptiveRouting) > 0 {
		target.AdaptiveRouting = transformAdaptiveRouting(ctx, source.AdaptiveRouting[0])
	} else {
		target.AdaptiveRouting = customfield.NullObject[TargetAdaptiveRoutingModel](ctx)
	}

	// Transform location_strategy (Array[0] → single object)
	if len(source.LocationStrategy) > 0 {
		target.LocationStrategy = transformLocationStrategy(ctx, source.LocationStrategy[0])
	} else {
		target.LocationStrategy = customfield.NullObject[TargetLocationStrategyModel](ctx)
	}

	// Transform random_steering (Array[0] → single object)
	if len(source.RandomSteering) > 0 {
		target.RandomSteering = transformRandomSteering(ctx, source.RandomSteering[0])
	} else {
		target.RandomSteering = customfield.NullObject[TargetRandomSteeringModel](ctx)
	}

	// Transform pool blocks (Set → Map)
	target.RegionPools = transformPoolsToMap(ctx, source.RegionPools, "region")
	target.POPPools = transformPoolsToMap(ctx, source.PopPools, "pop")
	target.CountryPools = transformPoolsToMap(ctx, source.CountryPools, "country")

	return customfield.NewObjectMust(ctx, &target)
}

// Helper: Convert Int64 to Float64 (nullable)
func convertInt64ToFloat64Nullable(val types.Int64) types.Float64 {
	if val.IsNull() {
		return types.Float64Null()
	}
	if val.IsUnknown() {
		return types.Float64Unknown()
	}
	return types.Float64Value(float64(val.ValueInt64()))
}
