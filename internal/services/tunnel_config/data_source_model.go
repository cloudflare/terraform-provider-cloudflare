// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tunnel_config

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TunnelConfigResultDataSourceEnvelope struct {
	Result TunnelConfigDataSourceModel `json:"result,computed"`
}

type TunnelConfigDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	TunnelID  types.String `tfsdk:"tunnel_id" path:"tunnel_id"`
}
