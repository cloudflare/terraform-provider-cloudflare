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
	Action      *RateLimitsItemsActionDataSourceModel    `tfsdk:"action" json:"action"`
	Bypass      *[]*RateLimitsItemsBypassDataSourceModel `tfsdk:"bypass" json:"bypass"`
	Description types.String                             `tfsdk:"description" json:"description"`
	Disabled    types.Bool                               `tfsdk:"disabled" json:"disabled"`
	Match       *RateLimitsItemsMatchDataSourceModel     `tfsdk:"match" json:"match"`
	Period      types.Float64                            `tfsdk:"period" json:"period"`
	Threshold   types.Float64                            `tfsdk:"threshold" json:"threshold"`
}

type RateLimitsItemsActionDataSourceModel struct {
	Mode     types.String                                  `tfsdk:"mode" json:"mode"`
	Response *RateLimitsItemsActionResponseDataSourceModel `tfsdk:"response" json:"response"`
	Timeout  types.Float64                                 `tfsdk:"timeout" json:"timeout"`
}

type RateLimitsItemsActionResponseDataSourceModel struct {
	Body        types.String `tfsdk:"body" json:"body"`
	ContentType types.String `tfsdk:"content_type" json:"content_type"`
}

type RateLimitsItemsBypassDataSourceModel struct {
	Name  types.String `tfsdk:"name" json:"name"`
	Value types.String `tfsdk:"value" json:"value"`
}

type RateLimitsItemsMatchDataSourceModel struct {
	Headers  *[]*RateLimitsItemsMatchHeadersDataSourceModel `tfsdk:"headers" json:"headers"`
	Request  *RateLimitsItemsMatchRequestDataSourceModel    `tfsdk:"request" json:"request"`
	Response *RateLimitsItemsMatchResponseDataSourceModel   `tfsdk:"response" json:"response"`
}

type RateLimitsItemsMatchHeadersDataSourceModel struct {
	Name  types.String `tfsdk:"name" json:"name"`
	Op    types.String `tfsdk:"op" json:"op"`
	Value types.String `tfsdk:"value" json:"value"`
}

type RateLimitsItemsMatchRequestDataSourceModel struct {
	Methods *[]types.String `tfsdk:"methods" json:"methods"`
	Schemes *[]types.String `tfsdk:"schemes" json:"schemes"`
	URL     types.String    `tfsdk:"url" json:"url"`
}

type RateLimitsItemsMatchResponseDataSourceModel struct {
	OriginTraffic types.Bool `tfsdk:"origin_traffic" json:"origin_traffic"`
}
