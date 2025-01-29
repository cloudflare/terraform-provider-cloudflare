// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_subscription

import (
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountSubscriptionResultEnvelope struct {
	Result AccountSubscriptionModel `json:"result"`
}

type AccountSubscriptionModel struct {
	AccountID              types.String                                               `tfsdk:"account_id" path:"account_id,required"`
	SubscriptionIdentifier types.String                                               `tfsdk:"subscription_identifier" path:"subscription_identifier,optional"`
	Frequency              types.String                                               `tfsdk:"frequency" json:"frequency,optional"`
	RatePlan               customfield.NestedObject[AccountSubscriptionRatePlanModel] `tfsdk:"rate_plan" json:"rate_plan,computed_optional"`
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
