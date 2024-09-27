// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_role

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/accounts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountRoleResultDataSourceEnvelope struct {
	Result AccountRoleDataSourceModel `json:"result,computed"`
}

type AccountRoleResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[AccountRoleDataSourceModel] `json:"result,computed"`
}

type AccountRoleDataSourceModel struct {
	AccountID   types.String                         `tfsdk:"account_id" path:"account_id,optional"`
	RoleID      types.String                         `tfsdk:"role_id" path:"role_id,optional"`
	Description types.String                         `tfsdk:"description" json:"description,optional"`
	ID          types.String                         `tfsdk:"id" json:"id,optional"`
	Name        types.String                         `tfsdk:"name" json:"name,optional"`
	Permissions *[]types.String                      `tfsdk:"permissions" json:"permissions,optional"`
	Filter      *AccountRoleFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *AccountRoleDataSourceModel) toReadParams(_ context.Context) (params accounts.RoleGetParams, diags diag.Diagnostics) {
	params = accounts.RoleGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *AccountRoleDataSourceModel) toListParams(_ context.Context) (params accounts.RoleListParams, diags diag.Diagnostics) {
	params = accounts.RoleListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type AccountRoleFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
}
