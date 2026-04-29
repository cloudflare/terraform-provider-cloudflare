package v500

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x - SDKv2)
// ============================================================================

// SourceCloudflareTeamsRuleModel represents the legacy cloudflare_teams_rule state from v4.x provider.
// Schema version: 0 (implicit in v4 SDKv2)
// Resource type: cloudflare_teams_rule
//
// IMPORTANT: v4 SDKv2 stores all TypeList MaxItems:1 blocks as ARRAYS in state.
// All nested structures within rule_settings are arrays that need to be converted to objects.
type SourceCloudflareTeamsRuleModel struct {
	// Core fields
	ID            types.String `tfsdk:"id"`
	AccountID     types.String `tfsdk:"account_id"`
	Name          types.String `tfsdk:"name"`
	Description   types.String `tfsdk:"description"`
	Precedence    types.Int64  `tfsdk:"precedence"`
	Enabled       types.Bool   `tfsdk:"enabled"`
	Action        types.String `tfsdk:"action"`
	Filters       types.List   `tfsdk:"filters"` // List of strings
	Traffic       types.String `tfsdk:"traffic"`
	Identity      types.String `tfsdk:"identity"`
	DevicePosture types.String `tfsdk:"device_posture"`
	Version       types.Int64  `tfsdk:"version"` // Computed field, may not be in v5 schema

	// rule_settings - TypeList MaxItems:1 in v4, stored as array
	RuleSettings []SourceRuleSettingsModel `tfsdk:"rule_settings"`
}

// SourceRuleSettingsModel represents the rule_settings block from v4.
// In v4, this is TypeList MaxItems:1, stored as array with single element.
// In v5, this becomes SingleNestedAttribute (object).
type SourceRuleSettingsModel struct {
	// Simple scalar fields
	BlockPageEnabled                types.Bool              `tfsdk:"block_page_enabled"`
	BlockPageReason                 types.String            `tfsdk:"block_page_reason"` // Renamed to block_reason in v5
	OverrideIPs                     types.List              `tfsdk:"override_ips"`      // List of strings
	OverrideHost                    types.String            `tfsdk:"override_host"`
	IPCategories                    types.Bool              `tfsdk:"ip_categories"`
	IgnoreCNAMECategoryMatches      types.Bool              `tfsdk:"ignore_cname_category_matches"`
	AllowChildBypass                types.Bool              `tfsdk:"allow_child_bypass"`
	BypassParentRule                types.Bool              `tfsdk:"bypass_parent_rule"`
	InsecureDisableDNSSECValidation types.Bool              `tfsdk:"insecure_disable_dnssec_validation"`
	ResolveDNSThroughCloudflare     types.Bool              `tfsdk:"resolve_dns_through_cloudflare"`
	AddHeaders                      *map[string]types.String `tfsdk:"add_headers"` // v4: Map[string]string

	// All nested structures as arrays (TypeList MaxItems:1 in SDKv2)
	AuditSSH             []SourceAuditSSHModel             `tfsdk:"audit_ssh"`
	L4override           []SourceL4overrideModel           `tfsdk:"l4override"`
	BISOAdminControls    []SourceBISOAdminControlsModel    `tfsdk:"biso_admin_controls"`
	CheckSession         []SourceCheckSessionModel         `tfsdk:"check_session"`
	Egress               []SourceEgressModel               `tfsdk:"egress"`
	UntrustedCERT        []SourceUntrustedCERTModel        `tfsdk:"untrusted_cert"`
	PayloadLog           []SourcePayloadLogModel           `tfsdk:"payload_log"`
	NotificationSettings []SourceNotificationSettingsModel `tfsdk:"notification_settings"`
	DNSResolvers         []SourceDNSResolversModel         `tfsdk:"dns_resolvers"`
	ResolveDNSInternally []SourceResolveDNSInternallyModel `tfsdk:"resolve_dns_internally"`
}

// SourceAuditSSHModel represents audit_ssh settings from v4.
type SourceAuditSSHModel struct {
	CommandLogging types.Bool `tfsdk:"command_logging"`
}

