// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_subscription

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/zones"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneSubscriptionResultDataSourceEnvelope struct {
	Result ZoneSubscriptionDataSourceModel `json:"result,computed"`
}

type ZoneSubscriptionDataSourceModel struct {
	ZoneID             types.String                                                      `tfsdk:"zone_id" path:"zone_id,required"`
	Currency           types.String                                                      `tfsdk:"currency" json:"currency,computed"`
	CurrentPeriodEnd   timetypes.RFC3339                                                 `tfsdk:"current_period_end" json:"current_period_end,computed" format:"date-time"`
	CurrentPeriodStart timetypes.RFC3339                                                 `tfsdk:"current_period_start" json:"current_period_start,computed" format:"date-time"`
	Frequency          types.String                                                      `tfsdk:"frequency" json:"frequency,computed"`
	ID                 types.String                                                      `tfsdk:"id" json:"id,computed"`
	Price              types.Float64                                                     `tfsdk:"price" json:"price,computed"`
	State              types.String                                                      `tfsdk:"state" json:"state,computed"`
	RatePlan           customfield.NestedObject[ZoneSubscriptionRatePlanDataSourceModel] `tfsdk:"rate_plan" json:"rate_plan,computed"`
}

func (m *ZoneSubscriptionDataSourceModel) toReadParams(_ context.Context) (params zones.SubscriptionGetParams, diags diag.Diagnostics) {
	params = zones.SubscriptionGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type ZoneSubscriptionRatePlanDataSourceModel struct {
	ID                types.String                   `tfsdk:"id" json:"id,computed"`
	Currency          types.String                   `tfsdk:"currency" json:"currency,computed"`
	ExternallyManaged types.Bool                     `tfsdk:"externally_managed" json:"externally_managed,computed"`
	IsContract        types.Bool                     `tfsdk:"is_contract" json:"is_contract,computed"`
	PublicName        types.String                   `tfsdk:"public_name" json:"public_name,computed"`
	Scope             types.String                   `tfsdk:"scope" json:"scope,computed"`
	Sets              customfield.List[types.String] `tfsdk:"sets" json:"sets,computed"`
}
