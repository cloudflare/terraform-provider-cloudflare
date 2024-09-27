// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_subscription

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountSubscriptionResultEnvelope struct {
	Result AccountSubscriptionModel `json:"result"`
}

type AccountSubscriptionModel struct {
	AccountID              types.String                      `tfsdk:"account_id" path:"account_id,required"`
	SubscriptionIdentifier types.String                      `tfsdk:"subscription_identifier" path:"subscription_identifier,optional"`
	Frequency              types.String                      `tfsdk:"frequency" json:"frequency,optional"`
	RatePlan               *AccountSubscriptionRatePlanModel `tfsdk:"rate_plan" json:"rate_plan,optional"`
	SubscriptionID         types.String                      `tfsdk:"subscription_id" json:"subscription_id,computed"`
}

type AccountSubscriptionRatePlanModel struct {
	ID                types.String                   `tfsdk:"id" json:"id,computed_optional"`
	Currency          types.String                   `tfsdk:"currency" json:"currency,computed_optional"`
	ExternallyManaged types.Bool                     `tfsdk:"externally_managed" json:"externally_managed,computed_optional"`
	IsContract        types.Bool                     `tfsdk:"is_contract" json:"is_contract,computed_optional"`
	PublicName        types.String                   `tfsdk:"public_name" json:"public_name,computed_optional"`
	Scope             types.String                   `tfsdk:"scope" json:"scope,computed_optional"`
	Sets              customfield.List[types.String] `tfsdk:"sets" json:"sets,computed_optional"`
}
