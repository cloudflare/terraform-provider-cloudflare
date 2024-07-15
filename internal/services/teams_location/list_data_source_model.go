// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package teams_location

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TeamsLocationsResultListDataSourceEnvelope struct {
	Result *[]*TeamsLocationsItemsDataSourceModel `json:"result,computed"`
}

type TeamsLocationsDataSourceModel struct {
	AccountID types.String                           `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                            `tfsdk:"max_items"`
	Items     *[]*TeamsLocationsItemsDataSourceModel `tfsdk:"items"`
}

type TeamsLocationsItemsDataSourceModel struct {
	ID            types.String                                   `tfsdk:"id" json:"id"`
	ClientDefault types.Bool                                     `tfsdk:"client_default" json:"client_default"`
	CreatedAt     timetypes.RFC3339                              `tfsdk:"created_at" json:"created_at"`
	DohSubdomain  types.String                                   `tfsdk:"doh_subdomain" json:"doh_subdomain"`
	EcsSupport    types.Bool                                     `tfsdk:"ecs_support" json:"ecs_support"`
	IP            types.String                                   `tfsdk:"ip" json:"ip"`
	Name          types.String                                   `tfsdk:"name" json:"name"`
	Networks      *[]*TeamsLocationsItemsNetworksDataSourceModel `tfsdk:"networks" json:"networks"`
	UpdatedAt     timetypes.RFC3339                              `tfsdk:"updated_at" json:"updated_at"`
}

type TeamsLocationsItemsNetworksDataSourceModel struct {
	Network types.String `tfsdk:"network" json:"network,computed"`
}
