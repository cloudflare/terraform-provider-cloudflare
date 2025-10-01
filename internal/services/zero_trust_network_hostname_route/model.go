// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_network_hostname_route

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustNetworkHostnameRouteResultEnvelope struct {
	Result ZeroTrustNetworkHostnameRouteModel `json:"result"`
}

type ZeroTrustNetworkHostnameRouteModel struct {
	ID         types.String      `tfsdk:"id" json:"id,computed"`
	AccountID  types.String      `tfsdk:"account_id" path:"account_id,required"`
	Comment    types.String      `tfsdk:"comment" json:"comment,optional"`
	Hostname   types.String      `tfsdk:"hostname" json:"hostname,optional"`
	TunnelID   types.String      `tfsdk:"tunnel_id" json:"tunnel_id,optional"`
	CreatedAt  timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DeletedAt  timetypes.RFC3339 `tfsdk:"deleted_at" json:"deleted_at,computed" format:"date-time"`
	TunnelName types.String      `tfsdk:"tunnel_name" json:"tunnel_name,computed"`
}

func (m ZeroTrustNetworkHostnameRouteModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustNetworkHostnameRouteModel) MarshalJSONForUpdate(state ZeroTrustNetworkHostnameRouteModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
