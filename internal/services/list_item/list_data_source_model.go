// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list_item

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ListItemsResultListDataSourceEnvelope struct {
	Result *[]*ListItemsItemsDataSourceModel `json:"result,computed"`
}

type ListItemsDataSourceModel struct {
	AccountID types.String                      `tfsdk:"account_id" path:"account_id"`
	ListID    types.String                      `tfsdk:"list_id" path:"list_id"`
	Cursor    types.String                      `tfsdk:"cursor" query:"cursor"`
	PerPage   types.Int64                       `tfsdk:"per_page" query:"per_page"`
	Search    types.String                      `tfsdk:"search" query:"search"`
	MaxItems  types.Int64                       `tfsdk:"max_items"`
	Items     *[]*ListItemsItemsDataSourceModel `tfsdk:"items"`
}

type ListItemsItemsDataSourceModel struct {
}
