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
	ModifiedOn timetypes.RFC3339                      `tfsdk:"modified_on" json:"modified_on"`
	Value      *ZoneCacheVariantsValueDataSourceModel `tfsdk:"value" json:"value"`
}

func (m *ZoneCacheVariantsDataSourceModel) toReadParams() (params cache.VariantGetParams, diags diag.Diagnostics) {
	params = cache.VariantGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type ZoneCacheVariantsValueDataSourceModel struct {
	AVIF *[]types.String `tfsdk:"avif" json:"avif"`
	BMP  *[]types.String `tfsdk:"bmp" json:"bmp"`
	GIF  *[]types.String `tfsdk:"gif" json:"gif"`
	JP2  *[]types.String `tfsdk:"jp2" json:"jp2"`
	JPEG *[]types.String `tfsdk:"jpeg" json:"jpeg"`
	JPG  *[]types.String `tfsdk:"jpg" json:"jpg"`
	JPG2 *[]types.String `tfsdk:"jpg2" json:"jpg2"`
	PNG  *[]types.String `tfsdk:"png" json:"png"`
	TIF  *[]types.String `tfsdk:"tif" json:"tif"`
	TIFF *[]types.String `tfsdk:"tiff" json:"tiff"`
	WebP *[]types.String `tfsdk:"webp" json:"webp"`
}
