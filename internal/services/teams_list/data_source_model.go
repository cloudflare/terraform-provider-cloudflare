// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package teams_list

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TeamsListResultDataSourceEnvelope struct {
	Result TeamsListDataSourceModel `json:"result,computed"`
}

type TeamsListResultListDataSourceEnvelope struct {
	Result *[]*TeamsListDataSourceModel `json:"result,computed"`
}

type TeamsListDataSourceModel struct {
	AccountID   types.String                       `tfsdk:"account_id" path:"account_id"`
	ListID      types.String                       `tfsdk:"list_id" path:"list_id"`
	ID          types.String                       `tfsdk:"id" json:"id"`
	ListCount   types.Float64                      `tfsdk:"list_count" json:"count"`
	CreatedAt   types.String                       `tfsdk:"created_at" json:"created_at"`
	Description types.String                       `tfsdk:"description" json:"description"`
	Name        types.String                       `tfsdk:"name" json:"name"`
	Type        types.String                       `tfsdk:"type" json:"type"`
	UpdatedAt   types.String                       `tfsdk:"updated_at" json:"updated_at"`
	FindOneBy   *TeamsListFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

type TeamsListFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	Type      types.String `tfsdk:"type" query:"type"`
}
