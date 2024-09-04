// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_cache_reserve

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneCacheReserveResultEnvelope struct {
	Result ZoneCacheReserveModel `json:"result"`
}

type ZoneCacheReserveModel struct {
	ID            types.String      `tfsdk:"id" json:"-,computed"`
	ZoneID        types.String      `tfsdk:"zone_id" path:"zone_id,required"`
	Value         types.String      `tfsdk:"value" json:"value,computed_optional"`
	ModifiedOn    timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	ZoneSettingID types.String      `tfsdk:"zone_setting_id" json:"id,computed"`
}
