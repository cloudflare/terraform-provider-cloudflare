// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tiered_cache

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/cache"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TieredCacheResultDataSourceEnvelope struct {
	Result TieredCacheDataSourceModel `json:"result,computed"`
}

type TieredCacheDataSourceModel struct {
	ID         types.String      `tfsdk:"id" path:"zone_id,computed"`
	ZoneID     types.String      `tfsdk:"zone_id" path:"zone_id,required"`
	Editable   types.Bool        `tfsdk:"editable" json:"editable,computed"`
	ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Value      types.String      `tfsdk:"value" json:"value,computed"`
}

func (m *TieredCacheDataSourceModel) toReadParams(_ context.Context) (params cache.SmartTieredCacheGetParams, diags diag.Diagnostics) {
	params = cache.SmartTieredCacheGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
