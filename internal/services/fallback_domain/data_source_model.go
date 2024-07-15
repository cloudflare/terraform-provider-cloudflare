// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package fallback_domain

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FallbackDomainResultDataSourceEnvelope struct {
	Result FallbackDomainDataSourceModel `json:"result,computed"`
}

type FallbackDomainResultListDataSourceEnvelope struct {
	Result *[]*FallbackDomainDataSourceModel `json:"result,computed"`
}

type FallbackDomainDataSourceModel struct {
	AccountID   types.String                            `tfsdk:"account_id" path:"account_id"`
	PolicyID    types.String                            `tfsdk:"policy_id" path:"policy_id"`
	Suffix      types.String                            `tfsdk:"suffix" json:"suffix"`
	Description types.String                            `tfsdk:"description" json:"description"`
	DNSServer   *[]jsontypes.Normalized                 `tfsdk:"dns_server" json:"dns_server"`
	FindOneBy   *FallbackDomainFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

type FallbackDomainFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
