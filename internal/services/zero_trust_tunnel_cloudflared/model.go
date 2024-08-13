// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustTunnelCloudflaredResultEnvelope struct {
	Result ZeroTrustTunnelCloudflaredModel `json:"result,computed"`
}

type ZeroTrustTunnelCloudflaredModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	AccountID    types.String `tfsdk:"account_id" path:"account_id"`
	TunnelID     types.String `tfsdk:"tunnel_id" path:"tunnel_id"`
	Name         types.String `tfsdk:"name" json:"name"`
	TunnelSecret types.String `tfsdk:"tunnel_secret" json:"tunnel_secret"`
}
