// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package regional_tiered_cache

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RegionalTieredCacheResultEnvelope struct {
	Result RegionalTieredCacheModel `json:"result,computed"`
}

type RegionalTieredCacheModel struct {
	ID     types.String `tfsdk:"id" json:"id,computed"`
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id"`
	Value  types.String `tfsdk:"value" json:"value"`
}
