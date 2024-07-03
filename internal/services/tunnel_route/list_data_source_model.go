// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tunnel_route

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TunnelRoutesResultListDataSourceEnvelope struct {
	Result *[]*TunnelRoutesItemsDataSourceModel `json:"result,computed"`
}

type TunnelRoutesDataSourceModel struct {
	AccountID        types.String                         `tfsdk:"account_id" path:"account_id"`
	Comment          types.String                         `tfsdk:"comment" query:"comment"`
	ExistedAt        types.String                         `tfsdk:"existed_at" query:"existed_at"`
	IsDeleted        types.Bool                           `tfsdk:"is_deleted" query:"is_deleted"`
	NetworkSubset    types.String                         `tfsdk:"network_subset" query:"network_subset"`
	NetworkSuperset  types.String                         `tfsdk:"network_superset" query:"network_superset"`
	Page             types.Float64                        `tfsdk:"page" query:"page"`
	PerPage          types.Float64                        `tfsdk:"per_page" query:"per_page"`
	RouteID          types.String                         `tfsdk:"route_id" query:"route_id"`
	TunTypes         types.String                         `tfsdk:"tun_types" query:"tun_types"`
	TunnelID         types.String                         `tfsdk:"tunnel_id" query:"tunnel_id"`
	VirtualNetworkID types.String                         `tfsdk:"virtual_network_id" query:"virtual_network_id"`
	MaxItems         types.Int64                          `tfsdk:"max_items"`
	Items            *[]*TunnelRoutesItemsDataSourceModel `tfsdk:"items"`
}

type TunnelRoutesItemsDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	Comment            types.String `tfsdk:"comment" json:"comment,computed"`
	CreatedAt          types.String `tfsdk:"created_at" json:"created_at,computed"`
	DeletedAt          types.String `tfsdk:"deleted_at" json:"deleted_at,computed"`
	Network            types.String `tfsdk:"network" json:"network,computed"`
	TunType            types.String `tfsdk:"tun_type" json:"tun_type,computed"`
	TunnelID           types.String `tfsdk:"tunnel_id" json:"tunnel_id,computed"`
	TunnelName         types.String `tfsdk:"tunnel_name" json:"tunnel_name,computed"`
	VirtualNetworkID   types.String `tfsdk:"virtual_network_id" json:"virtual_network_id,computed"`
	VirtualNetworkName types.String `tfsdk:"virtual_network_name" json:"virtual_network_name,computed"`
}
