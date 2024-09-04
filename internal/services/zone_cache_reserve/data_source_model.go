// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_cache_reserve

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/cache"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneCacheReserveResultDataSourceEnvelope struct {
	Result ZoneCacheReserveDataSourceModel `json:"result,computed"`
}

type ZoneCacheReserveDataSourceModel struct {
	ZoneID        types.String      `tfsdk:"zone_id" path:"zone_id,required"`
	ModifiedOn    timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,optional" format:"date-time"`
	ZoneSettingID types.String      `tfsdk:"zone_setting_id" json:"id,optional"`
	Value         types.String      `tfsdk:"value" json:"value,computed_optional"`
}

func (m *ZoneCacheReserveDataSourceModel) toReadParams(_ context.Context) (params cache.CacheReserveGetParams, diags diag.Diagnostics) {
	params = cache.CacheReserveGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
