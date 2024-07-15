// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package teams_location

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TeamsLocationResultDataSourceEnvelope struct {
	Result TeamsLocationDataSourceModel `json:"result,computed"`
}

type TeamsLocationResultListDataSourceEnvelope struct {
	Result *[]*TeamsLocationDataSourceModel `json:"result,computed"`
}

type TeamsLocationDataSourceModel struct {
	AccountID             types.String                             `tfsdk:"account_id" path:"account_id"`
	LocationID            types.String                             `tfsdk:"location_id" path:"location_id"`
	ID                    types.String                             `tfsdk:"id" json:"id"`
	ClientDefault         types.Bool                               `tfsdk:"client_default" json:"client_default"`
	CreatedAt             timetypes.RFC3339                        `tfsdk:"created_at" json:"created_at"`
	DNSDestinationIPsID   types.String                             `tfsdk:"dns_destination_ips_id" json:"dns_destination_ips_id"`
	DohSubdomain          types.String                             `tfsdk:"doh_subdomain" json:"doh_subdomain"`
	EcsSupport            types.Bool                               `tfsdk:"ecs_support" json:"ecs_support"`
	Endpoints             *TeamsLocationEndpointsDataSourceModel   `tfsdk:"endpoints" json:"endpoints"`
	IP                    types.String                             `tfsdk:"ip" json:"ip"`
	IPV4Destination       types.String                             `tfsdk:"ipv4_destination" json:"ipv4_destination"`
	IPV4DestinationBackup types.String                             `tfsdk:"ipv4_destination_backup" json:"ipv4_destination_backup"`
	Name                  types.String                             `tfsdk:"name" json:"name"`
	Networks              *[]*TeamsLocationNetworksDataSourceModel `tfsdk:"networks" json:"networks"`
	UpdatedAt             timetypes.RFC3339                        `tfsdk:"updated_at" json:"updated_at"`
	FindOneBy             *TeamsLocationFindOneByDataSourceModel   `tfsdk:"find_one_by"`
}

type TeamsLocationEndpointsDataSourceModel struct {
	Doh  *TeamsLocationEndpointsDohDataSourceModel  `tfsdk:"doh" json:"doh"`
	Dot  *TeamsLocationEndpointsDotDataSourceModel  `tfsdk:"dot" json:"dot"`
	IPV4 *TeamsLocationEndpointsIPV4DataSourceModel `tfsdk:"ipv4" json:"ipv4"`
	IPV6 *TeamsLocationEndpointsIPV6DataSourceModel `tfsdk:"ipv6" json:"ipv6"`
}

type TeamsLocationEndpointsDohDataSourceModel struct {
	Enabled      types.Bool                                           `tfsdk:"enabled" json:"enabled"`
	Networks     *[]*TeamsLocationEndpointsDohNetworksDataSourceModel `tfsdk:"networks" json:"networks"`
	RequireToken types.Bool                                           `tfsdk:"require_token" json:"require_token"`
}

type TeamsLocationEndpointsDohNetworksDataSourceModel struct {
	Network types.String `tfsdk:"network" json:"network,computed"`
}

type TeamsLocationEndpointsDotDataSourceModel struct {
	Enabled  types.Bool                                           `tfsdk:"enabled" json:"enabled"`
	Networks *[]*TeamsLocationEndpointsDotNetworksDataSourceModel `tfsdk:"networks" json:"networks"`
}

type TeamsLocationEndpointsDotNetworksDataSourceModel struct {
	Network types.String `tfsdk:"network" json:"network,computed"`
}

type TeamsLocationEndpointsIPV4DataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled"`
}

type TeamsLocationEndpointsIPV6DataSourceModel struct {
	Enabled  types.Bool                                            `tfsdk:"enabled" json:"enabled"`
	Networks *[]*TeamsLocationEndpointsIPV6NetworksDataSourceModel `tfsdk:"networks" json:"networks"`
}

type TeamsLocationEndpointsIPV6NetworksDataSourceModel struct {
	Network types.String `tfsdk:"network" json:"network,computed"`
}

type TeamsLocationNetworksDataSourceModel struct {
	Network types.String `tfsdk:"network" json:"network,computed"`
}

type TeamsLocationFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
