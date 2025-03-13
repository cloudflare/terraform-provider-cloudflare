// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_kv_namespace

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/kv"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersKVNamespacesResultListDataSourceEnvelope struct {
Result customfield.NestedObjectList[WorkersKVNamespacesResultDataSourceModel] `json:"result,computed"`
}

type WorkersKVNamespacesDataSourceModel struct {
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
Direction types.String `tfsdk:"direction" query:"direction,optional"`
Order types.String `tfsdk:"order" query:"order,optional"`
MaxItems types.Int64 `tfsdk:"max_items"`
Result customfield.NestedObjectList[WorkersKVNamespacesResultDataSourceModel] `tfsdk:"result"`
}

func (m *WorkersKVNamespacesDataSourceModel) toListParams(_ context.Context) (params kv.NamespaceListParams, diags diag.Diagnostics) {
  params = kv.NamespaceListParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  if !m.Direction.IsNull() {
    params.Direction = cloudflare.F(kv.NamespaceListParamsDirection(m.Direction.ValueString()))
  }
  if !m.Order.IsNull() {
    params.Order = cloudflare.F(kv.NamespaceListParamsOrder(m.Order.ValueString()))
  }

  return
}

type WorkersKVNamespacesResultDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
Title types.String `tfsdk:"title" json:"title,computed"`
SupportsURLEncoding types.Bool `tfsdk:"supports_url_encoding" json:"supports_url_encoding,computed"`
}
