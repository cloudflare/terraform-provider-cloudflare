// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_identity_provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessIdentityProvidersResultListDataSourceEnvelope struct {
	Result *[]*ZeroTrustAccessIdentityProvidersResultDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessIdentityProvidersDataSourceModel struct {
	AccountID types.String                                              `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String                                              `tfsdk:"zone_id" path:"zone_id"`
	MaxItems  types.Int64                                               `tfsdk:"max_items"`
	Result    *[]*ZeroTrustAccessIdentityProvidersResultDataSourceModel `tfsdk:"result"`
}

type ZeroTrustAccessIdentityProvidersResultDataSourceModel struct {
}
