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
	Precedence    types.Int64                                                       `tfsdk:"precedence" json:"precedence,optional"`
	Filters       *[]types.String                                                   `tfsdk:"filters" json:"filters,optional"`
	Schedule      *ZeroTrustGatewayPolicyScheduleModel                              `tfsdk:"schedule" json:"schedule,optional"`
	DevicePosture types.String                                                      `tfsdk:"device_posture" json:"device_posture,computed_optional"`
	Enabled       types.Bool                                                        `tfsdk:"enabled" json:"enabled,computed_optional"`
	Identity      types.String                                                      `tfsdk:"identity" json:"identity,computed_optional"`
	Traffic       types.String                                                      `tfsdk:"traffic" json:"traffic,computed_optional"`
	Expiration    customfield.NestedObject[ZeroTrustGatewayPolicyExpirationModel]   `tfsdk:"expiration" json:"expiration,computed_optional"`
	RuleSettings  customfield.NestedObject[ZeroTrustGatewayPolicyRuleSettingsModel] `tfsdk:"rule_settings" json:"rule_settings,computed_optional"`
	CreatedAt     timetypes.RFC3339                                                 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DeletedAt     timetypes.RFC3339                                                 `tfsdk:"deleted_at" json:"deleted_at,computed" format:"date-time"`
	UpdatedAt     timetypes.RFC3339                                                 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Version       types.Int64                                                       `tfsdk:"version" json:"version,computed"`
}

func (m ZeroTrustGatewayPolicyModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustGatewayPolicyModel) MarshalJSONForUpdate(state ZeroTrustGatewayPolicyModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
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

type ZeroTrustGatewayPolicyExpirationModel struct {
	ExpiresAt timetypes.RFC3339 `tfsdk:"expires_at" json:"expires_at,required" format:"date-time"`
	Duration  types.Int64       `tfsdk:"duration" json:"duration,optional"`
	Expired   types.Bool        `tfsdk:"expired" json:"expired,computed"`
}

type ZeroTrustGatewayPolicyRuleSettingsModel struct {
	AddHeaders                      *map[string]types.String                                                           `tfsdk:"add_headers" json:"add_headers,optional"`
	AllowChildBypass                types.Bool                                                                         `tfsdk:"allow_child_bypass" json:"allow_child_bypass,optional"`
	AuditSSH                        *ZeroTrustGatewayPolicyRuleSettingsAuditSSHModel                                   `tfsdk:"audit_ssh" json:"audit_ssh,optional"`
	BISOAdminControls               customfield.NestedObject[ZeroTrustGatewayPolicyRuleSettingsBISOAdminControlsModel] `tfsdk:"biso_admin_controls" json:"biso_admin_controls,computed_optional"`
	BlockPage                       *ZeroTrustGatewayPolicyRuleSettingsBlockPageModel                                  `tfsdk:"block_page" json:"block_page,optional"`
	BlockPageEnabled                types.Bool                                                                         `tfsdk:"block_page_enabled" json:"block_page_enabled,optional"`
	BlockReason                     types.String                                                                       `tfsdk:"block_reason" json:"block_reason,optional"`
	BypassParentRule                types.Bool                                                                         `tfsdk:"bypass_parent_rule" json:"bypass_parent_rule,optional"`
	CheckSession                    *ZeroTrustGatewayPolicyRuleSettingsCheckSessionModel                               `tfsdk:"check_session" json:"check_session,optional"`
	DNSResolvers                    *ZeroTrustGatewayPolicyRuleSettingsDNSResolversModel                               `tfsdk:"dns_resolvers" json:"dns_resolvers,optional"`
	Egress                          *ZeroTrustGatewayPolicyRuleSettingsEgressModel                                     `tfsdk:"egress" json:"egress,optional"`
	IgnoreCNAMECategoryMatches      types.Bool                                                                         `tfsdk:"ignore_cname_category_matches" json:"ignore_cname_category_matches,optional"`
	InsecureDisableDNSSECValidation types.Bool                                                                         `tfsdk:"insecure_disable_dnssec_validation" json:"insecure_disable_dnssec_validation,optional"`
	IPCategories                    types.Bool                                                                         `tfsdk:"ip_categories" json:"ip_categories,optional"`
	IPIndicatorFeeds                types.Bool                                                                         `tfsdk:"ip_indicator_feeds" json:"ip_indicator_feeds,optional"`
	L4override                      *ZeroTrustGatewayPolicyRuleSettingsL4overrideModel                                 `tfsdk:"l4override" json:"l4override,optional"`
	NotificationSettings            *ZeroTrustGatewayPolicyRuleSettingsNotificationSettingsModel                       `tfsdk:"notification_settings" json:"notification_settings,optional"`
	OverrideHost                    types.String                                                                       `tfsdk:"override_host" json:"override_host,optional"`
	OverrideIPs                     *[]types.String                                                                    `tfsdk:"override_ips" json:"override_ips,optional"`
	PayloadLog                      *ZeroTrustGatewayPolicyRuleSettingsPayloadLogModel                                 `tfsdk:"payload_log" json:"payload_log,optional"`
	Quarantine                      *ZeroTrustGatewayPolicyRuleSettingsQuarantineModel                                 `tfsdk:"quarantine" json:"quarantine,optional"`
	Redirect                        *ZeroTrustGatewayPolicyRuleSettingsRedirectModel                                   `tfsdk:"redirect" json:"redirect,optional"`
	ResolveDNSInternally            *ZeroTrustGatewayPolicyRuleSettingsResolveDNSInternallyModel                       `tfsdk:"resolve_dns_internally" json:"resolve_dns_internally,optional"`
	ResolveDNSThroughCloudflare     types.Bool                                                                         `tfsdk:"resolve_dns_through_cloudflare" json:"resolve_dns_through_cloudflare,optional"`
	UntrustedCERT                   *ZeroTrustGatewayPolicyRuleSettingsUntrustedCERTModel                              `tfsdk:"untrusted_cert" json:"untrusted_cert,optional"`
}

type ZeroTrustGatewayPolicyRuleSettingsAuditSSHModel struct {
	CommandLogging types.Bool `tfsdk:"command_logging" json:"command_logging,optional"`
}

type ZeroTrustGatewayPolicyRuleSettingsBISOAdminControlsModel struct {
	Copy     types.String `tfsdk:"copy" json:"copy,optional"`
	DCP      types.Bool   `tfsdk:"dcp" json:"dcp,computed_optional"`
	DD       types.Bool   `tfsdk:"dd" json:"dd,computed_optional"`
	DK       types.Bool   `tfsdk:"dk" json:"dk,computed_optional"`
	Download types.String `tfsdk:"download" json:"download,optional"`
	DP       types.Bool   `tfsdk:"dp" json:"dp,computed_optional"`
	DU       types.Bool   `tfsdk:"du" json:"du,computed_optional"`
	Keyboard types.String `tfsdk:"keyboard" json:"keyboard,optional"`
	Paste    types.String `tfsdk:"paste" json:"paste,optional"`
	Printing types.String `tfsdk:"printing" json:"printing,optional"`
	Upload   types.String `tfsdk:"upload" json:"upload,optional"`
	Version  types.String `tfsdk:"version" json:"version,computed_optional"`
}

type ZeroTrustGatewayPolicyRuleSettingsBlockPageModel struct {
	TargetURI      types.String `tfsdk:"target_uri" json:"target_uri,required"`
	IncludeContext types.Bool   `tfsdk:"include_context" json:"include_context,optional"`
}

type ZeroTrustGatewayPolicyRuleSettingsCheckSessionModel struct {
	Duration types.String `tfsdk:"duration" json:"duration,optional"`
	Enforce  types.Bool   `tfsdk:"enforce" json:"enforce,optional"`
}

type ZeroTrustGatewayPolicyRuleSettingsDNSResolversModel struct {
	IPV4 *[]*ZeroTrustGatewayPolicyRuleSettingsDNSResolversIPV4Model `tfsdk:"ipv4" json:"ipv4,optional"`
	IPV6 *[]*ZeroTrustGatewayPolicyRuleSettingsDNSResolversIPV6Model `tfsdk:"ipv6" json:"ipv6,optional"`
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
	Enabled        types.Bool   `tfsdk:"enabled" json:"enabled,optional"`
	IncludeContext types.Bool   `tfsdk:"include_context" json:"include_context,optional"`
	Msg            types.String `tfsdk:"msg" json:"msg,optional"`
	SupportURL     types.String `tfsdk:"support_url" json:"support_url,optional"`
}

type ZeroTrustGatewayPolicyRuleSettingsPayloadLogModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,optional"`
}

type ZeroTrustGatewayPolicyRuleSettingsQuarantineModel struct {
	FileTypes *[]types.String `tfsdk:"file_types" json:"file_types,optional"`
}

type ZeroTrustGatewayPolicyRuleSettingsRedirectModel struct {
	TargetURI            types.String `tfsdk:"target_uri" json:"target_uri,required"`
	IncludeContext       types.Bool   `tfsdk:"include_context" json:"include_context,optional"`
	PreservePathAndQuery types.Bool   `tfsdk:"preserve_path_and_query" json:"preserve_path_and_query,optional"`
}

type ZeroTrustGatewayPolicyRuleSettingsResolveDNSInternallyModel struct {
	Fallback types.String `tfsdk:"fallback" json:"fallback,optional"`
	ViewID   types.String `tfsdk:"view_id" json:"view_id,optional"`
}

type ZeroTrustGatewayPolicyRuleSettingsUntrustedCERTModel struct {
	Action types.String `tfsdk:"action" json:"action,optional"`
}
