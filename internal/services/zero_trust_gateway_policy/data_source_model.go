// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_policy

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustGatewayPolicyResultDataSourceEnvelope struct {
	Result ZeroTrustGatewayPolicyDataSourceModel `json:"result,computed"`
}

type ZeroTrustGatewayPolicyDataSourceModel struct {
	ID            types.String                                                                `tfsdk:"id" path:"rule_id,computed"`
	RuleID        types.String                                                                `tfsdk:"rule_id" path:"rule_id,optional"`
	AccountID     types.String                                                                `tfsdk:"account_id" path:"account_id,required"`
	Action        types.String                                                                `tfsdk:"action" json:"action,computed"`
	CreatedAt     timetypes.RFC3339                                                           `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DeletedAt     timetypes.RFC3339                                                           `tfsdk:"deleted_at" json:"deleted_at,computed" format:"date-time"`
	Description   types.String                                                                `tfsdk:"description" json:"description,computed"`
	DevicePosture types.String                                                                `tfsdk:"device_posture" json:"device_posture,computed"`
	Enabled       types.Bool                                                                  `tfsdk:"enabled" json:"enabled,computed"`
	Identity      types.String                                                                `tfsdk:"identity" json:"identity,computed"`
	Name          types.String                                                                `tfsdk:"name" json:"name,computed"`
	Precedence    types.Int64                                                                 `tfsdk:"precedence" json:"precedence,computed"`
	Traffic       types.String                                                                `tfsdk:"traffic" json:"traffic,computed"`
	UpdatedAt     timetypes.RFC3339                                                           `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Version       types.Int64                                                                 `tfsdk:"version" json:"version,computed"`
	WarningStatus types.String                                                                `tfsdk:"warning_status" json:"warning_status,computed"`
	Filters       customfield.List[types.String]                                              `tfsdk:"filters" json:"filters,computed"`
	Expiration    customfield.NestedObject[ZeroTrustGatewayPolicyExpirationDataSourceModel]   `tfsdk:"expiration" json:"expiration,computed"`
	RuleSettings  customfield.NestedObject[ZeroTrustGatewayPolicyRuleSettingsDataSourceModel] `tfsdk:"rule_settings" json:"rule_settings,computed"`
	Schedule      customfield.NestedObject[ZeroTrustGatewayPolicyScheduleDataSourceModel]     `tfsdk:"schedule" json:"schedule,computed"`
}

