// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_token_permissions_groups

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/accounts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APITokenPermissionsGroupsListResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[APITokenPermissionsGroupsListResultDataSourceModel] `json:"result,computed"`
}

type APITokenPermissionsGroupsListDataSourceModel struct {
	AccountID types.String                                                                     `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                                      `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[APITokenPermissionsGroupsListResultDataSourceModel] `tfsdk:"result"`
}

func (m *APITokenPermissionsGroupsListDataSourceModel) toListParams(_ context.Context) (params accounts.TokenPermissionGroupListParams, diags diag.Diagnostics) {
	params = accounts.TokenPermissionGroupListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type APITokenPermissionsGroupsListResultDataSourceModel struct {
}
