// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package connectivity_directory_service

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/connectivity"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ConnectivityDirectoryServiceResultDataSourceEnvelope struct {
	Result ConnectivityDirectoryServiceDataSourceModel `json:"result,computed"`
}

type ConnectivityDirectoryServiceDataSourceModel struct {
	ID        types.String                                                              `tfsdk:"id" path:"service_id,computed"`
	ServiceID types.String                                                              `tfsdk:"service_id" path:"service_id,computed_optional"`
	AccountID types.String                                                              `tfsdk:"account_id" path:"account_id,required"`
	CreatedAt timetypes.RFC3339                                                         `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	HTTPPort  types.Int64                                                               `tfsdk:"http_port" json:"http_port,computed"`
	HTTPSPort types.Int64                                                               `tfsdk:"https_port" json:"https_port,computed"`
	Name      types.String                                                              `tfsdk:"name" json:"name,computed"`
	Type      types.String                                                              `tfsdk:"type" json:"type,computed"`
	UpdatedAt timetypes.RFC3339                                                         `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Host      customfield.NestedObject[ConnectivityDirectoryServiceHostDataSourceModel] `tfsdk:"host" json:"host,computed"`
	Filter    *ConnectivityDirectoryServiceFindOneByDataSourceModel                     `tfsdk:"filter"`
}

func (m *ConnectivityDirectoryServiceDataSourceModel) toReadParams(_ context.Context) (params connectivity.DirectoryServiceGetParams, diags diag.Diagnostics) {
	params = connectivity.DirectoryServiceGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ConnectivityDirectoryServiceDataSourceModel) toListParams(_ context.Context) (params connectivity.DirectoryServiceListParams, diags diag.Diagnostics) {
	params = connectivity.DirectoryServiceListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Filter.Type.IsNull() {
		params.Type = cloudflare.F(connectivity.DirectoryServiceListParamsType(m.Filter.Type.ValueString()))
	}

	return
}

type ConnectivityDirectoryServiceHostDataSourceModel struct {
	IPV4            types.String                                                                             `tfsdk:"ipv4" json:"ipv4,computed"`
	Network         customfield.NestedObject[ConnectivityDirectoryServiceHostNetworkDataSourceModel]         `tfsdk:"network" json:"network,computed"`
	IPV6            types.String                                                                             `tfsdk:"ipv6" json:"ipv6,computed"`
	Hostname        types.String                                                                             `tfsdk:"hostname" json:"hostname,computed"`
	ResolverNetwork customfield.NestedObject[ConnectivityDirectoryServiceHostResolverNetworkDataSourceModel] `tfsdk:"resolver_network" json:"resolver_network,computed"`
}

type ConnectivityDirectoryServiceHostNetworkDataSourceModel struct {
	TunnelID types.String `tfsdk:"tunnel_id" json:"tunnel_id,computed"`
}

type ConnectivityDirectoryServiceHostResolverNetworkDataSourceModel struct {
	TunnelID    types.String                   `tfsdk:"tunnel_id" json:"tunnel_id,computed"`
	ResolverIPs customfield.List[types.String] `tfsdk:"resolver_ips" json:"resolver_ips,computed"`
}

type ConnectivityDirectoryServiceFindOneByDataSourceModel struct {
	Type types.String `tfsdk:"type" query:"type,optional"`
}