func (m *ZeroTrustGatewayPolicyDataSourceModel) toReadParams(_ context.Context) (params zero_trust.GatewayRuleGetParams, diags diag.Diagnostics) {
	params = zero_trust.GatewayRuleGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustGatewayPolicyExpirationDataSourceModel struct {
	ExpiresAt timetypes.RFC3339 `tfsdk:"expires_at" json:"expires_at,computed" format:"date-time"`
	Duration  types.Int64       `tfsdk:"duration" json:"duration,computed"`
	Expired   types.Bool        `tfsdk:"expired" json:"expired,computed"`
}

type ZeroTrustGatewayPolicyRuleSettingsDataSourceModel struct {
	AddHeaders                      customfield.Map[types.String]                                                                   `tfsdk:"add_headers" json:"add_headers,computed"`
	AllowChildBypass                types.Bool                                                                                      `tfsdk:"allow_child_bypass" json:"allow_child_bypass,computed"`
	AuditSSH                        customfield.NestedObject[ZeroTrustGatewayPolicyRuleSettingsAuditSSHDataSourceModel]             `tfsdk:"audit_ssh" json:"audit_ssh,computed"`
	BISOAdminControls               customfield.NestedObject[ZeroTrustGatewayPolicyRuleSettingsBISOAdminControlsDataSourceModel]    `tfsdk:"biso_admin_controls" json:"biso_admin_controls,computed"`
	BlockPage                       customfield.NestedObject[ZeroTrustGatewayPolicyRuleSettingsBlockPageDataSourceModel]            `tfsdk:"block_page" json:"block_page,computed"`
	BlockPageEnabled                types.Bool                                                                                      `tfsdk:"block_page_enabled" json:"block_page_enabled,computed"`
	BlockReason                     types.String                                                                                    `tfsdk:"block_reason" json:"block_reason,computed"`
	BypassParentRule                types.Bool                                                                                      `tfsdk:"bypass_parent_rule" json:"bypass_parent_rule,computed"`
	CheckSession                    customfield.NestedObject[ZeroTrustGatewayPolicyRuleSettingsCheckSessionDataSourceModel]         `tfsdk:"check_session" json:"check_session,computed"`
	DNSResolvers                    customfield.NestedObject[ZeroTrustGatewayPolicyRuleSettingsDNSResolversDataSourceModel]         `tfsdk:"dns_resolvers" json:"dns_resolvers,computed"`
	Egress                          customfield.NestedObject[ZeroTrustGatewayPolicyRuleSettingsEgressDataSourceModel]               `tfsdk:"egress" json:"egress,computed"`
	IgnoreCNAMECategoryMatches      types.Bool                                                                                      `tfsdk:"ignore_cname_category_matches" json:"ignore_cname_category_matches,computed"`
	InsecureDisableDNSSECValidation types.Bool                                                                                      `tfsdk:"insecure_disable_dnssec_validation" json:"insecure_disable_dnssec_validation,computed"`
	IPCategories                    types.Bool                                                                                      `tfsdk:"ip_categories" json:"ip_categories,computed"`
	IPIndicatorFeeds                types.Bool                                                                                      `tfsdk:"ip_indicator_feeds" json:"ip_indicator_feeds,computed"`
	L4override                      customfield.NestedObject[ZeroTrustGatewayPolicyRuleSettingsL4overrideDataSourceModel]           `tfsdk:"l4override" json:"l4override,computed"`
	NotificationSettings            customfield.NestedObject[ZeroTrustGatewayPolicyRuleSettingsNotificationSettingsDataSourceModel] `tfsdk:"notification_settings" json:"notification_settings,computed"`
	OverrideHost                    types.String                                                                                    `tfsdk:"override_host" json:"override_host,computed"`
	OverrideIPs                     customfield.List[types.String]                                                                  `tfsdk:"override_ips" json:"override_ips,computed"`
	PayloadLog                      customfield.NestedObject[ZeroTrustGatewayPolicyRuleSettingsPayloadLogDataSourceModel]           `tfsdk:"payload_log" json:"payload_log,computed"`
	Quarantine                      customfield.NestedObject[ZeroTrustGatewayPolicyRuleSettingsQuarantineDataSourceModel]           `tfsdk:"quarantine" json:"quarantine,computed"`
	Redirect                        customfield.NestedObject[ZeroTrustGatewayPolicyRuleSettingsRedirectDataSourceModel]             `tfsdk:"redirect" json:"redirect,computed"`
	ResolveDNSInternally            customfield.NestedObject[ZeroTrustGatewayPolicyRuleSettingsResolveDNSInternallyDataSourceModel] `tfsdk:"resolve_dns_internally" json:"resolve_dns_internally,computed"`
	ResolveDNSThroughCloudflare     types.Bool                                                                                      `tfsdk:"resolve_dns_through_cloudflare" json:"resolve_dns_through_cloudflare,computed"`
	UntrustedCERT                   customfield.NestedObject[ZeroTrustGatewayPolicyRuleSettingsUntrustedCERTDataSourceModel]        `tfsdk:"untrusted_cert" json:"untrusted_cert,computed"`
}

type ZeroTrustGatewayPolicyRuleSettingsAuditSSHDataSourceModel struct {
	CommandLogging types.Bool `tfsdk:"command_logging" json:"command_logging,computed"`
}

type ZeroTrustGatewayPolicyRuleSettingsBISOAdminControlsDataSourceModel struct {
	Copy     types.String `tfsdk:"copy" json:"copy,computed"`
	DCP      types.Bool   `tfsdk:"dcp" json:"dcp,computed"`
	DD       types.Bool   `tfsdk:"dd" json:"dd,computed"`
	DK       types.Bool   `tfsdk:"dk" json:"dk,computed"`
	Download types.String `tfsdk:"download" json:"download,computed"`
	DP       types.Bool   `tfsdk:"dp" json:"dp,computed"`
	DU       types.Bool   `tfsdk:"du" json:"du,computed"`
	Keyboard types.String `tfsdk:"keyboard" json:"keyboard,computed"`
	Paste    types.String `tfsdk:"paste" json:"paste,computed"`
	Printing types.String `tfsdk:"printing" json:"printing,computed"`
	Upload   types.String `tfsdk:"upload" json:"upload,computed"`
	Version  types.String `tfsdk:"version" json:"version,computed"`
}

type ZeroTrustGatewayPolicyRuleSettingsBlockPageDataSourceModel struct {
	TargetURI      types.String `tfsdk:"target_uri" json:"target_uri,computed"`
	IncludeContext types.Bool   `tfsdk:"include_context" json:"include_context,computed"`
}

type ZeroTrustGatewayPolicyRuleSettingsCheckSessionDataSourceModel struct {
	Duration types.String `tfsdk:"duration" json:"duration,computed"`
	Enforce  types.Bool   `tfsdk:"enforce" json:"enforce,computed"`
}

type ZeroTrustGatewayPolicyRuleSettingsDNSResolversDataSourceModel struct {
	IPV4 customfield.NestedObjectList[ZeroTrustGatewayPolicyRuleSettingsDNSResolversIPV4DataSourceModel] `tfsdk:"ipv4" json:"ipv4,computed"`
	IPV6 customfield.NestedObjectList[ZeroTrustGatewayPolicyRuleSettingsDNSResolversIPV6DataSourceModel] `tfsdk:"ipv6" json:"ipv6,computed"`
}

type ZeroTrustGatewayPolicyRuleSettingsDNSResolversIPV4DataSourceModel struct {
	IP                         types.String `tfsdk:"ip" json:"ip,computed"`
	Port                       types.Int64  `tfsdk:"port" json:"port,computed"`
	RouteThroughPrivateNetwork types.Bool   `tfsdk:"route_through_private_network" json:"route_through_private_network,computed"`
	VnetID                     types.String `tfsdk:"vnet_id" json:"vnet_id,computed"`
}

type ZeroTrustGatewayPolicyRuleSettingsDNSResolversIPV6DataSourceModel struct {
	IP                         types.String `tfsdk:"ip" json:"ip,computed"`
	Port                       types.Int64  `tfsdk:"port" json:"port,computed"`
	RouteThroughPrivateNetwork types.Bool   `tfsdk:"route_through_private_network" json:"route_through_private_network,computed"`
	VnetID                     types.String `tfsdk:"vnet_id" json:"vnet_id,computed"`
}

type ZeroTrustGatewayPolicyRuleSettingsEgressDataSourceModel struct {
	IPV4         types.String `tfsdk:"ipv4" json:"ipv4,computed"`
	IPV4Fallback types.String `tfsdk:"ipv4_fallback" json:"ipv4_fallback,computed"`
	IPV6         types.String `tfsdk:"ipv6" json:"ipv6,computed"`
}

type ZeroTrustGatewayPolicyRuleSettingsL4overrideDataSourceModel struct {
	IP   types.String `tfsdk:"ip" json:"ip,computed"`
	Port types.Int64  `tfsdk:"port" json:"port,computed"`
}

type ZeroTrustGatewayPolicyRuleSettingsNotificationSettingsDataSourceModel struct {
	Enabled        types.Bool   `tfsdk:"enabled" json:"enabled,computed"`
	IncludeContext types.Bool   `tfsdk:"include_context" json:"include_context,computed"`
	Msg            types.String `tfsdk:"msg" json:"msg,computed"`
	SupportURL     types.String `tfsdk:"support_url" json:"support_url,computed"`
}

type ZeroTrustGatewayPolicyRuleSettingsPayloadLogDataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed"`
}

