package v500

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x)
// ============================================================================

// SourceCloudflareLoadBalancerModel represents the legacy cloudflare_load_balancer state from v4.x provider.
// Schema version: 1 (SDKv2 default for resources with migrations)
// Resource type: cloudflare_load_balancer
//
// Key differences from v5:
// - default_pool_ids (v4) → default_pools (v5)
// - fallback_pool_id (v4) → fallback_pool (v5)
// - ttl as Int64 (v4) → Float64 (v5)
// - session_affinity_ttl as Int64 (v4) → Float64 (v5)
// - Nested objects stored as TypeList MaxItems:1 (arrays) in v4, single objects in v5
// - Pool fields stored as TypeSet of blocks in v4, maps in v5
type SourceCloudflareLoadBalancerModel struct {
	ID                        types.String                                `tfsdk:"id"`
	ZoneID                    types.String                                `tfsdk:"zone_id"`
	Name                      types.String                                `tfsdk:"name"`
	DefaultPoolIds            types.List                                  `tfsdk:"default_pool_ids"` // Renamed to default_pools in v5
	FallbackPoolId            types.String                                `tfsdk:"fallback_pool_id"` // Renamed to fallback_pool in v5
	Description               types.String                                `tfsdk:"description"`
	TTL                       types.Int64                                 `tfsdk:"ttl"`                        // Int64 in v4, Float64 in v5
	SessionAffinity           types.String                                `tfsdk:"session_affinity"`
	SessionAffinityTTL        types.Int64                                 `tfsdk:"session_affinity_ttl"`       // Int64 in v4, Float64 in v5
	SessionAffinityAttributes []SourceSessionAffinityAttributesModel      `tfsdk:"session_affinity_attributes"` // Array in v4, single object in v5
	Proxied                   types.Bool                                  `tfsdk:"proxied"`
	Enabled                   types.Bool                                  `tfsdk:"enabled"`
	SteeringPolicy            types.String                                `tfsdk:"steering_policy"`
	AdaptiveRouting           []SourceAdaptiveRoutingModel                `tfsdk:"adaptive_routing"`    // Array in v4, single object in v5
	LocationStrategy          []SourceLocationStrategyModel               `tfsdk:"location_strategy"`   // Array in v4, single object in v5
	RandomSteering            []SourceRandomSteeringModel                 `tfsdk:"random_steering"`     // Array in v4, single object in v5
	RegionPools               []SourceRegionPoolModel                     `tfsdk:"region_pools"`        // Set of blocks in v4, map in v5
	PopPools                  []SourcePopPoolModel                        `tfsdk:"pop_pools"`           // Set of blocks in v4, map in v5
	CountryPools              []SourceCountryPoolModel                    `tfsdk:"country_pools"`       // Set of blocks in v4, map in v5
	Rules                     []SourceRulesModel                          `tfsdk:"rules"`
	CreatedOn                 types.String                                `tfsdk:"created_on"`
	ModifiedOn                types.String                                `tfsdk:"modified_on"`
}

// SourceSessionAffinityAttributesModel represents v4 session_affinity_attributes structure.
// Stored as TypeList MaxItems:1 in v4, will become single object in v5.
type SourceSessionAffinityAttributesModel struct {
	Samesite             types.String `tfsdk:"samesite"`
	Secure               types.String `tfsdk:"secure"`
	DrainDuration        types.Int64  `tfsdk:"drain_duration"` // Int64 in v4, Float64 in v5
	ZeroDowntimeFailover types.String `tfsdk:"zero_downtime_failover"`
	Headers              types.List   `tfsdk:"headers"` // List of strings
	RequireAllHeaders    types.Bool   `tfsdk:"require_all_headers"`
}

// SourceAdaptiveRoutingModel represents v4 adaptive_routing structure.
// Stored as TypeList MaxItems:1 in v4, will become single object in v5.
type SourceAdaptiveRoutingModel struct {
	FailoverAcrossPools types.Bool `tfsdk:"failover_across_pools"`
}

// SourceLocationStrategyModel represents v4 location_strategy structure.
// Stored as TypeList MaxItems:1 in v4, will become single object in v5.
type SourceLocationStrategyModel struct {
	PreferECS types.String `tfsdk:"prefer_ecs"`
	Mode      types.String `tfsdk:"mode"`
}

// SourceRandomSteeringModel represents v4 random_steering structure.
// Stored as TypeList MaxItems:1 in v4, will become single object in v5.
type SourceRandomSteeringModel struct {
	PoolWeights   types.Map     `tfsdk:"pool_weights"` // Map of pool_id -> weight (Float)
	DefaultWeight types.Float64 `tfsdk:"default_weight"`
}

// SourceRegionPoolModel represents v4 region_pools block structure.
// Stored as TypeSet of blocks in v4, will become map in v5.
type SourceRegionPoolModel struct {
	Region  types.String `tfsdk:"region"`
	PoolIds types.List   `tfsdk:"pool_ids"` // List of strings
}

