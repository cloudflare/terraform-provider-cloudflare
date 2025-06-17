// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_default_profile

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDeviceDefaultProfileResultEnvelope struct {
	Result ZeroTrustDeviceDefaultProfileModel `json:"result"`
}

type ZeroTrustDeviceDefaultProfileModel struct {
	ID                         types.String                                                                    `tfsdk:"id" json:"-,computed"`
	AccountID                  types.String                                                                    `tfsdk:"account_id" path:"account_id,required"`
	LANAllowMinutes            types.Float64                                                                   `tfsdk:"lan_allow_minutes" json:"lan_allow_minutes,optional,no_refresh"`
	LANAllowSubnetSize         types.Float64                                                                   `tfsdk:"lan_allow_subnet_size" json:"lan_allow_subnet_size,optional,no_refresh"`
	Exclude                    *[]*ZeroTrustDeviceDefaultProfileExcludeModel                                   `tfsdk:"exclude" json:"exclude,optional"`
	Include                    *[]*ZeroTrustDeviceDefaultProfileIncludeModel                                   `tfsdk:"include" json:"include,optional"`
	ServiceModeV2              *ZeroTrustDeviceDefaultProfileServiceModeV2Model                                `tfsdk:"service_mode_v2" json:"service_mode_v2,optional"`
	AllowModeSwitch            types.Bool                                                                      `tfsdk:"allow_mode_switch" json:"allow_mode_switch,computed_optional"`
	AllowUpdates               types.Bool                                                                      `tfsdk:"allow_updates" json:"allow_updates,computed_optional"`
	AllowedToLeave             types.Bool                                                                      `tfsdk:"allowed_to_leave" json:"allowed_to_leave,computed_optional"`
	AutoConnect                types.Float64                                                                   `tfsdk:"auto_connect" json:"auto_connect,computed_optional"`
	CaptivePortal              types.Float64                                                                   `tfsdk:"captive_portal" json:"captive_portal,computed_optional"`
	DisableAutoFallback        types.Bool                                                                      `tfsdk:"disable_auto_fallback" json:"disable_auto_fallback,computed_optional"`
	ExcludeOfficeIPs           types.Bool                                                                      `tfsdk:"exclude_office_ips" json:"exclude_office_ips,computed_optional"`
	RegisterInterfaceIPWithDNS types.Bool                                                                      `tfsdk:"register_interface_ip_with_dns" json:"register_interface_ip_with_dns,computed_optional"`
	SccmVpnBoundarySupport     types.Bool                                                                      `tfsdk:"sccm_vpn_boundary_support" json:"sccm_vpn_boundary_support,computed_optional"`
	SupportURL                 types.String                                                                    `tfsdk:"support_url" json:"support_url,computed_optional"`
	SwitchLocked               types.Bool                                                                      `tfsdk:"switch_locked" json:"switch_locked,computed_optional"`
	TunnelProtocol             types.String                                                                    `tfsdk:"tunnel_protocol" json:"tunnel_protocol,computed_optional"`
	Default                    types.Bool                                                                      `tfsdk:"default" json:"default,computed"`
	Enabled                    types.Bool                                                                      `tfsdk:"enabled" json:"enabled,computed"`
	GatewayUniqueID            types.String                                                                    `tfsdk:"gateway_unique_id" json:"gateway_unique_id,computed"`
	FallbackDomains            customfield.NestedObjectList[ZeroTrustDeviceDefaultProfileFallbackDomainsModel] `tfsdk:"fallback_domains" json:"fallback_domains,computed"`
}

func (m ZeroTrustDeviceDefaultProfileModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustDeviceDefaultProfileModel) MarshalJSONForUpdate(state ZeroTrustDeviceDefaultProfileModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type ZeroTrustDeviceDefaultProfileExcludeModel struct {
	Address     types.String `tfsdk:"address" json:"address,optional"`
	Description types.String `tfsdk:"description" json:"description,optional"`
	Host        types.String `tfsdk:"host" json:"host,optional"`
}

type ZeroTrustDeviceDefaultProfileIncludeModel struct {
	Address     types.String `tfsdk:"address" json:"address,optional"`
	Description types.String `tfsdk:"description" json:"description,optional"`
	Host        types.String `tfsdk:"host" json:"host,optional"`
}

type ZeroTrustDeviceDefaultProfileServiceModeV2Model struct {
	Mode types.String  `tfsdk:"mode" json:"mode,optional"`
	Port types.Float64 `tfsdk:"port" json:"port,optional"`
}

type ZeroTrustDeviceDefaultProfileFallbackDomainsModel struct {
	Suffix      types.String                   `tfsdk:"suffix" json:"suffix,computed"`
	Description types.String                   `tfsdk:"description" json:"description,computed"`
	DNSServer   customfield.List[types.String] `tfsdk:"dns_server" json:"dns_server,computed"`
}
