// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_connectivity_settings

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustConnectivitySettingsResultEnvelope struct {
	Result ZeroTrustConnectivitySettingsModel `json:"result"`
}

type ZeroTrustConnectivitySettingsModel struct {
	ID                 types.String `tfsdk:"id" json:"-,computed"`
	AccountID          types.String `tfsdk:"account_id" path:"account_id,required"`
	IcmpProxyEnabled   types.Bool   `tfsdk:"icmp_proxy_enabled" json:"icmp_proxy_enabled,optional"`
	OfframpWARPEnabled types.Bool   `tfsdk:"offramp_warp_enabled" json:"offramp_warp_enabled,optional"`
}

func (m ZeroTrustConnectivitySettingsModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustConnectivitySettingsModel) MarshalJSONForUpdate(state ZeroTrustConnectivitySettingsModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
