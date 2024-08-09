// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_key_configuration

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessKeyConfigurationResultDataSourceEnvelope struct {
	Result ZeroTrustAccessKeyConfigurationDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessKeyConfigurationDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
