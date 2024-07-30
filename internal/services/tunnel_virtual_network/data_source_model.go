// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tunnel_virtual_network

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TunnelVirtualNetworkResultListDataSourceEnvelope struct {
	Result *[]*TunnelVirtualNetworkDataSourceModel `json:"result,computed"`
}

type TunnelVirtualNetworkDataSourceModel struct {
	ID               types.String                                  `tfsdk:"id" json:"id"`
	Comment          types.String                                  `tfsdk:"comment" json:"comment"`
	CreatedAt        timetypes.RFC3339                             `tfsdk:"created_at" json:"created_at"`
	IsDefaultNetwork types.Bool                                    `tfsdk:"is_default_network" json:"is_default_network"`
	Name             types.String                                  `tfsdk:"name" json:"name"`
	DeletedAt        timetypes.RFC3339                             `tfsdk:"deleted_at" json:"deleted_at"`
	Filter           *TunnelVirtualNetworkFindOneByDataSourceModel `tfsdk:"filter"`
}

type TunnelVirtualNetworkFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	ID        types.String `tfsdk:"id" query:"id"`
	IsDefault types.Bool   `tfsdk:"is_default" query:"is_default"`
	IsDeleted types.Bool   `tfsdk:"is_deleted" query:"is_deleted"`
	Name      types.String `tfsdk:"name" query:"name"`
}
