// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_address

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailRoutingAddressResultDataSourceEnvelope struct {
	Result EmailRoutingAddressDataSourceModel `json:"result,computed"`
}

type EmailRoutingAddressResultListDataSourceEnvelope struct {
	Result *[]*EmailRoutingAddressDataSourceModel `json:"result,computed"`
}

type EmailRoutingAddressDataSourceModel struct {
	AccountIdentifier            types.String                                 `tfsdk:"account_identifier" path:"account_identifier"`
	DestinationAddressIdentifier types.String                                 `tfsdk:"destination_address_identifier" path:"destination_address_identifier"`
	ID                           types.String                                 `tfsdk:"id" json:"id,computed"`
	Created                      types.String                                 `tfsdk:"created" json:"created,computed"`
	Email                        types.String                                 `tfsdk:"email" json:"email"`
	Modified                     types.String                                 `tfsdk:"modified" json:"modified,computed"`
	Tag                          types.String                                 `tfsdk:"tag" json:"tag,computed"`
	Verified                     types.String                                 `tfsdk:"verified" json:"verified,computed"`
	FindOneBy                    *EmailRoutingAddressFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

type EmailRoutingAddressFindOneByDataSourceModel struct {
	AccountIdentifier types.String  `tfsdk:"account_identifier" path:"account_identifier"`
	Direction         types.String  `tfsdk:"direction" query:"direction"`
	Page              types.Float64 `tfsdk:"page" query:"page"`
	PerPage           types.Float64 `tfsdk:"per_page" query:"per_page"`
	Verified          types.Bool    `tfsdk:"verified" query:"verified"`
}
