// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list_item

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/rules"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ListItemResultDataSourceEnvelope struct {
	Result ListItemDataSourceModel `json:"result,computed"`
}

type ListItemDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
	ItemID    types.String `tfsdk:"item_id" path:"item_id,required"`
	ListID    types.String `tfsdk:"list_id" path:"list_id,required"`
}

func (m *ListItemDataSourceModel) toReadParams(_ context.Context) (params rules.ListItemGetParams, diags diag.Diagnostics) {
	params = rules.ListItemGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
