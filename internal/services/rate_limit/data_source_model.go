// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rate_limit

import (
	"context"

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
	ZoneIdentifier types.String                       `tfsdk:"zone_identifier" path:"zone_identifier,optional"`
	ID             types.String                       `tfsdk:"id" path:"id,computed_optional"`
	Description    types.String                       `tfsdk:"description" json:"description,optional"`
	Disabled       types.Bool                         `tfsdk:"disabled" json:"disabled,optional"`
	Period         types.Float64                      `tfsdk:"period" json:"period,optional"`
	Threshold      types.Float64                      `tfsdk:"threshold" json:"threshold,optional"`
	Action         *RateLimitActionDataSourceModel    `tfsdk:"action" json:"action,optional"`
	Bypass         *[]*RateLimitBypassDataSourceModel `tfsdk:"bypass" json:"bypass,optional"`
	Match          *RateLimitMatchDataSourceModel     `tfsdk:"match" json:"match,optional"`
	Filter         *RateLimitFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *RateLimitDataSourceModel) toListParams(_ context.Context) (params rate_limits.RateLimitListParams, diags diag.Diagnostics) {
	params = rate_limits.RateLimitListParams{}

	return
}

type RateLimitActionDataSourceModel struct {
	Mode     types.String                                                     `tfsdk:"mode" json:"mode,computed"`
	Response customfield.NestedObject[RateLimitActionResponseDataSourceModel] `tfsdk:"response" json:"response,computed"`
	Timeout  types.Float64                                                    `tfsdk:"timeout" json:"timeout,computed"`
}

type RateLimitActionResponseDataSourceModel struct {
	Body        types.String `tfsdk:"body" json:"body,computed"`
	ContentType types.String `tfsdk:"content_type" json:"content_type,computed"`
}

type RateLimitBypassDataSourceModel struct {
	Name  types.String `tfsdk:"name" json:"name,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}

type RateLimitMatchDataSourceModel struct {
	Headers  customfield.NestedObjectList[RateLimitMatchHeadersDataSourceModel] `tfsdk:"headers" json:"headers,computed"`
	Request  customfield.NestedObject[RateLimitMatchRequestDataSourceModel]     `tfsdk:"request" json:"request,computed"`
	Response customfield.NestedObject[RateLimitMatchResponseDataSourceModel]    `tfsdk:"response" json:"response,computed"`
}

type RateLimitMatchHeadersDataSourceModel struct {
	Name  types.String `tfsdk:"name" json:"name,computed"`
	Op    types.String `tfsdk:"op" json:"op,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}

type RateLimitMatchRequestDataSourceModel struct {
	Methods customfield.List[types.String] `tfsdk:"methods" json:"methods,computed"`
	Schemes customfield.List[types.String] `tfsdk:"schemes" json:"schemes,computed"`
	URL     types.String                   `tfsdk:"url" json:"url,computed"`
}

type RateLimitMatchResponseDataSourceModel struct {
	OriginTraffic types.Bool `tfsdk:"origin_traffic" json:"origin_traffic,computed"`
}

type RateLimitFindOneByDataSourceModel struct {
	ZoneIdentifier types.String `tfsdk:"zone_identifier" path:"zone_identifier,required"`
}
