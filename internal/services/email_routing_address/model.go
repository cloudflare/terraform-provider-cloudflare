// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_address

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailRoutingAddressResultEnvelope struct {
	Result EmailRoutingAddressModel `json:"result,computed"`
}

type EmailRoutingAddressResultDataSourceEnvelope struct {
	Result EmailRoutingAddressDataSourceModel `json:"result,computed"`
}

type EmailRoutingAddressesResultDataSourceEnvelope struct {
	Result EmailRoutingAddressesDataSourceModel `json:"result,computed"`
}

type EmailRoutingAddressModel struct {
	ID                types.String `tfsdk:"id" json:"id"`
	AccountIdentifier types.String `tfsdk:"account_identifier" path:"account_identifier"`
	Email             types.String `tfsdk:"email" json:"email"`
	Created           types.String `tfsdk:"created" json:"created,computed"`
	Modified          types.String `tfsdk:"modified" json:"modified,computed"`
	Tag               types.String `tfsdk:"tag" json:"tag,computed"`
	Verified          types.String `tfsdk:"verified" json:"verified,computed"`
}

type EmailRoutingAddressDataSourceModel struct {
}

type EmailRoutingAddressesDataSourceModel struct {
}
