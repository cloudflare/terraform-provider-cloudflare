// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rate_limit

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3/rate_limits"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RateLimitsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[RateLimitsResultDataSourceModel] `json:"result,computed"`
}

type RateLimitsDataSourceModel struct {
	ZoneIdentifier types.String                                                  `tfsdk:"zone_identifier" path:"zone_identifier,required"`
	MaxItems       types.Int64                                                   `tfsdk:"max_items"`
	Result         customfield.NestedObjectList[RateLimitsResultDataSourceModel] `tfsdk:"result"`
}

func (m *RateLimitsDataSourceModel) toListParams(_ context.Context) (params rate_limits.RateLimitListParams, diags diag.Diagnostics) {
	params = rate_limits.RateLimitListParams{}

	return
}

type RateLimitsResultDataSourceModel struct {
	ID          types.String                                                  `tfsdk:"id" json:"id,computed"`
	Action      customfield.NestedObject[RateLimitsActionDataSourceModel]     `tfsdk:"action" json:"action,computed"`
	Bypass      customfield.NestedObjectList[RateLimitsBypassDataSourceModel] `tfsdk:"bypass" json:"bypass,computed"`
	Description types.String                                                  `tfsdk:"description" json:"description,computed"`
	Disabled    types.Bool                                                    `tfsdk:"disabled" json:"disabled,computed"`
	Match       customfield.NestedObject[RateLimitsMatchDataSourceModel]      `tfsdk:"match" json:"match,computed"`
	Period      types.Float64                                                 `tfsdk:"period" json:"period,computed"`
	Threshold   types.Float64                                                 `tfsdk:"threshold" json:"threshold,computed"`
}

type RateLimitsActionDataSourceModel struct {
	Mode     types.String                                                      `tfsdk:"mode" json:"mode,computed"`
	Response customfield.NestedObject[RateLimitsActionResponseDataSourceModel] `tfsdk:"response" json:"response,computed"`
	Timeout  types.Float64                                                     `tfsdk:"timeout" json:"timeout,computed"`
}

type RateLimitsActionResponseDataSourceModel struct {
	Body        types.String `tfsdk:"body" json:"body,computed"`
	ContentType types.String `tfsdk:"content_type" json:"content_type,computed"`
}

type RateLimitsBypassDataSourceModel struct {
	Name  types.String `tfsdk:"name" json:"name,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}

type RateLimitsMatchDataSourceModel struct {
	Headers  customfield.NestedObjectList[RateLimitsMatchHeadersDataSourceModel] `tfsdk:"headers" json:"headers,computed"`
	Request  customfield.NestedObject[RateLimitsMatchRequestDataSourceModel]     `tfsdk:"request" json:"request,computed"`
	Response customfield.NestedObject[RateLimitsMatchResponseDataSourceModel]    `tfsdk:"response" json:"response,computed"`
}

type RateLimitsMatchHeadersDataSourceModel struct {
	Name  types.String `tfsdk:"name" json:"name,computed"`
	Op    types.String `tfsdk:"op" json:"op,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}

type RateLimitsMatchRequestDataSourceModel struct {
	Methods customfield.List[types.String] `tfsdk:"methods" json:"methods,computed"`
	Schemes customfield.List[types.String] `tfsdk:"schemes" json:"schemes,computed"`
	URL     types.String                   `tfsdk:"url" json:"url,computed"`
}

type RateLimitsMatchResponseDataSourceModel struct {
	OriginTraffic types.Bool `tfsdk:"origin_traffic" json:"origin_traffic,computed"`
}
