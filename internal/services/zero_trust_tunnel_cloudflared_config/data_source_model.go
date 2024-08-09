// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared_config

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustTunnelCloudflaredConfigResultDataSourceEnvelope struct {
	Result ZeroTrustTunnelCloudflaredConfigDataSourceModel `json:"result,computed"`
}

type ZeroTrustTunnelCloudflaredConfigDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	TunnelID  types.String `tfsdk:"tunnel_id" path:"tunnel_id"`
}
