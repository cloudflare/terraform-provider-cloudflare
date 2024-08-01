// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LoadBalancerResultDataSourceEnvelope struct {
	Result LoadBalancerDataSourceModel `json:"result,computed"`
}

type LoadBalancerResultListDataSourceEnvelope struct {
	Result *[]*LoadBalancerDataSourceModel `json:"result,computed"`
}

type LoadBalancerDataSourceModel struct {
	LoadBalancerID            types.String                                          `tfsdk:"load_balancer_id" path:"load_balancer_id"`
	ZoneID                    types.String                                          `tfsdk:"zone_id" path:"zone_id"`
	CreatedOn                 timetypes.RFC3339                                     `tfsdk:"created_on" json:"created_on,computed"`
	Enabled                   types.Bool                                            `tfsdk:"enabled" json:"enabled,computed"`
	ModifiedOn                timetypes.RFC3339                                     `tfsdk:"modified_on" json:"modified_on,computed"`
	Proxied                   types.Bool                                            `tfsdk:"proxied" json:"proxied,computed"`
	SessionAffinity           types.String                                          `tfsdk:"session_affinity" json:"session_affinity,computed"`
	SteeringPolicy            types.String                                          `tfsdk:"steering_policy" json:"steering_policy,computed"`
	Description               types.String                                          `tfsdk:"description" json:"description"`
	ID                        types.String                                          `tfsdk:"id" json:"id"`
	Name                      types.String                                          `tfsdk:"name" json:"name"`
	SessionAffinityTTL        types.Float64                                         `tfsdk:"session_affinity_ttl" json:"session_affinity_ttl"`
	TTL                       types.Float64                                         `tfsdk:"ttl" json:"ttl"`
	DefaultPools              *[]types.String                                       `tfsdk:"default_pools" json:"default_pools"`
	AdaptiveRouting           *LoadBalancerAdaptiveRoutingDataSourceModel           `tfsdk:"adaptive_routing" json:"adaptive_routing"`
	LocationStrategy          *LoadBalancerLocationStrategyDataSourceModel          `tfsdk:"location_strategy" json:"location_strategy"`
	RandomSteering            *LoadBalancerRandomSteeringDataSourceModel            `tfsdk:"random_steering" json:"random_steering"`
	Rules                     *[]*LoadBalancerRulesDataSourceModel                  `tfsdk:"rules" json:"rules"`
	SessionAffinityAttributes *LoadBalancerSessionAffinityAttributesDataSourceModel `tfsdk:"session_affinity_attributes" json:"session_affinity_attributes"`
	CountryPools              jsontypes.Normalized                                  `tfsdk:"country_pools" json:"country_pools"`
	FallbackPool              jsontypes.Normalized                                  `tfsdk:"fallback_pool" json:"fallback_pool"`
	PopPools                  jsontypes.Normalized                                  `tfsdk:"pop_pools" json:"pop_pools"`
	RegionPools               jsontypes.Normalized                                  `tfsdk:"region_pools" json:"region_pools"`
	Filter                    *LoadBalancerFindOneByDataSourceModel                 `tfsdk:"filter"`
}

type LoadBalancerAdaptiveRoutingDataSourceModel struct {
	FailoverAcrossPools types.Bool `tfsdk:"failover_across_pools" json:"failover_across_pools,computed"`
}

type LoadBalancerLocationStrategyDataSourceModel struct {
	Mode      types.String `tfsdk:"mode" json:"mode,computed"`
	PreferECS types.String `tfsdk:"prefer_ecs" json:"prefer_ecs,computed"`
}

type LoadBalancerRandomSteeringDataSourceModel struct {
	DefaultWeight types.Float64        `tfsdk:"default_weight" json:"default_weight,computed"`
	PoolWeights   jsontypes.Normalized `tfsdk:"pool_weights" json:"pool_weights"`
}

type LoadBalancerRulesDataSourceModel struct {
	Condition     types.String                                   `tfsdk:"condition" json:"condition"`
	Disabled      types.Bool                                     `tfsdk:"disabled" json:"disabled,computed"`
	FixedResponse *LoadBalancerRulesFixedResponseDataSourceModel `tfsdk:"fixed_response" json:"fixed_response"`
	Name          types.String                                   `tfsdk:"name" json:"name"`
	Overrides     *LoadBalancerRulesOverridesDataSourceModel     `tfsdk:"overrides" json:"overrides"`
	Priority      types.Int64                                    `tfsdk:"priority" json:"priority,computed"`
	Terminates    types.Bool                                     `tfsdk:"terminates" json:"terminates,computed"`
}

