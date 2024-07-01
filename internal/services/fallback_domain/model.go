// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package fallback_domain

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FallbackDomainResultEnvelope struct {
	Result FallbackDomainModel `json:"result,computed"`
}

type FallbackDomainResultDataSourceEnvelope struct {
	Result FallbackDomainDataSourceModel `json:"result,computed"`
}

type FallbackDomainsResultDataSourceEnvelope struct {
	Result FallbackDomainsDataSourceModel `json:"result,computed"`
}

type FallbackDomainModel struct {
	AccountID   types.String    `tfsdk:"account_id" path:"account_id"`
	PolicyID    types.String    `tfsdk:"policy_id" path:"policy_id"`
	Suffix      types.String    `tfsdk:"suffix" json:"suffix"`
	Description types.String    `tfsdk:"description" json:"description"`
	DNSServer   *[]types.String `tfsdk:"dns_server" json:"dns_server"`
}

type FallbackDomainDataSourceModel struct {
}

type FallbackDomainsDataSourceModel struct {
}
