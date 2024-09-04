// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_address

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailRoutingAddressResultEnvelope struct {
	Result EmailRoutingAddressModel `json:"result"`
}

type EmailRoutingAddressModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	AccountID types.String      `tfsdk:"account_id" path:"account_id,required"`
	Email     types.String      `tfsdk:"email" json:"email,required"`
	Created   timetypes.RFC3339 `tfsdk:"created" json:"created,computed" format:"date-time"`
	Modified  timetypes.RFC3339 `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	Tag       types.String      `tfsdk:"tag" json:"tag,computed"`
	Verified  timetypes.RFC3339 `tfsdk:"verified" json:"verified,computed" format:"date-time"`
}
