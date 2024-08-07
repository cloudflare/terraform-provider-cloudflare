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
	MaxItems  types.Int64                        `tfsdk:"max_items"`
	Result    *[]*APITokensResultDataSourceModel `tfsdk:"result"`
}

type APITokensResultDataSourceModel struct {
}
