// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_proxy_endpoint

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustGatewayProxyEndpointsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustGatewayProxyEndpointsResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustGatewayProxyEndpointsDataSourceModel struct {
	AccountID types.String                                                                      `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                                       `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustGatewayProxyEndpointsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustGatewayProxyEndpointsDataSourceModel) toListParams(_ context.Context) (params zero_trust.GatewayProxyEndpointListParams, diags diag.Diagnostics) {
	params = zero_trust.GatewayProxyEndpointListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustGatewayProxyEndpointsResultDataSourceModel struct {
	IPs       customfield.List[types.String] `tfsdk:"ips" json:"ips,computed"`
	Name      types.String                   `tfsdk:"name" json:"name,computed"`
	ID        types.String                   `tfsdk:"id" json:"id,computed"`
	CreatedAt timetypes.RFC3339              `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Kind      types.String                   `tfsdk:"kind" json:"kind,computed"`
	Subdomain types.String                   `tfsdk:"subdomain" json:"subdomain,computed"`
	UpdatedAt timetypes.RFC3339              `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}
