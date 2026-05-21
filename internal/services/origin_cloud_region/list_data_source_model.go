// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_cloud_region

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/cache"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type OriginCloudRegionsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[OriginCloudRegionsResultDataSourceModel] `json:"result,computed"`
}

type OriginCloudRegionsDataSourceModel struct {
	ZoneID   types.String                                                          `tfsdk:"zone_id" path:"zone_id,required"`
	MaxItems types.Int64                                                           `tfsdk:"max_items"`
	Result   customfield.NestedObjectList[OriginCloudRegionsResultDataSourceModel] `tfsdk:"result"`
}

func (m *OriginCloudRegionsDataSourceModel) toListParams(_ context.Context) (params cache.OriginCloudRegionListParams, diags diag.Diagnostics) {
	params = cache.OriginCloudRegionListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type OriginCloudRegionsResultDataSourceModel struct {
	ID         types.String      `tfsdk:"id" json:"origin_ip,computed"`
	OriginIP   types.String      `tfsdk:"origin_ip" json:"origin_ip,computed"`
	Region     types.String      `tfsdk:"region" json:"region,computed"`
	Vendor     types.String      `tfsdk:"vendor" json:"vendor,computed"`
	ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
}
