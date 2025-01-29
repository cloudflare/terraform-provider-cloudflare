// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package regional_hostname

import (
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RegionalHostnameResultEnvelope struct {
	Result RegionalHostnameModel `json:"result"`
}

type RegionalHostnameModel struct {
	ID        types.String      `tfsdk:"id" json:"-,computed"`
	Hostname  types.String      `tfsdk:"hostname" json:"hostname,required"`
	ZoneID    types.String      `tfsdk:"zone_id" path:"zone_id,required"`
	RegionKey types.String      `tfsdk:"region_key" json:"region_key,required"`
	CreatedOn timetypes.RFC3339 `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
}

func (m RegionalHostnameModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m RegionalHostnameModel) MarshalJSONForUpdate(state RegionalHostnameModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
