// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_settings

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustGatewaySettingsResultEnvelope struct {
	Result ZeroTrustGatewaySettingsModel `json:"result"`
}

type ZeroTrustGatewaySettingsModel struct {
	ID        types.String                           `tfsdk:"id" json:"-,computed"`
	AccountID types.String                           `tfsdk:"account_id" path:"account_id"`
	Settings  *ZeroTrustGatewaySettingsSettingsModel `tfsdk:"settings" json:"settings"`
	CreatedAt timetypes.RFC3339                      `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UpdatedAt timetypes.RFC3339                      `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

type ZeroTrustGatewaySettingsSettingsModel struct {
	ActivityLog           *ZeroTrustGatewaySettingsSettingsActivityLogModel           `tfsdk:"activity_log" json:"activity_log"`
	Antivirus             *ZeroTrustGatewaySettingsSettingsAntivirusModel             `tfsdk:"antivirus" json:"antivirus"`
	BlockPage             *ZeroTrustGatewaySettingsSettingsBlockPageModel             `tfsdk:"block_page" json:"block_page"`
	BodyScanning          *ZeroTrustGatewaySettingsSettingsBodyScanningModel          `tfsdk:"body_scanning" json:"body_scanning"`
	BrowserIsolation      *ZeroTrustGatewaySettingsSettingsBrowserIsolationModel      `tfsdk:"browser_isolation" json:"browser_isolation"`
	Certificate           *ZeroTrustGatewaySettingsSettingsCertificateModel           `tfsdk:"certificate" json:"certificate"`
	CustomCertificate     *ZeroTrustGatewaySettingsSettingsCustomCertificateModel     `tfsdk:"custom_certificate" json:"custom_certificate"`
	ExtendedEmailMatching *ZeroTrustGatewaySettingsSettingsExtendedEmailMatchingModel `tfsdk:"extended_email_matching" json:"extended_email_matching"`
	Fips                  *ZeroTrustGatewaySettingsSettingsFipsModel                  `tfsdk:"fips" json:"fips"`
	ProtocolDetection     *ZeroTrustGatewaySettingsSettingsProtocolDetectionModel     `tfsdk:"protocol_detection" json:"protocol_detection"`
	TLSDecrypt            *ZeroTrustGatewaySettingsSettingsTLSDecryptModel            `tfsdk:"tls_decrypt" json:"tls_decrypt"`
}

type ZeroTrustGatewaySettingsSettingsActivityLogModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled"`
}

type ZeroTrustGatewaySettingsSettingsAntivirusModel struct {
	EnabledDownloadPhase types.Bool                                                          `tfsdk:"enabled_download_phase" json:"enabled_download_phase"`
	EnabledUploadPhase   types.Bool                                                          `tfsdk:"enabled_upload_phase" json:"enabled_upload_phase"`
	FailClosed           types.Bool                                                          `tfsdk:"fail_closed" json:"fail_closed"`
	NotificationSettings *ZeroTrustGatewaySettingsSettingsAntivirusNotificationSettingsModel `tfsdk:"notification_settings" json:"notification_settings"`
}

type ZeroTrustGatewaySettingsSettingsAntivirusNotificationSettingsModel struct {
	Enabled    types.Bool   `tfsdk:"enabled" json:"enabled"`
	Msg        types.String `tfsdk:"msg" json:"msg"`
	SupportURL types.String `tfsdk:"support_url" json:"support_url"`
}

type ZeroTrustGatewaySettingsSettingsBlockPageModel struct {
	BackgroundColor types.String `tfsdk:"background_color" json:"background_color"`
	Enabled         types.Bool   `tfsdk:"enabled" json:"enabled"`
	FooterText      types.String `tfsdk:"footer_text" json:"footer_text"`
	HeaderText      types.String `tfsdk:"header_text" json:"header_text"`
	LogoPath        types.String `tfsdk:"logo_path" json:"logo_path"`
	MailtoAddress   types.String `tfsdk:"mailto_address" json:"mailto_address"`
	MailtoSubject   types.String `tfsdk:"mailto_subject" json:"mailto_subject"`
	Name            types.String `tfsdk:"name" json:"name"`
	SuppressFooter  types.Bool   `tfsdk:"suppress_footer" json:"suppress_footer"`
}

type ZeroTrustGatewaySettingsSettingsBodyScanningModel struct {
	InspectionMode types.String `tfsdk:"inspection_mode" json:"inspection_mode"`
}

type ZeroTrustGatewaySettingsSettingsBrowserIsolationModel struct {
	NonIdentityEnabled         types.Bool `tfsdk:"non_identity_enabled" json:"non_identity_enabled"`
	URLBrowserIsolationEnabled types.Bool `tfsdk:"url_browser_isolation_enabled" json:"url_browser_isolation_enabled"`
}

type ZeroTrustGatewaySettingsSettingsCertificateModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type ZeroTrustGatewaySettingsSettingsCustomCertificateModel struct {
	Enabled       types.Bool        `tfsdk:"enabled" json:"enabled"`
	ID            types.String      `tfsdk:"id" json:"id"`
	BindingStatus types.String      `tfsdk:"binding_status" json:"binding_status,computed"`
	UpdatedAt     timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

type ZeroTrustGatewaySettingsSettingsExtendedEmailMatchingModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled"`
}

type ZeroTrustGatewaySettingsSettingsFipsModel struct {
	TLS types.Bool `tfsdk:"tls" json:"tls"`
}

type ZeroTrustGatewaySettingsSettingsProtocolDetectionModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled"`
}

type ZeroTrustGatewaySettingsSettingsTLSDecryptModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled"`
}
