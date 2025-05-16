// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_default_profile

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDeviceDefaultProfileResultDataSourceEnvelope struct {
	Result ZeroTrustDeviceDefaultProfileDataSourceModel `json:"result,computed"`
}

type ZeroTrustDeviceDefaultProfileDataSourceModel struct {
	AccountID                  types.String                                                                              `tfsdk:"account_id" path:"account_id,required"`
	AllowModeSwitch            types.Bool                                                                                `tfsdk:"allow_mode_switch" json:"allow_mode_switch,computed"`
	AllowUpdates               types.Bool                                                                                `tfsdk:"allow_updates" json:"allow_updates,computed"`
	AllowedToLeave             types.Bool                                                                                `tfsdk:"allowed_to_leave" json:"allowed_to_leave,computed"`
	AutoConnect                types.Float64                                                                             `tfsdk:"auto_connect" json:"auto_connect,computed"`
	CaptivePortal              types.Float64                                                                             `tfsdk:"captive_portal" json:"captive_portal,computed"`
	Default                    types.Bool                                                                                `tfsdk:"default" json:"default,computed"`
	DisableAutoFallback        types.Bool                                                                                `tfsdk:"disable_auto_fallback" json:"disable_auto_fallback,computed"`
	Enabled                    types.Bool                                                                                `tfsdk:"enabled" json:"enabled,computed"`
	ExcludeOfficeIPs           types.Bool                                                                                `tfsdk:"exclude_office_ips" json:"exclude_office_ips,computed"`
	GatewayUniqueID            types.String                                                                              `tfsdk:"gateway_unique_id" json:"gateway_unique_id,computed"`
	RegisterInterfaceIPWithDNS types.Bool                                                                                `tfsdk:"register_interface_ip_with_dns" json:"register_interface_ip_with_dns,computed"`
	SccmVpnBoundarySupport     types.Bool                                                                                `tfsdk:"sccm_vpn_boundary_support" json:"sccm_vpn_boundary_support,computed"`
	SupportURL                 types.String                                                                              `tfsdk:"support_url" json:"support_url,computed"`
	SwitchLocked               types.Bool                                                                                `tfsdk:"switch_locked" json:"switch_locked,computed"`
	TunnelProtocol             types.String                                                                              `tfsdk:"tunnel_protocol" json:"tunnel_protocol,computed"`
	Exclude                    customfield.NestedObjectList[ZeroTrustDeviceDefaultProfileExcludeDataSourceModel]         `tfsdk:"exclude" json:"exclude,computed"`
	FallbackDomains            customfield.NestedObjectList[ZeroTrustDeviceDefaultProfileFallbackDomainsDataSourceModel] `tfsdk:"fallback_domains" json:"fallback_domains,computed"`
	Include                    customfield.NestedObjectList[ZeroTrustDeviceDefaultProfileIncludeDataSourceModel]         `tfsdk:"include" json:"include,computed"`
	ServiceModeV2              customfield.NestedObject[ZeroTrustDeviceDefaultProfileServiceModeV2DataSourceModel]       `tfsdk:"service_mode_v2" json:"service_mode_v2,computed"`
}

func (m *ZeroTrustDeviceDefaultProfileDataSourceModel) toReadParams(_ context.Context) (params zero_trust.DevicePolicyDefaultGetParams, diags diag.Diagnostics) {
	params = zero_trust.DevicePolicyDefaultGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustDeviceDefaultProfileExcludeDataSourceModel struct {
	Address     types.String `tfsdk:"address" json:"address,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Host        types.String `tfsdk:"host" json:"host,computed"`
}

type ZeroTrustDeviceDefaultProfileFallbackDomainsDataSourceModel struct {
	Suffix      types.String                   `tfsdk:"suffix" json:"suffix,computed"`
	Description types.String                   `tfsdk:"description" json:"description,computed"`
	DNSServer   customfield.List[types.String] `tfsdk:"dns_server" json:"dns_server,computed"`
}

type ZeroTrustDeviceDefaultProfileIncludeDataSourceModel struct {
	Address     types.String `tfsdk:"address" json:"address,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Host        types.String `tfsdk:"host" json:"host,computed"`
}

type ZeroTrustDeviceDefaultProfileServiceModeV2DataSourceModel struct {
	Mode types.String  `tfsdk:"mode" json:"mode,computed"`
	Port types.Float64 `tfsdk:"port" json:"port,computed"`
}
