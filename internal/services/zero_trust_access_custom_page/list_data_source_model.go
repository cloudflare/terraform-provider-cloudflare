// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_custom_page

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/zero_trust"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessCustomPagesResultListDataSourceEnvelope struct {
Result customfield.NestedObjectList[ZeroTrustAccessCustomPagesResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustAccessCustomPagesDataSourceModel struct {
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
MaxItems types.Int64 `tfsdk:"max_items"`
Result customfield.NestedObjectList[ZeroTrustAccessCustomPagesResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustAccessCustomPagesDataSourceModel) toListParams(_ context.Context) (params zero_trust.AccessCustomPageListParams, diags diag.Diagnostics) {
  params = zero_trust.AccessCustomPageListParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  return
}

type ZeroTrustAccessCustomPagesResultDataSourceModel struct {
Name types.String `tfsdk:"name" json:"name,computed"`
Type types.String `tfsdk:"type" json:"type,computed"`
AppCount types.Int64 `tfsdk:"app_count" json:"app_count,computed"`
CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
UID types.String `tfsdk:"uid" json:"uid,computed"`
UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}
