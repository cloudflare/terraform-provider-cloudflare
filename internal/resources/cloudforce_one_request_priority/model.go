// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_request_priority

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CloudforceOneRequestPriorityResultEnvelope struct {
	Result CloudforceOneRequestPriorityModel `json:"result,computed"`
}

type CloudforceOneRequestPriorityModel struct {
	ID                types.String    `tfsdk:"id" json:"id,computed"`
	AccountIdentifier types.String    `tfsdk:"account_identifier" path:"account_identifier"`
	Labels            *[]types.String `tfsdk:"labels" json:"labels"`
	Priority          types.Int64     `tfsdk:"priority" json:"priority"`
	Requirement       types.String    `tfsdk:"requirement" json:"requirement"`
	Tlp               types.String    `tfsdk:"tlp" json:"tlp"`
}
