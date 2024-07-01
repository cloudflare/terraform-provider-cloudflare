// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rate_limit

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RateLimitResultEnvelope struct {
	Result RateLimitModel `json:"result,computed"`
}

type RateLimitResultDataSourceEnvelope struct {
	Result RateLimitDataSourceModel `json:"result,computed"`
}

type RateLimitsResultDataSourceEnvelope struct {
	Result RateLimitsDataSourceModel `json:"result,computed"`
}

type RateLimitModel struct {
	ZoneIdentifier types.String `tfsdk:"zone_identifier" path:"zone_identifier"`
	ID             types.String `tfsdk:"id" path:"id"`
}

type RateLimitDataSourceModel struct {
}

type RateLimitsDataSourceModel struct {
}
