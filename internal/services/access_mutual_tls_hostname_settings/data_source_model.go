// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_mutual_tls_hostname_settings

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessMutualTLSHostnameSettingsResultDataSourceEnvelope struct {
	Result AccessMutualTLSHostnameSettingsDataSourceModel `json:"result,computed"`
}

type AccessMutualTLSHostnameSettingsDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id"`
}
