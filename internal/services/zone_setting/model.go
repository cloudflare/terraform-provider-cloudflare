// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_setting

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneSettingResultEnvelope struct {
	Result ZoneSettingModel `json:"result"`
}

type ZoneSettingModel struct {
	ID            types.String                       `tfsdk:"id" json:"-,computed"`
	SettingID     types.String                       `tfsdk:"setting_id" path:"setting_id,required"`
	ZoneID        types.String                       `tfsdk:"zone_id" path:"zone_id,required"`
	Value         customfield.NormalizedDynamicValue `tfsdk:"value" json:"value,required"`
	Enabled       types.Bool                         `tfsdk:"enabled" json:"enabled,computed_optional"`
	Editable      types.Bool                         `tfsdk:"editable" json:"editable,computed"`
	ModifiedOn    timetypes.RFC3339                  `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	TimeRemaining types.Float64                      `tfsdk:"time_remaining" json:"time_remaining,computed"`
}

func (m ZoneSettingModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

// MarshalJSONForUpdate sends the full planned value rather than a JSON merge
// patch. The zone settings edit endpoint replaces `value` wholesale, so a patch
// containing only the changed keys of an object-valued setting (e.g. `aegis`)
// drops the unchanged ones. For aegis that meant `{"value":{"pool_id":"..."}}`
// was sent without `enabled: false`, and the setting was enabled.
func (m ZoneSettingModel) MarshalJSONForUpdate(state ZoneSettingModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
