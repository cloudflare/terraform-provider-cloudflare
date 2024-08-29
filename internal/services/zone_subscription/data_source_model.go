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

type ZoneSubscriptionResultDataSourceEnvelope struct {
	Result ZoneSubscriptionDataSourceModel `json:"result,computed"`
}

type ZoneSubscriptionResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZoneSubscriptionDataSourceModel] `json:"result,computed"`
}

type ZoneSubscriptionDataSourceModel struct {
	Identifier         types.String                                       `tfsdk:"identifier" path:"identifier"`
	Currency           types.String                                       `tfsdk:"currency" json:"currency"`
	CurrentPeriodEnd   timetypes.RFC3339                                  `tfsdk:"current_period_end" json:"current_period_end" format:"date-time"`
	CurrentPeriodStart timetypes.RFC3339                                  `tfsdk:"current_period_start" json:"current_period_start" format:"date-time"`
	Frequency          types.String                                       `tfsdk:"frequency" json:"frequency"`
	ID                 types.String                                       `tfsdk:"id" json:"id"`
	Price              types.Float64                                      `tfsdk:"price" json:"price"`
	State              types.String                                       `tfsdk:"state" json:"state"`
	App                *ZoneSubscriptionAppDataSourceModel                `tfsdk:"app" json:"app"`
	ComponentValues    *[]*ZoneSubscriptionComponentValuesDataSourceModel `tfsdk:"component_values" json:"component_values"`
	RatePlan           *ZoneSubscriptionRatePlanDataSourceModel           `tfsdk:"rate_plan" json:"rate_plan"`
	Zone               *ZoneSubscriptionZoneDataSourceModel               `tfsdk:"zone" json:"zone"`
	Filter             *ZoneSubscriptionFindOneByDataSourceModel          `tfsdk:"filter"`
}

func (m *ZoneSubscriptionDataSourceModel) toListParams() (params zones.SubscriptionListParams, diags diag.Diagnostics) {
	params = zones.SubscriptionListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type ZoneSubscriptionAppDataSourceModel struct {
	InstallID types.String `tfsdk:"install_id" json:"install_id,computed"`
}

type ZoneSubscriptionComponentValuesDataSourceModel struct {
	Default types.Float64 `tfsdk:"default" json:"default,computed"`
	Name    types.String  `tfsdk:"name" json:"name,computed"`
	Price   types.Float64 `tfsdk:"price" json:"price,computed"`
	Value   types.Float64 `tfsdk:"value" json:"value,computed"`
}

type ZoneSubscriptionRatePlanDataSourceModel struct {
	ID                types.String `tfsdk:"id" json:"id,computed"`
	Currency          types.String `tfsdk:"currency" json:"currency,computed"`
	ExternallyManaged types.Bool   `tfsdk:"externally_managed" json:"externally_managed,computed"`
	IsContract        types.Bool   `tfsdk:"is_contract" json:"is_contract,computed"`
	PublicName        types.String `tfsdk:"public_name" json:"public_name,computed"`
	Scope             types.String `tfsdk:"scope" json:"scope,computed"`
	Sets              types.List   `tfsdk:"sets" json:"sets,computed"`
}

type ZoneSubscriptionZoneDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type ZoneSubscriptionFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
