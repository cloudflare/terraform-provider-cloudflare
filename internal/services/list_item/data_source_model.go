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

type ListItemResultDataSourceEnvelope struct {
	Result ListItemDataSourceModel `json:"result,computed"`
}

type ListItemResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ListItemDataSourceModel] `json:"result,computed"`
}

type ListItemDataSourceModel struct {
	AccountIdentifier types.String                      `tfsdk:"account_identifier" path:"account_identifier"`
	ItemID            types.String                      `tfsdk:"item_id" path:"item_id"`
	ListID            types.String                      `tfsdk:"list_id" path:"list_id"`
	Filter            *ListItemFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *ListItemDataSourceModel) toListParams(_ context.Context) (params rules.ListItemListParams, diags diag.Diagnostics) {
	params = rules.ListItemListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	if !m.Filter.Search.IsNull() {
		params.Search = cloudflare.F(m.Filter.Search.ValueString())
	}

	return
}

type ListItemFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	ListID    types.String `tfsdk:"list_id" path:"list_id"`
	Search    types.String `tfsdk:"search" query:"search"`
}
