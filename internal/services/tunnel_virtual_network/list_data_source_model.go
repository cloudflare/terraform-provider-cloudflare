// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tunnel_virtual_network

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TunnelVirtualNetworksResultListDataSourceEnvelope struct {
	Result *[]*TunnelVirtualNetworksItemsDataSourceModel `json:"result,computed"`
}

type TunnelVirtualNetworksDataSourceModel struct {
	AccountID types.String                                  `tfsdk:"account_id" path:"account_id"`
	ID        types.String                                  `tfsdk:"id" query:"id"`
	IsDefault types.Bool                                    `tfsdk:"is_default" query:"is_default"`
	IsDeleted types.Bool                                    `tfsdk:"is_deleted" query:"is_deleted"`
	Name      types.String                                  `tfsdk:"name" query:"name"`
	MaxItems  types.Int64                                   `tfsdk:"max_items"`
	Items     *[]*TunnelVirtualNetworksItemsDataSourceModel `tfsdk:"items"`
}

type TunnelVirtualNetworksItemsDataSourceModel struct {
	ID               types.String      `tfsdk:"id" json:"id,computed"`
	Comment          types.String      `tfsdk:"comment" json:"comment,computed"`
	CreatedAt        timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed"`
	IsDefaultNetwork types.Bool        `tfsdk:"is_default_network" json:"is_default_network,computed"`
	Name             types.String      `tfsdk:"name" json:"name,computed"`
	DeletedAt        timetypes.RFC3339 `tfsdk:"deleted_at" json:"deleted_at"`
}
