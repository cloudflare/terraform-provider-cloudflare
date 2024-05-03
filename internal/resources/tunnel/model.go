// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tunnel

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TunnelResultEnvelope struct {
	Result TunnelModel `json:"result,computed"`
}

type TunnelModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	AccountID    types.String `tfsdk:"account_id" path:"account_id"`
	Name         types.String `tfsdk:"name" json:"name"`
	TunnelSecret types.String `tfsdk:"tunnel_secret" json:"tunnel_secret"`
}
