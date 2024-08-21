// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_policy

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustGatewayPolicyResultDataSourceEnvelope struct {
	Result ZeroTrustGatewayPolicyDataSourceModel `json:"result,computed"`
}

type ZeroTrustGatewayPolicyResultListDataSourceEnvelope struct {
	Result *[]*ZeroTrustGatewayPolicyDataSourceModel `json:"result,computed"`
}

type ZeroTrustGatewayPolicyDataSourceModel struct {
	AccountID     types.String                                       `tfsdk:"account_id" path:"account_id"`
	RuleID        types.String                                       `tfsdk:"rule_id" path:"rule_id"`
	CreatedAt     timetypes.RFC3339                                  `tfsdk:"created_at" json:"created_at,computed"`
	DeletedAt     timetypes.RFC3339                                  `tfsdk:"deleted_at" json:"deleted_at,computed"`
	UpdatedAt     timetypes.RFC3339                                  `tfsdk:"updated_at" json:"updated_at,computed"`
	Action        types.String                                       `tfsdk:"action" json:"action"`
	Description   types.String                                       `tfsdk:"description" json:"description"`
	DevicePosture types.String                                       `tfsdk:"device_posture" json:"device_posture"`
	Enabled       types.Bool                                         `tfsdk:"enabled" json:"enabled"`
	ID            types.String                                       `tfsdk:"id" json:"id"`
	Identity      types.String                                       `tfsdk:"identity" json:"identity"`
	Name          types.String                                       `tfsdk:"name" json:"name"`
	Precedence    types.Int64                                        `tfsdk:"precedence" json:"precedence"`
	Traffic       types.String                                       `tfsdk:"traffic" json:"traffic"`
	Filters       *[]types.String                                    `tfsdk:"filters" json:"filters"`
	RuleSettings  *ZeroTrustGatewayPolicyRuleSettingsDataSourceModel `tfsdk:"rule_settings" json:"rule_settings"`
	Schedule      *ZeroTrustGatewayPolicyScheduleDataSourceModel     `tfsdk:"schedule" json:"schedule"`
	Filter        *ZeroTrustGatewayPolicyFindOneByDataSourceModel    `tfsdk:"filter"`
}

