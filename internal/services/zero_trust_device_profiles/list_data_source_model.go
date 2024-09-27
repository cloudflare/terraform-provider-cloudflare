// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_profiles

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDeviceProfilesListResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustDeviceProfilesListResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustDeviceProfilesListDataSourceModel struct {
	AccountID types.String                                                                   `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                                    `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustDeviceProfilesListResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustDeviceProfilesListDataSourceModel) toListParams(_ context.Context) (params zero_trust.DevicePolicyListParams, diags diag.Diagnostics) {
	params = zero_trust.DevicePolicyListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustDeviceProfilesListResultDataSourceModel struct {
	AllowModeSwitch     types.Bool                                                                              `tfsdk:"allow_mode_switch" json:"allow_mode_switch,computed"`
	AllowUpdates        types.Bool                                                                              `tfsdk:"allow_updates" json:"allow_updates,computed"`
	AllowedToLeave      types.Bool                                                                              `tfsdk:"allowed_to_leave" json:"allowed_to_leave,computed"`
	AutoConnect         types.Float64                                                                           `tfsdk:"auto_connect" json:"auto_connect,computed"`
	CaptivePortal       types.Float64                                                                           `tfsdk:"captive_portal" json:"captive_portal,computed"`
	Default             types.Bool                                                                              `tfsdk:"default" json:"default,computed"`
	Description         types.String                                                                            `tfsdk:"description" json:"description,computed"`
	DisableAutoFallback types.Bool                                                                              `tfsdk:"disable_auto_fallback" json:"disable_auto_fallback,computed"`
	Enabled             types.Bool                                                                              `tfsdk:"enabled" json:"enabled,computed"`
	Exclude             customfield.NestedObjectList[ZeroTrustDeviceProfilesListExcludeDataSourceModel]         `tfsdk:"exclude" json:"exclude,computed"`
	ExcludeOfficeIPs    types.Bool                                                                              `tfsdk:"exclude_office_ips" json:"exclude_office_ips,computed"`
	FallbackDomains     customfield.NestedObjectList[ZeroTrustDeviceProfilesListFallbackDomainsDataSourceModel] `tfsdk:"fallback_domains" json:"fallback_domains,computed"`
	GatewayUniqueID     types.String                                                                            `tfsdk:"gateway_unique_id" json:"gateway_unique_id,computed"`
	Include             customfield.NestedObjectList[ZeroTrustDeviceProfilesListIncludeDataSourceModel]         `tfsdk:"include" json:"include,computed"`
	LANAllowMinutes     types.Float64                                                                           `tfsdk:"lan_allow_minutes" json:"lan_allow_minutes,computed"`
	LANAllowSubnetSize  types.Float64                                                                           `tfsdk:"lan_allow_subnet_size" json:"lan_allow_subnet_size,computed"`
	Match               types.String                                                                            `tfsdk:"match" json:"match,computed"`
	Name                types.String                                                                            `tfsdk:"name" json:"name,computed"`
	PolicyID            types.String                                                                            `tfsdk:"policy_id" json:"policy_id,computed"`
	Precedence          types.Float64                                                                           `tfsdk:"precedence" json:"precedence,computed"`
	ServiceModeV2       customfield.NestedObject[ZeroTrustDeviceProfilesListServiceModeV2DataSourceModel]       `tfsdk:"service_mode_v2" json:"service_mode_v2,computed"`
	SupportURL          types.String                                                                            `tfsdk:"support_url" json:"support_url,computed"`
	SwitchLocked        types.Bool                                                                              `tfsdk:"switch_locked" json:"switch_locked,computed"`
	TargetTests         customfield.NestedObjectList[ZeroTrustDeviceProfilesListTargetTestsDataSourceModel]     `tfsdk:"target_tests" json:"target_tests,computed"`
	TunnelProtocol      types.String                                                                            `tfsdk:"tunnel_protocol" json:"tunnel_protocol,computed"`
}

type ZeroTrustDeviceProfilesListExcludeDataSourceModel struct {
	Address     types.String `tfsdk:"address" json:"address,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Host        types.String `tfsdk:"host" json:"host,computed"`
}

type ZeroTrustDeviceProfilesListFallbackDomainsDataSourceModel struct {
	Suffix      types.String                   `tfsdk:"suffix" json:"suffix,computed"`
	Description types.String                   `tfsdk:"description" json:"description,computed"`
	DNSServer   customfield.List[types.String] `tfsdk:"dns_server" json:"dns_server,computed"`
}

type ZeroTrustDeviceProfilesListIncludeDataSourceModel struct {
	Address     types.String `tfsdk:"address" json:"address,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Host        types.String `tfsdk:"host" json:"host,computed"`
}

type ZeroTrustDeviceProfilesListServiceModeV2DataSourceModel struct {
	Mode types.String  `tfsdk:"mode" json:"mode,computed"`
	Port types.Float64 `tfsdk:"port" json:"port,computed"`
}

type ZeroTrustDeviceProfilesListTargetTestsDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}
