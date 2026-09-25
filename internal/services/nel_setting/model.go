// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nel_setting

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NELSettingResultEnvelope struct {
	Result NELSettingModel `json:"result"`
}

type NELSettingModel struct {
	ID         types.String          `tfsdk:"id" json:"id,computed"`
	ZoneID     types.String          `tfsdk:"zone_id" path:"zone_id,required"`
	Value      *NELSettingValueModel `tfsdk:"value" json:"value,required"`
	Editable   types.Bool            `tfsdk:"editable" json:"editable,computed"`
	ModifiedOn timetypes.RFC3339     `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
}

func (m NELSettingModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m NELSettingModel) MarshalJSONForUpdate(state NELSettingModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type NELSettingValueModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,required"`
}
