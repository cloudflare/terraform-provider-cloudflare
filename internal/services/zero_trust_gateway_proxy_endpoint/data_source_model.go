// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_proxy_endpoint

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustGatewayProxyEndpointResultDataSourceEnvelope struct {
	Result ZeroTrustGatewayProxyEndpointDataSourceModel `json:"result,computed"`
}

type ZeroTrustGatewayProxyEndpointDataSourceModel struct {
	ID              types.String                                           `tfsdk:"id" path:"proxy_endpoint_id,computed"`
	ProxyEndpointID types.String                                           `tfsdk:"proxy_endpoint_id" path:"proxy_endpoint_id,optional"`
	AccountID       types.String                                           `tfsdk:"account_id" path:"account_id,required"`
	CreatedAt       timetypes.RFC3339                                      `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Kind            types.String                                           `tfsdk:"kind" json:"kind,computed"`
	Name            types.String                                           `tfsdk:"name" json:"name,computed"`
	Subdomain       types.String                                           `tfsdk:"subdomain" json:"subdomain,computed"`
	UpdatedAt       timetypes.RFC3339                                      `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	IPs             customfield.List[types.String]                         `tfsdk:"ips" json:"ips,computed"`
	Filter          *ZeroTrustGatewayProxyEndpointFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *ZeroTrustGatewayProxyEndpointDataSourceModel) toReadParams(_ context.Context) (params zero_trust.GatewayProxyEndpointGetParams, diags diag.Diagnostics) {
	params = zero_trust.GatewayProxyEndpointGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ZeroTrustGatewayProxyEndpointDataSourceModel) toListParams(_ context.Context) (params zero_trust.GatewayProxyEndpointListParams, diags diag.Diagnostics) {
	mFilterFilter := []interface{}{}
	if m.Filter.Filter != nil {
		for _, item := range *m.Filter.Filter {
			mFilterFilter = append(mFilterFilter, item.ValueString())
		}
	}

	params = zero_trust.GatewayProxyEndpointListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
		Filter:    cloudflare.F(mFilterFilter),
	}

	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(zero_trust.GatewayProxyEndpointListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.OrderBy.IsNull() {
		params.OrderBy = cloudflare.F(zero_trust.GatewayProxyEndpointListParamsOrderBy(m.Filter.OrderBy.ValueString()))
	}
	if !m.Filter.Search.IsNull() {
		params.Search = cloudflare.F(m.Filter.Search.ValueString())
	}

	return
}

type ZeroTrustGatewayProxyEndpointFindOneByDataSourceModel struct {
	Direction types.String    `tfsdk:"direction" query:"direction,optional"`
	Filter    *[]types.String `tfsdk:"filter" query:"filter,optional"`
	OrderBy   types.String    `tfsdk:"order_by" query:"order_by,optional"`
	Search    types.String    `tfsdk:"search" query:"search,optional"`
}
