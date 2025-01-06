// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_cache_variants

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/cache"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneCacheVariantsResultDataSourceEnvelope struct {
	Result ZoneCacheVariantsDataSourceModel `json:"result,computed"`
}

type ZoneCacheVariantsDataSourceModel struct {
	ZoneID     types.String      `tfsdk:"zone_id" path:"zone_id,required"`
	Editable   types.Bool        `tfsdk:"editable" json:"editable,computed"`
	ID         types.String      `tfsdk:"id" json:"id,computed"`
	ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Value      types.String      `tfsdk:"value" json:"value,computed"`
}

func (m *ZoneCacheVariantsDataSourceModel) toReadParams(_ context.Context) (params cache.VariantGetParams, diags diag.Diagnostics) {
	params = cache.VariantGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
