// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list_item

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ListItemResultDataSourceEnvelope struct {
	Result ListItemDataSourceModel `json:"result,computed"`
}

type ListItemResultListDataSourceEnvelope struct {
	Result *[]*ListItemDataSourceModel `json:"result,computed"`
}

type ListItemDataSourceModel struct {
	AccountIdentifier types.String                      `tfsdk:"account_identifier" path:"account_identifier"`
	ListID            types.String                      `tfsdk:"list_id" path:"list_id"`
	ItemID            types.String                      `tfsdk:"item_id" path:"item_id"`
	FindOneBy         *ListItemFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

type ListItemFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	ListID    types.String `tfsdk:"list_id" path:"list_id"`
	Cursor    types.String `tfsdk:"cursor" query:"cursor"`
	PerPage   types.Int64  `tfsdk:"per_page" query:"per_page"`
	Search    types.String `tfsdk:"search" query:"search"`
}
