// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_cache_variants

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneCacheVariantsResultEnvelope struct {
	Result ZoneCacheVariantsModel `json:"result"`
}

type ZoneCacheVariantsModel struct {
	ID         types.String                 `tfsdk:"id" json:"-,computed"`
	ZoneID     types.String                 `tfsdk:"zone_id" path:"zone_id,required"`
	Value      *ZoneCacheVariantsValueModel `tfsdk:"value" json:"value,required"`
	ModifiedOn timetypes.RFC3339            `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
}

type ZoneCacheVariantsValueModel struct {
	AVIF *[]types.String `tfsdk:"avif" json:"avif,optional"`
	BMP  *[]types.String `tfsdk:"bmp" json:"bmp,optional"`
	GIF  *[]types.String `tfsdk:"gif" json:"gif,optional"`
	JP2  *[]types.String `tfsdk:"jp2" json:"jp2,optional"`
	JPEG *[]types.String `tfsdk:"jpeg" json:"jpeg,optional"`
	JPG  *[]types.String `tfsdk:"jpg" json:"jpg,optional"`
	JPG2 *[]types.String `tfsdk:"jpg2" json:"jpg2,optional"`
	PNG  *[]types.String `tfsdk:"png" json:"png,optional"`
	TIF  *[]types.String `tfsdk:"tif" json:"tif,optional"`
	TIFF *[]types.String `tfsdk:"tiff" json:"tiff,optional"`
	WebP *[]types.String `tfsdk:"webp" json:"webp,optional"`
}
