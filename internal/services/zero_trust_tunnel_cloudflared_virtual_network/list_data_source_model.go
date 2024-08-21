// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared_virtual_network

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustTunnelCloudflaredVirtualNetworksResultListDataSourceEnvelope struct {
	Result *[]*ZeroTrustTunnelCloudflaredVirtualNetworksResultDataSourceModel `json:"result,computed"`
}

type ZeroTrustTunnelCloudflaredVirtualNetworksDataSourceModel struct {
	AccountID types.String                                                       `tfsdk:"account_id" path:"account_id"`
	ID        types.String                                                       `tfsdk:"id" query:"id"`
	IsDefault types.Bool                                                         `tfsdk:"is_default" query:"is_default"`
	IsDeleted types.Bool                                                         `tfsdk:"is_deleted" query:"is_deleted"`
	Name      types.String                                                       `tfsdk:"name" query:"name"`
	MaxItems  types.Int64                                                        `tfsdk:"max_items"`
	Result    *[]*ZeroTrustTunnelCloudflaredVirtualNetworksResultDataSourceModel `tfsdk:"result"`
}

func (m *ZeroTrustTunnelCloudflaredVirtualNetworksDataSourceModel) toListParams() (params zero_trust.NetworkVirtualNetworkListParams, diags diag.Diagnostics) {
	params = zero_trust.NetworkVirtualNetworkListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.ID.IsNull() {
		params.ID = cloudflare.F(m.ID.ValueString())
	}
	if !m.IsDefault.IsNull() {
		params.IsDefault = cloudflare.F(m.IsDefault.ValueBool())
	}
	if !m.IsDeleted.IsNull() {
		params.IsDeleted = cloudflare.F(m.IsDeleted.ValueBool())
	}
	if !m.Name.IsNull() {
		params.Name = cloudflare.F(m.Name.ValueString())
	}

	return
}

type ZeroTrustTunnelCloudflaredVirtualNetworksResultDataSourceModel struct {
	ID               types.String      `tfsdk:"id" json:"id,computed"`
	Comment          types.String      `tfsdk:"comment" json:"comment,computed"`
	CreatedAt        timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed"`
	IsDefaultNetwork types.Bool        `tfsdk:"is_default_network" json:"is_default_network,computed"`
	Name             types.String      `tfsdk:"name" json:"name,computed"`
	DeletedAt        timetypes.RFC3339 `tfsdk:"deleted_at" json:"deleted_at,computed"`
}
