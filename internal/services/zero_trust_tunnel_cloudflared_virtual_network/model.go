// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared_virtual_network

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type ZeroTrustTunnelCloudflaredVirtualNetworkResultEnvelope struct {
	Result ZeroTrustTunnelCloudflaredVirtualNetworkModel `json:"result"`
}

type ZeroTrustTunnelCloudflaredVirtualNetworkModel struct {
	ID               types.String      `tfsdk:"id" json:"id,computed"`
	AccountID        types.String      `tfsdk:"account_id" path:"account_id,required"`
	Name             types.String      `tfsdk:"name" json:"name,required"`
	Comment          types.String      `tfsdk:"comment" json:"comment,optional"`
	IsDefaultNetwork types.Bool        `tfsdk:"is_default_network" json:"is_default_network,computed_optional"`
	CreatedAt        timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DeletedAt        timetypes.RFC3339 `tfsdk:"deleted_at" json:"deleted_at,computed" format:"date-time"`
}

type newZeroTrustTunnelCloudflaredVirtualNetworkModel struct {
	Name      types.String `json:"name,required"`
	Comment   types.String `json:"comment,optional"`
	IsDefault types.Bool   `json:"is_default,required"`
}

func (m ZeroTrustTunnelCloudflaredVirtualNetworkModel) MarshalJSON() (data []byte, err error) {
	body := &newZeroTrustTunnelCloudflaredVirtualNetworkModel{
		Name:      m.Name,
		Comment:   m.Comment,
		IsDefault: basetypes.NewBoolValue(m.IsDefaultNetwork.ValueBool()),
	}
	return apijson.MarshalRoot(body)
}

func (m ZeroTrustTunnelCloudflaredVirtualNetworkModel) MarshalJSONForUpdate(state ZeroTrustTunnelCloudflaredVirtualNetworkModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
