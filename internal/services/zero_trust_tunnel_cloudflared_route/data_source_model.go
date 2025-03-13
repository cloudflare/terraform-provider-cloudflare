// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared_route

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/zero_trust"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustTunnelCloudflaredRouteResultDataSourceEnvelope struct {
Result ZeroTrustTunnelCloudflaredRouteDataSourceModel `json:"result,computed"`
}

type ZeroTrustTunnelCloudflaredRouteDataSourceModel struct {
ID types.String `tfsdk:"id" json:"-,computed"`
RouteID types.String `tfsdk:"route_id" path:"route_id,optional"`
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
Comment types.String `tfsdk:"comment" json:"comment,computed"`
CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
DeletedAt timetypes.RFC3339 `tfsdk:"deleted_at" json:"deleted_at,computed" format:"date-time"`
Network types.String `tfsdk:"network" json:"network,computed"`
TunnelID types.String `tfsdk:"tunnel_id" json:"tunnel_id,computed"`
VirtualNetworkID types.String `tfsdk:"virtual_network_id" json:"virtual_network_id,computed"`
Filter *ZeroTrustTunnelCloudflaredRouteFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *ZeroTrustTunnelCloudflaredRouteDataSourceModel) toReadParams(_ context.Context) (params zero_trust.NetworkRouteGetParams, diags diag.Diagnostics) {
  params = zero_trust.NetworkRouteGetParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  return
}

func (m *ZeroTrustTunnelCloudflaredRouteDataSourceModel) toListParams(_ context.Context) (params zero_trust.NetworkRouteListParams, diags diag.Diagnostics) {
  mFilterTunTypes := []zero_trust.NetworkRouteListParamsTunType{}
  for _, item := range *m.Filter.TunTypes {
    mFilterTunTypes = append(mFilterTunTypes, zero_trust.NetworkRouteListParamsTunType(item.ValueString()))
  }

  params = zero_trust.NetworkRouteListParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
    TunTypes: cloudflare.F(mFilterTunTypes),
  }

  if !m.Filter.Comment.IsNull() {
    params.Comment = cloudflare.F(m.Filter.Comment.ValueString())
  }
  if !m.Filter.ExistedAt.IsNull() {
    params.ExistedAt = cloudflare.F(m.Filter.ExistedAt.ValueString())
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
  if !m.Filter.TunnelID.IsNull() {
    params.TunnelID = cloudflare.F(m.Filter.TunnelID.ValueString())
  }
  if !m.Filter.VirtualNetworkID.IsNull() {
    params.VirtualNetworkID = cloudflare.F(m.Filter.VirtualNetworkID.ValueString())
  }

  return
}

type ZeroTrustTunnelCloudflaredRouteFindOneByDataSourceModel struct {
Comment types.String `tfsdk:"comment" query:"comment,optional"`
ExistedAt types.String `tfsdk:"existed_at" query:"existed_at,optional"`
IsDeleted types.Bool `tfsdk:"is_deleted" query:"is_deleted,optional"`
NetworkSubset types.String `tfsdk:"network_subset" query:"network_subset,optional"`
NetworkSuperset types.String `tfsdk:"network_superset" query:"network_superset,optional"`
RouteID types.String `tfsdk:"route_id" query:"route_id,optional"`
TunTypes *[]types.String `tfsdk:"tun_types" query:"tun_types,optional"`
TunnelID types.String `tfsdk:"tunnel_id" query:"tunnel_id,optional"`
VirtualNetworkID types.String `tfsdk:"virtual_network_id" query:"virtual_network_id,optional"`
}
