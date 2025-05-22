// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_settings

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDeviceSettingsResultDataSourceEnvelope struct {
	Result ZeroTrustDeviceSettingsDataSourceModel `json:"result,computed"`
}

type ZeroTrustDeviceSettingsDataSourceModel struct {
	AccountID                          types.String  `tfsdk:"account_id" path:"account_id,required"`
	DisableForTime                     types.Float64 `tfsdk:"disable_for_time" json:"disable_for_time,computed"`
	GatewayProxyEnabled                types.Bool    `tfsdk:"gateway_proxy_enabled" json:"gateway_proxy_enabled,computed"`
	GatewayUdpProxyEnabled             types.Bool    `tfsdk:"gateway_udp_proxy_enabled" json:"gateway_udp_proxy_enabled,computed"`
	RootCertificateInstallationEnabled types.Bool    `tfsdk:"root_certificate_installation_enabled" json:"root_certificate_installation_enabled,computed"`
	UseZtVirtualIP                     types.Bool    `tfsdk:"use_zt_virtual_ip" json:"use_zt_virtual_ip,computed"`
}

func (m *ZeroTrustDeviceSettingsDataSourceModel) toReadParams(_ context.Context) (params zero_trust.DeviceSettingGetParams, diags diag.Diagnostics) {
	params = zero_trust.DeviceSettingGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
