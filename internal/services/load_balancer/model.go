// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
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
	Networks                  *[]types.String                             `tfsdk:"networks" json:"networks"`
	PopPools                  map[string]*[]types.String                  `tfsdk:"pop_pools" json:"pop_pools"`
	RegionPools               map[string]*[]types.String                  `tfsdk:"region_pools" json:"region_pools"`
	AdaptiveRouting           *LoadBalancerAdaptiveRoutingModel           `tfsdk:"adaptive_routing" json:"adaptive_routing"`
	LocationStrategy          *LoadBalancerLocationStrategyModel          `tfsdk:"location_strategy" json:"location_strategy"`
	RandomSteering            *LoadBalancerRandomSteeringModel            `tfsdk:"random_steering" json:"random_steering"`
	Rules                     *[]*LoadBalancerRulesModel                  `tfsdk:"rules" json:"rules"`
	SessionAffinityAttributes *LoadBalancerSessionAffinityAttributesModel `tfsdk:"session_affinity_attributes" json:"session_affinity_attributes"`
	Enabled                   types.Bool                                  `tfsdk:"enabled" json:"enabled,computed_optional"`
	Proxied                   types.Bool                                  `tfsdk:"proxied" json:"proxied,computed_optional"`
	SessionAffinity           types.String                                `tfsdk:"session_affinity" json:"session_affinity,computed_optional"`
	SteeringPolicy            types.String                                `tfsdk:"steering_policy" json:"steering_policy,computed_optional"`
	CreatedOn                 timetypes.RFC3339                           `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ModifiedOn                timetypes.RFC3339                           `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
}

type LoadBalancerAdaptiveRoutingModel struct {
	FailoverAcrossPools types.Bool `tfsdk:"failover_across_pools" json:"failover_across_pools,computed_optional"`
}

type LoadBalancerLocationStrategyModel struct {
	Mode      types.String `tfsdk:"mode" json:"mode,computed_optional"`
	PreferECS types.String `tfsdk:"prefer_ecs" json:"prefer_ecs,computed_optional"`
}

type LoadBalancerRandomSteeringModel struct {
	DefaultWeight types.Float64                                                        `tfsdk:"default_weight" json:"default_weight,computed_optional"`
	PoolWeights   customfield.NestedObject[LoadBalancerRandomSteeringPoolWeightsModel] `tfsdk:"pool_weights" json:"pool_weights,computed_optional"`
}

type LoadBalancerRandomSteeringPoolWeightsModel struct {
	Key   types.String  `tfsdk:"key" json:"key,computed_optional"`
	Value types.Float64 `tfsdk:"value" json:"value,computed_optional"`
}

type LoadBalancerRulesModel struct {
	Condition     types.String                                                  `tfsdk:"condition" json:"condition,computed_optional"`
	Disabled      types.Bool                                                    `tfsdk:"disabled" json:"disabled,computed_optional"`
	FixedResponse customfield.NestedObject[LoadBalancerRulesFixedResponseModel] `tfsdk:"fixed_response" json:"fixed_response,computed_optional"`
	Name          types.String                                                  `tfsdk:"name" json:"name,computed_optional"`
	Overrides     customfield.NestedObject[LoadBalancerRulesOverridesModel]     `tfsdk:"overrides" json:"overrides,computed_optional"`
	Priority      types.Int64                                                   `tfsdk:"priority" json:"priority,computed_optional"`
	Terminates    types.Bool                                                    `tfsdk:"terminates" json:"terminates,computed_optional"`
}

type LoadBalancerRulesFixedResponseModel struct {
	ContentType types.String `tfsdk:"content_type" json:"content_type,computed_optional"`
	Location    types.String `tfsdk:"location" json:"location,computed_optional"`
	MessageBody types.String `tfsdk:"message_body" json:"message_body,computed_optional"`
	StatusCode  types.Int64  `tfsdk:"status_code" json:"status_code,computed_optional"`
}

