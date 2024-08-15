// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer_pool

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LoadBalancerPoolsResultListDataSourceEnvelope struct {
	Result *[]*LoadBalancerPoolsResultDataSourceModel `json:"result,computed"`
}

type LoadBalancerPoolsDataSourceModel struct {
	AccountID types.String                               `tfsdk:"account_id" path:"account_id"`
	Monitor   types.String                               `tfsdk:"monitor" query:"monitor"`
	MaxItems  types.Int64                                `tfsdk:"max_items"`
	Result    *[]*LoadBalancerPoolsResultDataSourceModel `tfsdk:"result"`
}

type LoadBalancerPoolsResultDataSourceModel struct {
	ID                 types.String                                        `tfsdk:"id" json:"id"`
	CheckRegions       *[]types.String                                     `tfsdk:"check_regions" json:"check_regions"`
	CreatedOn          timetypes.RFC3339                                   `tfsdk:"created_on" json:"created_on,computed"`
	Description        types.String                                        `tfsdk:"description" json:"description"`
	DisabledAt         timetypes.RFC3339                                   `tfsdk:"disabled_at" json:"disabled_at,computed"`
	Enabled            types.Bool                                          `tfsdk:"enabled" json:"enabled,computed"`
	Latitude           types.Float64                                       `tfsdk:"latitude" json:"latitude"`
	LoadShedding       *LoadBalancerPoolsLoadSheddingDataSourceModel       `tfsdk:"load_shedding" json:"load_shedding"`
	Longitude          types.Float64                                       `tfsdk:"longitude" json:"longitude"`
	MinimumOrigins     types.Int64                                         `tfsdk:"minimum_origins" json:"minimum_origins,computed"`
	ModifiedOn         timetypes.RFC3339                                   `tfsdk:"modified_on" json:"modified_on,computed"`
	Monitor            types.String                                        `tfsdk:"monitor" json:"monitor"`
	Name               types.String                                        `tfsdk:"name" json:"name"`
	NotificationEmail  types.String                                        `tfsdk:"notification_email" json:"notification_email"`
	NotificationFilter *LoadBalancerPoolsNotificationFilterDataSourceModel `tfsdk:"notification_filter" json:"notification_filter"`
	OriginSteering     *LoadBalancerPoolsOriginSteeringDataSourceModel     `tfsdk:"origin_steering" json:"origin_steering"`
	Origins            *[]*LoadBalancerPoolsOriginsDataSourceModel         `tfsdk:"origins" json:"origins"`
}

type LoadBalancerPoolsLoadSheddingDataSourceModel struct {
	DefaultPercent types.Float64 `tfsdk:"default_percent" json:"default_percent,computed"`
	DefaultPolicy  types.String  `tfsdk:"default_policy" json:"default_policy,computed"`
	SessionPercent types.Float64 `tfsdk:"session_percent" json:"session_percent,computed"`
	SessionPolicy  types.String  `tfsdk:"session_policy" json:"session_policy,computed"`
}

type LoadBalancerPoolsNotificationFilterDataSourceModel struct {
	Origin *LoadBalancerPoolsNotificationFilterOriginDataSourceModel `tfsdk:"origin" json:"origin"`
	Pool   *LoadBalancerPoolsNotificationFilterPoolDataSourceModel   `tfsdk:"pool" json:"pool"`
}

type LoadBalancerPoolsNotificationFilterOriginDataSourceModel struct {
	Disable types.Bool `tfsdk:"disable" json:"disable,computed"`
	Healthy types.Bool `tfsdk:"healthy" json:"healthy"`
}

type LoadBalancerPoolsNotificationFilterPoolDataSourceModel struct {
	Disable types.Bool `tfsdk:"disable" json:"disable,computed"`
	Healthy types.Bool `tfsdk:"healthy" json:"healthy"`
}

type LoadBalancerPoolsOriginSteeringDataSourceModel struct {
	Policy types.String `tfsdk:"policy" json:"policy,computed"`
}

type LoadBalancerPoolsOriginsDataSourceModel struct {
	Address          types.String                                   `tfsdk:"address" json:"address"`
	DisabledAt       timetypes.RFC3339                              `tfsdk:"disabled_at" json:"disabled_at,computed"`
	Enabled          types.Bool                                     `tfsdk:"enabled" json:"enabled,computed"`
	Header           *LoadBalancerPoolsOriginsHeaderDataSourceModel `tfsdk:"header" json:"header"`
	Name             types.String                                   `tfsdk:"name" json:"name"`
	VirtualNetworkID types.String                                   `tfsdk:"virtual_network_id" json:"virtual_network_id"`
	Weight           types.Float64                                  `tfsdk:"weight" json:"weight,computed"`
}

type LoadBalancerPoolsOriginsHeaderDataSourceModel struct {
	Host *[]types.String `tfsdk:"host" json:"Host"`
}
