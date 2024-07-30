// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package fallback_domain

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FallbackDomainsResultListDataSourceEnvelope struct {
	Result *[]*FallbackDomainsResultDataSourceModel `json:"result,computed"`
}

type FallbackDomainsDataSourceModel struct {
	AccountID types.String                             `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                              `tfsdk:"max_items"`
	Result    *[]*FallbackDomainsResultDataSourceModel `tfsdk:"result"`
}

type FallbackDomainsResultDataSourceModel struct {
	Suffix      types.String            `tfsdk:"suffix" json:"suffix,computed"`
	Description types.String            `tfsdk:"description" json:"description"`
	DNSServer   *[]jsontypes.Normalized `tfsdk:"dns_server" json:"dns_server"`
}
