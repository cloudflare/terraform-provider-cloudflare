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
	ID                         types.String                                                                   `tfsdk:"id" json:"-,computed"`
	PolicyID                   types.String                                                                   `tfsdk:"policy_id" json:"policy_id,computed"`
	AccountID                  types.String                                                                   `tfsdk:"account_id" path:"account_id,required"`
	Match                      types.String                                                                   `tfsdk:"match" json:"match,required"`
	Name                       types.String                                                                   `tfsdk:"name" json:"name,required"`
	Precedence                 types.Float64                                                                  `tfsdk:"precedence" json:"precedence,required"`
	Description                types.String                                                                   `tfsdk:"description" json:"description,optional"`
	Enabled                    types.Bool                                                                     `tfsdk:"enabled" json:"enabled,optional"`
	LANAllowMinutes            types.Float64                                                                  `tfsdk:"lan_allow_minutes" json:"lan_allow_minutes,optional"`
	LANAllowSubnetSize         types.Float64                                                                  `tfsdk:"lan_allow_subnet_size" json:"lan_allow_subnet_size,optional"`
	Exclude                    *[]*ZeroTrustDeviceCustomProfileExcludeModel                                   `tfsdk:"exclude" json:"exclude,optional"`
	Include                    *[]*ZeroTrustDeviceCustomProfileIncludeModel                                   `tfsdk:"include" json:"include,optional"`
	ServiceModeV2              *ZeroTrustDeviceCustomProfileServiceModeV2Model                                `tfsdk:"service_mode_v2" json:"service_mode_v2,optional"`
	AllowModeSwitch            types.Bool                                                                     `tfsdk:"allow_mode_switch" json:"allow_mode_switch,computed_optional"`
	AllowUpdates               types.Bool                                                                     `tfsdk:"allow_updates" json:"allow_updates,computed_optional"`
	AllowedToLeave             types.Bool                                                                     `tfsdk:"allowed_to_leave" json:"allowed_to_leave,computed_optional"`
	AutoConnect                types.Float64                                                                  `tfsdk:"auto_connect" json:"auto_connect,computed_optional"`
	CaptivePortal              types.Float64                                                                  `tfsdk:"captive_portal" json:"captive_portal,computed_optional"`
	DisableAutoFallback        types.Bool                                                                     `tfsdk:"disable_auto_fallback" json:"disable_auto_fallback,computed_optional"`
	ExcludeOfficeIPs           types.Bool                                                                     `tfsdk:"exclude_office_ips" json:"exclude_office_ips,computed_optional"`
	RegisterInterfaceIPWithDNS types.Bool                                                                     `tfsdk:"register_interface_ip_with_dns" json:"register_interface_ip_with_dns,computed_optional"`
	SccmVpnBoundarySupport     types.Bool                                                                     `tfsdk:"sccm_vpn_boundary_support" json:"sccm_vpn_boundary_support,computed_optional"`
	SupportURL                 types.String                                                                   `tfsdk:"support_url" json:"support_url,computed_optional"`
	SwitchLocked               types.Bool                                                                     `tfsdk:"switch_locked" json:"switch_locked,computed_optional"`
	TunnelProtocol             types.String                                                                   `tfsdk:"tunnel_protocol" json:"tunnel_protocol,computed_optional"`
	Default                    types.Bool                                                                     `tfsdk:"default" json:"default,computed"`
	GatewayUniqueID            types.String                                                                   `tfsdk:"gateway_unique_id" json:"gateway_unique_id,computed"`
	FallbackDomains            customfield.NestedObjectList[ZeroTrustDeviceCustomProfileFallbackDomainsModel] `tfsdk:"fallback_domains" json:"fallback_domains,computed"`
	TargetTests                customfield.NestedObjectList[ZeroTrustDeviceCustomProfileTargetTestsModel]     `tfsdk:"target_tests" json:"target_tests,computed"`
}

func (m ZeroTrustDeviceCustomProfileModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustDeviceCustomProfileModel) MarshalJSONForUpdate(state ZeroTrustDeviceCustomProfileModel) (data []byte, err error) {
	// For computed_optional fields that have default values from the server,
	// we need to avoid sending null values when the user hasn't explicitly set them.
	// This prevents API errors like "bad device request" (code 2004).

	// For pure optional fields (like description), when user sets them to null
	// we should NOT send them to the API at all, rather than sending null,
	// because some API endpoints don't accept null for these fields.

	// Create a copy of the plan model to modify
	planCopy := m

	// Special handling for description field:
	// If user wants to remove description (plan is null but state has value),
	// we send empty string instead of null because API doesn't accept null.
	if m.Description.IsNull() && !state.Description.IsNull() {
		planCopy.Description = types.StringValue("") // Send empty string to remove description
	}

	// Handle other pure optional fields similarly
	if m.Enabled.IsNull() && !state.Enabled.IsNull() {
		planCopy.Enabled = state.Enabled
	}
	if m.LANAllowMinutes.IsNull() && !state.LANAllowMinutes.IsNull() {
		planCopy.LANAllowMinutes = state.LANAllowMinutes
	}
	if m.LANAllowSubnetSize.IsNull() && !state.LANAllowSubnetSize.IsNull() {
		planCopy.LANAllowSubnetSize = state.LANAllowSubnetSize
	}

	// For each computed_optional field that is null in plan but has a value in state,
	// preserve the state value to avoid nullifying server defaults
	if m.AllowModeSwitch.IsNull() && !state.AllowModeSwitch.IsNull() {
		planCopy.AllowModeSwitch = state.AllowModeSwitch
	}
	if m.AllowUpdates.IsNull() && !state.AllowUpdates.IsNull() {
		planCopy.AllowUpdates = state.AllowUpdates
	}
	if m.AllowedToLeave.IsNull() && !state.AllowedToLeave.IsNull() {
		planCopy.AllowedToLeave = state.AllowedToLeave
	}
	if m.AutoConnect.IsNull() && !state.AutoConnect.IsNull() {
		planCopy.AutoConnect = state.AutoConnect
	}
	if m.CaptivePortal.IsNull() && !state.CaptivePortal.IsNull() {
		planCopy.CaptivePortal = state.CaptivePortal
	}
	if m.DisableAutoFallback.IsNull() && !state.DisableAutoFallback.IsNull() {
		planCopy.DisableAutoFallback = state.DisableAutoFallback
	}
	if m.ExcludeOfficeIPs.IsNull() && !state.ExcludeOfficeIPs.IsNull() {
		planCopy.ExcludeOfficeIPs = state.ExcludeOfficeIPs
	}
	if m.RegisterInterfaceIPWithDNS.IsNull() && !state.RegisterInterfaceIPWithDNS.IsNull() {
		planCopy.RegisterInterfaceIPWithDNS = state.RegisterInterfaceIPWithDNS
	}
	if m.SccmVpnBoundarySupport.IsNull() && !state.SccmVpnBoundarySupport.IsNull() {
		planCopy.SccmVpnBoundarySupport = state.SccmVpnBoundarySupport
	}
	if m.SupportURL.IsNull() && !state.SupportURL.IsNull() {
		planCopy.SupportURL = state.SupportURL
	}
	if m.SwitchLocked.IsNull() && !state.SwitchLocked.IsNull() {
		planCopy.SwitchLocked = state.SwitchLocked
	}
	if m.TunnelProtocol.IsNull() && !state.TunnelProtocol.IsNull() {
		planCopy.TunnelProtocol = state.TunnelProtocol
	} // For optional fields that are lists/objects, handle them carefully
	if m.Exclude == nil && state.Exclude != nil {
		planCopy.Exclude = state.Exclude
	}
	if m.Include == nil && state.Include != nil {
		planCopy.Include = state.Include
	}
	if m.ServiceModeV2 == nil && state.ServiceModeV2 != nil {
		planCopy.ServiceModeV2 = state.ServiceModeV2
	}

	// Computed-only fields should always preserve state values
	if !state.FallbackDomains.IsNull() {
		planCopy.FallbackDomains = state.FallbackDomains
	}
	if !state.Default.IsNull() {
		planCopy.Default = state.Default
	}
	if !state.GatewayUniqueID.IsNull() {
		planCopy.GatewayUniqueID = state.GatewayUniqueID
	}
	if !state.TargetTests.IsNull() {
		planCopy.TargetTests = state.TargetTests
	}

	result, err := apijson.MarshalForPatch(planCopy, state)
	return result, err
}

type ZeroTrustDeviceCustomProfileExcludeModel struct {
	Address     types.String `tfsdk:"address" json:"address,optional"`
	Description types.String `tfsdk:"description" json:"description,optional"`
	Host        types.String `tfsdk:"host" json:"host,optional"`
}

type ZeroTrustDeviceCustomProfileIncludeModel struct {
	Address     types.String `tfsdk:"address" json:"address,optional"`
	Description types.String `tfsdk:"description" json:"description,optional"`
	Host        types.String `tfsdk:"host" json:"host,optional"`
}

type ZeroTrustDeviceCustomProfileServiceModeV2Model struct {
	Mode types.String  `tfsdk:"mode" json:"mode,optional"`
	Port types.Float64 `tfsdk:"port" json:"port,optional"`
}

type ZeroTrustDeviceCustomProfileFallbackDomainsModel struct {
	Suffix      types.String                   `tfsdk:"suffix" json:"suffix,computed"`
	Description types.String                   `tfsdk:"description" json:"description,computed"`
	DNSServer   customfield.List[types.String] `tfsdk:"dns_server" json:"dns_server,computed"`
}

type ZeroTrustDeviceCustomProfileTargetTestsModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}
