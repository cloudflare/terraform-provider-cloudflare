// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_local_domain_fallback

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustLocalDomainFallbackResultEnvelope struct {
	Result ZeroTrustLocalDomainFallbackModel `json:"result"`
}

type ZeroTrustLocalDomainFallbackModel struct {
	ID          types.String                   `tfsdk:"id" json:"-,computed"`
	PolicyID    types.String                   `tfsdk:"policy_id" path:"policy_id"`
	AccountID   types.String                   `tfsdk:"account_id" path:"account_id"`
	Suffix      types.String                   `tfsdk:"suffix" json:"suffix"`
	Description types.String                   `tfsdk:"description" json:"description,computed_optional"`
	DNSServer   customfield.List[types.String] `tfsdk:"dns_server" json:"dns_server,computed_optional"`
}
