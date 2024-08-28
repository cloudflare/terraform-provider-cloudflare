// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer_pool

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/load_balancers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LoadBalancerPoolsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[LoadBalancerPoolsResultDataSourceModel] `json:"result,computed"`
}

type LoadBalancerPoolsDataSourceModel struct {
	AccountID types.String                                                         `tfsdk:"account_id" path:"account_id"`
	Monitor   types.String                                                         `tfsdk:"monitor" query:"monitor"`
	MaxItems  types.Int64                                                          `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[LoadBalancerPoolsResultDataSourceModel] `tfsdk:"result"`
}

func (m *LoadBalancerPoolsDataSourceModel) toListParams() (params load_balancers.PoolListParams, diags diag.Diagnostics) {
	params = load_balancers.PoolListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Monitor.IsNull() {
		params.Monitor = cloudflare.F(m.Monitor.ValueString())
	}

	return
}

type LoadBalancerPoolsResultDataSourceModel struct {
	ID                 types.String                                        `tfsdk:"id" json:"id,computed_optional"`
	CheckRegions       *[]types.String                                     `tfsdk:"check_regions" json:"check_regions,computed_optional"`
	CreatedOn          timetypes.RFC3339                                   `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Description        types.String                                        `tfsdk:"description" json:"description,computed_optional"`
	DisabledAt         timetypes.RFC3339                                   `tfsdk:"disabled_at" json:"disabled_at,computed" format:"date-time"`
	Enabled            types.Bool                                          `tfsdk:"enabled" json:"enabled,computed"`
	Latitude           types.Float64                                       `tfsdk:"latitude" json:"latitude,computed_optional"`
	LoadShedding       *LoadBalancerPoolsLoadSheddingDataSourceModel       `tfsdk:"load_shedding" json:"load_shedding,computed_optional"`
	Longitude          types.Float64                                       `tfsdk:"longitude" json:"longitude,computed_optional"`
	MinimumOrigins     types.Int64                                         `tfsdk:"minimum_origins" json:"minimum_origins,computed"`
	ModifiedOn         timetypes.RFC3339                                   `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Monitor            types.String                                        `tfsdk:"monitor" json:"monitor,computed_optional"`
	Name               types.String                                        `tfsdk:"name" json:"name,computed_optional"`
	Networks           *[]types.String                                     `tfsdk:"networks" json:"networks,computed_optional"`
	NotificationEmail  types.String                                        `tfsdk:"notification_email" json:"notification_email,computed_optional"`
	NotificationFilter *LoadBalancerPoolsNotificationFilterDataSourceModel `tfsdk:"notification_filter" json:"notification_filter,computed_optional"`
	OriginSteering     *LoadBalancerPoolsOriginSteeringDataSourceModel     `tfsdk:"origin_steering" json:"origin_steering,computed_optional"`
	Origins            *[]*LoadBalancerPoolsOriginsDataSourceModel         `tfsdk:"origins" json:"origins,computed_optional"`
}

type LoadBalancerPoolsLoadSheddingDataSourceModel struct {
	DefaultPercent types.Float64 `tfsdk:"default_percent" json:"default_percent,computed"`
	DefaultPolicy  types.String  `tfsdk:"default_policy" json:"default_policy,computed"`
	SessionPercent types.Float64 `tfsdk:"session_percent" json:"session_percent,computed"`
	SessionPolicy  types.String  `tfsdk:"session_policy" json:"session_policy,computed"`
}

type LoadBalancerPoolsNotificationFilterDataSourceModel struct {
	Origin *LoadBalancerPoolsNotificationFilterOriginDataSourceModel `tfsdk:"origin" json:"origin,computed_optional"`
	Pool   *LoadBalancerPoolsNotificationFilterPoolDataSourceModel   `tfsdk:"pool" json:"pool,computed_optional"`
}

type LoadBalancerPoolsNotificationFilterOriginDataSourceModel struct {
	Disable types.Bool `tfsdk:"disable" json:"disable,computed"`
	Healthy types.Bool `tfsdk:"healthy" json:"healthy,computed_optional"`
}

type LoadBalancerPoolsNotificationFilterPoolDataSourceModel struct {
	Disable types.Bool `tfsdk:"disable" json:"disable,computed"`
	Healthy types.Bool `tfsdk:"healthy" json:"healthy,computed_optional"`
}

type LoadBalancerPoolsOriginSteeringDataSourceModel struct {
	Policy types.String `tfsdk:"policy" json:"policy,computed"`
}

type LoadBalancerPoolsOriginsDataSourceModel struct {
	Address          types.String                                   `tfsdk:"address" json:"address,computed_optional"`
	DisabledAt       timetypes.RFC3339                              `tfsdk:"disabled_at" json:"disabled_at,computed" format:"date-time"`
	Enabled          types.Bool                                     `tfsdk:"enabled" json:"enabled,computed"`
	Header           *LoadBalancerPoolsOriginsHeaderDataSourceModel `tfsdk:"header" json:"header,computed_optional"`
	Name             types.String                                   `tfsdk:"name" json:"name,computed_optional"`
	VirtualNetworkID types.String                                   `tfsdk:"virtual_network_id" json:"virtual_network_id,computed_optional"`
	Weight           types.Float64                                  `tfsdk:"weight" json:"weight,computed"`
}

type LoadBalancerPoolsOriginsHeaderDataSourceModel struct {
	Host *[]types.String `tfsdk:"host" json:"Host,computed_optional"`
}