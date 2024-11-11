// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_custom_profile

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDeviceCustomProfileResultEnvelope struct {
	Result ZeroTrustDeviceCustomProfileModel `json:"result"`
}

type ZeroTrustDeviceCustomProfileModel struct {
	AccountID           types.String                                                                   `tfsdk:"account_id" path:"account_id,required"`
	PolicyID            types.String                                                                   `tfsdk:"policy_id" path:"policy_id,optional"`
	LANAllowMinutes     types.Float64                                                                  `tfsdk:"lan_allow_minutes" json:"lan_allow_minutes,optional"`
	LANAllowSubnetSize  types.Float64                                                                  `tfsdk:"lan_allow_subnet_size" json:"lan_allow_subnet_size,optional"`
	Match               types.String                                                                   `tfsdk:"match" json:"match,required"`
	Name                types.String                                                                   `tfsdk:"name" json:"name,required"`
	Precedence          types.Float64                                                                  `tfsdk:"precedence" json:"precedence,required"`
	AllowModeSwitch     types.Bool                                                                     `tfsdk:"allow_mode_switch" json:"allow_mode_switch,optional"`
	AllowUpdates        types.Bool                                                                     `tfsdk:"allow_updates" json:"allow_updates,optional"`
	AllowedToLeave      types.Bool                                                                     `tfsdk:"allowed_to_leave" json:"allowed_to_leave,optional"`
	AutoConnect         types.Float64                                                                  `tfsdk:"auto_connect" json:"auto_connect,optional"`
	CaptivePortal       types.Float64                                                                  `tfsdk:"captive_portal" json:"captive_portal,optional"`
	Description         types.String                                                                   `tfsdk:"description" json:"description,optional"`
	DisableAutoFallback types.Bool                                                                     `tfsdk:"disable_auto_fallback" json:"disable_auto_fallback,optional"`
	Enabled             types.Bool                                                                     `tfsdk:"enabled" json:"enabled,optional"`
	ExcludeOfficeIPs    types.Bool                                                                     `tfsdk:"exclude_office_ips" json:"exclude_office_ips,optional"`
	SupportURL          types.String                                                                   `tfsdk:"support_url" json:"support_url,optional"`
	SwitchLocked        types.Bool                                                                     `tfsdk:"switch_locked" json:"switch_locked,optional"`
	TunnelProtocol      types.String                                                                   `tfsdk:"tunnel_protocol" json:"tunnel_protocol,optional"`
	ServiceModeV2       customfield.NestedObject[ZeroTrustDeviceCustomProfileServiceModeV2Model]       `tfsdk:"service_mode_v2" json:"service_mode_v2,computed_optional"`
	Default             types.Bool                                                                     `tfsdk:"default" json:"default,computed"`
	GatewayUniqueID     types.String                                                                   `tfsdk:"gateway_unique_id" json:"gateway_unique_id,computed"`
	Exclude             customfield.NestedObjectList[ZeroTrustDeviceCustomProfileExcludeModel]         `tfsdk:"exclude" json:"exclude,computed"`
	FallbackDomains     customfield.NestedObjectList[ZeroTrustDeviceCustomProfileFallbackDomainsModel] `tfsdk:"fallback_domains" json:"fallback_domains,computed"`
	Include             customfield.NestedObjectList[ZeroTrustDeviceCustomProfileIncludeModel]         `tfsdk:"include" json:"include,computed"`
	TargetTests         customfield.NestedObjectList[ZeroTrustDeviceCustomProfileTargetTestsModel]     `tfsdk:"target_tests" json:"target_tests,computed"`
}

func (m ZeroTrustDeviceCustomProfileModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustDeviceCustomProfileModel) MarshalJSONForUpdate(state ZeroTrustDeviceCustomProfileModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustDeviceCustomProfileServiceModeV2Model struct {
	Mode types.String  `tfsdk:"mode" json:"mode,optional"`
	Port types.Float64 `tfsdk:"port" json:"port,optional"`
}

type ZeroTrustDeviceCustomProfileExcludeModel struct {
	Address     types.String `tfsdk:"address" json:"address,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Host        types.String `tfsdk:"host" json:"host,computed"`
}

type ZeroTrustDeviceCustomProfileFallbackDomainsModel struct {
	Suffix      types.String                   `tfsdk:"suffix" json:"suffix,computed"`
	Description types.String                   `tfsdk:"description" json:"description,computed"`
	DNSServer   customfield.List[types.String] `tfsdk:"dns_server" json:"dns_server,computed"`
}

type ZeroTrustDeviceCustomProfileIncludeModel struct {
	Address     types.String `tfsdk:"address" json:"address,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Host        types.String `tfsdk:"host" json:"host,computed"`
}

type ZeroTrustDeviceCustomProfileTargetTestsModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}
