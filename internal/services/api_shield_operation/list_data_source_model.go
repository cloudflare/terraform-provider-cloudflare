// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_operation

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/api_gateway"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APIShieldOperationsResultListDataSourceEnvelope struct {
	Result *[]*APIShieldOperationsResultDataSourceModel `json:"result,computed"`
}

type APIShieldOperationsDataSourceModel struct {
	ZoneID    types.String                                 `tfsdk:"zone_id" path:"zone_id"`
	Diff      types.Bool                                   `tfsdk:"diff" query:"diff"`
	Direction types.String                                 `tfsdk:"direction" query:"direction"`
	Endpoint  types.String                                 `tfsdk:"endpoint" query:"endpoint"`
	Order     types.String                                 `tfsdk:"order" query:"order"`
	Origin    types.String                                 `tfsdk:"origin" query:"origin"`
	State     types.String                                 `tfsdk:"state" query:"state"`
	Host      *[]types.String                              `tfsdk:"host" query:"host"`
	Method    *[]types.String                              `tfsdk:"method" query:"method"`
	MaxItems  types.Int64                                  `tfsdk:"max_items"`
	Result    *[]*APIShieldOperationsResultDataSourceModel `tfsdk:"result"`
}

func (m *APIShieldOperationsDataSourceModel) toListParams() (params api_gateway.DiscoveryOperationListParams, diags diag.Diagnostics) {
	mHost := []string{}
	for _, item := range *m.Host {
		mHost = append(mHost, item.ValueString())
	}
	mMethod := []string{}
	for _, item := range *m.Method {
		mMethod = append(mMethod, item.ValueString())
	}

	params = api_gateway.DiscoveryOperationListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
		Host:   cloudflare.F(mHost),
		Method: cloudflare.F(mMethod),
	}

	if !m.Diff.IsNull() {
		params.Diff = cloudflare.F(m.Diff.ValueBool())
	}
	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(api_gateway.DiscoveryOperationListParamsDirection(m.Direction.ValueString()))
	}
	if !m.Endpoint.IsNull() {
		params.Endpoint = cloudflare.F(m.Endpoint.ValueString())
	}
	if !m.Order.IsNull() {
		params.Order = cloudflare.F(api_gateway.DiscoveryOperationListParamsOrder(m.Order.ValueString()))
	}
	if !m.Origin.IsNull() {
		params.Origin = cloudflare.F(api_gateway.DiscoveryOperationListParamsOrigin(m.Origin.ValueString()))
	}
	if !m.State.IsNull() {
		params.State = cloudflare.F(api_gateway.DiscoveryOperationListParamsState(m.State.ValueString()))
	}

	return
}

type APIShieldOperationsResultDataSourceModel struct {
	ID          types.String                                                         `tfsdk:"id" json:"id,computed"`
	Endpoint    types.String                                                         `tfsdk:"endpoint" json:"endpoint,computed"`
	Host        types.String                                                         `tfsdk:"host" json:"host,computed"`
	LastUpdated timetypes.RFC3339                                                    `tfsdk:"last_updated" json:"last_updated,computed"`
	Method      types.String                                                         `tfsdk:"method" json:"method,computed"`
	Origin      *[]types.String                                                      `tfsdk:"origin" json:"origin,computed"`
	State       types.String                                                         `tfsdk:"state" json:"state,computed"`
	Features    customfield.NestedObject[APIShieldOperationsFeaturesDataSourceModel] `tfsdk:"features" json:"features,computed"`
}

type APIShieldOperationsFeaturesDataSourceModel struct {
	TrafficStats *APIShieldOperationsFeaturesTrafficStatsDataSourceModel `tfsdk:"traffic_stats" json:"traffic_stats"`
}

type APIShieldOperationsFeaturesTrafficStatsDataSourceModel struct {
	LastUpdated   timetypes.RFC3339 `tfsdk:"last_updated" json:"last_updated,computed"`
	PeriodSeconds types.Int64       `tfsdk:"period_seconds" json:"period_seconds,computed"`
	Requests      types.Float64     `tfsdk:"requests" json:"requests,computed"`
}
