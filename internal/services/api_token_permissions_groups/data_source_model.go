// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_token_permissions_groups

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/accounts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APITokenPermissionsGroupsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[APITokenPermissionsGroupsDataSourceModel] `json:"result,computed"`
}

type APITokenPermissionsGroupsDataSourceModel struct {
	Filter *APITokenPermissionsGroupsFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *APITokenPermissionsGroupsDataSourceModel) toListParams(_ context.Context) (params accounts.TokenPermissionGroupListParams, diags diag.Diagnostics) {
	params = accounts.TokenPermissionGroupListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type APITokenPermissionsGroupsFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
}
