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
	AVIF *[]types.String `tfsdk:"avif" json:"avif,computed_optional"`
	BMP  *[]types.String `tfsdk:"bmp" json:"bmp,computed_optional"`
	GIF  *[]types.String `tfsdk:"gif" json:"gif,computed_optional"`
	JP2  *[]types.String `tfsdk:"jp2" json:"jp2,computed_optional"`
	JPEG *[]types.String `tfsdk:"jpeg" json:"jpeg,computed_optional"`
	JPG  *[]types.String `tfsdk:"jpg" json:"jpg,computed_optional"`
	JPG2 *[]types.String `tfsdk:"jpg2" json:"jpg2,computed_optional"`
	PNG  *[]types.String `tfsdk:"png" json:"png,computed_optional"`
	TIF  *[]types.String `tfsdk:"tif" json:"tif,computed_optional"`
	TIFF *[]types.String `tfsdk:"tiff" json:"tiff,computed_optional"`
	WebP *[]types.String `tfsdk:"webp" json:"webp,computed_optional"`
}
