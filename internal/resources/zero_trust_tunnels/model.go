// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnels

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustTunnelsResultEnvelope struct {
	Result ZeroTrustTunnelsModel `json:"result,computed"`
}

type ZeroTrustTunnelsModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	AccountID    types.String `tfsdk:"account_id" path:"account_id"`
	Name         types.String `tfsdk:"name" json:"name"`
	TunnelSecret types.String `tfsdk:"tunnel_secret" json:"tunnel_secret"`
}