// SourcePopPoolModel represents v4 pop_pools block structure.
// Stored as TypeSet of blocks in v4, will become map in v5.
type SourcePopPoolModel struct {
	Pop     types.String `tfsdk:"pop"`
	PoolIds types.List   `tfsdk:"pool_ids"` // List of strings
}

// SourceCountryPoolModel represents v4 country_pools block structure.
// Stored as TypeSet of blocks in v4, will become map in v5.
type SourceCountryPoolModel struct {
	Country types.String `tfsdk:"country"`
	PoolIds types.List   `tfsdk:"pool_ids"` // List of strings
}

// SourceRulesModel represents v4 rules structure.
type SourceRulesModel struct {
	Name          types.String                `tfsdk:"name"`
	Priority      types.Int64                 `tfsdk:"priority"`
	Disabled      types.Bool                  `tfsdk:"disabled"`
	Condition     types.String                `tfsdk:"condition"`
	Terminates    types.Bool                  `tfsdk:"terminates"`
	Overrides     []SourceRulesOverridesModel `tfsdk:"overrides"`      // Array MaxItems:1 in v4
	FixedResponse []SourceFixedResponseModel  `tfsdk:"fixed_response"` // Array MaxItems:1 in v4
}

// SourceRulesOverridesModel represents v4 rules.overrides structure.
// Mirrors top-level transformations (same nested structure changes).
type SourceRulesOverridesModel struct {
	SessionAffinity           types.String                                `tfsdk:"session_affinity"`
	SessionAffinityTTL        types.Int64                                 `tfsdk:"session_affinity_ttl"`        // Int64 in v4, Float64 in v5
	SessionAffinityAttributes []SourceSessionAffinityAttributesModel      `tfsdk:"session_affinity_attributes"` // Array in v4, single object in v5
	TTL                       types.Int64                                 `tfsdk:"ttl"`                         // Int64 in v4, Float64 in v5
	SteeringPolicy            types.String                                `tfsdk:"steering_policy"`
	FallbackPool              types.String                                `tfsdk:"fallback_pool"`
	DefaultPools              types.List                                  `tfsdk:"default_pools"`
	AdaptiveRouting           []SourceAdaptiveRoutingModel                `tfsdk:"adaptive_routing"`    // Array in v4, single object in v5
	LocationStrategy          []SourceLocationStrategyModel               `tfsdk:"location_strategy"`   // Array in v4, single object in v5
	RandomSteering            []SourceRandomSteeringModel                 `tfsdk:"random_steering"`     // Array in v4, single object in v5
	RegionPools               []SourceRegionPoolModel                     `tfsdk:"region_pools"`        // Set of blocks in v4, map in v5
	PopPools                  []SourcePopPoolModel                        `tfsdk:"pop_pools"`           // Set of blocks in v4, map in v5
	CountryPools              []SourceCountryPoolModel                    `tfsdk:"country_pools"`       // Set of blocks in v4, map in v5
}

// SourceFixedResponseModel represents v4 fixed_response structure.
type SourceFixedResponseModel struct {
	MessageBody types.String `tfsdk:"message_body"`
	StatusCode  types.Int64  `tfsdk:"status_code"`
	ContentType types.String `tfsdk:"content_type"`
	Location    types.String `tfsdk:"location"`
}

// ============================================================================
// Target Models (Current Provider - v5.x+)
// ============================================================================
// These mirror the models in the parent package (../model.go) but are
// duplicated here to keep the migration package self-contained.

// TargetLoadBalancerModel represents the target cloudflare_load_balancer state (v500).
type TargetLoadBalancerModel struct {
	ID                        types.String                                                         `tfsdk:"id"`
	ZoneID                    types.String                                                         `tfsdk:"zone_id"`
	FallbackPool              types.String                                                         `tfsdk:"fallback_pool"` // Renamed from fallback_pool_id
	Name                      types.String                                                         `tfsdk:"name"`
	DefaultPools              *[]types.String                                                      `tfsdk:"default_pools"` // Renamed from default_pool_ids
	Description               types.String                                                         `tfsdk:"description"`
	SessionAffinityTTL        types.Float64                                                        `tfsdk:"session_affinity_ttl"` // Converted from Int64
	TTL                       types.Float64                                                        `tfsdk:"ttl"`                  // Converted from Int64
	CountryPools              customfield.Map[customfield.List[types.String]]                      `tfsdk:"country_pools"`        // Converted from set of blocks
	Networks                  customfield.List[types.String]                                       `tfsdk:"networks"`             // New field in v5
	POPPools                  customfield.Map[customfield.List[types.String]]                      `tfsdk:"pop_pools"`            // Converted from set of blocks
	RegionPools               customfield.Map[customfield.List[types.String]]                      `tfsdk:"region_pools"`         // Converted from set of blocks
	Enabled                   types.Bool                                                           `tfsdk:"enabled"`
	Proxied                   types.Bool                                                           `tfsdk:"proxied"`
	SessionAffinity           types.String                                                         `tfsdk:"session_affinity"`
	SteeringPolicy            types.String                                                         `tfsdk:"steering_policy"`
	AdaptiveRouting           customfield.NestedObject[TargetAdaptiveRoutingModel]                 `tfsdk:"adaptive_routing"`             // Converted from array
	LocationStrategy          customfield.NestedObject[TargetLocationStrategyModel]                `tfsdk:"location_strategy"`            // Converted from array
	RandomSteering            customfield.NestedObject[TargetRandomSteeringModel]                  `tfsdk:"random_steering"`              // Converted from array
	Rules                     customfield.NestedObjectList[TargetRulesModel]                       `tfsdk:"rules"`
	SessionAffinityAttributes customfield.NestedObject[TargetSessionAffinityAttributesModel]       `tfsdk:"session_affinity_attributes"`  // Converted from array
	CreatedOn                 types.String                                                         `tfsdk:"created_on"`
	ModifiedOn                types.String                                                         `tfsdk:"modified_on"`
	ZoneName                  types.String                                                         `tfsdk:"zone_name"` // Computed, may not exist in v4
}

