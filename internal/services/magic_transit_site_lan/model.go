// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_site_lan

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicTransitSiteLANResultEnvelope struct {
	Result MagicTransitSiteLANModel `json:"result"`
}

type MagicTransitSiteLANModel struct {
	ID               types.String                              `tfsdk:"id" json:"id,computed"`
	AccountID        types.String                              `tfsdk:"account_id" path:"account_id,required"`
	SiteID           types.String                              `tfsdk:"site_id" path:"site_id,required"`
	HaLink           types.Bool                                `tfsdk:"ha_link" json:"ha_link,optional"`
	Physport         types.Int64                               `tfsdk:"physport" json:"physport,required"`
	Name             types.String                              `tfsdk:"name" json:"name,optional"`
	VlanTag          types.Int64                               `tfsdk:"vlan_tag" json:"vlan_tag,optional"`
	Nat              *MagicTransitSiteLANNatModel              `tfsdk:"nat" json:"nat,optional"`
	RoutedSubnets    *[]*MagicTransitSiteLANRoutedSubnetsModel `tfsdk:"routed_subnets" json:"routed_subnets,optional"`
	StaticAddressing *MagicTransitSiteLANStaticAddressingModel `tfsdk:"static_addressing" json:"static_addressing,optional"`
}

func (m MagicTransitSiteLANModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m MagicTransitSiteLANModel) MarshalJSONForUpdate(state MagicTransitSiteLANModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type MagicTransitSiteLANNatModel struct {
	StaticPrefix types.String `tfsdk:"static_prefix" json:"static_prefix,optional"`
}

type MagicTransitSiteLANRoutedSubnetsModel struct {
	NextHop types.String                              `tfsdk:"next_hop" json:"next_hop,required"`
	Prefix  types.String                              `tfsdk:"prefix" json:"prefix,required"`
	Nat     *MagicTransitSiteLANRoutedSubnetsNatModel `tfsdk:"nat" json:"nat,optional"`
}

type MagicTransitSiteLANRoutedSubnetsNatModel struct {
	StaticPrefix types.String `tfsdk:"static_prefix" json:"static_prefix,optional"`
}

type MagicTransitSiteLANStaticAddressingModel struct {
	Address          types.String                                        `tfsdk:"address" json:"address,required"`
	DHCPRelay        *MagicTransitSiteLANStaticAddressingDHCPRelayModel  `tfsdk:"dhcp_relay" json:"dhcp_relay,optional"`
	DHCPServer       *MagicTransitSiteLANStaticAddressingDHCPServerModel `tfsdk:"dhcp_server" json:"dhcp_server,optional"`
	SecondaryAddress types.String                                        `tfsdk:"secondary_address" json:"secondary_address,optional"`
	VirtualAddress   types.String                                        `tfsdk:"virtual_address" json:"virtual_address,optional"`
}

type MagicTransitSiteLANStaticAddressingDHCPRelayModel struct {
	ServerAddresses *[]types.String `tfsdk:"server_addresses" json:"server_addresses,optional"`
}

type MagicTransitSiteLANStaticAddressingDHCPServerModel struct {
	DHCPPoolEnd   types.String             `tfsdk:"dhcp_pool_end" json:"dhcp_pool_end,optional"`
	DHCPPoolStart types.String             `tfsdk:"dhcp_pool_start" json:"dhcp_pool_start,optional"`
	DNSServer     types.String             `tfsdk:"dns_server" json:"dns_server,optional"`
	DNSServers    *[]types.String          `tfsdk:"dns_servers" json:"dns_servers,optional"`
	Reservations  *map[string]types.String `tfsdk:"reservations" json:"reservations,optional"`
}
