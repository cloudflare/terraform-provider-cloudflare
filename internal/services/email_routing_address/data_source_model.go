// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_address

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
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
	Created                      timetypes.RFC3339                            `tfsdk:"created" json:"created,computed"`
	ID                           types.String                                 `tfsdk:"id" json:"id,computed"`
	Modified                     timetypes.RFC3339                            `tfsdk:"modified" json:"modified,computed"`
	Tag                          types.String                                 `tfsdk:"tag" json:"tag,computed"`
	Verified                     timetypes.RFC3339                            `tfsdk:"verified" json:"verified,computed"`
	Email                        types.String                                 `tfsdk:"email" json:"email"`
	Filter                       *EmailRoutingAddressFindOneByDataSourceModel `tfsdk:"filter"`
}

type EmailRoutingAddressFindOneByDataSourceModel struct {
	AccountIdentifier types.String `tfsdk:"account_identifier" path:"account_identifier"`
	Direction         types.String `tfsdk:"direction" query:"direction"`
	Verified          types.Bool   `tfsdk:"verified" query:"verified"`
}
