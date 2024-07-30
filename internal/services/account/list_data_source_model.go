// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountsResultListDataSourceEnvelope struct {
	Result *[]*AccountsResultDataSourceModel `json:"result,computed"`
}

type AccountsDataSourceModel struct {
	Direction types.String                      `tfsdk:"direction" query:"direction"`
	Name      types.String                      `tfsdk:"name" query:"name"`
	Page      types.Float64                     `tfsdk:"page" query:"page"`
	PerPage   types.Float64                     `tfsdk:"per_page" query:"per_page"`
	MaxItems  types.Int64                       `tfsdk:"max_items"`
	Result    *[]*AccountsResultDataSourceModel `tfsdk:"result"`
}

type AccountsResultDataSourceModel struct {
}
