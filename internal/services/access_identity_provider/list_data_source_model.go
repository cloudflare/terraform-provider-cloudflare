// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_identity_provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessIdentityProvidersResultListDataSourceEnvelope struct {
	Result *[]*AccessIdentityProvidersResultDataSourceModel `json:"result,computed"`
}

type AccessIdentityProvidersDataSourceModel struct {
	AccountID types.String                                     `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String                                     `tfsdk:"zone_id" path:"zone_id"`
	MaxItems  types.Int64                                      `tfsdk:"max_items"`
	Result    *[]*AccessIdentityProvidersResultDataSourceModel `tfsdk:"result"`
}

type AccessIdentityProvidersResultDataSourceModel struct {
}
