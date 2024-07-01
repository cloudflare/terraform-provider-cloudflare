// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package teams_location

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TeamsLocationResultEnvelope struct {
	Result TeamsLocationModel `json:"result,computed"`
}

type TeamsLocationResultDataSourceEnvelope struct {
	Result TeamsLocationDataSourceModel `json:"result,computed"`
}

type TeamsLocationsResultDataSourceEnvelope struct {
	Result TeamsLocationsDataSourceModel `json:"result,computed"`
}

type TeamsLocationModel struct {
	ID            types.String                   `tfsdk:"id" json:"id,computed"`
	AccountID     types.String                   `tfsdk:"account_id" path:"account_id"`
	Name          types.String                   `tfsdk:"name" json:"name"`
	ClientDefault types.Bool                     `tfsdk:"client_default" json:"client_default"`
	EcsSupport    types.Bool                     `tfsdk:"ecs_support" json:"ecs_support"`
	Networks      *[]*TeamsLocationNetworksModel `tfsdk:"networks" json:"networks"`
	CreatedAt     types.String                   `tfsdk:"created_at" json:"created_at,computed"`
	DohSubdomain  types.String                   `tfsdk:"doh_subdomain" json:"doh_subdomain,computed"`
	IP            types.String                   `tfsdk:"ip" json:"ip,computed"`
	UpdatedAt     types.String                   `tfsdk:"updated_at" json:"updated_at,computed"`
}

type TeamsLocationNetworksModel struct {
	Network types.String `tfsdk:"network" json:"network"`
}

type TeamsLocationDataSourceModel struct {
}

type TeamsLocationsDataSourceModel struct {
}
