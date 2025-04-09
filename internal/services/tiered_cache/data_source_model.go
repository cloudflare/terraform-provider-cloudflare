// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tiered_cache

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/cache"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TieredCacheResultDataSourceEnvelope struct {
	Result TieredCacheDataSourceModel `json:"result,computed"`
}

type TieredCacheDataSourceModel struct {
	ZoneID     types.String      `tfsdk:"zone_id" path:"zone_id,required"`
	Editable   types.Bool        `tfsdk:"editable" json:"editable,computed"`
	ID         types.String      `tfsdk:"id" json:"id,computed"`
	ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Value      types.String      `tfsdk:"value" json:"value,computed"`
}

func (m *TieredCacheDataSourceModel) toReadParams(_ context.Context) (params cache.SmartTieredCacheGetParams, diags diag.Diagnostics) {
	params = cache.SmartTieredCacheGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
