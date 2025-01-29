// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_subscription

import (
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneSubscriptionResultEnvelope struct {
	Result ZoneSubscriptionModel `json:"result"`
}

type ZoneSubscriptionModel struct {
	Identifier types.String                                            `tfsdk:"identifier" path:"identifier,required"`
	Frequency  types.String                                            `tfsdk:"frequency" json:"frequency,optional"`
	RatePlan   customfield.NestedObject[ZoneSubscriptionRatePlanModel] `tfsdk:"rate_plan" json:"rate_plan,computed_optional"`
}

func (m ZoneSubscriptionModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZoneSubscriptionModel) MarshalJSONForUpdate(state ZoneSubscriptionModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZoneSubscriptionRatePlanModel struct {
	ID                types.String    `tfsdk:"id" json:"id,optional"`
	Currency          types.String    `tfsdk:"currency" json:"currency,optional"`
	ExternallyManaged types.Bool      `tfsdk:"externally_managed" json:"externally_managed,optional"`
	IsContract        types.Bool      `tfsdk:"is_contract" json:"is_contract,optional"`
	PublicName        types.String    `tfsdk:"public_name" json:"public_name,optional"`
	Scope             types.String    `tfsdk:"scope" json:"scope,optional"`
	Sets              *[]types.String `tfsdk:"sets" json:"sets,optional"`
}
