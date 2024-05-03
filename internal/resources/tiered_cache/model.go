// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tiered_cache

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TieredCacheResultEnvelope struct {
	Result TieredCacheModel `json:"result,computed"`
}

type TieredCacheModel struct {
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id"`
	Value  types.String `tfsdk:"value" json:"value"`
}
