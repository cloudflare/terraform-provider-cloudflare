// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package teams_account

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TeamsAccountResultDataSourceEnvelope struct {
	Result TeamsAccountDataSourceModel `json:"result,computed"`
}

type TeamsAccountDataSourceModel struct {
	AccountID types.String                         `tfsdk:"account_id" path:"account_id"`
	CreatedAt timetypes.RFC3339                    `tfsdk:"created_at" json:"created_at"`
	UpdatedAt timetypes.RFC3339                    `tfsdk:"updated_at" json:"updated_at"`
	Settings  *TeamsAccountSettingsDataSourceModel `tfsdk:"settings" json:"settings"`
}

type TeamsAccountSettingsDataSourceModel struct {
	ActivityLog           *TeamsAccountSettingsActivityLogDataSourceModel           `tfsdk:"activity_log" json:"activity_log"`
	Antivirus             *TeamsAccountSettingsAntivirusDataSourceModel             `tfsdk:"antivirus" json:"antivirus"`
	BlockPage             *TeamsAccountSettingsBlockPageDataSourceModel             `tfsdk:"block_page" json:"block_page"`
	BodyScanning          *TeamsAccountSettingsBodyScanningDataSourceModel          `tfsdk:"body_scanning" json:"body_scanning"`
	BrowserIsolation      *TeamsAccountSettingsBrowserIsolationDataSourceModel      `tfsdk:"browser_isolation" json:"browser_isolation"`
	Certificate           *TeamsAccountSettingsCertificateDataSourceModel           `tfsdk:"certificate" json:"certificate"`
	CustomCertificate     *TeamsAccountSettingsCustomCertificateDataSourceModel     `tfsdk:"custom_certificate" json:"custom_certificate"`
	ExtendedEmailMatching *TeamsAccountSettingsExtendedEmailMatchingDataSourceModel `tfsdk:"extended_email_matching" json:"extended_email_matching"`
	Fips                  *TeamsAccountSettingsFipsDataSourceModel                  `tfsdk:"fips" json:"fips"`
	ProtocolDetection     *TeamsAccountSettingsProtocolDetectionDataSourceModel     `tfsdk:"protocol_detection" json:"protocol_detection"`
	TLSDecrypt            *TeamsAccountSettingsTLSDecryptDataSourceModel            `tfsdk:"tls_decrypt" json:"tls_decrypt"`
}

type TeamsAccountSettingsActivityLogDataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled"`
}

type TeamsAccountSettingsAntivirusDataSourceModel struct {
	EnabledDownloadPhase types.Bool                                                        `tfsdk:"enabled_download_phase" json:"enabled_download_phase"`
	EnabledUploadPhase   types.Bool                                                        `tfsdk:"enabled_upload_phase" json:"enabled_upload_phase"`
	FailClosed           types.Bool                                                        `tfsdk:"fail_closed" json:"fail_closed"`
	NotificationSettings *TeamsAccountSettingsAntivirusNotificationSettingsDataSourceModel `tfsdk:"notification_settings" json:"notification_settings"`
}

type TeamsAccountSettingsAntivirusNotificationSettingsDataSourceModel struct {
	Enabled    types.Bool   `tfsdk:"enabled" json:"enabled"`
	Msg        types.String `tfsdk:"msg" json:"msg"`
	SupportURL types.String `tfsdk:"support_url" json:"support_url"`
}

type TeamsAccountSettingsBlockPageDataSourceModel struct {
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

type TeamsAccountSettingsBodyScanningDataSourceModel struct {
	InspectionMode types.String `tfsdk:"inspection_mode" json:"inspection_mode"`
}

type TeamsAccountSettingsBrowserIsolationDataSourceModel struct {
	NonIdentityEnabled         types.Bool `tfsdk:"non_identity_enabled" json:"non_identity_enabled"`
	URLBrowserIsolationEnabled types.Bool `tfsdk:"url_browser_isolation_enabled" json:"url_browser_isolation_enabled"`
}

type TeamsAccountSettingsCertificateDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type TeamsAccountSettingsCustomCertificateDataSourceModel struct {
	Enabled       types.Bool        `tfsdk:"enabled" json:"enabled,computed"`
	ID            types.String      `tfsdk:"id" json:"id"`
	BindingStatus types.String      `tfsdk:"binding_status" json:"binding_status,computed"`
	UpdatedAt     timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed"`
}

type TeamsAccountSettingsExtendedEmailMatchingDataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled"`
}

type TeamsAccountSettingsFipsDataSourceModel struct {
	TLS types.Bool `tfsdk:"tls" json:"tls"`
}

type TeamsAccountSettingsProtocolDetectionDataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled"`
}

type TeamsAccountSettingsTLSDecryptDataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled"`
}
