// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_api_token_permission_groups

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/accounts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountAPITokenPermissionGroupsListResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[AccountAPITokenPermissionGroupsListResultDataSourceModel] `json:"result,computed"`
}

type AccountAPITokenPermissionGroupsListDataSourceModel struct {
	AccountID types.String                                                                           `tfsdk:"account_id" path:"account_id,required"`
	Name      types.String                                                                           `tfsdk:"name" query:"name,optional"`
	Scope     types.String                                                                           `tfsdk:"scope" query:"scope,optional"`
	MaxItems  types.Int64                                                                            `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[AccountAPITokenPermissionGroupsListResultDataSourceModel] `tfsdk:"result"`
}

func (m *AccountAPITokenPermissionGroupsListDataSourceModel) toListParams(_ context.Context) (params accounts.TokenPermissionGroupListParams, diags diag.Diagnostics) {
	params = accounts.TokenPermissionGroupListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Name.IsNull() {
		params.Name = cloudflare.F(m.Name.ValueString())
	}
	if !m.Scope.IsNull() {
		params.Scope = cloudflare.F(m.Scope.ValueString())
	}

	return
}

type AccountAPITokenPermissionGroupsListResultDataSourceModel struct {
	ID     types.String                   `tfsdk:"id" json:"id,computed"`
	Name   types.String                   `tfsdk:"name" json:"name,computed"`
	Scopes customfield.List[types.String] `tfsdk:"scopes" json:"scopes,computed"`
}
