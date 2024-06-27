// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package teams_proxy_endpoint

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TeamsProxyEndpointResultEnvelope struct {
	Result TeamsProxyEndpointModel `json:"result,computed"`
}

type TeamsProxyEndpointModel struct {
	ID        types.String    `tfsdk:"id" json:"id,computed"`
	AccountID types.String    `tfsdk:"account_id" path:"account_id"`
	IPs       *[]types.String `tfsdk:"ips" json:"ips"`
	Name      types.String    `tfsdk:"name" json:"name"`
	CreatedAt types.String    `tfsdk:"created_at" json:"created_at,computed"`
	Subdomain types.String    `tfsdk:"subdomain" json:"subdomain,computed"`
	UpdatedAt types.String    `tfsdk:"updated_at" json:"updated_at,computed"`
}
