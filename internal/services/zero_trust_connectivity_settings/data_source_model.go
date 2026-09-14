// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_connectivity_settings

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustConnectivitySettingsResultDataSourceEnvelope struct {
	Result ZeroTrustConnectivitySettingsDataSourceModel `json:"result,computed"`
}

type ZeroTrustConnectivitySettingsDataSourceModel struct {
	ID                 types.String `tfsdk:"id" path:"account_id,computed"`
	AccountID          types.String `tfsdk:"account_id" path:"account_id,required"`
	IcmpProxyEnabled   types.Bool   `tfsdk:"icmp_proxy_enabled" json:"icmp_proxy_enabled,computed"`
	OfframpWARPEnabled types.Bool   `tfsdk:"offramp_warp_enabled" json:"offramp_warp_enabled,computed"`
}

func (m *ZeroTrustConnectivitySettingsDataSourceModel) toReadParams(_ context.Context) (params zero_trust.ConnectivitySettingGetParams, diags diag.Diagnostics) {
	params = zero_trust.ConnectivitySettingGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
