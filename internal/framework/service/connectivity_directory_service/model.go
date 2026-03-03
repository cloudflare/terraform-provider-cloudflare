package connectivity_directory_service

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ConnectivityDirectoryServiceModel struct {
	ID        types.String                           `tfsdk:"id"`
	AccountID types.String                           `tfsdk:"account_id"`
	ServiceID types.String                           `tfsdk:"service_id"`
	Name      types.String                           `tfsdk:"name"`
	Type      types.String                           `tfsdk:"type"`
	Host      *ConnectivityDirectoryServiceHostModel `tfsdk:"host"`
	HTTPPort  types.Int64                            `tfsdk:"http_port"`
	HTTPSPort types.Int64                            `tfsdk:"https_port"`
	CreatedAt types.String                           `tfsdk:"created_at"`
	UpdatedAt types.String                           `tfsdk:"updated_at"`
}

type ConnectivityDirectoryServiceHostModel struct {
	IPV4            types.String                                          `tfsdk:"ipv4"`
	IPV6            types.String                                          `tfsdk:"ipv6"`
	Hostname        types.String                                          `tfsdk:"hostname"`
	Network         *ConnectivityDirectoryServiceHostNetworkModel         `tfsdk:"network"`
	ResolverNetwork *ConnectivityDirectoryServiceHostResolverNetworkModel `tfsdk:"resolver_network"`
}

type ConnectivityDirectoryServiceHostNetworkModel struct {
	TunnelID types.String `tfsdk:"tunnel_id"`
}

type ConnectivityDirectoryServiceHostResolverNetworkModel struct {
	TunnelID    types.String `tfsdk:"tunnel_id"`
	ResolverIPs types.List   `tfsdk:"resolver_ips"`
}
