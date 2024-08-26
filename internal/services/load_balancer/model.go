// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LoadBalancerResultEnvelope struct {
	Result LoadBalancerModel `json:"result"`
}

type LoadBalancerModel struct {
	ID                        types.String                                `tfsdk:"id" json:"id,computed"`
	ZoneID                    types.String                                `tfsdk:"zone_id" path:"zone_id"`
	FallbackPool              types.String                                `tfsdk:"fallback_pool" json:"fallback_pool"`
	Name                      types.String                                `tfsdk:"name" json:"name"`
	DefaultPools              *[]types.String                             `tfsdk:"default_pools" json:"default_pools"`
	Description               types.String                                `tfsdk:"description" json:"description"`
	SessionAffinityTTL        types.Float64                               `tfsdk:"session_affinity_ttl" json:"session_affinity_ttl"`
	TTL                       types.Float64                               `tfsdk:"ttl" json:"ttl"`
	CountryPools              map[string]*[]types.String                  `tfsdk:"country_pools" json:"country_pools"`
	PopPools                  map[string]*[]types.String                  `tfsdk:"pop_pools" json:"pop_pools"`
	RegionPools               map[string]*[]types.String                  `tfsdk:"region_pools" json:"region_pools"`
	AdaptiveRouting           *LoadBalancerAdaptiveRoutingModel           `tfsdk:"adaptive_routing" json:"adaptive_routing"`
	LocationStrategy          *LoadBalancerLocationStrategyModel          `tfsdk:"location_strategy" json:"location_strategy"`
	RandomSteering            *LoadBalancerRandomSteeringModel            `tfsdk:"random_steering" json:"random_steering"`
	Rules                     *[]*LoadBalancerRulesModel                  `tfsdk:"rules" json:"rules"`
	SessionAffinityAttributes *LoadBalancerSessionAffinityAttributesModel `tfsdk:"session_affinity_attributes" json:"session_affinity_attributes"`
	Enabled                   types.Bool                                  `tfsdk:"enabled" json:"enabled"`
	Proxied                   types.Bool                                  `tfsdk:"proxied" json:"proxied"`
	SessionAffinity           types.String                                `tfsdk:"session_affinity" json:"session_affinity"`
	SteeringPolicy            types.String                                `tfsdk:"steering_policy" json:"steering_policy"`
	CreatedOn                 timetypes.RFC3339                           `tfsdk:"created_on" json:"created_on,computed"`
	ModifiedOn                timetypes.RFC3339                           `tfsdk:"modified_on" json:"modified_on,computed"`
}

type LoadBalancerAdaptiveRoutingModel struct {
	FailoverAcrossPools types.Bool `tfsdk:"failover_across_pools" json:"failover_across_pools"`
}

type LoadBalancerLocationStrategyModel struct {
	Mode      types.String `tfsdk:"mode" json:"mode"`
	PreferECS types.String `tfsdk:"prefer_ecs" json:"prefer_ecs"`
}

type LoadBalancerRandomSteeringModel struct {
	DefaultWeight types.Float64                               `tfsdk:"default_weight" json:"default_weight"`
	PoolWeights   *LoadBalancerRandomSteeringPoolWeightsModel `tfsdk:"pool_weights" json:"pool_weights"`
}

type LoadBalancerRandomSteeringPoolWeightsModel struct {
	Key   types.String  `tfsdk:"key" json:"key"`
	Value types.Float64 `tfsdk:"value" json:"value"`
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
	CountryPools              map[string]*[]types.String                                `tfsdk:"country_pools" json:"country_pools"`
	DefaultPools              *[]types.String                                           `tfsdk:"default_pools" json:"default_pools"`
	FallbackPool              types.String                                              `tfsdk:"fallback_pool" json:"fallback_pool"`
	LocationStrategy          *LoadBalancerRulesOverridesLocationStrategyModel          `tfsdk:"location_strategy" json:"location_strategy"`
	PopPools                  map[string]*[]types.String                                `tfsdk:"pop_pools" json:"pop_pools"`
	RandomSteering            *LoadBalancerRulesOverridesRandomSteeringModel            `tfsdk:"random_steering" json:"random_steering"`
	RegionPools               map[string]*[]types.String                                `tfsdk:"region_pools" json:"region_pools"`
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
	PreferECS types.String `tfsdk:"prefer_ecs" json:"prefer_ecs"`
}

type LoadBalancerRulesOverridesRandomSteeringModel struct {
	DefaultWeight types.Float64                                             `tfsdk:"default_weight" json:"default_weight"`
	PoolWeights   *LoadBalancerRulesOverridesRandomSteeringPoolWeightsModel `tfsdk:"pool_weights" json:"pool_weights"`
}

type LoadBalancerRulesOverridesRandomSteeringPoolWeightsModel struct {
	Key   types.String  `tfsdk:"key" json:"key"`
	Value types.Float64 `tfsdk:"value" json:"value"`
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
