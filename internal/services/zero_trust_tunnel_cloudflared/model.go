// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustTunnelCloudflaredResultEnvelope struct {
	Result ZeroTrustTunnelCloudflaredModel `json:"result"`
}

type ZeroTrustTunnelCloudflaredModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	AccountID    types.String `tfsdk:"account_id" path:"account_id,required"`
	ConfigSrc    types.String `tfsdk:"config_src" json:"config_src,computed_optional"`
	Name         types.String `tfsdk:"name" json:"name,required"`
	TunnelSecret types.String `tfsdk:"tunnel_secret" json:"tunnel_secret,optional"`
}

func (m ZeroTrustTunnelCloudflaredModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustTunnelCloudflaredModel) MarshalJSONForUpdate(state ZeroTrustTunnelCloudflaredModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
