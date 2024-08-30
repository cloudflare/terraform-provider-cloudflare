// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_profiles

import (
	"context"

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
	AccountID           types.String                                                                        `tfsdk:"account_id" path:"account_id"`
	PolicyID            types.String                                                                        `tfsdk:"policy_id" path:"policy_id,computed_optional"`
	AllowModeSwitch     types.Bool                                                                          `tfsdk:"allow_mode_switch" json:"allow_mode_switch,computed"`
	AllowUpdates        types.Bool                                                                          `tfsdk:"allow_updates" json:"allow_updates,computed"`
	AllowedToLeave      types.Bool                                                                          `tfsdk:"allowed_to_leave" json:"allowed_to_leave,computed"`
	AutoConnect         types.Float64                                                                       `tfsdk:"auto_connect" json:"auto_connect,computed"`
	CaptivePortal       types.Float64                                                                       `tfsdk:"captive_portal" json:"captive_portal,computed"`
	Default             types.Bool                                                                          `tfsdk:"default" json:"default,computed"`
	Description         types.String                                                                        `tfsdk:"description" json:"description,computed"`
	DisableAutoFallback types.Bool                                                                          `tfsdk:"disable_auto_fallback" json:"disable_auto_fallback,computed"`
	Enabled             types.Bool                                                                          `tfsdk:"enabled" json:"enabled,computed"`
	ExcludeOfficeIPs    types.Bool                                                                          `tfsdk:"exclude_office_ips" json:"exclude_office_ips,computed"`
	GatewayUniqueID     types.String                                                                        `tfsdk:"gateway_unique_id" json:"gateway_unique_id,computed"`
	LANAllowMinutes     types.Float64                                                                       `tfsdk:"lan_allow_minutes" json:"lan_allow_minutes,computed"`
	LANAllowSubnetSize  types.Float64                                                                       `tfsdk:"lan_allow_subnet_size" json:"lan_allow_subnet_size,computed"`
	Match               types.String                                                                        `tfsdk:"match" json:"match,computed"`
	Name                types.String                                                                        `tfsdk:"name" json:"name,computed"`
	Precedence          types.Float64                                                                       `tfsdk:"precedence" json:"precedence,computed"`
	SupportURL          types.String                                                                        `tfsdk:"support_url" json:"support_url,computed"`
	SwitchLocked        types.Bool                                                                          `tfsdk:"switch_locked" json:"switch_locked,computed"`
	TunnelProtocol      types.String                                                                        `tfsdk:"tunnel_protocol" json:"tunnel_protocol,computed"`
	Exclude             customfield.NestedObjectList[ZeroTrustDeviceProfilesExcludeDataSourceModel]         `tfsdk:"exclude" json:"exclude,computed"`
	FallbackDomains     customfield.NestedObjectList[ZeroTrustDeviceProfilesFallbackDomainsDataSourceModel] `tfsdk:"fallback_domains" json:"fallback_domains,computed"`
	Include             customfield.NestedObjectList[ZeroTrustDeviceProfilesIncludeDataSourceModel]         `tfsdk:"include" json:"include,computed"`
	ServiceModeV2       customfield.NestedObject[ZeroTrustDeviceProfilesServiceModeV2DataSourceModel]       `tfsdk:"service_mode_v2" json:"service_mode_v2,computed"`
	TargetTests         customfield.NestedObjectList[ZeroTrustDeviceProfilesTargetTestsDataSourceModel]     `tfsdk:"target_tests" json:"target_tests,computed"`
	Filter              *ZeroTrustDeviceProfilesFindOneByDataSourceModel                                    `tfsdk:"filter"`
}

func (m *ZeroTrustDeviceProfilesDataSourceModel) toReadParams(_ context.Context) (params zero_trust.DevicePolicyGetParams, diags diag.Diagnostics) {
	params = zero_trust.DevicePolicyGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ZeroTrustDeviceProfilesDataSourceModel) toListParams(_ context.Context) (params zero_trust.DevicePolicyListParams, diags diag.Diagnostics) {
	params = zero_trust.DevicePolicyListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type ZeroTrustDeviceProfilesExcludeDataSourceModel struct {
	Address     types.String `tfsdk:"address" json:"address,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Host        types.String `tfsdk:"host" json:"host,computed"`
}

type ZeroTrustDeviceProfilesFallbackDomainsDataSourceModel struct {
	Suffix      types.String `tfsdk:"suffix" json:"suffix,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	DNSServer   types.List   `tfsdk:"dns_server" json:"dns_server,computed"`
}

type ZeroTrustDeviceProfilesIncludeDataSourceModel struct {
	Address     types.String `tfsdk:"address" json:"address,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Host        types.String `tfsdk:"host" json:"host,computed"`
}

type ZeroTrustDeviceProfilesServiceModeV2DataSourceModel struct {
	Mode types.String  `tfsdk:"mode" json:"mode,computed"`
	Port types.Float64 `tfsdk:"port" json:"port,computed"`
}

type ZeroTrustDeviceProfilesTargetTestsDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustDeviceProfilesFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
