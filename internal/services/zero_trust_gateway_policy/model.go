// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_policy

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustGatewayPolicyResultEnvelope struct {
	Result ZeroTrustGatewayPolicyModel `json:"result,computed"`
}

type ZeroTrustGatewayPolicyModel struct {
	ID            types.String                             `tfsdk:"id" json:"id,computed"`
	AccountID     types.String                             `tfsdk:"account_id" path:"account_id"`
	Action        types.String                             `tfsdk:"action" json:"action"`
	Name          types.String                             `tfsdk:"name" json:"name"`
	Description   types.String                             `tfsdk:"description" json:"description"`
	DevicePosture types.String                             `tfsdk:"device_posture" json:"device_posture"`
	Enabled       types.Bool                               `tfsdk:"enabled" json:"enabled"`
	Identity      types.String                             `tfsdk:"identity" json:"identity"`
	Precedence    types.Int64                              `tfsdk:"precedence" json:"precedence"`
	Traffic       types.String                             `tfsdk:"traffic" json:"traffic"`
	Filters       *[]types.String                          `tfsdk:"filters" json:"filters"`
	RuleSettings  *ZeroTrustGatewayPolicyRuleSettingsModel `tfsdk:"rule_settings" json:"rule_settings"`
	Schedule      *ZeroTrustGatewayPolicyScheduleModel     `tfsdk:"schedule" json:"schedule"`
	CreatedAt     timetypes.RFC3339                        `tfsdk:"created_at" json:"created_at,computed"`
	DeletedAt     timetypes.RFC3339                        `tfsdk:"deleted_at" json:"deleted_at,computed"`
	UpdatedAt     timetypes.RFC3339                        `tfsdk:"updated_at" json:"updated_at,computed"`
}

type ZeroTrustGatewayPolicyRuleSettingsModel struct {
	AddHeaders                      map[string]types.String                                      `tfsdk:"add_headers" json:"add_headers"`
	AllowChildBypass                types.Bool                                                   `tfsdk:"allow_child_bypass" json:"allow_child_bypass"`
	AuditSSH                        *ZeroTrustGatewayPolicyRuleSettingsAuditSSHModel             `tfsdk:"audit_ssh" json:"audit_ssh"`
	BISOAdminControls               *ZeroTrustGatewayPolicyRuleSettingsBISOAdminControlsModel    `tfsdk:"biso_admin_controls" json:"biso_admin_controls"`
	BlockPageEnabled                types.Bool                                                   `tfsdk:"block_page_enabled" json:"block_page_enabled"`
	BlockReason                     types.String                                                 `tfsdk:"block_reason" json:"block_reason"`
	BypassParentRule                types.Bool                                                   `tfsdk:"bypass_parent_rule" json:"bypass_parent_rule"`
	CheckSession                    *ZeroTrustGatewayPolicyRuleSettingsCheckSessionModel         `tfsdk:"check_session" json:"check_session"`
	DNSResolvers                    *ZeroTrustGatewayPolicyRuleSettingsDNSResolversModel         `tfsdk:"dns_resolvers" json:"dns_resolvers"`
	Egress                          *ZeroTrustGatewayPolicyRuleSettingsEgressModel               `tfsdk:"egress" json:"egress"`
	IgnoreCNAMECategoryMatches      types.Bool                                                   `tfsdk:"ignore_cname_category_matches" json:"ignore_cname_category_matches"`
	InsecureDisableDNSSECValidation types.Bool                                                   `tfsdk:"insecure_disable_dnssec_validation" json:"insecure_disable_dnssec_validation"`
	IPCategories                    types.Bool                                                   `tfsdk:"ip_categories" json:"ip_categories"`
	IPIndicatorFeeds                types.Bool                                                   `tfsdk:"ip_indicator_feeds" json:"ip_indicator_feeds"`
	L4override                      *ZeroTrustGatewayPolicyRuleSettingsL4overrideModel           `tfsdk:"l4override" json:"l4override"`
	NotificationSettings            *ZeroTrustGatewayPolicyRuleSettingsNotificationSettingsModel `tfsdk:"notification_settings" json:"notification_settings"`
	OverrideHost                    types.String                                                 `tfsdk:"override_host" json:"override_host"`
	OverrideIPs                     *[]types.String                                              `tfsdk:"override_ips" json:"override_ips"`
	PayloadLog                      *ZeroTrustGatewayPolicyRuleSettingsPayloadLogModel           `tfsdk:"payload_log" json:"payload_log"`
	ResolveDNSThroughCloudflare     types.Bool                                                   `tfsdk:"resolve_dns_through_cloudflare" json:"resolve_dns_through_cloudflare"`
	UntrustedCERT                   *ZeroTrustGatewayPolicyRuleSettingsUntrustedCERTModel        `tfsdk:"untrusted_cert" json:"untrusted_cert"`
}

