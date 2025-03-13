// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package leaked_credential_check_rule

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/leaked_credential_checks"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type LeakedCredentialCheckRulesResultListDataSourceEnvelope struct {
Result customfield.NestedObjectList[LeakedCredentialCheckRulesResultDataSourceModel] `json:"result,computed"`
}

type LeakedCredentialCheckRulesDataSourceModel struct {
ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
MaxItems types.Int64 `tfsdk:"max_items"`
Result customfield.NestedObjectList[LeakedCredentialCheckRulesResultDataSourceModel] `tfsdk:"result"`
}

func (m *LeakedCredentialCheckRulesDataSourceModel) toListParams(_ context.Context) (params leaked_credential_checks.DetectionListParams, diags diag.Diagnostics) {
  params = leaked_credential_checks.DetectionListParams{
    ZoneID: cloudflare.F(m.ZoneID.ValueString()),
  }

  return
}

type LeakedCredentialCheckRulesResultDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
Password types.String `tfsdk:"password" json:"password,computed"`
Username types.String `tfsdk:"username" json:"username,computed"`
}
