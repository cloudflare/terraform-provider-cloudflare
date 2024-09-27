// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_address

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/email_routing"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailRoutingAddressResultDataSourceEnvelope struct {
	Result EmailRoutingAddressDataSourceModel `json:"result,computed"`
}

type EmailRoutingAddressResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[EmailRoutingAddressDataSourceModel] `json:"result,computed"`
}

type EmailRoutingAddressDataSourceModel struct {
	AccountID                    types.String                                 `tfsdk:"account_id" path:"account_id,optional"`
	DestinationAddressIdentifier types.String                                 `tfsdk:"destination_address_identifier" path:"destination_address_identifier,optional"`
	Created                      timetypes.RFC3339                            `tfsdk:"created" json:"created,computed" format:"date-time"`
	Email                        types.String                                 `tfsdk:"email" json:"email,computed"`
	ID                           types.String                                 `tfsdk:"id" json:"id,computed"`
	Modified                     timetypes.RFC3339                            `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	Tag                          types.String                                 `tfsdk:"tag" json:"tag,computed"`
	Verified                     timetypes.RFC3339                            `tfsdk:"verified" json:"verified,computed" format:"date-time"`
	Filter                       *EmailRoutingAddressFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *EmailRoutingAddressDataSourceModel) toReadParams(_ context.Context) (params email_routing.AddressGetParams, diags diag.Diagnostics) {
	params = email_routing.AddressGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *EmailRoutingAddressDataSourceModel) toListParams(_ context.Context) (params email_routing.AddressListParams, diags diag.Diagnostics) {
	params = email_routing.AddressListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(email_routing.AddressListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.Verified.IsNull() {
		params.Verified = cloudflare.F(email_routing.AddressListParamsVerified(m.Filter.Verified.ValueBool()))
	}

	return
}

type EmailRoutingAddressFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
	Direction types.String `tfsdk:"direction" query:"direction,computed_optional"`
	Verified  types.Bool   `tfsdk:"verified" query:"verified,computed_optional"`
}
