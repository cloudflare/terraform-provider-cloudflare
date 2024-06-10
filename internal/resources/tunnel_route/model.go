// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tunnel_route

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TunnelRouteResultEnvelope struct {
	Result TunnelRouteModel `json:"result,computed"`
}

type TunnelRouteModel struct {
	ID               types.String `tfsdk:"id" json:"id,computed"`
	AccountID        types.String `tfsdk:"account_id" path:"account_id"`
	Network          types.String `tfsdk:"network" json:"network"`
	Comment          types.String `tfsdk:"comment" json:"comment"`
	VirtualNetworkID types.String `tfsdk:"virtual_network_id" json:"virtual_network_id"`
}
