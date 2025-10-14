// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_connectivity_directory_service

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustConnectivityDirectoryServiceResultEnvelope struct {
	Result ZeroTrustConnectivityDirectoryServiceModel `json:"result"`
}

type ZeroTrustConnectivityDirectoryServiceModel struct {
	ID        types.String                                    `tfsdk:"id" json:"-,computed"`
	ServiceID types.String                                    `tfsdk:"service_id" json:"service_id,computed"`
	AccountID types.String                                    `tfsdk:"account_id" path:"account_id,required"`
	Name      types.String                                    `tfsdk:"name" json:"name,required"`
	Type      types.String                                    `tfsdk:"type" json:"type,required"`
	Host      *ZeroTrustConnectivityDirectoryServiceHostModel `tfsdk:"host" json:"host,required"`
	HTTPPort  types.Int64                                     `tfsdk:"http_port" json:"http_port,optional"`
	HTTPSPort types.Int64                                     `tfsdk:"https_port" json:"https_port,optional"`
	CreatedAt timetypes.RFC3339                               `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UpdatedAt timetypes.RFC3339                               `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m ZeroTrustConnectivityDirectoryServiceModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustConnectivityDirectoryServiceModel) MarshalJSONForUpdate(state ZeroTrustConnectivityDirectoryServiceModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustConnectivityDirectoryServiceHostModel struct {
	IPV4            types.String                                                   `tfsdk:"ipv4" json:"ipv4,optional"`
	Network         *ZeroTrustConnectivityDirectoryServiceHostNetworkModel         `tfsdk:"network" json:"network,optional"`
	IPV6            types.String                                                   `tfsdk:"ipv6" json:"ipv6,optional"`
	Hostname        types.String                                                   `tfsdk:"hostname" json:"hostname,optional"`
	ResolverNetwork *ZeroTrustConnectivityDirectoryServiceHostResolverNetworkModel `tfsdk:"resolver_network" json:"resolver_network,optional"`
}

type ZeroTrustConnectivityDirectoryServiceHostNetworkModel struct {
	TunnelID types.String `tfsdk:"tunnel_id" json:"tunnel_id,required"`
}

type ZeroTrustConnectivityDirectoryServiceHostResolverNetworkModel struct {
	TunnelID    types.String    `tfsdk:"tunnel_id" json:"tunnel_id,required"`
	ResolverIPs *[]types.String `tfsdk:"resolver_ips" json:"resolver_ips,optional"`
}
