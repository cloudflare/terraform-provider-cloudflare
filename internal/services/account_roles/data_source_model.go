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

type AccountRolesResultDataSourceEnvelope struct {
	Result AccountRolesDataSourceModel `json:"result,computed"`
}

type AccountRolesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[AccountRolesDataSourceModel] `json:"result,computed"`
}

type AccountRolesDataSourceModel struct {
	AccountID   types.String                          `tfsdk:"account_id" path:"account_id,optional"`
	RoleID      types.String                          `tfsdk:"role_id" path:"role_id,optional"`
	Description types.String                          `tfsdk:"description" json:"description,optional"`
	ID          types.String                          `tfsdk:"id" json:"id,optional"`
	Name        types.String                          `tfsdk:"name" json:"name,optional"`
	Permissions *[]types.String                       `tfsdk:"permissions" json:"permissions,optional"`
	Filter      *AccountRolesFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *AccountRolesDataSourceModel) toReadParams(_ context.Context) (params accounts.RoleGetParams, diags diag.Diagnostics) {
	params = accounts.RoleGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *AccountRolesDataSourceModel) toListParams(_ context.Context) (params accounts.RoleListParams, diags diag.Diagnostics) {
	params = accounts.RoleListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type AccountRolesFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
}
