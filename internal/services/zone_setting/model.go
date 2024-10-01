// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_setting

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneSettingResultEnvelope struct {
	Result ZoneSettingModel `json:"result"`
}

type ZoneSettingModel struct {
	SettingID types.String `tfsdk:"setting_id" path:"setting_id,required"`
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id,required"`
	ID        types.String `tfsdk:"id" json:"id,optional"`
	Enabled   types.Bool   `tfsdk:"enabled" json:"enabled,computed_optional"`
	Value     types.String `tfsdk:"value" json:"value,computed_optional"`
}

func (m ZoneSettingModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZoneSettingModel) MarshalJSONForUpdate(state ZoneSettingModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
