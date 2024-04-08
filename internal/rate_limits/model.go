// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rate_limits

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RateLimitsResultEnvelope struct {
	Result RateLimitsModel `json:"result,computed"`
}

type RateLimitsModel struct {
	ZoneIdentifier types.String `tfsdk:"zone_identifier" path:"zone_identifier"`
	ID             types.String `tfsdk:"id" path:"id"`
}
