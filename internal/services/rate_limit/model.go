// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rate_limit

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RateLimitResultEnvelope struct {
	Result RateLimitModel `json:"result"`
}

type RateLimitModel struct {
	ID          types.String                                       `tfsdk:"id" json:"id,computed"`
	ZoneID      types.String                                       `tfsdk:"zone_id" path:"zone_id,required"`
	Period      types.Float64                                      `tfsdk:"period" json:"period,required"`
	Threshold   types.Float64                                      `tfsdk:"threshold" json:"threshold,required"`
	Action      *RateLimitActionModel                              `tfsdk:"action" json:"action,required"`
	Match       *RateLimitMatchModel                               `tfsdk:"match" json:"match,required"`
	Description types.String                                       `tfsdk:"description" json:"description,computed"`
	Disabled    types.Bool                                         `tfsdk:"disabled" json:"disabled,computed"`
	Bypass      customfield.NestedObjectList[RateLimitBypassModel] `tfsdk:"bypass" json:"bypass,computed"`
}

func (m RateLimitModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m RateLimitModel) MarshalJSONForUpdate(state RateLimitModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type RateLimitActionModel struct {
	Mode     types.String                  `tfsdk:"mode" json:"mode,optional"`
	Response *RateLimitActionResponseModel `tfsdk:"response" json:"response,optional"`
	Timeout  types.Float64                 `tfsdk:"timeout" json:"timeout,optional"`
}

type RateLimitActionResponseModel struct {
	Body        types.String `tfsdk:"body" json:"body,optional"`
	ContentType types.String `tfsdk:"content_type" json:"content_type,optional"`
}

type RateLimitMatchModel struct {
	Headers  *[]*RateLimitMatchHeadersModel `tfsdk:"headers" json:"headers,optional"`
	Request  *RateLimitMatchRequestModel    `tfsdk:"request" json:"request,optional"`
	Response *RateLimitMatchResponseModel   `tfsdk:"response" json:"response,optional"`
}

type RateLimitMatchHeadersModel struct {
	Name  types.String `tfsdk:"name" json:"name,optional"`
	Op    types.String `tfsdk:"op" json:"op,optional"`
	Value types.String `tfsdk:"value" json:"value,optional"`
}

type RateLimitMatchRequestModel struct {
	Methods *[]types.String `tfsdk:"methods" json:"methods,optional"`
	Schemes *[]types.String `tfsdk:"schemes" json:"schemes,optional"`
	URL     types.String    `tfsdk:"url" json:"url,optional"`
}

type RateLimitMatchResponseModel struct {
	OriginTraffic types.Bool `tfsdk:"origin_traffic" json:"origin_traffic,optional"`
}

type RateLimitBypassModel struct {
	Name  types.String `tfsdk:"name" json:"name,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}
