// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/load_balancers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
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
	LoadBalancerID            types.String                                                                   `tfsdk:"load_balancer_id" path:"load_balancer_id,optional"`
	ZoneID                    types.String                                                                   `tfsdk:"zone_id" path:"zone_id,optional"`
	CreatedOn                 types.String                                                                   `tfsdk:"created_on" json:"created_on,computed"`
	Description               types.String                                                                   `tfsdk:"description" json:"description,computed"`
	Enabled                   types.Bool                                                                     `tfsdk:"enabled" json:"enabled,computed"`
	FallbackPool              types.String                                                                   `tfsdk:"fallback_pool" json:"fallback_pool,computed"`
	ID                        types.String                                                                   `tfsdk:"id" json:"id,computed"`
	ModifiedOn                types.String                                                                   `tfsdk:"modified_on" json:"modified_on,computed"`
	Name                      types.String                                                                   `tfsdk:"name" json:"name,computed"`
	Proxied                   types.Bool                                                                     `tfsdk:"proxied" json:"proxied,computed"`
	SessionAffinity           types.String                                                                   `tfsdk:"session_affinity" json:"session_affinity,computed"`
	SessionAffinityTTL        types.Float64                                                                  `tfsdk:"session_affinity_ttl" json:"session_affinity_ttl,computed_optional"`
	SteeringPolicy            types.String                                                                   `tfsdk:"steering_policy" json:"steering_policy,computed_optional"`
	TTL                       types.Float64                                                                  `tfsdk:"ttl" json:"ttl,computed_optional"`
	CountryPools              customfield.Map[customfield.List[types.String]]                                `tfsdk:"country_pools" json:"country_pools,computed_optional"`
	DefaultPools              customfield.List[types.String]                                                 `tfsdk:"default_pools" json:"default_pools,computed"`
	Networks                  customfield.List[types.String]                                                 `tfsdk:"networks" json:"networks,computed"`
	POPPools                  customfield.Map[customfield.List[types.String]]                                `tfsdk:"pop_pools" json:"pop_pools,computed_optional"`
	RegionPools               customfield.Map[customfield.List[types.String]]                                `tfsdk:"region_pools" json:"region_pools,computed"`
	AdaptiveRouting           customfield.NestedObject[LoadBalancerAdaptiveRoutingDataSourceModel]           `tfsdk:"adaptive_routing" json:"adaptive_routing,computed"`
	LocationStrategy          customfield.NestedObject[LoadBalancerLocationStrategyDataSourceModel]          `tfsdk:"location_strategy" json:"location_strategy,computed"`
	RandomSteering            customfield.NestedObject[LoadBalancerRandomSteeringDataSourceModel]            `tfsdk:"random_steering" json:"random_steering,computed"`
	Rules                     customfield.NestedObjectList[LoadBalancerRulesDataSourceModel]                 `tfsdk:"rules" json:"rules,computed"`
	SessionAffinityAttributes customfield.NestedObject[LoadBalancerSessionAffinityAttributesDataSourceModel] `tfsdk:"session_affinity_attributes" json:"session_affinity_attributes,computed"`
	Filter                    *LoadBalancerFindOneByDataSourceModel                                          `tfsdk:"filter"`
}

