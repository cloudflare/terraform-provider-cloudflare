// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package teams_location

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TeamsLocationsResultListDataSourceEnvelope struct {
	Result *[]*TeamsLocationsResultDataSourceModel `json:"result,computed"`
}

type TeamsLocationsDataSourceModel struct {
	AccountID types.String                            `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                             `tfsdk:"max_items"`
	Result    *[]*TeamsLocationsResultDataSourceModel `tfsdk:"result"`
}

type TeamsLocationsResultDataSourceModel struct {
	ID                    types.String                              `tfsdk:"id" json:"id"`
	ClientDefault         types.Bool                                `tfsdk:"client_default" json:"client_default"`
	CreatedAt             timetypes.RFC3339                         `tfsdk:"created_at" json:"created_at,computed"`
	DNSDestinationIPsID   types.String                              `tfsdk:"dns_destination_ips_id" json:"dns_destination_ips_id"`
	DOHSubdomain          types.String                              `tfsdk:"doh_subdomain" json:"doh_subdomain"`
	ECSSupport            types.Bool                                `tfsdk:"ecs_support" json:"ecs_support"`
	Endpoints             *TeamsLocationsEndpointsDataSourceModel   `tfsdk:"endpoints" json:"endpoints"`
	IP                    types.String                              `tfsdk:"ip" json:"ip"`
	IPV4Destination       types.String                              `tfsdk:"ipv4_destination" json:"ipv4_destination"`
	IPV4DestinationBackup types.String                              `tfsdk:"ipv4_destination_backup" json:"ipv4_destination_backup"`
	Name                  types.String                              `tfsdk:"name" json:"name"`
	Networks              *[]*TeamsLocationsNetworksDataSourceModel `tfsdk:"networks" json:"networks"`
	UpdatedAt             timetypes.RFC3339                         `tfsdk:"updated_at" json:"updated_at,computed"`
}

type TeamsLocationsEndpointsDataSourceModel struct {
	DOH  *TeamsLocationsEndpointsDOHDataSourceModel  `tfsdk:"doh" json:"doh"`
	DOT  *TeamsLocationsEndpointsDOTDataSourceModel  `tfsdk:"dot" json:"dot"`
	IPV4 *TeamsLocationsEndpointsIPV4DataSourceModel `tfsdk:"ipv4" json:"ipv4"`
	IPV6 *TeamsLocationsEndpointsIPV6DataSourceModel `tfsdk:"ipv6" json:"ipv6"`
}

type TeamsLocationsEndpointsDOHDataSourceModel struct {
	Enabled      types.Bool                                            `tfsdk:"enabled" json:"enabled"`
	Networks     *[]*TeamsLocationsEndpointsDOHNetworksDataSourceModel `tfsdk:"networks" json:"networks"`
	RequireToken types.Bool                                            `tfsdk:"require_token" json:"require_token"`
}

type TeamsLocationsEndpointsDOHNetworksDataSourceModel struct {
	Network types.String `tfsdk:"network" json:"network,computed"`
}

type TeamsLocationsEndpointsDOTDataSourceModel struct {
	Enabled  types.Bool                                            `tfsdk:"enabled" json:"enabled"`
	Networks *[]*TeamsLocationsEndpointsDOTNetworksDataSourceModel `tfsdk:"networks" json:"networks"`
}

type TeamsLocationsEndpointsDOTNetworksDataSourceModel struct {
	Network types.String `tfsdk:"network" json:"network,computed"`
}

type TeamsLocationsEndpointsIPV4DataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled"`
}

type TeamsLocationsEndpointsIPV6DataSourceModel struct {
	Enabled  types.Bool                                             `tfsdk:"enabled" json:"enabled"`
	Networks *[]*TeamsLocationsEndpointsIPV6NetworksDataSourceModel `tfsdk:"networks" json:"networks"`
}

type TeamsLocationsEndpointsIPV6NetworksDataSourceModel struct {
	Network types.String `tfsdk:"network" json:"network,computed"`
}

type TeamsLocationsNetworksDataSourceModel struct {
	Network types.String `tfsdk:"network" json:"network,computed"`
}
