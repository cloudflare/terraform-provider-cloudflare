// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_custom_profile_local_domain_fallback

import (
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDeviceCustomProfileLocalDomainFallbackResultEnvelope struct {
Result *[]*ZeroTrustDeviceCustomProfileLocalDomainFallbackDomainsModel `json:"result"`
}

type ZeroTrustDeviceCustomProfileLocalDomainFallbackModel struct {
ID types.String `tfsdk:"id" json:"-,computed"`
PolicyID types.String `tfsdk:"policy_id" path:"policy_id,required"`
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
Domains *[]*ZeroTrustDeviceCustomProfileLocalDomainFallbackDomainsModel `tfsdk:"domains" json:"domains,required"`
Description types.String `tfsdk:"description" json:"description,computed"`
Suffix types.String `tfsdk:"suffix" json:"suffix,computed"`
DNSServer customfield.List[types.String] `tfsdk:"dns_server" json:"dns_server,computed"`
}

func (m ZeroTrustDeviceCustomProfileLocalDomainFallbackModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m.Domains)
}

func (m ZeroTrustDeviceCustomProfileLocalDomainFallbackModel) MarshalJSONForUpdate(state ZeroTrustDeviceCustomProfileLocalDomainFallbackModel) (data []byte, err error) {
  return apijson.MarshalForUpdate(m.Domains, state.Domains)
}

type ZeroTrustDeviceCustomProfileLocalDomainFallbackDomainsModel struct {
Suffix types.String `tfsdk:"suffix" json:"suffix,required"`
Description types.String `tfsdk:"description" json:"description,optional"`
DNSServer *[]types.String `tfsdk:"dns_server" json:"dns_server,optional"`
}
