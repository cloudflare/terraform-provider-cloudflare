// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared_virtual_network

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustTunnelCloudflaredVirtualNetworkResultListDataSourceEnvelope struct {
	Result *[]*ZeroTrustTunnelCloudflaredVirtualNetworkDataSourceModel `json:"result,computed"`
}

type ZeroTrustTunnelCloudflaredVirtualNetworkDataSourceModel struct {
	Comment          types.String                                                      `tfsdk:"comment" json:"comment"`
	CreatedAt        timetypes.RFC3339                                                 `tfsdk:"created_at" json:"created_at"`
	DeletedAt        timetypes.RFC3339                                                 `tfsdk:"deleted_at" json:"deleted_at"`
	ID               types.String                                                      `tfsdk:"id" json:"id"`
	IsDefaultNetwork types.Bool                                                        `tfsdk:"is_default_network" json:"is_default_network"`
	Name             types.String                                                      `tfsdk:"name" json:"name"`
	Filter           *ZeroTrustTunnelCloudflaredVirtualNetworkFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *ZeroTrustTunnelCloudflaredVirtualNetworkDataSourceModel) toListParams() (params zero_trust.NetworkVirtualNetworkListParams, diags diag.Diagnostics) {
	params = zero_trust.NetworkVirtualNetworkListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
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
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	ID        types.String `tfsdk:"id" query:"id"`
	IsDefault types.Bool   `tfsdk:"is_default" query:"is_default"`
	IsDeleted types.Bool   `tfsdk:"is_deleted" query:"is_deleted"`
	Name      types.String `tfsdk:"name" query:"name"`
}
