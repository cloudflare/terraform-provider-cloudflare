// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_subscription

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/accounts"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountSubscriptionResultDataSourceEnvelope struct {
Result AccountSubscriptionDataSourceModel `json:"result,computed"`
}

type AccountSubscriptionDataSourceModel struct {
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
Currency types.String `tfsdk:"currency" json:"currency,computed"`
CurrentPeriodEnd timetypes.RFC3339 `tfsdk:"current_period_end" json:"current_period_end,computed" format:"date-time"`
CurrentPeriodStart timetypes.RFC3339 `tfsdk:"current_period_start" json:"current_period_start,computed" format:"date-time"`
Frequency types.String `tfsdk:"frequency" json:"frequency,computed"`
ID types.String `tfsdk:"id" json:"id,computed"`
Price types.Float64 `tfsdk:"price" json:"price,computed"`
State types.String `tfsdk:"state" json:"state,computed"`
RatePlan customfield.NestedObject[AccountSubscriptionRatePlanDataSourceModel] `tfsdk:"rate_plan" json:"rate_plan,computed"`
}

func (m *AccountSubscriptionDataSourceModel) toReadParams(_ context.Context) (params accounts.SubscriptionGetParams, diags diag.Diagnostics) {
  params = accounts.SubscriptionGetParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  return
}

type AccountSubscriptionRatePlanDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
Currency types.String `tfsdk:"currency" json:"currency,computed"`
ExternallyManaged types.Bool `tfsdk:"externally_managed" json:"externally_managed,computed"`
IsContract types.Bool `tfsdk:"is_contract" json:"is_contract,computed"`
PublicName types.String `tfsdk:"public_name" json:"public_name,computed"`
Scope types.String `tfsdk:"scope" json:"scope,computed"`
Sets customfield.List[types.String] `tfsdk:"sets" json:"sets,computed"`
}
