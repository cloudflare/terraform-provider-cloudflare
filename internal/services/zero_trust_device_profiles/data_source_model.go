// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_profiles

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDeviceProfilesResultDataSourceEnvelope struct {
	Result ZeroTrustDeviceProfilesDataSourceModel `json:"result,computed"`
}

type ZeroTrustDeviceProfilesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustDeviceProfilesDataSourceModel] `json:"result,computed"`
}

type ZeroTrustDeviceProfilesDataSourceModel struct {
	AccountID           types.String                                              `tfsdk:"account_id" path:"account_id"`
	PolicyID            types.String                                              `tfsdk:"policy_id" path:"policy_id"`
	AllowModeSwitch     types.Bool                                                `tfsdk:"allow_mode_switch" json:"allow_mode_switch"`
	AllowUpdates        types.Bool                                                `tfsdk:"allow_updates" json:"allow_updates"`
	AllowedToLeave      types.Bool                                                `tfsdk:"allowed_to_leave" json:"allowed_to_leave"`
	AutoConnect         types.Float64                                             `tfsdk:"auto_connect" json:"auto_connect"`
	CaptivePortal       types.Float64                                             `tfsdk:"captive_portal" json:"captive_portal"`
	Default             types.Bool                                                `tfsdk:"default" json:"default"`
	Description         types.String                                              `tfsdk:"description" json:"description"`
	DisableAutoFallback types.Bool                                                `tfsdk:"disable_auto_fallback" json:"disable_auto_fallback"`
	Enabled             types.Bool                                                `tfsdk:"enabled" json:"enabled"`
	ExcludeOfficeIPs    types.Bool                                                `tfsdk:"exclude_office_ips" json:"exclude_office_ips"`
	GatewayUniqueID     types.String                                              `tfsdk:"gateway_unique_id" json:"gateway_unique_id"`
	LANAllowMinutes     types.Float64                                             `tfsdk:"lan_allow_minutes" json:"lan_allow_minutes"`
	LANAllowSubnetSize  types.Float64                                             `tfsdk:"lan_allow_subnet_size" json:"lan_allow_subnet_size"`
	Match               types.String                                              `tfsdk:"match" json:"match"`
	Name                types.String                                              `tfsdk:"name" json:"name"`
	Precedence          types.Float64                                             `tfsdk:"precedence" json:"precedence"`
	SupportURL          types.String                                              `tfsdk:"support_url" json:"support_url"`
	SwitchLocked        types.Bool                                                `tfsdk:"switch_locked" json:"switch_locked"`
	TunnelProtocol      types.String                                              `tfsdk:"tunnel_protocol" json:"tunnel_protocol"`
	Exclude             *[]*ZeroTrustDeviceProfilesExcludeDataSourceModel         `tfsdk:"exclude" json:"exclude"`
	FallbackDomains     *[]*ZeroTrustDeviceProfilesFallbackDomainsDataSourceModel `tfsdk:"fallback_domains" json:"fallback_domains"`
	Include             *[]*ZeroTrustDeviceProfilesIncludeDataSourceModel         `tfsdk:"include" json:"include"`
	ServiceModeV2       *ZeroTrustDeviceProfilesServiceModeV2DataSourceModel      `tfsdk:"service_mode_v2" json:"service_mode_v2"`
	TargetTests         *[]*ZeroTrustDeviceProfilesTargetTestsDataSourceModel     `tfsdk:"target_tests" json:"target_tests"`
	Filter              *ZeroTrustDeviceProfilesFindOneByDataSourceModel          `tfsdk:"filter"`
}

func (m *ZeroTrustDeviceProfilesDataSourceModel) toReadParams() (params zero_trust.DevicePolicyGetParams, diags diag.Diagnostics) {
	params = zero_trust.DevicePolicyGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ZeroTrustDeviceProfilesDataSourceModel) toListParams() (params zero_trust.DevicePolicyListParams, diags diag.Diagnostics) {
	params = zero_trust.DevicePolicyListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type ZeroTrustDeviceProfilesExcludeDataSourceModel struct {
	Address     types.String `tfsdk:"address" json:"address,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Host        types.String `tfsdk:"host" json:"host"`
}

type ZeroTrustDeviceProfilesFallbackDomainsDataSourceModel struct {
	Suffix      types.String    `tfsdk:"suffix" json:"suffix,computed"`
	Description types.String    `tfsdk:"description" json:"description"`
	DNSServer   *[]types.String `tfsdk:"dns_server" json:"dns_server"`
}

type ZeroTrustDeviceProfilesIncludeDataSourceModel struct {
	Address     types.String `tfsdk:"address" json:"address,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Host        types.String `tfsdk:"host" json:"host"`
}

type ZeroTrustDeviceProfilesServiceModeV2DataSourceModel struct {
	Mode types.String  `tfsdk:"mode" json:"mode"`
	Port types.Float64 `tfsdk:"port" json:"port"`
}

type ZeroTrustDeviceProfilesTargetTestsDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id"`
	Name types.String `tfsdk:"name" json:"name"`
}

type ZeroTrustDeviceProfilesFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
