// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ListsResultListDataSourceEnvelope struct {
	Result *[]*ListsItemsDataSourceModel `json:"result,computed"`
}

type ListsDataSourceModel struct {
	AccountID types.String                  `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                   `tfsdk:"max_items"`
	Items     *[]*ListsItemsDataSourceModel `tfsdk:"items"`
}

type ListsItemsDataSourceModel struct {
	ID                    types.String  `tfsdk:"id" json:"id,computed"`
	CreatedOn             types.String  `tfsdk:"created_on" json:"created_on,computed"`
	Kind                  types.String  `tfsdk:"kind" json:"kind,computed"`
	ModifiedOn            types.String  `tfsdk:"modified_on" json:"modified_on,computed"`
	Name                  types.String  `tfsdk:"name" json:"name,computed"`
	NumItems              types.Float64 `tfsdk:"num_items" json:"num_items,computed"`
	Description           types.String  `tfsdk:"description" json:"description,computed"`
	NumReferencingFilters types.Float64 `tfsdk:"num_referencing_filters" json:"num_referencing_filters,computed"`
}
