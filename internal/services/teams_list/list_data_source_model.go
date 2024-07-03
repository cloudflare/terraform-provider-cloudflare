// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package teams_list

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TeamsListsResultListDataSourceEnvelope struct {
	Result *[]*TeamsListsItemsDataSourceModel `json:"result,computed"`
}

type TeamsListsDataSourceModel struct {
	AccountID types.String                       `tfsdk:"account_id" path:"account_id"`
	Type      types.String                       `tfsdk:"type" query:"type"`
	MaxItems  types.Int64                        `tfsdk:"max_items"`
	Items     *[]*TeamsListsItemsDataSourceModel `tfsdk:"items"`
}

type TeamsListsItemsDataSourceModel struct {
	ID          types.String  `tfsdk:"id" json:"id,computed"`
	Count       types.Float64 `tfsdk:"count" json:"count,computed"`
	CreatedAt   types.String  `tfsdk:"created_at" json:"created_at,computed"`
	Description types.String  `tfsdk:"description" json:"description,computed"`
	Name        types.String  `tfsdk:"name" json:"name,computed"`
	Type        types.String  `tfsdk:"type" json:"type,computed"`
	UpdatedAt   types.String  `tfsdk:"updated_at" json:"updated_at,computed"`
}
