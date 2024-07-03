// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package teams_location

import (
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
	ID            types.String                                   `tfsdk:"id" json:"id,computed"`
	ClientDefault types.Bool                                     `tfsdk:"client_default" json:"client_default,computed"`
	CreatedAt     types.String                                   `tfsdk:"created_at" json:"created_at,computed"`
	DohSubdomain  types.String                                   `tfsdk:"doh_subdomain" json:"doh_subdomain,computed"`
	EcsSupport    types.Bool                                     `tfsdk:"ecs_support" json:"ecs_support,computed"`
	IP            types.String                                   `tfsdk:"ip" json:"ip,computed"`
	Name          types.String                                   `tfsdk:"name" json:"name,computed"`
	Networks      *[]*TeamsLocationsItemsNetworksDataSourceModel `tfsdk:"networks" json:"networks,computed"`
	UpdatedAt     types.String                                   `tfsdk:"updated_at" json:"updated_at,computed"`
}

type TeamsLocationsItemsNetworksDataSourceModel struct {
	Network types.String `tfsdk:"network" json:"network,computed"`
}
