// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dns_location

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDNSLocationsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustDNSLocationsResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustDNSLocationsDataSourceModel struct {
	AccountID types.String                                                             `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                                                              `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustDNSLocationsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustDNSLocationsDataSourceModel) toListParams() (params zero_trust.GatewayLocationListParams, diags diag.Diagnostics) {
	params = zero_trust.GatewayLocationListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustDNSLocationsResultDataSourceModel struct {
	ID                    types.String                                     `tfsdk:"id" json:"id"`
	ClientDefault         types.Bool                                       `tfsdk:"client_default" json:"client_default"`
	CreatedAt             timetypes.RFC3339                                `tfsdk:"created_at" json:"created_at,computed"`
	DNSDestinationIPsID   types.String                                     `tfsdk:"dns_destination_ips_id" json:"dns_destination_ips_id"`
	DOHSubdomain          types.String                                     `tfsdk:"doh_subdomain" json:"doh_subdomain"`
	ECSSupport            types.Bool                                       `tfsdk:"ecs_support" json:"ecs_support"`
	Endpoints             *ZeroTrustDNSLocationsEndpointsDataSourceModel   `tfsdk:"endpoints" json:"endpoints"`
	IP                    types.String                                     `tfsdk:"ip" json:"ip"`
	IPV4Destination       types.String                                     `tfsdk:"ipv4_destination" json:"ipv4_destination"`
	IPV4DestinationBackup types.String                                     `tfsdk:"ipv4_destination_backup" json:"ipv4_destination_backup"`
	Name                  types.String                                     `tfsdk:"name" json:"name"`
	Networks              *[]*ZeroTrustDNSLocationsNetworksDataSourceModel `tfsdk:"networks" json:"networks"`
	UpdatedAt             timetypes.RFC3339                                `tfsdk:"updated_at" json:"updated_at,computed"`
}

type ZeroTrustDNSLocationsEndpointsDataSourceModel struct {
	DOH  *ZeroTrustDNSLocationsEndpointsDOHDataSourceModel  `tfsdk:"doh" json:"doh"`
	DOT  *ZeroTrustDNSLocationsEndpointsDOTDataSourceModel  `tfsdk:"dot" json:"dot"`
	IPV4 *ZeroTrustDNSLocationsEndpointsIPV4DataSourceModel `tfsdk:"ipv4" json:"ipv4"`
	IPV6 *ZeroTrustDNSLocationsEndpointsIPV6DataSourceModel `tfsdk:"ipv6" json:"ipv6"`
}

type ZeroTrustDNSLocationsEndpointsDOHDataSourceModel struct {
	Enabled      types.Bool                                                   `tfsdk:"enabled" json:"enabled"`
	Networks     *[]*ZeroTrustDNSLocationsEndpointsDOHNetworksDataSourceModel `tfsdk:"networks" json:"networks"`
	RequireToken types.Bool                                                   `tfsdk:"require_token" json:"require_token"`
}

type ZeroTrustDNSLocationsEndpointsDOHNetworksDataSourceModel struct {
	Network types.String `tfsdk:"network" json:"network,computed"`
}

type ZeroTrustDNSLocationsEndpointsDOTDataSourceModel struct {
	Enabled  types.Bool                                                   `tfsdk:"enabled" json:"enabled"`
	Networks *[]*ZeroTrustDNSLocationsEndpointsDOTNetworksDataSourceModel `tfsdk:"networks" json:"networks"`
}

type ZeroTrustDNSLocationsEndpointsDOTNetworksDataSourceModel struct {
	Network types.String `tfsdk:"network" json:"network,computed"`
}

type ZeroTrustDNSLocationsEndpointsIPV4DataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled"`
}

type ZeroTrustDNSLocationsEndpointsIPV6DataSourceModel struct {
	Enabled  types.Bool                                                    `tfsdk:"enabled" json:"enabled"`
	Networks *[]*ZeroTrustDNSLocationsEndpointsIPV6NetworksDataSourceModel `tfsdk:"networks" json:"networks"`
}

type ZeroTrustDNSLocationsEndpointsIPV6NetworksDataSourceModel struct {
	Network types.String `tfsdk:"network" json:"network,computed"`
}

type ZeroTrustDNSLocationsNetworksDataSourceModel struct {
	Network types.String `tfsdk:"network" json:"network,computed"`
}
