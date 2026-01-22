// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_api_token_permission_groups

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/accounts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountAPITokenPermissionGroupsResultDataSourceEnvelope struct {
	Result AccountAPITokenPermissionGroupsDataSourceModel `json:"result,computed"`
}

type AccountAPITokenPermissionGroupsDataSourceModel struct {
	AccountID        types.String                                                                                 `tfsdk:"account_id" path:"account_id,required"`
	Name             types.String                                                                                 `tfsdk:"name" query:"name,optional"`
	Scope            types.String                                                                                 `tfsdk:"scope" query:"scope,optional"`
	PermissionGroups customfield.NestedObjectList[AccountAPITokenPermissionGroupsPermissionGroupsDataSourceModel] `tfsdk:"permission_groups" json:"permission_groups,computed"`
}

func (m *AccountAPITokenPermissionGroupsDataSourceModel) toReadParams(_ context.Context) (params accounts.TokenPermissionGroupGetParams, diags diag.Diagnostics) {
	params = accounts.TokenPermissionGroupGetParams{
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

type AccountAPITokenPermissionGroupsPermissionGroupsDataSourceModel struct {
	ID     types.String                   `tfsdk:"id" json:"id,computed"`
	Name   types.String                   `tfsdk:"name" json:"name,computed"`
	Scopes customfield.List[types.String] `tfsdk:"scopes" json:"scopes,computed"`
}
