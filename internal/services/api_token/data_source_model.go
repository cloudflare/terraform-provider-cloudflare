// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_token

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APITokenResultDataSourceEnvelope struct {
	Result APITokenDataSourceModel `json:"result,computed"`
}

type APITokenResultListDataSourceEnvelope struct {
	Result *[]*APITokenDataSourceModel `json:"result,computed"`
}

type APITokenDataSourceModel struct {
	TokenID   jsontypes.Normalized              `tfsdk:"token_id" path:"token_id"`
	FindOneBy *APITokenFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

type APITokenFindOneByDataSourceModel struct {
	Direction types.String  `tfsdk:"direction" query:"direction"`
	Page      types.Float64 `tfsdk:"page" query:"page"`
	PerPage   types.Float64 `tfsdk:"per_page" query:"per_page"`
}
