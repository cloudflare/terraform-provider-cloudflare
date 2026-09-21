// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dns_location

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDNSLocationResultDataSourceEnvelope struct {
	Result ZeroTrustDNSLocationDataSourceModel `json:"result,computed"`
}

type ZeroTrustDNSLocationDataSourceModel struct {
	ID                        types.String                                                              `tfsdk:"id" path:"location_id,computed"`
	LocationID                types.String                                                              `tfsdk:"location_id" path:"location_id,required"`
	AccountID                 types.String                                                              `tfsdk:"account_id" path:"account_id,optional"`
	ClientDefault             types.Bool                                                                `tfsdk:"client_default" json:"client_default,computed"`
	CreatedAt                 timetypes.RFC3339                                                         `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DNSDestinationIPsID       types.String                                                              `tfsdk:"dns_destination_ips_id" json:"dns_destination_ips_id,computed"`
	DNSDestinationIPV6BlockID types.String                                                              `tfsdk:"dns_destination_ipv6_block_id" json:"dns_destination_ipv6_block_id,computed"`
	DOHSubdomain              types.String                                                              `tfsdk:"doh_subdomain" json:"doh_subdomain,computed"`
	ECSSupport                types.Bool                                                                `tfsdk:"ecs_support" json:"ecs_support,computed"`
	IP                        types.String                                                              `tfsdk:"ip" json:"ip,computed"`
	IPV4Destination           types.String                                                              `tfsdk:"ipv4_destination" json:"ipv4_destination,computed"`
	IPV4DestinationBackup     types.String                                                              `tfsdk:"ipv4_destination_backup" json:"ipv4_destination_backup,computed"`
	Name                      types.String                                                              `tfsdk:"name" json:"name,computed"`
	UpdatedAt                 timetypes.RFC3339                                                         `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Endpoints                 customfield.NestedObject[ZeroTrustDNSLocationEndpointsDataSourceModel]    `tfsdk:"endpoints" json:"endpoints,computed"`
	MaxTTL                    customfield.NestedObject[ZeroTrustDNSLocationMaxTTLDataSourceModel]       `tfsdk:"max_ttl" json:"max_ttl,computed"`
	Networks                  customfield.NestedObjectList[ZeroTrustDNSLocationNetworksDataSourceModel] `tfsdk:"networks" json:"networks,computed"`
	Filter                    *ZeroTrustDNSLocationFindOneByDataSourceModel                             `tfsdk:"filter"`
}

func (m *ZeroTrustDNSLocationDataSourceModel) toReadParams(_ context.Context) (params zero_trust.GatewayLocationGetParams, diags diag.Diagnostics) {
	params = zero_trust.GatewayLocationGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ZeroTrustDNSLocationDataSourceModel) toListParams(_ context.Context) (params zero_trust.GatewayLocationListParams, diags diag.Diagnostics) {
	mFilterFilter := []interface{}{}
	if m.Filter.Filter != nil {
		for _, item := range *m.Filter.Filter {
			mFilterFilter = append(mFilterFilter, item.ValueString())
		}
	}

	params = zero_trust.GatewayLocationListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
		Filter:    cloudflare.F(mFilterFilter),
	}

	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(zero_trust.GatewayLocationListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.OrderBy.IsNull() {
		params.OrderBy = cloudflare.F(zero_trust.GatewayLocationListParamsOrderBy(m.Filter.OrderBy.ValueString()))
	}
	if !m.Filter.Search.IsNull() {
		params.Search = cloudflare.F(m.Filter.Search.ValueString())
	}

	return
}

type ZeroTrustDNSLocationEndpointsDataSourceModel struct {
	DOH  customfield.NestedObject[ZeroTrustDNSLocationEndpointsDOHDataSourceModel]  `tfsdk:"doh" json:"doh,computed"`
	DOT  customfield.NestedObject[ZeroTrustDNSLocationEndpointsDOTDataSourceModel]  `tfsdk:"dot" json:"dot,computed"`
	IPV4 customfield.NestedObject[ZeroTrustDNSLocationEndpointsIPV4DataSourceModel] `tfsdk:"ipv4" json:"ipv4,computed"`
	IPV6 customfield.NestedObject[ZeroTrustDNSLocationEndpointsIPV6DataSourceModel] `tfsdk:"ipv6" json:"ipv6,computed"`
}

type ZeroTrustDNSLocationEndpointsDOHDataSourceModel struct {
	Enabled      types.Bool                                                                            `tfsdk:"enabled" json:"enabled,computed"`
	Networks     customfield.NestedObjectList[ZeroTrustDNSLocationEndpointsDOHNetworksDataSourceModel] `tfsdk:"networks" json:"networks,computed"`
	RequireToken types.Bool                                                                            `tfsdk:"require_token" json:"require_token,computed"`
}

type ZeroTrustDNSLocationEndpointsDOHNetworksDataSourceModel struct {
	Network types.String `tfsdk:"network" json:"network,computed"`
}

type ZeroTrustDNSLocationEndpointsDOTDataSourceModel struct {
	Enabled  types.Bool                                                                            `tfsdk:"enabled" json:"enabled,computed"`
	Networks customfield.NestedObjectList[ZeroTrustDNSLocationEndpointsDOTNetworksDataSourceModel] `tfsdk:"networks" json:"networks,computed"`
}

type ZeroTrustDNSLocationEndpointsDOTNetworksDataSourceModel struct {
	Network types.String `tfsdk:"network" json:"network,computed"`
}

type ZeroTrustDNSLocationEndpointsIPV4DataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed"`
}

type ZeroTrustDNSLocationEndpointsIPV6DataSourceModel struct {
	Enabled  types.Bool                                                                             `tfsdk:"enabled" json:"enabled,computed"`
	Networks customfield.NestedObjectList[ZeroTrustDNSLocationEndpointsIPV6NetworksDataSourceModel] `tfsdk:"networks" json:"networks,computed"`
}

type ZeroTrustDNSLocationEndpointsIPV6NetworksDataSourceModel struct {
	Network types.String `tfsdk:"network" json:"network,computed"`
}

type ZeroTrustDNSLocationMaxTTLDataSourceModel struct {
	Mode    types.String `tfsdk:"mode" json:"mode,computed"`
	TTLSecs types.Int64  `tfsdk:"ttl_secs" json:"ttl_secs,computed"`
}

type ZeroTrustDNSLocationNetworksDataSourceModel struct {
	Network types.String `tfsdk:"network" json:"network,computed"`
}

type ZeroTrustDNSLocationFindOneByDataSourceModel struct {
	Direction types.String    `tfsdk:"direction" query:"direction,optional"`
	Filter    *[]types.String `tfsdk:"filter" query:"filter,optional"`
	OrderBy   types.String    `tfsdk:"order_by" query:"order_by,optional"`
	Search    types.String    `tfsdk:"search" query:"search,optional"`
}
