// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_proxy_endpoint

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustGatewayProxyEndpointsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustGatewayProxyEndpointsResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustGatewayProxyEndpointsDataSourceModel struct {
	AccountID types.String                                                                      `tfsdk:"account_id" path:"account_id,optional"`
	Direction types.String                                                                      `tfsdk:"direction" query:"direction,optional"`
	OrderBy   types.String                                                                      `tfsdk:"order_by" query:"order_by,optional"`
	Search    types.String                                                                      `tfsdk:"search" query:"search,optional"`
	Filter    *[]jsontypes.Normalized                                                           `tfsdk:"filter" query:"filter,optional"`
	MaxItems  types.Int64                                                                       `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustGatewayProxyEndpointsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustGatewayProxyEndpointsDataSourceModel) toListParams(_ context.Context) (params zero_trust.GatewayProxyEndpointListParams, diags diag.Diagnostics) {
	mFilter := []interface{}{}
	if m.Filter != nil {
		for _, item := range *m.Filter {
			mFilter = append(mFilter, item.ValueString())
		}
	}

	params = zero_trust.GatewayProxyEndpointListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
		Filter:    cloudflare.F(mFilter),
	}

	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(zero_trust.GatewayProxyEndpointListParamsDirection(m.Direction.ValueString()))
	}
	if !m.OrderBy.IsNull() {
		params.OrderBy = cloudflare.F(zero_trust.GatewayProxyEndpointListParamsOrderBy(m.OrderBy.ValueString()))
	}
	if !m.Search.IsNull() {
		params.Search = cloudflare.F(m.Search.ValueString())
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