type ZeroTrustGatewayPolicyRuleSettingsQuarantineDataSourceModel struct {
	FileTypes customfield.List[types.String] `tfsdk:"file_types" json:"file_types,computed"`
}

type ZeroTrustGatewayPolicyRuleSettingsRedirectDataSourceModel struct {
	TargetURI            types.String `tfsdk:"target_uri" json:"target_uri,computed"`
	IncludeContext       types.Bool   `tfsdk:"include_context" json:"include_context,computed"`
	PreservePathAndQuery types.Bool   `tfsdk:"preserve_path_and_query" json:"preserve_path_and_query,computed"`
}

type ZeroTrustGatewayPolicyRuleSettingsResolveDNSInternallyDataSourceModel struct {
	Fallback types.String `tfsdk:"fallback" json:"fallback,computed"`
	ViewID   types.String `tfsdk:"view_id" json:"view_id,computed"`
}

type ZeroTrustGatewayPolicyRuleSettingsUntrustedCERTDataSourceModel struct {
	Action types.String `tfsdk:"action" json:"action,computed"`
}

type ZeroTrustGatewayPolicyScheduleDataSourceModel struct {
	Fri      types.String `tfsdk:"fri" json:"fri,computed"`
	Mon      types.String `tfsdk:"mon" json:"mon,computed"`
	Sat      types.String `tfsdk:"sat" json:"sat,computed"`
	Sun      types.String `tfsdk:"sun" json:"sun,computed"`
	Thu      types.String `tfsdk:"thu" json:"thu,computed"`
	TimeZone types.String `tfsdk:"time_zone" json:"time_zone,computed"`
	Tue      types.String `tfsdk:"tue" json:"tue,computed"`
	Wed      types.String `tfsdk:"wed" json:"wed,computed"`
}
