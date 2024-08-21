// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_proxy_endpoint

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustGatewayProxyEndpointResultDataSourceEnvelope struct {
	Result ZeroTrustGatewayProxyEndpointDataSourceModel `json:"result,computed"`
}

type ZeroTrustGatewayProxyEndpointDataSourceModel struct {
	AccountID       types.String `tfsdk:"account_id" path:"account_id"`
	ProxyEndpointID types.String `tfsdk:"proxy_endpoint_id" path:"proxy_endpoint_id"`
}

func (m *ZeroTrustGatewayProxyEndpointDataSourceModel) toReadParams() (params zero_trust.GatewayProxyEndpointGetParams, diags diag.Diagnostics) {
	params = zero_trust.GatewayProxyEndpointGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
