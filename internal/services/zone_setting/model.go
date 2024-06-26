// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_setting

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneSettingResultEnvelope struct {
	Result ZoneSettingModel `json:"result,computed"`
}

type ZoneSettingModel struct {
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id"`
	SettingID types.String `tfsdk:"setting_id" path:"setting_id"`
}
