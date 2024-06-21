// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer_pool

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LoadBalancerPoolResultEnvelope struct {
	Result LoadBalancerPoolModel `json:"result,computed"`
}

type LoadBalancerPoolModel struct {
	ID                 types.String                             `tfsdk:"id" json:"id,computed"`
	AccountID          types.String                             `tfsdk:"account_id" path:"account_id"`
	Name               types.String                             `tfsdk:"name" json:"name"`
	Origins            *[]*LoadBalancerPoolOriginsModel         `tfsdk:"origins" json:"origins"`
	Description        types.String                             `tfsdk:"description" json:"description"`
	Enabled            types.Bool                               `tfsdk:"enabled" json:"enabled"`
	Latitude           types.Float64                            `tfsdk:"latitude" json:"latitude"`
	LoadShedding       *LoadBalancerPoolLoadSheddingModel       `tfsdk:"load_shedding" json:"load_shedding"`
	Longitude          types.Float64                            `tfsdk:"longitude" json:"longitude"`
	MinimumOrigins     types.Int64                              `tfsdk:"minimum_origins" json:"minimum_origins"`
	Monitor            types.String                             `tfsdk:"monitor" json:"monitor"`
	NotificationEmail  types.String                             `tfsdk:"notification_email" json:"notification_email"`
	NotificationFilter *LoadBalancerPoolNotificationFilterModel `tfsdk:"notification_filter" json:"notification_filter"`
	OriginSteering     *LoadBalancerPoolOriginSteeringModel     `tfsdk:"origin_steering" json:"origin_steering"`
}

type LoadBalancerPoolOriginsModel struct {
	Address          types.String                        `tfsdk:"address" json:"address"`
	DisabledAt       types.String                        `tfsdk:"disabled_at" json:"disabled_at,computed"`
	Enabled          types.Bool                          `tfsdk:"enabled" json:"enabled"`
	Header           *LoadBalancerPoolOriginsHeaderModel `tfsdk:"header" json:"header"`
	Name             types.String                        `tfsdk:"name" json:"name"`
	VirtualNetworkID types.String                        `tfsdk:"virtual_network_id" json:"virtual_network_id"`
	Weight           types.Float64                       `tfsdk:"weight" json:"weight"`
}

type LoadBalancerPoolOriginsHeaderModel struct {
	Host *[]types.String `tfsdk:"host" json:"Host"`
}

type LoadBalancerPoolLoadSheddingModel struct {
	DefaultPercent types.Float64 `tfsdk:"default_percent" json:"default_percent"`
	DefaultPolicy  types.String  `tfsdk:"default_policy" json:"default_policy"`
	SessionPercent types.Float64 `tfsdk:"session_percent" json:"session_percent"`
	SessionPolicy  types.String  `tfsdk:"session_policy" json:"session_policy"`
}

type LoadBalancerPoolNotificationFilterModel struct {
	Origin *LoadBalancerPoolNotificationFilterOriginModel `tfsdk:"origin" json:"origin"`
	Pool   *LoadBalancerPoolNotificationFilterPoolModel   `tfsdk:"pool" json:"pool"`
}

type LoadBalancerPoolNotificationFilterOriginModel struct {
	Disable types.Bool `tfsdk:"disable" json:"disable"`
	Healthy types.Bool `tfsdk:"healthy" json:"healthy"`
}

type LoadBalancerPoolNotificationFilterPoolModel struct {
	Disable types.Bool `tfsdk:"disable" json:"disable"`
	Healthy types.Bool `tfsdk:"healthy" json:"healthy"`
}

type LoadBalancerPoolOriginSteeringModel struct {
	Policy types.String `tfsdk:"policy" json:"policy"`
}
