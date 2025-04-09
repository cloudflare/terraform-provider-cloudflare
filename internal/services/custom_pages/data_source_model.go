// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_pages

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/custom_pages"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomPagesResultDataSourceEnvelope struct {
Result CustomPagesDataSourceModel `json:"result,computed"`
}

type CustomPagesDataSourceModel struct {
Identifier types.String `tfsdk:"identifier" path:"identifier,required"`
AccountID types.String `tfsdk:"account_id" path:"account_id,optional"`
ZoneID types.String `tfsdk:"zone_id" path:"zone_id,optional"`
}

func (m *CustomPagesDataSourceModel) toReadParams(_ context.Context) (params custom_pages.CustomPageGetParams, diags diag.Diagnostics) {
  params = custom_pages.CustomPageGetParams{

  }

  if !m.AccountID.IsNull() {
    params.AccountID = cloudflare.F(m.AccountID.ValueString())
  } else {
    params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
  }

  return
}
