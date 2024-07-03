// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LoadBalancersResultListDataSourceEnvelope struct {
	Result *[]*LoadBalancersItemsDataSourceModel `json:"result,computed"`
}

type LoadBalancersDataSourceModel struct {
	ZoneID   types.String                          `tfsdk:"zone_id" path:"zone_id"`
	MaxItems types.Int64                           `tfsdk:"max_items"`
	Items    *[]*LoadBalancersItemsDataSourceModel `tfsdk:"items"`
}

type LoadBalancersItemsDataSourceModel struct {
	ID                 types.String                               `tfsdk:"id" json:"id,computed"`
	CountryPools       types.String                               `tfsdk:"country_pools" json:"country_pools,computed"`
	CreatedOn          types.String                               `tfsdk:"created_on" json:"created_on,computed"`
	DefaultPools       *[]types.String                            `tfsdk:"default_pools" json:"default_pools,computed"`
	Description        types.String                               `tfsdk:"description" json:"description,computed"`
	Enabled            types.Bool                                 `tfsdk:"enabled" json:"enabled,computed"`
	FallbackPool       types.String                               `tfsdk:"fallback_pool" json:"fallback_pool,computed"`
	ModifiedOn         types.String                               `tfsdk:"modified_on" json:"modified_on,computed"`
	Name               types.String                               `tfsdk:"name" json:"name,computed"`
	PopPools           types.String                               `tfsdk:"pop_pools" json:"pop_pools,computed"`
	Proxied            types.Bool                                 `tfsdk:"proxied" json:"proxied,computed"`
	RegionPools        types.String                               `tfsdk:"region_pools" json:"region_pools,computed"`
	Rules              *[]*LoadBalancersItemsRulesDataSourceModel `tfsdk:"rules" json:"rules,computed"`
	SessionAffinity    types.String                               `tfsdk:"session_affinity" json:"session_affinity,computed"`
	SessionAffinityTTL types.Float64                              `tfsdk:"session_affinity_ttl" json:"session_affinity_ttl,computed"`
	SteeringPolicy     types.String                               `tfsdk:"steering_policy" json:"steering_policy,computed"`
	TTL                types.Float64                              `tfsdk:"ttl" json:"ttl,computed"`
}

type LoadBalancersItemsAdaptiveRoutingDataSourceModel struct {
	FailoverAcrossPools types.Bool `tfsdk:"failover_across_pools" json:"failover_across_pools,computed"`
}

type LoadBalancersItemsLocationStrategyDataSourceModel struct {
	Mode      types.String `tfsdk:"mode" json:"mode,computed"`
	PreferEcs types.String `tfsdk:"prefer_ecs" json:"prefer_ecs,computed"`
}

type LoadBalancersItemsRandomSteeringDataSourceModel struct {
	DefaultWeight types.Float64 `tfsdk:"default_weight" json:"default_weight,computed"`
	PoolWeights   types.String  `tfsdk:"pool_weights" json:"pool_weights,computed"`
}

type LoadBalancersItemsRulesDataSourceModel struct {
	Condition  types.String `tfsdk:"condition" json:"condition,computed"`
	Disabled   types.Bool   `tfsdk:"disabled" json:"disabled,computed"`
	Name       types.String `tfsdk:"name" json:"name,computed"`
	Priority   types.Int64  `tfsdk:"priority" json:"priority,computed"`
	Terminates types.Bool   `tfsdk:"terminates" json:"terminates,computed"`
}

type LoadBalancersItemsRulesFixedResponseDataSourceModel struct {
	ContentType types.String `tfsdk:"content_type" json:"content_type,computed"`
	Location    types.String `tfsdk:"location" json:"location,computed"`
	MessageBody types.String `tfsdk:"message_body" json:"message_body,computed"`
	StatusCode  types.Int64  `tfsdk:"status_code" json:"status_code,computed"`
}

type LoadBalancersItemsRulesOverridesDataSourceModel struct {
	CountryPools       types.String    `tfsdk:"country_pools" json:"country_pools,computed"`
	DefaultPools       *[]types.String `tfsdk:"default_pools" json:"default_pools,computed"`
	FallbackPool       types.String    `tfsdk:"fallback_pool" json:"fallback_pool,computed"`
	PopPools           types.String    `tfsdk:"pop_pools" json:"pop_pools,computed"`
	RegionPools        types.String    `tfsdk:"region_pools" json:"region_pools,computed"`
	SessionAffinity    types.String    `tfsdk:"session_affinity" json:"session_affinity,computed"`
	SessionAffinityTTL types.Float64   `tfsdk:"session_affinity_ttl" json:"session_affinity_ttl,computed"`
	SteeringPolicy     types.String    `tfsdk:"steering_policy" json:"steering_policy,computed"`
	TTL                types.Float64   `tfsdk:"ttl" json:"ttl,computed"`
}

type LoadBalancersItemsRulesOverridesAdaptiveRoutingDataSourceModel struct {
	FailoverAcrossPools types.Bool `tfsdk:"failover_across_pools" json:"failover_across_pools,computed"`
}

type LoadBalancersItemsRulesOverridesLocationStrategyDataSourceModel struct {
	Mode      types.String `tfsdk:"mode" json:"mode,computed"`
	PreferEcs types.String `tfsdk:"prefer_ecs" json:"prefer_ecs,computed"`
}

type LoadBalancersItemsRulesOverridesRandomSteeringDataSourceModel struct {
	DefaultWeight types.Float64 `tfsdk:"default_weight" json:"default_weight,computed"`
	PoolWeights   types.String  `tfsdk:"pool_weights" json:"pool_weights,computed"`
}

type LoadBalancersItemsRulesOverridesSessionAffinityAttributesDataSourceModel struct {
	DrainDuration        types.Float64   `tfsdk:"drain_duration" json:"drain_duration,computed"`
	Headers              *[]types.String `tfsdk:"headers" json:"headers,computed"`
	RequireAllHeaders    types.Bool      `tfsdk:"require_all_headers" json:"require_all_headers,computed"`
	Samesite             types.String    `tfsdk:"samesite" json:"samesite,computed"`
	Secure               types.String    `tfsdk:"secure" json:"secure,computed"`
	ZeroDowntimeFailover types.String    `tfsdk:"zero_downtime_failover" json:"zero_downtime_failover,computed"`
}

type LoadBalancersItemsSessionAffinityAttributesDataSourceModel struct {
	DrainDuration        types.Float64   `tfsdk:"drain_duration" json:"drain_duration,computed"`
	Headers              *[]types.String `tfsdk:"headers" json:"headers,computed"`
	RequireAllHeaders    types.Bool      `tfsdk:"require_all_headers" json:"require_all_headers,computed"`
	Samesite             types.String    `tfsdk:"samesite" json:"samesite,computed"`
	Secure               types.String    `tfsdk:"secure" json:"secure,computed"`
	ZeroDowntimeFailover types.String    `tfsdk:"zero_downtime_failover" json:"zero_downtime_failover,computed"`
}
