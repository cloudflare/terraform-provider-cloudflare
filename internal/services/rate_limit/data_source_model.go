// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rate_limit

import (
	"github.com/cloudflare/cloudflare-go/v2/rate_limits"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RateLimitResultDataSourceEnvelope struct {
	Result RateLimitDataSourceModel `json:"result,computed"`
}

type RateLimitResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[RateLimitDataSourceModel] `json:"result,computed"`
}

type RateLimitDataSourceModel struct {
	ZoneIdentifier types.String                       `tfsdk:"zone_identifier" path:"zone_identifier"`
	ID             types.String                       `tfsdk:"id" path:"id,computed_optional"`
	Description    types.String                       `tfsdk:"description" json:"description"`
	Disabled       types.Bool                         `tfsdk:"disabled" json:"disabled"`
	Period         types.Float64                      `tfsdk:"period" json:"period"`
	Threshold      types.Float64                      `tfsdk:"threshold" json:"threshold"`
	Action         *RateLimitActionDataSourceModel    `tfsdk:"action" json:"action"`
	Bypass         *[]*RateLimitBypassDataSourceModel `tfsdk:"bypass" json:"bypass"`
	Match          *RateLimitMatchDataSourceModel     `tfsdk:"match" json:"match"`
	Filter         *RateLimitFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *RateLimitDataSourceModel) toListParams() (params rate_limits.RateLimitListParams, diags diag.Diagnostics) {
	params = rate_limits.RateLimitListParams{}

	return
}

type RateLimitActionDataSourceModel struct {
	Mode     types.String                            `tfsdk:"mode" json:"mode,computed_optional"`
	Response *RateLimitActionResponseDataSourceModel `tfsdk:"response" json:"response,computed_optional"`
	Timeout  types.Float64                           `tfsdk:"timeout" json:"timeout,computed_optional"`
}

type RateLimitActionResponseDataSourceModel struct {
	Body        types.String `tfsdk:"body" json:"body,computed_optional"`
	ContentType types.String `tfsdk:"content_type" json:"content_type,computed_optional"`
}

type RateLimitBypassDataSourceModel struct {
	Name  types.String `tfsdk:"name" json:"name,computed_optional"`
	Value types.String `tfsdk:"value" json:"value,computed_optional"`
}

type RateLimitMatchDataSourceModel struct {
	Headers  *[]*RateLimitMatchHeadersDataSourceModel `tfsdk:"headers" json:"headers,computed_optional"`
	Request  *RateLimitMatchRequestDataSourceModel    `tfsdk:"request" json:"request,computed_optional"`
	Response *RateLimitMatchResponseDataSourceModel   `tfsdk:"response" json:"response,computed_optional"`
}

type RateLimitMatchHeadersDataSourceModel struct {
	Name  types.String `tfsdk:"name" json:"name,computed_optional"`
	Op    types.String `tfsdk:"op" json:"op,computed_optional"`
	Value types.String `tfsdk:"value" json:"value,computed_optional"`
}

type RateLimitMatchRequestDataSourceModel struct {
	Methods *[]types.String `tfsdk:"methods" json:"methods,computed_optional"`
	Schemes *[]types.String `tfsdk:"schemes" json:"schemes,computed_optional"`
	URL     types.String    `tfsdk:"url" json:"url,computed_optional"`
}

type RateLimitMatchResponseDataSourceModel struct {
	OriginTraffic types.Bool `tfsdk:"origin_traffic" json:"origin_traffic,computed_optional"`
}

type RateLimitFindOneByDataSourceModel struct {
	ZoneIdentifier types.String `tfsdk:"zone_identifier" path:"zone_identifier"`
}
