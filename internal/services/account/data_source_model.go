// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountResultDataSourceEnvelope struct {
	Result AccountDataSourceModel `json:"result,computed"`
}

type AccountResultListDataSourceEnvelope struct {
	Result *[]*AccountDataSourceModel `json:"result,computed"`
}

type AccountDataSourceModel struct {
	AccountID types.String                     `tfsdk:"account_id" path:"account_id"`
	Filter    *AccountFindOneByDataSourceModel `tfsdk:"filter"`
}

type AccountFindOneByDataSourceModel struct {
	Direction types.String `tfsdk:"direction" query:"direction"`
	Name      types.String `tfsdk:"name" query:"name"`
}
