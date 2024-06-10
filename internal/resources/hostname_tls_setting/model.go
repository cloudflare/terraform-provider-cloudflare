// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hostname_tls_setting

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type HostnameTLSSettingResultEnvelope struct {
	Result HostnameTLSSettingModel `json:"result,computed"`
}

type HostnameTLSSettingModel struct {
	ZoneID    types.String  `tfsdk:"zone_id" path:"zone_id"`
	SettingID types.String  `tfsdk:"setting_id" path:"setting_id"`
	Hostname  types.String  `tfsdk:"hostname" path:"hostname"`
	Value     types.Float64 `tfsdk:"value" json:"value"`
}
