// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_mtls_hostname_settings

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessMTLSHostnameSettingsResultDataSourceEnvelope struct {
	Result ZeroTrustAccessMTLSHostnameSettingsDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessMTLSHostnameSettingsDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id"`
}

func (m *ZeroTrustAccessMTLSHostnameSettingsDataSourceModel) toReadParams() (params zero_trust.AccessCertificateSettingGetParams, diags diag.Diagnostics) {
	params = zero_trust.AccessCertificateSettingGetParams{}

	if !m.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
	}

	return
}
