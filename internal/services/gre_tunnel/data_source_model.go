// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gre_tunnel

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type GRETunnelResultDataSourceEnvelope struct {
	Result GRETunnelDataSourceModel `json:"result,computed"`
}

type GRETunnelDataSourceModel struct {
	AccountID   types.String         `tfsdk:"account_id" path:"account_id"`
	GRETunnelID types.String         `tfsdk:"gre_tunnel_id" path:"gre_tunnel_id"`
	GRETunnel   jsontypes.Normalized `tfsdk:"gre_tunnel" json:"gre_tunnel"`
}
