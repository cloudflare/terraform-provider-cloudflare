// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/load_balancers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LoadBalancersResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[LoadBalancersResultDataSourceModel] `json:"result,computed"`
}

type LoadBalancersDataSourceModel struct {
	ZoneID   types.String                                                     `tfsdk:"zone_id" path:"zone_id,required"`
	MaxItems types.Int64                                                      `tfsdk:"max_items"`
	Result   customfield.NestedObjectList[LoadBalancersResultDataSourceModel] `tfsdk:"result"`
}

func (m *LoadBalancersDataSourceModel) toListParams(_ context.Context) (params load_balancers.LoadBalancerListParams, diags diag.Diagnostics) {
	params = load_balancers.LoadBalancerListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type LoadBalancersResultDataSourceModel struct {
	ID                        types.String                                                                    `tfsdk:"id" json:"id,computed"`
	AdaptiveRouting           customfield.NestedObject[LoadBalancersAdaptiveRoutingDataSourceModel]           `tfsdk:"adaptive_routing" json:"adaptive_routing,computed"`
	CountryPools              customfield.Map[customfield.List[types.String]]                                 `tfsdk:"country_pools" json:"country_pools,computed"`
	CreatedOn                 types.String                                                                    `tfsdk:"created_on" json:"created_on,computed"`
	DefaultPools              customfield.List[types.String]                                                  `tfsdk:"default_pools" json:"default_pools,computed"`
	Description               types.String                                                                    `tfsdk:"description" json:"description,computed"`
	Enabled                   types.Bool                                                                      `tfsdk:"enabled" json:"enabled,computed"`
	FallbackPool              types.String                                                                    `tfsdk:"fallback_pool" json:"fallback_pool,computed"`
	LocationStrategy          customfield.NestedObject[LoadBalancersLocationStrategyDataSourceModel]          `tfsdk:"location_strategy" json:"location_strategy,computed"`
	ModifiedOn                types.String                                                                    `tfsdk:"modified_on" json:"modified_on,computed"`
	Name                      types.String                                                                    `tfsdk:"name" json:"name,computed"`
	Networks                  customfield.List[types.String]                                                  `tfsdk:"networks" json:"networks,computed"`
	POPPools                  customfield.Map[customfield.List[types.String]]                                 `tfsdk:"pop_pools" json:"pop_pools,computed"`
	Proxied                   types.Bool                                                                      `tfsdk:"proxied" json:"proxied,computed"`
	RandomSteering            customfield.NestedObject[LoadBalancersRandomSteeringDataSourceModel]            `tfsdk:"random_steering" json:"random_steering,computed"`
	RegionPools               customfield.Map[customfield.List[types.String]]                                 `tfsdk:"region_pools" json:"region_pools,computed"`
	Rules                     customfield.NestedObjectList[LoadBalancersRulesDataSourceModel]                 `tfsdk:"rules" json:"rules,computed"`
	SessionAffinity           types.String                                                                    `tfsdk:"session_affinity" json:"session_affinity,computed"`
	SessionAffinityAttributes customfield.NestedObject[LoadBalancersSessionAffinityAttributesDataSourceModel] `tfsdk:"session_affinity_attributes" json:"session_affinity_attributes,computed"`
	SessionAffinityTTL        types.Float64                                                                   `tfsdk:"session_affinity_ttl" json:"session_affinity_ttl,computed"`
	SteeringPolicy            types.String                                                                    `tfsdk:"steering_policy" json:"steering_policy,computed"`
	TTL                       types.Float64                                                                   `tfsdk:"ttl" json:"ttl,computed"`
}

type LoadBalancersAdaptiveRoutingDataSourceModel struct {
	FailoverAcrossPools types.Bool `tfsdk:"failover_across_pools" json:"failover_across_pools,computed"`
}

type LoadBalancersLocationStrategyDataSourceModel struct {
	Mode      types.String `tfsdk:"mode" json:"mode,computed"`
	PreferECS types.String `tfsdk:"prefer_ecs" json:"prefer_ecs,computed"`
}

type LoadBalancersRandomSteeringDataSourceModel struct {
	DefaultWeight types.Float64                  `tfsdk:"default_weight" json:"default_weight,computed"`
	PoolWeights   customfield.Map[types.Float64] `tfsdk:"pool_weights" json:"pool_weights,computed"`
}

type LoadBalancersRulesDataSourceModel struct {
	Condition     types.String                                                             `tfsdk:"condition" json:"condition,computed"`
	Disabled      types.Bool                                                               `tfsdk:"disabled" json:"disabled,computed"`
	FixedResponse customfield.NestedObject[LoadBalancersRulesFixedResponseDataSourceModel] `tfsdk:"fixed_response" json:"fixed_response,computed"`
	Name          types.String                                                             `tfsdk:"name" json:"name,computed"`
	Overrides     customfield.NestedObject[LoadBalancersRulesOverridesDataSourceModel]     `tfsdk:"overrides" json:"overrides,computed"`
	Priority      types.Int64                                                              `tfsdk:"priority" json:"priority,computed"`
	Terminates    types.Bool                                                               `tfsdk:"terminates" json:"terminates,computed"`
}

type LoadBalancersRulesFixedResponseDataSourceModel struct {
	ContentType types.String `tfsdk:"content_type" json:"content_type,computed"`
	Location    types.String `tfsdk:"location" json:"location,computed"`
	MessageBody types.String `tfsdk:"message_body" json:"message_body,computed"`
	StatusCode  types.Int64  `tfsdk:"status_code" json:"status_code,computed"`
}

type LoadBalancersRulesOverridesDataSourceModel struct {
	AdaptiveRouting           customfield.NestedObject[LoadBalancersRulesOverridesAdaptiveRoutingDataSourceModel]           `tfsdk:"adaptive_routing" json:"adaptive_routing,computed"`
	CountryPools              customfield.Map[customfield.List[types.String]]                                               `tfsdk:"country_pools" json:"country_pools,computed"`
	DefaultPools              customfield.List[types.String]                                                                `tfsdk:"default_pools" json:"default_pools,computed"`
	FallbackPool              types.String                                                                                  `tfsdk:"fallback_pool" json:"fallback_pool,computed"`
	LocationStrategy          customfield.NestedObject[LoadBalancersRulesOverridesLocationStrategyDataSourceModel]          `tfsdk:"location_strategy" json:"location_strategy,computed"`
	POPPools                  customfield.Map[customfield.List[types.String]]                                               `tfsdk:"pop_pools" json:"pop_pools,computed"`
	RandomSteering            customfield.NestedObject[LoadBalancersRulesOverridesRandomSteeringDataSourceModel]            `tfsdk:"random_steering" json:"random_steering,computed"`
	RegionPools               customfield.Map[customfield.List[types.String]]                                               `tfsdk:"region_pools" json:"region_pools,computed"`
	SessionAffinity           types.String                                                                                  `tfsdk:"session_affinity" json:"session_affinity,computed"`
	SessionAffinityAttributes customfield.NestedObject[LoadBalancersRulesOverridesSessionAffinityAttributesDataSourceModel] `tfsdk:"session_affinity_attributes" json:"session_affinity_attributes,computed"`
	SessionAffinityTTL        types.Float64                                                                                 `tfsdk:"session_affinity_ttl" json:"session_affinity_ttl,computed"`
	SteeringPolicy            types.String                                                                                  `tfsdk:"steering_policy" json:"steering_policy,computed"`
	TTL                       types.Float64                                                                                 `tfsdk:"ttl" json:"ttl,computed"`
}

type LoadBalancersRulesOverridesAdaptiveRoutingDataSourceModel struct {
	FailoverAcrossPools types.Bool `tfsdk:"failover_across_pools" json:"failover_across_pools,computed"`
}

type LoadBalancersRulesOverridesLocationStrategyDataSourceModel struct {
	Mode      types.String `tfsdk:"mode" json:"mode,computed"`
	PreferECS types.String `tfsdk:"prefer_ecs" json:"prefer_ecs,computed"`
}

type LoadBalancersRulesOverridesRandomSteeringDataSourceModel struct {
	DefaultWeight types.Float64                  `tfsdk:"default_weight" json:"default_weight,computed"`
	PoolWeights   customfield.Map[types.Float64] `tfsdk:"pool_weights" json:"pool_weights,computed"`
}

type LoadBalancersRulesOverridesSessionAffinityAttributesDataSourceModel struct {
	DrainDuration        types.Float64                  `tfsdk:"drain_duration" json:"drain_duration,computed"`
	Headers              customfield.List[types.String] `tfsdk:"headers" json:"headers,computed"`
	RequireAllHeaders    types.Bool                     `tfsdk:"require_all_headers" json:"require_all_headers,computed"`
	Samesite             types.String                   `tfsdk:"samesite" json:"samesite,computed"`
	Secure               types.String                   `tfsdk:"secure" json:"secure,computed"`
	ZeroDowntimeFailover types.String                   `tfsdk:"zero_downtime_failover" json:"zero_downtime_failover,computed"`
}

type LoadBalancersSessionAffinityAttributesDataSourceModel struct {
	DrainDuration        types.Float64                  `tfsdk:"drain_duration" json:"drain_duration,computed"`
	Headers              customfield.List[types.String] `tfsdk:"headers" json:"headers,computed"`
	RequireAllHeaders    types.Bool                     `tfsdk:"require_all_headers" json:"require_all_headers,computed"`
	Samesite             types.String                   `tfsdk:"samesite" json:"samesite,computed"`
	Secure               types.String                   `tfsdk:"secure" json:"secure,computed"`
	ZeroDowntimeFailover types.String                   `tfsdk:"zero_downtime_failover" json:"zero_downtime_failover,computed"`
}
