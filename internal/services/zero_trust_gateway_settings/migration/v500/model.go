package v500

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source V4 Models (Legacy SDKv2 Provider - cloudflare_teams_account)
// ============================================================================

// SourceV4ZeroTrustGatewaySettingsModel represents the teams_account state from
// v4.x provider (SDKv2). Schema version: 0 (SDKv2 default).
//
// All TypeList MaxItems:1 blocks are stored as arrays in SDKv2 state.
// All fields from v4 schema are included, even those dropped in v5 (logging,
// proxy, ssh_session_log, payload_log), to allow the framework to parse v4 state.
type SourceV4ZeroTrustGatewaySettingsModel struct {
	AccountID types.String `tfsdk:"account_id"`

	// Flat boolean fields (v4 top-level → v5 nested under settings.*)
	ActivityLogEnabled                types.Bool `tfsdk:"activity_log_enabled"`
	TLSDecryptEnabled                 types.Bool `tfsdk:"tls_decrypt_enabled"`
	ProtocolDetectionEnabled          types.Bool `tfsdk:"protocol_detection_enabled"`
	URLBrowserIsolationEnabled        types.Bool `tfsdk:"url_browser_isolation_enabled"`
	NonIdentityBrowserIsolationEnabled types.Bool `tfsdk:"non_identity_browser_isolation_enabled"`

	// MaxItems:1 blocks (stored as arrays in v4 SDKv2 state)
	BlockPage             []SourceV4BlockPageModel             `tfsdk:"block_page"`
	BodyScanning          []SourceV4BodyScanningModel          `tfsdk:"body_scanning"`
	Fips                  []SourceV4FipsModel                  `tfsdk:"fips"`
	Antivirus             []SourceV4AntivirusModel             `tfsdk:"antivirus"`
	ExtendedEmailMatching []SourceV4ExtendedEmailMatchingModel `tfsdk:"extended_email_matching"`
	CustomCertificate     []SourceV4CustomCertificateModel     `tfsdk:"custom_certificate"`
	Certificate           []SourceV4CertificateModel           `tfsdk:"certificate"`

	// Blocks dropped in v5 (migrated to separate resources or removed entirely)
	// These must be in the source schema so the framework can parse v4 state.
	Logging       []SourceV4LoggingModel       `tfsdk:"logging"`
	Proxy         []SourceV4ProxyModel         `tfsdk:"proxy"`
	SSHSessionLog []SourceV4SSHSessionLogModel `tfsdk:"ssh_session_log"`
	PayloadLog    []SourceV4PayloadLogModel    `tfsdk:"payload_log"`
}

type SourceV4BlockPageModel struct {
	BackgroundColor types.String `tfsdk:"background_color"`
	Enabled         types.Bool   `tfsdk:"enabled"`
	FooterText      types.String `tfsdk:"footer_text"`
	HeaderText      types.String `tfsdk:"header_text"`
	LogoPath        types.String `tfsdk:"logo_path"`
	MailtoAddress   types.String `tfsdk:"mailto_address"`
	MailtoSubject   types.String `tfsdk:"mailto_subject"`
	Name            types.String `tfsdk:"name"`
}

type SourceV4BodyScanningModel struct {
	InspectionMode types.String `tfsdk:"inspection_mode"`
}

type SourceV4FipsModel struct {
	TLS types.Bool `tfsdk:"tls"`
}

// SourceV4AntivirusModel represents the antivirus block in v4.
// notification_settings is a nested TypeList MaxItems:1 → stored as array.
type SourceV4AntivirusModel struct {
	EnabledDownloadPhase types.Bool                       `tfsdk:"enabled_download_phase"`
	EnabledUploadPhase   types.Bool                       `tfsdk:"enabled_upload_phase"`
	FailClosed           types.Bool                       `tfsdk:"fail_closed"`
	NotificationSettings []SourceV4NotificationSettingsModel `tfsdk:"notification_settings"`
}

// SourceV4NotificationSettingsModel represents the notification_settings block in v4.
// Note: v4 field is "message"; v5 renames this to "msg".
type SourceV4NotificationSettingsModel struct {
	Enabled    types.Bool   `tfsdk:"enabled"`
	Message    types.String `tfsdk:"message"`
	SupportURL types.String `tfsdk:"support_url"`
}