type LoadBalancerRulesFixedResponseDataSourceModel struct {
	ContentType types.String `tfsdk:"content_type" json:"content_type"`
	Location    types.String `tfsdk:"location" json:"location"`
	MessageBody types.String `tfsdk:"message_body" json:"message_body"`
	StatusCode  types.Int64  `tfsdk:"status_code" json:"status_code"`
}

type LoadBalancerRulesOverridesDataSourceModel struct {
	AdaptiveRouting           *LoadBalancerRulesOverridesAdaptiveRoutingDataSourceModel           `tfsdk:"adaptive_routing" json:"adaptive_routing"`
	CountryPools              jsontypes.Normalized                                                `tfsdk:"country_pools" json:"country_pools"`
	DefaultPools              *[]types.String                                                     `tfsdk:"default_pools" json:"default_pools"`
	FallbackPool              jsontypes.Normalized                                                `tfsdk:"fallback_pool" json:"fallback_pool"`
	LocationStrategy          *LoadBalancerRulesOverridesLocationStrategyDataSourceModel          `tfsdk:"location_strategy" json:"location_strategy"`
	PopPools                  jsontypes.Normalized                                                `tfsdk:"pop_pools" json:"pop_pools"`
	RandomSteering            *LoadBalancerRulesOverridesRandomSteeringDataSourceModel            `tfsdk:"random_steering" json:"random_steering"`
	RegionPools               jsontypes.Normalized                                                `tfsdk:"region_pools" json:"region_pools"`
	SessionAffinity           types.String                                                        `tfsdk:"session_affinity" json:"session_affinity,computed"`
	SessionAffinityAttributes *LoadBalancerRulesOverridesSessionAffinityAttributesDataSourceModel `tfsdk:"session_affinity_attributes" json:"session_affinity_attributes"`
	SessionAffinityTTL        types.Float64                                                       `tfsdk:"session_affinity_ttl" json:"session_affinity_ttl"`
	SteeringPolicy            types.String                                                        `tfsdk:"steering_policy" json:"steering_policy,computed"`
	TTL                       types.Float64                                                       `tfsdk:"ttl" json:"ttl"`
}

type LoadBalancerRulesOverridesAdaptiveRoutingDataSourceModel struct {
	FailoverAcrossPools types.Bool `tfsdk:"failover_across_pools" json:"failover_across_pools,computed"`
}

type LoadBalancerRulesOverridesLocationStrategyDataSourceModel struct {
	Mode      types.String `tfsdk:"mode" json:"mode,computed"`
	PreferECS types.String `tfsdk:"prefer_ecs" json:"prefer_ecs,computed"`
}

type LoadBalancerRulesOverridesRandomSteeringDataSourceModel struct {
	DefaultWeight types.Float64        `tfsdk:"default_weight" json:"default_weight,computed"`
	PoolWeights   jsontypes.Normalized `tfsdk:"pool_weights" json:"pool_weights"`
}

type LoadBalancerRulesOverridesSessionAffinityAttributesDataSourceModel struct {
	DrainDuration        types.Float64   `tfsdk:"drain_duration" json:"drain_duration"`
	Headers              *[]types.String `tfsdk:"headers" json:"headers"`
	RequireAllHeaders    types.Bool      `tfsdk:"require_all_headers" json:"require_all_headers,computed"`
	Samesite             types.String    `tfsdk:"samesite" json:"samesite,computed"`
	Secure               types.String    `tfsdk:"secure" json:"secure,computed"`
	ZeroDowntimeFailover types.String    `tfsdk:"zero_downtime_failover" json:"zero_downtime_failover,computed"`
}

type LoadBalancerSessionAffinityAttributesDataSourceModel struct {
	DrainDuration        types.Float64   `tfsdk:"drain_duration" json:"drain_duration"`
	Headers              *[]types.String `tfsdk:"headers" json:"headers"`
	RequireAllHeaders    types.Bool      `tfsdk:"require_all_headers" json:"require_all_headers,computed"`
	Samesite             types.String    `tfsdk:"samesite" json:"samesite,computed"`
	Secure               types.String    `tfsdk:"secure" json:"secure,computed"`
	ZeroDowntimeFailover types.String    `tfsdk:"zero_downtime_failover" json:"zero_downtime_failover,computed"`
}

type LoadBalancerFindOneByDataSourceModel struct {
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id"`
}
