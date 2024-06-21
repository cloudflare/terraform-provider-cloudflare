// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_cache_variants

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneCacheVariantsResultEnvelope struct {
	Result ZoneCacheVariantsModel `json:"result,computed"`
}

type ZoneCacheVariantsModel struct {
	ID     types.String                 `tfsdk:"id" json:"id"`
	ZoneID types.String                 `tfsdk:"zone_id" path:"zone_id"`
	Value  *ZoneCacheVariantsValueModel `tfsdk:"value" json:"value"`
}

type ZoneCacheVariantsValueModel struct {
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
