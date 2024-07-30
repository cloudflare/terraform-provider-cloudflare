// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_operation

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APIShieldOperationResultListDataSourceEnvelope struct {
	Result *[]*APIShieldOperationDataSourceModel `json:"result,computed"`
}

type APIShieldOperationDataSourceModel struct {
	ID          types.String                                `tfsdk:"id" json:"id"`
	Endpoint    types.String                                `tfsdk:"endpoint" json:"endpoint"`
	Host        types.String                                `tfsdk:"host" json:"host"`
	LastUpdated timetypes.RFC3339                           `tfsdk:"last_updated" json:"last_updated"`
	Method      types.String                                `tfsdk:"method" json:"method"`
	Origin      *[]types.String                             `tfsdk:"origin" json:"origin"`
	State       types.String                                `tfsdk:"state" json:"state"`
	Features    *APIShieldOperationFeaturesDataSourceModel  `tfsdk:"features" json:"features"`
	Filter      *APIShieldOperationFindOneByDataSourceModel `tfsdk:"filter"`
}

type APIShieldOperationFeaturesDataSourceModel struct {
	TrafficStats *APIShieldOperationFeaturesTrafficStatsDataSourceModel `tfsdk:"traffic_stats" json:"traffic_stats"`
}

type APIShieldOperationFeaturesTrafficStatsDataSourceModel struct {
	LastUpdated   timetypes.RFC3339 `tfsdk:"last_updated" json:"last_updated,computed"`
	PeriodSeconds types.Int64       `tfsdk:"period_seconds" json:"period_seconds,computed"`
	Requests      types.Float64     `tfsdk:"requests" json:"requests,computed"`
}

type APIShieldOperationFindOneByDataSourceModel struct {
	ZoneID    types.String    `tfsdk:"zone_id" path:"zone_id"`
	Diff      types.Bool      `tfsdk:"diff" query:"diff"`
	Direction types.String    `tfsdk:"direction" query:"direction"`
	Endpoint  types.String    `tfsdk:"endpoint" query:"endpoint"`
	Host      *[]types.String `tfsdk:"host" query:"host"`
	Method    *[]types.String `tfsdk:"method" query:"method"`
	Order     types.String    `tfsdk:"order" query:"order"`
	Origin    types.String    `tfsdk:"origin" query:"origin"`
	Page      types.Int64     `tfsdk:"page" query:"page"`
	PerPage   types.Int64     `tfsdk:"per_page" query:"per_page"`
	State     types.String    `tfsdk:"state" query:"state"`
}
