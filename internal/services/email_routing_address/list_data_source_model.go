// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_address

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/email_routing"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailRoutingAddressesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[EmailRoutingAddressesResultDataSourceModel] `json:"result,computed"`
}

type EmailRoutingAddressesDataSourceModel struct {
	AccountIdentifier types.String                                                             `tfsdk:"account_identifier" path:"account_identifier"`
	Direction         types.String                                                             `tfsdk:"direction" query:"direction,computed_optional"`
	Verified          types.Bool                                                               `tfsdk:"verified" query:"verified,computed_optional"`
	MaxItems          types.Int64                                                              `tfsdk:"max_items"`
	Result            customfield.NestedObjectList[EmailRoutingAddressesResultDataSourceModel] `tfsdk:"result"`
}

func (m *EmailRoutingAddressesDataSourceModel) toListParams() (params email_routing.AddressListParams, diags diag.Diagnostics) {
	params = email_routing.AddressListParams{}

	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(email_routing.AddressListParamsDirection(m.Direction.ValueString()))
	}
	if !m.Verified.IsNull() {
		params.Verified = cloudflare.F(email_routing.AddressListParamsVerified(m.Verified.ValueBool()))
	}

	return
}

type EmailRoutingAddressesResultDataSourceModel struct {
	ID       types.String      `tfsdk:"id" json:"id,computed"`
	Created  timetypes.RFC3339 `tfsdk:"created" json:"created,computed" format:"date-time"`
	Email    types.String      `tfsdk:"email" json:"email,computed"`
	Modified timetypes.RFC3339 `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	Tag      types.String      `tfsdk:"tag" json:"tag,computed"`
	Verified timetypes.RFC3339 `tfsdk:"verified" json:"verified,computed" format:"date-time"`
}
