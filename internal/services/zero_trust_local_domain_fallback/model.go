// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_local_domain_fallback

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustLocalDomainFallbackResultEnvelope struct {
	Result *[]*ZeroTrustLocalDomainFallbackDomainsModel `json:"result"`
}

type ZeroTrustLocalDomainFallbackModel struct {
	ID        types.String                                 `tfsdk:"id" json:"-,computed"`
	PolicyID  types.String                                 `tfsdk:"policy_id" path:"policy_id,required"`
	AccountID types.String                                 `tfsdk:"account_id" path:"account_id,required"`
	Domains   *[]*ZeroTrustLocalDomainFallbackDomainsModel `tfsdk:"domains" json:"domains,required"`
}

func (m ZeroTrustLocalDomainFallbackModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m.Domains)
}

func (m ZeroTrustLocalDomainFallbackModel) MarshalJSONForUpdate(state ZeroTrustLocalDomainFallbackModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m.Domains, state.Domains)
}

type ZeroTrustLocalDomainFallbackDomainsModel struct {
	Suffix      types.String    `tfsdk:"suffix" json:"suffix,required"`
	Description types.String    `tfsdk:"description" json:"description,optional"`
	DNSServer   *[]types.String `tfsdk:"dns_server" json:"dns_server,optional"`
}
