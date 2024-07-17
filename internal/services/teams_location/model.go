// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package teams_location

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TeamsLocationResultEnvelope struct {
	Result TeamsLocationModel `json:"result,computed"`
}

type TeamsLocationModel struct {
	ID                    types.String                   `tfsdk:"id" json:"id,computed"`
	AccountID             types.String                   `tfsdk:"account_id" path:"account_id"`
	Name                  types.String                   `tfsdk:"name" json:"name"`
	ClientDefault         types.Bool                     `tfsdk:"client_default" json:"client_default"`
	DNSDestinationIPsID   types.String                   `tfsdk:"dns_destination_ips_id" json:"dns_destination_ips_id"`
	ECSSupport            types.Bool                     `tfsdk:"ecs_support" json:"ecs_support"`
	Endpoints             *TeamsLocationEndpointsModel   `tfsdk:"endpoints" json:"endpoints"`
	Networks              *[]*TeamsLocationNetworksModel `tfsdk:"networks" json:"networks"`
	CreatedAt             timetypes.RFC3339              `tfsdk:"created_at" json:"created_at,computed"`
	DOHSubdomain          types.String                   `tfsdk:"doh_subdomain" json:"doh_subdomain,computed"`
	IP                    types.String                   `tfsdk:"ip" json:"ip,computed"`
	IPV4Destination       types.String                   `tfsdk:"ipv4_destination" json:"ipv4_destination,computed"`
	IPV4DestinationBackup types.String                   `tfsdk:"ipv4_destination_backup" json:"ipv4_destination_backup,computed"`
	UpdatedAt             timetypes.RFC3339              `tfsdk:"updated_at" json:"updated_at,computed"`
}

type TeamsLocationEndpointsModel struct {
	DOH  *TeamsLocationEndpointsDOHModel  `tfsdk:"doh" json:"doh"`
	DOT  *TeamsLocationEndpointsDOTModel  `tfsdk:"dot" json:"dot"`
	IPV4 *TeamsLocationEndpointsIPV4Model `tfsdk:"ipv4" json:"ipv4"`
	IPV6 *TeamsLocationEndpointsIPV6Model `tfsdk:"ipv6" json:"ipv6"`
}

type TeamsLocationEndpointsDOHModel struct {
	Enabled      types.Bool                                 `tfsdk:"enabled" json:"enabled"`
	Networks     *[]*TeamsLocationEndpointsDOHNetworksModel `tfsdk:"networks" json:"networks"`
	RequireToken types.Bool                                 `tfsdk:"require_token" json:"require_token"`
}

type TeamsLocationEndpointsDOHNetworksModel struct {
	Network types.String `tfsdk:"network" json:"network"`
}

type TeamsLocationEndpointsDOTModel struct {
	Enabled  types.Bool                                 `tfsdk:"enabled" json:"enabled"`
	Networks *[]*TeamsLocationEndpointsDOTNetworksModel `tfsdk:"networks" json:"networks"`
}

type TeamsLocationEndpointsDOTNetworksModel struct {
	Network types.String `tfsdk:"network" json:"network"`
}

type TeamsLocationEndpointsIPV4Model struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled"`
}

type TeamsLocationEndpointsIPV6Model struct {
	Enabled  types.Bool                                  `tfsdk:"enabled" json:"enabled"`
	Networks *[]*TeamsLocationEndpointsIPV6NetworksModel `tfsdk:"networks" json:"networks"`
}

type TeamsLocationEndpointsIPV6NetworksModel struct {
	Network types.String `tfsdk:"network" json:"network"`
}

type TeamsLocationNetworksModel struct {
	Network types.String `tfsdk:"network" json:"network"`
}
