// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_token

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APITokensResultListDataSourceEnvelope struct {
	Result *[]*APITokensResultDataSourceModel `json:"result,computed"`
}

type APITokensDataSourceModel struct {
	Direction types.String                       `tfsdk:"direction" query:"direction"`
	Page      types.Float64                      `tfsdk:"page" query:"page"`
	PerPage   types.Float64                      `tfsdk:"per_page" query:"per_page"`
	MaxItems  types.Int64                        `tfsdk:"max_items"`
	Result    *[]*APITokensResultDataSourceModel `tfsdk:"result"`
}

type APITokensResultDataSourceModel struct {
}
