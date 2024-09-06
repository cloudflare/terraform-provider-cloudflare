// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package regional_tiered_cache

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/cache"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RegionalTieredCacheResultDataSourceEnvelope struct {
	Result RegionalTieredCacheDataSourceModel `json:"result,computed"`
}

type RegionalTieredCacheDataSourceModel struct {
	ZoneID     types.String                             `tfsdk:"zone_id" path:"zone_id,required"`
	ID         types.String                             `tfsdk:"id" json:"id,optional"`
	ModifiedOn timetypes.RFC3339                        `tfsdk:"modified_on" json:"modified_on,optional" format:"date-time"`
	Value      *RegionalTieredCacheValueDataSourceModel `tfsdk:"value" json:"value,optional"`
}

func (m *RegionalTieredCacheDataSourceModel) toReadParams(_ context.Context) (params cache.RegionalTieredCacheGetParams, diags diag.Diagnostics) {
	params = cache.RegionalTieredCacheGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type RegionalTieredCacheValueDataSourceModel struct {
	ID         types.String      `tfsdk:"id" json:"id,computed"`
	ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
}
