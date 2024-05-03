// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_address

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailRoutingAddressResultEnvelope struct {
	Result EmailRoutingAddressModel `json:"result,computed"`
}

type EmailRoutingAddressModel struct {
	ID                types.String `tfsdk:"id" json:"id,computed"`
	AccountIdentifier types.String `tfsdk:"account_identifier" path:"account_identifier"`
	Email             types.String `tfsdk:"email" json:"email"`
}
