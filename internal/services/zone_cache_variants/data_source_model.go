// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_cache_variants

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/cache"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneCacheVariantsResultDataSourceEnvelope struct {
	Result ZoneCacheVariantsDataSourceModel `json:"result,computed"`
}

type ZoneCacheVariantsDataSourceModel struct {
	ZoneID     types.String                           `tfsdk:"zone_id" path:"zone_id"`
	ID         types.String                           `tfsdk:"id" json:"id"`
	ModifiedOn timetypes.RFC3339                      `tfsdk:"modified_on" json:"modified_on" format:"date-time"`
	Value      *ZoneCacheVariantsValueDataSourceModel `tfsdk:"value" json:"value"`
}

func (m *ZoneCacheVariantsDataSourceModel) toReadParams() (params cache.VariantGetParams, diags diag.Diagnostics) {
	params = cache.VariantGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type ZoneCacheVariantsValueDataSourceModel struct {
	AVIF types.List `tfsdk:"avif" json:"avif,computed"`
	BMP  types.List `tfsdk:"bmp" json:"bmp,computed"`
	GIF  types.List `tfsdk:"gif" json:"gif,computed"`
	JP2  types.List `tfsdk:"jp2" json:"jp2,computed"`
	JPEG types.List `tfsdk:"jpeg" json:"jpeg,computed"`
	JPG  types.List `tfsdk:"jpg" json:"jpg,computed"`
	JPG2 types.List `tfsdk:"jpg2" json:"jpg2,computed"`
	PNG  types.List `tfsdk:"png" json:"png,computed"`
	TIF  types.List `tfsdk:"tif" json:"tif,computed"`
	TIFF types.List `tfsdk:"tiff" json:"tiff,computed"`
	WebP types.List `tfsdk:"webp" json:"webp,computed"`
}
