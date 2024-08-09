// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_application

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessApplicationResultDataSourceEnvelope struct {
	Result ZeroTrustAccessApplicationDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessApplicationResultListDataSourceEnvelope struct {
	Result *[]*ZeroTrustAccessApplicationDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessApplicationDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	AppID     types.String `tfsdk:"app_id" path:"app_id"`
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id"`
}
