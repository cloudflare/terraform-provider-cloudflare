// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_domain

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PagesDomainsResultListDataSourceEnvelope struct {
	Result *[]*PagesDomainsResultDataSourceModel `json:"result,computed"`
}

type PagesDomainsDataSourceModel struct {
	AccountID   types.String                          `tfsdk:"account_id" path:"account_id"`
	ProjectName types.String                          `tfsdk:"project_name" path:"project_name"`
	MaxItems    types.Int64                           `tfsdk:"max_items"`
	Result      *[]*PagesDomainsResultDataSourceModel `tfsdk:"result"`
}

type PagesDomainsResultDataSourceModel struct {
}
