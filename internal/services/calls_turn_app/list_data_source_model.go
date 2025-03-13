// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package calls_turn_app

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/calls"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type CallsTURNAppsResultListDataSourceEnvelope struct {
Result customfield.NestedObjectList[CallsTURNAppsResultDataSourceModel] `json:"result,computed"`
}

type CallsTURNAppsDataSourceModel struct {
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
MaxItems types.Int64 `tfsdk:"max_items"`
Result customfield.NestedObjectList[CallsTURNAppsResultDataSourceModel] `tfsdk:"result"`
}

func (m *CallsTURNAppsDataSourceModel) toListParams(_ context.Context) (params calls.TURNListParams, diags diag.Diagnostics) {
  params = calls.TURNListParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  return
}

type CallsTURNAppsResultDataSourceModel struct {
Created timetypes.RFC3339 `tfsdk:"created" json:"created,computed" format:"date-time"`
Modified timetypes.RFC3339 `tfsdk:"modified" json:"modified,computed" format:"date-time"`
Name types.String `tfsdk:"name" json:"name,computed"`
UID types.String `tfsdk:"uid" json:"uid,computed"`
}
