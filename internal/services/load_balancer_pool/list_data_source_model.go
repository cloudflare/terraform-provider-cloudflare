// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer_pool

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LoadBalancerPoolsResultListDataSourceEnvelope struct {
	Result *[]*LoadBalancerPoolsItemsDataSourceModel `json:"result,computed"`
}

type LoadBalancerPoolsDataSourceModel struct {
	AccountID types.String                              `tfsdk:"account_id" path:"account_id"`
	Monitor   types.String                              `tfsdk:"monitor" query:"monitor"`
	MaxItems  types.Int64                               `tfsdk:"max_items"`
	Items     *[]*LoadBalancerPoolsItemsDataSourceModel `tfsdk:"items"`
}

type LoadBalancerPoolsItemsDataSourceModel struct {
	ID                types.String                                     `tfsdk:"id" json:"id,computed"`
	CheckRegions      *[]types.String                                  `tfsdk:"check_regions" json:"check_regions,computed"`
	CreatedOn         types.String                                     `tfsdk:"created_on" json:"created_on,computed"`
	Description       types.String                                     `tfsdk:"description" json:"description,computed"`
	DisabledAt        types.String                                     `tfsdk:"disabled_at" json:"disabled_at,computed"`
	Enabled           types.Bool                                       `tfsdk:"enabled" json:"enabled,computed"`
	Latitude          types.Float64                                    `tfsdk:"latitude" json:"latitude,computed"`
	Longitude         types.Float64                                    `tfsdk:"longitude" json:"longitude,computed"`
	MinimumOrigins    types.Int64                                      `tfsdk:"minimum_origins" json:"minimum_origins,computed"`
	ModifiedOn        types.String                                     `tfsdk:"modified_on" json:"modified_on,computed"`
	Monitor           types.String                                     `tfsdk:"monitor" json:"monitor,computed"`
	Name              types.String                                     `tfsdk:"name" json:"name,computed"`
	NotificationEmail types.String                                     `tfsdk:"notification_email" json:"notification_email,computed"`
	Origins           *[]*LoadBalancerPoolsItemsOriginsDataSourceModel `tfsdk:"origins" json:"origins,computed"`
}

type LoadBalancerPoolsItemsLoadSheddingDataSourceModel struct {
	DefaultPercent types.Float64 `tfsdk:"default_percent" json:"default_percent,computed"`
	DefaultPolicy  types.String  `tfsdk:"default_policy" json:"default_policy,computed"`
	SessionPercent types.Float64 `tfsdk:"session_percent" json:"session_percent,computed"`
	SessionPolicy  types.String  `tfsdk:"session_policy" json:"session_policy,computed"`
}

type LoadBalancerPoolsItemsNotificationFilterDataSourceModel struct {
}

type LoadBalancerPoolsItemsNotificationFilterOriginDataSourceModel struct {
	Disable types.Bool `tfsdk:"disable" json:"disable,computed"`
	Healthy types.Bool `tfsdk:"healthy" json:"healthy,computed"`
}

type LoadBalancerPoolsItemsNotificationFilterPoolDataSourceModel struct {
	Disable types.Bool `tfsdk:"disable" json:"disable,computed"`
	Healthy types.Bool `tfsdk:"healthy" json:"healthy,computed"`
}

type LoadBalancerPoolsItemsOriginSteeringDataSourceModel struct {
	Policy types.String `tfsdk:"policy" json:"policy,computed"`
}

type LoadBalancerPoolsItemsOriginsDataSourceModel struct {
	Address          types.String  `tfsdk:"address" json:"address,computed"`
	DisabledAt       types.String  `tfsdk:"disabled_at" json:"disabled_at,computed"`
	Enabled          types.Bool    `tfsdk:"enabled" json:"enabled,computed"`
	Name             types.String  `tfsdk:"name" json:"name,computed"`
	VirtualNetworkID types.String  `tfsdk:"virtual_network_id" json:"virtual_network_id,computed"`
	Weight           types.Float64 `tfsdk:"weight" json:"weight,computed"`
}

type LoadBalancerPoolsItemsOriginsHeaderDataSourceModel struct {
	Host *[]types.String `tfsdk:"host" json:"Host,computed"`
}
