// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_roles

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/accounts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountRolesListResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[AccountRolesListResultDataSourceModel] `json:"result,computed"`
}

type AccountRolesListDataSourceModel struct {
	AccountID types.String                                                        `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                         `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[AccountRolesListResultDataSourceModel] `tfsdk:"result"`
}

func (m *AccountRolesListDataSourceModel) toListParams(_ context.Context) (params accounts.RoleListParams, diags diag.Diagnostics) {
	params = accounts.RoleListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type AccountRolesListResultDataSourceModel struct {
	ID          types.String                   `tfsdk:"id" json:"id,computed"`
	Description types.String                   `tfsdk:"description" json:"description,computed"`
	Name        types.String                   `tfsdk:"name" json:"name,computed"`
	Permissions customfield.List[types.String] `tfsdk:"permissions" json:"permissions,computed"`
}
