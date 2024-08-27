// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer_pool

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LoadBalancerPoolResultEnvelope struct {
	Result LoadBalancerPoolModel `json:"result"`
}

type LoadBalancerPoolModel struct {
	ID                 types.String                             `tfsdk:"id" json:"id,computed"`
	AccountID          types.String                             `tfsdk:"account_id" path:"account_id"`
	Name               types.String                             `tfsdk:"name" json:"name"`
	Origins            *[]*LoadBalancerPoolOriginsModel         `tfsdk:"origins" json:"origins"`
	Description        types.String                             `tfsdk:"description" json:"description"`
	Latitude           types.Float64                            `tfsdk:"latitude" json:"latitude"`
	Longitude          types.Float64                            `tfsdk:"longitude" json:"longitude"`
	Monitor            types.String                             `tfsdk:"monitor" json:"monitor"`
	NotificationEmail  types.String                             `tfsdk:"notification_email" json:"notification_email"`
	CheckRegions       *[]types.String                          `tfsdk:"check_regions" json:"check_regions"`
	LoadShedding       *LoadBalancerPoolLoadSheddingModel       `tfsdk:"load_shedding" json:"load_shedding"`
	NotificationFilter *LoadBalancerPoolNotificationFilterModel `tfsdk:"notification_filter" json:"notification_filter"`
	OriginSteering     *LoadBalancerPoolOriginSteeringModel     `tfsdk:"origin_steering" json:"origin_steering"`
	Enabled            types.Bool                               `tfsdk:"enabled" json:"enabled,computed_optional"`
	MinimumOrigins     types.Int64                              `tfsdk:"minimum_origins" json:"minimum_origins,computed_optional"`
	CreatedOn          timetypes.RFC3339                        `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	DisabledAt         timetypes.RFC3339                        `tfsdk:"disabled_at" json:"disabled_at,computed" format:"date-time"`
	ModifiedOn         timetypes.RFC3339                        `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
}

type LoadBalancerPoolOriginsModel struct {
	Address          types.String                        `tfsdk:"address" json:"address"`
	DisabledAt       timetypes.RFC3339                   `tfsdk:"disabled_at" json:"disabled_at,computed" format:"date-time"`
	Enabled          types.Bool                          `tfsdk:"enabled" json:"enabled,computed_optional"`
	Header           *LoadBalancerPoolOriginsHeaderModel `tfsdk:"header" json:"header"`
	Name             types.String                        `tfsdk:"name" json:"name"`
	VirtualNetworkID types.String                        `tfsdk:"virtual_network_id" json:"virtual_network_id"`
	Weight           types.Float64                       `tfsdk:"weight" json:"weight,computed_optional"`
}

type LoadBalancerPoolOriginsHeaderModel struct {
	Host *[]types.String `tfsdk:"host" json:"Host"`
}

type LoadBalancerPoolLoadSheddingModel struct {
	DefaultPercent types.Float64 `tfsdk:"default_percent" json:"default_percent,computed_optional"`
	DefaultPolicy  types.String  `tfsdk:"default_policy" json:"default_policy,computed_optional"`
	SessionPercent types.Float64 `tfsdk:"session_percent" json:"session_percent,computed_optional"`
	SessionPolicy  types.String  `tfsdk:"session_policy" json:"session_policy,computed_optional"`
}

type LoadBalancerPoolNotificationFilterModel struct {
	Origin *LoadBalancerPoolNotificationFilterOriginModel `tfsdk:"origin" json:"origin"`
	Pool   *LoadBalancerPoolNotificationFilterPoolModel   `tfsdk:"pool" json:"pool"`
}

type LoadBalancerPoolNotificationFilterOriginModel struct {
	Disable types.Bool `tfsdk:"disable" json:"disable,computed_optional"`
	Healthy types.Bool `tfsdk:"healthy" json:"healthy"`
}

type LoadBalancerPoolNotificationFilterPoolModel struct {
	Disable types.Bool `tfsdk:"disable" json:"disable,computed_optional"`
	Healthy types.Bool `tfsdk:"healthy" json:"healthy"`
}

type LoadBalancerPoolOriginSteeringModel struct {
	Policy types.String `tfsdk:"policy" json:"policy,computed_optional"`
}
