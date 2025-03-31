// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_cache_reserve

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneCacheReserveResultEnvelope struct {
	Result ZoneCacheReserveModel `json:"result"`
}

type ZoneCacheReserveModel struct {
	ID            types.String         `tfsdk:"id" json:"-,computed"`
	ZoneID        types.String         `tfsdk:"zone_id" path:"zone_id,required"`
	Value         types.String         `tfsdk:"value" json:"value,computed_optional"`
	Editable      types.Bool           `tfsdk:"editable" json:"editable,computed"`
	ModifiedOn    timetypes.RFC3339    `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	ZoneSettingID jsontypes.Normalized `tfsdk:"zone_setting_id" json:"id,computed"`
}

func (m ZoneCacheReserveModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZoneCacheReserveModel) MarshalJSONForUpdate(state ZoneCacheReserveModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
