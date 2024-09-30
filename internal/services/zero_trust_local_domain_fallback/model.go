// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_local_domain_fallback

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustLocalDomainFallbackResultEnvelope struct {
	Result ZeroTrustLocalDomainFallbackModel `json:"result"`
}

type ZeroTrustLocalDomainFallbackModel struct {
	ID          types.String    `tfsdk:"id" json:"-,computed"`
	PolicyID    types.String    `tfsdk:"policy_id" path:"policy_id,required"`
	AccountID   types.String    `tfsdk:"account_id" path:"account_id,required"`
	Suffix      types.String    `tfsdk:"suffix" json:"suffix,required"`
	Description types.String    `tfsdk:"description" json:"description,optional"`
	DNSServer   *[]types.String `tfsdk:"dns_server" json:"dns_server,optional"`
}
