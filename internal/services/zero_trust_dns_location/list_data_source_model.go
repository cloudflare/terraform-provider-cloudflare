// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dns_location

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDNSLocationsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustDNSLocationsResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustDNSLocationsDataSourceModel struct {
	AccountID types.String                                                             `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                              `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustDNSLocationsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustDNSLocationsDataSourceModel) toListParams(_ context.Context) (params zero_trust.GatewayLocationListParams, diags diag.Diagnostics) {
	params = zero_trust.GatewayLocationListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustDNSLocationsResultDataSourceModel struct {
	ID                        types.String                                                               `tfsdk:"id" json:"id,computed"`
	ClientDefault             types.Bool                                                                 `tfsdk:"client_default" json:"client_default,computed"`
	CreatedAt                 timetypes.RFC3339                                                          `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DNSDestinationIPsID       types.String                                                               `tfsdk:"dns_destination_ips_id" json:"dns_destination_ips_id,computed"`
	DNSDestinationIPV6BlockID types.String                                                               `tfsdk:"dns_destination_ipv6_block_id" json:"dns_destination_ipv6_block_id,computed"`
	DOHSubdomain              types.String                                                               `tfsdk:"doh_subdomain" json:"doh_subdomain,computed"`
	ECSSupport                types.Bool                                                                 `tfsdk:"ecs_support" json:"ecs_support,computed"`
	Endpoints                 customfield.NestedObject[ZeroTrustDNSLocationsEndpointsDataSourceModel]    `tfsdk:"endpoints" json:"endpoints,computed"`
	IP                        types.String                                                               `tfsdk:"ip" json:"ip,computed"`
	IPV4Destination           types.String                                                               `tfsdk:"ipv4_destination" json:"ipv4_destination,computed"`
	IPV4DestinationBackup     types.String                                                               `tfsdk:"ipv4_destination_backup" json:"ipv4_destination_backup,computed"`
	Name                      types.String                                                               `tfsdk:"name" json:"name,computed"`
	Networks                  customfield.NestedObjectList[ZeroTrustDNSLocationsNetworksDataSourceModel] `tfsdk:"networks" json:"networks,computed"`
	UpdatedAt                 timetypes.RFC3339                                                          `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

type ZeroTrustDNSLocationsEndpointsDataSourceModel struct {
	DOH  customfield.NestedObject[ZeroTrustDNSLocationsEndpointsDOHDataSourceModel]  `tfsdk:"doh" json:"doh,computed"`
	DOT  customfield.NestedObject[ZeroTrustDNSLocationsEndpointsDOTDataSourceModel]  `tfsdk:"dot" json:"dot,computed"`
	IPV4 customfield.NestedObject[ZeroTrustDNSLocationsEndpointsIPV4DataSourceModel] `tfsdk:"ipv4" json:"ipv4,computed"`
	IPV6 customfield.NestedObject[ZeroTrustDNSLocationsEndpointsIPV6DataSourceModel] `tfsdk:"ipv6" json:"ipv6,computed"`
}

type ZeroTrustDNSLocationsEndpointsDOHDataSourceModel struct {
	Enabled      types.Bool                                                                             `tfsdk:"enabled" json:"enabled,computed"`
	Networks     customfield.NestedObjectList[ZeroTrustDNSLocationsEndpointsDOHNetworksDataSourceModel] `tfsdk:"networks" json:"networks,computed"`
	RequireToken types.Bool                                                                             `tfsdk:"require_token" json:"require_token,computed"`
}

type ZeroTrustDNSLocationsEndpointsDOHNetworksDataSourceModel struct {
	Network types.String `tfsdk:"network" json:"network,computed"`
}

type ZeroTrustDNSLocationsEndpointsDOTDataSourceModel struct {
	Enabled  types.Bool                                                                             `tfsdk:"enabled" json:"enabled,computed"`
	Networks customfield.NestedObjectList[ZeroTrustDNSLocationsEndpointsDOTNetworksDataSourceModel] `tfsdk:"networks" json:"networks,computed"`
}

type ZeroTrustDNSLocationsEndpointsDOTNetworksDataSourceModel struct {
	Network types.String `tfsdk:"network" json:"network,computed"`
}

type ZeroTrustDNSLocationsEndpointsIPV4DataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed"`
}

type ZeroTrustDNSLocationsEndpointsIPV6DataSourceModel struct {
	Enabled  types.Bool                                                                              `tfsdk:"enabled" json:"enabled,computed"`
	Networks customfield.NestedObjectList[ZeroTrustDNSLocationsEndpointsIPV6NetworksDataSourceModel] `tfsdk:"networks" json:"networks,computed"`
}

type ZeroTrustDNSLocationsEndpointsIPV6NetworksDataSourceModel struct {
	Network types.String `tfsdk:"network" json:"network,computed"`
}

type ZeroTrustDNSLocationsNetworksDataSourceModel struct {
	Network types.String `tfsdk:"network" json:"network,computed"`
}
