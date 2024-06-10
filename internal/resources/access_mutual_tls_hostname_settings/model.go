// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_mutual_tls_hostname_settings

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessMutualTLSHostnameSettingsResultEnvelope struct {
	Result AccessMutualTLSHostnameSettingsModel `json:"result,computed"`
}

type AccessMutualTLSHostnameSettingsModel struct {
	AccountID types.String                                     `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String                                     `tfsdk:"zone_id" path:"zone_id"`
	Settings  *[]*AccessMutualTLSHostnameSettingsSettingsModel `tfsdk:"settings" json:"settings"`
}

type AccessMutualTLSHostnameSettingsSettingsModel struct {
	ChinaNetwork                types.Bool   `tfsdk:"china_network" json:"china_network"`
	ClientCertificateForwarding types.Bool   `tfsdk:"client_certificate_forwarding" json:"client_certificate_forwarding"`
	Hostname                    types.String `tfsdk:"hostname" json:"hostname"`
}
