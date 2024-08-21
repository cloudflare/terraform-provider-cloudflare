// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_settings

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustGatewaySettingsResultDataSourceEnvelope struct {
	Result ZeroTrustGatewaySettingsDataSourceModel `json:"result,computed"`
}

type ZeroTrustGatewaySettingsDataSourceModel struct {
	AccountID types.String                                     `tfsdk:"account_id" path:"account_id"`
	CreatedAt timetypes.RFC3339                                `tfsdk:"created_at" json:"created_at"`
	UpdatedAt timetypes.RFC3339                                `tfsdk:"updated_at" json:"updated_at"`
	Settings  *ZeroTrustGatewaySettingsSettingsDataSourceModel `tfsdk:"settings" json:"settings"`
}

func (m *ZeroTrustGatewaySettingsDataSourceModel) toReadParams() (params zero_trust.GatewayConfigurationGetParams, diags diag.Diagnostics) {
	params = zero_trust.GatewayConfigurationGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustGatewaySettingsSettingsDataSourceModel struct {
	ActivityLog           *ZeroTrustGatewaySettingsSettingsActivityLogDataSourceModel           `tfsdk:"activity_log" json:"activity_log"`
	Antivirus             *ZeroTrustGatewaySettingsSettingsAntivirusDataSourceModel             `tfsdk:"antivirus" json:"antivirus"`
	BlockPage             *ZeroTrustGatewaySettingsSettingsBlockPageDataSourceModel             `tfsdk:"block_page" json:"block_page"`
	BodyScanning          *ZeroTrustGatewaySettingsSettingsBodyScanningDataSourceModel          `tfsdk:"body_scanning" json:"body_scanning"`
	BrowserIsolation      *ZeroTrustGatewaySettingsSettingsBrowserIsolationDataSourceModel      `tfsdk:"browser_isolation" json:"browser_isolation"`
	Certificate           *ZeroTrustGatewaySettingsSettingsCertificateDataSourceModel           `tfsdk:"certificate" json:"certificate"`
	CustomCertificate     *ZeroTrustGatewaySettingsSettingsCustomCertificateDataSourceModel     `tfsdk:"custom_certificate" json:"custom_certificate"`
	ExtendedEmailMatching *ZeroTrustGatewaySettingsSettingsExtendedEmailMatchingDataSourceModel `tfsdk:"extended_email_matching" json:"extended_email_matching"`
	Fips                  *ZeroTrustGatewaySettingsSettingsFipsDataSourceModel                  `tfsdk:"fips" json:"fips"`
	ProtocolDetection     *ZeroTrustGatewaySettingsSettingsProtocolDetectionDataSourceModel     `tfsdk:"protocol_detection" json:"protocol_detection"`
	TLSDecrypt            *ZeroTrustGatewaySettingsSettingsTLSDecryptDataSourceModel            `tfsdk:"tls_decrypt" json:"tls_decrypt"`
}

type ZeroTrustGatewaySettingsSettingsActivityLogDataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled"`
}

type ZeroTrustGatewaySettingsSettingsAntivirusDataSourceModel struct {
	EnabledDownloadPhase types.Bool                                                                    `tfsdk:"enabled_download_phase" json:"enabled_download_phase"`
	EnabledUploadPhase   types.Bool                                                                    `tfsdk:"enabled_upload_phase" json:"enabled_upload_phase"`
	FailClosed           types.Bool                                                                    `tfsdk:"fail_closed" json:"fail_closed"`
	NotificationSettings *ZeroTrustGatewaySettingsSettingsAntivirusNotificationSettingsDataSourceModel `tfsdk:"notification_settings" json:"notification_settings"`
}

type ZeroTrustGatewaySettingsSettingsAntivirusNotificationSettingsDataSourceModel struct {
	Enabled    types.Bool   `tfsdk:"enabled" json:"enabled"`
	Msg        types.String `tfsdk:"msg" json:"msg"`
	SupportURL types.String `tfsdk:"support_url" json:"support_url"`
}

type ZeroTrustGatewaySettingsSettingsBlockPageDataSourceModel struct {
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

type ZeroTrustGatewaySettingsSettingsBodyScanningDataSourceModel struct {
	InspectionMode types.String `tfsdk:"inspection_mode" json:"inspection_mode"`
}

type ZeroTrustGatewaySettingsSettingsBrowserIsolationDataSourceModel struct {
	NonIdentityEnabled         types.Bool `tfsdk:"non_identity_enabled" json:"non_identity_enabled"`
	URLBrowserIsolationEnabled types.Bool `tfsdk:"url_browser_isolation_enabled" json:"url_browser_isolation_enabled"`
}

type ZeroTrustGatewaySettingsSettingsCertificateDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustGatewaySettingsSettingsCustomCertificateDataSourceModel struct {
	Enabled       types.Bool        `tfsdk:"enabled" json:"enabled,computed"`
	ID            types.String      `tfsdk:"id" json:"id"`
	BindingStatus types.String      `tfsdk:"binding_status" json:"binding_status,computed"`
	UpdatedAt     timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed"`
}

type ZeroTrustGatewaySettingsSettingsExtendedEmailMatchingDataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled"`
}

type ZeroTrustGatewaySettingsSettingsFipsDataSourceModel struct {
	TLS types.Bool `tfsdk:"tls" json:"tls"`
}

type ZeroTrustGatewaySettingsSettingsProtocolDetectionDataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled"`
}

type ZeroTrustGatewaySettingsSettingsTLSDecryptDataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled"`
}