type ZeroTrustGatewayPolicyRuleSettingsAuditSSHModel struct {
	CommandLogging types.Bool `tfsdk:"command_logging" json:"command_logging"`
}

type ZeroTrustGatewayPolicyRuleSettingsBISOAdminControlsModel struct {
	DCP types.Bool `tfsdk:"dcp" json:"dcp"`
	DD  types.Bool `tfsdk:"dd" json:"dd"`
	DK  types.Bool `tfsdk:"dk" json:"dk"`
	DP  types.Bool `tfsdk:"dp" json:"dp"`
	DU  types.Bool `tfsdk:"du" json:"du"`
}

type ZeroTrustGatewayPolicyRuleSettingsCheckSessionModel struct {
	Duration types.String `tfsdk:"duration" json:"duration"`
	Enforce  types.Bool   `tfsdk:"enforce" json:"enforce"`
}

type ZeroTrustGatewayPolicyRuleSettingsDNSResolversModel struct {
	IPV4 *[]*ZeroTrustGatewayPolicyRuleSettingsDNSResolversIPV4Model `tfsdk:"ipv4" json:"ipv4"`
	IPV6 *[]*ZeroTrustGatewayPolicyRuleSettingsDNSResolversIPV6Model `tfsdk:"ipv6" json:"ipv6"`
}

type ZeroTrustGatewayPolicyRuleSettingsDNSResolversIPV4Model struct {
	IP                         types.String `tfsdk:"ip" json:"ip"`
	Port                       types.Int64  `tfsdk:"port" json:"port"`
	RouteThroughPrivateNetwork types.Bool   `tfsdk:"route_through_private_network" json:"route_through_private_network"`
	VnetID                     types.String `tfsdk:"vnet_id" json:"vnet_id"`
}

type ZeroTrustGatewayPolicyRuleSettingsDNSResolversIPV6Model struct {
	IP                         types.String `tfsdk:"ip" json:"ip"`
	Port                       types.Int64  `tfsdk:"port" json:"port"`
	RouteThroughPrivateNetwork types.Bool   `tfsdk:"route_through_private_network" json:"route_through_private_network"`
	VnetID                     types.String `tfsdk:"vnet_id" json:"vnet_id"`
}

type ZeroTrustGatewayPolicyRuleSettingsEgressModel struct {
	IPV4         types.String `tfsdk:"ipv4" json:"ipv4"`
	IPV4Fallback types.String `tfsdk:"ipv4_fallback" json:"ipv4_fallback"`
	IPV6         types.String `tfsdk:"ipv6" json:"ipv6"`
}

type ZeroTrustGatewayPolicyRuleSettingsL4overrideModel struct {
	IP   types.String `tfsdk:"ip" json:"ip"`
	Port types.Int64  `tfsdk:"port" json:"port"`
}

type ZeroTrustGatewayPolicyRuleSettingsNotificationSettingsModel struct {
	Enabled    types.Bool   `tfsdk:"enabled" json:"enabled"`
	Msg        types.String `tfsdk:"msg" json:"msg"`
	SupportURL types.String `tfsdk:"support_url" json:"support_url"`
}

type ZeroTrustGatewayPolicyRuleSettingsPayloadLogModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled"`
}

type ZeroTrustGatewayPolicyRuleSettingsUntrustedCERTModel struct {
	Action types.String `tfsdk:"action" json:"action"`
}

type ZeroTrustGatewayPolicyScheduleModel struct {
	Fri      types.String `tfsdk:"fri" json:"fri"`
	Mon      types.String `tfsdk:"mon" json:"mon"`
	Sat      types.String `tfsdk:"sat" json:"sat"`
	Sun      types.String `tfsdk:"sun" json:"sun"`
	Thu      types.String `tfsdk:"thu" json:"thu"`
	TimeZone types.String `tfsdk:"time_zone" json:"time_zone"`
	Tue      types.String `tfsdk:"tue" json:"tue"`
	Wed      types.String `tfsdk:"wed" json:"wed"`
}
