// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ListResultDataSourceEnvelope struct {
	Result ListDataSourceModel `json:"result,computed"`
}

type ListResultListDataSourceEnvelope struct {
	Result *[]*ListDataSourceModel `json:"result,computed"`
}

type ListDataSourceModel struct {
	AccountID             types.String                  `tfsdk:"account_id" path:"account_id"`
	ListID                types.String                  `tfsdk:"list_id" path:"list_id"`
	CreatedOn             types.String                  `tfsdk:"created_on" json:"created_on"`
	Description           types.String                  `tfsdk:"description" json:"description"`
	ID                    types.String                  `tfsdk:"id" json:"id"`
	Kind                  types.String                  `tfsdk:"kind" json:"kind"`
	ModifiedOn            types.String                  `tfsdk:"modified_on" json:"modified_on"`
	Name                  types.String                  `tfsdk:"name" json:"name"`
	NumItems              types.Float64                 `tfsdk:"num_items" json:"num_items"`
	NumReferencingFilters types.Float64                 `tfsdk:"num_referencing_filters" json:"num_referencing_filters"`
	Filter                *ListFindOneByDataSourceModel `tfsdk:"filter"`
}

type ListFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
