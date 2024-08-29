// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_subscription

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zones"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneSubscriptionsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZoneSubscriptionsResultDataSourceModel] `json:"result,computed"`
}

type ZoneSubscriptionsDataSourceModel struct {
	AccountID types.String                                                         `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                                                          `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZoneSubscriptionsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZoneSubscriptionsDataSourceModel) toListParams() (params zones.SubscriptionListParams, diags diag.Diagnostics) {
	params = zones.SubscriptionListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZoneSubscriptionsResultDataSourceModel struct {
	ID                 types.String                                                                  `tfsdk:"id" json:"id,computed"`
	App                customfield.NestedObject[ZoneSubscriptionsAppDataSourceModel]                 `tfsdk:"app" json:"app,computed"`
	ComponentValues    customfield.NestedObjectList[ZoneSubscriptionsComponentValuesDataSourceModel] `tfsdk:"component_values" json:"component_values,computed"`
	Currency           types.String                                                                  `tfsdk:"currency" json:"currency,computed"`
	CurrentPeriodEnd   timetypes.RFC3339                                                             `tfsdk:"current_period_end" json:"current_period_end,computed" format:"date-time"`
	CurrentPeriodStart timetypes.RFC3339                                                             `tfsdk:"current_period_start" json:"current_period_start,computed" format:"date-time"`
	Frequency          types.String                                                                  `tfsdk:"frequency" json:"frequency,computed"`
	Price              types.Float64                                                                 `tfsdk:"price" json:"price,computed"`
	RatePlan           customfield.NestedObject[ZoneSubscriptionsRatePlanDataSourceModel]            `tfsdk:"rate_plan" json:"rate_plan,computed"`
	State              types.String                                                                  `tfsdk:"state" json:"state,computed"`
	Zone               customfield.NestedObject[ZoneSubscriptionsZoneDataSourceModel]                `tfsdk:"zone" json:"zone,computed"`
}

type ZoneSubscriptionsAppDataSourceModel struct {
	InstallID types.String `tfsdk:"install_id" json:"install_id,computed"`
}

type ZoneSubscriptionsComponentValuesDataSourceModel struct {
	Default types.Float64 `tfsdk:"default" json:"default,computed"`
	Name    types.String  `tfsdk:"name" json:"name,computed"`
	Price   types.Float64 `tfsdk:"price" json:"price,computed"`
	Value   types.Float64 `tfsdk:"value" json:"value,computed"`
}

type ZoneSubscriptionsRatePlanDataSourceModel struct {
	ID                types.String `tfsdk:"id" json:"id,computed"`
	Currency          types.String `tfsdk:"currency" json:"currency,computed"`
	ExternallyManaged types.Bool   `tfsdk:"externally_managed" json:"externally_managed,computed"`
	IsContract        types.Bool   `tfsdk:"is_contract" json:"is_contract,computed"`
	PublicName        types.String `tfsdk:"public_name" json:"public_name,computed"`
	Scope             types.String `tfsdk:"scope" json:"scope,computed"`
	Sets              types.List   `tfsdk:"sets" json:"sets,computed"`
}

type ZoneSubscriptionsZoneDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}
