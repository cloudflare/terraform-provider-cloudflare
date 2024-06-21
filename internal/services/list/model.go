// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ListResultEnvelope struct {
	Result ListModel `json:"result,computed"`
}

type ListModel struct {
	AccountID   types.String `tfsdk:"account_id" path:"account_id"`
	ListID      types.String `tfsdk:"list_id" path:"list_id"`
	Kind        types.String `tfsdk:"kind" json:"kind"`
	Name        types.String `tfsdk:"name" json:"name"`
	Description types.String `tfsdk:"description" json:"description"`
}
