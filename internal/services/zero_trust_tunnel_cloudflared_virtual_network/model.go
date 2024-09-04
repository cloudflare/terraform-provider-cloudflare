// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared_virtual_network

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustTunnelCloudflaredVirtualNetworkResultEnvelope struct {
	Result ZeroTrustTunnelCloudflaredVirtualNetworkModel `json:"result"`
}

type ZeroTrustTunnelCloudflaredVirtualNetworkModel struct {
	ID               types.String      `tfsdk:"id" json:"id,computed"`
	AccountID        types.String      `tfsdk:"account_id" path:"account_id,required"`
	IsDefault        types.Bool        `tfsdk:"is_default" json:"is_default,optional"`
	Comment          types.String      `tfsdk:"comment" json:"comment,computed_optional"`
	IsDefaultNetwork types.Bool        `tfsdk:"is_default_network" json:"is_default_network,computed_optional"`
	Name             types.String      `tfsdk:"name" json:"name,computed_optional"`
	CreatedAt        timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DeletedAt        timetypes.RFC3339 `tfsdk:"deleted_at" json:"deleted_at,computed" format:"date-time"`
}
