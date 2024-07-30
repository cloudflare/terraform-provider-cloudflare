// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountResultDataSourceEnvelope struct {
	Result AccountDataSourceModel `json:"result,computed"`
}

type AccountResultListDataSourceEnvelope struct {
	Result *[]*AccountDataSourceModel `json:"result,computed"`
}

type AccountDataSourceModel struct {
	AccountID jsontypes.Normalized             `tfsdk:"account_id" path:"account_id"`
	Filter    *AccountFindOneByDataSourceModel `tfsdk:"filter"`
}

type AccountFindOneByDataSourceModel struct {
	Direction types.String  `tfsdk:"direction" query:"direction"`
	Name      types.String  `tfsdk:"name" query:"name"`
	Page      types.Float64 `tfsdk:"page" query:"page"`
	PerPage   types.Float64 `tfsdk:"per_page" query:"per_page"`
}
