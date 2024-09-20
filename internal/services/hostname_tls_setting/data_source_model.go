// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hostname_tls_setting

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/hostnames"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type HostnameTLSSettingResultDataSourceEnvelope struct {
	Result HostnameTLSSettingDataSourceModel `json:"result,computed"`
}

type HostnameTLSSettingDataSourceModel struct {
	SettingID types.String `tfsdk:"setting_id" path:"setting_id,required"`
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id,required"`
}

func (m *HostnameTLSSettingDataSourceModel) toReadParams(_ context.Context) (params hostnames.SettingTLSGetParams, diags diag.Diagnostics) {
	params = hostnames.SettingTLSGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
