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

type ZeroTrustDNSLocationResultDataSourceEnvelope struct {
	Result ZeroTrustDNSLocationDataSourceModel `json:"result,computed"`
}

type ZeroTrustDNSLocationResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustDNSLocationDataSourceModel] `json:"result,computed"`
}

type ZeroTrustDNSLocationDataSourceModel struct {
	AccountID             types.String                                    `tfsdk:"account_id" path:"account_id"`
	LocationID            types.String                                    `tfsdk:"location_id" path:"location_id"`
	CreatedAt             timetypes.RFC3339                               `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UpdatedAt             timetypes.RFC3339                               `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	ClientDefault         types.Bool                                      `tfsdk:"client_default" json:"client_default,computed_optional"`
	DNSDestinationIPsID   types.String                                    `tfsdk:"dns_destination_ips_id" json:"dns_destination_ips_id,computed_optional"`
	DOHSubdomain          types.String                                    `tfsdk:"doh_subdomain" json:"doh_subdomain,computed_optional"`
	ECSSupport            types.Bool                                      `tfsdk:"ecs_support" json:"ecs_support,computed_optional"`
	ID                    types.String                                    `tfsdk:"id" json:"id,computed_optional"`
	IP                    types.String                                    `tfsdk:"ip" json:"ip,computed_optional"`
	IPV4Destination       types.String                                    `tfsdk:"ipv4_destination" json:"ipv4_destination,computed_optional"`
	IPV4DestinationBackup types.String                                    `tfsdk:"ipv4_destination_backup" json:"ipv4_destination_backup,computed_optional"`
	Name                  types.String                                    `tfsdk:"name" json:"name,computed_optional"`
	Endpoints             *ZeroTrustDNSLocationEndpointsDataSourceModel   `tfsdk:"endpoints" json:"endpoints,computed_optional"`
	Networks              *[]*ZeroTrustDNSLocationNetworksDataSourceModel `tfsdk:"networks" json:"networks,computed_optional"`
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
	DOH  *ZeroTrustDNSLocationEndpointsDOHDataSourceModel  `tfsdk:"doh" json:"doh,computed_optional"`
	DOT  *ZeroTrustDNSLocationEndpointsDOTDataSourceModel  `tfsdk:"dot" json:"dot,computed_optional"`
	IPV4 *ZeroTrustDNSLocationEndpointsIPV4DataSourceModel `tfsdk:"ipv4" json:"ipv4,computed_optional"`
	IPV6 *ZeroTrustDNSLocationEndpointsIPV6DataSourceModel `tfsdk:"ipv6" json:"ipv6,computed_optional"`
}

type ZeroTrustDNSLocationEndpointsDOHDataSourceModel struct {
	Enabled      types.Bool                                                  `tfsdk:"enabled" json:"enabled,computed_optional"`
	Networks     *[]*ZeroTrustDNSLocationEndpointsDOHNetworksDataSourceModel `tfsdk:"networks" json:"networks,computed_optional"`
	RequireToken types.Bool                                                  `tfsdk:"require_token" json:"require_token,computed_optional"`
}

type ZeroTrustDNSLocationEndpointsDOHNetworksDataSourceModel struct {
	Network types.String `tfsdk:"network" json:"network,computed"`
}

type ZeroTrustDNSLocationEndpointsDOTDataSourceModel struct {
	Enabled  types.Bool                                                  `tfsdk:"enabled" json:"enabled,computed_optional"`
	Networks *[]*ZeroTrustDNSLocationEndpointsDOTNetworksDataSourceModel `tfsdk:"networks" json:"networks,computed_optional"`
}

type ZeroTrustDNSLocationEndpointsDOTNetworksDataSourceModel struct {
	Network types.String `tfsdk:"network" json:"network,computed"`
}

type ZeroTrustDNSLocationEndpointsIPV4DataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed_optional"`
}

type ZeroTrustDNSLocationEndpointsIPV6DataSourceModel struct {
	Enabled  types.Bool                                                   `tfsdk:"enabled" json:"enabled,computed_optional"`
	Networks *[]*ZeroTrustDNSLocationEndpointsIPV6NetworksDataSourceModel `tfsdk:"networks" json:"networks,computed_optional"`
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