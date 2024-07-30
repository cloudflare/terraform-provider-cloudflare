// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rate_limit

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RateLimitResultDataSourceEnvelope struct {
	Result RateLimitDataSourceModel `json:"result,computed"`
}

type RateLimitResultListDataSourceEnvelope struct {
	Result *[]*RateLimitDataSourceModel `json:"result,computed"`
}

type RateLimitDataSourceModel struct {
	ZoneIdentifier types.String                       `tfsdk:"zone_identifier" path:"zone_identifier"`
	ID             types.String                       `tfsdk:"id" path:"id"`
	Action         *RateLimitActionDataSourceModel    `tfsdk:"action" json:"action"`
	Bypass         *[]*RateLimitBypassDataSourceModel `tfsdk:"bypass" json:"bypass"`
	Description    types.String                       `tfsdk:"description" json:"description"`
	Disabled       types.Bool                         `tfsdk:"disabled" json:"disabled"`
	Match          *RateLimitMatchDataSourceModel     `tfsdk:"match" json:"match"`
	Period         types.Float64                      `tfsdk:"period" json:"period"`
	Threshold      types.Float64                      `tfsdk:"threshold" json:"threshold"`
	Filter         *RateLimitFindOneByDataSourceModel `tfsdk:"filter"`
}

type RateLimitActionDataSourceModel struct {
	Mode     types.String                            `tfsdk:"mode" json:"mode"`
	Response *RateLimitActionResponseDataSourceModel `tfsdk:"response" json:"response"`
	Timeout  types.Float64                           `tfsdk:"timeout" json:"timeout"`
}

type RateLimitActionResponseDataSourceModel struct {
	Body        types.String `tfsdk:"body" json:"body"`
	ContentType types.String `tfsdk:"content_type" json:"content_type"`
}

type RateLimitBypassDataSourceModel struct {
	Name  types.String `tfsdk:"name" json:"name"`
	Value types.String `tfsdk:"value" json:"value"`
}

type RateLimitMatchDataSourceModel struct {
	Headers  *[]*RateLimitMatchHeadersDataSourceModel `tfsdk:"headers" json:"headers"`
	Request  *RateLimitMatchRequestDataSourceModel    `tfsdk:"request" json:"request"`
	Response *RateLimitMatchResponseDataSourceModel   `tfsdk:"response" json:"response"`
}

type RateLimitMatchHeadersDataSourceModel struct {
	Name  types.String `tfsdk:"name" json:"name"`
	Op    types.String `tfsdk:"op" json:"op"`
	Value types.String `tfsdk:"value" json:"value"`
}

type RateLimitMatchRequestDataSourceModel struct {
	Methods *[]types.String `tfsdk:"methods" json:"methods"`
	Schemes *[]types.String `tfsdk:"schemes" json:"schemes"`
	URL     types.String    `tfsdk:"url" json:"url"`
}

type RateLimitMatchResponseDataSourceModel struct {
	OriginTraffic types.Bool `tfsdk:"origin_traffic" json:"origin_traffic"`
}

type RateLimitFindOneByDataSourceModel struct {
	ZoneIdentifier types.String  `tfsdk:"zone_identifier" path:"zone_identifier"`
	Page           types.Float64 `tfsdk:"page" query:"page"`
	PerPage        types.Float64 `tfsdk:"per_page" query:"per_page"`
}
