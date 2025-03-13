// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_custom_profile_local_domain_fallback

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/zero_trust"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDeviceCustomProfileLocalDomainFallbackResultDataSourceEnvelope struct {
Result ZeroTrustDeviceCustomProfileLocalDomainFallbackDataSourceModel `json:"result,computed"`
}

type ZeroTrustDeviceCustomProfileLocalDomainFallbackDataSourceModel struct {
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
PolicyID types.String `tfsdk:"policy_id" path:"policy_id,required"`
Description types.String `tfsdk:"description" json:"description,computed"`
Suffix types.String `tfsdk:"suffix" json:"suffix,computed"`
DNSServer customfield.List[types.String] `tfsdk:"dns_server" json:"dns_server,computed"`
}

func (m *ZeroTrustDeviceCustomProfileLocalDomainFallbackDataSourceModel) toReadParams(_ context.Context) (params zero_trust.DevicePolicyCustomFallbackDomainGetParams, diags diag.Diagnostics) {
  params = zero_trust.DevicePolicyCustomFallbackDomainGetParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  return
}
