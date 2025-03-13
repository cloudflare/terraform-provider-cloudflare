// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/rules"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type ListsResultListDataSourceEnvelope struct {
Result customfield.NestedObjectList[ListsResultDataSourceModel] `json:"result,computed"`
}

type ListsDataSourceModel struct {
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
MaxItems types.Int64 `tfsdk:"max_items"`
Result customfield.NestedObjectList[ListsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ListsDataSourceModel) toListParams(_ context.Context) (params rules.ListListParams, diags diag.Diagnostics) {
  params = rules.ListListParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  return
}

type ListsResultDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
CreatedOn types.String `tfsdk:"created_on" json:"created_on,computed"`
Description types.String `tfsdk:"description" json:"description,computed"`
Kind types.String `tfsdk:"kind" json:"kind,computed"`
ModifiedOn types.String `tfsdk:"modified_on" json:"modified_on,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
NumItems types.Float64 `tfsdk:"num_items" json:"num_items,computed"`
NumReferencingFilters types.Float64 `tfsdk:"num_referencing_filters" json:"num_referencing_filters,computed"`
}
