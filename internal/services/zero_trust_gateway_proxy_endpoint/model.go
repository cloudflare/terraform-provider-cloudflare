// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_proxy_endpoint

import (
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustGatewayProxyEndpointResultEnvelope struct {
Result ZeroTrustGatewayProxyEndpointModel `json:"result"`
}

type ZeroTrustGatewayProxyEndpointModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
Name types.String `tfsdk:"name" json:"name,required"`
IPs *[]types.String `tfsdk:"ips" json:"ips,required"`
CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
Subdomain types.String `tfsdk:"subdomain" json:"subdomain,computed"`
UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m ZeroTrustGatewayProxyEndpointModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m ZeroTrustGatewayProxyEndpointModel) MarshalJSONForUpdate(state ZeroTrustGatewayProxyEndpointModel) (data []byte, err error) {
  return apijson.MarshalForPatch(m, state)
}
