// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list_item

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/rules"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ListItemsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ListItemsResultDataSourceModel] `json:"result,computed"`
}

type ListItemsDataSourceModel struct {
	AccountID types.String                                                 `tfsdk:"account_id" path:"account_id,required"`
	ListID    types.String                                                 `tfsdk:"list_id" path:"list_id,required"`
	Search    types.String                                                 `tfsdk:"search" query:"search,optional"`
	MaxItems  types.Int64                                                  `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ListItemsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ListItemsDataSourceModel) toListParams(_ context.Context) (params rules.ListItemListParams, diags diag.Diagnostics) {
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