// SourceL4overrideModel represents l4override settings from v4.
type SourceL4overrideModel struct {
	IP   types.String `tfsdk:"ip"`
	Port types.Int64  `tfsdk:"port"` // Int64 in v4, Int64 in v5 (stored as float64 in state)
}

// SourceBISOAdminControlsModel represents biso_admin_controls from v4.
// IMPORTANT: Contains BOTH v1 (deprecated disable_*) and v2 (string-based) fields.
// v1 fields will be REMOVED during transformation.
type SourceBISOAdminControlsModel struct {
	Version types.String `tfsdk:"version"` // "v1" or "v2"

	// v1 fields (will be RENAMED to shortened versions in transformation)
	// These map to: dp, dcp, dd, dk, du in v5
	DisablePrinting             types.Bool `tfsdk:"disable_printing"`              // → dp
	DisableCopyPaste            types.Bool `tfsdk:"disable_copy_paste"`            // → dcp
	DisableDownload             types.Bool `tfsdk:"disable_download"`              // → dd
	DisableKeyboard             types.Bool `tfsdk:"disable_keyboard"`              // → dk
	DisableUpload               types.Bool `tfsdk:"disable_upload"`                // → du
	DisableClipboardRedirection types.Bool `tfsdk:"disable_clipboard_redirection"` // → removed (no v5 equivalent)

	// v2 fields (unchanged - exist in both v4 and v5 with same names)
	Copy     types.String `tfsdk:"copy"`     // "enabled", "disabled", "remote_only"
	Download types.String `tfsdk:"download"` // "enabled", "disabled", "remote_only"
	Keyboard types.String `tfsdk:"keyboard"` // "enabled", "disabled"
	Paste    types.String `tfsdk:"paste"`    // "enabled", "disabled", "remote_only"
	Printing types.String `tfsdk:"printing"` // "enabled", "disabled"
	Upload   types.String `tfsdk:"upload"`   // "enabled", "disabled"
}

// SourceCheckSessionModel represents check_session settings from v4.
type SourceCheckSessionModel struct {
	Enforce  types.Bool   `tfsdk:"enforce"`
	Duration types.String `tfsdk:"duration"` // e.g., "24h0m0s" (may be normalized to "24h" by API)
}

// SourceEgressModel represents egress settings from v4.
type SourceEgressModel struct {
	IPV6         types.String `tfsdk:"ipv6"`
	IPV4         types.String `tfsdk:"ipv4"`
	IPV4Fallback types.String `tfsdk:"ipv4_fallback"`
}

// SourceUntrustedCERTModel represents untrusted_cert settings from v4.
type SourceUntrustedCERTModel struct {
	Action types.String `tfsdk:"action"`
}

// SourcePayloadLogModel represents payload_log settings from v4.
type SourcePayloadLogModel struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

// SourceNotificationSettingsModel represents notification_settings from v4.
type SourceNotificationSettingsModel struct {
	Enabled    types.Bool   `tfsdk:"enabled"`
	Message    types.String `tfsdk:"message"` // Renamed to msg in v5
	SupportURL types.String `tfsdk:"support_url"`
}

// SourceDNSResolversModel represents dns_resolvers from v4.
// IMPORTANT: This is TypeList MaxItems:1 at the dns_resolvers level,
// but IPV4 and IPV6 are actual arrays (not MaxItems:1).
type SourceDNSResolversModel struct {
	IPV4 []SourceDNSResolversIPV4Model `tfsdk:"ipv4"` // Array of resolvers
	IPV6 []SourceDNSResolversIPV6Model `tfsdk:"ipv6"` // Array of resolvers
}

// SourceDNSResolversIPV4Model represents a single IPv4 resolver address.
type SourceDNSResolversIPV4Model struct {
	IP                         types.String `tfsdk:"ip"`
	Port                       types.Int64  `tfsdk:"port"` // Int64 in v4
	VnetID                     types.String `tfsdk:"vnet_id"`
	RouteThroughPrivateNetwork types.Bool   `tfsdk:"route_through_private_network"`
}

