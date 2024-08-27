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

type LoadBalancersResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[LoadBalancersResultDataSourceModel] `json:"result,computed"`
}

type LoadBalancersDataSourceModel struct {
	ZoneID   types.String                                                     `tfsdk:"zone_id" path:"zone_id"`
	MaxItems types.Int64                                                      `tfsdk:"max_items"`
	Result   customfield.NestedObjectList[LoadBalancersResultDataSourceModel] `tfsdk:"result"`
}

func (m *LoadBalancersDataSourceModel) toListParams() (params load_balancers.LoadBalancerListParams, diags diag.Diagnostics) {
	params = load_balancers.LoadBalancerListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type LoadBalancersResultDataSourceModel struct {
	ID                        types.String                                           `tfsdk:"id" json:"id,computed_optional"`
	AdaptiveRouting           *LoadBalancersAdaptiveRoutingDataSourceModel           `tfsdk:"adaptive_routing" json:"adaptive_routing,computed_optional"`
	CountryPools              map[string]*[]types.String                             `tfsdk:"country_pools" json:"country_pools,computed_optional"`
	CreatedOn                 timetypes.RFC3339                                      `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	DefaultPools              *[]types.String                                        `tfsdk:"default_pools" json:"default_pools,computed_optional"`
	Description               types.String                                           `tfsdk:"description" json:"description,computed_optional"`
	Enabled                   types.Bool                                             `tfsdk:"enabled" json:"enabled,computed"`
	FallbackPool              types.String                                           `tfsdk:"fallback_pool" json:"fallback_pool,computed_optional"`
	LocationStrategy          *LoadBalancersLocationStrategyDataSourceModel          `tfsdk:"location_strategy" json:"location_strategy,computed_optional"`
	ModifiedOn                timetypes.RFC3339                                      `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Name                      types.String                                           `tfsdk:"name" json:"name,computed_optional"`
	PopPools                  map[string]*[]types.String                             `tfsdk:"pop_pools" json:"pop_pools,computed_optional"`
	Proxied                   types.Bool                                             `tfsdk:"proxied" json:"proxied,computed"`
	RandomSteering            *LoadBalancersRandomSteeringDataSourceModel            `tfsdk:"random_steering" json:"random_steering,computed_optional"`
	RegionPools               map[string]*[]types.String                             `tfsdk:"region_pools" json:"region_pools,computed_optional"`
	Rules                     *[]*LoadBalancersRulesDataSourceModel                  `tfsdk:"rules" json:"rules,computed_optional"`
	SessionAffinity           types.String                                           `tfsdk:"session_affinity" json:"session_affinity,computed"`
	SessionAffinityAttributes *LoadBalancersSessionAffinityAttributesDataSourceModel `tfsdk:"session_affinity_attributes" json:"session_affinity_attributes,computed_optional"`
	SessionAffinityTTL        types.Float64                                          `tfsdk:"session_affinity_ttl" json:"session_affinity_ttl,computed_optional"`
	SteeringPolicy            types.String                                           `tfsdk:"steering_policy" json:"steering_policy,computed"`
	TTL                       types.Float64                                          `tfsdk:"ttl" json:"ttl,computed_optional"`
}

type LoadBalancersAdaptiveRoutingDataSourceModel struct {
	FailoverAcrossPools types.Bool `tfsdk:"failover_across_pools" json:"failover_across_pools,computed"`
}

type LoadBalancersLocationStrategyDataSourceModel struct {
	Mode      types.String `tfsdk:"mode" json:"mode,computed"`
	PreferECS types.String `tfsdk:"prefer_ecs" json:"prefer_ecs,computed"`
}

type LoadBalancersRandomSteeringDataSourceModel struct {
	DefaultWeight types.Float64                                          `tfsdk:"default_weight" json:"default_weight,computed"`
	PoolWeights   *LoadBalancersRandomSteeringPoolWeightsDataSourceModel `tfsdk:"pool_weights" json:"pool_weights,computed_optional"`
}

type LoadBalancersRandomSteeringPoolWeightsDataSourceModel struct {
	Key   types.String  `tfsdk:"key" json:"key,computed_optional"`
	Value types.Float64 `tfsdk:"value" json:"value,computed_optional"`
}

type LoadBalancersRulesDataSourceModel struct {
	Condition     types.String                                    `tfsdk:"condition" json:"condition,computed_optional"`
	Disabled      types.Bool                                      `tfsdk:"disabled" json:"disabled,computed"`
	FixedResponse *LoadBalancersRulesFixedResponseDataSourceModel `tfsdk:"fixed_response" json:"fixed_response,computed_optional"`
	Name          types.String                                    `tfsdk:"name" json:"name,computed_optional"`
	Overrides     *LoadBalancersRulesOverridesDataSourceModel     `tfsdk:"overrides" json:"overrides,computed_optional"`
	Priority      types.Int64                                     `tfsdk:"priority" json:"priority,computed"`
	Terminates    types.Bool                                      `tfsdk:"terminates" json:"terminates,computed"`
}

type LoadBalancersRulesFixedResponseDataSourceModel struct {
	ContentType types.String `tfsdk:"content_type" json:"content_type,computed_optional"`
	Location    types.String `tfsdk:"location" json:"location,computed_optional"`
	MessageBody types.String `tfsdk:"message_body" json:"message_body,computed_optional"`
	StatusCode  types.Int64  `tfsdk:"status_code" json:"status_code,computed_optional"`
}

type LoadBalancersRulesOverridesDataSourceModel struct {
	AdaptiveRouting           *LoadBalancersRulesOverridesAdaptiveRoutingDataSourceModel           `tfsdk:"adaptive_routing" json:"adaptive_routing,computed_optional"`
	CountryPools              map[string]*[]types.String                                           `tfsdk:"country_pools" json:"country_pools,computed_optional"`
	DefaultPools              *[]types.String                                                      `tfsdk:"default_pools" json:"default_pools,computed_optional"`
	FallbackPool              types.String                                                         `tfsdk:"fallback_pool" json:"fallback_pool,computed_optional"`
	LocationStrategy          *LoadBalancersRulesOverridesLocationStrategyDataSourceModel          `tfsdk:"location_strategy" json:"location_strategy,computed_optional"`
	PopPools                  map[string]*[]types.String                                           `tfsdk:"pop_pools" json:"pop_pools,computed_optional"`
	RandomSteering            *LoadBalancersRulesOverridesRandomSteeringDataSourceModel            `tfsdk:"random_steering" json:"random_steering,computed_optional"`
	RegionPools               map[string]*[]types.String                                           `tfsdk:"region_pools" json:"region_pools,computed_optional"`
	SessionAffinity           types.String                                                         `tfsdk:"session_affinity" json:"session_affinity,computed"`
	SessionAffinityAttributes *LoadBalancersRulesOverridesSessionAffinityAttributesDataSourceModel `tfsdk:"session_affinity_attributes" json:"session_affinity_attributes,computed_optional"`
	SessionAffinityTTL        types.Float64                                                        `tfsdk:"session_affinity_ttl" json:"session_affinity_ttl,computed_optional"`
	SteeringPolicy            types.String                                                         `tfsdk:"steering_policy" json:"steering_policy,computed"`
	TTL                       types.Float64                                                        `tfsdk:"ttl" json:"ttl,computed_optional"`
}

type LoadBalancersRulesOverridesAdaptiveRoutingDataSourceModel struct {
	FailoverAcrossPools types.Bool `tfsdk:"failover_across_pools" json:"failover_across_pools,computed"`
}

type LoadBalancersRulesOverridesLocationStrategyDataSourceModel struct {
	Mode      types.String `tfsdk:"mode" json:"mode,computed"`
	PreferECS types.String `tfsdk:"prefer_ecs" json:"prefer_ecs,computed"`
}

type LoadBalancersRulesOverridesRandomSteeringDataSourceModel struct {
	DefaultWeight types.Float64                                                        `tfsdk:"default_weight" json:"default_weight,computed"`
	PoolWeights   *LoadBalancersRulesOverridesRandomSteeringPoolWeightsDataSourceModel `tfsdk:"pool_weights" json:"pool_weights,computed_optional"`
}

type LoadBalancersRulesOverridesRandomSteeringPoolWeightsDataSourceModel struct {
	Key   types.String  `tfsdk:"key" json:"key,computed_optional"`
	Value types.Float64 `tfsdk:"value" json:"value,computed_optional"`
}

type LoadBalancersRulesOverridesSessionAffinityAttributesDataSourceModel struct {
	DrainDuration        types.Float64   `tfsdk:"drain_duration" json:"drain_duration,computed_optional"`
	Headers              *[]types.String `tfsdk:"headers" json:"headers,computed_optional"`
	RequireAllHeaders    types.Bool      `tfsdk:"require_all_headers" json:"require_all_headers,computed"`
	Samesite             types.String    `tfsdk:"samesite" json:"samesite,computed"`
	Secure               types.String    `tfsdk:"secure" json:"secure,computed"`
	ZeroDowntimeFailover types.String    `tfsdk:"zero_downtime_failover" json:"zero_downtime_failover,computed"`
}

type LoadBalancersSessionAffinityAttributesDataSourceModel struct {
	DrainDuration        types.Float64   `tfsdk:"drain_duration" json:"drain_duration,computed_optional"`
	Headers              *[]types.String `tfsdk:"headers" json:"headers,computed_optional"`
	RequireAllHeaders    types.Bool      `tfsdk:"require_all_headers" json:"require_all_headers,computed"`
	Samesite             types.String    `tfsdk:"samesite" json:"samesite,computed"`
	Secure               types.String    `tfsdk:"secure" json:"secure,computed"`
	ZeroDowntimeFailover types.String    `tfsdk:"zero_downtime_failover" json:"zero_downtime_failover,computed"`
}