type LoadBalancerRulesOverridesModel struct {
	AdaptiveRouting           customfield.NestedObject[LoadBalancerRulesOverridesAdaptiveRoutingModel]           `tfsdk:"adaptive_routing" json:"adaptive_routing,computed_optional"`
	CountryPools              map[string]types.List                                                              `tfsdk:"country_pools" json:"country_pools,computed_optional"`
	DefaultPools              types.List                                                                         `tfsdk:"default_pools" json:"default_pools,computed_optional"`
	FallbackPool              types.String                                                                       `tfsdk:"fallback_pool" json:"fallback_pool,computed_optional"`
	LocationStrategy          customfield.NestedObject[LoadBalancerRulesOverridesLocationStrategyModel]          `tfsdk:"location_strategy" json:"location_strategy,computed_optional"`
	PopPools                  map[string]types.List                                                              `tfsdk:"pop_pools" json:"pop_pools,computed_optional"`
	RandomSteering            customfield.NestedObject[LoadBalancerRulesOverridesRandomSteeringModel]            `tfsdk:"random_steering" json:"random_steering,computed_optional"`
	RegionPools               map[string]types.List                                                              `tfsdk:"region_pools" json:"region_pools,computed_optional"`
	SessionAffinity           types.String                                                                       `tfsdk:"session_affinity" json:"session_affinity,computed_optional"`
	SessionAffinityAttributes customfield.NestedObject[LoadBalancerRulesOverridesSessionAffinityAttributesModel] `tfsdk:"session_affinity_attributes" json:"session_affinity_attributes,computed_optional"`
	SessionAffinityTTL        types.Float64                                                                      `tfsdk:"session_affinity_ttl" json:"session_affinity_ttl,computed_optional"`
	SteeringPolicy            types.String                                                                       `tfsdk:"steering_policy" json:"steering_policy,computed_optional"`
	TTL                       types.Float64                                                                      `tfsdk:"ttl" json:"ttl,computed_optional"`
}

type LoadBalancerRulesOverridesAdaptiveRoutingModel struct {
	FailoverAcrossPools types.Bool `tfsdk:"failover_across_pools" json:"failover_across_pools,computed_optional"`
}

type LoadBalancerRulesOverridesLocationStrategyModel struct {
	Mode      types.String `tfsdk:"mode" json:"mode,computed_optional"`
	PreferECS types.String `tfsdk:"prefer_ecs" json:"prefer_ecs,computed_optional"`
}

type LoadBalancerRulesOverridesRandomSteeringModel struct {
	DefaultWeight types.Float64                                                                      `tfsdk:"default_weight" json:"default_weight,computed_optional"`
	PoolWeights   customfield.NestedObject[LoadBalancerRulesOverridesRandomSteeringPoolWeightsModel] `tfsdk:"pool_weights" json:"pool_weights,computed_optional"`
}

type LoadBalancerRulesOverridesRandomSteeringPoolWeightsModel struct {
	Key   types.String  `tfsdk:"key" json:"key,computed_optional"`
	Value types.Float64 `tfsdk:"value" json:"value,computed_optional"`
}

type LoadBalancerRulesOverridesSessionAffinityAttributesModel struct {
	DrainDuration        types.Float64 `tfsdk:"drain_duration" json:"drain_duration,computed_optional"`
	Headers              types.List    `tfsdk:"headers" json:"headers,computed_optional"`
	RequireAllHeaders    types.Bool    `tfsdk:"require_all_headers" json:"require_all_headers,computed_optional"`
	Samesite             types.String  `tfsdk:"samesite" json:"samesite,computed_optional"`
	Secure               types.String  `tfsdk:"secure" json:"secure,computed_optional"`
	ZeroDowntimeFailover types.String  `tfsdk:"zero_downtime_failover" json:"zero_downtime_failover,computed_optional"`
}

type LoadBalancerSessionAffinityAttributesModel struct {
	DrainDuration        types.Float64 `tfsdk:"drain_duration" json:"drain_duration,computed_optional"`
	Headers              types.List    `tfsdk:"headers" json:"headers,computed_optional"`
	RequireAllHeaders    types.Bool    `tfsdk:"require_all_headers" json:"require_all_headers,computed_optional"`
	Samesite             types.String  `tfsdk:"samesite" json:"samesite,computed_optional"`
	Secure               types.String  `tfsdk:"secure" json:"secure,computed_optional"`
	ZeroDowntimeFailover types.String  `tfsdk:"zero_downtime_failover" json:"zero_downtime_failover,computed_optional"`
}
