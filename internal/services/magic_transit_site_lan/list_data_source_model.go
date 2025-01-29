// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_site_lan

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/magic_transit"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicTransitSiteLANsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[MagicTransitSiteLANsResultDataSourceModel] `json:"result,computed"`
}

type MagicTransitSiteLANsDataSourceModel struct {
	AccountID types.String                                                            `tfsdk:"account_id" path:"account_id,required"`
	SiteID    types.String                                                            `tfsdk:"site_id" path:"site_id,required"`
	MaxItems  types.Int64                                                             `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[MagicTransitSiteLANsResultDataSourceModel] `tfsdk:"result"`
}

func (m *MagicTransitSiteLANsDataSourceModel) toListParams(_ context.Context) (params magic_transit.SiteLANListParams, diags diag.Diagnostics) {
	params = magic_transit.SiteLANListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type MagicTransitSiteLANsResultDataSourceModel struct {
	ID               types.String                                                                   `tfsdk:"id" json:"id,computed"`
	HaLink           types.Bool                                                                     `tfsdk:"ha_link" json:"ha_link,computed"`
	Name             types.String                                                                   `tfsdk:"name" json:"name,computed"`
	Nat              customfield.NestedObject[MagicTransitSiteLANsNatDataSourceModel]               `tfsdk:"nat" json:"nat,computed"`
	Physport         types.Int64                                                                    `tfsdk:"physport" json:"physport,computed"`
	RoutedSubnets    customfield.NestedObjectList[MagicTransitSiteLANsRoutedSubnetsDataSourceModel] `tfsdk:"routed_subnets" json:"routed_subnets,computed"`
	SiteID           types.String                                                                   `tfsdk:"site_id" json:"site_id,computed"`
	StaticAddressing customfield.NestedObject[MagicTransitSiteLANsStaticAddressingDataSourceModel]  `tfsdk:"static_addressing" json:"static_addressing,computed"`
	VlanTag          types.Int64                                                                    `tfsdk:"vlan_tag" json:"vlan_tag,computed"`
}

type MagicTransitSiteLANsNatDataSourceModel struct {
	StaticPrefix types.String `tfsdk:"static_prefix" json:"static_prefix,computed"`
}

type MagicTransitSiteLANsRoutedSubnetsDataSourceModel struct {
	NextHop types.String                                                                  `tfsdk:"next_hop" json:"next_hop,computed"`
	Prefix  types.String                                                                  `tfsdk:"prefix" json:"prefix,computed"`
	Nat     customfield.NestedObject[MagicTransitSiteLANsRoutedSubnetsNatDataSourceModel] `tfsdk:"nat" json:"nat,computed"`
}

type MagicTransitSiteLANsRoutedSubnetsNatDataSourceModel struct {
	StaticPrefix types.String `tfsdk:"static_prefix" json:"static_prefix,computed"`
}

type MagicTransitSiteLANsStaticAddressingDataSourceModel struct {
	Address          types.String                                                                            `tfsdk:"address" json:"address,computed"`
	DHCPRelay        customfield.NestedObject[MagicTransitSiteLANsStaticAddressingDHCPRelayDataSourceModel]  `tfsdk:"dhcp_relay" json:"dhcp_relay,computed"`
	DHCPServer       customfield.NestedObject[MagicTransitSiteLANsStaticAddressingDHCPServerDataSourceModel] `tfsdk:"dhcp_server" json:"dhcp_server,computed"`
	SecondaryAddress types.String                                                                            `tfsdk:"secondary_address" json:"secondary_address,computed"`
	VirtualAddress   types.String                                                                            `tfsdk:"virtual_address" json:"virtual_address,computed"`
}

type MagicTransitSiteLANsStaticAddressingDHCPRelayDataSourceModel struct {
	ServerAddresses customfield.List[types.String] `tfsdk:"server_addresses" json:"server_addresses,computed"`
}

type MagicTransitSiteLANsStaticAddressingDHCPServerDataSourceModel struct {
	DHCPPoolEnd   types.String                  `tfsdk:"dhcp_pool_end" json:"dhcp_pool_end,computed"`
	DHCPPoolStart types.String                  `tfsdk:"dhcp_pool_start" json:"dhcp_pool_start,computed"`
	DNSServer     types.String                  `tfsdk:"dns_server" json:"dns_server,computed"`
	Reservations  customfield.Map[types.String] `tfsdk:"reservations" json:"reservations,computed"`
}
