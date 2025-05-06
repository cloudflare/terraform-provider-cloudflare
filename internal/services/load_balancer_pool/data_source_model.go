// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer_pool

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/load_balancers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LoadBalancerPoolResultDataSourceEnvelope struct {
	Result LoadBalancerPoolDataSourceModel `json:"result,computed"`
}

type LoadBalancerPoolDataSourceModel struct {
	ID                 types.String                                                                `tfsdk:"id" path:"pool_id,computed"`
	PoolID             types.String                                                                `tfsdk:"pool_id" path:"pool_id,optional"`
	AccountID          types.String                                                                `tfsdk:"account_id" path:"account_id,required"`
	CreatedOn          types.String                                                                `tfsdk:"created_on" json:"created_on,computed"`
	Description        types.String                                                                `tfsdk:"description" json:"description,computed"`
	DisabledAt         timetypes.RFC3339                                                           `tfsdk:"disabled_at" json:"disabled_at,computed" format:"date-time"`
	Enabled            types.Bool                                                                  `tfsdk:"enabled" json:"enabled,computed"`
	Latitude           types.Float64                                                               `tfsdk:"latitude" json:"latitude,computed"`
	Longitude          types.Float64                                                               `tfsdk:"longitude" json:"longitude,computed"`
	MinimumOrigins     types.Int64                                                                 `tfsdk:"minimum_origins" json:"minimum_origins,computed"`
	ModifiedOn         types.String                                                                `tfsdk:"modified_on" json:"modified_on,computed"`
	Monitor            types.String                                                                `tfsdk:"monitor" json:"monitor,computed"`
	Name               types.String                                                                `tfsdk:"name" json:"name,computed"`
	NotificationEmail  types.String                                                                `tfsdk:"notification_email" json:"notification_email,computed"`
	CheckRegions       customfield.List[types.String]                                              `tfsdk:"check_regions" json:"check_regions,computed"`
	Networks           customfield.List[types.String]                                              `tfsdk:"networks" json:"networks,computed"`
	LoadShedding       customfield.NestedObject[LoadBalancerPoolLoadSheddingDataSourceModel]       `tfsdk:"load_shedding" json:"load_shedding,computed"`
	NotificationFilter customfield.NestedObject[LoadBalancerPoolNotificationFilterDataSourceModel] `tfsdk:"notification_filter" json:"notification_filter,computed"`
	OriginSteering     customfield.NestedObject[LoadBalancerPoolOriginSteeringDataSourceModel]     `tfsdk:"origin_steering" json:"origin_steering,computed"`
	Origins            customfield.NestedObjectList[LoadBalancerPoolOriginsDataSourceModel]        `tfsdk:"origins" json:"origins,computed"`
	Filter             *LoadBalancerPoolFindOneByDataSourceModel                                   `tfsdk:"filter"`
}

func (m *LoadBalancerPoolDataSourceModel) toReadParams(_ context.Context) (params load_balancers.PoolGetParams, diags diag.Diagnostics) {
	params = load_balancers.PoolGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *LoadBalancerPoolDataSourceModel) toListParams(_ context.Context) (params load_balancers.PoolListParams, diags diag.Diagnostics) {
	params = load_balancers.PoolListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Filter.Monitor.IsNull() {
		params.Monitor = cloudflare.F(m.Filter.Monitor.ValueString())
	}

	return
}

type LoadBalancerPoolLoadSheddingDataSourceModel struct {
	DefaultPercent types.Float64 `tfsdk:"default_percent" json:"default_percent,computed"`
	DefaultPolicy  types.String  `tfsdk:"default_policy" json:"default_policy,computed"`
	SessionPercent types.Float64 `tfsdk:"session_percent" json:"session_percent,computed"`
	SessionPolicy  types.String  `tfsdk:"session_policy" json:"session_policy,computed"`
}

type LoadBalancerPoolNotificationFilterDataSourceModel struct {
	Origin customfield.NestedObject[LoadBalancerPoolNotificationFilterOriginDataSourceModel] `tfsdk:"origin" json:"origin,computed"`
	Pool   customfield.NestedObject[LoadBalancerPoolNotificationFilterPoolDataSourceModel]   `tfsdk:"pool" json:"pool,computed"`
}

type LoadBalancerPoolNotificationFilterOriginDataSourceModel struct {
	Disable types.Bool `tfsdk:"disable" json:"disable,computed"`
	Healthy types.Bool `tfsdk:"healthy" json:"healthy,computed"`
}

type LoadBalancerPoolNotificationFilterPoolDataSourceModel struct {
	Disable types.Bool `tfsdk:"disable" json:"disable,computed"`
	Healthy types.Bool `tfsdk:"healthy" json:"healthy,computed"`
}

type LoadBalancerPoolOriginSteeringDataSourceModel struct {
	Policy types.String `tfsdk:"policy" json:"policy,computed"`
}

type LoadBalancerPoolOriginsDataSourceModel struct {
	Address          types.String                                                           `tfsdk:"address" json:"address,computed"`
	DisabledAt       timetypes.RFC3339                                                      `tfsdk:"disabled_at" json:"disabled_at,computed" format:"date-time"`
	Enabled          types.Bool                                                             `tfsdk:"enabled" json:"enabled,computed"`
	Header           customfield.NestedObject[LoadBalancerPoolOriginsHeaderDataSourceModel] `tfsdk:"header" json:"header,computed"`
	Name             types.String                                                           `tfsdk:"name" json:"name,computed"`
	Port             types.Int64                                                            `tfsdk:"port" json:"port,computed"`
	VirtualNetworkID types.String                                                           `tfsdk:"virtual_network_id" json:"virtual_network_id,computed"`
	Weight           types.Float64                                                          `tfsdk:"weight" json:"weight,computed"`
}

type LoadBalancerPoolOriginsHeaderDataSourceModel struct {
	Host customfield.List[types.String] `tfsdk:"host" json:"Host,computed"`
}

type LoadBalancerPoolFindOneByDataSourceModel struct {
	Monitor types.String `tfsdk:"monitor" query:"monitor,optional"`
}
