// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_operation

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/api_gateway"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APIShieldOperationResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[APIShieldOperationDataSourceModel] `json:"result,computed"`
}

type APIShieldOperationDataSourceModel struct {
	Endpoint    types.String                                                        `tfsdk:"endpoint" json:"endpoint,computed"`
	Host        types.String                                                        `tfsdk:"host" json:"host,computed"`
	ID          types.String                                                        `tfsdk:"id" json:"id,computed"`
	LastUpdated timetypes.RFC3339                                                   `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
	Method      types.String                                                        `tfsdk:"method" json:"method,computed"`
	State       types.String                                                        `tfsdk:"state" json:"state,computed"`
	Origin      customfield.List[types.String]                                      `tfsdk:"origin" json:"origin,computed"`
	Features    customfield.NestedObject[APIShieldOperationFeaturesDataSourceModel] `tfsdk:"features" json:"features,computed"`
	Filter      *APIShieldOperationFindOneByDataSourceModel                         `tfsdk:"filter"`
}

func (m *APIShieldOperationDataSourceModel) toListParams(_ context.Context) (params api_gateway.DiscoveryOperationListParams, diags diag.Diagnostics) {
	mFilterHost := []string{}
	for _, item := range *m.Filter.Host {
		mFilterHost = append(mFilterHost, item.ValueString())
	}
	mFilterMethod := []string{}
	for _, item := range *m.Filter.Method {
		mFilterMethod = append(mFilterMethod, item.ValueString())
	}

	params = api_gateway.DiscoveryOperationListParams{
		ZoneID: cloudflare.F(m.Filter.ZoneID.ValueString()),
		Host:   cloudflare.F(mFilterHost),
		Method: cloudflare.F(mFilterMethod),
	}

	if !m.Filter.Diff.IsNull() {
		params.Diff = cloudflare.F(m.Filter.Diff.ValueBool())
	}
	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(api_gateway.DiscoveryOperationListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.Endpoint.IsNull() {
		params.Endpoint = cloudflare.F(m.Filter.Endpoint.ValueString())
	}
	if !m.Filter.Order.IsNull() {
		params.Order = cloudflare.F(api_gateway.DiscoveryOperationListParamsOrder(m.Filter.Order.ValueString()))
	}
	if !m.Filter.Origin.IsNull() {
		params.Origin = cloudflare.F(api_gateway.DiscoveryOperationListParamsOrigin(m.Filter.Origin.ValueString()))
	}
	if !m.Filter.State.IsNull() {
		params.State = cloudflare.F(api_gateway.DiscoveryOperationListParamsState(m.Filter.State.ValueString()))
	}

	return
}

type APIShieldOperationFeaturesDataSourceModel struct {
	TrafficStats customfield.NestedObject[APIShieldOperationFeaturesTrafficStatsDataSourceModel] `tfsdk:"traffic_stats" json:"traffic_stats,computed"`
}

type APIShieldOperationFeaturesTrafficStatsDataSourceModel struct {
	LastUpdated   timetypes.RFC3339 `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
	PeriodSeconds types.Int64       `tfsdk:"period_seconds" json:"period_seconds,computed"`
	Requests      types.Float64     `tfsdk:"requests" json:"requests,computed"`
}

type APIShieldOperationFindOneByDataSourceModel struct {
	ZoneID    types.String    `tfsdk:"zone_id" path:"zone_id,required"`
	Diff      types.Bool      `tfsdk:"diff" query:"diff,optional"`
	Direction types.String    `tfsdk:"direction" query:"direction,optional"`
	Endpoint  types.String    `tfsdk:"endpoint" query:"endpoint,optional"`
	Host      *[]types.String `tfsdk:"host" query:"host,optional"`
	Method    *[]types.String `tfsdk:"method" query:"method,optional"`
	Order     types.String    `tfsdk:"order" query:"order,optional"`
	Origin    types.String    `tfsdk:"origin" query:"origin,optional"`
	State     types.String    `tfsdk:"state" query:"state,optional"`
}
