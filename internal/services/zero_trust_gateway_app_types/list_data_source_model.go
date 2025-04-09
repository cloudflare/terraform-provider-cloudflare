// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_app_types

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/zero_trust"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustGatewayAppTypesListResultListDataSourceEnvelope struct {
Result customfield.NestedObjectList[ZeroTrustGatewayAppTypesListResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustGatewayAppTypesListDataSourceModel struct {
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
MaxItems types.Int64 `tfsdk:"max_items"`
Result customfield.NestedObjectList[ZeroTrustGatewayAppTypesListResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustGatewayAppTypesListDataSourceModel) toListParams(_ context.Context) (params zero_trust.GatewayAppTypeListParams, diags diag.Diagnostics) {
  params = zero_trust.GatewayAppTypeListParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  return
}

type ZeroTrustGatewayAppTypesListResultDataSourceModel struct {
ID types.Int64 `tfsdk:"id" json:"id,computed"`
ApplicationTypeID types.Int64 `tfsdk:"application_type_id" json:"application_type_id,computed"`
CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
Name types.String `tfsdk:"name" json:"name,computed"`
Description types.String `tfsdk:"description" json:"description,computed"`
}
