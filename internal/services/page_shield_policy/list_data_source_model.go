// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield_policy

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/page_shield"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type PageShieldPoliciesResultListDataSourceEnvelope struct {
Result customfield.NestedObjectList[PageShieldPoliciesResultDataSourceModel] `json:"result,computed"`
}

type PageShieldPoliciesDataSourceModel struct {
ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
MaxItems types.Int64 `tfsdk:"max_items"`
Result customfield.NestedObjectList[PageShieldPoliciesResultDataSourceModel] `tfsdk:"result"`
}

func (m *PageShieldPoliciesDataSourceModel) toListParams(_ context.Context) (params page_shield.PolicyListParams, diags diag.Diagnostics) {
  params = page_shield.PolicyListParams{
    ZoneID: cloudflare.F(m.ZoneID.ValueString()),
  }

  return
}

type PageShieldPoliciesResultDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
Action types.String `tfsdk:"action" json:"action,computed"`
Description types.String `tfsdk:"description" json:"description,computed"`
Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed"`
Expression types.String `tfsdk:"expression" json:"expression,computed"`
Value types.String `tfsdk:"value" json:"value,computed"`
}
