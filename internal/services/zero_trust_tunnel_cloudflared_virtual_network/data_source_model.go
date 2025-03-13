// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared_virtual_network

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/zero_trust"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustTunnelCloudflaredVirtualNetworkResultDataSourceEnvelope struct {
Result ZeroTrustTunnelCloudflaredVirtualNetworkDataSourceModel `json:"result,computed"`
}

type ZeroTrustTunnelCloudflaredVirtualNetworkDataSourceModel struct {
ID types.String `tfsdk:"id" json:"-,computed"`
VirtualNetworkID types.String `tfsdk:"virtual_network_id" path:"virtual_network_id,optional"`
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
Comment types.String `tfsdk:"comment" json:"comment,computed"`
CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
DeletedAt timetypes.RFC3339 `tfsdk:"deleted_at" json:"deleted_at,computed" format:"date-time"`
IsDefaultNetwork types.Bool `tfsdk:"is_default_network" json:"is_default_network,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
Filter *ZeroTrustTunnelCloudflaredVirtualNetworkFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *ZeroTrustTunnelCloudflaredVirtualNetworkDataSourceModel) toReadParams(_ context.Context) (params zero_trust.NetworkVirtualNetworkGetParams, diags diag.Diagnostics) {
  params = zero_trust.NetworkVirtualNetworkGetParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  return
}

func (m *ZeroTrustTunnelCloudflaredVirtualNetworkDataSourceModel) toListParams(_ context.Context) (params zero_trust.NetworkVirtualNetworkListParams, diags diag.Diagnostics) {
  params = zero_trust.NetworkVirtualNetworkListParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  if !m.Filter.ID.IsNull() {
    params.ID = cloudflare.F(m.Filter.ID.ValueString())
  }
  if !m.Filter.IsDefault.IsNull() {
    params.IsDefault = cloudflare.F(m.Filter.IsDefault.ValueBool())
  }
  if !m.Filter.IsDeleted.IsNull() {
    params.IsDeleted = cloudflare.F(m.Filter.IsDeleted.ValueBool())
  }
  if !m.Filter.Name.IsNull() {
    params.Name = cloudflare.F(m.Filter.Name.ValueString())
  }

  return
}

type ZeroTrustTunnelCloudflaredVirtualNetworkFindOneByDataSourceModel struct {
ID types.String `tfsdk:"id" query:"id,optional"`
IsDefault types.Bool `tfsdk:"is_default" query:"is_default,optional"`
IsDeleted types.Bool `tfsdk:"is_deleted" query:"is_deleted,optional"`
Name types.String `tfsdk:"name" query:"name,optional"`
}
