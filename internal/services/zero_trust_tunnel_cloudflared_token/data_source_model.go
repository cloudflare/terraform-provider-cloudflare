// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared_token

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/zero_trust"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustTunnelCloudflaredTokenResultDataSourceEnvelope struct {
Result ZeroTrustTunnelCloudflaredTokenDataSourceModel `json:"result,computed"`
}

type ZeroTrustTunnelCloudflaredTokenDataSourceModel struct {
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
TunnelID types.String `tfsdk:"tunnel_id" path:"tunnel_id,required"`
}

func (m *ZeroTrustTunnelCloudflaredTokenDataSourceModel) toReadParams(_ context.Context) (params zero_trust.TunnelCloudflaredTokenGetParams, diags diag.Diagnostics) {
  params = zero_trust.TunnelCloudflaredTokenGetParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  return
}
