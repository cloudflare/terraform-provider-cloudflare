// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package teams_account

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TeamsAccountResultEnvelope struct {
	Result TeamsAccountModel `json:"result,computed"`
}

type TeamsAccountModel struct {
	AccountID types.String               `tfsdk:"account_id" path:"account_id"`
	Settings  *TeamsAccountSettingsModel `tfsdk:"settings" json:"settings"`
}

type TeamsAccountSettingsModel struct {
	ActivityLog           *TeamsAccountSettingsActivityLogModel           `tfsdk:"activity_log" json:"activity_log"`
	Antivirus             *TeamsAccountSettingsAntivirusModel             `tfsdk:"antivirus" json:"antivirus"`
	BlockPage             *TeamsAccountSettingsBlockPageModel             `tfsdk:"block_page" json:"block_page"`
	BodyScanning          *TeamsAccountSettingsBodyScanningModel          `tfsdk:"body_scanning" json:"body_scanning"`
	BrowserIsolation      *TeamsAccountSettingsBrowserIsolationModel      `tfsdk:"browser_isolation" json:"browser_isolation"`
	CustomCertificate     *TeamsAccountSettingsCustomCertificateModel     `tfsdk:"custom_certificate" json:"custom_certificate"`
	ExtendedEmailMatching *TeamsAccountSettingsExtendedEmailMatchingModel `tfsdk:"extended_email_matching" json:"extended_email_matching"`
	Fips                  *TeamsAccountSettingsFipsModel                  `tfsdk:"fips" json:"fips"`
	ProtocolDetection     *TeamsAccountSettingsProtocolDetectionModel     `tfsdk:"protocol_detection" json:"protocol_detection"`
	TLSDecrypt            *TeamsAccountSettingsTLSDecryptModel            `tfsdk:"tls_decrypt" json:"tls_decrypt"`
}

type TeamsAccountSettingsActivityLogModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled"`
}

type TeamsAccountSettingsAntivirusModel struct {
	EnabledDownloadPhase types.Bool                                              `tfsdk:"enabled_download_phase" json:"enabled_download_phase"`
	EnabledUploadPhase   types.Bool                                              `tfsdk:"enabled_upload_phase" json:"enabled_upload_phase"`
	FailClosed           types.Bool                                              `tfsdk:"fail_closed" json:"fail_closed"`
	NotificationSettings *TeamsAccountSettingsAntivirusNotificationSettingsModel `tfsdk:"notification_settings" json:"notification_settings"`
}

type TeamsAccountSettingsAntivirusNotificationSettingsModel struct {
	Enabled    types.Bool   `tfsdk:"enabled" json:"enabled"`
	Msg        types.String `tfsdk:"msg" json:"msg"`
	SupportURL types.String `tfsdk:"support_url" json:"support_url"`
}

type TeamsAccountSettingsBlockPageModel struct {
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

type TeamsAccountSettingsBodyScanningModel struct {
	InspectionMode types.String `tfsdk:"inspection_mode" json:"inspection_mode"`
}

type TeamsAccountSettingsBrowserIsolationModel struct {
	NonIdentityEnabled         types.Bool `tfsdk:"non_identity_enabled" json:"non_identity_enabled"`
	URLBrowserIsolationEnabled types.Bool `tfsdk:"url_browser_isolation_enabled" json:"url_browser_isolation_enabled"`
}

type TeamsAccountSettingsCustomCertificateModel struct {
	Enabled       types.Bool   `tfsdk:"enabled" json:"enabled"`
	ID            types.String `tfsdk:"id" json:"id"`
	BindingStatus types.String `tfsdk:"binding_status" json:"binding_status,computed"`
	UpdatedAt     types.String `tfsdk:"updated_at" json:"updated_at,computed"`
}

type TeamsAccountSettingsExtendedEmailMatchingModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled"`
}

type TeamsAccountSettingsFipsModel struct {
	TLS types.Bool `tfsdk:"tls" json:"tls"`
}

type TeamsAccountSettingsProtocolDetectionModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled"`
}

type TeamsAccountSettingsTLSDecryptModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled"`
}
