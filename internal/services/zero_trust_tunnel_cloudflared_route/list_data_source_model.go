// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared_route

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustTunnelCloudflaredRoutesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustTunnelCloudflaredRoutesResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustTunnelCloudflaredRoutesDataSourceModel struct {
	AccountID        types.String                                                                        `tfsdk:"account_id" path:"account_id,required"`
	Comment          types.String                                                                        `tfsdk:"comment" query:"comment,optional"`
	ExistedAt        timetypes.RFC3339                                                                   `tfsdk:"existed_at" query:"existed_at,optional" format:"date-time"`
	IsDeleted        types.Bool                                                                          `tfsdk:"is_deleted" query:"is_deleted,optional"`
	NetworkSubset    types.String                                                                        `tfsdk:"network_subset" query:"network_subset,optional"`
	NetworkSuperset  types.String                                                                        `tfsdk:"network_superset" query:"network_superset,optional"`
	RouteID          types.String                                                                        `tfsdk:"route_id" query:"route_id,optional"`
	TunTypes         types.String                                                                        `tfsdk:"tun_types" query:"tun_types,optional"`
	TunnelID         types.String                                                                        `tfsdk:"tunnel_id" query:"tunnel_id,optional"`
	VirtualNetworkID types.String                                                                        `tfsdk:"virtual_network_id" query:"virtual_network_id,optional"`
	MaxItems         types.Int64                                                                         `tfsdk:"max_items"`
	Result           customfield.NestedObjectList[ZeroTrustTunnelCloudflaredRoutesResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustTunnelCloudflaredRoutesDataSourceModel) toListParams(_ context.Context) (params zero_trust.NetworkRouteListParams, diags diag.Diagnostics) {
	mExistedAt, errs := m.ExistedAt.ValueRFC3339Time()
	diags.Append(errs...)

	params = zero_trust.NetworkRouteListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Comment.IsNull() {
		params.Comment = cloudflare.F(m.Comment.ValueString())
	}
	if !m.ExistedAt.IsNull() {
		params.ExistedAt = cloudflare.F(mExistedAt)
	}
	if !m.IsDeleted.IsNull() {
		params.IsDeleted = cloudflare.F(m.IsDeleted.ValueBool())
	}
	if !m.NetworkSubset.IsNull() {
		params.NetworkSubset = cloudflare.F(m.NetworkSubset.ValueString())
	}
	if !m.NetworkSuperset.IsNull() {
		params.NetworkSuperset = cloudflare.F(m.NetworkSuperset.ValueString())
	}
	if !m.RouteID.IsNull() {
		params.RouteID = cloudflare.F(m.RouteID.ValueString())
	}
	if !m.TunTypes.IsNull() {
		params.TunTypes = cloudflare.F(m.TunTypes.ValueString())
	}
	if !m.TunnelID.IsNull() {
		params.TunnelID = cloudflare.F(m.TunnelID.ValueString())
	}
	if !m.VirtualNetworkID.IsNull() {
		params.VirtualNetworkID = cloudflare.F(m.VirtualNetworkID.ValueString())
	}

	return
}

type ZeroTrustTunnelCloudflaredRoutesResultDataSourceModel struct {
	ID                 types.String      `tfsdk:"id" json:"id,computed"`
	Comment            types.String      `tfsdk:"comment" json:"comment,computed"`
	CreatedAt          timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DeletedAt          timetypes.RFC3339 `tfsdk:"deleted_at" json:"deleted_at,computed" format:"date-time"`
	Network            types.String      `tfsdk:"network" json:"network,computed"`
	TunType            types.String      `tfsdk:"tun_type" json:"tun_type,computed"`
	TunnelID           types.String      `tfsdk:"tunnel_id" json:"tunnel_id,computed"`
	TunnelName         types.String      `tfsdk:"tunnel_name" json:"tunnel_name,computed"`
	VirtualNetworkID   types.String      `tfsdk:"virtual_network_id" json:"virtual_network_id,computed"`
	VirtualNetworkName types.String      `tfsdk:"virtual_network_name" json:"virtual_network_name,computed"`
}
