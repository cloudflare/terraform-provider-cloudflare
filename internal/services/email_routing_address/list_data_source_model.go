// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_address

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailRoutingAddressesResultListDataSourceEnvelope struct {
	Result *[]*EmailRoutingAddressesResultDataSourceModel `json:"result,computed"`
}

type EmailRoutingAddressesDataSourceModel struct {
	AccountIdentifier types.String                                   `tfsdk:"account_identifier" path:"account_identifier"`
	Direction         types.String                                   `tfsdk:"direction" query:"direction"`
	Verified          types.Bool                                     `tfsdk:"verified" query:"verified"`
	MaxItems          types.Int64                                    `tfsdk:"max_items"`
	Result            *[]*EmailRoutingAddressesResultDataSourceModel `tfsdk:"result"`
}

type EmailRoutingAddressesResultDataSourceModel struct {
	ID       types.String      `tfsdk:"id" json:"id,computed"`
	Created  timetypes.RFC3339 `tfsdk:"created" json:"created,computed"`
	Email    types.String      `tfsdk:"email" json:"email"`
	Modified timetypes.RFC3339 `tfsdk:"modified" json:"modified,computed"`
	Tag      types.String      `tfsdk:"tag" json:"tag,computed"`
	Verified timetypes.RFC3339 `tfsdk:"verified" json:"verified,computed"`
}