func (m *LoadBalancerDataSourceModel) toReadParams(_ context.Context) (params load_balancers.LoadBalancerGetParams, diags diag.Diagnostics) {
	params = load_balancers.LoadBalancerGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

func (m *LoadBalancerDataSourceModel) toListParams(_ context.Context) (params load_balancers.LoadBalancerListParams, diags diag.Diagnostics) {
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
	DefaultWeight types.Float64                  `tfsdk:"default_weight" json:"default_weight,computed"`
	PoolWeights   customfield.Map[types.Float64] `tfsdk:"pool_weights" json:"pool_weights,computed"`
}

type LoadBalancerRulesDataSourceModel struct {
	Condition     types.String                                                            `tfsdk:"condition" json:"condition,computed"`
	Disabled      types.Bool                                                              `tfsdk:"disabled" json:"disabled,computed"`
	FixedResponse customfield.NestedObject[LoadBalancerRulesFixedResponseDataSourceModel] `tfsdk:"fixed_response" json:"fixed_response,computed"`
	Name          types.String                                                            `tfsdk:"name" json:"name,computed"`
	Overrides     customfield.NestedObject[LoadBalancerRulesOverridesDataSourceModel]     `tfsdk:"overrides" json:"overrides,computed"`
	Priority      types.Int64                                                             `tfsdk:"priority" json:"priority,computed"`
	Terminates    types.Bool                                                              `tfsdk:"terminates" json:"terminates,computed"`
}

type LoadBalancerRulesFixedResponseDataSourceModel struct {
	ContentType types.String `tfsdk:"content_type" json:"content_type,computed"`
	Location    types.String `tfsdk:"location" json:"location,computed"`
	MessageBody types.String `tfsdk:"message_body" json:"message_body,computed"`
	StatusCode  types.Int64  `tfsdk:"status_code" json:"status_code,computed"`
}

type LoadBalancerRulesOverridesDataSourceModel struct {
	AdaptiveRouting           customfield.NestedObject[LoadBalancerRulesOverridesAdaptiveRoutingDataSourceModel]           `tfsdk:"adaptive_routing" json:"adaptive_routing,computed"`
	CountryPools              customfield.Map[customfield.List[types.String]]                                              `tfsdk:"country_pools" json:"country_pools,computed_optional"`
	DefaultPools              customfield.List[types.String]                                                               `tfsdk:"default_pools" json:"default_pools,computed"`
	FallbackPool              types.String                                                                                 `tfsdk:"fallback_pool" json:"fallback_pool,computed"`
	LocationStrategy          customfield.NestedObject[LoadBalancerRulesOverridesLocationStrategyDataSourceModel]          `tfsdk:"location_strategy" json:"location_strategy,computed"`
<<<<<<< HEAD
<<<<<<< HEAD
	POPPools                  customfield.Map[customfield.List[types.String]]                                              `tfsdk:"pop_pools" json:"pop_pools,computed"`
=======
	PopPools                  customfield.Map[customfield.List[types.String]]                                              `tfsdk:"pop_pools" json:"pop_pools,optional"`
>>>>>>> c47923fde (DELETE ME: Fix incorrect optional fields)
=======
	PopPools                  customfield.Map[customfield.List[types.String]]                                              `tfsdk:"pop_pools" json:"pop_pools,computed_optional"`
>>>>>>> ba9bc8219 (DELETE ME: Fix incorrect computed_optional pop_pools)
	RandomSteering            customfield.NestedObject[LoadBalancerRulesOverridesRandomSteeringDataSourceModel]            `tfsdk:"random_steering" json:"random_steering,computed"`
	RegionPools               customfield.Map[customfield.List[types.String]]                                              `tfsdk:"region_pools" json:"region_pools,computed_optional"`
	SessionAffinity           types.String                                                                                 `tfsdk:"session_affinity" json:"session_affinity,computed"`
	SessionAffinityAttributes customfield.NestedObject[LoadBalancerRulesOverridesSessionAffinityAttributesDataSourceModel] `tfsdk:"session_affinity_attributes" json:"session_affinity_attributes,computed"`
	SessionAffinityTTL        types.Float64                                                                                `tfsdk:"session_affinity_ttl" json:"session_affinity_ttl,computed_optional"`
	SteeringPolicy            types.String                                                                                 `tfsdk:"steering_policy" json:"steering_policy,computed_optional"`
	TTL                       types.Float64                                                                                `tfsdk:"ttl" json:"ttl,computed_optional"`
}

type LoadBalancerRulesOverridesAdaptiveRoutingDataSourceModel struct {
	FailoverAcrossPools types.Bool `tfsdk:"failover_across_pools" json:"failover_across_pools,computed"`
}

type LoadBalancerRulesOverridesLocationStrategyDataSourceModel struct {
	Mode      types.String `tfsdk:"mode" json:"mode,computed"`
	PreferECS types.String `tfsdk:"prefer_ecs" json:"prefer_ecs,computed"`
}

type LoadBalancerRulesOverridesRandomSteeringDataSourceModel struct {
	DefaultWeight types.Float64                  `tfsdk:"default_weight" json:"default_weight,computed"`
	PoolWeights   customfield.Map[types.Float64] `tfsdk:"pool_weights" json:"pool_weights,computed"`
}

type LoadBalancerRulesOverridesSessionAffinityAttributesDataSourceModel struct {
	DrainDuration        types.Float64                  `tfsdk:"drain_duration" json:"drain_duration,computed_optional"`
	Headers              customfield.List[types.String] `tfsdk:"headers" json:"headers,computed"`
	RequireAllHeaders    types.Bool                     `tfsdk:"require_all_headers" json:"require_all_headers,computed"`
	Samesite             types.String                   `tfsdk:"samesite" json:"samesite,computed"`
	Secure               types.String                   `tfsdk:"secure" json:"secure,computed"`
	ZeroDowntimeFailover types.String                   `tfsdk:"zero_downtime_failover" json:"zero_downtime_failover,computed"`
}

type LoadBalancerSessionAffinityAttributesDataSourceModel struct {
	DrainDuration        types.Float64                  `tfsdk:"drain_duration" json:"drain_duration,computed_optional"`
	Headers              customfield.List[types.String] `tfsdk:"headers" json:"headers,computed"`
	RequireAllHeaders    types.Bool                     `tfsdk:"require_all_headers" json:"require_all_headers,computed"`
	Samesite             types.String                   `tfsdk:"samesite" json:"samesite,computed"`
	Secure               types.String                   `tfsdk:"secure" json:"secure,computed"`
	ZeroDowntimeFailover types.String                   `tfsdk:"zero_downtime_failover" json:"zero_downtime_failover,computed"`
}

type LoadBalancerFindOneByDataSourceModel struct {
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
}
