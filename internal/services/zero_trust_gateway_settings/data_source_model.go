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
	ActivityLog           *ZeroTrustGatewaySettingsSettingsActivityLogDataSourceModel           `tfsdk:"activity_log" json:"activity_log,computed_optional"`
	Antivirus             *ZeroTrustGatewaySettingsSettingsAntivirusDataSourceModel             `tfsdk:"antivirus" json:"antivirus,computed_optional"`
	BlockPage             *ZeroTrustGatewaySettingsSettingsBlockPageDataSourceModel             `tfsdk:"block_page" json:"block_page,computed_optional"`
	BodyScanning          *ZeroTrustGatewaySettingsSettingsBodyScanningDataSourceModel          `tfsdk:"body_scanning" json:"body_scanning,computed_optional"`
	BrowserIsolation      *ZeroTrustGatewaySettingsSettingsBrowserIsolationDataSourceModel      `tfsdk:"browser_isolation" json:"browser_isolation,computed_optional"`
	Certificate           *ZeroTrustGatewaySettingsSettingsCertificateDataSourceModel           `tfsdk:"certificate" json:"certificate,computed_optional"`
	CustomCertificate     *ZeroTrustGatewaySettingsSettingsCustomCertificateDataSourceModel     `tfsdk:"custom_certificate" json:"custom_certificate,computed_optional"`
	ExtendedEmailMatching *ZeroTrustGatewaySettingsSettingsExtendedEmailMatchingDataSourceModel `tfsdk:"extended_email_matching" json:"extended_email_matching,computed_optional"`
	Fips                  *ZeroTrustGatewaySettingsSettingsFipsDataSourceModel                  `tfsdk:"fips" json:"fips,computed_optional"`
	ProtocolDetection     *ZeroTrustGatewaySettingsSettingsProtocolDetectionDataSourceModel     `tfsdk:"protocol_detection" json:"protocol_detection,computed_optional"`
	TLSDecrypt            *ZeroTrustGatewaySettingsSettingsTLSDecryptDataSourceModel            `tfsdk:"tls_decrypt" json:"tls_decrypt,computed_optional"`
}

type ZeroTrustGatewaySettingsSettingsActivityLogDataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed_optional"`
}

type ZeroTrustGatewaySettingsSettingsAntivirusDataSourceModel struct {
	EnabledDownloadPhase types.Bool                                                                    `tfsdk:"enabled_download_phase" json:"enabled_download_phase,computed_optional"`
	EnabledUploadPhase   types.Bool                                                                    `tfsdk:"enabled_upload_phase" json:"enabled_upload_phase,computed_optional"`
	FailClosed           types.Bool                                                                    `tfsdk:"fail_closed" json:"fail_closed,computed_optional"`
	NotificationSettings *ZeroTrustGatewaySettingsSettingsAntivirusNotificationSettingsDataSourceModel `tfsdk:"notification_settings" json:"notification_settings,computed_optional"`
}

type ZeroTrustGatewaySettingsSettingsAntivirusNotificationSettingsDataSourceModel struct {
	Enabled    types.Bool   `tfsdk:"enabled" json:"enabled,computed_optional"`
	Msg        types.String `tfsdk:"msg" json:"msg,computed_optional"`
	SupportURL types.String `tfsdk:"support_url" json:"support_url,computed_optional"`
}

type ZeroTrustGatewaySettingsSettingsBlockPageDataSourceModel struct {
	BackgroundColor types.String `tfsdk:"background_color" json:"background_color,computed_optional"`
	Enabled         types.Bool   `tfsdk:"enabled" json:"enabled,computed_optional"`
	FooterText      types.String `tfsdk:"footer_text" json:"footer_text,computed_optional"`
	HeaderText      types.String `tfsdk:"header_text" json:"header_text,computed_optional"`
	LogoPath        types.String `tfsdk:"logo_path" json:"logo_path,computed_optional"`
	MailtoAddress   types.String `tfsdk:"mailto_address" json:"mailto_address,computed_optional"`
	MailtoSubject   types.String `tfsdk:"mailto_subject" json:"mailto_subject,computed_optional"`
	Name            types.String `tfsdk:"name" json:"name,computed_optional"`
	SuppressFooter  types.Bool   `tfsdk:"suppress_footer" json:"suppress_footer,computed_optional"`
}

type ZeroTrustGatewaySettingsSettingsBodyScanningDataSourceModel struct {
	InspectionMode types.String `tfsdk:"inspection_mode" json:"inspection_mode,computed_optional"`
}

type ZeroTrustGatewaySettingsSettingsBrowserIsolationDataSourceModel struct {
	NonIdentityEnabled         types.Bool `tfsdk:"non_identity_enabled" json:"non_identity_enabled,computed_optional"`
	URLBrowserIsolationEnabled types.Bool `tfsdk:"url_browser_isolation_enabled" json:"url_browser_isolation_enabled,computed_optional"`
}

type ZeroTrustGatewaySettingsSettingsCertificateDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustGatewaySettingsSettingsCustomCertificateDataSourceModel struct {
	Enabled       types.Bool        `tfsdk:"enabled" json:"enabled,computed"`
	ID            types.String      `tfsdk:"id" json:"id,computed_optional"`
	BindingStatus types.String      `tfsdk:"binding_status" json:"binding_status,computed"`
	UpdatedAt     timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed"`
}

type ZeroTrustGatewaySettingsSettingsExtendedEmailMatchingDataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed_optional"`
}

type ZeroTrustGatewaySettingsSettingsFipsDataSourceModel struct {
	TLS types.Bool `tfsdk:"tls" json:"tls,computed_optional"`
}

type ZeroTrustGatewaySettingsSettingsProtocolDetectionDataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed_optional"`
}

type ZeroTrustGatewaySettingsSettingsTLSDecryptDataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed_optional"`
}
