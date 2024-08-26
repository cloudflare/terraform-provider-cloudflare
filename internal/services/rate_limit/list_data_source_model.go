// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rate_limit

import (
	"github.com/cloudflare/cloudflare-go/v2/rate_limits"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RateLimitsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[RateLimitsResultDataSourceModel] `json:"result,computed"`
}

type RateLimitsDataSourceModel struct {
	ZoneIdentifier types.String                                                  `tfsdk:"zone_identifier" path:"zone_identifier"`
	MaxItems       types.Int64                                                   `tfsdk:"max_items"`
	Result         customfield.NestedObjectList[RateLimitsResultDataSourceModel] `tfsdk:"result"`
}

func (m *RateLimitsDataSourceModel) toListParams() (params rate_limits.RateLimitListParams, diags diag.Diagnostics) {
	params = rate_limits.RateLimitListParams{}

	return
}

type RateLimitsResultDataSourceModel struct {
	ID          types.String                        `tfsdk:"id" json:"id,computed"`
	Action      *RateLimitsActionDataSourceModel    `tfsdk:"action" json:"action,computed_optional"`
	Bypass      *[]*RateLimitsBypassDataSourceModel `tfsdk:"bypass" json:"bypass,computed_optional"`
	Description types.String                        `tfsdk:"description" json:"description,computed_optional"`
	Disabled    types.Bool                          `tfsdk:"disabled" json:"disabled,computed_optional"`
	Match       *RateLimitsMatchDataSourceModel     `tfsdk:"match" json:"match,computed_optional"`
	Period      types.Float64                       `tfsdk:"period" json:"period,computed_optional"`
	Threshold   types.Float64                       `tfsdk:"threshold" json:"threshold,computed_optional"`
}

type RateLimitsActionDataSourceModel struct {
	Mode     types.String                             `tfsdk:"mode" json:"mode,computed_optional"`
	Response *RateLimitsActionResponseDataSourceModel `tfsdk:"response" json:"response,computed_optional"`
	Timeout  types.Float64                            `tfsdk:"timeout" json:"timeout,computed_optional"`
}

type RateLimitsActionResponseDataSourceModel struct {
	Body        types.String `tfsdk:"body" json:"body,computed_optional"`
	ContentType types.String `tfsdk:"content_type" json:"content_type,computed_optional"`
}

type RateLimitsBypassDataSourceModel struct {
	Name  types.String `tfsdk:"name" json:"name,computed_optional"`
	Value types.String `tfsdk:"value" json:"value,computed_optional"`
}

type RateLimitsMatchDataSourceModel struct {
	Headers  *[]*RateLimitsMatchHeadersDataSourceModel `tfsdk:"headers" json:"headers,computed_optional"`
	Request  *RateLimitsMatchRequestDataSourceModel    `tfsdk:"request" json:"request,computed_optional"`
	Response *RateLimitsMatchResponseDataSourceModel   `tfsdk:"response" json:"response,computed_optional"`
}

type RateLimitsMatchHeadersDataSourceModel struct {
	Name  types.String `tfsdk:"name" json:"name,computed_optional"`
	Op    types.String `tfsdk:"op" json:"op,computed_optional"`
	Value types.String `tfsdk:"value" json:"value,computed_optional"`
}

type RateLimitsMatchRequestDataSourceModel struct {
	Methods *[]types.String `tfsdk:"methods" json:"methods,computed_optional"`
	Schemes *[]types.String `tfsdk:"schemes" json:"schemes,computed_optional"`
	URL     types.String    `tfsdk:"url" json:"url,computed_optional"`
}

type RateLimitsMatchResponseDataSourceModel struct {
	OriginTraffic types.Bool `tfsdk:"origin_traffic" json:"origin_traffic,computed_optional"`
}
