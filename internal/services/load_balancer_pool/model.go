// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer_pool

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LoadBalancerPoolResultEnvelope struct {
	Result LoadBalancerPoolModel `json:"result"`
}

type LoadBalancerPoolModel struct {
	ID                 types.String                                                      `tfsdk:"id" json:"id,computed"`
	AccountID          types.String                                                      `tfsdk:"account_id" path:"account_id,required"`
	Name               types.String                                                      `tfsdk:"name" json:"name,required"`
	Origins            *[]*LoadBalancerPoolOriginsModel                                  `tfsdk:"origins" json:"origins,required"`
	Description        types.String                                                      `tfsdk:"description" json:"description,computed_optional"`
	Enabled            types.Bool                                                        `tfsdk:"enabled" json:"enabled,computed_optional"`
	Latitude           types.Float64                                                     `tfsdk:"latitude" json:"latitude,computed_optional"`
	Longitude          types.Float64                                                     `tfsdk:"longitude" json:"longitude,computed_optional"`
	MinimumOrigins     types.Int64                                                       `tfsdk:"minimum_origins" json:"minimum_origins,computed_optional"`
	Monitor            types.String                                                      `tfsdk:"monitor" json:"monitor,computed_optional"`
	NotificationEmail  types.String                                                      `tfsdk:"notification_email" json:"notification_email,computed_optional"`
	CheckRegions       customfield.List[types.String]                                    `tfsdk:"check_regions" json:"check_regions,computed_optional"`
	LoadShedding       customfield.NestedObject[LoadBalancerPoolLoadSheddingModel]       `tfsdk:"load_shedding" json:"load_shedding,computed_optional"`
	NotificationFilter customfield.NestedObject[LoadBalancerPoolNotificationFilterModel] `tfsdk:"notification_filter" json:"notification_filter,computed_optional"`
	OriginSteering     customfield.NestedObject[LoadBalancerPoolOriginSteeringModel]     `tfsdk:"origin_steering" json:"origin_steering,computed_optional"`
	CreatedOn          timetypes.RFC3339                                                 `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	DisabledAt         timetypes.RFC3339                                                 `tfsdk:"disabled_at" json:"disabled_at,computed" format:"date-time"`
	ModifiedOn         timetypes.RFC3339                                                 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Networks           customfield.List[types.String]                                    `tfsdk:"networks" json:"networks,computed"`
}

type LoadBalancerPoolOriginsModel struct {
	Address          types.String                                                 `tfsdk:"address" json:"address,computed_optional"`
	DisabledAt       timetypes.RFC3339                                            `tfsdk:"disabled_at" json:"disabled_at,computed" format:"date-time"`
	Enabled          types.Bool                                                   `tfsdk:"enabled" json:"enabled,computed_optional"`
	Header           customfield.NestedObject[LoadBalancerPoolOriginsHeaderModel] `tfsdk:"header" json:"header,computed_optional"`
	Name             types.String                                                 `tfsdk:"name" json:"name,computed_optional"`
	VirtualNetworkID types.String                                                 `tfsdk:"virtual_network_id" json:"virtual_network_id,computed_optional"`
	Weight           types.Float64                                                `tfsdk:"weight" json:"weight,computed_optional"`
}

type LoadBalancerPoolOriginsHeaderModel struct {
	Host customfield.List[types.String] `tfsdk:"host" json:"Host,computed_optional"`
}

type LoadBalancerPoolLoadSheddingModel struct {
	DefaultPercent types.Float64 `tfsdk:"default_percent" json:"default_percent,computed_optional"`
	DefaultPolicy  types.String  `tfsdk:"default_policy" json:"default_policy,computed_optional"`
	SessionPercent types.Float64 `tfsdk:"session_percent" json:"session_percent,computed_optional"`
	SessionPolicy  types.String  `tfsdk:"session_policy" json:"session_policy,computed_optional"`
}

type LoadBalancerPoolNotificationFilterModel struct {
	Origin customfield.NestedObject[LoadBalancerPoolNotificationFilterOriginModel] `tfsdk:"origin" json:"origin,computed_optional"`
	Pool   customfield.NestedObject[LoadBalancerPoolNotificationFilterPoolModel]   `tfsdk:"pool" json:"pool,computed_optional"`
}

type LoadBalancerPoolNotificationFilterOriginModel struct {
	Disable types.Bool `tfsdk:"disable" json:"disable,computed_optional"`
	Healthy types.Bool `tfsdk:"healthy" json:"healthy,computed_optional"`
}

type LoadBalancerPoolNotificationFilterPoolModel struct {
	Disable types.Bool `tfsdk:"disable" json:"disable,computed_optional"`
	Healthy types.Bool `tfsdk:"healthy" json:"healthy,computed_optional"`
}

type LoadBalancerPoolOriginSteeringModel struct {
	Policy types.String `tfsdk:"policy" json:"policy,computed_optional"`
}
