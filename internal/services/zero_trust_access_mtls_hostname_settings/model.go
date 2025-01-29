// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_mtls_hostname_settings

import (
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessMTLSHostnameSettingsResultEnvelope struct {
	Result ZeroTrustAccessMTLSHostnameSettingsModel `json:"result"`
}

type ZeroTrustAccessMTLSHostnameSettingsModel struct {
	AccountID types.String                                         `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID    types.String                                         `tfsdk:"zone_id" path:"zone_id,optional"`
	Settings  *[]*ZeroTrustAccessMTLSHostnameSettingsSettingsModel `tfsdk:"settings" json:"settings,required"`
}

func (m ZeroTrustAccessMTLSHostnameSettingsModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustAccessMTLSHostnameSettingsModel) MarshalJSONForUpdate(state ZeroTrustAccessMTLSHostnameSettingsModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustAccessMTLSHostnameSettingsSettingsModel struct {
	ChinaNetwork                types.Bool   `tfsdk:"china_network" json:"china_network,required"`
	ClientCertificateForwarding types.Bool   `tfsdk:"client_certificate_forwarding" json:"client_certificate_forwarding,required"`
	Hostname                    types.String `tfsdk:"hostname" json:"hostname,required"`
}
