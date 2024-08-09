// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_local_domain_fallback

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustLocalDomainFallbacksResultListDataSourceEnvelope struct {
	Result *[]*ZeroTrustLocalDomainFallbacksResultDataSourceModel `json:"result,computed"`
}

type ZeroTrustLocalDomainFallbacksDataSourceModel struct {
	AccountID types.String                                           `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                                            `tfsdk:"max_items"`
	Result    *[]*ZeroTrustLocalDomainFallbacksResultDataSourceModel `tfsdk:"result"`
}

type ZeroTrustLocalDomainFallbacksResultDataSourceModel struct {
	Suffix      types.String            `tfsdk:"suffix" json:"suffix,computed"`
	Description types.String            `tfsdk:"description" json:"description"`
	DNSServer   *[]jsontypes.Normalized `tfsdk:"dns_server" json:"dns_server"`
}
