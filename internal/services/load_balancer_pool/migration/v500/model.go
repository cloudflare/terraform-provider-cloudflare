// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v500

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x)
// ============================================================================

// SourceCloudflareLoadBalancerPoolModel represents the legacy load_balancer_pool resource state from v4.x provider.
// Schema version: 0 (v4 had no schema version)
// Resource type: cloudflare_load_balancer_pool
//
// Note: SDK v2 storage quirks:
// - TypeSet fields stored as arrays in state (order not guaranteed)
// - TypeList MaxItems:1 fields stored as arrays with single element: [{...}]
type SourceCloudflareLoadBalancerPoolModel struct {
	ID                types.String              `tfsdk:"id"`
	AccountID         types.String              `tfsdk:"account_id"`
	Name              types.String              `tfsdk:"name"`
	Origins           []SourceOriginsModel      `tfsdk:"origins"` // TypeSet stored as array
	Enabled           types.Bool                `tfsdk:"enabled"`
	MinimumOrigins    types.Int64               `tfsdk:"minimum_origins"` // Int in v4, Int64 in v5
	Latitude          types.Float64             `tfsdk:"latitude"`
	Longitude         types.Float64             `tfsdk:"longitude"`
	CheckRegions      types.Set                 `tfsdk:"check_regions"` // TypeSet, will convert to List
	Description       types.String              `tfsdk:"description"`
	Monitor           types.String              `tfsdk:"monitor"`
	NotificationEmail types.String              `tfsdk:"notification_email"`
	LoadShedding      []SourceLoadSheddingModel `tfsdk:"load_shedding"` // TypeSet MaxItems:1 stored as array
	OriginSteering    []SourceOriginSteeringModel `tfsdk:"origin_steering"` // TypeSet MaxItems:1 stored as array
	CreatedOn         types.String              `tfsdk:"created_on"`
	ModifiedOn        types.String              `tfsdk:"modified_on"`
}

// SourceOriginsModel represents a single origin within the pool from v4.x provider.
// Stored in TypeSet, so in state it's an array of these objects.
type SourceOriginsModel struct {
	Name             types.String       `tfsdk:"name"`
	Address          types.String       `tfsdk:"address"`
	VirtualNetworkID types.String       `tfsdk:"virtual_network_id"`
	Weight           types.Float64      `tfsdk:"weight"`
	Enabled          types.Bool         `tfsdk:"enabled"`
	Header           []SourceHeaderModel `tfsdk:"header"` // TypeSet stored as array
}

// SourceHeaderModel represents HTTP request headers from v4.x provider.
// v4 format: { header: "Host", values: ["value1", "value2"] }
// v5 format: { host: ["value1", "value2"] }
type SourceHeaderModel struct {
	Header types.String `tfsdk:"header"` // Header name (e.g., "Host")
	Values types.Set    `tfsdk:"values"` // Set of values
}

// SourceLoadSheddingModel represents load shedding configuration from v4.x provider.
// SDK v2 stores TypeList MaxItems:1 as array: [{...}]
type SourceLoadSheddingModel struct {
	DefaultPercent types.Float64 `tfsdk:"default_percent"`
	DefaultPolicy  types.String  `tfsdk:"default_policy"` // Default: "" in v4, "random" in v5
	SessionPercent types.Float64 `tfsdk:"session_percent"`
	SessionPolicy  types.String  `tfsdk:"session_policy"` // Default: "" in v4, "hash" in v5
}

// SourceOriginSteeringModel represents origin steering policy from v4.x provider.
// SDK v2 stores TypeList MaxItems:1 as array: [{...}]
type SourceOriginSteeringModel struct {
	Policy types.String `tfsdk:"policy"` // Default: "random" in both v4 and v5
}

// ============================================================================
// Target Models (Current Provider - v5.x+)
// ============================================================================

