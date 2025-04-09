// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/rules"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type ListResultDataSourceEnvelope struct {
Result ListDataSourceModel `json:"result,computed"`
}

type ListDataSourceModel struct {
ID types.String `tfsdk:"id" path:"list_id,computed"`
ListID types.String `tfsdk:"list_id" path:"list_id,optional"`
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
CreatedOn types.String `tfsdk:"created_on" json:"created_on,computed"`
Description types.String `tfsdk:"description" json:"description,computed"`
Kind types.String `tfsdk:"kind" json:"kind,computed"`
ModifiedOn types.String `tfsdk:"modified_on" json:"modified_on,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
NumItems types.Float64 `tfsdk:"num_items" json:"num_items,computed"`
NumReferencingFilters types.Float64 `tfsdk:"num_referencing_filters" json:"num_referencing_filters,computed"`
}

func (m *ListDataSourceModel) toReadParams(_ context.Context) (params rules.ListGetParams, diags diag.Diagnostics) {
  params = rules.ListGetParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  return
}
