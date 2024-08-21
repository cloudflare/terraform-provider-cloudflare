// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dns_location

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDNSLocationResultDataSourceEnvelope struct {
	Result ZeroTrustDNSLocationDataSourceModel `json:"result,computed"`
}

type ZeroTrustDNSLocationResultListDataSourceEnvelope struct {
	Result *[]*ZeroTrustDNSLocationDataSourceModel `json:"result,computed"`
}

type ZeroTrustDNSLocationDataSourceModel struct {
	AccountID             types.String                                    `tfsdk:"account_id" path:"account_id"`
	LocationID            types.String                                    `tfsdk:"location_id" path:"location_id"`
	CreatedAt             timetypes.RFC3339                               `tfsdk:"created_at" json:"created_at,computed"`
	UpdatedAt             timetypes.RFC3339                               `tfsdk:"updated_at" json:"updated_at,computed"`
	ClientDefault         types.Bool                                      `tfsdk:"client_default" json:"client_default"`
	DNSDestinationIPsID   types.String                                    `tfsdk:"dns_destination_ips_id" json:"dns_destination_ips_id"`
	DOHSubdomain          types.String                                    `tfsdk:"doh_subdomain" json:"doh_subdomain"`
	ECSSupport            types.Bool                                      `tfsdk:"ecs_support" json:"ecs_support"`
	ID                    types.String                                    `tfsdk:"id" json:"id"`
	IP                    types.String                                    `tfsdk:"ip" json:"ip"`
	IPV4Destination       types.String                                    `tfsdk:"ipv4_destination" json:"ipv4_destination"`
	IPV4DestinationBackup types.String                                    `tfsdk:"ipv4_destination_backup" json:"ipv4_destination_backup"`
	Name                  types.String                                    `tfsdk:"name" json:"name"`
	Endpoints             *ZeroTrustDNSLocationEndpointsDataSourceModel   `tfsdk:"endpoints" json:"endpoints"`
	Networks              *[]*ZeroTrustDNSLocationNetworksDataSourceModel `tfsdk:"networks" json:"networks"`
	Filter                *ZeroTrustDNSLocationFindOneByDataSourceModel   `tfsdk:"filter"`
}

func (m *ZeroTrustDNSLocationDataSourceModel) toReadParams() (params zero_trust.GatewayLocationGetParams, diags diag.Diagnostics) {
	params = zero_trust.GatewayLocationGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ZeroTrustDNSLocationDataSourceModel) toListParams() (params zero_trust.GatewayLocationListParams, diags diag.Diagnostics) {
	params = zero_trust.GatewayLocationListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type ZeroTrustDNSLocationEndpointsDataSourceModel struct {
	DOH  *ZeroTrustDNSLocationEndpointsDOHDataSourceModel  `tfsdk:"doh" json:"doh"`
	DOT  *ZeroTrustDNSLocationEndpointsDOTDataSourceModel  `tfsdk:"dot" json:"dot"`
	IPV4 *ZeroTrustDNSLocationEndpointsIPV4DataSourceModel `tfsdk:"ipv4" json:"ipv4"`
	IPV6 *ZeroTrustDNSLocationEndpointsIPV6DataSourceModel `tfsdk:"ipv6" json:"ipv6"`
}

type ZeroTrustDNSLocationEndpointsDOHDataSourceModel struct {
	Enabled      types.Bool                                                  `tfsdk:"enabled" json:"enabled"`
	Networks     *[]*ZeroTrustDNSLocationEndpointsDOHNetworksDataSourceModel `tfsdk:"networks" json:"networks"`
	RequireToken types.Bool                                                  `tfsdk:"require_token" json:"require_token"`
}

type ZeroTrustDNSLocationEndpointsDOHNetworksDataSourceModel struct {
	Network types.String `tfsdk:"network" json:"network,computed"`
}

type ZeroTrustDNSLocationEndpointsDOTDataSourceModel struct {
	Enabled  types.Bool                                                  `tfsdk:"enabled" json:"enabled"`
	Networks *[]*ZeroTrustDNSLocationEndpointsDOTNetworksDataSourceModel `tfsdk:"networks" json:"networks"`
}

type ZeroTrustDNSLocationEndpointsDOTNetworksDataSourceModel struct {
	Network types.String `tfsdk:"network" json:"network,computed"`
}

type ZeroTrustDNSLocationEndpointsIPV4DataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled"`
}

type ZeroTrustDNSLocationEndpointsIPV6DataSourceModel struct {
	Enabled  types.Bool                                                   `tfsdk:"enabled" json:"enabled"`
	Networks *[]*ZeroTrustDNSLocationEndpointsIPV6NetworksDataSourceModel `tfsdk:"networks" json:"networks"`
}

type ZeroTrustDNSLocationEndpointsIPV6NetworksDataSourceModel struct {
	Network types.String `tfsdk:"network" json:"network,computed"`
}

type ZeroTrustDNSLocationNetworksDataSourceModel struct {
	Network types.String `tfsdk:"network" json:"network,computed"`
}

type ZeroTrustDNSLocationFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