type SourceV4ExtendedEmailMatchingModel struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type SourceV4CustomCertificateModel struct {
	Enabled   types.Bool   `tfsdk:"enabled"`
	ID        types.String `tfsdk:"id"`
	UpdatedAt types.String `tfsdk:"updated_at"` // Plain string in v4
}

type SourceV4CertificateModel struct {
	ID types.String `tfsdk:"id"`
}

// SourceV4LoggingModel is included for state parsing only; logging is
// migrated to cloudflare_zero_trust_gateway_logging by tf-migrate.
type SourceV4LoggingModel struct {
	RedactPII         types.Bool                          `tfsdk:"redact_pii"`
	SettingsByRuleType []SourceV4LoggingSettingsByRuleTypeModel `tfsdk:"settings_by_rule_type"`
}

type SourceV4LoggingSettingsByRuleTypeModel struct {
	DNS  []SourceV4LoggingEnabledModel `tfsdk:"dns"`
	HTTP []SourceV4LoggingEnabledModel `tfsdk:"http"`
	L4   []SourceV4LoggingEnabledModel `tfsdk:"l4"`
}

type SourceV4LoggingEnabledModel struct {
	LogAll    types.Bool `tfsdk:"log_all"`
	LogBlocks types.Bool `tfsdk:"log_blocks"`
}

// SourceV4ProxyModel is included for state parsing only; proxy is migrated to
// cloudflare_zero_trust_device_settings by tf-migrate.
type SourceV4ProxyModel struct {
	TCP           types.Bool  `tfsdk:"tcp"`
	UDP           types.Bool  `tfsdk:"udp"`
	RootCA        types.Bool  `tfsdk:"root_ca"`
	VirtualIP     types.Bool  `tfsdk:"virtual_ip"`
	DisableForTime types.Int64 `tfsdk:"disable_for_time"`
}

// SourceV4SSHSessionLogModel is included for state parsing only; dropped in v5.
type SourceV4SSHSessionLogModel struct {
	PublicKey types.String `tfsdk:"public_key"`
}

// SourceV4PayloadLogModel is included for state parsing only; dropped in v5.
type SourceV4PayloadLogModel struct {
	PublicKey types.String `tfsdk:"public_key"`
}

// ============================================================================
// Target V5 Models (Current Plugin Framework Provider)
// ============================================================================

// TargetV5ZeroTrustGatewaySettingsModel represents the zero_trust_gateway_settings
// state for v5.x+ provider (Plugin Framework). Schema version: 500.
type TargetV5ZeroTrustGatewaySettingsModel struct {
	ID        types.String                         `tfsdk:"id"`
	AccountID types.String                         `tfsdk:"account_id"`
	Settings  *TargetV5SettingsModel               `tfsdk:"settings"`
	CreatedAt timetypes.RFC3339                    `tfsdk:"created_at"`
	UpdatedAt timetypes.RFC3339                    `tfsdk:"updated_at"`
}

type TargetV5SettingsModel struct {
	ActivityLog           *TargetV5ActivityLogModel           `tfsdk:"activity_log"`
	Antivirus             *TargetV5AntivirusModel             `tfsdk:"antivirus"`
	BlockPage             *TargetV5BlockPageModel             `tfsdk:"block_page"`
	BodyScanning          *TargetV5BodyScanningModel          `tfsdk:"body_scanning"`
	BrowserIsolation      *TargetV5BrowserIsolationModel      `tfsdk:"browser_isolation"`
	Certificate           *TargetV5CertificateModel           `tfsdk:"certificate"`
	CustomCertificate     *TargetV5CustomCertificateModel     `tfsdk:"custom_certificate"`
	ExtendedEmailMatching *TargetV5ExtendedEmailMatchingModel `tfsdk:"extended_email_matching"`
	Fips                  *TargetV5FipsModel                  `tfsdk:"fips"`
	HostSelector          *TargetV5HostSelectorModel          `tfsdk:"host_selector"`
	Inspection            *TargetV5InspectionModel            `tfsdk:"inspection"`
	ProtocolDetection     *TargetV5ProtocolDetectionModel     `tfsdk:"protocol_detection"`
	Sandbox               *TargetV5SandboxModel               `tfsdk:"sandbox"`
	TLSDecrypt            *TargetV5TLSDecryptModel            `tfsdk:"tls_decrypt"`
}

