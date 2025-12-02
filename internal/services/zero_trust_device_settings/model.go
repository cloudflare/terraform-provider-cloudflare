// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_settings

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDeviceSettingsResultEnvelope struct {
	Result ZeroTrustDeviceSettingsModel `json:"result"`
}

type ZeroTrustDeviceSettingsModel struct {
	AccountID                          types.String  `tfsdk:"account_id" path:"account_id,required"`
	DisableForTime                     types.Float64 `tfsdk:"disable_for_time" json:"disable_for_time,optional"`
	ExternalEmergencySignalEnabled     types.Bool    `tfsdk:"external_emergency_signal_enabled" json:"external_emergency_signal_enabled,optional"`
	ExternalEmergencySignalFingerprint types.String  `tfsdk:"external_emergency_signal_fingerprint" json:"external_emergency_signal_fingerprint,optional"`
	ExternalEmergencySignalInterval    types.String  `tfsdk:"external_emergency_signal_interval" json:"external_emergency_signal_interval,optional"`
	ExternalEmergencySignalURL         types.String  `tfsdk:"external_emergency_signal_url" json:"external_emergency_signal_url,optional"`
	GatewayProxyEnabled                types.Bool    `tfsdk:"gateway_proxy_enabled" json:"gateway_proxy_enabled,optional"`
	GatewayUdpProxyEnabled             types.Bool    `tfsdk:"gateway_udp_proxy_enabled" json:"gateway_udp_proxy_enabled,optional"`
	RootCertificateInstallationEnabled types.Bool    `tfsdk:"root_certificate_installation_enabled" json:"root_certificate_installation_enabled,optional"`
	UseZtVirtualIP                     types.Bool    `tfsdk:"use_zt_virtual_ip" json:"use_zt_virtual_ip,optional"`
}

func (m ZeroTrustDeviceSettingsModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustDeviceSettingsModel) MarshalJSONForUpdate(state ZeroTrustDeviceSettingsModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
