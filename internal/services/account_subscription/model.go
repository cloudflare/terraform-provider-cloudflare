// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_subscription

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountSubscriptionResultEnvelope struct {
	Result AccountSubscriptionModel `json:"result"`
}

type AccountSubscriptionModel struct {
	ID                 types.String                      `tfsdk:"id" json:"id,computed"`
	AccountID          types.String                      `tfsdk:"account_id" path:"account_id,required"`
	Frequency          types.String                      `tfsdk:"frequency" json:"frequency,optional"`
	RatePlan           *AccountSubscriptionRatePlanModel `tfsdk:"rate_plan" json:"rate_plan,optional"`
	Currency           types.String                      `tfsdk:"currency" json:"currency,computed"`
	CurrentPeriodEnd   timetypes.RFC3339                 `tfsdk:"current_period_end" json:"current_period_end,computed" format:"date-time"`
	CurrentPeriodStart timetypes.RFC3339                 `tfsdk:"current_period_start" json:"current_period_start,computed" format:"date-time"`
	Price              types.Float64                     `tfsdk:"price" json:"price,computed"`
	State              types.String                      `tfsdk:"state" json:"state,computed"`
}

func (m AccountSubscriptionModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AccountSubscriptionModel) MarshalJSONForUpdate(state AccountSubscriptionModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type AccountSubscriptionRatePlanModel struct {
	ID                types.String    `tfsdk:"id" json:"id,optional"`
	Currency          types.String    `tfsdk:"currency" json:"currency,optional"`
	ExternallyManaged types.Bool      `tfsdk:"externally_managed" json:"externally_managed,optional"`
	IsContract        types.Bool      `tfsdk:"is_contract" json:"is_contract,optional"`
	PublicName        types.String    `tfsdk:"public_name" json:"public_name,optional"`
	Scope             types.String    `tfsdk:"scope" json:"scope,optional"`
	Sets              *[]types.String `tfsdk:"sets" json:"sets,optional"`
}