type TargetV5ActivityLogModel struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

// TargetV5AntivirusModel represents the v5 antivirus settings.
// notification_settings uses customfield.NestedObject because the v5 schema
// uses customfield.NewNestedObjectType for this attribute.
type TargetV5AntivirusModel struct {
	EnabledDownloadPhase types.Bool                                                         `tfsdk:"enabled_download_phase"`
	EnabledUploadPhase   types.Bool                                                         `tfsdk:"enabled_upload_phase"`
	FailClosed           types.Bool                                                         `tfsdk:"fail_closed"`
	NotificationSettings customfield.NestedObject[TargetV5NotificationSettingsModel]        `tfsdk:"notification_settings"`
}

// TargetV5NotificationSettingsModel represents the v5 notification_settings.
// Note: v4 had "message" but v5 uses "msg" (field rename).
// Also new in v5: include_context (not present in v4, will be null).
type TargetV5NotificationSettingsModel struct {
	Enabled        types.Bool   `tfsdk:"enabled"`
	IncludeContext types.Bool   `tfsdk:"include_context"`
	Msg            types.String `tfsdk:"msg"`
	SupportURL     types.String `tfsdk:"support_url"`
}

type TargetV5BlockPageModel struct {
	BackgroundColor types.String `tfsdk:"background_color"`
	Enabled         types.Bool   `tfsdk:"enabled"`
	FooterText      types.String `tfsdk:"footer_text"`
	HeaderText      types.String `tfsdk:"header_text"`
	IncludeContext  types.Bool   `tfsdk:"include_context"`
	LogoPath        types.String `tfsdk:"logo_path"`
	MailtoAddress   types.String `tfsdk:"mailto_address"`
	MailtoSubject   types.String `tfsdk:"mailto_subject"`
	Mode            types.String `tfsdk:"mode"`
	Name            types.String `tfsdk:"name"`
	ReadOnly        types.Bool   `tfsdk:"read_only"`
	SourceAccount   types.String `tfsdk:"source_account"`
	SuppressFooter  types.Bool   `tfsdk:"suppress_footer"`
	TargetURI       types.String `tfsdk:"target_uri"`
	Version         types.Int64  `tfsdk:"version"`
}

type TargetV5BodyScanningModel struct {
	InspectionMode types.String `tfsdk:"inspection_mode"`
}

type TargetV5BrowserIsolationModel struct {
	NonIdentityEnabled         types.Bool `tfsdk:"non_identity_enabled"`
	URLBrowserIsolationEnabled types.Bool `tfsdk:"url_browser_isolation_enabled"`
}

type TargetV5CertificateModel struct {
	ID types.String `tfsdk:"id"`
}

type TargetV5CustomCertificateModel struct {
	Enabled       types.Bool        `tfsdk:"enabled"`
	ID            types.String      `tfsdk:"id"`
	BindingStatus types.String      `tfsdk:"binding_status"`
	UpdatedAt     timetypes.RFC3339 `tfsdk:"updated_at"`
}

type TargetV5ExtendedEmailMatchingModel struct {
	Enabled       types.Bool   `tfsdk:"enabled"`
	ReadOnly      types.Bool   `tfsdk:"read_only"`
	SourceAccount types.String `tfsdk:"source_account"`
	Version       types.Int64  `tfsdk:"version"`
}

type TargetV5FipsModel struct {
	TLS types.Bool `tfsdk:"tls"`
}

type TargetV5HostSelectorModel struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type TargetV5InspectionModel struct {
	Mode types.String `tfsdk:"mode"`
}

type TargetV5ProtocolDetectionModel struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type TargetV5SandboxModel struct {
	Enabled        types.Bool   `tfsdk:"enabled"`
	FallbackAction types.String `tfsdk:"fallback_action"`
}

type TargetV5TLSDecryptModel struct {
	Enabled types.Bool `tfsdk:"enabled"`
}
