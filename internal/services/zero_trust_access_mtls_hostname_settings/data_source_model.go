// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_mtls_hostname_settings

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessMTLSHostnameSettingsResultDataSourceEnvelope struct {
	Result ZeroTrustAccessMTLSHostnameSettingsDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessMTLSHostnameSettingsDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id"`
}
