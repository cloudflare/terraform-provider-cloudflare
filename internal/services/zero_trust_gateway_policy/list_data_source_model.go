// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_policy

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustGatewayPoliciesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustGatewayPoliciesResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustGatewayPoliciesDataSourceModel struct {
	AccountID types.String                                                                `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                                                                 `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustGatewayPoliciesResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustGatewayPoliciesDataSourceModel) toListParams() (params zero_trust.GatewayRuleListParams, diags diag.Diagnostics) {
	params = zero_trust.GatewayRuleListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustGatewayPoliciesResultDataSourceModel struct {
	ID            types.String                                         `tfsdk:"id" json:"id,computed_optional"`
	Action        types.String                                         `tfsdk:"action" json:"action,computed_optional"`
	CreatedAt     timetypes.RFC3339                                    `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DeletedAt     timetypes.RFC3339                                    `tfsdk:"deleted_at" json:"deleted_at,computed" format:"date-time"`
	Description   types.String                                         `tfsdk:"description" json:"description,computed_optional"`
	DevicePosture types.String                                         `tfsdk:"device_posture" json:"device_posture,computed_optional"`
	Enabled       types.Bool                                           `tfsdk:"enabled" json:"enabled,computed_optional"`
	Filters       *[]types.String                                      `tfsdk:"filters" json:"filters,computed_optional"`
	Identity      types.String                                         `tfsdk:"identity" json:"identity,computed_optional"`
	Name          types.String                                         `tfsdk:"name" json:"name,computed_optional"`
	Precedence    types.Int64                                          `tfsdk:"precedence" json:"precedence,computed_optional"`
	RuleSettings  *ZeroTrustGatewayPoliciesRuleSettingsDataSourceModel `tfsdk:"rule_settings" json:"rule_settings,computed_optional"`
	Schedule      *ZeroTrustGatewayPoliciesScheduleDataSourceModel     `tfsdk:"schedule" json:"schedule,computed_optional"`
	Traffic       types.String                                         `tfsdk:"traffic" json:"traffic,computed_optional"`
	UpdatedAt     timetypes.RFC3339                                    `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

type ZeroTrustGatewayPoliciesRuleSettingsDataSourceModel struct {
	AddHeaders                      map[string]types.String                                                  `tfsdk:"add_headers" json:"add_headers,computed_optional"`
	AllowChildBypass                types.Bool                                                               `tfsdk:"allow_child_bypass" json:"allow_child_bypass,computed_optional"`
	AuditSSH                        *ZeroTrustGatewayPoliciesRuleSettingsAuditSSHDataSourceModel             `tfsdk:"audit_ssh" json:"audit_ssh,computed_optional"`
	BISOAdminControls               *ZeroTrustGatewayPoliciesRuleSettingsBISOAdminControlsDataSourceModel    `tfsdk:"biso_admin_controls" json:"biso_admin_controls,computed_optional"`
	BlockPageEnabled                types.Bool                                                               `tfsdk:"block_page_enabled" json:"block_page_enabled,computed_optional"`
	BlockReason                     types.String                                                             `tfsdk:"block_reason" json:"block_reason,computed_optional"`
	BypassParentRule                types.Bool                                                               `tfsdk:"bypass_parent_rule" json:"bypass_parent_rule,computed_optional"`
	CheckSession                    *ZeroTrustGatewayPoliciesRuleSettingsCheckSessionDataSourceModel         `tfsdk:"check_session" json:"check_session,computed_optional"`
	DNSResolvers                    *ZeroTrustGatewayPoliciesRuleSettingsDNSResolversDataSourceModel         `tfsdk:"dns_resolvers" json:"dns_resolvers,computed_optional"`
	Egress                          *ZeroTrustGatewayPoliciesRuleSettingsEgressDataSourceModel               `tfsdk:"egress" json:"egress,computed_optional"`
	IgnoreCNAMECategoryMatches      types.Bool                                                               `tfsdk:"ignore_cname_category_matches" json:"ignore_cname_category_matches,computed_optional"`
	InsecureDisableDNSSECValidation types.Bool                                                               `tfsdk:"insecure_disable_dnssec_validation" json:"insecure_disable_dnssec_validation,computed_optional"`
	IPCategories                    types.Bool                                                               `tfsdk:"ip_categories" json:"ip_categories,computed_optional"`
	IPIndicatorFeeds                types.Bool                                                               `tfsdk:"ip_indicator_feeds" json:"ip_indicator_feeds,computed_optional"`
	L4override                      *ZeroTrustGatewayPoliciesRuleSettingsL4overrideDataSourceModel           `tfsdk:"l4override" json:"l4override,computed_optional"`
	NotificationSettings            *ZeroTrustGatewayPoliciesRuleSettingsNotificationSettingsDataSourceModel `tfsdk:"notification_settings" json:"notification_settings,computed_optional"`
	OverrideHost                    types.String                                                             `tfsdk:"override_host" json:"override_host,computed_optional"`
	OverrideIPs                     *[]types.String                                                          `tfsdk:"override_ips" json:"override_ips,computed_optional"`
	PayloadLog                      *ZeroTrustGatewayPoliciesRuleSettingsPayloadLogDataSourceModel           `tfsdk:"payload_log" json:"payload_log,computed_optional"`
	ResolveDNSThroughCloudflare     types.Bool                                                               `tfsdk:"resolve_dns_through_cloudflare" json:"resolve_dns_through_cloudflare,computed_optional"`
	UntrustedCERT                   *ZeroTrustGatewayPoliciesRuleSettingsUntrustedCERTDataSourceModel        `tfsdk:"untrusted_cert" json:"untrusted_cert,computed_optional"`
}

type ZeroTrustGatewayPoliciesRuleSettingsAuditSSHDataSourceModel struct {
	CommandLogging types.Bool `tfsdk:"command_logging" json:"command_logging,computed_optional"`
}

type ZeroTrustGatewayPoliciesRuleSettingsBISOAdminControlsDataSourceModel struct {
	DCP types.Bool `tfsdk:"dcp" json:"dcp,computed_optional"`
	DD  types.Bool `tfsdk:"dd" json:"dd,computed_optional"`
	DK  types.Bool `tfsdk:"dk" json:"dk,computed_optional"`
	DP  types.Bool `tfsdk:"dp" json:"dp,computed_optional"`
	DU  types.Bool `tfsdk:"du" json:"du,computed_optional"`
}

type ZeroTrustGatewayPoliciesRuleSettingsCheckSessionDataSourceModel struct {
	Duration types.String `tfsdk:"duration" json:"duration,computed_optional"`
	Enforce  types.Bool   `tfsdk:"enforce" json:"enforce,computed_optional"`
}

type ZeroTrustGatewayPoliciesRuleSettingsDNSResolversDataSourceModel struct {
	IPV4 *[]*ZeroTrustGatewayPoliciesRuleSettingsDNSResolversIPV4DataSourceModel `tfsdk:"ipv4" json:"ipv4,computed_optional"`
	IPV6 *[]*ZeroTrustGatewayPoliciesRuleSettingsDNSResolversIPV6DataSourceModel `tfsdk:"ipv6" json:"ipv6,computed_optional"`
}

type ZeroTrustGatewayPoliciesRuleSettingsDNSResolversIPV4DataSourceModel struct {
	IP                         types.String `tfsdk:"ip" json:"ip,computed"`
	Port                       types.Int64  `tfsdk:"port" json:"port,computed_optional"`
	RouteThroughPrivateNetwork types.Bool   `tfsdk:"route_through_private_network" json:"route_through_private_network,computed_optional"`
	VnetID                     types.String `tfsdk:"vnet_id" json:"vnet_id,computed_optional"`
}

type ZeroTrustGatewayPoliciesRuleSettingsDNSResolversIPV6DataSourceModel struct {
	IP                         types.String `tfsdk:"ip" json:"ip,computed"`
	Port                       types.Int64  `tfsdk:"port" json:"port,computed_optional"`
	RouteThroughPrivateNetwork types.Bool   `tfsdk:"route_through_private_network" json:"route_through_private_network,computed_optional"`
	VnetID                     types.String `tfsdk:"vnet_id" json:"vnet_id,computed_optional"`
}

type ZeroTrustGatewayPoliciesRuleSettingsEgressDataSourceModel struct {
	IPV4         types.String `tfsdk:"ipv4" json:"ipv4,computed_optional"`
	IPV4Fallback types.String `tfsdk:"ipv4_fallback" json:"ipv4_fallback,computed_optional"`
	IPV6         types.String `tfsdk:"ipv6" json:"ipv6,computed_optional"`
}

type ZeroTrustGatewayPoliciesRuleSettingsL4overrideDataSourceModel struct {
	IP   types.String `tfsdk:"ip" json:"ip,computed_optional"`
	Port types.Int64  `tfsdk:"port" json:"port,computed_optional"`
}

type ZeroTrustGatewayPoliciesRuleSettingsNotificationSettingsDataSourceModel struct {
	Enabled    types.Bool   `tfsdk:"enabled" json:"enabled,computed_optional"`
	Msg        types.String `tfsdk:"msg" json:"msg,computed_optional"`
	SupportURL types.String `tfsdk:"support_url" json:"support_url,computed_optional"`
}

type ZeroTrustGatewayPoliciesRuleSettingsPayloadLogDataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed_optional"`
}

type ZeroTrustGatewayPoliciesRuleSettingsUntrustedCERTDataSourceModel struct {
	Action types.String `tfsdk:"action" json:"action,computed_optional"`
}

type ZeroTrustGatewayPoliciesScheduleDataSourceModel struct {
	Fri      types.String `tfsdk:"fri" json:"fri,computed_optional"`
	Mon      types.String `tfsdk:"mon" json:"mon,computed_optional"`
	Sat      types.String `tfsdk:"sat" json:"sat,computed_optional"`
	Sun      types.String `tfsdk:"sun" json:"sun,computed_optional"`
	Thu      types.String `tfsdk:"thu" json:"thu,computed_optional"`
	TimeZone types.String `tfsdk:"time_zone" json:"time_zone,computed_optional"`
	Tue      types.String `tfsdk:"tue" json:"tue,computed_optional"`
	Wed      types.String `tfsdk:"wed" json:"wed,computed_optional"`
}