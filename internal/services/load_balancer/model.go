// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LoadBalancerResultEnvelope struct {
	Result LoadBalancerModel `json:"result,computed"`
}

type LoadBalancerModel struct {
	ID                        types.String                                `tfsdk:"id" json:"id,computed"`
	ZoneID                    types.String                                `tfsdk:"zone_id" path:"zone_id"`
	DefaultPools              *[]types.String                             `tfsdk:"default_pools" json:"default_pools"`
	FallbackPool              types.String                                `tfsdk:"fallback_pool" json:"fallback_pool"`
	Name                      types.String                                `tfsdk:"name" json:"name"`
	AdaptiveRouting           *LoadBalancerAdaptiveRoutingModel           `tfsdk:"adaptive_routing" json:"adaptive_routing"`
	CountryPools              types.String                                `tfsdk:"country_pools" json:"country_pools"`
	Description               types.String                                `tfsdk:"description" json:"description"`
	LocationStrategy          *LoadBalancerLocationStrategyModel          `tfsdk:"location_strategy" json:"location_strategy"`
	PopPools                  types.String                                `tfsdk:"pop_pools" json:"pop_pools"`
	Proxied                   types.Bool                                  `tfsdk:"proxied" json:"proxied"`
	RandomSteering            *LoadBalancerRandomSteeringModel            `tfsdk:"random_steering" json:"random_steering"`
	RegionPools               types.String                                `tfsdk:"region_pools" json:"region_pools"`
	Rules                     *[]*LoadBalancerRulesModel                  `tfsdk:"rules" json:"rules"`
	SessionAffinity           types.String                                `tfsdk:"session_affinity" json:"session_affinity"`
	SessionAffinityAttributes *LoadBalancerSessionAffinityAttributesModel `tfsdk:"session_affinity_attributes" json:"session_affinity_attributes"`
	SessionAffinityTTL        types.Float64                               `tfsdk:"session_affinity_ttl" json:"session_affinity_ttl"`
	SteeringPolicy            types.String                                `tfsdk:"steering_policy" json:"steering_policy"`
	TTL                       types.Float64                               `tfsdk:"ttl" json:"ttl"`
	Enabled                   types.Bool                                  `tfsdk:"enabled" json:"enabled"`
	CreatedOn                 types.String                                `tfsdk:"created_on" json:"created_on,computed"`
	ModifiedOn                types.String                                `tfsdk:"modified_on" json:"modified_on,computed"`
}

type LoadBalancerAdaptiveRoutingModel struct {
	FailoverAcrossPools types.Bool `tfsdk:"failover_across_pools" json:"failover_across_pools"`
}

type LoadBalancerLocationStrategyModel struct {
	Mode      types.String `tfsdk:"mode" json:"mode"`
	PreferEcs types.String `tfsdk:"prefer_ecs" json:"prefer_ecs"`
}

type LoadBalancerRandomSteeringModel struct {
	DefaultWeight types.Float64 `tfsdk:"default_weight" json:"default_weight"`
	PoolWeights   types.String  `tfsdk:"pool_weights" json:"pool_weights"`
}

type LoadBalancerRulesModel struct {
	Condition     types.String                         `tfsdk:"condition" json:"condition"`
	Disabled      types.Bool                           `tfsdk:"disabled" json:"disabled"`
	FixedResponse *LoadBalancerRulesFixedResponseModel `tfsdk:"fixed_response" json:"fixed_response"`
	Name          types.String                         `tfsdk:"name" json:"name"`
	Overrides     *LoadBalancerRulesOverridesModel     `tfsdk:"overrides" json:"overrides"`
	Priority      types.Int64                          `tfsdk:"priority" json:"priority"`
	Terminates    types.Bool                           `tfsdk:"terminates" json:"terminates"`
}

type LoadBalancerRulesFixedResponseModel struct {
	ContentType types.String `tfsdk:"content_type" json:"content_type"`
	Location    types.String `tfsdk:"location" json:"location"`
	MessageBody types.String `tfsdk:"message_body" json:"message_body"`
	StatusCode  types.Int64  `tfsdk:"status_code" json:"status_code"`
}

type LoadBalancerRulesOverridesModel struct {
	AdaptiveRouting           *LoadBalancerRulesOverridesAdaptiveRoutingModel           `tfsdk:"adaptive_routing" json:"adaptive_routing"`
	CountryPools              types.String                                              `tfsdk:"country_pools" json:"country_pools"`
	DefaultPools              *[]types.String                                           `tfsdk:"default_pools" json:"default_pools"`
	FallbackPool              types.String                                              `tfsdk:"fallback_pool" json:"fallback_pool"`
	LocationStrategy          *LoadBalancerRulesOverridesLocationStrategyModel          `tfsdk:"location_strategy" json:"location_strategy"`
	PopPools                  types.String                                              `tfsdk:"pop_pools" json:"pop_pools"`
	RandomSteering            *LoadBalancerRulesOverridesRandomSteeringModel            `tfsdk:"random_steering" json:"random_steering"`
	RegionPools               types.String                                              `tfsdk:"region_pools" json:"region_pools"`
	SessionAffinity           types.String                                              `tfsdk:"session_affinity" json:"session_affinity"`
	SessionAffinityAttributes *LoadBalancerRulesOverridesSessionAffinityAttributesModel `tfsdk:"session_affinity_attributes" json:"session_affinity_attributes"`
	SessionAffinityTTL        types.Float64                                             `tfsdk:"session_affinity_ttl" json:"session_affinity_ttl"`
	SteeringPolicy            types.String                                              `tfsdk:"steering_policy" json:"steering_policy"`
	TTL                       types.Float64                                             `tfsdk:"ttl" json:"ttl"`
}

type LoadBalancerRulesOverridesAdaptiveRoutingModel struct {
	FailoverAcrossPools types.Bool `tfsdk:"failover_across_pools" json:"failover_across_pools"`
}

type LoadBalancerRulesOverridesLocationStrategyModel struct {
	Mode      types.String `tfsdk:"mode" json:"mode"`
	PreferEcs types.String `tfsdk:"prefer_ecs" json:"prefer_ecs"`
}

type LoadBalancerRulesOverridesRandomSteeringModel struct {
	DefaultWeight types.Float64 `tfsdk:"default_weight" json:"default_weight"`
	PoolWeights   types.String  `tfsdk:"pool_weights" json:"pool_weights"`
}

type LoadBalancerRulesOverridesSessionAffinityAttributesModel struct {
	DrainDuration        types.Float64   `tfsdk:"drain_duration" json:"drain_duration"`
	Headers              *[]types.String `tfsdk:"headers" json:"headers"`
	RequireAllHeaders    types.Bool      `tfsdk:"require_all_headers" json:"require_all_headers"`
	Samesite             types.String    `tfsdk:"samesite" json:"samesite"`
	Secure               types.String    `tfsdk:"secure" json:"secure"`
	ZeroDowntimeFailover types.String    `tfsdk:"zero_downtime_failover" json:"zero_downtime_failover"`
}

type LoadBalancerSessionAffinityAttributesModel struct {
	DrainDuration        types.Float64   `tfsdk:"drain_duration" json:"drain_duration"`
	Headers              *[]types.String `tfsdk:"headers" json:"headers"`
	RequireAllHeaders    types.Bool      `tfsdk:"require_all_headers" json:"require_all_headers"`
	Samesite             types.String    `tfsdk:"samesite" json:"samesite"`
	Secure               types.String    `tfsdk:"secure" json:"secure"`
	ZeroDowntimeFailover types.String    `tfsdk:"zero_downtime_failover" json:"zero_downtime_failover"`
}
