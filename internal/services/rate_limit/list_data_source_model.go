// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rate_limit

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RateLimitsResultListDataSourceEnvelope struct {
	Result *[]*RateLimitsResultDataSourceModel `json:"result,computed"`
}

type RateLimitsDataSourceModel struct {
	ZoneIdentifier types.String                        `tfsdk:"zone_identifier" path:"zone_identifier"`
	MaxItems       types.Int64                         `tfsdk:"max_items"`
	Result         *[]*RateLimitsResultDataSourceModel `tfsdk:"result"`
}

type RateLimitsResultDataSourceModel struct {
	ID          types.String                        `tfsdk:"id" json:"id,computed"`
	Action      *RateLimitsActionDataSourceModel    `tfsdk:"action" json:"action"`
	Bypass      *[]*RateLimitsBypassDataSourceModel `tfsdk:"bypass" json:"bypass"`
	Description types.String                        `tfsdk:"description" json:"description"`
	Disabled    types.Bool                          `tfsdk:"disabled" json:"disabled"`
	Match       *RateLimitsMatchDataSourceModel     `tfsdk:"match" json:"match"`
	Period      types.Float64                       `tfsdk:"period" json:"period"`
	Threshold   types.Float64                       `tfsdk:"threshold" json:"threshold"`
}

type RateLimitsActionDataSourceModel struct {
	Mode     types.String                             `tfsdk:"mode" json:"mode"`
	Response *RateLimitsActionResponseDataSourceModel `tfsdk:"response" json:"response"`
	Timeout  types.Float64                            `tfsdk:"timeout" json:"timeout"`
}

type RateLimitsActionResponseDataSourceModel struct {
	Body        types.String `tfsdk:"body" json:"body"`
	ContentType types.String `tfsdk:"content_type" json:"content_type"`
}

type RateLimitsBypassDataSourceModel struct {
	Name  types.String `tfsdk:"name" json:"name"`
	Value types.String `tfsdk:"value" json:"value"`
}

type RateLimitsMatchDataSourceModel struct {
	Headers  *[]*RateLimitsMatchHeadersDataSourceModel `tfsdk:"headers" json:"headers"`
	Request  *RateLimitsMatchRequestDataSourceModel    `tfsdk:"request" json:"request"`
	Response *RateLimitsMatchResponseDataSourceModel   `tfsdk:"response" json:"response"`
}

type RateLimitsMatchHeadersDataSourceModel struct {
	Name  types.String `tfsdk:"name" json:"name"`
	Op    types.String `tfsdk:"op" json:"op"`
	Value types.String `tfsdk:"value" json:"value"`
}

type RateLimitsMatchRequestDataSourceModel struct {
	Methods *[]types.String `tfsdk:"methods" json:"methods"`
	Schemes *[]types.String `tfsdk:"schemes" json:"schemes"`
	URL     types.String    `tfsdk:"url" json:"url"`
}

type RateLimitsMatchResponseDataSourceModel struct {
	OriginTraffic types.Bool `tfsdk:"origin_traffic" json:"origin_traffic"`
}
