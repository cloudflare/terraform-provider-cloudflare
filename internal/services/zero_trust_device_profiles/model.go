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
	PolicyID            types.String                                                              `tfsdk:"policy_id" path:"policy_id,computed_optional"`
	LANAllowMinutes     types.Float64                                                             `tfsdk:"lan_allow_minutes" json:"lan_allow_minutes,computed_optional"`
	LANAllowSubnetSize  types.Float64                                                             `tfsdk:"lan_allow_subnet_size" json:"lan_allow_subnet_size,computed_optional"`
	AllowModeSwitch     types.Bool                                                                `tfsdk:"allow_mode_switch" json:"allow_mode_switch,computed_optional"`
	AllowUpdates        types.Bool                                                                `tfsdk:"allow_updates" json:"allow_updates,computed_optional"`
	AllowedToLeave      types.Bool                                                                `tfsdk:"allowed_to_leave" json:"allowed_to_leave,computed_optional"`
	AutoConnect         types.Float64                                                             `tfsdk:"auto_connect" json:"auto_connect,computed_optional"`
	CaptivePortal       types.Float64                                                             `tfsdk:"captive_portal" json:"captive_portal,computed_optional"`
	Description         types.String                                                              `tfsdk:"description" json:"description,computed_optional"`
	DisableAutoFallback types.Bool                                                                `tfsdk:"disable_auto_fallback" json:"disable_auto_fallback,computed_optional"`
	Enabled             types.Bool                                                                `tfsdk:"enabled" json:"enabled,computed_optional"`
	ExcludeOfficeIPs    types.Bool                                                                `tfsdk:"exclude_office_ips" json:"exclude_office_ips,computed_optional"`
	Match               types.String                                                              `tfsdk:"match" json:"match,computed_optional"`
	Name                types.String                                                              `tfsdk:"name" json:"name,computed_optional"`
	Precedence          types.Float64                                                             `tfsdk:"precedence" json:"precedence,computed_optional"`
	SupportURL          types.String                                                              `tfsdk:"support_url" json:"support_url,computed_optional"`
	SwitchLocked        types.Bool                                                                `tfsdk:"switch_locked" json:"switch_locked,computed_optional"`
	TunnelProtocol      types.String                                                              `tfsdk:"tunnel_protocol" json:"tunnel_protocol,computed_optional"`
	ServiceModeV2       customfield.NestedObject[ZeroTrustDeviceProfilesServiceModeV2Model]       `tfsdk:"service_mode_v2" json:"service_mode_v2,computed_optional"`
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
	Address     types.String `tfsdk:"address" json:"address,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Host        types.String `tfsdk:"host" json:"host,computed"`
}

type ZeroTrustDeviceProfilesFallbackDomainsModel struct {
	Suffix      types.String                   `tfsdk:"suffix" json:"suffix,computed"`
	Description types.String                   `tfsdk:"description" json:"description,computed"`
	DNSServer   customfield.List[types.String] `tfsdk:"dns_server" json:"dns_server,computed"`
}

type ZeroTrustDeviceProfilesIncludeModel struct {
	Address     types.String `tfsdk:"address" json:"address,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Host        types.String `tfsdk:"host" json:"host,computed"`
}

type ZeroTrustDeviceProfilesTargetTestsModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}
