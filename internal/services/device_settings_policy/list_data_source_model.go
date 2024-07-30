// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package device_settings_policy

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DeviceSettingsPoliciesResultListDataSourceEnvelope struct {
	Result *[]*DeviceSettingsPoliciesResultDataSourceModel `json:"result,computed"`
}

type DeviceSettingsPoliciesDataSourceModel struct {
	AccountID types.String                                    `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                                     `tfsdk:"max_items"`
	Result    *[]*DeviceSettingsPoliciesResultDataSourceModel `tfsdk:"result"`
}

type DeviceSettingsPoliciesResultDataSourceModel struct {
	AllowModeSwitch     types.Bool                                               `tfsdk:"allow_mode_switch" json:"allow_mode_switch"`
	AllowUpdates        types.Bool                                               `tfsdk:"allow_updates" json:"allow_updates"`
	AllowedToLeave      types.Bool                                               `tfsdk:"allowed_to_leave" json:"allowed_to_leave"`
	AutoConnect         types.Float64                                            `tfsdk:"auto_connect" json:"auto_connect"`
	CaptivePortal       types.Float64                                            `tfsdk:"captive_portal" json:"captive_portal"`
	Default             types.Bool                                               `tfsdk:"default" json:"default"`
	Description         types.String                                             `tfsdk:"description" json:"description"`
	DisableAutoFallback types.Bool                                               `tfsdk:"disable_auto_fallback" json:"disable_auto_fallback"`
	Enabled             types.Bool                                               `tfsdk:"enabled" json:"enabled"`
	Exclude             *[]*DeviceSettingsPoliciesExcludeDataSourceModel         `tfsdk:"exclude" json:"exclude"`
	ExcludeOfficeIPs    types.Bool                                               `tfsdk:"exclude_office_ips" json:"exclude_office_ips"`
	FallbackDomains     *[]*DeviceSettingsPoliciesFallbackDomainsDataSourceModel `tfsdk:"fallback_domains" json:"fallback_domains"`
	GatewayUniqueID     types.String                                             `tfsdk:"gateway_unique_id" json:"gateway_unique_id"`
	Include             *[]*DeviceSettingsPoliciesIncludeDataSourceModel         `tfsdk:"include" json:"include"`
	LANAllowMinutes     types.Float64                                            `tfsdk:"lan_allow_minutes" json:"lan_allow_minutes"`
	LANAllowSubnetSize  types.Float64                                            `tfsdk:"lan_allow_subnet_size" json:"lan_allow_subnet_size"`
	Match               types.String                                             `tfsdk:"match" json:"match"`
	Name                types.String                                             `tfsdk:"name" json:"name"`
	PolicyID            types.String                                             `tfsdk:"policy_id" json:"policy_id"`
	Precedence          types.Float64                                            `tfsdk:"precedence" json:"precedence"`
	ServiceModeV2       *DeviceSettingsPoliciesServiceModeV2DataSourceModel      `tfsdk:"service_mode_v2" json:"service_mode_v2"`
	SupportURL          types.String                                             `tfsdk:"support_url" json:"support_url"`
	SwitchLocked        types.Bool                                               `tfsdk:"switch_locked" json:"switch_locked"`
	TargetTests         *[]*DeviceSettingsPoliciesTargetTestsDataSourceModel     `tfsdk:"target_tests" json:"target_tests"`
	TunnelProtocol      types.String                                             `tfsdk:"tunnel_protocol" json:"tunnel_protocol"`
}

type DeviceSettingsPoliciesExcludeDataSourceModel struct {
	Address     types.String `tfsdk:"address" json:"address,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Host        types.String `tfsdk:"host" json:"host"`
}

type DeviceSettingsPoliciesFallbackDomainsDataSourceModel struct {
	Suffix      types.String            `tfsdk:"suffix" json:"suffix,computed"`
	Description types.String            `tfsdk:"description" json:"description"`
	DNSServer   *[]jsontypes.Normalized `tfsdk:"dns_server" json:"dns_server"`
}

type DeviceSettingsPoliciesIncludeDataSourceModel struct {
	Address     types.String `tfsdk:"address" json:"address,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Host        types.String `tfsdk:"host" json:"host"`
}

type DeviceSettingsPoliciesServiceModeV2DataSourceModel struct {
	Mode types.String  `tfsdk:"mode" json:"mode"`
	Port types.Float64 `tfsdk:"port" json:"port"`
}

type DeviceSettingsPoliciesTargetTestsDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id"`
	Name types.String `tfsdk:"name" json:"name"`
}
