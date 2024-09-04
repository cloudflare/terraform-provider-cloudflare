// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ListResultEnvelope struct {
	Result ListModel `json:"result"`
}

type ListModel struct {
	AccountID   types.String `tfsdk:"account_id" path:"account_id,required"`
	ListID      types.String `tfsdk:"list_id" path:"list_id,optional"`
	Kind        types.String `tfsdk:"kind" json:"kind,required"`
	Name        types.String `tfsdk:"name" json:"name,required"`
	Description types.String `tfsdk:"description" json:"description,optional"`
	ID          types.String `tfsdk:"id" json:"id,computed"`
}
