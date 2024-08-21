// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list_item

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/rules"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ListItemsResultListDataSourceEnvelope struct {
	Result *[]*ListItemsResultDataSourceModel `json:"result,computed"`
}

type ListItemsDataSourceModel struct {
	AccountID types.String                       `tfsdk:"account_id" path:"account_id"`
	ListID    types.String                       `tfsdk:"list_id" path:"list_id"`
	Search    types.String                       `tfsdk:"search" query:"search"`
	MaxItems  types.Int64                        `tfsdk:"max_items"`
	Result    *[]*ListItemsResultDataSourceModel `tfsdk:"result"`
}

func (m *ListItemsDataSourceModel) toListParams() (params rules.ListItemListParams, diags diag.Diagnostics) {
	params = rules.ListItemListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Search.IsNull() {
		params.Search = cloudflare.F(m.Search.ValueString())
	}

	return
}

type ListItemsResultDataSourceModel struct {
}
