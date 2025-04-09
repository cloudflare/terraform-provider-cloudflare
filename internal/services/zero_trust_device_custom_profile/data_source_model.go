// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_custom_profile

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/zero_trust"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDeviceCustomProfileResultDataSourceEnvelope struct {
Result ZeroTrustDeviceCustomProfileDataSourceModel `json:"result,computed"`
}

type ZeroTrustDeviceCustomProfileDataSourceModel struct {
ID types.String `tfsdk:"id" path:"policy_id,computed"`
PolicyID types.String `tfsdk:"policy_id" path:"policy_id,computed_optional"`
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
AllowModeSwitch types.Bool `tfsdk:"allow_mode_switch" json:"allow_mode_switch,computed"`
AllowUpdates types.Bool `tfsdk:"allow_updates" json:"allow_updates,computed"`
AllowedToLeave types.Bool `tfsdk:"allowed_to_leave" json:"allowed_to_leave,computed"`
AutoConnect types.Float64 `tfsdk:"auto_connect" json:"auto_connect,computed"`
CaptivePortal types.Float64 `tfsdk:"captive_portal" json:"captive_portal,computed"`
Default types.Bool `tfsdk:"default" json:"default,computed"`
Description types.String `tfsdk:"description" json:"description,computed"`
DisableAutoFallback types.Bool `tfsdk:"disable_auto_fallback" json:"disable_auto_fallback,computed"`
Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed"`
ExcludeOfficeIPs types.Bool `tfsdk:"exclude_office_ips" json:"exclude_office_ips,computed"`
GatewayUniqueID types.String `tfsdk:"gateway_unique_id" json:"gateway_unique_id,computed"`
LANAllowMinutes types.Float64 `tfsdk:"lan_allow_minutes" json:"lan_allow_minutes,computed"`
LANAllowSubnetSize types.Float64 `tfsdk:"lan_allow_subnet_size" json:"lan_allow_subnet_size,computed"`
Match types.String `tfsdk:"match" json:"match,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
Precedence types.Float64 `tfsdk:"precedence" json:"precedence,computed"`
RegisterInterfaceIPWithDNS types.Bool `tfsdk:"register_interface_ip_with_dns" json:"register_interface_ip_with_dns,computed"`
SupportURL types.String `tfsdk:"support_url" json:"support_url,computed"`
SwitchLocked types.Bool `tfsdk:"switch_locked" json:"switch_locked,computed"`
TunnelProtocol types.String `tfsdk:"tunnel_protocol" json:"tunnel_protocol,computed"`
Exclude customfield.NestedObjectList[ZeroTrustDeviceCustomProfileExcludeDataSourceModel] `tfsdk:"exclude" json:"exclude,computed"`
FallbackDomains customfield.NestedObjectList[ZeroTrustDeviceCustomProfileFallbackDomainsDataSourceModel] `tfsdk:"fallback_domains" json:"fallback_domains,computed"`
Include customfield.NestedObjectList[ZeroTrustDeviceCustomProfileIncludeDataSourceModel] `tfsdk:"include" json:"include,computed"`
ServiceModeV2 customfield.NestedObject[ZeroTrustDeviceCustomProfileServiceModeV2DataSourceModel] `tfsdk:"service_mode_v2" json:"service_mode_v2,computed"`
TargetTests customfield.NestedObjectList[ZeroTrustDeviceCustomProfileTargetTestsDataSourceModel] `tfsdk:"target_tests" json:"target_tests,computed"`
}

func (m *ZeroTrustDeviceCustomProfileDataSourceModel) toReadParams(_ context.Context) (params zero_trust.DevicePolicyCustomGetParams, diags diag.Diagnostics) {
  params = zero_trust.DevicePolicyCustomGetParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  return
}

type ZeroTrustDeviceCustomProfileExcludeDataSourceModel struct {
Address types.String `tfsdk:"address" json:"address,computed"`
Description types.String `tfsdk:"description" json:"description,computed"`
Host types.String `tfsdk:"host" json:"host,computed"`
}

type ZeroTrustDeviceCustomProfileFallbackDomainsDataSourceModel struct {
Suffix types.String `tfsdk:"suffix" json:"suffix,computed"`
Description types.String `tfsdk:"description" json:"description,computed"`
DNSServer customfield.List[types.String] `tfsdk:"dns_server" json:"dns_server,computed"`
}

type ZeroTrustDeviceCustomProfileIncludeDataSourceModel struct {
Address types.String `tfsdk:"address" json:"address,computed"`
Description types.String `tfsdk:"description" json:"description,computed"`
Host types.String `tfsdk:"host" json:"host,computed"`
}

type ZeroTrustDeviceCustomProfileServiceModeV2DataSourceModel struct {
Mode types.String `tfsdk:"mode" json:"mode,computed"`
Port types.Float64 `tfsdk:"port" json:"port,computed"`
}

type ZeroTrustDeviceCustomProfileTargetTestsDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
}
