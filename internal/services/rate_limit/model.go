// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rate_limit

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RateLimitResultEnvelope struct {
	Result RateLimitModel `json:"result"`
}

type RateLimitModel struct {
	ZoneIdentifier types.String          `tfsdk:"zone_identifier" path:"zone_identifier,required"`
	ID             types.String          `tfsdk:"id" path:"id,computed_optional"`
	Period         types.Float64         `tfsdk:"period" json:"period,required"`
	Threshold      types.Float64         `tfsdk:"threshold" json:"threshold,required"`
	Action         *RateLimitActionModel `tfsdk:"action" json:"action,required"`
	Match          *RateLimitMatchModel  `tfsdk:"match" json:"match,required"`
}

type RateLimitActionModel struct {
	Mode     types.String                                           `tfsdk:"mode" json:"mode,computed_optional"`
	Response customfield.NestedObject[RateLimitActionResponseModel] `tfsdk:"response" json:"response,computed_optional"`
	Timeout  types.Float64                                          `tfsdk:"timeout" json:"timeout,computed_optional"`
}

type RateLimitActionResponseModel struct {
	Body        types.String `tfsdk:"body" json:"body,computed_optional"`
	ContentType types.String `tfsdk:"content_type" json:"content_type,computed_optional"`
}

type RateLimitMatchModel struct {
	Headers  customfield.NestedObjectList[RateLimitMatchHeadersModel] `tfsdk:"headers" json:"headers,computed_optional"`
	Request  customfield.NestedObject[RateLimitMatchRequestModel]     `tfsdk:"request" json:"request,computed_optional"`
	Response customfield.NestedObject[RateLimitMatchResponseModel]    `tfsdk:"response" json:"response,computed_optional"`
}

type RateLimitMatchHeadersModel struct {
	Name  types.String `tfsdk:"name" json:"name,computed_optional"`
	Op    types.String `tfsdk:"op" json:"op,computed_optional"`
	Value types.String `tfsdk:"value" json:"value,computed_optional"`
}

type RateLimitMatchRequestModel struct {
	Methods customfield.List[types.String] `tfsdk:"methods" json:"methods,computed_optional"`
	Schemes customfield.List[types.String] `tfsdk:"schemes" json:"schemes,computed_optional"`
	URL     types.String                   `tfsdk:"url" json:"url,computed_optional"`
}

type RateLimitMatchResponseModel struct {
	OriginTraffic types.Bool `tfsdk:"origin_traffic" json:"origin_traffic,computed_optional"`
}
