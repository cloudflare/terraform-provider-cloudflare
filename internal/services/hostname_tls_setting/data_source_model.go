// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hostname_tls_setting

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type HostnameTLSSettingResultDataSourceEnvelope struct {
	Result HostnameTLSSettingDataSourceModel `json:"result,computed"`
}

type HostnameTLSSettingDataSourceModel struct {
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id"`
	SettingID types.String `tfsdk:"setting_id" path:"setting_id"`
}
