// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package teams_rule

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TeamsRulesResultListDataSourceEnvelope struct {
	Result *[]*TeamsRulesItemsDataSourceModel `json:"result,computed"`
}

type TeamsRulesDataSourceModel struct {
	AccountID types.String                       `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                        `tfsdk:"max_items"`
	Items     *[]*TeamsRulesItemsDataSourceModel `tfsdk:"items"`
}

type TeamsRulesItemsDataSourceModel struct {
	ID            types.String    `tfsdk:"id" json:"id,computed"`
	Action        types.String    `tfsdk:"action" json:"action,computed"`
	CreatedAt     types.String    `tfsdk:"created_at" json:"created_at,computed"`
	DeletedAt     types.String    `tfsdk:"deleted_at" json:"deleted_at,computed"`
	Description   types.String    `tfsdk:"description" json:"description,computed"`
	DevicePosture types.String    `tfsdk:"device_posture" json:"device_posture,computed"`
	Enabled       types.Bool      `tfsdk:"enabled" json:"enabled,computed"`
	Filters       *[]types.String `tfsdk:"filters" json:"filters,computed"`
	Identity      types.String    `tfsdk:"identity" json:"identity,computed"`
	Name          types.String    `tfsdk:"name" json:"name,computed"`
	Precedence    types.Int64     `tfsdk:"precedence" json:"precedence,computed"`
	Traffic       types.String    `tfsdk:"traffic" json:"traffic,computed"`
	UpdatedAt     types.String    `tfsdk:"updated_at" json:"updated_at,computed"`
}

type TeamsRulesItemsRuleSettingsDataSourceModel struct {
	AddHeaders                      types.String    `tfsdk:"add_headers" json:"add_headers,computed"`
	AllowChildBypass                types.Bool      `tfsdk:"allow_child_bypass" json:"allow_child_bypass,computed"`
	BlockPageEnabled                types.Bool      `tfsdk:"block_page_enabled" json:"block_page_enabled,computed"`
	BlockReason                     types.String    `tfsdk:"block_reason" json:"block_reason,computed"`
	BypassParentRule                types.Bool      `tfsdk:"bypass_parent_rule" json:"bypass_parent_rule,computed"`
	IgnoreCNAMECategoryMatches      types.Bool      `tfsdk:"ignore_cname_category_matches" json:"ignore_cname_category_matches,computed"`
	InsecureDisableDNSSECValidation types.Bool      `tfsdk:"insecure_disable_dnssec_validation" json:"insecure_disable_dnssec_validation,computed"`
	IPCategories                    types.Bool      `tfsdk:"ip_categories" json:"ip_categories,computed"`
	IPIndicatorFeeds                types.Bool      `tfsdk:"ip_indicator_feeds" json:"ip_indicator_feeds,computed"`
	OverrideHost                    types.String    `tfsdk:"override_host" json:"override_host,computed"`
	OverrideIPs                     *[]types.String `tfsdk:"override_ips" json:"override_ips,computed"`
	ResolveDNSThroughCloudflare     types.Bool      `tfsdk:"resolve_dns_through_cloudflare" json:"resolve_dns_through_cloudflare,computed"`
}

type TeamsRulesItemsRuleSettingsAuditSSHDataSourceModel struct {
	CommandLogging types.Bool `tfsdk:"command_logging" json:"command_logging,computed"`
}

type TeamsRulesItemsRuleSettingsBisoAdminControlsDataSourceModel struct {
	DCP types.Bool `tfsdk:"dcp" json:"dcp,computed"`
	DD  types.Bool `tfsdk:"dd" json:"dd,computed"`
	DK  types.Bool `tfsdk:"dk" json:"dk,computed"`
	DP  types.Bool `tfsdk:"dp" json:"dp,computed"`
	DU  types.Bool `tfsdk:"du" json:"du,computed"`
}

type TeamsRulesItemsRuleSettingsCheckSessionDataSourceModel struct {
	Duration types.String `tfsdk:"duration" json:"duration,computed"`
	Enforce  types.Bool   `tfsdk:"enforce" json:"enforce,computed"`
}

type TeamsRulesItemsRuleSettingsDNSResolversDataSourceModel struct {
	IPV4 *[]*TeamsRulesItemsRuleSettingsDNSResolversIPV4DataSourceModel `tfsdk:"ipv4" json:"ipv4,computed"`
	IPV6 *[]*TeamsRulesItemsRuleSettingsDNSResolversIPV6DataSourceModel `tfsdk:"ipv6" json:"ipv6,computed"`
}

type TeamsRulesItemsRuleSettingsDNSResolversIPV4DataSourceModel struct {
	IP                         types.String `tfsdk:"ip" json:"ip,computed"`
	Port                       types.Int64  `tfsdk:"port" json:"port,computed"`
	RouteThroughPrivateNetwork types.Bool   `tfsdk:"route_through_private_network" json:"route_through_private_network,computed"`
	VnetID                     types.String `tfsdk:"vnet_id" json:"vnet_id,computed"`
}

type TeamsRulesItemsRuleSettingsDNSResolversIPV6DataSourceModel struct {
	IP                         types.String `tfsdk:"ip" json:"ip,computed"`
	Port                       types.Int64  `tfsdk:"port" json:"port,computed"`
	RouteThroughPrivateNetwork types.Bool   `tfsdk:"route_through_private_network" json:"route_through_private_network,computed"`
	VnetID                     types.String `tfsdk:"vnet_id" json:"vnet_id,computed"`
}

type TeamsRulesItemsRuleSettingsEgressDataSourceModel struct {
	IPV4         types.String `tfsdk:"ipv4" json:"ipv4,computed"`
	IPV4Fallback types.String `tfsdk:"ipv4_fallback" json:"ipv4_fallback,computed"`
	IPV6         types.String `tfsdk:"ipv6" json:"ipv6,computed"`
}

type TeamsRulesItemsRuleSettingsL4overrideDataSourceModel struct {
	IP   types.String `tfsdk:"ip" json:"ip,computed"`
	Port types.Int64  `tfsdk:"port" json:"port,computed"`
}

type TeamsRulesItemsRuleSettingsNotificationSettingsDataSourceModel struct {
	Enabled    types.Bool   `tfsdk:"enabled" json:"enabled,computed"`
	Msg        types.String `tfsdk:"msg" json:"msg,computed"`
	SupportURL types.String `tfsdk:"support_url" json:"support_url,computed"`
}

type TeamsRulesItemsRuleSettingsPayloadLogDataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed"`
}

type TeamsRulesItemsRuleSettingsUntrustedCERTDataSourceModel struct {
	Action types.String `tfsdk:"action" json:"action,computed"`
}

type TeamsRulesItemsScheduleDataSourceModel struct {
	Fri      types.String `tfsdk:"fri" json:"fri,computed"`
	Mon      types.String `tfsdk:"mon" json:"mon,computed"`
	Sat      types.String `tfsdk:"sat" json:"sat,computed"`
	Sun      types.String `tfsdk:"sun" json:"sun,computed"`
	Thu      types.String `tfsdk:"thu" json:"thu,computed"`
	TimeZone types.String `tfsdk:"time_zone" json:"time_zone,computed"`
	Tue      types.String `tfsdk:"tue" json:"tue,computed"`
	Wed      types.String `tfsdk:"wed" json:"wed,computed"`
}
