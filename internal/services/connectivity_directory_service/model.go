// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package connectivity_directory_service

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ConnectivityDirectoryServiceModel struct {
	AccountID types.String                           `tfsdk:"account_id" path:"account_id,required"`
	ServiceID types.String                           `tfsdk:"service_id" path:"service_id,optional"`
	Name      types.String                           `tfsdk:"name" json:"name,required,no_refresh"`
	Type      types.String                           `tfsdk:"type" json:"type,required,no_refresh"`
	Host      *ConnectivityDirectoryServiceHostModel `tfsdk:"host" json:"host,required,no_refresh"`
	HTTPPort  types.Int64                            `tfsdk:"http_port" json:"http_port,optional,no_refresh"`
	HTTPSPort types.Int64                            `tfsdk:"https_port" json:"https_port,optional,no_refresh"`
}

func (m ConnectivityDirectoryServiceModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ConnectivityDirectoryServiceModel) MarshalJSONForUpdate(state ConnectivityDirectoryServiceModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ConnectivityDirectoryServiceHostModel struct {
	Hostname        types.String         `tfsdk:"hostname" json:"hostname,optional"`
	IPV4            types.String         `tfsdk:"ipv4" json:"ipv4,optional"`
	IPV6            types.String         `tfsdk:"ipv6" json:"ipv6,optional"`
	Network         jsontypes.Normalized `tfsdk:"network" json:"network,optional"`
	ResolverNetwork jsontypes.Normalized `tfsdk:"resolver_network" json:"resolver_network,optional"`
}
