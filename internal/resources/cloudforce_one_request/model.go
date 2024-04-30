// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_request

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CloudforceOneRequestResultEnvelope struct {
	Result CloudforceOneRequestModel `json:"result,computed"`
}

type CloudforceOneRequestModel struct {
	ID                types.String `tfsdk:"id" json:"id,computed"`
	AccountIdentifier types.String `tfsdk:"account_identifier" path:"account_identifier"`
	Content           types.String `tfsdk:"content" json:"content"`
	Priority          types.String `tfsdk:"priority" json:"priority"`
	RequestType       types.String `tfsdk:"request_type" json:"request_type"`
	Summary           types.String `tfsdk:"summary" json:"summary"`
	Tlp               types.String `tfsdk:"tlp" json:"tlp"`
}
