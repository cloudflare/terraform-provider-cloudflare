// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_local_domain_fallback

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework/diag"
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

func (m *ZeroTrustLocalDomainFallbacksDataSourceModel) toListParams() (params zero_trust.DevicePolicyFallbackDomainListParams, diags diag.Diagnostics) {
	params = zero_trust.DevicePolicyFallbackDomainListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustLocalDomainFallbacksResultDataSourceModel struct {
	Suffix      types.String    `tfsdk:"suffix" json:"suffix,computed"`
	Description types.String    `tfsdk:"description" json:"description"`
	DNSServer   *[]types.String `tfsdk:"dns_server" json:"dns_server"`
}
