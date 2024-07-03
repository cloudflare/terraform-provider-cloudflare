// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rate_limit

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RateLimitsResultListDataSourceEnvelope struct {
	Result *[]*RateLimitsItemsDataSourceModel `json:"result,computed"`
}

type RateLimitsDataSourceModel struct {
	ZoneIdentifier types.String                       `tfsdk:"zone_identifier" path:"zone_identifier"`
	Page           types.Float64                      `tfsdk:"page" query:"page"`
	PerPage        types.Float64                      `tfsdk:"per_page" query:"per_page"`
	MaxItems       types.Int64                        `tfsdk:"max_items"`
	Items          *[]*RateLimitsItemsDataSourceModel `tfsdk:"items"`
}

type RateLimitsItemsDataSourceModel struct {
	ID          types.String                             `tfsdk:"id" json:"id,computed"`
	Bypass      *[]*RateLimitsItemsBypassDataSourceModel `tfsdk:"bypass" json:"bypass,computed"`
	Description types.String                             `tfsdk:"description" json:"description,computed"`
	Disabled    types.Bool                               `tfsdk:"disabled" json:"disabled,computed"`
	Period      types.Float64                            `tfsdk:"period" json:"period,computed"`
	Threshold   types.Float64                            `tfsdk:"threshold" json:"threshold,computed"`
}

type RateLimitsItemsActionDataSourceModel struct {
	Mode    types.String  `tfsdk:"mode" json:"mode,computed"`
	Timeout types.Float64 `tfsdk:"timeout" json:"timeout,computed"`
}

type RateLimitsItemsActionResponseDataSourceModel struct {
	Body        types.String `tfsdk:"body" json:"body,computed"`
	ContentType types.String `tfsdk:"content_type" json:"content_type,computed"`
}

type RateLimitsItemsBypassDataSourceModel struct {
	Name  types.String `tfsdk:"name" json:"name,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}

type RateLimitsItemsMatchDataSourceModel struct {
	Headers *[]*RateLimitsItemsMatchHeadersDataSourceModel `tfsdk:"headers" json:"headers,computed"`
}

type RateLimitsItemsMatchHeadersDataSourceModel struct {
	Name  types.String `tfsdk:"name" json:"name,computed"`
	Op    types.String `tfsdk:"op" json:"op,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}

type RateLimitsItemsMatchRequestDataSourceModel struct {
	Methods *[]types.String `tfsdk:"methods" json:"methods,computed"`
	Schemes *[]types.String `tfsdk:"schemes" json:"schemes,computed"`
	URL     types.String    `tfsdk:"url" json:"url,computed"`
}

type RateLimitsItemsMatchResponseDataSourceModel struct {
	OriginTraffic types.Bool `tfsdk:"origin_traffic" json:"origin_traffic,computed"`
}
