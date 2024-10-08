// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_policy

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustGatewayPolicyResultEnvelope struct {
	Result ZeroTrustGatewayPolicyModel `json:"result"`
}

type ZeroTrustGatewayPolicyModel struct {
	ID            types.String                                                      `tfsdk:"id" json:"id,computed"`
	AccountID     types.String                                                      `tfsdk:"account_id" path:"account_id,required"`
	Action        types.String                                                      `tfsdk:"action" json:"action,required"`
	Name          types.String                                                      `tfsdk:"name" json:"name,required"`
	Description   types.String                                                      `tfsdk:"description" json:"description,optional"`
	DevicePosture types.String                                                      `tfsdk:"device_posture" json:"device_posture,optional"`
	Enabled       types.Bool                                                        `tfsdk:"enabled" json:"enabled,optional"`
	Identity      types.String                                                      `tfsdk:"identity" json:"identity,optional"`
	Precedence    types.Int64                                                       `tfsdk:"precedence" json:"precedence,optional"`
	Traffic       types.String                                                      `tfsdk:"traffic" json:"traffic,optional"`
	Filters       *[]types.String                                                   `tfsdk:"filters" json:"filters,optional"`
	RuleSettings  customfield.NestedObject[ZeroTrustGatewayPolicyRuleSettingsModel] `tfsdk:"rule_settings" json:"rule_settings,computed_optional"`
	Schedule      customfield.NestedObject[ZeroTrustGatewayPolicyScheduleModel]     `tfsdk:"schedule" json:"schedule,computed_optional"`
	CreatedAt     timetypes.RFC3339                                                 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DeletedAt     timetypes.RFC3339                                                 `tfsdk:"deleted_at" json:"deleted_at,computed" format:"date-time"`
	UpdatedAt     timetypes.RFC3339                                                 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m ZeroTrustGatewayPolicyModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustGatewayPolicyModel) MarshalJSONForUpdate(state ZeroTrustGatewayPolicyModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustGatewayPolicyRuleSettingsModel struct {
	AddHeaders                      map[string]types.String                                                               `tfsdk:"add_headers" json:"add_headers,optional"`
	AllowChildBypass                types.Bool                                                                            `tfsdk:"allow_child_bypass" json:"allow_child_bypass,optional"`
	AuditSSH                        customfield.NestedObject[ZeroTrustGatewayPolicyRuleSettingsAuditSSHModel]             `tfsdk:"audit_ssh" json:"audit_ssh,computed_optional"`
	BISOAdminControls               customfield.NestedObject[ZeroTrustGatewayPolicyRuleSettingsBISOAdminControlsModel]    `tfsdk:"biso_admin_controls" json:"biso_admin_controls,computed_optional"`
	BlockPageEnabled                types.Bool                                                                            `tfsdk:"block_page_enabled" json:"block_page_enabled,optional"`
	BlockReason                     types.String                                                                          `tfsdk:"block_reason" json:"block_reason,optional"`
	BypassParentRule                types.Bool                                                                            `tfsdk:"bypass_parent_rule" json:"bypass_parent_rule,optional"`
	CheckSession                    customfield.NestedObject[ZeroTrustGatewayPolicyRuleSettingsCheckSessionModel]         `tfsdk:"check_session" json:"check_session,computed_optional"`
	DNSResolvers                    customfield.NestedObject[ZeroTrustGatewayPolicyRuleSettingsDNSResolversModel]         `tfsdk:"dns_resolvers" json:"dns_resolvers,computed_optional"`
	Egress                          customfield.NestedObject[ZeroTrustGatewayPolicyRuleSettingsEgressModel]               `tfsdk:"egress" json:"egress,computed_optional"`
	IgnoreCNAMECategoryMatches      types.Bool                                                                            `tfsdk:"ignore_cname_category_matches" json:"ignore_cname_category_matches,optional"`
	InsecureDisableDNSSECValidation types.Bool                                                                            `tfsdk:"insecure_disable_dnssec_validation" json:"insecure_disable_dnssec_validation,optional"`
	IPCategories                    types.Bool                                                                            `tfsdk:"ip_categories" json:"ip_categories,optional"`
	IPIndicatorFeeds                types.Bool                                                                            `tfsdk:"ip_indicator_feeds" json:"ip_indicator_feeds,optional"`
	L4override                      customfield.NestedObject[ZeroTrustGatewayPolicyRuleSettingsL4overrideModel]           `tfsdk:"l4override" json:"l4override,computed_optional"`
	NotificationSettings            customfield.NestedObject[ZeroTrustGatewayPolicyRuleSettingsNotificationSettingsModel] `tfsdk:"notification_settings" json:"notification_settings,computed_optional"`
	OverrideHost                    types.String                                                                          `tfsdk:"override_host" json:"override_host,optional"`
	OverrideIPs                     *[]types.String                                                                       `tfsdk:"override_ips" json:"override_ips,optional"`
	PayloadLog                      customfield.NestedObject[ZeroTrustGatewayPolicyRuleSettingsPayloadLogModel]           `tfsdk:"payload_log" json:"payload_log,computed_optional"`
	Quarantine                      customfield.NestedObject[ZeroTrustGatewayPolicyRuleSettingsQuarantineModel]           `tfsdk:"quarantine" json:"quarantine,computed_optional"`
	ResolveDNSThroughCloudflare     types.Bool                                                                            `tfsdk:"resolve_dns_through_cloudflare" json:"resolve_dns_through_cloudflare,optional"`
	UntrustedCERT                   customfield.NestedObject[ZeroTrustGatewayPolicyRuleSettingsUntrustedCERTModel]        `tfsdk:"untrusted_cert" json:"untrusted_cert,computed_optional"`
}

type ZeroTrustGatewayPolicyRuleSettingsAuditSSHModel struct {
	CommandLogging types.Bool `tfsdk:"command_logging" json:"command_logging,optional"`
}

type ZeroTrustGatewayPolicyRuleSettingsBISOAdminControlsModel struct {
	DCP types.Bool `tfsdk:"dcp" json:"dcp,optional"`
	DD  types.Bool `tfsdk:"dd" json:"dd,optional"`
	DK  types.Bool `tfsdk:"dk" json:"dk,optional"`
	DP  types.Bool `tfsdk:"dp" json:"dp,optional"`
	DU  types.Bool `tfsdk:"du" json:"du,optional"`
}

type ZeroTrustGatewayPolicyRuleSettingsCheckSessionModel struct {
	Duration types.String `tfsdk:"duration" json:"duration,optional"`
	Enforce  types.Bool   `tfsdk:"enforce" json:"enforce,optional"`
}

type ZeroTrustGatewayPolicyRuleSettingsDNSResolversModel struct {
	IPV4 customfield.NestedObjectList[ZeroTrustGatewayPolicyRuleSettingsDNSResolversIPV4Model] `tfsdk:"ipv4" json:"ipv4,computed_optional"`
	IPV6 customfield.NestedObjectList[ZeroTrustGatewayPolicyRuleSettingsDNSResolversIPV6Model] `tfsdk:"ipv6" json:"ipv6,computed_optional"`
}

type ZeroTrustGatewayPolicyRuleSettingsDNSResolversIPV4Model struct {
	IP                         types.String `tfsdk:"ip" json:"ip,required"`
	Port                       types.Int64  `tfsdk:"port" json:"port,optional"`
	RouteThroughPrivateNetwork types.Bool   `tfsdk:"route_through_private_network" json:"route_through_private_network,optional"`
	VnetID                     types.String `tfsdk:"vnet_id" json:"vnet_id,optional"`
}

type ZeroTrustGatewayPolicyRuleSettingsDNSResolversIPV6Model struct {
	IP                         types.String `tfsdk:"ip" json:"ip,required"`
	Port                       types.Int64  `tfsdk:"port" json:"port,optional"`
	RouteThroughPrivateNetwork types.Bool   `tfsdk:"route_through_private_network" json:"route_through_private_network,optional"`
	VnetID                     types.String `tfsdk:"vnet_id" json:"vnet_id,optional"`
}

type ZeroTrustGatewayPolicyRuleSettingsEgressModel struct {
	IPV4         types.String `tfsdk:"ipv4" json:"ipv4,optional"`
	IPV4Fallback types.String `tfsdk:"ipv4_fallback" json:"ipv4_fallback,optional"`
	IPV6         types.String `tfsdk:"ipv6" json:"ipv6,optional"`
}

type ZeroTrustGatewayPolicyRuleSettingsL4overrideModel struct {
	IP   types.String `tfsdk:"ip" json:"ip,optional"`
	Port types.Int64  `tfsdk:"port" json:"port,optional"`
}

type ZeroTrustGatewayPolicyRuleSettingsNotificationSettingsModel struct {
	Enabled    types.Bool   `tfsdk:"enabled" json:"enabled,optional"`
	Msg        types.String `tfsdk:"msg" json:"msg,optional"`
	SupportURL types.String `tfsdk:"support_url" json:"support_url,optional"`
}

type ZeroTrustGatewayPolicyRuleSettingsPayloadLogModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,optional"`
}

type ZeroTrustGatewayPolicyRuleSettingsQuarantineModel struct {
	FileTypes *[]types.String `tfsdk:"file_types" json:"file_types,optional"`
}

type ZeroTrustGatewayPolicyRuleSettingsUntrustedCERTModel struct {
	Action types.String `tfsdk:"action" json:"action,optional"`
}

type ZeroTrustGatewayPolicyScheduleModel struct {
	Fri      types.String `tfsdk:"fri" json:"fri,optional"`
	Mon      types.String `tfsdk:"mon" json:"mon,optional"`
	Sat      types.String `tfsdk:"sat" json:"sat,optional"`
	Sun      types.String `tfsdk:"sun" json:"sun,optional"`
	Thu      types.String `tfsdk:"thu" json:"thu,optional"`
	TimeZone types.String `tfsdk:"time_zone" json:"time_zone,optional"`
	Tue      types.String `tfsdk:"tue" json:"tue,optional"`
	Wed      types.String `tfsdk:"wed" json:"wed,optional"`
}
