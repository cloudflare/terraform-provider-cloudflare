// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_application

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessApplicationResultDataSourceEnvelope struct {
	Result AccessApplicationDataSourceModel `json:"result,computed"`
}

type AccessApplicationResultListDataSourceEnvelope struct {
	Result *[]*AccessApplicationDataSourceModel `json:"result,computed"`
}

type AccessApplicationDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	AppID     types.String `tfsdk:"app_id" path:"app_id"`
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id"`
}
