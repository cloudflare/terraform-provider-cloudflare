// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_local_domain_fallback

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustLocalDomainFallbackResultDataSourceEnvelope struct {
	Result ZeroTrustLocalDomainFallbackDataSourceModel `json:"result,computed"`
}

type ZeroTrustLocalDomainFallbackResultListDataSourceEnvelope struct {
	Result *[]*ZeroTrustLocalDomainFallbackDataSourceModel `json:"result,computed"`
}

type ZeroTrustLocalDomainFallbackDataSourceModel struct {
	AccountID   types.String                                          `tfsdk:"account_id" path:"account_id"`
	PolicyID    types.String                                          `tfsdk:"policy_id" path:"policy_id"`
	Description types.String                                          `tfsdk:"description" json:"description"`
	Suffix      types.String                                          `tfsdk:"suffix" json:"suffix"`
	DNSServer   *[]types.String                                       `tfsdk:"dns_server" json:"dns_server"`
	Filter      *ZeroTrustLocalDomainFallbackFindOneByDataSourceModel `tfsdk:"filter"`
}

type ZeroTrustLocalDomainFallbackFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
