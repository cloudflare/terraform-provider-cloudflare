// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tunnel_virtual_network

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TunnelVirtualNetworkResultEnvelope struct {
	Result TunnelVirtualNetworkModel `json:"result,computed"`
}

type TunnelVirtualNetworksResultDataSourceEnvelope struct {
	Result TunnelVirtualNetworksDataSourceModel `json:"result,computed"`
}

type TunnelVirtualNetworkModel struct {
	AccountID        types.String `tfsdk:"account_id" path:"account_id"`
	VirtualNetworkID types.String `tfsdk:"virtual_network_id" path:"virtual_network_id"`
	Name             types.String `tfsdk:"name" json:"name"`
	Comment          types.String `tfsdk:"comment" json:"comment"`
	IsDefault        types.Bool   `tfsdk:"is_default" json:"is_default"`
	IsDefaultNetwork types.Bool   `tfsdk:"is_default_network" json:"is_default_network"`
}

type TunnelVirtualNetworksDataSourceModel struct {
}
