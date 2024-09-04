// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared_route

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustTunnelCloudflaredRouteResultEnvelope struct {
	Result ZeroTrustTunnelCloudflaredRouteModel `json:"result"`
}

type ZeroTrustTunnelCloudflaredRouteModel struct {
	ID               types.String      `tfsdk:"id" json:"id,computed"`
	AccountID        types.String      `tfsdk:"account_id" path:"account_id"`
	Comment          types.String      `tfsdk:"comment" json:"comment,computed_optional"`
	Network          types.String      `tfsdk:"network" json:"network,computed_optional"`
	TunnelID         types.String      `tfsdk:"tunnel_id" json:"tunnel_id,computed_optional"`
	VirtualNetworkID types.String      `tfsdk:"virtual_network_id" json:"virtual_network_id,computed_optional"`
	CreatedAt        timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DeletedAt        timetypes.RFC3339 `tfsdk:"deleted_at" json:"deleted_at,computed" format:"date-time"`
}
