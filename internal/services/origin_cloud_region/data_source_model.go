// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_cloud_region

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/cache"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type OriginCloudRegionResultDataSourceEnvelope struct {
	Result OriginCloudRegionDataSourceModel `json:"result,computed"`
}

type OriginCloudRegionDataSourceModel struct {
	ID         types.String      `tfsdk:"id" path:"origin_ip,computed"`
	OriginIP   types.String      `tfsdk:"origin_ip" path:"origin_ip,required"`
	ZoneID     types.String      `tfsdk:"zone_id" path:"zone_id,required"`
	ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Region     types.String      `tfsdk:"region" json:"region,computed"`
	Vendor     types.String      `tfsdk:"vendor" json:"vendor,computed"`
}

func (m *OriginCloudRegionDataSourceModel) toReadParams(_ context.Context) (params cache.OriginCloudRegionGetParams, diags diag.Diagnostics) {
	params = cache.OriginCloudRegionGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
