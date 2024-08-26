// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/load_balancers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LoadBalancerResultDataSourceEnvelope struct {
	Result LoadBalancerDataSourceModel `json:"result,computed"`
}

type LoadBalancerResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[LoadBalancerDataSourceModel] `json:"result,computed"`
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
	Description               types.String                                          `tfsdk:"description" json:"description,computed_optional"`
	FallbackPool              types.String                                          `tfsdk:"fallback_pool" json:"fallback_pool,computed_optional"`
	ID                        types.String                                          `tfsdk:"id" json:"id,computed_optional"`
	Name                      types.String                                          `tfsdk:"name" json:"name,computed_optional"`
	SessionAffinityTTL        types.Float64                                         `tfsdk:"session_affinity_ttl" json:"session_affinity_ttl,computed_optional"`
	TTL                       types.Float64                                         `tfsdk:"ttl" json:"ttl,computed_optional"`
	CountryPools              map[string]*[]types.String                            `tfsdk:"country_pools" json:"country_pools,computed_optional"`
	DefaultPools              *[]types.String                                       `tfsdk:"default_pools" json:"default_pools,computed_optional"`
	PopPools                  map[string]*[]types.String                            `tfsdk:"pop_pools" json:"pop_pools,computed_optional"`
	RegionPools               map[string]*[]types.String                            `tfsdk:"region_pools" json:"region_pools,computed_optional"`
	AdaptiveRouting           *LoadBalancerAdaptiveRoutingDataSourceModel           `tfsdk:"adaptive_routing" json:"adaptive_routing,computed_optional"`
	LocationStrategy          *LoadBalancerLocationStrategyDataSourceModel          `tfsdk:"location_strategy" json:"location_strategy,computed_optional"`
	RandomSteering            *LoadBalancerRandomSteeringDataSourceModel            `tfsdk:"random_steering" json:"random_steering,computed_optional"`
	Rules                     *[]*LoadBalancerRulesDataSourceModel                  `tfsdk:"rules" json:"rules,computed_optional"`
	SessionAffinityAttributes *LoadBalancerSessionAffinityAttributesDataSourceModel `tfsdk:"session_affinity_attributes" json:"session_affinity_attributes,computed_optional"`
	Filter                    *LoadBalancerFindOneByDataSourceModel                 `tfsdk:"filter"`
}

