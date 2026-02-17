package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy v4 SDKv2 Provider)
// ============================================================================

// SourceMTLSHostnameSettingsModel represents the v4 cloudflare_zero_trust_access_mtls_hostname_settings state.
// Schema version: 0 (SDKv2 default)
type SourceMTLSHostnameSettingsModel struct {
	AccountID                   types.String          `tfsdk:"account_id"`
	ZoneID                      types.String          `tfsdk:"zone_id"`
	Settings                    []SourceSettingsModel `tfsdk:"settings"` // v4 block stored as list
	ChinaNetwork                types.Bool            `tfsdk:"china_network"`
	ClientCertificateForwarding types.Bool            `tfsdk:"client_certificate_forwarding"`
	Hostname                    types.String          `tfsdk:"hostname"`
}

// SourceSettingsModel represents each settings block in v4.
// china_network and client_certificate_forwarding were Optional in v4 (may be null).
type SourceSettingsModel struct {
	ChinaNetwork                types.Bool   `tfsdk:"china_network"`
	ClientCertificateForwarding types.Bool   `tfsdk:"client_certificate_forwarding"`
	Hostname                    types.String `tfsdk:"hostname"`
}

// ============================================================================
// Target Models (Current v5 Provider)
// ============================================================================

// TargetMTLSHostnameSettingsModel represents the v5 cloudflare_zero_trust_access_mtls_hostname_settings state.
// Schema version: 500
type TargetMTLSHostnameSettingsModel struct {
	AccountID                   types.String           `tfsdk:"account_id"`
	ZoneID                      types.String           `tfsdk:"zone_id"`
	Settings                    *[]*TargetSettingsModel `tfsdk:"settings"`
	ChinaNetwork                types.Bool             `tfsdk:"china_network"`
	ClientCertificateForwarding types.Bool             `tfsdk:"client_certificate_forwarding"`
	Hostname                    types.String           `tfsdk:"hostname"`
}

// TargetSettingsModel represents each settings entry in v5.
// china_network and client_certificate_forwarding are Required in v5.
type TargetSettingsModel struct {
	ChinaNetwork                types.Bool   `tfsdk:"china_network"`
	ClientCertificateForwarding types.Bool   `tfsdk:"client_certificate_forwarding"`
	Hostname                    types.String `tfsdk:"hostname"`
}
