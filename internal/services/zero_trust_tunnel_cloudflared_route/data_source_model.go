// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared_route

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustTunnelCloudflaredRouteResultListDataSourceEnvelope struct {
	Result *[]*ZeroTrustTunnelCloudflaredRouteDataSourceModel `json:"result,computed"`
}

type ZeroTrustTunnelCloudflaredRouteDataSourceModel struct {
	Comment            types.String                                             `tfsdk:"comment" json:"comment"`
	CreatedAt          timetypes.RFC3339                                        `tfsdk:"created_at" json:"created_at"`
	DeletedAt          timetypes.RFC3339                                        `tfsdk:"deleted_at" json:"deleted_at"`
	ID                 types.String                                             `tfsdk:"id" json:"id"`
	Network            types.String                                             `tfsdk:"network" json:"network"`
	TunType            types.String                                             `tfsdk:"tun_type" json:"tun_type"`
	TunnelID           types.String                                             `tfsdk:"tunnel_id" json:"tunnel_id"`
	TunnelName         types.String                                             `tfsdk:"tunnel_name" json:"tunnel_name"`
	VirtualNetworkID   types.String                                             `tfsdk:"virtual_network_id" json:"virtual_network_id"`
	VirtualNetworkName types.String                                             `tfsdk:"virtual_network_name" json:"virtual_network_name"`
	Filter             *ZeroTrustTunnelCloudflaredRouteFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *ZeroTrustTunnelCloudflaredRouteDataSourceModel) toListParams() (params zero_trust.NetworkRouteListParams, diags diag.Diagnostics) {
	mFilterExistedAt, errs := m.Filter.ExistedAt.ValueRFC3339Time()
	diags.Append(errs...)

	params = zero_trust.NetworkRouteListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	if !m.Filter.Comment.IsNull() {
		params.Comment = cloudflare.F(m.Filter.Comment.ValueString())
	}
	if !m.Filter.ExistedAt.IsNull() {
		params.ExistedAt = cloudflare.F(mFilterExistedAt)
	}
	if !m.Filter.IsDeleted.IsNull() {
		params.IsDeleted = cloudflare.F(m.Filter.IsDeleted.ValueBool())
	}
	if !m.Filter.NetworkSubset.IsNull() {
		params.NetworkSubset = cloudflare.F(m.Filter.NetworkSubset.ValueString())
	}
	if !m.Filter.NetworkSuperset.IsNull() {
		params.NetworkSuperset = cloudflare.F(m.Filter.NetworkSuperset.ValueString())
	}
	if !m.Filter.RouteID.IsNull() {
		params.RouteID = cloudflare.F(m.Filter.RouteID.ValueString())
	}
	if !m.Filter.TunTypes.IsNull() {
		params.TunTypes = cloudflare.F(m.Filter.TunTypes.ValueString())
	}
	if !m.Filter.TunnelID.IsNull() {
		params.TunnelID = cloudflare.F(m.Filter.TunnelID.ValueString())
	}
	if !m.Filter.VirtualNetworkID.IsNull() {
		params.VirtualNetworkID = cloudflare.F(m.Filter.VirtualNetworkID.ValueString())
	}

	return
}

type ZeroTrustTunnelCloudflaredRouteFindOneByDataSourceModel struct {
	AccountID        types.String      `tfsdk:"account_id" path:"account_id"`
	Comment          types.String      `tfsdk:"comment" query:"comment"`
	ExistedAt        timetypes.RFC3339 `tfsdk:"existed_at" query:"existed_at"`
	IsDeleted        types.Bool        `tfsdk:"is_deleted" query:"is_deleted"`
	NetworkSubset    types.String      `tfsdk:"network_subset" query:"network_subset"`
	NetworkSuperset  types.String      `tfsdk:"network_superset" query:"network_superset"`
	RouteID          types.String      `tfsdk:"route_id" query:"route_id"`
	TunTypes         types.String      `tfsdk:"tun_types" query:"tun_types"`
	TunnelID         types.String      `tfsdk:"tunnel_id" query:"tunnel_id"`
	VirtualNetworkID types.String      `tfsdk:"virtual_network_id" query:"virtual_network_id"`
}