func (m *LoadBalancerDataSourceModel) toReadParams() (params load_balancers.LoadBalancerGetParams, diags diag.Diagnostics) {
	params = load_balancers.LoadBalancerGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

func (m *LoadBalancerDataSourceModel) toListParams() (params load_balancers.LoadBalancerListParams, diags diag.Diagnostics) {
	params = load_balancers.LoadBalancerListParams{
		ZoneID: cloudflare.F(m.Filter.ZoneID.ValueString()),
	}

	return
}

type LoadBalancerAdaptiveRoutingDataSourceModel struct {
	FailoverAcrossPools types.Bool `tfsdk:"failover_across_pools" json:"failover_across_pools,computed"`
}

type LoadBalancerLocationStrategyDataSourceModel struct {
	Mode      types.String `tfsdk:"mode" json:"mode,computed"`
	PreferECS types.String `tfsdk:"prefer_ecs" json:"prefer_ecs,computed"`
}

type LoadBalancerRandomSteeringDataSourceModel struct {
	DefaultWeight types.Float64                                         `tfsdk:"default_weight" json:"default_weight,computed"`
	PoolWeights   *LoadBalancerRandomSteeringPoolWeightsDataSourceModel `tfsdk:"pool_weights" json:"pool_weights,computed_optional"`
}

type LoadBalancerRandomSteeringPoolWeightsDataSourceModel struct {
	Key   types.String  `tfsdk:"key" json:"key,computed_optional"`
	Value types.Float64 `tfsdk:"value" json:"value,computed_optional"`
}

type LoadBalancerRulesDataSourceModel struct {
	Condition     types.String                                   `tfsdk:"condition" json:"condition,computed_optional"`
	Disabled      types.Bool                                     `tfsdk:"disabled" json:"disabled,computed"`
	FixedResponse *LoadBalancerRulesFixedResponseDataSourceModel `tfsdk:"fixed_response" json:"fixed_response,computed_optional"`
	Name          types.String                                   `tfsdk:"name" json:"name,computed_optional"`
	Overrides     *LoadBalancerRulesOverridesDataSourceModel     `tfsdk:"overrides" json:"overrides,computed_optional"`
	Priority      types.Int64                                    `tfsdk:"priority" json:"priority,computed"`
	Terminates    types.Bool                                     `tfsdk:"terminates" json:"terminates,computed"`
}

type LoadBalancerRulesFixedResponseDataSourceModel struct {
	ContentType types.String `tfsdk:"content_type" json:"content_type,computed_optional"`
	Location    types.String `tfsdk:"location" json:"location,computed_optional"`
	MessageBody types.String `tfsdk:"message_body" json:"message_body,computed_optional"`
	StatusCode  types.Int64  `tfsdk:"status_code" json:"status_code,computed_optional"`
}

type LoadBalancerRulesOverridesDataSourceModel struct {
	AdaptiveRouting           *LoadBalancerRulesOverridesAdaptiveRoutingDataSourceModel           `tfsdk:"adaptive_routing" json:"adaptive_routing,computed_optional"`
	CountryPools              map[string]*[]types.String                                          `tfsdk:"country_pools" json:"country_pools,computed_optional"`
	DefaultPools              *[]types.String                                                     `tfsdk:"default_pools" json:"default_pools,computed_optional"`
	FallbackPool              types.String                                                        `tfsdk:"fallback_pool" json:"fallback_pool,computed_optional"`
	LocationStrategy          *LoadBalancerRulesOverridesLocationStrategyDataSourceModel          `tfsdk:"location_strategy" json:"location_strategy,computed_optional"`
	PopPools                  map[string]*[]types.String                                          `tfsdk:"pop_pools" json:"pop_pools,computed_optional"`
	RandomSteering            *LoadBalancerRulesOverridesRandomSteeringDataSourceModel            `tfsdk:"random_steering" json:"random_steering,computed_optional"`
	RegionPools               map[string]*[]types.String                                          `tfsdk:"region_pools" json:"region_pools,computed_optional"`
	SessionAffinity           types.String                                                        `tfsdk:"session_affinity" json:"session_affinity,computed"`
	SessionAffinityAttributes *LoadBalancerRulesOverridesSessionAffinityAttributesDataSourceModel `tfsdk:"session_affinity_attributes" json:"session_affinity_attributes,computed_optional"`
	SessionAffinityTTL        types.Float64                                                       `tfsdk:"session_affinity_ttl" json:"session_affinity_ttl,computed_optional"`
	SteeringPolicy            types.String                                                        `tfsdk:"steering_policy" json:"steering_policy,computed"`
	TTL                       types.Float64                                                       `tfsdk:"ttl" json:"ttl,computed_optional"`
}

type LoadBalancerRulesOverridesAdaptiveRoutingDataSourceModel struct {
	FailoverAcrossPools types.Bool `tfsdk:"failover_across_pools" json:"failover_across_pools,computed"`
}

type LoadBalancerRulesOverridesLocationStrategyDataSourceModel struct {
	Mode      types.String `tfsdk:"mode" json:"mode,computed"`
	PreferECS types.String `tfsdk:"prefer_ecs" json:"prefer_ecs,computed"`
}

type LoadBalancerRulesOverridesRandomSteeringDataSourceModel struct {
	DefaultWeight types.Float64                                                       `tfsdk:"default_weight" json:"default_weight,computed"`
	PoolWeights   *LoadBalancerRulesOverridesRandomSteeringPoolWeightsDataSourceModel `tfsdk:"pool_weights" json:"pool_weights,computed_optional"`
}

type LoadBalancerRulesOverridesRandomSteeringPoolWeightsDataSourceModel struct {
	Key   types.String  `tfsdk:"key" json:"key,computed_optional"`
	Value types.Float64 `tfsdk:"value" json:"value,computed_optional"`
}

type LoadBalancerRulesOverridesSessionAffinityAttributesDataSourceModel struct {
	DrainDuration        types.Float64   `tfsdk:"drain_duration" json:"drain_duration,computed_optional"`
	Headers              *[]types.String `tfsdk:"headers" json:"headers,computed_optional"`
	RequireAllHeaders    types.Bool      `tfsdk:"require_all_headers" json:"require_all_headers,computed"`
	Samesite             types.String    `tfsdk:"samesite" json:"samesite,computed"`
	Secure               types.String    `tfsdk:"secure" json:"secure,computed"`
	ZeroDowntimeFailover types.String    `tfsdk:"zero_downtime_failover" json:"zero_downtime_failover,computed"`
}

type LoadBalancerSessionAffinityAttributesDataSourceModel struct {
	DrainDuration        types.Float64   `tfsdk:"drain_duration" json:"drain_duration,computed_optional"`
	Headers              *[]types.String `tfsdk:"headers" json:"headers,computed_optional"`
	RequireAllHeaders    types.Bool      `tfsdk:"require_all_headers" json:"require_all_headers,computed"`
	Samesite             types.String    `tfsdk:"samesite" json:"samesite,computed"`
	Secure               types.String    `tfsdk:"secure" json:"secure,computed"`
	ZeroDowntimeFailover types.String    `tfsdk:"zero_downtime_failover" json:"zero_downtime_failover,computed"`
}

type LoadBalancerFindOneByDataSourceModel struct {
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id"`
}
