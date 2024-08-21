// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hostname_tls_setting

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/hostnames"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type HostnameTLSSettingResultDataSourceEnvelope struct {
	Result HostnameTLSSettingDataSourceModel `json:"result,computed"`
}

type HostnameTLSSettingDataSourceModel struct {
	SettingID types.String `tfsdk:"setting_id" path:"setting_id"`
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id"`
}

func (m *HostnameTLSSettingDataSourceModel) toReadParams() (params hostnames.SettingTLSGetParams, diags diag.Diagnostics) {
	params = hostnames.SettingTLSGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
