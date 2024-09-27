// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_site_lan

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicTransitSiteLANResultEnvelope struct {
	Result MagicTransitSiteLANModel `json:"result"`
}

type MagicTransitSiteLANModel struct {
	AccountID        types.String                                                        `tfsdk:"account_id" path:"account_id,required"`
	SiteID           types.String                                                        `tfsdk:"site_id" path:"site_id,required"`
	LANID            types.String                                                        `tfsdk:"lan_id" path:"lan_id,optional"`
	HaLink           types.Bool                                                          `tfsdk:"ha_link" json:"ha_link,computed_optional"`
	Physport         types.Int64                                                         `tfsdk:"physport" json:"physport,required"`
	VlanTag          types.Int64                                                         `tfsdk:"vlan_tag" json:"vlan_tag,required"`
	Name             types.String                                                        `tfsdk:"name" json:"name,computed_optional"`
	Nat              customfield.NestedObject[MagicTransitSiteLANNatModel]               `tfsdk:"nat" json:"nat,computed_optional"`
	RoutedSubnets    customfield.NestedObjectList[MagicTransitSiteLANRoutedSubnetsModel] `tfsdk:"routed_subnets" json:"routed_subnets,computed_optional"`
	StaticAddressing customfield.NestedObject[MagicTransitSiteLANStaticAddressingModel]  `tfsdk:"static_addressing" json:"static_addressing,computed_optional"`
	ID               types.String                                                        `tfsdk:"id" json:"id,computed"`
}

type MagicTransitSiteLANNatModel struct {
	StaticPrefix types.String `tfsdk:"static_prefix" json:"static_prefix,computed_optional"`
}

type MagicTransitSiteLANRoutedSubnetsModel struct {
	NextHop types.String                                                       `tfsdk:"next_hop" json:"next_hop,required"`
	Prefix  types.String                                                       `tfsdk:"prefix" json:"prefix,required"`
	Nat     customfield.NestedObject[MagicTransitSiteLANRoutedSubnetsNatModel] `tfsdk:"nat" json:"nat,computed_optional"`
}

type MagicTransitSiteLANRoutedSubnetsNatModel struct {
	StaticPrefix types.String `tfsdk:"static_prefix" json:"static_prefix,computed_optional"`
}

type MagicTransitSiteLANStaticAddressingModel struct {
	Address          types.String                                                                 `tfsdk:"address" json:"address,required"`
	DHCPRelay        customfield.NestedObject[MagicTransitSiteLANStaticAddressingDHCPRelayModel]  `tfsdk:"dhcp_relay" json:"dhcp_relay,computed_optional"`
	DHCPServer       customfield.NestedObject[MagicTransitSiteLANStaticAddressingDHCPServerModel] `tfsdk:"dhcp_server" json:"dhcp_server,computed_optional"`
	SecondaryAddress types.String                                                                 `tfsdk:"secondary_address" json:"secondary_address,computed_optional"`
	VirtualAddress   types.String                                                                 `tfsdk:"virtual_address" json:"virtual_address,computed_optional"`
}

type MagicTransitSiteLANStaticAddressingDHCPRelayModel struct {
	ServerAddresses customfield.List[types.String] `tfsdk:"server_addresses" json:"server_addresses,computed_optional"`
}

type MagicTransitSiteLANStaticAddressingDHCPServerModel struct {
	DHCPPoolEnd   types.String                  `tfsdk:"dhcp_pool_end" json:"dhcp_pool_end,computed_optional"`
	DHCPPoolStart types.String                  `tfsdk:"dhcp_pool_start" json:"dhcp_pool_start,computed_optional"`
	DNSServer     types.String                  `tfsdk:"dns_server" json:"dns_server,computed_optional"`
	Reservations  customfield.Map[types.String] `tfsdk:"reservations" json:"reservations,computed_optional"`
}
