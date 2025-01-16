// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_settings

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustGatewaySettingsResultDataSourceEnvelope struct {
	Result ZeroTrustGatewaySettingsDataSourceModel `json:"result,computed"`
}

type ZeroTrustGatewaySettingsDataSourceModel struct {
	AccountID types.String                                                              `tfsdk:"account_id" path:"account_id,required"`
	CreatedAt timetypes.RFC3339                                                         `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UpdatedAt timetypes.RFC3339                                                         `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Settings  customfield.NestedObject[ZeroTrustGatewaySettingsSettingsDataSourceModel] `tfsdk:"settings" json:"settings,computed"`
}

func (m *ZeroTrustGatewaySettingsDataSourceModel) toReadParams(_ context.Context) (params zero_trust.GatewayConfigurationGetParams, diags diag.Diagnostics) {
	params = zero_trust.GatewayConfigurationGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustGatewaySettingsSettingsDataSourceModel struct {
	ActivityLog           customfield.NestedObject[ZeroTrustGatewaySettingsSettingsActivityLogDataSourceModel]           `tfsdk:"activity_log" json:"activity_log,computed"`
	Antivirus             customfield.NestedObject[ZeroTrustGatewaySettingsSettingsAntivirusDataSourceModel]             `tfsdk:"antivirus" json:"antivirus,computed"`
	BlockPage             customfield.NestedObject[ZeroTrustGatewaySettingsSettingsBlockPageDataSourceModel]             `tfsdk:"block_page" json:"block_page,computed"`
	BodyScanning          customfield.NestedObject[ZeroTrustGatewaySettingsSettingsBodyScanningDataSourceModel]          `tfsdk:"body_scanning" json:"body_scanning,computed"`
	BrowserIsolation      customfield.NestedObject[ZeroTrustGatewaySettingsSettingsBrowserIsolationDataSourceModel]      `tfsdk:"browser_isolation" json:"browser_isolation,computed"`
	Certificate           customfield.NestedObject[ZeroTrustGatewaySettingsSettingsCertificateDataSourceModel]           `tfsdk:"certificate" json:"certificate,computed"`
	CustomCertificate     customfield.NestedObject[ZeroTrustGatewaySettingsSettingsCustomCertificateDataSourceModel]     `tfsdk:"custom_certificate" json:"custom_certificate,computed"`
	ExtendedEmailMatching customfield.NestedObject[ZeroTrustGatewaySettingsSettingsExtendedEmailMatchingDataSourceModel] `tfsdk:"extended_email_matching" json:"extended_email_matching,computed"`
	Fips                  customfield.NestedObject[ZeroTrustGatewaySettingsSettingsFipsDataSourceModel]                  `tfsdk:"fips" json:"fips,computed"`
	ProtocolDetection     customfield.NestedObject[ZeroTrustGatewaySettingsSettingsProtocolDetectionDataSourceModel]     `tfsdk:"protocol_detection" json:"protocol_detection,computed"`
	Sandbox               customfield.NestedObject[ZeroTrustGatewaySettingsSettingsSandboxDataSourceModel]               `tfsdk:"sandbox" json:"sandbox,computed"`
	TLSDecrypt            customfield.NestedObject[ZeroTrustGatewaySettingsSettingsTLSDecryptDataSourceModel]            `tfsdk:"tls_decrypt" json:"tls_decrypt,computed"`
}

type ZeroTrustGatewaySettingsSettingsActivityLogDataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed"`
}

type ZeroTrustGatewaySettingsSettingsAntivirusDataSourceModel struct {
	EnabledDownloadPhase types.Bool                                                                                             `tfsdk:"enabled_download_phase" json:"enabled_download_phase,computed"`
	EnabledUploadPhase   types.Bool                                                                                             `tfsdk:"enabled_upload_phase" json:"enabled_upload_phase,computed"`
	FailClosed           types.Bool                                                                                             `tfsdk:"fail_closed" json:"fail_closed,computed"`
	NotificationSettings customfield.NestedObject[ZeroTrustGatewaySettingsSettingsAntivirusNotificationSettingsDataSourceModel] `tfsdk:"notification_settings" json:"notification_settings,computed"`
}

type ZeroTrustGatewaySettingsSettingsAntivirusNotificationSettingsDataSourceModel struct {
	Enabled    types.Bool   `tfsdk:"enabled" json:"enabled,computed"`
	Msg        types.String `tfsdk:"msg" json:"msg,computed"`
	SupportURL types.String `tfsdk:"support_url" json:"support_url,computed"`
}

type ZeroTrustGatewaySettingsSettingsBlockPageDataSourceModel struct {
	BackgroundColor types.String `tfsdk:"background_color" json:"background_color,computed"`
	Enabled         types.Bool   `tfsdk:"enabled" json:"enabled,computed"`
	FooterText      types.String `tfsdk:"footer_text" json:"footer_text,computed"`
	HeaderText      types.String `tfsdk:"header_text" json:"header_text,computed"`
	LogoPath        types.String `tfsdk:"logo_path" json:"logo_path,computed"`
	MailtoAddress   types.String `tfsdk:"mailto_address" json:"mailto_address,computed"`
	MailtoSubject   types.String `tfsdk:"mailto_subject" json:"mailto_subject,computed"`
	Name            types.String `tfsdk:"name" json:"name,computed"`
	SuppressFooter  types.Bool   `tfsdk:"suppress_footer" json:"suppress_footer,computed"`
}

type ZeroTrustGatewaySettingsSettingsBodyScanningDataSourceModel struct {
	InspectionMode types.String `tfsdk:"inspection_mode" json:"inspection_mode,computed"`
}

type ZeroTrustGatewaySettingsSettingsBrowserIsolationDataSourceModel struct {
	NonIdentityEnabled         types.Bool `tfsdk:"non_identity_enabled" json:"non_identity_enabled,computed"`
	URLBrowserIsolationEnabled types.Bool `tfsdk:"url_browser_isolation_enabled" json:"url_browser_isolation_enabled,computed"`
}

type ZeroTrustGatewaySettingsSettingsCertificateDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustGatewaySettingsSettingsCustomCertificateDataSourceModel struct {
	Enabled       types.Bool        `tfsdk:"enabled" json:"enabled,computed"`
	ID            types.String      `tfsdk:"id" json:"id,computed"`
	BindingStatus types.String      `tfsdk:"binding_status" json:"binding_status,computed"`
	UpdatedAt     timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

type ZeroTrustGatewaySettingsSettingsExtendedEmailMatchingDataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed"`
}

type ZeroTrustGatewaySettingsSettingsFipsDataSourceModel struct {
	TLS types.Bool `tfsdk:"tls" json:"tls,computed"`
}

type ZeroTrustGatewaySettingsSettingsProtocolDetectionDataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed"`
}

type ZeroTrustGatewaySettingsSettingsSandboxDataSourceModel struct {
	Enabled        types.Bool   `tfsdk:"enabled" json:"enabled,computed"`
	FallbackAction types.String `tfsdk:"fallback_action" json:"fallback_action,computed"`
}

type ZeroTrustGatewaySettingsSettingsTLSDecryptDataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed"`
}
