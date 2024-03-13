package access_mutual_tls_hostname_settings

import "github.com/hashicorp/terraform-plugin-framework/types"

type Settings struct {
	Hostname                    types.String `tfsdk:"hostname"`
	ChinaNetwork                types.Bool   `tfsdk:"china_network"`
	ClientCertificateForwarding types.Bool   `tfsdk:"client_certificate_forwarding"`
}

type AccessMutualTLSHostnameSettingsModel struct {
	AccountID types.String `tfsdk:"account_id"`
	ZoneID    types.String `tfsdk:"zone_id"`
	Settings  []Settings   `tfsdk:"settings"`
}