// SourceDNSResolversIPV6Model represents a single IPv6 resolver address.
type SourceDNSResolversIPV6Model struct {
	IP                         types.String `tfsdk:"ip"`
	Port                       types.Int64  `tfsdk:"port"` // Int64 in v4
	VnetID                     types.String `tfsdk:"vnet_id"`
	RouteThroughPrivateNetwork types.Bool   `tfsdk:"route_through_private_network"`
}

// SourceResolveDNSInternallyModel represents resolve_dns_internally from v4.
type SourceResolveDNSInternallyModel struct {
	ViewID   types.String `tfsdk:"view_id"`
	Fallback types.String `tfsdk:"fallback"`
}

// ============================================================================
// Target Models (Current Provider - v5.x+ - Framework)
// ============================================================================
//
// These models use type aliases to the actual v5 models from the parent package.
// This avoids duplication and ensures we match the exact structure.
//
// Note: We import the parent package types, which uses:
// - customfield.NestedObject for rule_settings at root
// - Pointers for all nested structures within rule_settings
// ============================================================================

// Import the actual v5 models from parent package
// (These are already defined in internal/services/zero_trust_gateway_policy/model.go)

// TargetZeroTrustGatewayPolicyModel is an alias to the actual v5 model.
// This matches ZeroTrustGatewayPolicyModel from the parent package.
type TargetZeroTrustGatewayPolicyModel struct {
	ID            types.String                                                      `tfsdk:"id"`
	AccountID     types.String                                                      `tfsdk:"account_id"`
	Action        types.String                                                      `tfsdk:"action"`
	Name          types.String                                                      `tfsdk:"name"`
	Description   types.String                                                      `tfsdk:"description"`
	Filters       *[]types.String                                                   `tfsdk:"filters"`
	DevicePosture types.String                                                      `tfsdk:"device_posture"`
	Enabled       types.Bool                                                        `tfsdk:"enabled"`
	Identity      types.String                                                      `tfsdk:"identity"`
	Precedence    types.Int64                                                       `tfsdk:"precedence"`
	Traffic       types.String                                                      `tfsdk:"traffic"`
	Expiration    customfield.NestedObject[TargetExpirationModel]                   `tfsdk:"expiration"`    // New in v5
	RuleSettings  customfield.NestedObject[TargetRuleSettingsModel]                 `tfsdk:"rule_settings"` // Changed from array to object
	Schedule      customfield.NestedObject[TargetScheduleModel]                     `tfsdk:"schedule"`      // New in v5
	CreatedAt     timetypes.RFC3339                                                 `tfsdk:"created_at"`    // New in v5
	DeletedAt     timetypes.RFC3339                                                 `tfsdk:"deleted_at"`    // New in v5
	ReadOnly      types.Bool                                                        `tfsdk:"read_only"`     // New in v5
	Sharable      types.Bool                                                        `tfsdk:"sharable"`      // New in v5
	SourceAccount types.String                                                      `tfsdk:"source_account"` // New in v5
	UpdatedAt     timetypes.RFC3339                                                 `tfsdk:"updated_at"`    // New in v5
	Version       types.Int64                                                       `tfsdk:"version"`       // Computed
	WarningStatus types.String                                                      `tfsdk:"warning_status"` // New in v5
}

// TargetExpirationModel represents expiration settings (new in v5).
type TargetExpirationModel struct {
	ExpiresAt timetypes.RFC3339 `tfsdk:"expires_at"`
	Duration  types.Int64       `tfsdk:"duration"`
	Expired   types.Bool        `tfsdk:"expired"`
}

