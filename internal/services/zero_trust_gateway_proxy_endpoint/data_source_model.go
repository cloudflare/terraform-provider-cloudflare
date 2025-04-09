// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_proxy_endpoint

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/zero_trust"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustGatewayProxyEndpointResultDataSourceEnvelope struct {
Result ZeroTrustGatewayProxyEndpointDataSourceModel `json:"result,computed"`
}

type ZeroTrustGatewayProxyEndpointDataSourceModel struct {
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
ProxyEndpointID types.String `tfsdk:"proxy_endpoint_id" path:"proxy_endpoint_id,required"`
CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
ID types.String `tfsdk:"id" json:"id,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
Subdomain types.String `tfsdk:"subdomain" json:"subdomain,computed"`
UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
IPs customfield.List[types.String] `tfsdk:"ips" json:"ips,computed"`
}

func (m *ZeroTrustGatewayProxyEndpointDataSourceModel) toReadParams(_ context.Context) (params zero_trust.GatewayProxyEndpointGetParams, diags diag.Diagnostics) {
  params = zero_trust.GatewayProxyEndpointGetParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  return
}
