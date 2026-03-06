// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v500

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Source models represent the v4 SDKv2 provider structure
// These match the v4 cloudflare_zero_trust_device_profiles / cloudflare_device_settings_policy resources

type SourceDeviceProfileModel struct {
	ID                  types.String  `tfsdk:"id"`
	AccountID           types.String  `tfsdk:"account_id"`
	Default             types.Bool    `tfsdk:"default"`
	Name                types.String  `tfsdk:"name"`
	Description         types.String  `tfsdk:"description"`
	Precedence          types.Int64   `tfsdk:"precedence"`
	Match               types.String  `tfsdk:"match"`
	Enabled             types.Bool    `tfsdk:"enabled"`
	DisableAutoFallback types.Bool    `tfsdk:"disable_auto_fallback"`
	CaptivePortal       types.Int64   `tfsdk:"captive_portal"`
	AllowModeSwitch     types.Bool    `tfsdk:"allow_mode_switch"`
	SwitchLocked        types.Bool    `tfsdk:"switch_locked"`
	AllowUpdates        types.Bool    `tfsdk:"allow_updates"`
	AutoConnect         types.Int64   `tfsdk:"auto_connect"`
	AllowedToLeave      types.Bool    `tfsdk:"allowed_to_leave"`
	SupportURL          types.String  `tfsdk:"support_url"`
	ServiceModeV2Mode   types.String  `tfsdk:"service_mode_v2_mode"`
	ServiceModeV2Port   types.Int64   `tfsdk:"service_mode_v2_port"`
	ExcludeOfficeIPs    types.Bool    `tfsdk:"exclude_office_ips"`
	TunnelProtocol      types.String  `tfsdk:"tunnel_protocol"`
	LANAllowMinutes     types.Float64 `tfsdk:"lan_allow_minutes"`
	LANAllowSubnetSize  types.Float64 `tfsdk:"lan_allow_subnet_size"`
	// Note: exclude and include are Optional+Computed in v4, present in some configs
	// Note: fallback_domains existed in v4 but will be removed during migration
}

// Target models represent the v5 Plugin Framework provider structure
// These match the cloudflare_zero_trust_device_default_profile resource

type TargetDefaultProfileModel struct {
	ID                         types.String                                                                       `tfsdk:"id"`
	AccountID                  types.String                                                                       `tfsdk:"account_id"`
	LANAllowMinutes            types.Float64                                                                      `tfsdk:"lan_allow_minutes"`
	LANAllowSubnetSize         types.Float64                                                                      `tfsdk:"lan_allow_subnet_size"`
	AllowModeSwitch            types.Bool                                                                         `tfsdk:"allow_mode_switch"`
	AllowUpdates               types.Bool                                                                         `tfsdk:"allow_updates"`
	AllowedToLeave             types.Bool                                                                         `tfsdk:"allowed_to_leave"`
	AutoConnect                types.Float64                                                                      `tfsdk:"auto_connect"`
	CaptivePortal              types.Float64                                                                      `tfsdk:"captive_portal"`
	DisableAutoFallback        types.Bool                                                                         `tfsdk:"disable_auto_fallback"`
	ExcludeOfficeIPs           types.Bool                                                                         `tfsdk:"exclude_office_ips"`
	RegisterInterfaceIPWithDNS types.Bool                                                                         `tfsdk:"register_interface_ip_with_dns"`
	SccmVpnBoundarySupport     types.Bool                                                                         `tfsdk:"sccm_vpn_boundary_support"`
	SupportURL                 types.String                                                                       `tfsdk:"support_url"`
	SwitchLocked               types.Bool                                                                         `tfsdk:"switch_locked"`
	TunnelProtocol             types.String                                                                       `tfsdk:"tunnel_protocol"`
	Exclude                    customfield.NestedObjectList[TargetDefaultProfileExcludeModel]                     `tfsdk:"exclude"`
	Include                    customfield.NestedObjectList[TargetDefaultProfileIncludeModel]                     `tfsdk:"include"`
	ServiceModeV2              customfield.NestedObject[TargetDefaultProfileServiceModeV2Model]                   `tfsdk:"service_mode_v2"`
	Default                    types.Bool                                                                         `tfsdk:"default"`
	Enabled                    types.Bool                                                                         `tfsdk:"enabled"`
	GatewayUniqueID            types.String                                                                       `tfsdk:"gateway_unique_id"`
	FallbackDomains            customfield.NestedObjectList[TargetDefaultProfileFallbackDomainsModel]             `tfsdk:"fallback_domains"`
}

type TargetDefaultProfileExcludeModel struct {
	Address     types.String `tfsdk:"address"`
	Description types.String `tfsdk:"description"`
	Host        types.String `tfsdk:"host"`
}

type TargetDefaultProfileIncludeModel struct {
	Address     types.String `tfsdk:"address"`
	Description types.String `tfsdk:"description"`
	Host        types.String `tfsdk:"host"`
}

type TargetDefaultProfileServiceModeV2Model struct {
	Mode types.String  `tfsdk:"mode"`
	Port types.Float64 `tfsdk:"port"`
}

type TargetDefaultProfileFallbackDomainsModel struct {
	Suffix      types.String                   `tfsdk:"suffix"`
	Description types.String                   `tfsdk:"description"`
	DNSServer   customfield.List[types.String] `tfsdk:"dns_server"`
}
