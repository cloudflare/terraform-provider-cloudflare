// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_application

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessApplicationsResultListDataSourceEnvelope struct {
	Result *[]*AccessApplicationsItemsDataSourceModel `json:"result,computed"`
}

type AccessApplicationsDataSourceModel struct {
	AccountID types.String                               `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String                               `tfsdk:"zone_id" path:"zone_id"`
	MaxItems  types.Int64                                `tfsdk:"max_items"`
	Items     *[]*AccessApplicationsItemsDataSourceModel `tfsdk:"items"`
}

type AccessApplicationsItemsDataSourceModel struct {
}
