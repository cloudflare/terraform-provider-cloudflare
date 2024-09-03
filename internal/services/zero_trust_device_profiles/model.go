// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_profiles

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDeviceProfilesResultEnvelope struct {
	Result ZeroTrustDeviceProfilesModel `json:"result"`
}

type ZeroTrustDeviceProfilesModel struct {
	AccountID           types.String                                                              `tfsdk:"account_id" path:"account_id"`
	PolicyID            types.String                                                              `tfsdk:"policy_id" path:"policy_id"`
	LANAllowMinutes     types.Float64                                                             `tfsdk:"lan_allow_minutes" json:"lan_allow_minutes"`
	LANAllowSubnetSize  types.Float64                                                             `tfsdk:"lan_allow_subnet_size" json:"lan_allow_subnet_size"`
	Match               types.String                                                              `tfsdk:"match" json:"match"`
	Name                types.String                                                              `tfsdk:"name" json:"name"`
	Precedence          types.Float64                                                             `tfsdk:"precedence" json:"precedence"`
	AllowModeSwitch     types.Bool                                                                `tfsdk:"allow_mode_switch" json:"allow_mode_switch"`
	AllowUpdates        types.Bool                                                                `tfsdk:"allow_updates" json:"allow_updates"`
	AllowedToLeave      types.Bool                                                                `tfsdk:"allowed_to_leave" json:"allowed_to_leave"`
	AutoConnect         types.Float64                                                             `tfsdk:"auto_connect" json:"auto_connect"`
	CaptivePortal       types.Float64                                                             `tfsdk:"captive_portal" json:"captive_portal"`
	Description         types.String                                                              `tfsdk:"description" json:"description"`
	DisableAutoFallback types.Bool                                                                `tfsdk:"disable_auto_fallback" json:"disable_auto_fallback"`
	Enabled             types.Bool                                                                `tfsdk:"enabled" json:"enabled"`
	ExcludeOfficeIPs    types.Bool                                                                `tfsdk:"exclude_office_ips" json:"exclude_office_ips"`
	SupportURL          types.String                                                              `tfsdk:"support_url" json:"support_url"`
	SwitchLocked        types.Bool                                                                `tfsdk:"switch_locked" json:"switch_locked"`
	TunnelProtocol      types.String                                                              `tfsdk:"tunnel_protocol" json:"tunnel_protocol"`
	ServiceModeV2       *ZeroTrustDeviceProfilesServiceModeV2Model                                `tfsdk:"service_mode_v2" json:"service_mode_v2"`
	Default             types.Bool                                                                `tfsdk:"default" json:"default,computed"`
	GatewayUniqueID     types.String                                                              `tfsdk:"gateway_unique_id" json:"gateway_unique_id,computed"`
	Exclude             customfield.NestedObjectList[ZeroTrustDeviceProfilesExcludeModel]         `tfsdk:"exclude" json:"exclude,computed"`
	FallbackDomains     customfield.NestedObjectList[ZeroTrustDeviceProfilesFallbackDomainsModel] `tfsdk:"fallback_domains" json:"fallback_domains,computed"`
	Include             customfield.NestedObjectList[ZeroTrustDeviceProfilesIncludeModel]         `tfsdk:"include" json:"include,computed"`
	TargetTests         customfield.NestedObjectList[ZeroTrustDeviceProfilesTargetTestsModel]     `tfsdk:"target_tests" json:"target_tests,computed"`
}

type ZeroTrustDeviceProfilesServiceModeV2Model struct {
	Mode types.String  `tfsdk:"mode" json:"mode,computed_optional"`
	Port types.Float64 `tfsdk:"port" json:"port,computed_optional"`
}

type ZeroTrustDeviceProfilesExcludeModel struct {
	Address     types.String `tfsdk:"address" json:"address"`
	Description types.String `tfsdk:"description" json:"description"`
	Host        types.String `tfsdk:"host" json:"host,computed_optional"`
}

type ZeroTrustDeviceProfilesFallbackDomainsModel struct {
	Suffix      types.String                   `tfsdk:"suffix" json:"suffix"`
	Description types.String                   `tfsdk:"description" json:"description,computed_optional"`
	DNSServer   customfield.List[types.String] `tfsdk:"dns_server" json:"dns_server,computed_optional"`
}

type ZeroTrustDeviceProfilesIncludeModel struct {
	Address     types.String `tfsdk:"address" json:"address"`
	Description types.String `tfsdk:"description" json:"description"`
	Host        types.String `tfsdk:"host" json:"host,computed_optional"`
}

type ZeroTrustDeviceProfilesTargetTestsModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed_optional"`
	Name types.String `tfsdk:"name" json:"name,computed_optional"`
}