// TargetLoadBalancerPoolModel represents the current load_balancer_pool resource state from v5.x+ provider.
// Schema version: 500
// Resource type: cloudflare_load_balancer_pool
//
// Note: This duplicates the model from the parent package to keep migration self-contained.
type TargetLoadBalancerPoolModel struct {
	ID                 types.String                                                          `tfsdk:"id"`
	AccountID          types.String                                                          `tfsdk:"account_id"`
	Name               types.String                                                          `tfsdk:"name"`
	Origins            *[]*TargetLoadBalancerPoolOriginsModel                                `tfsdk:"origins"` // Pointer to slice of pointers
	Latitude           types.Float64                                                         `tfsdk:"latitude"`
	Longitude          types.Float64                                                         `tfsdk:"longitude"`
	Monitor            types.String                                                          `tfsdk:"monitor"`
	MonitorGroup       types.String                                                          `tfsdk:"monitor_group"` // NEW in v5
	CheckRegions       *[]types.String                                                       `tfsdk:"check_regions"` // Set → List
	Description        types.String                                                          `tfsdk:"description"`
	Enabled            types.Bool                                                            `tfsdk:"enabled"`
	MinimumOrigins     types.Int64                                                           `tfsdk:"minimum_origins"`
	NotificationEmail  types.String                                                          `tfsdk:"notification_email"` // Deprecated in v5
	LoadShedding       customfield.NestedObject[TargetLoadBalancerPoolLoadSheddingModel]    `tfsdk:"load_shedding"` // Array → Object
	NotificationFilter customfield.NestedObject[TargetLoadBalancerPoolNotificationFilterModel] `tfsdk:"notification_filter"` // NEW in v5
	OriginSteering     customfield.NestedObject[TargetLoadBalancerPoolOriginSteeringModel]  `tfsdk:"origin_steering"` // Array → Object
	CreatedOn          types.String                                                          `tfsdk:"created_on"`
	DisabledAt         timetypes.RFC3339                                                     `tfsdk:"disabled_at"` // NEW in v5
	ModifiedOn         types.String                                                          `tfsdk:"modified_on"`
	Networks           customfield.List[types.String]                                        `tfsdk:"networks"` // NEW in v5
}

// TargetLoadBalancerPoolOriginsModel represents a single origin within the pool from v5.x+ provider.
type TargetLoadBalancerPoolOriginsModel struct {
	Address          types.String                                   `tfsdk:"address"`
	DisabledAt       timetypes.RFC3339                              `tfsdk:"disabled_at"` // NEW in v5
	Enabled          types.Bool                                     `tfsdk:"enabled"`
	Header           *TargetLoadBalancerPoolOriginsHeaderModel      `tfsdk:"header"` // Structure changed
	Name             types.String                                   `tfsdk:"name"`
	Port             types.Int64                                    `tfsdk:"port"` // NEW in v5
	VirtualNetworkID types.String                                   `tfsdk:"virtual_network_id"`
	Weight           types.Float64                                  `tfsdk:"weight"`
}

// TargetLoadBalancerPoolOriginsHeaderModel represents HTTP request headers from v5.x+ provider.
// v5 format: { host: ["value1", "value2"] }
type TargetLoadBalancerPoolOriginsHeaderModel struct {
	Host *[]types.String `tfsdk:"host"` // List of host values
}

// TargetLoadBalancerPoolLoadSheddingModel represents load shedding configuration from v5.x+ provider.
type TargetLoadBalancerPoolLoadSheddingModel struct {
	DefaultPercent types.Float64 `tfsdk:"default_percent"`
	DefaultPolicy  types.String  `tfsdk:"default_policy"` // Default: "random" in v5
	SessionPercent types.Float64 `tfsdk:"session_percent"`
	SessionPolicy  types.String  `tfsdk:"session_policy"` // Default: "hash" in v5
}

// TargetLoadBalancerPoolNotificationFilterModel represents notification filtering from v5.x+ provider.
// NEW in v5 - set to null/empty during migration.
type TargetLoadBalancerPoolNotificationFilterModel struct {
	Origin customfield.NestedObject[TargetLoadBalancerPoolNotificationFilterOriginModel] `tfsdk:"origin"`
	Pool   customfield.NestedObject[TargetLoadBalancerPoolNotificationFilterPoolModel]   `tfsdk:"pool"`
}

// TargetLoadBalancerPoolNotificationFilterOriginModel represents origin notification filter settings.
type TargetLoadBalancerPoolNotificationFilterOriginModel struct {
	Disable types.Bool `tfsdk:"disable"`
	Healthy types.Bool `tfsdk:"healthy"`
}

// TargetLoadBalancerPoolNotificationFilterPoolModel represents pool notification filter settings.
type TargetLoadBalancerPoolNotificationFilterPoolModel struct {
	Disable types.Bool `tfsdk:"disable"`
	Healthy types.Bool `tfsdk:"healthy"`
}

// TargetLoadBalancerPoolOriginSteeringModel represents origin steering policy from v5.x+ provider.
type TargetLoadBalancerPoolOriginSteeringModel struct {
	Policy types.String `tfsdk:"policy"` // Default: "random"
}
