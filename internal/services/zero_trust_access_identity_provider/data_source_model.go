// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_identity_provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessIdentityProviderResultDataSourceEnvelope struct {
	Result ZeroTrustAccessIdentityProviderDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessIdentityProviderResultListDataSourceEnvelope struct {
	Result *[]*ZeroTrustAccessIdentityProviderDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessIdentityProviderDataSourceModel struct {
	AccountID          types.String `tfsdk:"account_id" path:"account_id"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" path:"identity_provider_id"`
	ZoneID             types.String `tfsdk:"zone_id" path:"zone_id"`
}
