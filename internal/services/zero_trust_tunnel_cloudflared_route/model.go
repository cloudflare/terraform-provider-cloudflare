// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared_route

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustTunnelCloudflaredRouteResultEnvelope struct {
	Result ZeroTrustTunnelCloudflaredRouteModel `json:"result"`
}

type ZeroTrustTunnelCloudflaredRouteModel struct {
	ID               types.String      `tfsdk:"id" json:"id,computed"`
	AccountID        types.String      `tfsdk:"account_id" path:"account_id,required"`
	Network          types.String      `tfsdk:"network" json:"network,required"`
	TunnelID         types.String      `tfsdk:"tunnel_id" json:"tunnel_id,required"`
	Comment          types.String      `tfsdk:"comment" json:"comment,optional"`
	VirtualNetworkID types.String      `tfsdk:"virtual_network_id" json:"virtual_network_id,optional"`
	CreatedAt        timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DeletedAt        timetypes.RFC3339 `tfsdk:"deleted_at" json:"deleted_at,computed" format:"date-time"`
}

func (m ZeroTrustTunnelCloudflaredRouteModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustTunnelCloudflaredRouteModel) MarshalJSONForUpdate(state ZeroTrustTunnelCloudflaredRouteModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
