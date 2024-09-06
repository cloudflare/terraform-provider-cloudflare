// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_cache_variants

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
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
	AVIF customfield.List[types.String] `tfsdk:"avif" json:"avif,computed_optional"`
	BMP  customfield.List[types.String] `tfsdk:"bmp" json:"bmp,computed_optional"`
	GIF  customfield.List[types.String] `tfsdk:"gif" json:"gif,computed_optional"`
	JP2  customfield.List[types.String] `tfsdk:"jp2" json:"jp2,computed_optional"`
	JPEG customfield.List[types.String] `tfsdk:"jpeg" json:"jpeg,computed_optional"`
	JPG  customfield.List[types.String] `tfsdk:"jpg" json:"jpg,computed_optional"`
	JPG2 customfield.List[types.String] `tfsdk:"jpg2" json:"jpg2,computed_optional"`
	PNG  customfield.List[types.String] `tfsdk:"png" json:"png,computed_optional"`
	TIF  customfield.List[types.String] `tfsdk:"tif" json:"tif,computed_optional"`
	TIFF customfield.List[types.String] `tfsdk:"tiff" json:"tiff,computed_optional"`
	WebP customfield.List[types.String] `tfsdk:"webp" json:"webp,computed_optional"`
}
