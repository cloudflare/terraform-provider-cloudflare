// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_subscription

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneSubscriptionResultEnvelope struct {
	Result ZoneSubscriptionModel `json:"result"`
}

type ZoneSubscriptionModel struct {
	ID                 types.String                   `tfsdk:"id" json:"-,computed"`
	ZoneID             types.String                   `tfsdk:"zone_id" path:"zone_id,required"`
	Frequency          types.String                   `tfsdk:"frequency" json:"frequency,optional"`
	RatePlan           *ZoneSubscriptionRatePlanModel `tfsdk:"rate_plan" json:"rate_plan,optional"`
	Currency           types.String                   `tfsdk:"currency" json:"currency,computed"`
	CurrentPeriodEnd   timetypes.RFC3339              `tfsdk:"current_period_end" json:"current_period_end,computed" format:"date-time"`
	CurrentPeriodStart timetypes.RFC3339              `tfsdk:"current_period_start" json:"current_period_start,computed" format:"date-time"`
	Price              types.Float64                  `tfsdk:"price" json:"price,computed"`
	State              types.String                   `tfsdk:"state" json:"state,computed"`
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