// TargetRuleSettingsModel represents rule_settings in v5.
// This matches ZeroTrustGatewayPolicyRuleSettingsModel from the parent package.
type TargetRuleSettingsModel struct {
	AddHeaders                      *map[string]*[]types.String             `tfsdk:"add_headers"` // v5: Map[string][]string
	AllowChildBypass                types.Bool                              `tfsdk:"allow_child_bypass"`
	AuditSSH                        *TargetAuditSSHModel                    `tfsdk:"audit_ssh"`
	BISOAdminControls               *TargetBISOAdminControlsModel           `tfsdk:"biso_admin_controls"`
	BlockPage                       *TargetBlockPageModel                   `tfsdk:"block_page"`        // New in v5
	BlockPageEnabled                types.Bool                              `tfsdk:"block_page_enabled"`
	BlockReason                     types.String                            `tfsdk:"block_reason"`      // Renamed from block_page_reason
	BypassParentRule                types.Bool                              `tfsdk:"bypass_parent_rule"`
	CheckSession                    *TargetCheckSessionModel                `tfsdk:"check_session"`
	DNSResolvers                    *TargetDNSResolversModel                `tfsdk:"dns_resolvers"`
	Egress                          *TargetEgressModel                      `tfsdk:"egress"`
	ForensicCopy                    *TargetForensicCopyModel                `tfsdk:"forensic_copy"` // New in v5
	IgnoreCNAMECategoryMatches      types.Bool                              `tfsdk:"ignore_cname_category_matches"`
	InsecureDisableDNSSECValidation types.Bool                              `tfsdk:"insecure_disable_dnssec_validation"`
	IPCategories                    types.Bool                              `tfsdk:"ip_categories"`
	IPIndicatorFeeds                types.Bool                              `tfsdk:"ip_indicator_feeds"` // New in v5
	L4override                      *TargetL4overrideModel                  `tfsdk:"l4override"`
	NotificationSettings            *TargetNotificationSettingsModel        `tfsdk:"notification_settings"`
	OverrideHost                    types.String                            `tfsdk:"override_host"`
	OverrideIPs                     customfield.List[types.String]          `tfsdk:"override_ips"`
	PayloadLog                      *TargetPayloadLogModel                  `tfsdk:"payload_log"`
	Quarantine                      *TargetQuarantineModel                  `tfsdk:"quarantine"`        // New in v5
	Redirect                        *TargetRedirectModel                    `tfsdk:"redirect"`          // New in v5
	ResolveDNSInternally            *TargetResolveDNSInternallyModel        `tfsdk:"resolve_dns_internally"`
	ResolveDNSThroughCloudflare     types.Bool                              `tfsdk:"resolve_dns_through_cloudflare"`
	UntrustedCERT                   *TargetUntrustedCERTModel               `tfsdk:"untrusted_cert"`
}

// TargetAuditSSHModel represents audit_ssh in v5.
type TargetAuditSSHModel struct {
	CommandLogging types.Bool `tfsdk:"command_logging"`
}

// TargetBISOAdminControlsModel represents biso_admin_controls in v5.
// IMPORTANT: v1 deprecated fields (disable_*) are NOT present.
// Only v2 string-based fields and internal fields.
type TargetBISOAdminControlsModel struct {
	// v2 fields (string-based, unchanged from v4)
	Copy     types.String `tfsdk:"copy"`     // "enabled", "disabled", "remote_only" (v2 only)
	Download types.String `tfsdk:"download"` // "enabled", "disabled", "remote_only" (v2 only)
	Keyboard types.String `tfsdk:"keyboard"` // "enabled", "disabled" (v2 only)
	Paste    types.String `tfsdk:"paste"`    // "enabled", "disabled", "remote_only" (v2 only)
	Printing types.String `tfsdk:"printing"` // "enabled", "disabled" (v2 only)
	Upload   types.String `tfsdk:"upload"`   // "enabled", "disabled" (v2 only)

	// v1 fields (bool-based, renamed from v4 disable_* fields)
	DCP types.Bool `tfsdk:"dcp"` // disable_copy_paste → dcp (v1 only)
	DD  types.Bool `tfsdk:"dd"`  // disable_download → dd (v1 only)
	DK  types.Bool `tfsdk:"dk"`  // disable_keyboard → dk (v1 only)
	DP  types.Bool `tfsdk:"dp"`  // disable_printing → dp (v1 only)
	DU  types.Bool `tfsdk:"du"`  // disable_upload → du (v1 only)

	Version types.String `tfsdk:"version"` // "v1" or "v2"
}

