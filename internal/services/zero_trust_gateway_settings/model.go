// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_settings

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustGatewaySettingsResultEnvelope struct {
	Result ZeroTrustGatewaySettingsModel `json:"result"`
}

type ZeroTrustGatewaySettingsModel struct {
	ID        types.String                                                    `tfsdk:"id" json:"-,computed"`
	AccountID types.String                                                    `tfsdk:"account_id" path:"account_id,required"`
	Settings  customfield.NestedObject[ZeroTrustGatewaySettingsSettingsModel] `tfsdk:"settings" json:"settings,computed_optional"`
	CreatedAt timetypes.RFC3339                                               `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UpdatedAt timetypes.RFC3339                                               `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m ZeroTrustGatewaySettingsModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustGatewaySettingsModel) MarshalJSONForUpdate(state ZeroTrustGatewaySettingsModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustGatewaySettingsSettingsModel struct {
	ActivityLog           customfield.NestedObject[ZeroTrustGatewaySettingsSettingsActivityLogModel]           `tfsdk:"activity_log" json:"activity_log,computed_optional"`
	Antivirus             customfield.NestedObject[ZeroTrustGatewaySettingsSettingsAntivirusModel]             `tfsdk:"antivirus" json:"antivirus,computed_optional"`
	BlockPage             customfield.NestedObject[ZeroTrustGatewaySettingsSettingsBlockPageModel]             `tfsdk:"block_page" json:"block_page,computed_optional"`
	BodyScanning          customfield.NestedObject[ZeroTrustGatewaySettingsSettingsBodyScanningModel]          `tfsdk:"body_scanning" json:"body_scanning,computed_optional"`
	BrowserIsolation      customfield.NestedObject[ZeroTrustGatewaySettingsSettingsBrowserIsolationModel]      `tfsdk:"browser_isolation" json:"browser_isolation,computed_optional"`
	Certificate           customfield.NestedObject[ZeroTrustGatewaySettingsSettingsCertificateModel]           `tfsdk:"certificate" json:"certificate,computed_optional"`
	CustomCertificate     customfield.NestedObject[ZeroTrustGatewaySettingsSettingsCustomCertificateModel]     `tfsdk:"custom_certificate" json:"custom_certificate,computed_optional"`
	ExtendedEmailMatching customfield.NestedObject[ZeroTrustGatewaySettingsSettingsExtendedEmailMatchingModel] `tfsdk:"extended_email_matching" json:"extended_email_matching,computed_optional"`
	Fips                  customfield.NestedObject[ZeroTrustGatewaySettingsSettingsFipsModel]                  `tfsdk:"fips" json:"fips,computed_optional"`
	ProtocolDetection     customfield.NestedObject[ZeroTrustGatewaySettingsSettingsProtocolDetectionModel]     `tfsdk:"protocol_detection" json:"protocol_detection,computed_optional"`
	TLSDecrypt            customfield.NestedObject[ZeroTrustGatewaySettingsSettingsTLSDecryptModel]            `tfsdk:"tls_decrypt" json:"tls_decrypt,computed_optional"`
}

type ZeroTrustGatewaySettingsSettingsActivityLogModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,optional"`
}

type ZeroTrustGatewaySettingsSettingsAntivirusModel struct {
	EnabledDownloadPhase types.Bool                                                                                   `tfsdk:"enabled_download_phase" json:"enabled_download_phase,optional"`
	EnabledUploadPhase   types.Bool                                                                                   `tfsdk:"enabled_upload_phase" json:"enabled_upload_phase,optional"`
	FailClosed           types.Bool                                                                                   `tfsdk:"fail_closed" json:"fail_closed,optional"`
	NotificationSettings customfield.NestedObject[ZeroTrustGatewaySettingsSettingsAntivirusNotificationSettingsModel] `tfsdk:"notification_settings" json:"notification_settings,computed_optional"`
}

type ZeroTrustGatewaySettingsSettingsAntivirusNotificationSettingsModel struct {
	Enabled    types.Bool   `tfsdk:"enabled" json:"enabled,optional"`
	Msg        types.String `tfsdk:"msg" json:"msg,optional"`
	SupportURL types.String `tfsdk:"support_url" json:"support_url,optional"`
}

type ZeroTrustGatewaySettingsSettingsBlockPageModel struct {
	BackgroundColor types.String `tfsdk:"background_color" json:"background_color,optional"`
	Enabled         types.Bool   `tfsdk:"enabled" json:"enabled,optional"`
	FooterText      types.String `tfsdk:"footer_text" json:"footer_text,optional"`
	HeaderText      types.String `tfsdk:"header_text" json:"header_text,optional"`
	LogoPath        types.String `tfsdk:"logo_path" json:"logo_path,optional"`
	MailtoAddress   types.String `tfsdk:"mailto_address" json:"mailto_address,optional"`
	MailtoSubject   types.String `tfsdk:"mailto_subject" json:"mailto_subject,optional"`
	Name            types.String `tfsdk:"name" json:"name,optional"`
	SuppressFooter  types.Bool   `tfsdk:"suppress_footer" json:"suppress_footer,optional"`
}

type ZeroTrustGatewaySettingsSettingsBodyScanningModel struct {
	InspectionMode types.String `tfsdk:"inspection_mode" json:"inspection_mode,optional"`
}

type ZeroTrustGatewaySettingsSettingsBrowserIsolationModel struct {
	NonIdentityEnabled         types.Bool `tfsdk:"non_identity_enabled" json:"non_identity_enabled,optional"`
	URLBrowserIsolationEnabled types.Bool `tfsdk:"url_browser_isolation_enabled" json:"url_browser_isolation_enabled,optional"`
}

type ZeroTrustGatewaySettingsSettingsCertificateModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustGatewaySettingsSettingsCustomCertificateModel struct {
	Enabled       types.Bool        `tfsdk:"enabled" json:"enabled,required"`
	ID            types.String      `tfsdk:"id" json:"id,optional"`
	BindingStatus types.String      `tfsdk:"binding_status" json:"binding_status,computed"`
	UpdatedAt     timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

type ZeroTrustGatewaySettingsSettingsExtendedEmailMatchingModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,optional"`
}

type ZeroTrustGatewaySettingsSettingsFipsModel struct {
	TLS types.Bool `tfsdk:"tls" json:"tls,optional"`
}

type ZeroTrustGatewaySettingsSettingsProtocolDetectionModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,optional"`
}

type ZeroTrustGatewaySettingsSettingsTLSDecryptModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,optional"`
}