// TargetSessionAffinityAttributesModel represents v5 session_affinity_attributes structure.
type TargetSessionAffinityAttributesModel struct {
	DrainDuration        types.Float64   `tfsdk:"drain_duration"` // Converted from Int64
	Headers              *[]types.String `tfsdk:"headers"`
	RequireAllHeaders    types.Bool      `tfsdk:"require_all_headers"`
	Samesite             types.String    `tfsdk:"samesite"`
	Secure               types.String    `tfsdk:"secure"`
	ZeroDowntimeFailover types.String    `tfsdk:"zero_downtime_failover"`
}

// TargetAdaptiveRoutingModel represents v5 adaptive_routing structure.
type TargetAdaptiveRoutingModel struct {
	FailoverAcrossPools types.Bool `tfsdk:"failover_across_pools"`
}

// TargetLocationStrategyModel represents v5 location_strategy structure.
type TargetLocationStrategyModel struct {
	Mode      types.String `tfsdk:"mode"`
	PreferECS types.String `tfsdk:"prefer_ecs"`
}

// TargetRandomSteeringModel represents v5 random_steering structure.
type TargetRandomSteeringModel struct {
	DefaultWeight types.Float64             `tfsdk:"default_weight"`
	PoolWeights   *map[string]types.Float64 `tfsdk:"pool_weights"`
}

// TargetRulesModel represents v5 rules structure.
type TargetRulesModel struct {
	Condition     types.String                                              `tfsdk:"condition"`
	Disabled      types.Bool                                                `tfsdk:"disabled"`
	FixedResponse *TargetFixedResponseModel                                 `tfsdk:"fixed_response"` // Converted from array
	Name          types.String                                              `tfsdk:"name"`
	Overrides     customfield.NestedObject[TargetRulesOverridesModel]       `tfsdk:"overrides"` // Converted from array
	Priority      types.Int64                                               `tfsdk:"priority"`
	Terminates    types.Bool                                                `tfsdk:"terminates"`
}

// TargetFixedResponseModel represents v5 fixed_response structure.
type TargetFixedResponseModel struct {
	ContentType types.String `tfsdk:"content_type"`
	Location    types.String `tfsdk:"location"`
	MessageBody types.String `tfsdk:"message_body"`
	StatusCode  types.Int64  `tfsdk:"status_code"`
}

// TargetRulesOverridesModel represents v5 rules.overrides structure.
type TargetRulesOverridesModel struct {
	AdaptiveRouting           customfield.NestedObject[TargetAdaptiveRoutingModel]           `tfsdk:"adaptive_routing"`
	CountryPools              customfield.Map[customfield.List[types.String]]                `tfsdk:"country_pools"`
	DefaultPools              *[]types.String                                                `tfsdk:"default_pools"`
	FallbackPool              types.String                                                   `tfsdk:"fallback_pool"`
	LocationStrategy          customfield.NestedObject[TargetLocationStrategyModel]          `tfsdk:"location_strategy"`
	POPPools                  customfield.Map[customfield.List[types.String]]                `tfsdk:"pop_pools"`
	RandomSteering            customfield.NestedObject[TargetRandomSteeringModel]            `tfsdk:"random_steering"`
	RegionPools               customfield.Map[customfield.List[types.String]]                `tfsdk:"region_pools"`
	SessionAffinity           types.String                                                   `tfsdk:"session_affinity"`
	SessionAffinityAttributes customfield.NestedObject[TargetSessionAffinityAttributesModel] `tfsdk:"session_affinity_attributes"`
	SessionAffinityTTL        types.Float64                                                  `tfsdk:"session_affinity_ttl"` // Converted from Int64
	SteeringPolicy            types.String                                                   `tfsdk:"steering_policy"`
	TTL                       types.Float64                                                  `tfsdk:"ttl"` // Converted from Int64
}
