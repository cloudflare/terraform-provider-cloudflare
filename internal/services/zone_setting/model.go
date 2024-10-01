// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_setting

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneSettingResultEnvelope struct {
	Result ZoneSettingModel `json:"result"`
}

type ZoneSettingModel struct {
	SettingID types.String  `tfsdk:"setting_id" path:"setting_id,required"`
	ZoneID    types.String  `tfsdk:"zone_id" path:"zone_id,required"`
	ID        types.String  `tfsdk:"id" json:"id,optional"`
	Value     types.Dynamic `tfsdk:"value" json:"value,required"`
}
