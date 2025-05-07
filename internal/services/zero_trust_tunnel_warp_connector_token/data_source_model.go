// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_warp_connector_token

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustTunnelWARPConnectorTokenResultDataSourceEnvelope struct {
	Result ZeroTrustTunnelWARPConnectorTokenDataSourceModel `json:"result,computed"`
}

type ZeroTrustTunnelWARPConnectorTokenDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
	TunnelID  types.String `tfsdk:"tunnel_id" path:"tunnel_id,required"`
}

func (m *ZeroTrustTunnelWARPConnectorTokenDataSourceModel) toReadParams(_ context.Context) (params zero_trust.TunnelWARPConnectorTokenGetParams, diags diag.Diagnostics) {
	params = zero_trust.TunnelWARPConnectorTokenGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