// TargetBlockPageModel represents block_page settings (new in v5).
type TargetBlockPageModel struct {
	TargetURI      types.String `tfsdk:"target_uri"`
	IncludeContext types.Bool   `tfsdk:"include_context"`
}

// TargetCheckSessionModel represents check_session in v5.
type TargetCheckSessionModel struct {
	Duration types.String `tfsdk:"duration"`
	Enforce  types.Bool   `tfsdk:"enforce"`
}

// TargetDNSResolversModel represents dns_resolvers in v5.
type TargetDNSResolversModel struct {
	IPV4 *[]*TargetDNSResolversIPV4Model `tfsdk:"ipv4"`
	IPV6 *[]*TargetDNSResolversIPV6Model `tfsdk:"ipv6"`
}

// TargetDNSResolversIPV4Model represents IPv4 resolver.
type TargetDNSResolversIPV4Model struct {
	IP                         types.String `tfsdk:"ip"`
	Port                       types.Int64  `tfsdk:"port"`
	RouteThroughPrivateNetwork types.Bool   `tfsdk:"route_through_private_network"`
	VnetID                     types.String `tfsdk:"vnet_id"`
}

// TargetDNSResolversIPV6Model represents IPv6 resolver.
type TargetDNSResolversIPV6Model struct {
	IP                         types.String `tfsdk:"ip"`
	Port                       types.Int64  `tfsdk:"port"`
	RouteThroughPrivateNetwork types.Bool   `tfsdk:"route_through_private_network"`
	VnetID                     types.String `tfsdk:"vnet_id"`
}

// TargetEgressModel represents egress in v5.
type TargetEgressModel struct {
	IPV4         types.String `tfsdk:"ipv4"`
	IPV4Fallback types.String `tfsdk:"ipv4_fallback"`
	IPV6         types.String `tfsdk:"ipv6"`
}

// TargetForensicCopyModel represents forensic_copy (new in v5).
type TargetForensicCopyModel struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

// TargetL4overrideModel represents l4override in v5.
type TargetL4overrideModel struct {
	IP   types.String `tfsdk:"ip"`
	Port types.Int64  `tfsdk:"port"`
}

// TargetNotificationSettingsModel represents notification_settings in v5.
type TargetNotificationSettingsModel struct {
	Enabled        types.Bool   `tfsdk:"enabled"`
	IncludeContext types.Bool   `tfsdk:"include_context"` // New in v5
	Msg            types.String `tfsdk:"msg"`             // Renamed from message
	SupportURL     types.String `tfsdk:"support_url"`
}

// TargetPayloadLogModel represents payload_log in v5.
type TargetPayloadLogModel struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

// TargetQuarantineModel represents quarantine (new in v5).
type TargetQuarantineModel struct {
	FileTypes *[]types.String `tfsdk:"file_types"`
}

// TargetRedirectModel represents redirect (new in v5).
type TargetRedirectModel struct {
	TargetURI            types.String `tfsdk:"target_uri"`
	IncludeContext       types.Bool   `tfsdk:"include_context"`
	PreservePathAndQuery types.Bool   `tfsdk:"preserve_path_and_query"`
}

// TargetResolveDNSInternallyModel represents resolve_dns_internally in v5.
type TargetResolveDNSInternallyModel struct {
	Fallback types.String `tfsdk:"fallback"`
	ViewID   types.String `tfsdk:"view_id"`
}

// TargetUntrustedCERTModel represents untrusted_cert in v5.
type TargetUntrustedCERTModel struct {
	Action types.String `tfsdk:"action"`
}

// TargetScheduleModel represents schedule (new in v5).
type TargetScheduleModel struct {
	Fri      types.String `tfsdk:"fri"`
	Mon      types.String `tfsdk:"mon"`
	Sat      types.String `tfsdk:"sat"`
	Sun      types.String `tfsdk:"sun"`
	Thu      types.String `tfsdk:"thu"`
	TimeZone types.String `tfsdk:"time_zone"`
	Tue      types.String `tfsdk:"tue"`
	Wed      types.String `tfsdk:"wed"`
}
