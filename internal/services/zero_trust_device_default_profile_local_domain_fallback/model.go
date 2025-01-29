// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_default_profile_local_domain_fallback

import (
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDeviceDefaultProfileLocalDomainFallbackResultEnvelope struct {
	Result *[]*ZeroTrustDeviceDefaultProfileLocalDomainFallbackDomainsModel `json:"result"`
}

type ZeroTrustDeviceDefaultProfileLocalDomainFallbackModel struct {
	AccountID types.String                                                     `tfsdk:"account_id" path:"account_id,required"`
	Domains   *[]*ZeroTrustDeviceDefaultProfileLocalDomainFallbackDomainsModel `tfsdk:"domains" json:"domains,required"`
}

func (m ZeroTrustDeviceDefaultProfileLocalDomainFallbackModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m.Domains)
}

func (m ZeroTrustDeviceDefaultProfileLocalDomainFallbackModel) MarshalJSONForUpdate(state ZeroTrustDeviceDefaultProfileLocalDomainFallbackModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m.Domains, state.Domains)
}

type ZeroTrustDeviceDefaultProfileLocalDomainFallbackDomainsModel struct {
	Suffix      types.String    `tfsdk:"suffix" json:"suffix,required"`
	Description types.String    `tfsdk:"description" json:"description,optional"`
	DNSServer   *[]types.String `tfsdk:"dns_server" json:"dns_server,optional"`
}
