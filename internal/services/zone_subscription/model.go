// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_subscription

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneSubscriptionResultEnvelope struct {
	Result ZoneSubscriptionModel `json:"result"`
}

type ZoneSubscriptionModel struct {
	Identifier      types.String                             `tfsdk:"identifier" path:"identifier"`
	Frequency       types.String                             `tfsdk:"frequency" json:"frequency"`
	App             *ZoneSubscriptionAppModel                `tfsdk:"app" json:"app"`
	ComponentValues *[]*ZoneSubscriptionComponentValuesModel `tfsdk:"component_values" json:"component_values"`
	RatePlan        *ZoneSubscriptionRatePlanModel           `tfsdk:"rate_plan" json:"rate_plan"`
	Zone            *ZoneSubscriptionZoneModel               `tfsdk:"zone" json:"zone"`
}

type ZoneSubscriptionAppModel struct {
	InstallID types.String `tfsdk:"install_id" json:"install_id"`
}

type ZoneSubscriptionComponentValuesModel struct {
	Default types.Float64 `tfsdk:"default" json:"default"`
	Name    types.String  `tfsdk:"name" json:"name"`
	Price   types.Float64 `tfsdk:"price" json:"price"`
	Value   types.Float64 `tfsdk:"value" json:"value"`
}

type ZoneSubscriptionRatePlanModel struct {
	ID                types.String    `tfsdk:"id" json:"id"`
	Currency          types.String    `tfsdk:"currency" json:"currency"`
	ExternallyManaged types.Bool      `tfsdk:"externally_managed" json:"externally_managed"`
	IsContract        types.Bool      `tfsdk:"is_contract" json:"is_contract"`
	PublicName        types.String    `tfsdk:"public_name" json:"public_name"`
	Scope             types.String    `tfsdk:"scope" json:"scope"`
	Sets              *[]types.String `tfsdk:"sets" json:"sets"`
}

type ZoneSubscriptionZoneModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}
