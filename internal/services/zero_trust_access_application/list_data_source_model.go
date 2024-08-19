// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_application

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessApplicationsResultListDataSourceEnvelope struct {
	Result *[]*ZeroTrustAccessApplicationsResultDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessApplicationsDataSourceModel struct {
	AccountID types.String                                         `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String                                         `tfsdk:"zone_id" path:"zone_id"`
	MaxItems  types.Int64                                          `tfsdk:"max_items"`
	Result    *[]*ZeroTrustAccessApplicationsResultDataSourceModel `tfsdk:"result"`
}

type ZeroTrustAccessApplicationsResultDataSourceModel struct {
}
