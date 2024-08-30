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
	ZoneID     types.String                 `tfsdk:"zone_id" path:"zone_id"`
	Value      *ZoneCacheVariantsValueModel `tfsdk:"value" json:"value"`
	ModifiedOn timetypes.RFC3339            `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
}

type ZoneCacheVariantsValueModel struct {
	AVIF types.List `tfsdk:"avif" json:"avif,computed_optional"`
	BMP  types.List `tfsdk:"bmp" json:"bmp,computed_optional"`
	GIF  types.List `tfsdk:"gif" json:"gif,computed_optional"`
	JP2  types.List `tfsdk:"jp2" json:"jp2,computed_optional"`
	JPEG types.List `tfsdk:"jpeg" json:"jpeg,computed_optional"`
	JPG  types.List `tfsdk:"jpg" json:"jpg,computed_optional"`
	JPG2 types.List `tfsdk:"jpg2" json:"jpg2,computed_optional"`
	PNG  types.List `tfsdk:"png" json:"png,computed_optional"`
	TIF  types.List `tfsdk:"tif" json:"tif,computed_optional"`
	TIFF types.List `tfsdk:"tiff" json:"tiff,computed_optional"`
	WebP types.List `tfsdk:"webp" json:"webp,computed_optional"`
}
