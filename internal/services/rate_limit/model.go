// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rate_limit

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RateLimitResultEnvelope struct {
	Result RateLimitModel `json:"result,computed"`
}

type RateLimitModel struct {
	ZoneIdentifier types.String          `tfsdk:"zone_identifier" path:"zone_identifier"`
	ID             types.String          `tfsdk:"id" path:"id"`
	Period         types.Float64         `tfsdk:"period" json:"period"`
	Threshold      types.Float64         `tfsdk:"threshold" json:"threshold"`
	Action         *RateLimitActionModel `tfsdk:"action" json:"action"`
	Match          *RateLimitMatchModel  `tfsdk:"match" json:"match"`
}

type RateLimitActionModel struct {
	Mode     types.String                  `tfsdk:"mode" json:"mode"`
	Response *RateLimitActionResponseModel `tfsdk:"response" json:"response"`
	Timeout  types.Float64                 `tfsdk:"timeout" json:"timeout"`
}

type RateLimitActionResponseModel struct {
	Body        types.String `tfsdk:"body" json:"body"`
	ContentType types.String `tfsdk:"content_type" json:"content_type"`
}

type RateLimitMatchModel struct {
	Headers  *[]*RateLimitMatchHeadersModel `tfsdk:"headers" json:"headers"`
	Request  *RateLimitMatchRequestModel    `tfsdk:"request" json:"request"`
	Response *RateLimitMatchResponseModel   `tfsdk:"response" json:"response"`
}

type RateLimitMatchHeadersModel struct {
	Name  types.String `tfsdk:"name" json:"name"`
	Op    types.String `tfsdk:"op" json:"op"`
	Value types.String `tfsdk:"value" json:"value"`
}

type RateLimitMatchRequestModel struct {
	Methods *[]types.String `tfsdk:"methods" json:"methods"`
	Schemes *[]types.String `tfsdk:"schemes" json:"schemes"`
	URL     types.String    `tfsdk:"url" json:"url"`
}

type RateLimitMatchResponseModel struct {
	OriginTraffic types.Bool `tfsdk:"origin_traffic" json:"origin_traffic"`
}
