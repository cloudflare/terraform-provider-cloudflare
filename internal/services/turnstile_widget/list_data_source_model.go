// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package turnstile_widget

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/turnstile"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type TurnstileWidgetsResultListDataSourceEnvelope struct {
Result customfield.NestedObjectList[TurnstileWidgetsResultDataSourceModel] `json:"result,computed"`
}

type TurnstileWidgetsDataSourceModel struct {
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
Direction types.String `tfsdk:"direction" query:"direction,optional"`
Order types.String `tfsdk:"order" query:"order,optional"`
MaxItems types.Int64 `tfsdk:"max_items"`
Result customfield.NestedObjectList[TurnstileWidgetsResultDataSourceModel] `tfsdk:"result"`
}

func (m *TurnstileWidgetsDataSourceModel) toListParams(_ context.Context) (params turnstile.WidgetListParams, diags diag.Diagnostics) {
  params = turnstile.WidgetListParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  if !m.Direction.IsNull() {
    params.Direction = cloudflare.F(turnstile.WidgetListParamsDirection(m.Direction.ValueString()))
  }
  if !m.Order.IsNull() {
    params.Order = cloudflare.F(turnstile.WidgetListParamsOrder(m.Order.ValueString()))
  }

  return
}

type TurnstileWidgetsResultDataSourceModel struct {
BotFightMode types.Bool `tfsdk:"bot_fight_mode" json:"bot_fight_mode,computed"`
ClearanceLevel types.String `tfsdk:"clearance_level" json:"clearance_level,computed"`
CreatedOn timetypes.RFC3339 `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
Domains customfield.List[types.String] `tfsdk:"domains" json:"domains,computed"`
EphemeralID types.Bool `tfsdk:"ephemeral_id" json:"ephemeral_id,computed"`
Mode types.String `tfsdk:"mode" json:"mode,computed"`
ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
Name types.String `tfsdk:"name" json:"name,computed"`
Offlabel types.Bool `tfsdk:"offlabel" json:"offlabel,computed"`
Region types.String `tfsdk:"region" json:"region,computed"`
Sitekey types.String `tfsdk:"sitekey" json:"sitekey,computed"`
}
