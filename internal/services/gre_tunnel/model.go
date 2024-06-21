// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gre_tunnel

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type GRETunnelResultEnvelope struct {
	Result GRETunnelModel `json:"result,computed"`
}

type GRETunnelModel struct {
	AccountID   types.String `tfsdk:"account_id" path:"account_id"`
	GRETunnelID types.String `tfsdk:"gre_tunnel_id" path:"gre_tunnel_id"`
}
