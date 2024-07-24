// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package teams_location

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TeamsLocationResultDataSourceEnvelope struct {
	Result TeamsLocationDataSourceModel `json:"result,computed"`
}

type TeamsLocationResultListDataSourceEnvelope struct {
	Result *[]*TeamsLocationDataSourceModel `json:"result,computed"`
}

type TeamsLocationDataSourceModel struct {
	AccountID     types.String                             `tfsdk:"account_id" path:"account_id"`
	LocationID    types.String                             `tfsdk:"location_id" path:"location_id"`
	ID            types.String                             `tfsdk:"id" json:"id"`
	ClientDefault types.Bool                               `tfsdk:"client_default" json:"client_default"`
	CreatedAt             timetypes.RFC3339                        `tfsdk:"created_at" json:"created_at,computed"`
	CreatedAt     types.String                             `tfsdk:"created_at" json:"created_at"`
	DohSubdomain  types.String                             `tfsdk:"doh_subdomain" json:"doh_subdomain"`
	EcsSupport    types.Bool                               `tfsdk:"ecs_support" json:"ecs_support"`
	IP            types.String                             `tfsdk:"ip" json:"ip"`
	Name                  types.String                             `tfsdk:"name" json:"name"`
	Networks              *[]*TeamsLocationNetworksDataSourceModel `tfsdk:"networks" json:"networks"`
	UpdatedAt             timetypes.RFC3339                        `tfsdk:"updated_at" json:"updated_at,computed"`
	FindOneBy             *TeamsLocationFindOneByDataSourceModel   `tfsdk:"find_one_by"`
	Name          types.String                             `tfsdk:"name" json:"name"`
	Networks      *[]*TeamsLocationNetworksDataSourceModel `tfsdk:"networks" json:"networks"`
	UpdatedAt     types.String                             `tfsdk:"updated_at" json:"updated_at"`
	FindOneBy     *TeamsLocationFindOneByDataSourceModel   `tfsdk:"find_one_by"`
}

type TeamsLocationNetworksDataSourceModel struct {
	Network types.String `tfsdk:"network" json:"network,computed"`
}

type TeamsLocationFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
