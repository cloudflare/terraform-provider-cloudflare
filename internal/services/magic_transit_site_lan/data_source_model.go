// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_site_lan

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/magic_transit"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicTransitSiteLANResultDataSourceEnvelope struct {
	Result MagicTransitSiteLANDataSourceModel `json:"result,computed"`
}

type MagicTransitSiteLANDataSourceModel struct {
	ID               types.String                                                                  `tfsdk:"id" json:"-,computed"`
	LANID            types.String                                                                  `tfsdk:"lan_id" path:"lan_id,optional"`
	AccountID        types.String                                                                  `tfsdk:"account_id" path:"account_id,required"`
	SiteID           types.String                                                                  `tfsdk:"site_id" path:"site_id,computed"`
	HaLink           types.Bool                                                                    `tfsdk:"ha_link" json:"ha_link,computed"`
	Name             types.String                                                                  `tfsdk:"name" json:"name,computed"`
	Physport         types.Int64                                                                   `tfsdk:"physport" json:"physport,computed"`
	VlanTag          types.Int64                                                                   `tfsdk:"vlan_tag" json:"vlan_tag,computed"`
	Nat              customfield.NestedObject[MagicTransitSiteLANNatDataSourceModel]               `tfsdk:"nat" json:"nat,computed"`
	RoutedSubnets    customfield.NestedObjectList[MagicTransitSiteLANRoutedSubnetsDataSourceModel] `tfsdk:"routed_subnets" json:"routed_subnets,computed"`
	StaticAddressing customfield.NestedObject[MagicTransitSiteLANStaticAddressingDataSourceModel]  `tfsdk:"static_addressing" json:"static_addressing,computed"`
}

func (m *MagicTransitSiteLANDataSourceModel) toReadParams(_ context.Context) (params magic_transit.SiteLANGetParams, diags diag.Diagnostics) {
	params = magic_transit.SiteLANGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *MagicTransitSiteLANDataSourceModel) toListParams(_ context.Context) (params magic_transit.SiteLANListParams, diags diag.Diagnostics) {
	params = magic_transit.SiteLANListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type MagicTransitSiteLANNatDataSourceModel struct {
	StaticPrefix types.String `tfsdk:"static_prefix" json:"static_prefix,computed"`
}

type MagicTransitSiteLANRoutedSubnetsDataSourceModel struct {
	NextHop types.String                                                                 `tfsdk:"next_hop" json:"next_hop,computed"`
	Prefix  types.String                                                                 `tfsdk:"prefix" json:"prefix,computed"`
	Nat     customfield.NestedObject[MagicTransitSiteLANRoutedSubnetsNatDataSourceModel] `tfsdk:"nat" json:"nat,computed"`
}

type MagicTransitSiteLANRoutedSubnetsNatDataSourceModel struct {
	StaticPrefix types.String `tfsdk:"static_prefix" json:"static_prefix,computed"`
}

type MagicTransitSiteLANStaticAddressingDataSourceModel struct {
	Address          types.String                                                                           `tfsdk:"address" json:"address,computed"`
	DHCPRelay        customfield.NestedObject[MagicTransitSiteLANStaticAddressingDHCPRelayDataSourceModel]  `tfsdk:"dhcp_relay" json:"dhcp_relay,computed"`
	DHCPServer       customfield.NestedObject[MagicTransitSiteLANStaticAddressingDHCPServerDataSourceModel] `tfsdk:"dhcp_server" json:"dhcp_server,computed"`
	SecondaryAddress types.String                                                                           `tfsdk:"secondary_address" json:"secondary_address,computed"`
	VirtualAddress   types.String                                                                           `tfsdk:"virtual_address" json:"virtual_address,computed"`
}

type MagicTransitSiteLANStaticAddressingDHCPRelayDataSourceModel struct {
	ServerAddresses customfield.List[types.String] `tfsdk:"server_addresses" json:"server_addresses,computed"`
}

type MagicTransitSiteLANStaticAddressingDHCPServerDataSourceModel struct {
	DHCPPoolEnd   types.String                   `tfsdk:"dhcp_pool_end" json:"dhcp_pool_end,computed"`
	DHCPPoolStart types.String                   `tfsdk:"dhcp_pool_start" json:"dhcp_pool_start,computed"`
	DNSServer     types.String                   `tfsdk:"dns_server" json:"dns_server,computed"`
	DNSServers    customfield.List[types.String] `tfsdk:"dns_servers" json:"dns_servers,computed"`
	Reservations  customfield.Map[types.String]  `tfsdk:"reservations" json:"reservations,computed"`
}
