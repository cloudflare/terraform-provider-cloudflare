// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_member

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountMemberResultEnvelope struct {
	Result AccountMemberModel `json:"result,computed"`
}

type AccountMemberModel struct {
	ID        types.String    `tfsdk:"id" json:"id,computed"`
	AccountID types.String    `tfsdk:"account_id" path:"account_id"`
	Email     types.String    `tfsdk:"email" json:"email"`
	Roles     *[]types.String `tfsdk:"roles" json:"roles"`
	Status    types.String    `tfsdk:"status" json:"status"`
}
