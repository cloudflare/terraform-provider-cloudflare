// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_cache_variants

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/cache"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneCacheVariantsResultDataSourceEnvelope struct {
	Result ZoneCacheVariantsDataSourceModel `json:"result,computed"`
}

type ZoneCacheVariantsDataSourceModel struct {
	ZoneID     types.String                           `tfsdk:"zone_id" path:"zone_id,required"`
	ID         types.String                           `tfsdk:"id" json:"id,optional"`
	ModifiedOn timetypes.RFC3339                      `tfsdk:"modified_on" json:"modified_on,optional" format:"date-time"`
	Value      *ZoneCacheVariantsValueDataSourceModel `tfsdk:"value" json:"value,optional"`
}

func (m *ZoneCacheVariantsDataSourceModel) toReadParams(_ context.Context) (params cache.VariantGetParams, diags diag.Diagnostics) {
	params = cache.VariantGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type ZoneCacheVariantsValueDataSourceModel struct {
	AVIF customfield.List[types.String] `tfsdk:"avif" json:"avif,computed"`
	BMP  customfield.List[types.String] `tfsdk:"bmp" json:"bmp,computed"`
	GIF  customfield.List[types.String] `tfsdk:"gif" json:"gif,computed"`
	JP2  customfield.List[types.String] `tfsdk:"jp2" json:"jp2,computed"`
	JPEG customfield.List[types.String] `tfsdk:"jpeg" json:"jpeg,computed"`
	JPG  customfield.List[types.String] `tfsdk:"jpg" json:"jpg,computed"`
	JPG2 customfield.List[types.String] `tfsdk:"jpg2" json:"jpg2,computed"`
	PNG  customfield.List[types.String] `tfsdk:"png" json:"png,computed"`
	TIF  customfield.List[types.String] `tfsdk:"tif" json:"tif,computed"`
	TIFF customfield.List[types.String] `tfsdk:"tiff" json:"tiff,computed"`
	WebP customfield.List[types.String] `tfsdk:"webp" json:"webp,computed"`
}
