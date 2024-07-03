// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_address

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailRoutingAddressesResultListDataSourceEnvelope struct {
	Result *[]*EmailRoutingAddressesItemsDataSourceModel `json:"result,computed"`
}

type EmailRoutingAddressesDataSourceModel struct {
	AccountIdentifier types.String                                  `tfsdk:"account_identifier" path:"account_identifier"`
	Direction         types.String                                  `tfsdk:"direction" query:"direction"`
	Page              types.Float64                                 `tfsdk:"page" query:"page"`
	PerPage           types.Float64                                 `tfsdk:"per_page" query:"per_page"`
	Verified          types.Bool                                    `tfsdk:"verified" query:"verified"`
	MaxItems          types.Int64                                   `tfsdk:"max_items"`
	Items             *[]*EmailRoutingAddressesItemsDataSourceModel `tfsdk:"items"`
}

type EmailRoutingAddressesItemsDataSourceModel struct {
	ID       types.String `tfsdk:"id" json:"id,computed"`
	Created  types.String `tfsdk:"created" json:"created,computed"`
	Email    types.String `tfsdk:"email" json:"email,computed"`
	Modified types.String `tfsdk:"modified" json:"modified,computed"`
	Tag      types.String `tfsdk:"tag" json:"tag,computed"`
	Verified types.String `tfsdk:"verified" json:"verified,computed"`
}
