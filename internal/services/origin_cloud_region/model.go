// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_cloud_region

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type OriginCloudRegionResultEnvelope struct {
	Result OriginCloudRegionModel `json:"result"`
}

type OriginCloudRegionModel struct {
	ID         types.String      `tfsdk:"id" json:"-,computed"`
	OriginIP   types.String      `tfsdk:"origin_ip" json:"origin_ip,required"`
	ZoneID     types.String      `tfsdk:"zone_id" path:"zone_id,required"`
	Region     types.String      `tfsdk:"region" json:"region,required"`
	Vendor     types.String      `tfsdk:"vendor" json:"vendor,required"`
	ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
}

func (m OriginCloudRegionModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m OriginCloudRegionModel) MarshalJSONForUpdate(state OriginCloudRegionModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
