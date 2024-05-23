// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tiered_caching

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TieredCachingResultEnvelope struct {
	Result TieredCachingModel `json:"result,computed"`
}

type TieredCachingModel struct {
	ID     types.String `tfsdk:"id" json:"id"`
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id"`
	Value  types.String `tfsdk:"value" json:"value"`
}
