// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package argo_tunnel

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ArgoTunnelResultEnvelope struct {
	Result ArgoTunnelModel `json:"result"`
}

type ArgoTunnelModel struct {
	ID          types.String `tfsdk:"id" json:"id,computed"`
	AccountID   types.String `tfsdk:"account_id" path:"account_id,required"`
	Name        types.String `tfsdk:"name" json:"name,required"`
	Secret      types.String `tfsdk:"secret" json:"secret,required"`
	Cname       types.String `tfsdk:"cname" json:"cname,computed"`
	TunnelToken types.String `tfsdk:"tunnel_token" json:"tunnel_token,computed"`
}

func (m ArgoTunnelModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ArgoTunnelModel) MarshalJSONForUpdate(state ArgoTunnelModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
