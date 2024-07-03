// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_identity_provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessIdentityProviderResultDataSourceEnvelope struct {
	Result AccessIdentityProviderDataSourceModel `json:"result,computed"`
}

type AccessIdentityProviderResultListDataSourceEnvelope struct {
	Result *[]*AccessIdentityProviderDataSourceModel `json:"result,computed"`
}

type AccessIdentityProviderDataSourceModel struct {
	IdentityProviderID types.String                                    `tfsdk:"identity_provider_id" path:"identity_provider_id"`
	AccountID          types.String                                    `tfsdk:"account_id" path:"account_id"`
	ZoneID             types.String                                    `tfsdk:"zone_id" path:"zone_id"`
	FindOneBy          *AccessIdentityProviderFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

type AccessIdentityProviderFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id"`
}