func (m *ZeroTrustGatewayPolicyDataSourceModel) toReadParams() (params zero_trust.GatewayRuleGetParams, diags diag.Diagnostics) {
	params = zero_trust.GatewayRuleGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ZeroTrustGatewayPolicyDataSourceModel) toListParams() (params zero_trust.GatewayRuleListParams, diags diag.Diagnostics) {
	params = zero_trust.GatewayRuleListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type ZeroTrustGatewayPolicyRuleSettingsDataSourceModel struct {
	AddHeaders                      map[string]types.String                                                `tfsdk:"add_headers" json:"add_headers"`
	AllowChildBypass                types.Bool                                                             `tfsdk:"allow_child_bypass" json:"allow_child_bypass"`
	AuditSSH                        *ZeroTrustGatewayPolicyRuleSettingsAuditSSHDataSourceModel             `tfsdk:"audit_ssh" json:"audit_ssh"`
	BISOAdminControls               *ZeroTrustGatewayPolicyRuleSettingsBISOAdminControlsDataSourceModel    `tfsdk:"biso_admin_controls" json:"biso_admin_controls"`
	BlockPageEnabled                types.Bool                                                             `tfsdk:"block_page_enabled" json:"block_page_enabled"`
	BlockReason                     types.String                                                           `tfsdk:"block_reason" json:"block_reason"`
	BypassParentRule                types.Bool                                                             `tfsdk:"bypass_parent_rule" json:"bypass_parent_rule"`
	CheckSession                    *ZeroTrustGatewayPolicyRuleSettingsCheckSessionDataSourceModel         `tfsdk:"check_session" json:"check_session"`
	DNSResolvers                    *ZeroTrustGatewayPolicyRuleSettingsDNSResolversDataSourceModel         `tfsdk:"dns_resolvers" json:"dns_resolvers"`
	Egress                          *ZeroTrustGatewayPolicyRuleSettingsEgressDataSourceModel               `tfsdk:"egress" json:"egress"`
	IgnoreCNAMECategoryMatches      types.Bool                                                             `tfsdk:"ignore_cname_category_matches" json:"ignore_cname_category_matches"`
	InsecureDisableDNSSECValidation types.Bool                                                             `tfsdk:"insecure_disable_dnssec_validation" json:"insecure_disable_dnssec_validation"`
	IPCategories                    types.Bool                                                             `tfsdk:"ip_categories" json:"ip_categories"`
	IPIndicatorFeeds                types.Bool                                                             `tfsdk:"ip_indicator_feeds" json:"ip_indicator_feeds"`
	L4override                      *ZeroTrustGatewayPolicyRuleSettingsL4overrideDataSourceModel           `tfsdk:"l4override" json:"l4override"`
	NotificationSettings            *ZeroTrustGatewayPolicyRuleSettingsNotificationSettingsDataSourceModel `tfsdk:"notification_settings" json:"notification_settings"`
	OverrideHost                    types.String                                                           `tfsdk:"override_host" json:"override_host"`
	OverrideIPs                     *[]types.String                                                        `tfsdk:"override_ips" json:"override_ips"`
	PayloadLog                      *ZeroTrustGatewayPolicyRuleSettingsPayloadLogDataSourceModel           `tfsdk:"payload_log" json:"payload_log"`
	ResolveDNSThroughCloudflare     types.Bool                                                             `tfsdk:"resolve_dns_through_cloudflare" json:"resolve_dns_through_cloudflare"`
	UntrustedCERT                   *ZeroTrustGatewayPolicyRuleSettingsUntrustedCERTDataSourceModel        `tfsdk:"untrusted_cert" json:"untrusted_cert"`
}

type ZeroTrustGatewayPolicyRuleSettingsAuditSSHDataSourceModel struct {
	CommandLogging types.Bool `tfsdk:"command_logging" json:"command_logging"`
}

type ZeroTrustGatewayPolicyRuleSettingsBISOAdminControlsDataSourceModel struct {
	DCP types.Bool `tfsdk:"dcp" json:"dcp"`
	DD  types.Bool `tfsdk:"dd" json:"dd"`
	DK  types.Bool `tfsdk:"dk" json:"dk"`
	DP  types.Bool `tfsdk:"dp" json:"dp"`
	DU  types.Bool `tfsdk:"du" json:"du"`
}

type ZeroTrustGatewayPolicyRuleSettingsCheckSessionDataSourceModel struct {
	Duration types.String `tfsdk:"duration" json:"duration"`
	Enforce  types.Bool   `tfsdk:"enforce" json:"enforce"`
}

type ZeroTrustGatewayPolicyRuleSettingsDNSResolversDataSourceModel struct {
	IPV4 *[]*ZeroTrustGatewayPolicyRuleSettingsDNSResolversIPV4DataSourceModel `tfsdk:"ipv4" json:"ipv4"`
	IPV6 *[]*ZeroTrustGatewayPolicyRuleSettingsDNSResolversIPV6DataSourceModel `tfsdk:"ipv6" json:"ipv6"`
}

type ZeroTrustGatewayPolicyRuleSettingsDNSResolversIPV4DataSourceModel struct {
	IP                         types.String `tfsdk:"ip" json:"ip,computed"`
	Port                       types.Int64  `tfsdk:"port" json:"port"`
	RouteThroughPrivateNetwork types.Bool   `tfsdk:"route_through_private_network" json:"route_through_private_network"`
	VnetID                     types.String `tfsdk:"vnet_id" json:"vnet_id"`
}

type ZeroTrustGatewayPolicyRuleSettingsDNSResolversIPV6DataSourceModel struct {
	IP                         types.String `tfsdk:"ip" json:"ip,computed"`
	Port                       types.Int64  `tfsdk:"port" json:"port"`
	RouteThroughPrivateNetwork types.Bool   `tfsdk:"route_through_private_network" json:"route_through_private_network"`
	VnetID                     types.String `tfsdk:"vnet_id" json:"vnet_id"`
}

type ZeroTrustGatewayPolicyRuleSettingsEgressDataSourceModel struct {
	IPV4         types.String `tfsdk:"ipv4" json:"ipv4"`
	IPV4Fallback types.String `tfsdk:"ipv4_fallback" json:"ipv4_fallback"`
	IPV6         types.String `tfsdk:"ipv6" json:"ipv6"`
}

type ZeroTrustGatewayPolicyRuleSettingsL4overrideDataSourceModel struct {
	IP   types.String `tfsdk:"ip" json:"ip"`
	Port types.Int64  `tfsdk:"port" json:"port"`
}

type ZeroTrustGatewayPolicyRuleSettingsNotificationSettingsDataSourceModel struct {
	Enabled    types.Bool   `tfsdk:"enabled" json:"enabled"`
	Msg        types.String `tfsdk:"msg" json:"msg"`
	SupportURL types.String `tfsdk:"support_url" json:"support_url"`
}

type ZeroTrustGatewayPolicyRuleSettingsPayloadLogDataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled"`
}

type ZeroTrustGatewayPolicyRuleSettingsUntrustedCERTDataSourceModel struct {
	Action types.String `tfsdk:"action" json:"action"`
}

type ZeroTrustGatewayPolicyScheduleDataSourceModel struct {
	Fri      types.String `tfsdk:"fri" json:"fri"`
	Mon      types.String `tfsdk:"mon" json:"mon"`
	Sat      types.String `tfsdk:"sat" json:"sat"`
	Sun      types.String `tfsdk:"sun" json:"sun"`
	Thu      types.String `tfsdk:"thu" json:"thu"`
	TimeZone types.String `tfsdk:"time_zone" json:"time_zone"`
	Tue      types.String `tfsdk:"tue" json:"tue"`
	Wed      types.String `tfsdk:"wed" json:"wed"`
}

type ZeroTrustGatewayPolicyFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
